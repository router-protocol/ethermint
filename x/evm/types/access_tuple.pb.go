// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/evm/v1/access_tuple.proto

package types

import (
	fmt "fmt"
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

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	// address is a hex formatted ethereum address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// storage_keys are hex formatted hashes of the storage keys
	StorageKeys []string `protobuf:"bytes,2,rep,name=storage_keys,json=storageKeys,proto3" json:"storageKeys"`
}

func (m *AccessTuple) Reset()         { *m = AccessTuple{} }
func (m *AccessTuple) String() string { return proto.CompactTextString(m) }
func (*AccessTuple) ProtoMessage()    {}
func (*AccessTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8644a7fc00ba854, []int{0}
}
func (m *AccessTuple) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessTuple.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTuple.Merge(m, src)
}
func (m *AccessTuple) XXX_Size() int {
	return m.Size()
}
func (m *AccessTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTuple.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTuple proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AccessTuple)(nil), "ethermint.evm.v1.AccessTuple")
}

func init() {
	proto.RegisterFile("ethermint/evm/v1/access_tuple.proto", fileDescriptor_d8644a7fc00ba854)
}

var fileDescriptor_d8644a7fc00ba854 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2d, 0xc9, 0x48,
	0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x4f, 0x2d, 0xcb, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0x4c, 0x4e,
	0x4e, 0x2d, 0x2e, 0x8e, 0x2f, 0x29, 0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x80, 0x2b, 0xd2, 0x4b, 0x2d, 0xcb, 0xd5, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x4b, 0xea, 0x83, 0x58, 0x10, 0x75, 0x4a, 0x89, 0x5c, 0xdc, 0x8e, 0x60, 0xdd, 0x21, 0x20,
	0xcd, 0x42, 0x12, 0x5c, 0xec, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x11, 0x17, 0x4f, 0x71, 0x49, 0x7e, 0x51, 0x62, 0x7a, 0x6a,
	0x7c, 0x76, 0x6a, 0x65, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xa7, 0x13, 0xff, 0xab, 0x7b, 0xf2,
	0xdc, 0x50, 0x71, 0xef, 0xd4, 0xca, 0xe2, 0x20, 0x64, 0x8e, 0x15, 0x4b, 0xc7, 0x02, 0x79, 0x06,
	0x27, 0x87, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x52, 0x4b, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0x05, 0x79, 0x25, 0xbf, 0x58, 0x1f, 0xe1, 0xb5, 0x0a,
	0xb0, 0xe7, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x6e, 0x35, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x40, 0xb6, 0xb1, 0x5f, 0xfa, 0x00, 0x00, 0x00,
}

func (m *AccessTuple) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessTuple) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessTuple) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StorageKeys) > 0 {
		for iNdEx := len(m.StorageKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StorageKeys[iNdEx])
			copy(dAtA[i:], m.StorageKeys[iNdEx])
			i = encodeVarintAccessTuple(dAtA, i, uint64(len(m.StorageKeys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAccessTuple(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccessTuple(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccessTuple(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessTuple) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccessTuple(uint64(l))
	}
	if len(m.StorageKeys) > 0 {
		for _, s := range m.StorageKeys {
			l = len(s)
			n += 1 + l + sovAccessTuple(uint64(l))
		}
	}
	return n
}

func sovAccessTuple(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccessTuple(x uint64) (n int) {
	return sovAccessTuple(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessTuple) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccessTuple
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
			return fmt.Errorf("proto: AccessTuple: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessTuple: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTuple
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
				return ErrInvalidLengthAccessTuple
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTuple
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTuple
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
				return ErrInvalidLengthAccessTuple
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTuple
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageKeys = append(m.StorageKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccessTuple(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccessTuple
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
func skipAccessTuple(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccessTuple
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
					return 0, ErrIntOverflowAccessTuple
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
					return 0, ErrIntOverflowAccessTuple
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
				return 0, ErrInvalidLengthAccessTuple
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccessTuple
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccessTuple
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccessTuple        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccessTuple          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccessTuple = fmt.Errorf("proto: unexpected end of group")
)