// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package evilnet

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

type EvilNetMsg_TYPE int32

const (
	EvilNetMsg_REQREG  EvilNetMsg_TYPE = 0
	EvilNetMsg_RSPREG  EvilNetMsg_TYPE = 1
	EvilNetMsg_ROUTER  EvilNetMsg_TYPE = 2
	EvilNetMsg_REQCONN EvilNetMsg_TYPE = 3
	EvilNetMsg_RSPCONN EvilNetMsg_TYPE = 4
)

var EvilNetMsg_TYPE_name = map[int32]string{
	0: "REQREG",
	1: "RSPREG",
	2: "ROUTER",
	3: "REQCONN",
	4: "RSPCONN",
}

var EvilNetMsg_TYPE_value = map[string]int32{
	"REQREG":  0,
	"RSPREG":  1,
	"ROUTER":  2,
	"REQCONN": 3,
	"RSPCONN": 4,
}

func (x EvilNetMsg_TYPE) String() string {
	return proto.EnumName(EvilNetMsg_TYPE_name, int32(x))
}

func (EvilNetMsg_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5, 0}
}

type EvilNetReqRegMsg struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Localaddr            string   `protobuf:"bytes,3,opt,name=localaddr,proto3" json:"localaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvilNetReqRegMsg) Reset()         { *m = EvilNetReqRegMsg{} }
func (m *EvilNetReqRegMsg) String() string { return proto.CompactTextString(m) }
func (*EvilNetReqRegMsg) ProtoMessage()    {}
func (*EvilNetReqRegMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *EvilNetReqRegMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvilNetReqRegMsg.Unmarshal(m, b)
}
func (m *EvilNetReqRegMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvilNetReqRegMsg.Marshal(b, m, deterministic)
}
func (m *EvilNetReqRegMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvilNetReqRegMsg.Merge(m, src)
}
func (m *EvilNetReqRegMsg) XXX_Size() int {
	return xxx_messageInfo_EvilNetReqRegMsg.Size(m)
}
func (m *EvilNetReqRegMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EvilNetReqRegMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EvilNetReqRegMsg proto.InternalMessageInfo

func (m *EvilNetReqRegMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EvilNetReqRegMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EvilNetReqRegMsg) GetLocaladdr() string {
	if m != nil {
		return m.Localaddr
	}
	return ""
}

type EvilNetRspRegMsg struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Fathername           string   `protobuf:"bytes,2,opt,name=fathername,proto3" json:"fathername,omitempty"`
	Newname              string   `protobuf:"bytes,3,opt,name=newname,proto3" json:"newname,omitempty"`
	Localaddr            string   `protobuf:"bytes,4,opt,name=localaddr,proto3" json:"localaddr,omitempty"`
	Globaladdr           string   `protobuf:"bytes,5,opt,name=globaladdr,proto3" json:"globaladdr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvilNetRspRegMsg) Reset()         { *m = EvilNetRspRegMsg{} }
func (m *EvilNetRspRegMsg) String() string { return proto.CompactTextString(m) }
func (*EvilNetRspRegMsg) ProtoMessage()    {}
func (*EvilNetRspRegMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *EvilNetRspRegMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvilNetRspRegMsg.Unmarshal(m, b)
}
func (m *EvilNetRspRegMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvilNetRspRegMsg.Marshal(b, m, deterministic)
}
func (m *EvilNetRspRegMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvilNetRspRegMsg.Merge(m, src)
}
func (m *EvilNetRspRegMsg) XXX_Size() int {
	return xxx_messageInfo_EvilNetRspRegMsg.Size(m)
}
func (m *EvilNetRspRegMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EvilNetRspRegMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EvilNetRspRegMsg proto.InternalMessageInfo

func (m *EvilNetRspRegMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EvilNetRspRegMsg) GetFathername() string {
	if m != nil {
		return m.Fathername
	}
	return ""
}

func (m *EvilNetRspRegMsg) GetNewname() string {
	if m != nil {
		return m.Newname
	}
	return ""
}

func (m *EvilNetRspRegMsg) GetLocaladdr() string {
	if m != nil {
		return m.Localaddr
	}
	return ""
}

func (m *EvilNetRspRegMsg) GetGlobaladdr() string {
	if m != nil {
		return m.Globaladdr
	}
	return ""
}

type EvilNetRouterMsg struct {
	Src                  string   `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst                  string   `protobuf:"bytes,2,opt,name=dst,proto3" json:"dst,omitempty"`
	Newname              []byte   `protobuf:"bytes,3,opt,name=newname,proto3" json:"newname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvilNetRouterMsg) Reset()         { *m = EvilNetRouterMsg{} }
func (m *EvilNetRouterMsg) String() string { return proto.CompactTextString(m) }
func (*EvilNetRouterMsg) ProtoMessage()    {}
func (*EvilNetRouterMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *EvilNetRouterMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvilNetRouterMsg.Unmarshal(m, b)
}
func (m *EvilNetRouterMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvilNetRouterMsg.Marshal(b, m, deterministic)
}
func (m *EvilNetRouterMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvilNetRouterMsg.Merge(m, src)
}
func (m *EvilNetRouterMsg) XXX_Size() int {
	return xxx_messageInfo_EvilNetRouterMsg.Size(m)
}
func (m *EvilNetRouterMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EvilNetRouterMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EvilNetRouterMsg proto.InternalMessageInfo

func (m *EvilNetRouterMsg) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *EvilNetRouterMsg) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

func (m *EvilNetRouterMsg) GetNewname() []byte {
	if m != nil {
		return m.Newname
	}
	return nil
}

type EvilNetReqConnMsg struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Proto                string   `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvilNetReqConnMsg) Reset()         { *m = EvilNetReqConnMsg{} }
func (m *EvilNetReqConnMsg) String() string { return proto.CompactTextString(m) }
func (*EvilNetReqConnMsg) ProtoMessage()    {}
func (*EvilNetReqConnMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *EvilNetReqConnMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvilNetReqConnMsg.Unmarshal(m, b)
}
func (m *EvilNetReqConnMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvilNetReqConnMsg.Marshal(b, m, deterministic)
}
func (m *EvilNetReqConnMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvilNetReqConnMsg.Merge(m, src)
}
func (m *EvilNetReqConnMsg) XXX_Size() int {
	return xxx_messageInfo_EvilNetReqConnMsg.Size(m)
}
func (m *EvilNetReqConnMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EvilNetReqConnMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EvilNetReqConnMsg proto.InternalMessageInfo

func (m *EvilNetReqConnMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EvilNetReqConnMsg) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

type EvilNetRspConnMsg struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Proto                string   `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	Result               string   `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvilNetRspConnMsg) Reset()         { *m = EvilNetRspConnMsg{} }
func (m *EvilNetRspConnMsg) String() string { return proto.CompactTextString(m) }
func (*EvilNetRspConnMsg) ProtoMessage()    {}
func (*EvilNetRspConnMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *EvilNetRspConnMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvilNetRspConnMsg.Unmarshal(m, b)
}
func (m *EvilNetRspConnMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvilNetRspConnMsg.Marshal(b, m, deterministic)
}
func (m *EvilNetRspConnMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvilNetRspConnMsg.Merge(m, src)
}
func (m *EvilNetRspConnMsg) XXX_Size() int {
	return xxx_messageInfo_EvilNetRspConnMsg.Size(m)
}
func (m *EvilNetRspConnMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EvilNetRspConnMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EvilNetRspConnMsg proto.InternalMessageInfo

func (m *EvilNetRspConnMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EvilNetRspConnMsg) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *EvilNetRspConnMsg) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type EvilNetMsg struct {
	Type                 int32              `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	ReqRegMsg            *EvilNetReqRegMsg  `protobuf:"bytes,2,opt,name=reqRegMsg,proto3" json:"reqRegMsg,omitempty"`
	RspRegMsg            *EvilNetRspRegMsg  `protobuf:"bytes,3,opt,name=rspRegMsg,proto3" json:"rspRegMsg,omitempty"`
	RouterMsg            *EvilNetRouterMsg  `protobuf:"bytes,4,opt,name=routerMsg,proto3" json:"routerMsg,omitempty"`
	ReqConnMsg           *EvilNetReqConnMsg `protobuf:"bytes,5,opt,name=reqConnMsg,proto3" json:"reqConnMsg,omitempty"`
	RspConnMsg           *EvilNetRspConnMsg `protobuf:"bytes,6,opt,name=rspConnMsg,proto3" json:"rspConnMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EvilNetMsg) Reset()         { *m = EvilNetMsg{} }
func (m *EvilNetMsg) String() string { return proto.CompactTextString(m) }
func (*EvilNetMsg) ProtoMessage()    {}
func (*EvilNetMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}

func (m *EvilNetMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvilNetMsg.Unmarshal(m, b)
}
func (m *EvilNetMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvilNetMsg.Marshal(b, m, deterministic)
}
func (m *EvilNetMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvilNetMsg.Merge(m, src)
}
func (m *EvilNetMsg) XXX_Size() int {
	return xxx_messageInfo_EvilNetMsg.Size(m)
}
func (m *EvilNetMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EvilNetMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EvilNetMsg proto.InternalMessageInfo

func (m *EvilNetMsg) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *EvilNetMsg) GetReqRegMsg() *EvilNetReqRegMsg {
	if m != nil {
		return m.ReqRegMsg
	}
	return nil
}

func (m *EvilNetMsg) GetRspRegMsg() *EvilNetRspRegMsg {
	if m != nil {
		return m.RspRegMsg
	}
	return nil
}

func (m *EvilNetMsg) GetRouterMsg() *EvilNetRouterMsg {
	if m != nil {
		return m.RouterMsg
	}
	return nil
}

func (m *EvilNetMsg) GetReqConnMsg() *EvilNetReqConnMsg {
	if m != nil {
		return m.ReqConnMsg
	}
	return nil
}

func (m *EvilNetMsg) GetRspConnMsg() *EvilNetRspConnMsg {
	if m != nil {
		return m.RspConnMsg
	}
	return nil
}

func init() {
	proto.RegisterEnum("EvilNetMsg_TYPE", EvilNetMsg_TYPE_name, EvilNetMsg_TYPE_value)
	proto.RegisterType((*EvilNetReqRegMsg)(nil), "EvilNetReqRegMsg")
	proto.RegisterType((*EvilNetRspRegMsg)(nil), "EvilNetRspRegMsg")
	proto.RegisterType((*EvilNetRouterMsg)(nil), "EvilNetRouterMsg")
	proto.RegisterType((*EvilNetReqConnMsg)(nil), "EvilNetReqConnMsg")
	proto.RegisterType((*EvilNetRspConnMsg)(nil), "EvilNetRspConnMsg")
	proto.RegisterType((*EvilNetMsg)(nil), "EvilNetMsg")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x7d, 0x85, 0x52, 0xd2, 0xcb, 0x5b, 0x94, 0xc9, 0xcb, 0x4b, 0x17, 0xc6, 0x98, 0xae, 0x5c,
	0x61, 0x82, 0x4b, 0x77, 0x62, 0xe3, 0x4a, 0x28, 0x03, 0x9a, 0xe8, 0xae, 0xc0, 0x58, 0x89, 0xa5,
	0xad, 0x33, 0x03, 0x86, 0xff, 0xe2, 0x9f, 0xf3, 0x9f, 0x98, 0x5e, 0x66, 0x3a, 0xb5, 0xea, 0xc2,
	0xdd, 0xb9, 0x1f, 0xe7, 0xcc, 0xe9, 0xb9, 0x05, 0x77, 0x23, 0x92, 0x41, 0xc1, 0x73, 0x99, 0x07,
	0x77, 0xe0, 0x85, 0xbb, 0x75, 0x3a, 0x66, 0x92, 0xb2, 0x17, 0xca, 0x92, 0x1b, 0x91, 0x10, 0x0f,
	0xda, 0xcf, 0x6c, 0xef, 0x5b, 0x27, 0xd6, 0xa9, 0x4b, 0x4b, 0x48, 0x08, 0xd8, 0x59, 0xbc, 0x61,
	0x7e, 0x0b, 0x5b, 0x88, 0xc9, 0x11, 0xb8, 0x69, 0xbe, 0x8c, 0xd3, 0x78, 0xb5, 0xe2, 0x7e, 0x1b,
	0x07, 0xa6, 0x11, 0xbc, 0x59, 0x46, 0x58, 0x14, 0x3f, 0x0a, 0x1f, 0x03, 0x3c, 0xc6, 0xf2, 0x89,
	0xf1, 0x9a, 0x7c, 0xad, 0x43, 0x7c, 0xe8, 0x66, 0xec, 0x15, 0x87, 0x87, 0x27, 0x74, 0xf9, 0xf9,
	0x79, 0xbb, 0xf1, 0x7c, 0xa9, 0x9b, 0xa4, 0xf9, 0x42, 0x8d, 0x3b, 0x07, 0x5d, 0xd3, 0x09, 0x22,
	0xe3, 0x2e, 0xdf, 0x4a, 0xc6, 0x95, 0x3b, 0xc1, 0x97, 0xda, 0x9d, 0xe0, 0xcb, 0xb2, 0xb3, 0x12,
	0x52, 0xd9, 0x2a, 0x61, 0xd3, 0xcf, 0xdf, 0xca, 0x4f, 0x70, 0x01, 0x7d, 0x13, 0xe4, 0x28, 0xcf,
	0xb2, 0xef, 0x3f, 0xf8, 0x1f, 0x74, 0x30, 0x78, 0x25, 0x7a, 0x28, 0x82, 0x99, 0x21, 0x8b, 0xe2,
	0x97, 0x64, 0xf2, 0x1f, 0x1c, 0xce, 0xc4, 0x36, 0x95, 0x2a, 0x22, 0x55, 0x05, 0xef, 0x2d, 0x00,
	0xa5, 0x5a, 0xca, 0x11, 0xb0, 0xe5, 0xbe, 0x60, 0xa8, 0xd7, 0xa1, 0x88, 0xc9, 0x19, 0xb8, 0x5c,
	0x9f, 0x1d, 0x45, 0x7b, 0xc3, 0xfe, 0xa0, 0xf9, 0x3f, 0x50, 0xb3, 0x83, 0x04, 0x7d, 0x4e, 0x7c,
	0xae, 0x4e, 0xd0, 0x03, 0x6a, 0x76, 0x90, 0xa0, 0x13, 0xc6, 0x33, 0xd5, 0x09, 0x7a, 0x40, 0xcd,
	0x0e, 0x19, 0x02, 0xf0, 0x2a, 0x40, 0xbc, 0x5c, 0x6f, 0x48, 0x06, 0x5f, 0xa2, 0xa5, 0xb5, 0x2d,
	0xe4, 0x54, 0xb9, 0xf9, 0x4e, 0x83, 0x53, 0x4d, 0x68, 0x6d, 0x2b, 0xb8, 0x02, 0x7b, 0x7e, 0x1f,
	0x85, 0x04, 0xc0, 0xa1, 0xe1, 0x94, 0x86, 0xd7, 0xde, 0x1f, 0xc4, 0xb3, 0xa8, 0xc4, 0x16, 0xe2,
	0xc9, 0xed, 0x3c, 0xa4, 0x5e, 0x8b, 0xf4, 0xa0, 0x4b, 0xc3, 0xe9, 0x68, 0x32, 0x1e, 0x7b, 0x6d,
	0x2c, 0x66, 0x11, 0x16, 0xf6, 0xa5, 0xfb, 0xd0, 0x65, 0xbb, 0x75, 0x9a, 0x31, 0xb9, 0x70, 0xf0,
	0x1a, 0xe7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x7b, 0x66, 0x0e, 0x5d, 0x03, 0x00, 0x00,
}
