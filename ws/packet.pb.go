// Code generated by protoc-gen-go. DO NOT EDIT.
// source: packet.proto

package ws

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Call struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}
func (*Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{1}
}

func (m *Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call.Unmarshal(m, b)
}
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call.Marshal(b, m, deterministic)
}
func (m *Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call.Merge(m, src)
}
func (m *Call) XXX_Size() int {
	return xxx_messageInfo_Call.Size(m)
}
func (m *Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Call proto.InternalMessageInfo

func (m *Call) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Call) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Call) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Reply struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*Reply_Data
	//	*Reply_Error
	Result               isReply_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{2}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isReply_Result interface {
	isReply_Result()
}

type Reply_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type Reply_Error struct {
	Error *Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*Reply_Data) isReply_Result() {}

func (*Reply_Error) isReply_Result() {}

func (m *Reply) GetResult() isReply_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Reply) GetData() []byte {
	if x, ok := m.GetResult().(*Reply_Data); ok {
		return x.Data
	}
	return nil
}

func (m *Reply) GetError() *Error {
	if x, ok := m.GetResult().(*Reply_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Reply) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Reply_Data)(nil),
		(*Reply_Error)(nil),
	}
}

type Header struct {
	Entries              map[string]string `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{3}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetEntries() map[string]string {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Packet struct {
	// Types that are valid to be assigned to V:
	//	*Packet_Push
	//	*Packet_Call
	//	*Packet_Reply
	//	*Packet_Header
	V                    isPacket_V `protobuf_oneof:"v"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{4}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

type isPacket_V interface {
	isPacket_V()
}

type Packet_Push struct {
	Push []byte `protobuf:"bytes,1,opt,name=push,proto3,oneof"`
}

type Packet_Call struct {
	Call *Call `protobuf:"bytes,2,opt,name=call,proto3,oneof"`
}

type Packet_Reply struct {
	Reply *Reply `protobuf:"bytes,3,opt,name=reply,proto3,oneof"`
}

type Packet_Header struct {
	Header *Header `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}

func (*Packet_Push) isPacket_V() {}

func (*Packet_Call) isPacket_V() {}

func (*Packet_Reply) isPacket_V() {}

func (*Packet_Header) isPacket_V() {}

func (m *Packet) GetV() isPacket_V {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *Packet) GetPush() []byte {
	if x, ok := m.GetV().(*Packet_Push); ok {
		return x.Push
	}
	return nil
}

func (m *Packet) GetCall() *Call {
	if x, ok := m.GetV().(*Packet_Call); ok {
		return x.Call
	}
	return nil
}

func (m *Packet) GetReply() *Reply {
	if x, ok := m.GetV().(*Packet_Reply); ok {
		return x.Reply
	}
	return nil
}

func (m *Packet) GetHeader() *Header {
	if x, ok := m.GetV().(*Packet_Header); ok {
		return x.Header
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Packet) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Packet_Push)(nil),
		(*Packet_Call)(nil),
		(*Packet_Reply)(nil),
		(*Packet_Header)(nil),
	}
}

func init() {
	proto.RegisterType((*Error)(nil), "ws.Error")
	proto.RegisterType((*Call)(nil), "ws.Call")
	proto.RegisterType((*Reply)(nil), "ws.Reply")
	proto.RegisterType((*Header)(nil), "ws.Header")
	proto.RegisterMapType((map[string]string)(nil), "ws.Header.EntriesEntry")
	proto.RegisterType((*Packet)(nil), "ws.Packet")
}

func init() { proto.RegisterFile("packet.proto", fileDescriptor_e9ef1a6541f9f9e7) }

var fileDescriptor_e9ef1a6541f9f9e7 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x8b, 0xe2, 0x40,
	0x10, 0xb5, 0xf3, 0xa5, 0x96, 0x61, 0x59, 0x1a, 0x61, 0x83, 0x87, 0x25, 0x1b, 0xf6, 0x90, 0x53,
	0x64, 0x1c, 0x06, 0x06, 0x0f, 0x73, 0x70, 0x10, 0x72, 0x1c, 0xfa, 0xe8, 0xad, 0x4d, 0x1a, 0x0d,
	0x46, 0x13, 0xba, 0x13, 0x83, 0xff, 0x61, 0x7e, 0xf4, 0x50, 0x15, 0xc3, 0x08, 0x73, 0x4a, 0x55,
	0xbf, 0x57, 0x8f, 0xf7, 0x1e, 0x01, 0xbf, 0x96, 0xd9, 0x49, 0x35, 0x49, 0xad, 0xab, 0xa6, 0xe2,
	0x56, 0x67, 0xa2, 0x17, 0x70, 0xb7, 0x5a, 0x57, 0x9a, 0x73, 0x70, 0xb2, 0x2a, 0x57, 0x01, 0x0b,
	0x59, 0xec, 0x0a, 0x9a, 0x79, 0x00, 0xe3, 0xb3, 0x32, 0x46, 0x1e, 0x54, 0x60, 0x85, 0x2c, 0x9e,
	0x8a, 0x61, 0x8d, 0xde, 0xc0, 0x79, 0x97, 0x65, 0xc9, 0x7f, 0x81, 0x55, 0xe4, 0xf7, 0x1b, 0xab,
	0xc8, 0x51, 0xe5, 0x22, 0xcf, 0x03, 0x9d, 0x66, 0x7c, 0xcb, 0x65, 0x23, 0x03, 0x3b, 0x64, 0xb1,
	0x2f, 0x68, 0x8e, 0x76, 0xe0, 0x0a, 0x55, 0x97, 0xb7, 0x1f, 0x02, 0xf3, 0x3b, 0x19, 0x05, 0xfc,
	0x74, 0xd4, 0xd3, 0xf9, 0x3f, 0x70, 0x15, 0xba, 0x24, 0x8d, 0xd9, 0x6a, 0x9a, 0x74, 0x26, 0x21,
	0xdb, 0xe9, 0x48, 0xf4, 0xc8, 0x66, 0x02, 0x9e, 0x56, 0xa6, 0x2d, 0x9b, 0xa8, 0x03, 0x2f, 0x55,
	0x32, 0x57, 0x9a, 0x3f, 0xc1, 0x58, 0x5d, 0x1a, 0x5d, 0x28, 0x13, 0xb0, 0xd0, 0x8e, 0x67, 0xab,
	0x3f, 0x78, 0xd8, 0x83, 0xc9, 0xb6, 0x47, 0xf0, 0x73, 0x13, 0x03, 0x6f, 0xb1, 0x06, 0xff, 0x11,
	0xe0, 0xbf, 0xc1, 0x3e, 0xa9, 0x1b, 0x19, 0x9c, 0x0a, 0x1c, 0xf9, 0x1c, 0xdc, 0xab, 0x2c, 0xdb,
	0x21, 0x63, 0xbf, 0xac, 0xad, 0x57, 0x16, 0x7d, 0x32, 0xf0, 0x3e, 0xa8, 0x60, 0x8c, 0x51, 0xb7,
	0xe6, 0x48, 0x77, 0x14, 0x03, 0x37, 0xfe, 0x17, 0x9c, 0x4c, 0x96, 0x25, 0x5d, 0xce, 0x56, 0x13,
	0x34, 0x83, 0x2d, 0x22, 0x8e, 0xef, 0x18, 0x53, 0x63, 0x2b, 0x8f, 0x31, 0xa9, 0x26, 0x8c, 0x49,
	0x08, 0xff, 0x0f, 0xde, 0x91, 0xfc, 0x07, 0x0e, 0x71, 0xe0, 0x3b, 0x51, 0x3a, 0x12, 0x77, 0x6c,
	0x63, 0x03, 0xbb, 0x6e, 0x16, 0xbb, 0xe0, 0x50, 0x34, 0xc7, 0x76, 0x9f, 0x64, 0xd5, 0x79, 0x79,
	0xa8, 0xea, 0x76, 0xbf, 0xec, 0x8a, 0x8b, 0x5a, 0x76, 0x66, 0xef, 0xd1, 0x1f, 0xf0, 0xfc, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x5e, 0x85, 0xd3, 0x63, 0x11, 0x02, 0x00, 0x00,
}
