// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frame.proto

package frame

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

type FrameData_TYPE int32

const (
	FrameData_USER_DATA FrameData_TYPE = 0
	FrameData_CONN      FrameData_TYPE = 1
	FrameData_CONNRSP   FrameData_TYPE = 2
	FrameData_CLOSE     FrameData_TYPE = 3
)

var FrameData_TYPE_name = map[int32]string{
	0: "USER_DATA",
	1: "CONN",
	2: "CONNRSP",
	3: "CLOSE",
}

var FrameData_TYPE_value = map[string]int32{
	"USER_DATA": 0,
	"CONN":      1,
	"CONNRSP":   2,
	"CLOSE":     3,
}

func (x FrameData_TYPE) String() string {
	return proto.EnumName(FrameData_TYPE_name, int32(x))
}

func (FrameData_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{0, 0}
}

type Frame_TYPE int32

const (
	Frame_DATA Frame_TYPE = 0
	Frame_REQ  Frame_TYPE = 1
	Frame_ACK  Frame_TYPE = 2
	Frame_PING Frame_TYPE = 3
	Frame_PONG Frame_TYPE = 4
)

var Frame_TYPE_name = map[int32]string{
	0: "DATA",
	1: "REQ",
	2: "ACK",
	3: "PING",
	4: "PONG",
}

var Frame_TYPE_value = map[string]int32{
	"DATA": 0,
	"REQ":  1,
	"ACK":  2,
	"PING": 3,
	"PONG": 4,
}

func (x Frame_TYPE) String() string {
	return proto.EnumName(Frame_TYPE_name, int32(x))
}

func (Frame_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{1, 0}
}

type FrameData struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Compress             bool     `protobuf:"varint,3,opt,name=compress,proto3" json:"compress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrameData) Reset()         { *m = FrameData{} }
func (m *FrameData) String() string { return proto.CompactTextString(m) }
func (*FrameData) ProtoMessage()    {}
func (*FrameData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{0}
}

func (m *FrameData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameData.Unmarshal(m, b)
}
func (m *FrameData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameData.Marshal(b, m, deterministic)
}
func (m *FrameData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameData.Merge(m, src)
}
func (m *FrameData) XXX_Size() int {
	return xxx_messageInfo_FrameData.Size(m)
}
func (m *FrameData) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameData.DiscardUnknown(m)
}

var xxx_messageInfo_FrameData proto.InternalMessageInfo

func (m *FrameData) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *FrameData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FrameData) GetCompress() bool {
	if m != nil {
		return m.Compress
	}
	return false
}

type Frame struct {
	Type                 int32      `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Resend               bool       `protobuf:"varint,2,opt,name=resend,proto3" json:"resend,omitempty"`
	Sendtime             int64      `protobuf:"varint,3,opt,name=sendtime,proto3" json:"sendtime,omitempty"`
	Id                   int32      `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *FrameData `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Dataid               []int32    `protobuf:"varint,6,rep,packed,name=dataid,proto3" json:"dataid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{1}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Frame) GetResend() bool {
	if m != nil {
		return m.Resend
	}
	return false
}

func (m *Frame) GetSendtime() int64 {
	if m != nil {
		return m.Sendtime
	}
	return 0
}

func (m *Frame) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Frame) GetData() *FrameData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Frame) GetDataid() []int32 {
	if m != nil {
		return m.Dataid
	}
	return nil
}

func init() {
	proto.RegisterEnum("FrameData_TYPE", FrameData_TYPE_name, FrameData_TYPE_value)
	proto.RegisterEnum("Frame_TYPE", Frame_TYPE_name, Frame_TYPE_value)
	proto.RegisterType((*FrameData)(nil), "FrameData")
	proto.RegisterType((*Frame)(nil), "Frame")
}

func init() { proto.RegisterFile("frame.proto", fileDescriptor_5379e2b825e15002) }

var fileDescriptor_5379e2b825e15002 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0x85, 0x30,
	0x14, 0xc6, 0x2d, 0x6d, 0xf9, 0x73, 0x50, 0xd3, 0x74, 0xb8, 0x21, 0x0e, 0x86, 0x30, 0x31, 0xdd,
	0x41, 0x13, 0x9d, 0x91, 0x8b, 0x37, 0x46, 0x03, 0x58, 0xae, 0x83, 0x2e, 0x06, 0xa5, 0x26, 0x0c,
	0x08, 0x81, 0x2e, 0xbe, 0x81, 0x6f, 0xe7, 0x2b, 0x99, 0x36, 0x0d, 0xd3, 0x9d, 0xce, 0xf7, 0xb5,
	0xc9, 0xc9, 0xef, 0x77, 0x20, 0xfc, 0x9a, 0xdb, 0x41, 0x6e, 0xa7, 0x79, 0x54, 0x63, 0xf2, 0x8b,
	0x20, 0xb8, 0xd7, 0x7d, 0xd7, 0xaa, 0x96, 0x73, 0x20, 0xea, 0x67, 0x92, 0x11, 0x8a, 0x51, 0x4a,
	0x85, 0xc9, 0xfa, 0xad, 0x6b, 0x55, 0x1b, 0x39, 0x31, 0x4a, 0x4f, 0x85, 0xc9, 0xfc, 0x02, 0xfc,
	0xcf, 0x71, 0x98, 0x66, 0xb9, 0x2c, 0x11, 0x8e, 0x51, 0xea, 0x8b, 0xb5, 0x27, 0xb7, 0x40, 0x0e,
	0xaf, 0x75, 0xc1, 0xcf, 0x20, 0x78, 0x69, 0x0a, 0xf1, 0xbe, 0xcb, 0x0e, 0x19, 0x3b, 0xe1, 0x3e,
	0x90, 0xbc, 0x2a, 0x4b, 0x86, 0x78, 0x08, 0x9e, 0x4e, 0xa2, 0xa9, 0x99, 0xc3, 0x03, 0xa0, 0xf9,
	0x53, 0xd5, 0x14, 0x0c, 0x27, 0x7f, 0x08, 0xa8, 0x41, 0x39, 0x8a, 0xb1, 0x01, 0x77, 0x96, 0x8b,
	0xfc, 0xee, 0x0c, 0x88, 0x2f, 0x6c, 0xd3, 0x28, 0x7a, 0xaa, 0x7e, 0x90, 0x06, 0x05, 0x8b, 0xb5,
	0xf3, 0x73, 0x70, 0xfa, 0x2e, 0x22, 0x66, 0x8b, 0xd3, 0x77, 0xfc, 0xd2, 0xaa, 0xd0, 0x18, 0xa5,
	0xe1, 0x15, 0x6c, 0x57, 0x71, 0xab, 0xb5, 0x01, 0x57, 0xcf, 0xbe, 0x8b, 0xdc, 0x18, 0xa7, 0x54,
	0xd8, 0x96, 0xdc, 0x58, 0x25, 0x1f, 0x88, 0xb5, 0xf1, 0x00, 0x8b, 0xe2, 0x99, 0x21, 0x1d, 0xb2,
	0xfc, 0x91, 0x39, 0xfa, 0xaf, 0x7e, 0x28, 0xf7, 0x0c, 0x9b, 0x54, 0x95, 0x7b, 0x46, 0xee, 0xbc,
	0x37, 0x6a, 0x6e, 0xfd, 0xe1, 0x9a, 0x63, 0x5f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xc7,
	0xff, 0x0e, 0x7b, 0x01, 0x00, 0x00,
}
