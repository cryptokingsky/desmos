// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/magpie/v1beta1/models.proto

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

// SessionID represents a unique session id
type SessionID struct {
	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *SessionID) Reset()         { *m = SessionID{} }
func (m *SessionID) String() string { return proto.CompactTextString(m) }
func (*SessionID) ProtoMessage()    {}
func (*SessionID) Descriptor() ([]byte, []int) {
	return fileDescriptor_19ad13bf6c2fd1ac, []int{0}
}
func (m *SessionID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SessionID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionID.Merge(m, src)
}
func (m *SessionID) XXX_Size() int {
	return m.Size()
}
func (m *SessionID) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionID.DiscardUnknown(m)
}

var xxx_messageInfo_SessionID proto.InternalMessageInfo

func (m *SessionID) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Session struct {
	SessionId      SessionID `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id" yaml:"session_id"`
	Owner          string    `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	CreationTime   uint64    `protobuf:"varint,3,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty" yaml:"creation_time"`
	ExpirationTime uint64    `protobuf:"varint,4,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty" yaml:"expiration_time"`
	Namespace      string    `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty" yaml:"namespace"`
	ExternalOwner  string    `protobuf:"bytes,6,opt,name=external_owner,json=externalOwner,proto3" json:"external_owner,omitempty" yaml:"external_owner"`
	PublicKey      string    `protobuf:"bytes,7,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty" yaml:"pub_key"`
	Signature      string    `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty" yaml:"signature"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_19ad13bf6c2fd1ac, []int{1}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Session.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return m.Size()
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetSessionId() SessionID {
	if m != nil {
		return m.SessionId
	}
	return SessionID{}
}

func (m *Session) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Session) GetCreationTime() uint64 {
	if m != nil {
		return m.CreationTime
	}
	return 0
}

func (m *Session) GetExpirationTime() uint64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *Session) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Session) GetExternalOwner() string {
	if m != nil {
		return m.ExternalOwner
	}
	return ""
}

func (m *Session) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *Session) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionID)(nil), "desmos.magpie.v1beta1.SessionID")
	proto.RegisterType((*Session)(nil), "desmos.magpie.v1beta1.Session")
}

func init() {
	proto.RegisterFile("desmos/magpie/v1beta1/models.proto", fileDescriptor_19ad13bf6c2fd1ac)
}

var fileDescriptor_19ad13bf6c2fd1ac = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xb1, 0x6f, 0xd3, 0x40,
	0x14, 0xc6, 0x63, 0xda, 0xb4, 0xf5, 0xd1, 0x86, 0x62, 0xa5, 0xc8, 0xed, 0xe0, 0x8b, 0x6e, 0x40,
	0x95, 0x00, 0x5b, 0x29, 0x5b, 0x25, 0x24, 0x14, 0x10, 0x52, 0xc5, 0x80, 0x64, 0x98, 0xba, 0x44,
	0xe7, 0xe4, 0xc9, 0x9c, 0xf0, 0xf9, 0x2c, 0xdf, 0xb9, 0x24, 0xff, 0x05, 0x23, 0x63, 0xff, 0x9c,
	0x8e, 0x1d, 0x99, 0x4e, 0x28, 0x59, 0x98, 0x3d, 0x31, 0x22, 0xdf, 0xc5, 0x49, 0xa9, 0xba, 0xbd,
	0xbb, 0xef, 0xf7, 0x7d, 0xef, 0x9d, 0xde, 0x21, 0x32, 0x05, 0xc9, 0x85, 0x8c, 0x38, 0x4d, 0x0b,
	0x06, 0xd1, 0xd5, 0x30, 0x01, 0x45, 0x87, 0x11, 0x17, 0x53, 0xc8, 0x64, 0x58, 0x94, 0x42, 0x09,
	0xef, 0xc8, 0x32, 0xa1, 0x65, 0xc2, 0x15, 0x73, 0xd2, 0x4f, 0x45, 0x2a, 0x0c, 0x11, 0x35, 0x95,
	0x85, 0xc9, 0x0b, 0xe4, 0x7e, 0x06, 0x29, 0x99, 0xc8, 0x2f, 0xde, 0x7b, 0x7d, 0xd4, 0xbd, 0xa2,
	0x59, 0x05, 0xbe, 0x33, 0x70, 0x4e, 0xb7, 0x63, 0x7b, 0x38, 0xdf, 0xfb, 0x79, 0x8d, 0x9d, 0x3f,
	0xd7, 0xd8, 0x21, 0x7f, 0xb7, 0xd0, 0xee, 0x8a, 0xf6, 0x2e, 0x11, 0x92, 0xb6, 0x1c, 0xb3, 0xa9,
	0x31, 0x3c, 0x3e, 0x1b, 0x84, 0x0f, 0xb6, 0x0e, 0xd7, 0x1d, 0x46, 0xc7, 0x37, 0x1a, 0x77, 0x6a,
	0x8d, 0x9f, 0xce, 0x29, 0xcf, 0xce, 0xc9, 0x26, 0x81, 0xc4, 0xee, 0xea, 0x70, 0x31, 0xf5, 0x9e,
	0xa3, 0xae, 0xf8, 0x9e, 0x43, 0xe9, 0x3f, 0x1a, 0x38, 0xa7, 0xee, 0xe8, 0xb0, 0xd6, 0x78, 0xdf,
	0x1a, 0xcc, 0x35, 0x89, 0xad, 0xec, 0xbd, 0x41, 0x07, 0x93, 0x12, 0xa8, 0x6a, 0x22, 0x14, 0xe3,
	0xe0, 0x6f, 0x35, 0x73, 0x8f, 0xfc, 0x5a, 0xe3, 0xbe, 0xe5, 0xff, 0x93, 0x49, 0xbc, 0xdf, 0x9e,
	0xbf, 0x30, 0x0e, 0xde, 0x3b, 0xf4, 0x04, 0x66, 0x05, 0x2b, 0xef, 0x04, 0x6c, 0x9b, 0x80, 0x93,
	0x5a, 0xe3, 0x67, 0x36, 0xe0, 0x1e, 0x40, 0xe2, 0xde, 0xe6, 0xc6, 0x84, 0x9c, 0x21, 0x37, 0xa7,
	0x1c, 0x64, 0x41, 0x27, 0xe0, 0x77, 0xcd, 0xbc, 0xfd, 0x5a, 0xe3, 0x43, 0x6b, 0x5f, 0x4b, 0x24,
	0xde, 0x60, 0xde, 0x5b, 0xd4, 0x83, 0x99, 0x82, 0x32, 0xa7, 0xd9, 0xd8, 0x3e, 0x74, 0xc7, 0x18,
	0x8f, 0x6b, 0x8d, 0x8f, 0xda, 0xbe, 0x77, 0x75, 0x12, 0x1f, 0xb4, 0x17, 0x9f, 0xcc, 0xcb, 0x87,
	0x08, 0x15, 0x55, 0x92, 0xb1, 0xc9, 0xf8, 0x1b, 0xcc, 0xfd, 0x5d, 0xe3, 0xf6, 0x6a, 0x8d, 0x7b,
	0xd6, 0x5d, 0x54, 0x49, 0x23, 0x90, 0xd8, 0xb5, 0xd4, 0x47, 0x98, 0x37, 0x83, 0x4a, 0x96, 0xe6,
	0x54, 0x55, 0x25, 0xf8, 0x7b, 0xf7, 0x07, 0x5d, 0x4b, 0xcd, 0x22, 0xda, 0x7a, 0xb3, 0xfa, 0xd1,
	0x87, 0x9b, 0x45, 0xe0, 0xdc, 0x2e, 0x02, 0xe7, 0xf7, 0x22, 0x70, 0x7e, 0x2c, 0x83, 0xce, 0xed,
	0x32, 0xe8, 0xfc, 0x5a, 0x06, 0x9d, 0xcb, 0x97, 0x29, 0x53, 0x5f, 0xab, 0x24, 0x9c, 0x08, 0x1e,
	0xd9, 0xf5, 0xbf, 0xca, 0x68, 0x22, 0x57, 0x75, 0x34, 0x6b, 0xff, 0xaa, 0x9a, 0x17, 0x20, 0x93,
	0x1d, 0xf3, 0xed, 0x5e, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xfa, 0x30, 0xff, 0xc9, 0x02,
	0x00, 0x00,
}

func (this *SessionID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SessionID)
	if !ok {
		that2, ok := that.(SessionID)
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
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *Session) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Session)
	if !ok {
		that2, ok := that.(Session)
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
	if !this.SessionId.Equal(&that1.SessionId) {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.CreationTime != that1.CreationTime {
		return false
	}
	if this.ExpirationTime != that1.ExpirationTime {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.ExternalOwner != that1.ExternalOwner {
		return false
	}
	if this.PublicKey != that1.PublicKey {
		return false
	}
	if this.Signature != that1.Signature {
		return false
	}
	return true
}
func (m *SessionID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SessionID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Session) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintModels(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ExternalOwner) > 0 {
		i -= len(m.ExternalOwner)
		copy(dAtA[i:], m.ExternalOwner)
		i = encodeVarintModels(dAtA, i, uint64(len(m.ExternalOwner)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ExpirationTime != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.ExpirationTime))
		i--
		dAtA[i] = 0x20
	}
	if m.CreationTime != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.CreationTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.SessionId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModels(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *SessionID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovModels(uint64(m.Value))
	}
	return n
}

func (m *Session) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SessionId.Size()
	n += 1 + l + sovModels(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if m.CreationTime != 0 {
		n += 1 + sovModels(uint64(m.CreationTime))
	}
	if m.ExpirationTime != 0 {
		n += 1 + sovModels(uint64(m.ExpirationTime))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.ExternalOwner)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Signature)
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
func (m *SessionID) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SessionID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *Session) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SessionId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			m.CreationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			m.ExpirationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
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
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalOwner", wireType)
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
			m.ExternalOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
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
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
			m.Signature = string(dAtA[iNdEx:postIndex])
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