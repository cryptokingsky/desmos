// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/relationships/v1beta1/models.proto

package types

import (
	fmt "fmt"
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

// Relationship is the struct of a relationship.
// It represent the concept of "follow" of traditional social networks.
type Relationship struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
	Subspace  string `protobuf:"bytes,3,opt,name=subspace,proto3" json:"subspace,omitempty" yaml:"subspace"`
}

func (m *Relationship) Reset()         { *m = Relationship{} }
func (m *Relationship) String() string { return proto.CompactTextString(m) }
func (*Relationship) ProtoMessage()    {}
func (*Relationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_8511c516358f0efc, []int{0}
}
func (m *Relationship) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Relationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Relationship.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Relationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relationship.Merge(m, src)
}
func (m *Relationship) XXX_Size() int {
	return m.Size()
}
func (m *Relationship) XXX_DiscardUnknown() {
	xxx_messageInfo_Relationship.DiscardUnknown(m)
}

var xxx_messageInfo_Relationship proto.InternalMessageInfo

func (m *Relationship) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Relationship) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Relationship) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

// UserBlock represents the fact that the Blocker has blocked the given Blocked user.
// The Reason field represents the reason the user has been blocked for, and is optional.
type UserBlock struct {
	Blocker  string `protobuf:"bytes,1,opt,name=blocker,proto3" json:"blocker,omitempty" yaml:"blocker"`
	Blocked  string `protobuf:"bytes,2,opt,name=blocked,proto3" json:"blocked,omitempty" yaml:"blocked"`
	Reason   string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty" yaml:"reason"`
	Subspace string `protobuf:"bytes,4,opt,name=subspace,proto3" json:"subspace,omitempty" yaml:"subspace"`
}

func (m *UserBlock) Reset()         { *m = UserBlock{} }
func (m *UserBlock) String() string { return proto.CompactTextString(m) }
func (*UserBlock) ProtoMessage()    {}
func (*UserBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_8511c516358f0efc, []int{1}
}
func (m *UserBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBlock.Merge(m, src)
}
func (m *UserBlock) XXX_Size() int {
	return m.Size()
}
func (m *UserBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBlock.DiscardUnknown(m)
}

var xxx_messageInfo_UserBlock proto.InternalMessageInfo

func (m *UserBlock) GetBlocker() string {
	if m != nil {
		return m.Blocker
	}
	return ""
}

func (m *UserBlock) GetBlocked() string {
	if m != nil {
		return m.Blocked
	}
	return ""
}

func (m *UserBlock) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *UserBlock) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

func init() {
	proto.RegisterType((*Relationship)(nil), "desmos.relationships.v1beta1.Relationship")
	proto.RegisterType((*UserBlock)(nil), "desmos.relationships.v1beta1.UserBlock")
}

func init() {
	proto.RegisterFile("desmos/relationships/v1beta1/models.proto", fileDescriptor_8511c516358f0efc)
}

var fileDescriptor_8511c516358f0efc = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0x06, 0xf0, 0x9e, 0x4a, 0x6d, 0x0f, 0xff, 0xc6, 0x0e, 0x45, 0x24, 0x91, 0x9b, 0x2c, 0x68,
	0x8e, 0xda, 0xad, 0x63, 0x77, 0x97, 0x80, 0x8b, 0xdb, 0x5d, 0xf2, 0xd2, 0x06, 0x93, 0x5e, 0xb8,
	0xbb, 0x8a, 0xfd, 0x16, 0x8e, 0x8e, 0xdd, 0xfd, 0x22, 0x2e, 0x42, 0x47, 0xa7, 0x20, 0xed, 0xe2,
	0x9c, 0x4f, 0x20, 0x49, 0x2e, 0xad, 0x55, 0x17, 0xb7, 0x97, 0x3c, 0xbf, 0x17, 0x9e, 0x97, 0x0b,
	0xee, 0x04, 0xa0, 0x62, 0xa1, 0xa8, 0x84, 0x88, 0xe9, 0x50, 0x8c, 0xd5, 0x28, 0x4c, 0x14, 0x7d,
	0xe8, 0x72, 0xd0, 0xac, 0x4b, 0x63, 0x11, 0x40, 0xa4, 0xdc, 0x44, 0x0a, 0x2d, 0xac, 0xb3, 0x92,
	0xba, 0x1b, 0xd4, 0x35, 0xf4, 0xb4, 0x35, 0x14, 0x43, 0x51, 0x40, 0x9a, 0x4f, 0xe5, 0x0e, 0x79,
	0x41, 0x78, 0xcf, 0xfb, 0xe6, 0xad, 0x4b, 0xbc, 0xeb, 0x4b, 0x60, 0x5a, 0xc8, 0x36, 0x3a, 0x47,
	0x17, 0xcd, 0x81, 0x95, 0xa5, 0xce, 0xc1, 0x94, 0xc5, 0x51, 0x9f, 0x98, 0x80, 0x78, 0x15, 0xb1,
	0xae, 0x71, 0x53, 0x82, 0x1f, 0x26, 0x21, 0x8c, 0x75, 0x7b, 0xab, 0xf0, 0xad, 0x2c, 0x75, 0x8e,
	0x4a, 0xbf, 0x8a, 0x88, 0xb7, 0x66, 0x16, 0xc5, 0x0d, 0x35, 0xe1, 0x2a, 0x61, 0x3e, 0xb4, 0xb7,
	0x8b, 0x95, 0x93, 0x2c, 0x75, 0x0e, 0xcb, 0x95, 0x2a, 0x21, 0xde, 0x0a, 0xf5, 0x1b, 0xcf, 0x33,
	0x07, 0x7d, 0xce, 0x1c, 0x44, 0xde, 0x10, 0x6e, 0xde, 0x2a, 0x90, 0x83, 0x48, 0xf8, 0xf7, 0x79,
	0x55, 0x9e, 0x0f, 0xf0, 0x47, 0x55, 0x13, 0x10, 0xaf, 0x22, 0x6b, 0x1d, 0x98, 0xa2, 0xbf, 0x74,
	0xb0, 0xd2, 0x81, 0xd5, 0xc1, 0x75, 0x09, 0x4c, 0x89, 0xb1, 0xa9, 0x78, 0x9c, 0xa5, 0xce, 0x7e,
	0x75, 0x55, 0xfe, 0x9d, 0x78, 0x06, 0x6c, 0xdc, 0xb3, 0xf3, 0xaf, 0x7b, 0x06, 0x37, 0xaf, 0x0b,
	0x1b, 0xcd, 0x17, 0x36, 0xfa, 0x58, 0xd8, 0xe8, 0x69, 0x69, 0xd7, 0xe6, 0x4b, 0xbb, 0xf6, 0xbe,
	0xb4, 0x6b, 0x77, 0xbd, 0x61, 0xa8, 0x47, 0x13, 0xee, 0xfa, 0x22, 0xa6, 0xe5, 0xb3, 0x5e, 0x45,
	0x8c, 0x2b, 0x33, 0xd3, 0xc7, 0x1f, 0xff, 0x83, 0x9e, 0x26, 0xa0, 0x78, 0xbd, 0x78, 0xd3, 0xde,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xf4, 0x09, 0x8f, 0x34, 0x02, 0x00, 0x00,
}

func (this *Relationship) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Relationship)
	if !ok {
		that2, ok := that.(Relationship)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if this.Recipient != that1.Recipient {
		return false
	}
	if this.Subspace != that1.Subspace {
		return false
	}
	return true
}
func (this *UserBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserBlock)
	if !ok {
		that2, ok := that.(UserBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Blocker != that1.Blocker {
		return false
	}
	if this.Blocked != that1.Blocked {
		return false
	}
	if this.Reason != that1.Reason {
		return false
	}
	if this.Subspace != that1.Subspace {
		return false
	}
	return true
}
func (m *Relationship) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Relationship) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Relationship) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Blocked) > 0 {
		i -= len(m.Blocked)
		copy(dAtA[i:], m.Blocked)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Blocked)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Blocker) > 0 {
		i -= len(m.Blocker)
		copy(dAtA[i:], m.Blocker)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Blocker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModels(dAtA []byte, offset int, v uint64) int {
	offset -= sovModels(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Relationship) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	return n
}

func (m *UserBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Blocker)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Blocked)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	return n
}

func sovModels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModels(x uint64) (n int) {
	return sovModels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Relationship) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: Relationship: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Relationship: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModels
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
func (m *UserBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: UserBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocked = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModels
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
func skipModels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
				return 0, ErrInvalidLengthModels
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModels
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModels
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModels        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModels          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModels = fmt.Errorf("proto: unexpected end of group")
)