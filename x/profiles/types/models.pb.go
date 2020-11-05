// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/models.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// Pictures contains the data of a user profile's related pictures
type Pictures struct {
	Profile string `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty" yaml"profile"`
	Cover   string `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty" yaml"cover"`
}

func (m *Pictures) Reset()         { *m = Pictures{} }
func (m *Pictures) String() string { return proto.CompactTextString(m) }
func (*Pictures) ProtoMessage()    {}
func (*Pictures) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbf89d5cd53a9ca, []int{0}
}
func (m *Pictures) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pictures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pictures.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pictures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pictures.Merge(m, src)
}
func (m *Pictures) XXX_Size() int {
	return m.Size()
}
func (m *Pictures) XXX_DiscardUnknown() {
	xxx_messageInfo_Pictures.DiscardUnknown(m)
}

var xxx_messageInfo_Pictures proto.InternalMessageInfo

func (m *Pictures) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *Pictures) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

// Profile represents a generic first on Desmos, containing the information of a single user
type Profile struct {
	Dtag         string    `protobuf:"bytes,1,opt,name=dtag,proto3" json:"dtag,omitempty" yaml:"dtag"`
	Moniker      string    `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty" yaml:"moniker"`
	Bio          string    `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty" yaml:"bio"`
	Pictures     Pictures  `protobuf:"bytes,4,opt,name=pictures,proto3" json:"pictures" yaml:"pictures"`
	Creator      string    `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	CreationDate time.Time `protobuf:"bytes,6,opt,name=creation_date,json=creationDate,proto3,stdtime" json:"creation_date" yaml:"creation_date"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbf89d5cd53a9ca, []int{1}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return m.Size()
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetDtag() string {
	if m != nil {
		return m.Dtag
	}
	return ""
}

func (m *Profile) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Profile) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *Profile) GetPictures() Pictures {
	if m != nil {
		return m.Pictures
	}
	return Pictures{}
}

func (m *Profile) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Profile) GetCreationDate() time.Time {
	if m != nil {
		return m.CreationDate
	}
	return time.Time{}
}

// DTagTransferRequest represent a dtag transfer request between two users
type DTagTransferRequest struct {
	DtagToTrade string `protobuf:"bytes,1,opt,name=dtag_to_trade,json=dtagToTrade,proto3" json:"dtag_to_trade,omitempty" yaml:"dtag_to_trade"`
	Receiver    string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty" yaml:"receiver"`
	Sender      string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
}

func (m *DTagTransferRequest) Reset()         { *m = DTagTransferRequest{} }
func (m *DTagTransferRequest) String() string { return proto.CompactTextString(m) }
func (*DTagTransferRequest) ProtoMessage()    {}
func (*DTagTransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbf89d5cd53a9ca, []int{2}
}
func (m *DTagTransferRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DTagTransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DTagTransferRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DTagTransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTagTransferRequest.Merge(m, src)
}
func (m *DTagTransferRequest) XXX_Size() int {
	return m.Size()
}
func (m *DTagTransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DTagTransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DTagTransferRequest proto.InternalMessageInfo

func (m *DTagTransferRequest) GetDtagToTrade() string {
	if m != nil {
		return m.DtagToTrade
	}
	return ""
}

func (m *DTagTransferRequest) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *DTagTransferRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func init() {
	proto.RegisterType((*Pictures)(nil), "desmos.profiles.v1beta1.Pictures")
	proto.RegisterType((*Profile)(nil), "desmos.profiles.v1beta1.Profile")
	proto.RegisterType((*DTagTransferRequest)(nil), "desmos.profiles.v1beta1.DTagTransferRequest")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/models.proto", fileDescriptor_dcbf89d5cd53a9ca)
}

var fileDescriptor_dcbf89d5cd53a9ca = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0x13, 0x3f,
	0x14, 0xc5, 0x33, 0xfd, 0x48, 0xf2, 0x77, 0xfe, 0x69, 0x85, 0x5b, 0xa9, 0xa3, 0x2c, 0xc6, 0xc1,
	0x80, 0x54, 0x04, 0x78, 0x54, 0xd8, 0x45, 0xac, 0xa2, 0x2e, 0x58, 0x56, 0x56, 0xc4, 0x82, 0x4d,
	0xe4, 0xc9, 0x38, 0xd3, 0x11, 0x99, 0x38, 0xd8, 0x4e, 0x45, 0xdf, 0xa2, 0x4b, 0x96, 0x7d, 0x0d,
	0xde, 0xa0, 0xcb, 0x2e, 0x59, 0x0d, 0x28, 0xd9, 0xb0, 0x25, 0x4f, 0x80, 0xc6, 0x1f, 0x99, 0x82,
	0xc4, 0xce, 0xbe, 0xe7, 0xe7, 0x7b, 0x6f, 0xce, 0xc9, 0x80, 0xa7, 0x29, 0x57, 0x85, 0x50, 0xf1,
	0x42, 0x8a, 0x69, 0x3e, 0xe3, 0x2a, 0xbe, 0x3a, 0x4b, 0xb8, 0x66, 0x67, 0x71, 0x21, 0x52, 0x3e,
	0x53, 0x64, 0x21, 0x85, 0x16, 0xf0, 0xc4, 0x52, 0xc4, 0x53, 0xc4, 0x51, 0xbd, 0xe3, 0x4c, 0x64,
	0xc2, 0x30, 0x71, 0x75, 0xb2, 0x78, 0x0f, 0x65, 0x42, 0x64, 0x33, 0x1e, 0x9b, 0x5b, 0xb2, 0x9c,
	0xc6, 0x3a, 0x2f, 0xb8, 0xd2, 0xac, 0x58, 0x58, 0x00, 0x5f, 0x82, 0xf6, 0x45, 0x3e, 0xd1, 0x4b,
	0xc9, 0x15, 0x7c, 0x01, 0x5a, 0xae, 0x6d, 0x18, 0xf4, 0x83, 0xd3, 0xff, 0x86, 0x8f, 0x36, 0x25,
	0xea, 0x5e, 0xb3, 0x62, 0x86, 0x5d, 0x1d, 0x53, 0x4f, 0xc0, 0x67, 0x60, 0x7f, 0x22, 0xae, 0xb8,
	0x0c, 0x77, 0x0c, 0x7a, 0xb8, 0x29, 0x51, 0xc7, 0xa0, 0xa6, 0x8a, 0xa9, 0x55, 0x07, 0xed, 0x2f,
	0xb7, 0x28, 0xf8, 0x79, 0x8b, 0x02, 0xfc, 0x6b, 0x07, 0xb4, 0x2e, 0xdc, 0xe3, 0x27, 0x60, 0x2f,
	0xd5, 0x2c, 0x73, 0x63, 0xb6, 0x6f, 0x07, 0xb8, 0xaa, 0x62, 0x6a, 0x44, 0xf8, 0x12, 0xb4, 0x0a,
	0x31, 0xcf, 0x3f, 0x6e, 0x67, 0xc0, 0x4d, 0x89, 0x0e, 0x2c, 0xe7, 0x04, 0x4c, 0x3d, 0x02, 0xfb,
	0x60, 0x37, 0xc9, 0x45, 0xb8, 0x6b, 0xc8, 0x83, 0x4d, 0x89, 0x80, 0x25, 0x93, 0x5c, 0x60, 0x5a,
	0x49, 0xf0, 0x3d, 0x68, 0x2f, 0xdc, 0x4f, 0x0d, 0xf7, 0xfa, 0xc1, 0x69, 0xe7, 0xf5, 0x63, 0xf2,
	0x0f, 0x37, 0x89, 0xf7, 0x64, 0x78, 0x72, 0x57, 0xa2, 0xc6, 0xa6, 0x44, 0x87, 0xb6, 0x9b, 0x6f,
	0x80, 0xe9, 0xb6, 0x57, 0xb5, 0xe7, 0x44, 0x72, 0xa6, 0x85, 0x0c, 0xf7, 0xff, 0xde, 0xd3, 0x09,
	0x98, 0x7a, 0x04, 0x32, 0xd0, 0x35, 0xc7, 0x5c, 0xcc, 0xc7, 0x29, 0xd3, 0x3c, 0x6c, 0x9a, 0x55,
	0x7a, 0xc4, 0x26, 0x45, 0x7c, 0x52, 0x64, 0xe4, 0x93, 0x1a, 0xf6, 0xdd, 0x0e, 0xc7, 0x0f, 0x7a,
	0xfa, 0xe7, 0xf8, 0xe6, 0x3b, 0x0a, 0xe8, 0xff, 0xbe, 0x76, 0xce, 0x34, 0x7f, 0xe0, 0xf9, 0xd7,
	0x00, 0x1c, 0x9d, 0x8f, 0x58, 0x36, 0x92, 0x6c, 0xae, 0xa6, 0x5c, 0x52, 0xfe, 0x69, 0xc9, 0x95,
	0x86, 0x6f, 0x41, 0xb7, 0xb2, 0x78, 0xac, 0xc5, 0x58, 0x4b, 0x96, 0xfa, 0xbc, 0xc3, 0x7a, 0xc8,
	0x1f, 0x32, 0xa6, 0x9d, 0xea, 0x3e, 0x12, 0xa3, 0xea, 0x06, 0x63, 0xd0, 0x96, 0x7c, 0xc2, 0xf3,
	0x3a, 0xfd, 0xa3, 0xda, 0x21, 0xaf, 0x60, 0xba, 0x85, 0xe0, 0x73, 0xd0, 0x54, 0x7c, 0x9e, 0x72,
	0xe9, 0xe2, 0xd9, 0xfe, 0xaf, 0x06, 0xd8, 0xd6, 0x31, 0x75, 0x40, 0xbd, 0xfb, 0xf0, 0xdd, 0xdd,
	0x2a, 0x0a, 0xee, 0x57, 0x51, 0xf0, 0x63, 0x15, 0x05, 0x37, 0xeb, 0xa8, 0x71, 0xbf, 0x8e, 0x1a,
	0xdf, 0xd6, 0x51, 0xe3, 0x03, 0xc9, 0x72, 0x7d, 0xb9, 0x4c, 0xc8, 0x44, 0x14, 0xb1, 0x0d, 0xf0,
	0xd5, 0x8c, 0x25, 0xca, 0x9d, 0xe3, 0xcf, 0xf5, 0x27, 0xa4, 0xaf, 0x17, 0x5c, 0x25, 0x4d, 0xe3,
	0xe9, 0x9b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x35, 0x30, 0x85, 0x62, 0x03, 0x00, 0x00,
}

func (this *Pictures) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pictures)
	if !ok {
		that2, ok := that.(Pictures)
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
	if this.Profile != that1.Profile {
		return false
	}
	if this.Cover != that1.Cover {
		return false
	}
	return true
}
func (this *Profile) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Profile)
	if !ok {
		that2, ok := that.(Profile)
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
	if this.Dtag != that1.Dtag {
		return false
	}
	if this.Moniker != that1.Moniker {
		return false
	}
	if this.Bio != that1.Bio {
		return false
	}
	if !this.Pictures.Equal(&that1.Pictures) {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if !this.CreationDate.Equal(that1.CreationDate) {
		return false
	}
	return true
}
func (this *DTagTransferRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DTagTransferRequest)
	if !ok {
		that2, ok := that.(DTagTransferRequest)
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
	if this.DtagToTrade != that1.DtagToTrade {
		return false
	}
	if this.Receiver != that1.Receiver {
		return false
	}
	if this.Sender != that1.Sender {
		return false
	}
	return true
}
func (m *Pictures) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pictures) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pictures) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cover) > 0 {
		i -= len(m.Cover)
		copy(dAtA[i:], m.Cover)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Cover)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Profile) > 0 {
		i -= len(m.Profile)
		copy(dAtA[i:], m.Profile)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Profile)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Profile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Profile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Profile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreationDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationDate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintModels(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.Pictures.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModels(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Bio) > 0 {
		i -= len(m.Bio)
		copy(dAtA[i:], m.Bio)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Bio)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Dtag) > 0 {
		i -= len(m.Dtag)
		copy(dAtA[i:], m.Dtag)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Dtag)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DTagTransferRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DTagTransferRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DTagTransferRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DtagToTrade) > 0 {
		i -= len(m.DtagToTrade)
		copy(dAtA[i:], m.DtagToTrade)
		i = encodeVarintModels(dAtA, i, uint64(len(m.DtagToTrade)))
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
func (m *Pictures) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Profile)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Cover)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	return n
}

func (m *Profile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Dtag)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Bio)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = m.Pictures.Size()
	n += 1 + l + sovModels(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationDate)
	n += 1 + l + sovModels(uint64(l))
	return n
}

func (m *DTagTransferRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DtagToTrade)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Sender)
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
func (m *Pictures) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Pictures: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pictures: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
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
			m.Profile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cover", wireType)
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
			m.Cover = string(dAtA[iNdEx:postIndex])
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
func (m *Profile) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Profile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtag", wireType)
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
			m.Dtag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
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
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bio", wireType)
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
			m.Bio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pictures", wireType)
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
			if err := m.Pictures.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationDate", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreationDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *DTagTransferRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DTagTransferRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DTagTransferRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DtagToTrade", wireType)
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
			m.DtagToTrade = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
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
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
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
