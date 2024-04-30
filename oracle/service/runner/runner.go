package runner

import (
	"context"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/cometbft/cometbft/oracle/service/types"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	cs "github.com/cometbft/cometbft/consensus"
	oracleproto "github.com/cometbft/cometbft/proto/tendermint/oracle"
)

func RunProcessSignVoteQueue(oracleInfo *types.OracleInfo, consensusState *cs.State) {
	go func(oracleInfo *types.OracleInfo) {
		for {
			select {
			case <-oracleInfo.StopChannel:
				return
			default:
				interval := oracleInfo.Config.SignInterval
				if interval == 0 {
					interval = 100 * time.Millisecond
				}
				time.Sleep(interval)
				ProcessSignVoteQueue(oracleInfo, consensusState)
			}
		}
	}(oracleInfo)
}

func ProcessSignVoteQueue(oracleInfo *types.OracleInfo, consensusState *cs.State) {
	votes := []*oracleproto.Vote{}
	for {
		select {
		case vote := <-oracleInfo.SignVotesChan:
			votes = append(votes, vote)
			continue
		default:
		}
		break
	}

	if len(votes) == 0 {
		return
	}

	// batch sign the new votes, along with existing votes in gossipVoteBuffer, if any
	validatorIndex, _ := consensusState.Validators.GetByAddress(oracleInfo.PubKey.Address())
	if validatorIndex == -1 {
		log.Errorf("unable to find validator index")
		return
	}

	// append new batch into unsignedVotesBuffer, need to mutex lock as it will clash with concurrent pruning
	oracleInfo.UnsignedVoteBuffer.UpdateMtx.Lock()
	oracleInfo.UnsignedVoteBuffer.Buffer = append(oracleInfo.UnsignedVoteBuffer.Buffer, votes...)

	// batch sign the entire unsignedVoteBuffer and add to gossipBuffer
	newGossipVote := &oracleproto.GossipVote{
		ValidatorIndex:  validatorIndex,
		SignedTimestamp: time.Now().Unix(),
		Votes:           oracleInfo.UnsignedVoteBuffer.Buffer,
	}

	oracleInfo.UnsignedVoteBuffer.UpdateMtx.Unlock()

	// sort the votes so that we can rebuild it in a deterministic order, when uncompressing
	sort.Sort(ByVote(newGossipVote.Votes))

	// signing of vote should append the signature field of gossipVote
	if err := oracleInfo.PrivValidator.SignOracleVote("", newGossipVote); err != nil {
		log.Errorf("error signing oracle votes")
		// unlock here to prevent deadlock
		oracleInfo.GossipVoteBuffer.UpdateMtx.Unlock()
		return
	}

	// need to mutex lock as it will clash with concurrent gossip
	oracleInfo.GossipVoteBuffer.UpdateMtx.Lock()
	address := oracleInfo.PubKey.Address().String()
	oracleInfo.GossipVoteBuffer.Buffer[address] = newGossipVote
	oracleInfo.GossipVoteBuffer.UpdateMtx.Unlock()
}

func reverseInts(input []*oracleproto.Vote) []*oracleproto.Vote {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func PruneUnsignedVoteBuffer(oracleInfo *types.OracleInfo, consensusState *cs.State) {
	go func(oracleInfo *types.OracleInfo) {
		maxGossipVoteAge := oracleInfo.Config.MaxGossipVoteAge
		if maxGossipVoteAge == 0 {
			maxGossipVoteAge = 2
		}
		ticker := time.Tick(1 * time.Second)
		for range ticker {
			lastBlockTime := consensusState.GetState().LastBlockTime

			if !contains(oracleInfo.BlockTimestamps, lastBlockTime.Unix()) {
				oracleInfo.BlockTimestamps = append(oracleInfo.BlockTimestamps, lastBlockTime.Unix())
			}

			if len(oracleInfo.BlockTimestamps) < maxGossipVoteAge {
				continue
			}

			if len(oracleInfo.BlockTimestamps) > maxGossipVoteAge {
				oracleInfo.BlockTimestamps = oracleInfo.BlockTimestamps[1:]
			}

			oracleInfo.UnsignedVoteBuffer.UpdateMtx.Lock()
			// prune votes that are older than the maxGossipVoteAge (in terms of block height)
			newVotes := []*oracleproto.Vote{}
			unsignedVoteBuffer := oracleInfo.UnsignedVoteBuffer.Buffer
			for _, vote := range unsignedVoteBuffer {
				if vote.Timestamp >= oracleInfo.BlockTimestamps[0] {
					newVotes = append(newVotes, vote)
				} else {
					log.Infof("deleting vote timestamp: %v, block timestamp: %v", vote.Timestamp, oracleInfo.BlockTimestamps[0])
				}
			}
			oracleInfo.UnsignedVoteBuffer.Buffer = newVotes
			oracleInfo.UnsignedVoteBuffer.UpdateMtx.Unlock()
		}
	}(oracleInfo)
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func PruneGossipVoteBuffer(oracleInfo *types.OracleInfo) {
	go func(oracleInfo *types.OracleInfo) {
		interval := 60 * time.Second
		ticker := time.Tick(interval)
		for range ticker {
			oracleInfo.GossipVoteBuffer.UpdateMtx.Lock()
			currTime := time.Now().Unix()
			buffer := oracleInfo.GossipVoteBuffer.Buffer

			// prune gossip vote that have signed timestamps older than 60 secs
			for valAddr, gossipVote := range oracleInfo.GossipVoteBuffer.Buffer {
				if gossipVote.SignedTimestamp < currTime-int64(interval.Seconds()) {
					log.Infof("DELETING STALE GOSSIP BUFFER (%v) FOR VAL: %s", gossipVote.SignedTimestamp, valAddr)
					delete(buffer, valAddr)
				}
			}
			oracleInfo.GossipVoteBuffer.Buffer = buffer
			oracleInfo.GossipVoteBuffer.UpdateMtx.Unlock()
		}
	}(oracleInfo)
}

// Run run oracles
func Run(oracleInfo *types.OracleInfo, consensusState *cs.State) {
	log.Info("[oracle] Service started.")
	RunProcessSignVoteQueue(oracleInfo, consensusState)
	PruneUnsignedVoteBuffer(oracleInfo, consensusState)
	PruneGossipVoteBuffer(oracleInfo)
	// start to take votes from app
	for {
		res, err := oracleInfo.ProxyApp.PrepareOracleVotes(context.Background(), &abcitypes.RequestPrepareOracleVotes{})
		if err != nil {
			log.Errorf("app not ready: %v, retrying...", err)
			time.Sleep(1 * time.Second)
			continue
		}

		oracleInfo.SignVotesChan <- res.Vote
	}
}

type ByVote []*oracleproto.Vote

func (b ByVote) Len() int {
	return len(b)
}

func (b ByVote) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Define custom sorting rules
func (b ByVote) Less(i, j int) bool {
	if b[i].OracleId != b[j].OracleId {
		return b[i].OracleId < b[j].OracleId
	}
	if b[i].Timestamp != b[j].Timestamp {
		return b[i].Timestamp < b[j].Timestamp
	}
	return b[i].Data < b[j].Data
}
