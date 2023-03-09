// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/challenge/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the parameters for the module.
type Params struct {
	// Challenges which will be emitted in each block, including user triggered or randomly triggered.
	ChallengeCountPerBlock uint64 `protobuf:"varint,1,opt,name=challenge_count_per_block,json=challengeCountPerBlock,proto3" json:"challenge_count_per_block,omitempty" yaml:"challenge_count_per_block"`
	// The count of blocks to stand for the period in which the same storage and object info cannot be slashed again.
	SlashCoolingOffPeriod uint64 `protobuf:"varint,2,opt,name=slash_cooling_off_period,json=slashCoolingOffPeriod,proto3" json:"slash_cooling_off_period,omitempty" yaml:"slash_cooling_off_period"`
	// The slash coin amount will be calculated from the size of object info, and adjusted by this rate.
	SlashAmountSizeRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=slash_amount_size_rate,json=slashAmountSizeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_amount_size_rate" yaml:"slash_amount_size_rate"`
	// The minimal slash amount.
	SlashAmountMin github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=slash_amount_min,json=slashAmountMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_amount_min"`
	// The maximum slash amount.
	SlashAmountMax github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=slash_amount_max,json=slashAmountMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_amount_max"`
	// The ratio of slash amount for all validator rewards.
	RewardValidatorRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=reward_validator_ratio,json=rewardValidatorRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_validator_ratio" yaml:"reward_validator_ratio"`
	// The ratio of reward amount for submitter rewards.
	RewardSubmitterRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=reward_submitter_ratio,json=rewardSubmitterRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_submitter_ratio" yaml:"reward_challenger_ratio"`
	// The reward amount to submitter will be adjusted by the threshold.
	RewardSubmitterThreshold github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=reward_submitter_threshold,json=rewardSubmitterThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reward_submitter_threshold"`
	// Heartbeat interval defines the frequency of heartbeat based on challenges.
	HeartbeatInterval uint64 `protobuf:"varint,9,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty" yaml:"heartbeat_interval"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2396367ee53edf57, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetChallengeCountPerBlock() uint64 {
	if m != nil {
		return m.ChallengeCountPerBlock
	}
	return 0
}

func (m *Params) GetSlashCoolingOffPeriod() uint64 {
	if m != nil {
		return m.SlashCoolingOffPeriod
	}
	return 0
}

func (m *Params) GetHeartbeatInterval() uint64 {
	if m != nil {
		return m.HeartbeatInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "bnbchain.greenfield.challenge.Params")
}

func init() { proto.RegisterFile("greenfield/challenge/params.proto", fileDescriptor_2396367ee53edf57) }

var fileDescriptor_2396367ee53edf57 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x6f, 0xd3, 0x3e,
	0x1c, 0x6d, 0xfe, 0xff, 0xd2, 0xb1, 0x1c, 0x10, 0x84, 0x51, 0xa5, 0x95, 0x9a, 0x94, 0x80, 0xd0,
	0x2e, 0x6d, 0x85, 0xb8, 0x4d, 0x5c, 0xe8, 0xb8, 0x54, 0x80, 0x56, 0x65, 0x88, 0x03, 0x42, 0xb2,
	0x9c, 0xc4, 0x49, 0xac, 0x39, 0x76, 0x65, 0xbb, 0xa3, 0xeb, 0x19, 0x38, 0x73, 0xe4, 0xc8, 0x87,
	0xe0, 0x43, 0xec, 0x38, 0x71, 0x42, 0x1c, 0x22, 0xd4, 0x7e, 0x83, 0x7e, 0x01, 0x50, 0xec, 0xb6,
	0xeb, 0x5a, 0x76, 0x98, 0xd8, 0xa9, 0xcd, 0x7b, 0xcf, 0xef, 0x3d, 0xdb, 0xc9, 0xcf, 0xbc, 0x9f,
	0x70, 0x84, 0x68, 0x8c, 0x11, 0x89, 0x3a, 0x61, 0x0a, 0x09, 0x41, 0x34, 0x41, 0x9d, 0x01, 0xe4,
	0x30, 0x13, 0xed, 0x01, 0x67, 0x92, 0x59, 0x8d, 0x80, 0x06, 0x61, 0x0a, 0x31, 0x6d, 0x9f, 0x6b,
	0xdb, 0x4b, 0x6d, 0xbd, 0x16, 0x32, 0x91, 0x31, 0x01, 0x94, 0xb8, 0xa3, 0x1f, 0xf4, 0xca, 0xfa,
	0x4e, 0xc2, 0x12, 0xa6, 0xf1, 0xe2, 0x9f, 0x46, 0xbd, 0xdf, 0x5b, 0x66, 0xa5, 0xaf, 0x02, 0x2c,
	0x60, 0xd6, 0x96, 0x46, 0x20, 0x64, 0x43, 0x2a, 0xc1, 0x00, 0x71, 0x10, 0x10, 0x16, 0x1e, 0xd9,
	0x46, 0xd3, 0xd8, 0x2d, 0x77, 0x1f, 0xce, 0x72, 0xb7, 0x79, 0x02, 0x33, 0xb2, 0xe7, 0x5d, 0x2a,
	0xf5, 0xfc, 0xea, 0x92, 0xdb, 0x2f, 0xa8, 0x3e, 0xe2, 0xdd, 0x82, 0xb0, 0xde, 0x99, 0xb6, 0x20,
	0x50, 0xa4, 0x20, 0x64, 0x8c, 0x60, 0x9a, 0x00, 0x16, 0xc7, 0xc5, 0x3a, 0xcc, 0x22, 0xfb, 0x3f,
	0xe5, 0xff, 0x60, 0x96, 0xbb, 0xae, 0xf6, 0xbf, 0x4c, 0xe9, 0xf9, 0xf7, 0x14, 0xb5, 0xaf, 0x99,
	0x83, 0x38, 0xee, 0x2b, 0xdc, 0xfa, 0x60, 0x98, 0x55, 0xbd, 0x08, 0x66, 0xaa, 0x91, 0xc0, 0x63,
	0x04, 0x38, 0x94, 0xc8, 0xfe, 0xbf, 0x69, 0xec, 0x6e, 0x77, 0x0f, 0x4e, 0x73, 0xb7, 0xf4, 0x33,
	0x77, 0x1f, 0x25, 0x58, 0xa6, 0xc3, 0xa0, 0x1d, 0xb2, 0x6c, 0x7e, 0x42, 0xf3, 0x9f, 0x96, 0x88,
	0x8e, 0x3a, 0xf2, 0x64, 0x80, 0x44, 0xfb, 0x39, 0x0a, 0x67, 0xb9, 0xdb, 0x58, 0xad, 0xb2, 0xee,
	0xea, 0xf9, 0x77, 0x15, 0xf1, 0x4c, 0xe1, 0x87, 0x78, 0x8c, 0x7c, 0x28, 0x91, 0x15, 0x9b, 0xb7,
	0x2f, 0xe8, 0x33, 0x4c, 0xed, 0xb2, 0xca, 0x7f, 0x7a, 0x85, 0xfc, 0x1e, 0x95, 0xdf, 0xbf, 0xb5,
	0xcc, 0xf9, 0x05, 0xf6, 0xa8, 0xf4, 0x6f, 0xad, 0x84, 0xbd, 0xc2, 0x74, 0x33, 0x07, 0x8e, 0xec,
	0x1b, 0xd7, 0x9d, 0x03, 0x47, 0xd6, 0x47, 0xc3, 0xac, 0x72, 0xf4, 0x1e, 0xf2, 0x08, 0x1c, 0x43,
	0x82, 0x23, 0x28, 0x19, 0x2f, 0xf6, 0x8f, 0x99, 0x5d, 0xf9, 0xb7, 0x63, 0xfd, 0xbb, 0xab, 0xe7,
	0xef, 0x68, 0xe2, 0xcd, 0x02, 0xf7, 0x0b, 0xd8, 0xfa, 0x74, 0xde, 0x43, 0x0c, 0x83, 0x0c, 0x4b,
	0x89, 0x16, 0x3d, 0xb6, 0x54, 0x8f, 0xfe, 0x95, 0x7b, 0x38, 0x17, 0x7a, 0x2c, 0x5f, 0xda, 0xf5,
	0x22, 0x87, 0x8b, 0x38, 0x5d, 0x64, 0x6c, 0xd6, 0x37, 0x7a, 0xc8, 0x94, 0x23, 0x91, 0x32, 0x12,
	0xd9, 0x37, 0xaf, 0xe1, 0x0a, 0xec, 0xb5, 0xdc, 0xd7, 0x0b, 0x77, 0xeb, 0xa5, 0x69, 0xa5, 0x08,
	0x72, 0x19, 0x20, 0x28, 0x01, 0xa6, 0x12, 0xf1, 0x63, 0x48, 0xec, 0x6d, 0xf5, 0xed, 0x34, 0x66,
	0xb9, 0x5b, 0xd3, 0x3b, 0xda, 0xd4, 0x78, 0xfe, 0x9d, 0x25, 0xd8, 0x9b, 0x63, 0x7b, 0xe5, 0x2f,
	0x5f, 0xdd, 0x52, 0xf7, 0xc5, 0xe9, 0xc4, 0x31, 0xce, 0x26, 0x8e, 0xf1, 0x6b, 0xe2, 0x18, 0x9f,
	0xa7, 0x4e, 0xe9, 0x6c, 0xea, 0x94, 0x7e, 0x4c, 0x9d, 0xd2, 0xdb, 0xc7, 0x2b, 0xed, 0x03, 0x1a,
	0xb4, 0xd4, 0xdc, 0xe9, 0xac, 0xcc, 0xa8, 0xd1, 0xca, 0x94, 0x52, 0x9b, 0x09, 0x2a, 0x6a, 0xaa,
	0x3c, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x86, 0x5e, 0x7a, 0x44, 0xca, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HeartbeatInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HeartbeatInterval))
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.RewardSubmitterThreshold.Size()
		i -= size
		if _, err := m.RewardSubmitterThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.RewardSubmitterRatio.Size()
		i -= size
		if _, err := m.RewardSubmitterRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.RewardValidatorRatio.Size()
		i -= size
		if _, err := m.RewardValidatorRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.SlashAmountMax.Size()
		i -= size
		if _, err := m.SlashAmountMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SlashAmountMin.Size()
		i -= size
		if _, err := m.SlashAmountMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SlashAmountSizeRate.Size()
		i -= size
		if _, err := m.SlashAmountSizeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.SlashCoolingOffPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SlashCoolingOffPeriod))
		i--
		dAtA[i] = 0x10
	}
	if m.ChallengeCountPerBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ChallengeCountPerBlock))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChallengeCountPerBlock != 0 {
		n += 1 + sovParams(uint64(m.ChallengeCountPerBlock))
	}
	if m.SlashCoolingOffPeriod != 0 {
		n += 1 + sovParams(uint64(m.SlashCoolingOffPeriod))
	}
	l = m.SlashAmountSizeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashAmountMin.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashAmountMax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardValidatorRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardSubmitterRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardSubmitterThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.HeartbeatInterval != 0 {
		n += 1 + sovParams(uint64(m.HeartbeatInterval))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeCountPerBlock", wireType)
			}
			m.ChallengeCountPerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeCountPerBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashCoolingOffPeriod", wireType)
			}
			m.SlashCoolingOffPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashCoolingOffPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountSizeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashAmountSizeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashAmountMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashAmountMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardValidatorRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardValidatorRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardSubmitterRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardSubmitterRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardSubmitterThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardSubmitterThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatInterval", wireType)
			}
			m.HeartbeatInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeartbeatInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
