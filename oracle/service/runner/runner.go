package runner

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Switcheo/carbon-wallet-go/api"
	"github.com/Switcheo/carbon/constants"
	oracletypes "github.com/Switcheo/carbon/x/oracle/types"
	"github.com/cometbft/cometbft/oracle/service/adapters"
	"github.com/cometbft/cometbft/oracle/service/parser"
	"github.com/cometbft/cometbft/oracle/service/types"
	"github.com/cometbft/cometbft/redis"
)

var (
	OracleOverwriteData string
)

// OracleInfoResult oracle info result
type OracleInfoResult struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

// OracleInfoResponse oracle info response
type OracleInfoResponse struct {
	Height string           `json:"height"`
	Result OracleInfoResult `json:"result"`
}

// LastSubmissionTimeKey key for last submission time
const LastSubmissionTimeKey = "oracle:submitter:last-submission-time"

// LastStoreDataKey returns the key for the given adapter and job
func LastStoreDataKey(adapter types.Adapter, job types.OracleJob) string {
	return fmt.Sprintf("oracle:adapter-store:%s:%s", adapter.Id(), job.InputId)
}

// GetLastStoreData returns the last stored value for the given adapter and job
func GetLastStoreData(service redis.Service, adapter types.Adapter, job types.OracleJob) (data map[string]types.GenericValue, exists bool, err error) {
	key := LastStoreDataKey(adapter, job)
	value, exists, err := service.Get(key)
	if err != nil {
		return
	}
	data = make(map[string]types.GenericValue)
	if exists {
		err := json.Unmarshal([]byte(value.String()), &data)
		if err != nil {
			panic(err)
		}
	}
	return data, exists, nil
}

// SetLastStoreData sets the last store data for the given adapter and job
func SetLastStoreData(service redis.Service, adapter types.Adapter, job types.OracleJob, store types.AdapterStore) error {
	key := LastStoreDataKey(adapter, job)
	dataBytes, err := json.Marshal(&store.Data)
	if err != nil {
		panic(err)
	}
	err = service.Set(key, types.StringToGenericValue(string(dataBytes)), 0)
	if err != nil {
		return err
	}
	return nil
}

// GetOracleLockKey returns the lock key for a given oracle and time
func GetOracleLockKey(oracle types.Oracle, normalizedTime uint64) string {
	return fmt.Sprintf("oracle:oracle-lock:%s:%d", oracle.Id, normalizedTime)
}

func overwriteData(oracleId string, data string) string {
	if oracleId != "DXBT" { // if we want to overwrite DETH: `&& oracleID != "DETH"`
		return data
	}

	var min, max, interval int64
	switch oracleId {
	case "DXBT":
		min, max = 15000, 10000 // this was how it was before the refactor, maybe intended?
		interval = 20
	case "DETH":
		min, max = 500, 1500
		interval = 5
	}

	// create a price based on current system time
	t := time.Now().Unix()
	minute := t / 60
	seconds := t - (t/60)*60
	// round to the nearest 10th second, e.g. 10, 20, 30...
	roundedSeconds := (seconds / 10) * 10
	isEvenMinute := minute%2 == 0
	// if the minute is exactly an even minute
	if isEvenMinute {
		if roundedSeconds == 0 {
			return strconv.FormatUint(uint64(min), 10)
		}

		price := strconv.FormatUint(uint64(min+roundedSeconds*interval), 10)
		decimalPrice := strconv.FormatUint(uint64(seconds/10), 10)
		decimalPrice += strconv.FormatUint(10-uint64(seconds/10), 10)
		return price + "." + decimalPrice
	}

	if roundedSeconds == 0 {
		return strconv.FormatUint(uint64(max), 10)
	}

	price := strconv.FormatUint(uint64(max-roundedSeconds*interval), 10)
	decimalPrice := strconv.FormatUint(uint64(seconds/10)+4, 10)
	decimalPrice += strconv.FormatUint(10-uint64(seconds/10), 10)
	return price + "." + decimalPrice
}

// func msgResultCallback(response *sdktypes.TxResponse, msg sdktypes.Msg, err error) {
// 	var result float32
// 	result = 1
// 	if err != nil {
// 		result = 0
// 	}
// 	telemetry.SetGaugeWithLabels([]string{constants.METRIC_VOTE_STATUS}, result, []metrics.Label{telemetry.NewLabel("transaction_hash", response.TxHash)})
// }

// // SubmitVote submit oracle vote
// func SubmitVote(app types.App, msg oracletypes.MsgCreateVote) {
// 	power, err := oracleapi.GetSubAccountPower(app.Wallet.GRPCURL, app.Wallet.Bech32Addr, app.Wallet.ClientCtx)
// 	if err != nil {
// 		return
// 	}

// 	if power.IsZero() {
// 		return
// 	}

// 	app.Wallet.SubmitMsgAsync(&msg, msgResultCallback)
// 	// log.Infoln("Submitted vote", app.Wallet.AccAddress().String(), msg.OracleID, msg.Timestamp, msg.Data)
// }

// SyncOracles sync oracles with active on-chain oracles
func SyncOracles(app types.App) (oracles []types.Oracle, err error) {
	// fetch oracle list first
	grpcConn, err := api.GetGRPCConnection(app.Wallet.GRPCURL, app.Wallet.ClientCtx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer grpcConn.Close()

	oracleClient := oracletypes.NewQueryClient(grpcConn)
	oracleRes, err := oracleClient.OracleAll(
		context.Background(),
		&oracletypes.QueryAllOracleRequest{
			//Pagination: &sdkquerytypes.PageRequest{}
		},
	)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	oraclesData := oracleRes.Oracles

	for _, oracle := range oraclesData {
		var spec types.OracleSpec
		err = json.Unmarshal([]byte(oracle.Spec), &spec)
		if err != nil {
			log.Errorf("[oracle:%v] invalid oracle spec: %+v", oracle.Id, err)
			continue
		}
		spec, err := parser.ParseSpec(spec)
		if err != nil {
			log.Warnf("[oracle:%v] unable to unroll spec: %v", oracle.Id, err)
			continue
		}
		err = parser.ValidateOracleJobs(app, spec.Jobs)
		if err != nil {
			log.Warnf("[oracle: %v,] invalid oracle jobs: %v", oracle.Id, err)
			continue
		}
		oracles = append(oracles, types.Oracle{
			Id:         oracle.Id,
			Resolution: uint64(oracle.Resolution),
			Spec:       spec,
		})
	}
	log.Info("[oracle] Synced oracle specs")
	return oracles, err
}

func SaveOracleResult(price string, oracleId string, redisService redis.Service) {
	if price != "" {
		key := adapters.GetOracleResultKey(oracleId)
		data, err := json.Marshal(types.OracleCache{Price: price, Timestamp: types.JSONTime{Time: time.Now()}})
		if err != nil {
			panic(err)
		}
		jsonString := string(data)
		setErr := redisService.Set(key, types.StringToGenericValue(jsonString), 0)
		if setErr != nil {
			log.Error(err)
		}
	}
}

// RunOracle run oracle submission
func RunOracle(app types.App, oracle types.Oracle, currentTime uint64) error {
	red := app.Redis
	normalizedTime := (currentTime / oracle.Resolution) * oracle.Resolution
	lastSubmissionTime, exists, err := red.Get(LastSubmissionTimeKey)
	if err != nil {
		return err
	}
	if exists && normalizedTime <= lastSubmissionTime.Uint64() {
		return nil
	}
	lockKey := GetOracleLockKey(oracle, normalizedTime)
	err = red.SetNX(lockKey, types.StringToGenericValue("1"), time.Minute*5)
	//nolint:nilerr //already processed/processing
	if err != nil {
		return nil
	}

	jobs := oracle.Spec.Jobs
	shouldEarlyTerminate := oracle.Spec.ShouldEarlyTerminate
	result := types.NewAdapterResult()

	input := types.AdapterRunTimeInput{
		BeginTime: currentTime,
		Config:    app.Config,
	}

	for _, job := range jobs {
		adapter, ok := app.AdapterMap[job.Adapter]
		if !ok {
			panic("adapter should exist: " + job.Adapter)
		}
		input.LastStoreData, input.LastStoreDataExists, err = GetLastStoreData(red, adapter, job)
		if err != nil {
			return err
		}
		store := types.NewAdapterStore()
		result, err = adapter.Perform(job, result, input, &store)
		if err != nil {
			log.Error(fmt.Errorf("%s: %s: %s", oracle.Id, adapter.Id(), err.Error()))
			if shouldEarlyTerminate {
				break
			}
		}
		if store.ShouldPersist {
			if err := SetLastStoreData(red, adapter, job, store); err != nil {
				return err
			}
		}
	}

	err = red.Set(LastSubmissionTimeKey, types.Uint64ToGenericValue(normalizedTime), 0)
	if err != nil {
		return err
	}

	resultData := result.GetData(oracle.Spec.OutputId).String()

	SaveOracleResult(resultData, oracle.Id, red)

	if OracleOverwriteData == constants.True {
		resultData = overwriteData(oracle.Id, resultData) // if we want to override oracle price
		// resultData = overwriteDataV2(oracle.ID, resultData) // if we want to override oracle price
	}

	if resultData == "" {
		return errors.New("skipping submission for " + oracle.Id + " as result is empty")
	}

	msg := oracletypes.MsgCreateVote{
		Creator:   app.Wallet.Bech32Addr,
		OracleId:  oracle.Id,
		Timestamp: int64(normalizedTime),
		Data:      resultData,
	}

	SubmitVote(app, msg)

	return nil
}

// RunOracles run oracle submissions
func RunOracles(app types.App, t uint64) {
	for _, oracle := range app.Oracles {
		go func(currOracle types.Oracle) {
			err := RunOracle(app, currOracle, t)
			if err != nil {
				log.Warnln(err)
			}
		}(oracle)
	}
}

// Run run oracles
func Run(app types.App) {
	log.Info("[oracle] Service started.")
	count := 0
	for {
		if count == 0 { // on init, and every minute
			oracles, err := SyncOracles(app)
			if err != nil {
				log.Warn(err)
				time.Sleep(time.Second)
				continue
			}
			app.Oracles = oracles
		}

		RunOracles(app, uint64(time.Now().Unix()))
		time.Sleep(100 * time.Millisecond)

		count++
		if count > 600 { // 600 * 0.1s = 60s = every minute
			count = 0
		}
	}
}
