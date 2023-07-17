// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: downtime/v1/downtime.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters of the downtime module.
type Params struct {
	// downtime_duration defines the minimum time elapsed between blocks
	// that we consider the chain to be down.
	DowntimeDuration time.Duration `protobuf:"bytes,1,opt,name=downtime_duration,json=downtimeDuration,proto3,stdduration" json:"downtime_duration"`
	// garbage_collection_duration defines the time a downtime will be recorded
	// on the chain before it goes through garbage collection.
	GarbageCollectionDuration time.Duration `protobuf:"bytes,2,opt,name=garbage_collection_duration,json=garbageCollectionDuration,proto3,stdduration" json:"garbage_collection_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_62cbb07edbf8cff0, []int{0}
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

func (m *Params) GetDowntimeDuration() time.Duration {
	if m != nil {
		return m.DowntimeDuration
	}
	return 0
}

func (m *Params) GetGarbageCollectionDuration() time.Duration {
	if m != nil {
		return m.GarbageCollectionDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "lavanet.lava.downtime.v1.Params")
}

func init() { proto.RegisterFile("downtime/v1/downtime.proto", fileDescriptor_62cbb07edbf8cff0) }

var fileDescriptor_62cbb07edbf8cff0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xc9, 0x2f, 0xcf,
	0x2b, 0xc9, 0xcc, 0x4d, 0xd5, 0x2f, 0x33, 0xd4, 0x87, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x24, 0x72, 0x12, 0xcb, 0x12, 0xf3, 0x52, 0x4b, 0xf4, 0x40, 0xb4, 0x1e, 0x5c, 0xb2, 0xcc,
	0x50, 0x4a, 0x2e, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x1f, 0xac, 0x2e, 0xa9, 0x34, 0x4d, 0x3f,
	0xa5, 0xb4, 0x28, 0xb1, 0x24, 0x33, 0x3f, 0x0f, 0xa2, 0x53, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f,
	0xcc, 0xd4, 0x07, 0xb1, 0x20, 0xa2, 0x4a, 0xfb, 0x19, 0xb9, 0xd8, 0x02, 0x12, 0x8b, 0x12, 0x73,
	0x8b, 0x85, 0x02, 0xb8, 0x04, 0x61, 0xe6, 0xc5, 0xc3, 0xf4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x1b, 0x49, 0xea, 0x41, 0x0c, 0xd7, 0x83, 0x19, 0xae, 0xe7, 0x02, 0x55, 0xe0, 0xc4, 0x71, 0xe2,
	0x9e, 0x3c, 0xc3, 0x8c, 0xfb, 0xf2, 0x8c, 0x41, 0x02, 0x30, 0xdd, 0x30, 0x39, 0xa1, 0x64, 0x2e,
	0xe9, 0xf4, 0xc4, 0xa2, 0xa4, 0xc4, 0xf4, 0xd4, 0xf8, 0xe4, 0xfc, 0x9c, 0x9c, 0xd4, 0x64, 0x90,
	0x28, 0xc2, 0x6c, 0x26, 0xe2, 0xcd, 0x96, 0x84, 0x9a, 0xe3, 0x0c, 0x37, 0x06, 0xae, 0xc8, 0xfe,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x54, 0xd3, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xa1, 0xc1, 0x06, 0xa6, 0xf5, 0x2b, 0xf4, 0x91, 0x42, 0x38,
	0x89, 0x0d, 0x6c, 0xb1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x96, 0x64, 0x69, 0x77, 0x01,
	0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.GarbageCollectionDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.GarbageCollectionDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintDowntime(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DowntimeDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintDowntime(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDowntime(dAtA []byte, offset int, v uint64) int {
	offset -= sovDowntime(v)
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
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeDuration)
	n += 1 + l + sovDowntime(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.GarbageCollectionDuration)
	n += 1 + l + sovDowntime(uint64(l))
	return n
}

func sovDowntime(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDowntime(x uint64) (n int) {
	return sovDowntime(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDowntime
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DowntimeDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDowntime
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
				return ErrInvalidLengthDowntime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDowntime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DowntimeDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GarbageCollectionDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDowntime
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
				return ErrInvalidLengthDowntime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDowntime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.GarbageCollectionDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDowntime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDowntime
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
func skipDowntime(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDowntime
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
					return 0, ErrIntOverflowDowntime
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
					return 0, ErrIntOverflowDowntime
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
				return 0, ErrInvalidLengthDowntime
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDowntime
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDowntime
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDowntime        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDowntime          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDowntime = fmt.Errorf("proto: unexpected end of group")
)
