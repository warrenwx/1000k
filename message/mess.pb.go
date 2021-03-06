// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mess.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageHeader struct {
	MType                uint64   `protobuf:"varint,1,opt,name=MType,proto3" json:"MType,omitempty"`
	MLength              uint64   `protobuf:"varint,2,opt,name=MLength,proto3" json:"MLength,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageHeader) Reset()         { *m = MessageHeader{} }
func (m *MessageHeader) String() string { return proto.CompactTextString(m) }
func (*MessageHeader) ProtoMessage()    {}
func (*MessageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78b50d87ecacffe, []int{0}
}

func (m *MessageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageHeader.Unmarshal(m, b)
}
func (m *MessageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageHeader.Marshal(b, m, deterministic)
}
func (m *MessageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageHeader.Merge(m, src)
}
func (m *MessageHeader) XXX_Size() int {
	return xxx_messageInfo_MessageHeader.Size(m)
}
func (m *MessageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MessageHeader proto.InternalMessageInfo

func (m *MessageHeader) GetMType() uint64 {
	if m != nil {
		return m.MType
	}
	return 0
}

func (m *MessageHeader) GetMLength() uint64 {
	if m != nil {
		return m.MLength
	}
	return 0
}

type SimpleMessageC struct {
	A                    uint64   `protobuf:"varint,31,opt,name=A,proto3" json:"A,omitempty"`
	B                    string   `protobuf:"bytes,32,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleMessageC) Reset()         { *m = SimpleMessageC{} }
func (m *SimpleMessageC) String() string { return proto.CompactTextString(m) }
func (*SimpleMessageC) ProtoMessage()    {}
func (*SimpleMessageC) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78b50d87ecacffe, []int{1}
}

func (m *SimpleMessageC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMessageC.Unmarshal(m, b)
}
func (m *SimpleMessageC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMessageC.Marshal(b, m, deterministic)
}
func (m *SimpleMessageC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMessageC.Merge(m, src)
}
func (m *SimpleMessageC) XXX_Size() int {
	return xxx_messageInfo_SimpleMessageC.Size(m)
}
func (m *SimpleMessageC) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMessageC.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMessageC proto.InternalMessageInfo

func (m *SimpleMessageC) GetA() uint64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SimpleMessageC) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

type HeavyMessageC struct {
	A                    uint64   `protobuf:"varint,41,opt,name=A,proto3" json:"A,omitempty"`
	B                    uint64   `protobuf:"varint,42,opt,name=B,proto3" json:"B,omitempty"`
	C                    uint64   `protobuf:"varint,43,opt,name=C,proto3" json:"C,omitempty"`
	D                    string   `protobuf:"bytes,44,opt,name=D,proto3" json:"D,omitempty"`
	E                    string   `protobuf:"bytes,45,opt,name=E,proto3" json:"E,omitempty"`
	F                    string   `protobuf:"bytes,46,opt,name=F,proto3" json:"F,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeavyMessageC) Reset()         { *m = HeavyMessageC{} }
func (m *HeavyMessageC) String() string { return proto.CompactTextString(m) }
func (*HeavyMessageC) ProtoMessage()    {}
func (*HeavyMessageC) Descriptor() ([]byte, []int) {
	return fileDescriptor_c78b50d87ecacffe, []int{2}
}

func (m *HeavyMessageC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeavyMessageC.Unmarshal(m, b)
}
func (m *HeavyMessageC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeavyMessageC.Marshal(b, m, deterministic)
}
func (m *HeavyMessageC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeavyMessageC.Merge(m, src)
}
func (m *HeavyMessageC) XXX_Size() int {
	return xxx_messageInfo_HeavyMessageC.Size(m)
}
func (m *HeavyMessageC) XXX_DiscardUnknown() {
	xxx_messageInfo_HeavyMessageC.DiscardUnknown(m)
}

var xxx_messageInfo_HeavyMessageC proto.InternalMessageInfo

func (m *HeavyMessageC) GetA() uint64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *HeavyMessageC) GetB() uint64 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *HeavyMessageC) GetC() uint64 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *HeavyMessageC) GetD() string {
	if m != nil {
		return m.D
	}
	return ""
}

func (m *HeavyMessageC) GetE() string {
	if m != nil {
		return m.E
	}
	return ""
}

func (m *HeavyMessageC) GetF() string {
	if m != nil {
		return m.F
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageHeader)(nil), "message.MessageHeader")
	proto.RegisterType((*SimpleMessageC)(nil), "message.SimpleMessageC")
	proto.RegisterType((*HeavyMessageC)(nil), "message.HeavyMessageC")
}

func init() { proto.RegisterFile("mess.proto", fileDescriptor_c78b50d87ecacffe) }

var fileDescriptor_c78b50d87ecacffe = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2d, 0x2e,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0xb1, 0x13, 0xd3, 0x53, 0x95, 0xec, 0xb9,
	0x78, 0x7d, 0x21, 0x4c, 0x8f, 0xd4, 0xc4, 0x94, 0xd4, 0x22, 0x21, 0x11, 0x2e, 0x56, 0xdf, 0x90,
	0xca, 0x82, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0xdd,
	0xd7, 0x27, 0x35, 0x2f, 0xbd, 0x24, 0x43, 0x82, 0x09, 0x2c, 0x0e, 0xe3, 0x2a, 0xe9, 0x70, 0xf1,
	0x05, 0x67, 0xe6, 0x16, 0xe4, 0xa4, 0x42, 0x8d, 0x71, 0x16, 0xe2, 0xe1, 0x62, 0x74, 0x94, 0x90,
	0x07, 0xab, 0x62, 0x74, 0x04, 0xf1, 0x9c, 0x24, 0x14, 0x14, 0x18, 0x35, 0x38, 0x83, 0x18, 0x9d,
	0x94, 0x92, 0xb9, 0x78, 0x3d, 0x52, 0x13, 0xcb, 0x2a, 0x51, 0x15, 0x6b, 0xa2, 0x28, 0xd6, 0x82,
	0xf0, 0x9c, 0x40, 0x3c, 0x67, 0x09, 0x6d, 0x08, 0x0f, 0xac, 0xd2, 0x45, 0x42, 0x07, 0x62, 0x90,
	0x0b, 0x88, 0xe7, 0x2a, 0xa1, 0x0b, 0xe1, 0xb9, 0x82, 0x78, 0x6e, 0x12, 0x7a, 0x10, 0x9e, 0x5b,
	0x12, 0x1b, 0xd8, 0x8f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0xe7, 0xad, 0x03, 0xf1,
	0x00, 0x00, 0x00,
}
