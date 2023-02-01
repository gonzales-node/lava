// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pairing/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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
	MintCoinsPerCU           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=mintCoinsPerCU,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"mintCoinsPerCU" yaml:"mint_coins_per_cu"`
	BurnCoinsPerCU           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=burnCoinsPerCU,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"burnCoinsPerCU" yaml:"burn_coins_per_cu"`
	FraudStakeSlashingFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=fraudStakeSlashingFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fraudStakeSlashingFactor" yaml:"fraud_stake_slashing_factor"`
	FraudSlashingAmount      uint64                                 `protobuf:"varint,6,opt,name=fraudSlashingAmount,proto3" json:"fraudSlashingAmount,omitempty" yaml:"fraud_slashing_amount"`
	ServicersToPairCount     uint64                                 `protobuf:"varint,7,opt,name=servicersToPairCount,proto3" json:"servicersToPairCount,omitempty" yaml:"servicers_to_pair_count"`
	EpochBlocksOverlap       uint64                                 `protobuf:"varint,8,opt,name=epochBlocksOverlap,proto3" json:"epochBlocksOverlap,omitempty" yaml:"epoch_blocks_overlap"`
	StakeToMaxCUList         StakeToMaxCUList                       `protobuf:"bytes,9,opt,name=stakeToMaxCUList,proto3,customtype=StakeToMaxCUList" json:"stakeToMaxCUList" yaml:"stake_to_computeunits_list"`
	UnpayLimit               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=unpayLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"unpayLimit" yaml:"unpay_limit"`
	SlashLimit               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=slashLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slashLimit" yaml:"slash_limit"`
	DataReliabilityReward    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=dataReliabilityReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"dataReliabilityReward" yaml:"data_reliability_reward"`
	QoSWeight                github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=QoSWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"QoSWeight" yaml:"data_reliability_reward"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_72cc734580d3bc3a, []int{0}
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

func (m *Params) GetFraudSlashingAmount() uint64 {
	if m != nil {
		return m.FraudSlashingAmount
	}
	return 0
}

func (m *Params) GetServicersToPairCount() uint64 {
	if m != nil {
		return m.ServicersToPairCount
	}
	return 0
}

func (m *Params) GetEpochBlocksOverlap() uint64 {
	if m != nil {
		return m.EpochBlocksOverlap
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "lavanet.lava.pairing.Params")
}

func init() { proto.RegisterFile("pairing/params.proto", fileDescriptor_72cc734580d3bc3a) }

var fileDescriptor_72cc734580d3bc3a = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0xe3, 0x36, 0xbf, 0xfe, 0x9a, 0xe3, 0x8f, 0xa2, 0x23, 0x48, 0x16, 0x20, 0xbb, 0x78,
	0x80, 0x2e, 0xc4, 0x03, 0x5b, 0x25, 0x86, 0x26, 0x15, 0x43, 0x55, 0xd4, 0xe0, 0xb4, 0x20, 0xb1,
	0x9c, 0x2e, 0xce, 0xd5, 0x39, 0xc5, 0xf6, 0x59, 0x77, 0xe7, 0xd0, 0xbc, 0x01, 0xe6, 0x8e, 0x8c,
	0xbc, 0x16, 0xa6, 0x8e, 0x1d, 0x11, 0x83, 0x85, 0x92, 0x77, 0x90, 0x57, 0x80, 0xfc, 0xd8, 0x6d,
	0xd2, 0x10, 0x86, 0xa8, 0x62, 0x3a, 0xcb, 0xf7, 0x7d, 0x3e, 0x9f, 0xd3, 0x3d, 0xa7, 0x07, 0x35,
	0x12, 0xca, 0x25, 0x8f, 0x03, 0x37, 0xa1, 0x92, 0x46, 0xaa, 0x99, 0x48, 0xa1, 0x05, 0x6e, 0x84,
	0x74, 0x44, 0x63, 0xa6, 0x9b, 0xf9, 0xda, 0x2c, 0x23, 0x4f, 0x1a, 0x81, 0x08, 0x04, 0x04, 0xdc,
	0xfc, 0xab, 0xc8, 0x3a, 0xdf, 0x6b, 0x68, 0xab, 0x03, 0xc5, 0x58, 0xa2, 0x87, 0x11, 0x8f, 0x75,
	0x5b, 0xf0, 0x58, 0x75, 0x98, 0x6c, 0x9f, 0x9a, 0x9b, 0x3b, 0xc6, 0x6e, 0xad, 0x75, 0x78, 0x99,
	0xd9, 0x95, 0x9f, 0x99, 0xfd, 0x22, 0xe0, 0x7a, 0x90, 0xf6, 0x9a, 0xbe, 0x88, 0x5c, 0x5f, 0xa8,
	0x48, 0xa8, 0x72, 0x79, 0xa5, 0xfa, 0x43, 0x57, 0x8f, 0x13, 0xa6, 0x9a, 0x07, 0xcc, 0x9f, 0x65,
	0xb6, 0x39, 0xa6, 0x51, 0xb8, 0xe7, 0xe4, 0x34, 0xe2, 0xe7, 0x38, 0x92, 0x30, 0x49, 0xfc, 0xd4,
	0xf1, 0x96, 0x0c, 0xb9, 0xb3, 0x97, 0xca, 0x78, 0xc1, 0x59, 0xbd, 0x9b, 0x33, 0xa7, 0x2d, 0x3b,
	0x6f, 0x1b, 0xf0, 0x85, 0x81, 0xcc, 0x33, 0x49, 0xd3, 0x7e, 0x57, 0xd3, 0x21, 0xeb, 0x86, 0x54,
	0x0d, 0x78, 0x1c, 0xbc, 0xa5, 0xbe, 0x16, 0xd2, 0xfc, 0x0f, 0xf4, 0x27, 0x6b, 0xeb, 0x9d, 0x42,
	0x0f, 0x5c, 0xa2, 0x72, 0x30, 0x51, 0x25, 0x99, 0x9c, 0x01, 0xda, 0xf1, 0xfe, 0x6a, 0xc5, 0x1e,
	0x7a, 0x54, 0xec, 0x95, 0xbf, 0xf7, 0x23, 0x91, 0xc6, 0xda, 0xdc, 0xda, 0x31, 0x76, 0xab, 0xad,
	0x9d, 0x59, 0x66, 0x3f, 0xbb, 0x85, 0xbf, 0x06, 0x53, 0x88, 0x39, 0xde, 0xaa, 0x62, 0xfc, 0x01,
	0x35, 0x14, 0x93, 0x23, 0xee, 0x33, 0xa9, 0x4e, 0x44, 0x87, 0x72, 0xd9, 0x06, 0xe8, 0xff, 0x00,
	0x75, 0x66, 0x99, 0x6d, 0x15, 0xd0, 0x9b, 0x14, 0xd1, 0x82, 0xe4, 0xaf, 0x85, 0xf8, 0x05, 0x76,
	0x65, 0x3d, 0x3e, 0x46, 0x98, 0x25, 0xc2, 0x1f, 0xb4, 0x42, 0xe1, 0x0f, 0xd5, 0xf1, 0x88, 0xc9,
	0x90, 0x26, 0xe6, 0x36, 0x50, 0xed, 0x59, 0x66, 0x3f, 0x2d, 0xa8, 0x90, 0x21, 0x3d, 0x08, 0x11,
	0x51, 0xa4, 0x1c, 0x6f, 0x45, 0x29, 0xe6, 0xa8, 0x0e, 0x17, 0x76, 0x22, 0xde, 0xd1, 0xf3, 0xf6,
	0xe9, 0x11, 0x57, 0xda, 0xac, 0x41, 0x1b, 0xde, 0x94, 0x6d, 0xa8, 0x77, 0x97, 0xf6, 0x67, 0x99,
	0xfd, 0xbc, 0x3c, 0x3c, 0x5c, 0xb5, 0x16, 0xc4, 0x17, 0x51, 0x92, 0x6a, 0x96, 0xc6, 0x5c, 0x2b,
	0x12, 0x72, 0xa5, 0x1d, 0xef, 0x0f, 0x2c, 0xee, 0x23, 0x94, 0xc6, 0x09, 0x1d, 0x1f, 0xf1, 0x88,
	0x6b, 0x13, 0x81, 0xe4, 0x60, 0xed, 0x5e, 0xe3, 0x42, 0x0d, 0x24, 0x12, 0xe6, 0x28, 0xc7, 0x5b,
	0xe0, 0xe6, 0x16, 0x68, 0x51, 0x61, 0xb9, 0x77, 0x37, 0x0b, 0x90, 0x6e, 0x2c, 0x73, 0x2e, 0xfe,
	0x62, 0xa0, 0xc7, 0x7d, 0xaa, 0xa9, 0xc7, 0x42, 0x4e, 0x7b, 0x3c, 0xe4, 0x7a, 0xec, 0xb1, 0xcf,
	0x54, 0xf6, 0xcd, 0xfb, 0x60, 0xec, 0xac, 0x6d, 0x2c, 0xdf, 0x43, 0x0e, 0x25, 0x72, 0x4e, 0x25,
	0x12, 0xb0, 0x8e, 0xb7, 0x5a, 0x87, 0x63, 0x54, 0x7b, 0x2f, 0xba, 0x1f, 0x19, 0x0f, 0x06, 0xda,
	0x7c, 0xf0, 0x8f, 0xdc, 0x73, 0xc5, 0x5e, 0xf5, 0xeb, 0x37, 0xbb, 0x72, 0x58, 0xdd, 0x36, 0xea,
	0x1b, 0x87, 0xd5, 0xed, 0x8d, 0xfa, 0x66, 0x6b, 0xff, 0x72, 0x62, 0x19, 0x57, 0x13, 0xcb, 0xf8,
	0x35, 0xb1, 0x8c, 0x8b, 0xa9, 0x55, 0xb9, 0x9a, 0x5a, 0x95, 0x1f, 0x53, 0xab, 0xf2, 0xe9, 0xe5,
	0xc2, 0x01, 0xca, 0xa9, 0x08, 0xab, 0x7b, 0xee, 0x5e, 0x8f, 0x4e, 0x38, 0x45, 0x6f, 0x0b, 0xc6,
	0xe1, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0x2d, 0x65, 0x2c, 0x52, 0x05, 0x00, 0x00,
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
	{
		size := m.QoSWeight.Size()
		i -= size
		if _, err := m.QoSWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.DataReliabilityReward.Size()
		i -= size
		if _, err := m.DataReliabilityReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.SlashLimit.Size()
		i -= size
		if _, err := m.SlashLimit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.UnpayLimit.Size()
		i -= size
		if _, err := m.UnpayLimit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.StakeToMaxCUList.Size()
		i -= size
		if _, err := m.StakeToMaxCUList.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.EpochBlocksOverlap != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EpochBlocksOverlap))
		i--
		dAtA[i] = 0x40
	}
	if m.ServicersToPairCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ServicersToPairCount))
		i--
		dAtA[i] = 0x38
	}
	if m.FraudSlashingAmount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FraudSlashingAmount))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.FraudStakeSlashingFactor.Size()
		i -= size
		if _, err := m.FraudStakeSlashingFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.BurnCoinsPerCU.Size()
		i -= size
		if _, err := m.BurnCoinsPerCU.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.MintCoinsPerCU.Size()
		i -= size
		if _, err := m.MintCoinsPerCU.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
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
	l = m.MintCoinsPerCU.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.BurnCoinsPerCU.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.FraudStakeSlashingFactor.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.FraudSlashingAmount != 0 {
		n += 1 + sovParams(uint64(m.FraudSlashingAmount))
	}
	if m.ServicersToPairCount != 0 {
		n += 1 + sovParams(uint64(m.ServicersToPairCount))
	}
	if m.EpochBlocksOverlap != 0 {
		n += 1 + sovParams(uint64(m.EpochBlocksOverlap))
	}
	l = m.StakeToMaxCUList.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.UnpayLimit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashLimit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DataReliabilityReward.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.QoSWeight.Size()
	n += 1 + l + sovParams(uint64(l))
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintCoinsPerCU", wireType)
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
			if err := m.MintCoinsPerCU.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnCoinsPerCU", wireType)
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
			if err := m.BurnCoinsPerCU.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FraudStakeSlashingFactor", wireType)
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
			if err := m.FraudStakeSlashingFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FraudSlashingAmount", wireType)
			}
			m.FraudSlashingAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FraudSlashingAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServicersToPairCount", wireType)
			}
			m.ServicersToPairCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServicersToPairCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlocksOverlap", wireType)
			}
			m.EpochBlocksOverlap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochBlocksOverlap |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeToMaxCUList", wireType)
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
			if err := m.StakeToMaxCUList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnpayLimit", wireType)
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
			if err := m.UnpayLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashLimit", wireType)
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
			if err := m.SlashLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataReliabilityReward", wireType)
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
			if err := m.DataReliabilityReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QoSWeight", wireType)
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
			if err := m.QoSWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
