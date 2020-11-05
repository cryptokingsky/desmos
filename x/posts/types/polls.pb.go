// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/posts/v1beta1/polls.proto

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

// PollAnswer contains the data of a single poll answer inserted by the creator
type PollAnswer struct {
	ID   string `protobuf:"bytes,1,opt,name=answer_id,json=answerId,proto3" json:"id" yaml"id"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text" yaml"text"`
}

func (m *PollAnswer) Reset()         { *m = PollAnswer{} }
func (m *PollAnswer) String() string { return proto.CompactTextString(m) }
func (*PollAnswer) ProtoMessage()    {}
func (*PollAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_397fcd2757705694, []int{0}
}
func (m *PollAnswer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PollAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PollAnswer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PollAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollAnswer.Merge(m, src)
}
func (m *PollAnswer) XXX_Size() int {
	return m.Size()
}
func (m *PollAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_PollAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_PollAnswer proto.InternalMessageInfo

func (m *PollAnswer) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PollAnswer) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// PollAnswer contains the data of a single poll answer inserted by the creator inside a PollData object
type PollData struct {
	Question              string       `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	ProvidedAnswers       []PollAnswer `protobuf:"bytes,2,rep,name=provided_answers,json=providedAnswers,proto3" json:"provided_answers" yaml"provided_answers"`
	EndDate               time.Time    `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3,stdtime" json:"end_date" yaml"end_date"`
	AllowsMultipleAnswers bool         `protobuf:"varint,4,opt,name=allows_multiple_answers,json=allowsMultipleAnswers,proto3" json:"allows_multiple_answers,omitempty" yaml"allows_multiple_answers"`
	AllowsAnswerEdits     bool         `protobuf:"varint,5,opt,name=allows_answer_edits,json=allowsAnswerEdits,proto3" json:"allows_answer_edits,omitempty" yaml"allows_answer_edits"`
}

func (m *PollData) Reset()         { *m = PollData{} }
func (m *PollData) String() string { return proto.CompactTextString(m) }
func (*PollData) ProtoMessage()    {}
func (*PollData) Descriptor() ([]byte, []int) {
	return fileDescriptor_397fcd2757705694, []int{1}
}
func (m *PollData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PollData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PollData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PollData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollData.Merge(m, src)
}
func (m *PollData) XXX_Size() int {
	return m.Size()
}
func (m *PollData) XXX_DiscardUnknown() {
	xxx_messageInfo_PollData.DiscardUnknown(m)
}

var xxx_messageInfo_PollData proto.InternalMessageInfo

func (m *PollData) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *PollData) GetProvidedAnswers() []PollAnswer {
	if m != nil {
		return m.ProvidedAnswers
	}
	return nil
}

func (m *PollData) GetEndDate() time.Time {
	if m != nil {
		return m.EndDate
	}
	return time.Time{}
}

func (m *PollData) GetAllowsMultipleAnswers() bool {
	if m != nil {
		return m.AllowsMultipleAnswers
	}
	return false
}

func (m *PollData) GetAllowsAnswerEdits() bool {
	if m != nil {
		return m.AllowsAnswerEdits
	}
	return false
}

// UserAnswer contains the data of a user's answer to a poll
type UserAnswer struct {
	User    string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Answers []string `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (m *UserAnswer) Reset()         { *m = UserAnswer{} }
func (m *UserAnswer) String() string { return proto.CompactTextString(m) }
func (*UserAnswer) ProtoMessage()    {}
func (*UserAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_397fcd2757705694, []int{2}
}
func (m *UserAnswer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserAnswer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAnswer.Merge(m, src)
}
func (m *UserAnswer) XXX_Size() int {
	return m.Size()
}
func (m *UserAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_UserAnswer proto.InternalMessageInfo

func (m *UserAnswer) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserAnswer) GetAnswers() []string {
	if m != nil {
		return m.Answers
	}
	return nil
}

func init() {
	proto.RegisterType((*PollAnswer)(nil), "desmos.posts.v1beta1.PollAnswer")
	proto.RegisterType((*PollData)(nil), "desmos.posts.v1beta1.PollData")
	proto.RegisterType((*UserAnswer)(nil), "desmos.posts.v1beta1.UserAnswer")
}

func init() { proto.RegisterFile("desmos/posts/v1beta1/polls.proto", fileDescriptor_397fcd2757705694) }

var fileDescriptor_397fcd2757705694 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0xa6, 0x81, 0x3a, 0x5b, 0x89, 0x82, 0x29, 0xc5, 0x44, 0xaa, 0xd7, 0xf8, 0x80, 0x22,
	0x10, 0xb6, 0x5a, 0xc4, 0x85, 0x1b, 0x51, 0x72, 0xe8, 0xa1, 0x52, 0x65, 0xc1, 0x01, 0x24, 0x14,
	0x6d, 0xd8, 0x25, 0x58, 0x5a, 0x67, 0x8d, 0x77, 0xdd, 0x36, 0xff, 0xa2, 0xe2, 0xc4, 0xb1, 0x3f,
	0xa7, 0x27, 0xd4, 0x23, 0xa7, 0x05, 0x39, 0x17, 0x94, 0xa3, 0x7f, 0x01, 0xf2, 0xae, 0x9d, 0x86,
	0xaf, 0xdb, 0xcc, 0x9b, 0x37, 0x6f, 0x66, 0xe7, 0xd9, 0xd0, 0x23, 0x54, 0x24, 0x5c, 0x84, 0x29,
	0x17, 0x52, 0x84, 0x27, 0xfb, 0x13, 0x2a, 0xf1, 0x7e, 0x98, 0x72, 0xc6, 0x44, 0x90, 0x66, 0x5c,
	0x72, 0x7b, 0xc7, 0x30, 0x02, 0xcd, 0x08, 0x6a, 0x46, 0x6f, 0x67, 0xca, 0xa7, 0x5c, 0x13, 0xc2,
	0x2a, 0x32, 0xdc, 0x1e, 0x9a, 0x72, 0x3e, 0x65, 0x34, 0xd4, 0xd9, 0x24, 0xff, 0x10, 0xca, 0x38,
	0xa1, 0x42, 0xe2, 0x24, 0x35, 0x04, 0x7f, 0x0e, 0xe1, 0x31, 0x67, 0xec, 0xe5, 0x4c, 0x9c, 0xd2,
	0xcc, 0x7e, 0x0e, 0xbb, 0x58, 0x47, 0xe3, 0x98, 0x38, 0xc0, 0x03, 0xfd, 0xee, 0xc0, 0x29, 0x14,
	0x6a, 0x1f, 0x0e, 0x97, 0x0a, 0xb5, 0x63, 0x52, 0x2a, 0x64, 0xcd, 0x71, 0xc2, 0xfc, 0x98, 0xf8,
	0x91, 0x65, 0xa8, 0x87, 0xc4, 0x7e, 0x0c, 0x3b, 0x92, 0x9e, 0x49, 0xa7, 0xad, 0x3b, 0x76, 0x97,
	0x0a, 0xe9, 0xbc, 0x54, 0x08, 0x6a, 0x76, 0x95, 0xf8, 0x91, 0xc6, 0x5e, 0x58, 0x5f, 0x2e, 0x10,
	0xf8, 0x79, 0x81, 0x80, 0xff, 0x75, 0x03, 0x5a, 0xd5, 0xec, 0x21, 0x96, 0xd8, 0xee, 0x41, 0xeb,
	0x53, 0x4e, 0x85, 0x8c, 0xf9, 0xcc, 0x0c, 0x8e, 0x56, 0xb9, 0xfd, 0x19, 0xc0, 0xdb, 0x69, 0xc6,
	0x4f, 0x62, 0x42, 0xc9, 0xd8, 0x0c, 0x15, 0x4e, 0xdb, 0xdb, 0xe8, 0x6f, 0x1d, 0x78, 0xc1, 0xbf,
	0x8e, 0x11, 0x5c, 0x3f, 0x69, 0x30, 0xba, 0x54, 0xa8, 0x55, 0x28, 0xb4, 0x7d, 0x5c, 0x2b, 0x18,
	0x5c, 0x2c, 0x15, 0xfa, 0x4b, 0xb4, 0x54, 0x68, 0x57, 0x2f, 0xfc, 0x67, 0xc1, 0x8f, 0xb6, 0xd3,
	0xdf, 0xdb, 0xed, 0x77, 0xd0, 0xa2, 0x33, 0x32, 0x26, 0x58, 0x52, 0x67, 0xc3, 0x03, 0xfd, 0xad,
	0x83, 0x5e, 0x60, 0x8e, 0x1d, 0x34, 0xc7, 0x0e, 0x5e, 0x35, 0xc7, 0x1e, 0x3c, 0xaa, 0xb6, 0x58,
	0x2a, 0xb4, 0xea, 0x29, 0x15, 0xba, 0xa5, 0x47, 0x35, 0x80, 0x7f, 0xfe, 0x1d, 0x81, 0x68, 0x93,
	0xce, 0xc8, 0x10, 0x4b, 0x6a, 0xbf, 0x81, 0xf7, 0x31, 0x63, 0xfc, 0x54, 0x8c, 0x93, 0x9c, 0xc9,
	0x38, 0x65, 0x74, 0xf5, 0xf2, 0x8e, 0x07, 0xfa, 0xd6, 0xe0, 0x61, 0xa9, 0xd0, 0x9e, 0x56, 0xf8,
	0x0f, 0xcf, 0x8f, 0xee, 0x99, 0xca, 0x51, 0x5d, 0x68, 0x36, 0x3f, 0x82, 0x77, 0xeb, 0x96, 0xda,
	0x6b, 0x4a, 0x62, 0x29, 0x9c, 0x1b, 0x5a, 0x76, 0xaf, 0x54, 0xe8, 0xc1, 0xba, 0xec, 0x3a, 0xc7,
	0x8f, 0xee, 0x18, 0xd4, 0x48, 0x8d, 0x2a, 0x6c, 0xcd, 0xd0, 0x21, 0x84, 0xaf, 0x05, 0xcd, 0xea,
	0x6f, 0xc9, 0x86, 0x9d, 0x5c, 0xd0, 0xac, 0x76, 0x53, 0xc7, 0xb6, 0x03, 0x37, 0xd7, 0xfd, 0xeb,
	0x46, 0x4d, 0x7a, 0xad, 0x32, 0x18, 0x5d, 0x16, 0x2e, 0xb8, 0x2a, 0x5c, 0xf0, 0xa3, 0x70, 0xc1,
	0xf9, 0xc2, 0x6d, 0x5d, 0x2d, 0xdc, 0xd6, 0xb7, 0x85, 0xdb, 0x7a, 0xfb, 0x64, 0x1a, 0xcb, 0x8f,
	0xf9, 0x24, 0x78, 0xcf, 0x93, 0xd0, 0xd8, 0xfe, 0x94, 0xe1, 0x89, 0xa8, 0xe3, 0xf0, 0xac, 0xfe,
	0x67, 0xe4, 0x3c, 0xa5, 0x62, 0x72, 0x53, 0xbb, 0xf0, 0xec, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9c, 0xaf, 0x24, 0x9c, 0x50, 0x03, 0x00, 0x00,
}

func (this *PollAnswer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PollAnswer)
	if !ok {
		that2, ok := that.(PollAnswer)
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
	if this.ID != that1.ID {
		return false
	}
	if this.Text != that1.Text {
		return false
	}
	return true
}
func (this *PollData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PollData)
	if !ok {
		that2, ok := that.(PollData)
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
	if this.Question != that1.Question {
		return false
	}
	if len(this.ProvidedAnswers) != len(that1.ProvidedAnswers) {
		return false
	}
	for i := range this.ProvidedAnswers {
		if !this.ProvidedAnswers[i].Equal(&that1.ProvidedAnswers[i]) {
			return false
		}
	}
	if !this.EndDate.Equal(that1.EndDate) {
		return false
	}
	if this.AllowsMultipleAnswers != that1.AllowsMultipleAnswers {
		return false
	}
	if this.AllowsAnswerEdits != that1.AllowsAnswerEdits {
		return false
	}
	return true
}
func (this *UserAnswer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserAnswer)
	if !ok {
		that2, ok := that.(UserAnswer)
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
	if this.User != that1.User {
		return false
	}
	if len(this.Answers) != len(that1.Answers) {
		return false
	}
	for i := range this.Answers {
		if this.Answers[i] != that1.Answers[i] {
			return false
		}
	}
	return true
}
func (m *PollAnswer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PollAnswer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PollAnswer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintPolls(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintPolls(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PollData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PollData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PollData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllowsAnswerEdits {
		i--
		if m.AllowsAnswerEdits {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.AllowsMultipleAnswers {
		i--
		if m.AllowsMultipleAnswers {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndDate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPolls(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.ProvidedAnswers) > 0 {
		for iNdEx := len(m.ProvidedAnswers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProvidedAnswers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPolls(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Question) > 0 {
		i -= len(m.Question)
		copy(dAtA[i:], m.Question)
		i = encodeVarintPolls(dAtA, i, uint64(len(m.Question)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserAnswer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserAnswer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserAnswer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Answers) > 0 {
		for iNdEx := len(m.Answers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Answers[iNdEx])
			copy(dAtA[i:], m.Answers[iNdEx])
			i = encodeVarintPolls(dAtA, i, uint64(len(m.Answers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintPolls(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolls(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolls(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PollAnswer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovPolls(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovPolls(uint64(l))
	}
	return n
}

func (m *PollData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Question)
	if l > 0 {
		n += 1 + l + sovPolls(uint64(l))
	}
	if len(m.ProvidedAnswers) > 0 {
		for _, e := range m.ProvidedAnswers {
			l = e.Size()
			n += 1 + l + sovPolls(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndDate)
	n += 1 + l + sovPolls(uint64(l))
	if m.AllowsMultipleAnswers {
		n += 2
	}
	if m.AllowsAnswerEdits {
		n += 2
	}
	return n
}

func (m *UserAnswer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovPolls(uint64(l))
	}
	if len(m.Answers) > 0 {
		for _, s := range m.Answers {
			l = len(s)
			n += 1 + l + sovPolls(uint64(l))
		}
	}
	return n
}

func sovPolls(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolls(x uint64) (n int) {
	return sovPolls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PollAnswer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolls
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
			return fmt.Errorf("proto: PollAnswer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollAnswer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPolls
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
func (m *PollData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolls
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
			return fmt.Errorf("proto: PollData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Question", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Question = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvidedAnswers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProvidedAnswers = append(m.ProvidedAnswers, PollAnswer{})
			if err := m.ProvidedAnswers[len(m.ProvidedAnswers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowsMultipleAnswers", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowsMultipleAnswers = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowsAnswerEdits", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowsAnswerEdits = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPolls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPolls
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
func (m *UserAnswer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolls
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
			return fmt.Errorf("proto: UserAnswer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserAnswer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Answers = append(m.Answers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPolls
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
func skipPolls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolls
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
					return 0, ErrIntOverflowPolls
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
					return 0, ErrIntOverflowPolls
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
				return 0, ErrInvalidLengthPolls
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolls
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolls
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolls        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolls          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolls = fmt.Errorf("proto: unexpected end of group")
)
