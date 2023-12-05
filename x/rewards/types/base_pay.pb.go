// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/rewards/base_pay.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// aggregated rewards for the provider through out the month
type BasePay struct {
	Total github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=total,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total"`
}

func (m *BasePay) Reset()         { *m = BasePay{} }
func (m *BasePay) String() string { return proto.CompactTextString(m) }
func (*BasePay) ProtoMessage()    {}
func (*BasePay) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2fb0eb917a4ee4e, []int{0}
}
func (m *BasePay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasePay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasePay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasePay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasePay.Merge(m, src)
}
func (m *BasePay) XXX_Size() int {
	return m.Size()
}
func (m *BasePay) XXX_DiscardUnknown() {
	xxx_messageInfo_BasePay.DiscardUnknown(m)
}

var xxx_messageInfo_BasePay proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BasePay)(nil), "lavanet.lava.rewards.BasePay")
}

func init() {
	proto.RegisterFile("lavanet/lava/rewards/base_pay.proto", fileDescriptor_a2fb0eb917a4ee4e)
}

var fileDescriptor_a2fb0eb917a4ee4e = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0x45, 0xa9, 0xe5, 0x89, 0x45, 0x29, 0xc5, 0xfa, 0x49,
	0x89, 0xc5, 0xa9, 0xf1, 0x05, 0x89, 0x95, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x22, 0x50,
	0x45, 0x7a, 0x20, 0x5a, 0x0f, 0xaa, 0x48, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac, 0x40, 0x1f,
	0xc4, 0x82, 0xa8, 0x95, 0x92, 0x4b, 0xce, 0x2f, 0xce, 0xcd, 0x87, 0x18, 0xa1, 0x5f, 0x66, 0x98,
	0x94, 0x5a, 0x92, 0x68, 0xa8, 0x9f, 0x9c, 0x9f, 0x99, 0x07, 0x95, 0x97, 0x84, 0xc8, 0xc7, 0x43,
	0x34, 0x42, 0x38, 0x50, 0x29, 0xc1, 0xc4, 0xdc, 0xcc, 0xbc, 0x7c, 0x7d, 0x30, 0x09, 0x11, 0x52,
	0x8a, 0xe5, 0x62, 0x77, 0x4a, 0x2c, 0x4e, 0x0d, 0x48, 0xac, 0x14, 0x0a, 0xe2, 0x62, 0x2d, 0xc9,
	0x2f, 0x49, 0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0xb2, 0x39, 0x71, 0x4f, 0x9e, 0xe1,
	0xd6, 0x3d, 0x79, 0xb5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xa8, 0x69,
	0x50, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x58, 0xcf, 0x33, 0xaf, 0xe4,
	0xd2, 0x16, 0x5d, 0x2e, 0xa8, 0x65, 0x9e, 0x79, 0x25, 0x41, 0x10, 0xa3, 0x9c, 0x1c, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x1d, 0xc9, 0x58, 0x94, 0x20, 0xaa, 0x80,
	0x07, 0x12, 0xd8, 0xec, 0x24, 0x36, 0xb0, 0x43, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5,
	0xed, 0x03, 0x36, 0x49, 0x01, 0x00, 0x00,
}

func (m *BasePay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasePay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasePay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Total.Size()
		i -= size
		if _, err := m.Total.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBasePay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintBasePay(dAtA []byte, offset int, v uint64) int {
	offset -= sovBasePay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BasePay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Total.Size()
	n += 1 + l + sovBasePay(uint64(l))
	return n
}

func sovBasePay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBasePay(x uint64) (n int) {
	return sovBasePay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BasePay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBasePay
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
			return fmt.Errorf("proto: BasePay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasePay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasePay
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
				return ErrInvalidLengthBasePay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasePay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Total.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBasePay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBasePay
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
func skipBasePay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBasePay
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
					return 0, ErrIntOverflowBasePay
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
					return 0, ErrIntOverflowBasePay
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
				return 0, ErrInvalidLengthBasePay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBasePay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBasePay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBasePay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBasePay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBasePay = fmt.Errorf("proto: unexpected end of group")
)