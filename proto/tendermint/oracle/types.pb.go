// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/oracle/types.proto

package oracle

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Vote struct {
	Validator string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	OracleId  string `protobuf:"bytes,2,opt,name=oracle_id,json=oracleId,proto3" json:"oracle_id,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9227d272ed5d90, []int{0}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *Vote) GetOracleId() string {
	if m != nil {
		return m.OracleId
	}
	return ""
}

func (m *Vote) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Vote) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GossipedVotes struct {
	ValidatorIndex  int32   `protobuf:"varint,1,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	Votes           []*Vote `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
	SignedTimestamp int64   `protobuf:"varint,3,opt,name=signed_timestamp,json=signedTimestamp,proto3" json:"signed_timestamp,omitempty"`
	Signature       []byte  `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *GossipedVotes) Reset()         { *m = GossipedVotes{} }
func (m *GossipedVotes) String() string { return proto.CompactTextString(m) }
func (*GossipedVotes) ProtoMessage()    {}
func (*GossipedVotes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9227d272ed5d90, []int{1}
}
func (m *GossipedVotes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GossipedVotes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GossipedVotes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GossipedVotes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipedVotes.Merge(m, src)
}
func (m *GossipedVotes) XXX_Size() int {
	return m.Size()
}
func (m *GossipedVotes) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipedVotes.DiscardUnknown(m)
}

var xxx_messageInfo_GossipedVotes proto.InternalMessageInfo

func (m *GossipedVotes) GetValidatorIndex() int32 {
	if m != nil {
		return m.ValidatorIndex
	}
	return 0
}

func (m *GossipedVotes) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GossipedVotes) GetSignedTimestamp() int64 {
	if m != nil {
		return m.SignedTimestamp
	}
	return 0
}

func (m *GossipedVotes) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type CanonicalGossipedVotes struct {
	ValidatorIndex  int32   `protobuf:"varint,1,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	Votes           []*Vote `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
	SignedTimestamp int64   `protobuf:"varint,3,opt,name=signed_timestamp,json=signedTimestamp,proto3" json:"signed_timestamp,omitempty"`
}

func (m *CanonicalGossipedVotes) Reset()         { *m = CanonicalGossipedVotes{} }
func (m *CanonicalGossipedVotes) String() string { return proto.CompactTextString(m) }
func (*CanonicalGossipedVotes) ProtoMessage()    {}
func (*CanonicalGossipedVotes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9227d272ed5d90, []int{2}
}
func (m *CanonicalGossipedVotes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalGossipedVotes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalGossipedVotes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalGossipedVotes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalGossipedVotes.Merge(m, src)
}
func (m *CanonicalGossipedVotes) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalGossipedVotes) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalGossipedVotes.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalGossipedVotes proto.InternalMessageInfo

func (m *CanonicalGossipedVotes) GetValidatorIndex() int32 {
	if m != nil {
		return m.ValidatorIndex
	}
	return 0
}

func (m *CanonicalGossipedVotes) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *CanonicalGossipedVotes) GetSignedTimestamp() int64 {
	if m != nil {
		return m.SignedTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Vote)(nil), "tendermint.oracle.Vote")
	proto.RegisterType((*GossipedVotes)(nil), "tendermint.oracle.GossipedVotes")
	proto.RegisterType((*CanonicalGossipedVotes)(nil), "tendermint.oracle.CanonicalGossipedVotes")
}

func init() { proto.RegisterFile("tendermint/oracle/types.proto", fileDescriptor_ed9227d272ed5d90) }

var fileDescriptor_ed9227d272ed5d90 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x3d, 0x4e, 0xc3, 0x30,
	0x14, 0xc7, 0xeb, 0x7e, 0x20, 0x62, 0x3e, 0x0a, 0x1e, 0x20, 0x12, 0xc5, 0xaa, 0xba, 0x50, 0x06,
	0x12, 0x09, 0x38, 0x01, 0x0c, 0xa8, 0x0b, 0x43, 0x84, 0x18, 0x58, 0x2a, 0x27, 0x36, 0xc5, 0x52,
	0x62, 0x47, 0xf1, 0x6b, 0x05, 0xb7, 0x60, 0xe7, 0x12, 0x1c, 0x83, 0xb1, 0x23, 0x23, 0x4a, 0x2e,
	0x82, 0xec, 0x48, 0x89, 0x44, 0x2f, 0xc0, 0xf6, 0xf4, 0xfb, 0xbf, 0xe7, 0xf7, 0xb3, 0xf4, 0xf0,
	0x29, 0x08, 0xc5, 0x45, 0x91, 0x49, 0x05, 0xa1, 0x2e, 0x58, 0x92, 0x8a, 0x10, 0xde, 0x72, 0x61,
	0x82, 0xbc, 0xd0, 0xa0, 0xc9, 0x61, 0x1b, 0x07, 0x75, 0x3c, 0x31, 0xb8, 0xff, 0xa8, 0x41, 0x90,
	0x11, 0xf6, 0x56, 0x2c, 0x95, 0x9c, 0x81, 0x2e, 0x7c, 0x34, 0x46, 0x53, 0x2f, 0x6a, 0x01, 0x39,
	0xc1, 0x5e, 0xdd, 0x3f, 0x97, 0xdc, 0xef, 0xba, 0x74, 0xbb, 0x06, 0x33, 0x6e, 0x47, 0x41, 0x66,
	0xc2, 0x00, 0xcb, 0x72, 0xbf, 0x37, 0x46, 0xd3, 0x5e, 0xd4, 0x02, 0x42, 0x70, 0x9f, 0x33, 0x60,
	0x7e, 0xdf, 0x4d, 0xb9, 0x7a, 0xf2, 0x89, 0xf0, 0xde, 0x9d, 0x36, 0x46, 0xe6, 0x82, 0xdb, 0xed,
	0x86, 0x9c, 0xe1, 0x61, 0xb3, 0x6d, 0x2e, 0x15, 0x17, 0xaf, 0x4e, 0x62, 0x10, 0xed, 0x37, 0x78,
	0x66, 0x29, 0xb9, 0xc0, 0x83, 0x95, 0x9d, 0xf0, 0xbb, 0xe3, 0xde, 0x74, 0xe7, 0xf2, 0x38, 0xd8,
	0xf8, 0x52, 0x60, 0x5f, 0x8c, 0xea, 0x2e, 0x72, 0x8e, 0x0f, 0x8c, 0x5c, 0x28, 0xc1, 0xe7, 0x7f,
	0x15, 0x87, 0x35, 0x7f, 0x68, 0x44, 0x47, 0xd8, 0xb3, 0x88, 0xc1, 0xb2, 0x10, 0xce, 0x76, 0x37,
	0x6a, 0xc1, 0xe4, 0x03, 0xe1, 0xa3, 0x5b, 0xa6, 0xb4, 0x92, 0x09, 0x4b, 0xff, 0x9b, 0xfb, 0xcd,
	0xfd, 0x57, 0x49, 0xd1, 0xba, 0xa4, 0xe8, 0xa7, 0xa4, 0xe8, 0xbd, 0xa2, 0x9d, 0x75, 0x45, 0x3b,
	0xdf, 0x15, 0xed, 0x3c, 0x5d, 0x2f, 0x24, 0xbc, 0x2c, 0xe3, 0x20, 0xd1, 0x59, 0x98, 0xe8, 0x4c,
	0x40, 0xfc, 0x0c, 0x6d, 0xe1, 0xce, 0x22, 0xdc, 0x38, 0x9a, 0x78, 0xcb, 0x05, 0x57, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x64, 0xa4, 0x9b, 0xfc, 0x50, 0x02, 0x00, 0x00,
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if m.Timestamp != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.OracleId) > 0 {
		i -= len(m.OracleId)
		copy(dAtA[i:], m.OracleId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.OracleId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GossipedVotes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GossipedVotes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GossipedVotes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if m.SignedTimestamp != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SignedTimestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ValidatorIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ValidatorIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalGossipedVotes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalGossipedVotes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalGossipedVotes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignedTimestamp != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SignedTimestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ValidatorIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ValidatorIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.OracleId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTypes(uint64(m.Timestamp))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *GossipedVotes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidatorIndex != 0 {
		n += 1 + sovTypes(uint64(m.ValidatorIndex))
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.SignedTimestamp != 0 {
		n += 1 + sovTypes(uint64(m.SignedTimestamp))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *CanonicalGossipedVotes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidatorIndex != 0 {
		n += 1 + sovTypes(uint64(m.ValidatorIndex))
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.SignedTimestamp != 0 {
		n += 1 + sovTypes(uint64(m.SignedTimestamp))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GossipedVotes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GossipedVotes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GossipedVotes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorIndex", wireType)
			}
			m.ValidatorIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedTimestamp", wireType)
			}
			m.SignedTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CanonicalGossipedVotes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CanonicalGossipedVotes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalGossipedVotes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorIndex", wireType)
			}
			m.ValidatorIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedTimestamp", wireType)
			}
			m.SignedTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
