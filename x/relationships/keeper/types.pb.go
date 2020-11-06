// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/relationships/v1beta1/keeper/types.proto

package keeper

import (
	fmt "fmt"
	types "github.com/desmos-labs/desmos/x/relationships/types"
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

// WrappedRelationships wraps a list of Relationship objects
type WrappedRelationships struct {
	Relationships []types.Relationship `protobuf:"bytes,1,rep,name=relationships,proto3" json:"relationships"`
}

func (m *WrappedRelationships) Reset()         { *m = WrappedRelationships{} }
func (m *WrappedRelationships) String() string { return proto.CompactTextString(m) }
func (*WrappedRelationships) ProtoMessage()    {}
func (*WrappedRelationships) Descriptor() ([]byte, []int) {
	return fileDescriptor_503b65ce534fc241, []int{0}
}
func (m *WrappedRelationships) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedRelationships) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedRelationships.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedRelationships) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedRelationships.Merge(m, src)
}
func (m *WrappedRelationships) XXX_Size() int {
	return m.Size()
}
func (m *WrappedRelationships) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedRelationships.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedRelationships proto.InternalMessageInfo

func (m *WrappedRelationships) GetRelationships() []types.Relationship {
	if m != nil {
		return m.Relationships
	}
	return nil
}

// WrappedUserBlocks wraps a list of UserBlock objects
type WrappedUserBlocks struct {
	Blocks []types.UserBlock `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks"`
}

func (m *WrappedUserBlocks) Reset()         { *m = WrappedUserBlocks{} }
func (m *WrappedUserBlocks) String() string { return proto.CompactTextString(m) }
func (*WrappedUserBlocks) ProtoMessage()    {}
func (*WrappedUserBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_503b65ce534fc241, []int{1}
}
func (m *WrappedUserBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedUserBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedUserBlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedUserBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedUserBlocks.Merge(m, src)
}
func (m *WrappedUserBlocks) XXX_Size() int {
	return m.Size()
}
func (m *WrappedUserBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedUserBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedUserBlocks proto.InternalMessageInfo

func (m *WrappedUserBlocks) GetBlocks() []types.UserBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*WrappedRelationships)(nil), "desmos.relationships.v1beta1.keeper.WrappedRelationships")
	proto.RegisterType((*WrappedUserBlocks)(nil), "desmos.relationships.v1beta1.keeper.WrappedUserBlocks")
}

func init() {
	proto.RegisterFile("desmos/relationships/v1beta1/keeper/types.proto", fileDescriptor_503b65ce534fc241)
}

var fileDescriptor_503b65ce534fc241 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0x49, 0x2d, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0x4a, 0xcd, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xce, 0xc8, 0x2c, 0x28,
	0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0x4e, 0x4d, 0x2d, 0x48, 0x2d, 0xd2,
	0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x86, 0x68, 0xd0,
	0x43, 0xd1, 0xa0, 0x07, 0xd5, 0xa0, 0x07, 0xd1, 0x20, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56,
	0xaf, 0x0f, 0x62, 0x41, 0xb4, 0x4a, 0x69, 0xe2, 0xb5, 0x2b, 0x37, 0x3f, 0x25, 0x35, 0x07, 0x6a,
	0x8b, 0x52, 0x1e, 0x97, 0x48, 0x78, 0x51, 0x62, 0x41, 0x41, 0x6a, 0x4a, 0x10, 0xb2, 0x62, 0xa1,
	0x30, 0x2e, 0x5e, 0x14, 0xdd, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x5a, 0x7a, 0x78, 0x5d,
	0x85, 0x6c, 0x86, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0xa8, 0xc6, 0x28, 0x45, 0x71, 0x09,
	0x42, 0xed, 0x0b, 0x2d, 0x4e, 0x2d, 0x72, 0xca, 0xc9, 0x4f, 0xce, 0x2e, 0x16, 0x72, 0xe5, 0x62,
	0x4b, 0x02, 0xb3, 0xa0, 0xb6, 0xa8, 0xe3, 0xb7, 0x05, 0xae, 0x13, 0x6a, 0x05, 0x54, 0xb3, 0x93,
	0xdf, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0xa4, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0x42, 0xe3, 0x41, 0x37, 0x27, 0x31, 0xa9, 0x18, 0x16, 0x27,
	0x15, 0x68, 0x21, 0x05, 0x09, 0xdc, 0x24, 0x36, 0x70, 0x10, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x76, 0xe4, 0x7e, 0x12, 0xbb, 0x01, 0x00, 0x00,
}

func (m *WrappedRelationships) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedRelationships) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedRelationships) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relationships) > 0 {
		for iNdEx := len(m.Relationships) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Relationships[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *WrappedUserBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedUserBlocks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedUserBlocks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Blocks) > 0 {
		for iNdEx := len(m.Blocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Blocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *WrappedRelationships) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Relationships) > 0 {
		for _, e := range m.Relationships {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *WrappedUserBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Blocks) > 0 {
		for _, e := range m.Blocks {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WrappedRelationships) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WrappedRelationships: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedRelationships: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relationships", wireType)
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
			m.Relationships = append(m.Relationships, types.Relationship{})
			if err := m.Relationships[len(m.Relationships)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
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
func (m *WrappedUserBlocks) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WrappedUserBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedUserBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
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
			m.Blocks = append(m.Blocks, types.UserBlock{})
			if err := m.Blocks[len(m.Blocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
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