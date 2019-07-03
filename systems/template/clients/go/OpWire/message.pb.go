// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package OpWire

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

type Operation struct {
	// Types that are valid to be assigned to OpType:
	//	*Operation_Put
	//	*Operation_Get
	//	*Operation_Quit
	OpType               isOperation_OpType `protobuf_oneof:"op_type"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

type isOperation_OpType interface {
	isOperation_OpType()
}

type Operation_Put struct {
	Put *OperationOpPut `protobuf:"bytes,1,opt,name=put,proto3,oneof"`
}

type Operation_Get struct {
	Get *OperationOpGet `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type Operation_Quit struct {
	Quit *OperationOpQuit `protobuf:"bytes,3,opt,name=quit,proto3,oneof"`
}

func (*Operation_Put) isOperation_OpType() {}

func (*Operation_Get) isOperation_OpType() {}

func (*Operation_Quit) isOperation_OpType() {}

func (m *Operation) GetOpType() isOperation_OpType {
	if m != nil {
		return m.OpType
	}
	return nil
}

func (m *Operation) GetPut() *OperationOpPut {
	if x, ok := m.GetOpType().(*Operation_Put); ok {
		return x.Put
	}
	return nil
}

func (m *Operation) GetGet() *OperationOpGet {
	if x, ok := m.GetOpType().(*Operation_Get); ok {
		return x.Get
	}
	return nil
}

func (m *Operation) GetQuit() *OperationOpQuit {
	if x, ok := m.GetOpType().(*Operation_Quit); ok {
		return x.Quit
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Operation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Operation_Put)(nil),
		(*Operation_Get)(nil),
		(*Operation_Quit)(nil),
	}
}

type OperationOpPut struct {
	Key                  uint64   `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationOpPut) Reset()         { *m = OperationOpPut{} }
func (m *OperationOpPut) String() string { return proto.CompactTextString(m) }
func (*OperationOpPut) ProtoMessage()    {}
func (*OperationOpPut) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

func (m *OperationOpPut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationOpPut.Unmarshal(m, b)
}
func (m *OperationOpPut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationOpPut.Marshal(b, m, deterministic)
}
func (m *OperationOpPut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationOpPut.Merge(m, src)
}
func (m *OperationOpPut) XXX_Size() int {
	return xxx_messageInfo_OperationOpPut.Size(m)
}
func (m *OperationOpPut) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationOpPut.DiscardUnknown(m)
}

var xxx_messageInfo_OperationOpPut proto.InternalMessageInfo

func (m *OperationOpPut) GetKey() uint64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *OperationOpPut) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type OperationOpGet struct {
	Key                  uint64   `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationOpGet) Reset()         { *m = OperationOpGet{} }
func (m *OperationOpGet) String() string { return proto.CompactTextString(m) }
func (*OperationOpGet) ProtoMessage()    {}
func (*OperationOpGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}

func (m *OperationOpGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationOpGet.Unmarshal(m, b)
}
func (m *OperationOpGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationOpGet.Marshal(b, m, deterministic)
}
func (m *OperationOpGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationOpGet.Merge(m, src)
}
func (m *OperationOpGet) XXX_Size() int {
	return xxx_messageInfo_OperationOpGet.Size(m)
}
func (m *OperationOpGet) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationOpGet.DiscardUnknown(m)
}

var xxx_messageInfo_OperationOpGet proto.InternalMessageInfo

func (m *OperationOpGet) GetKey() uint64 {
	if m != nil {
		return m.Key
	}
	return 0
}

type OperationOpQuit struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationOpQuit) Reset()         { *m = OperationOpQuit{} }
func (m *OperationOpQuit) String() string { return proto.CompactTextString(m) }
func (*OperationOpQuit) ProtoMessage()    {}
func (*OperationOpQuit) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 2}
}

func (m *OperationOpQuit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationOpQuit.Unmarshal(m, b)
}
func (m *OperationOpQuit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationOpQuit.Marshal(b, m, deterministic)
}
func (m *OperationOpQuit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationOpQuit.Merge(m, src)
}
func (m *OperationOpQuit) XXX_Size() int {
	return xxx_messageInfo_OperationOpQuit.Size(m)
}
func (m *OperationOpQuit) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationOpQuit.DiscardUnknown(m)
}

var xxx_messageInfo_OperationOpQuit proto.InternalMessageInfo

func (m *OperationOpQuit) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Response struct {
	ResponseTime         float64  `protobuf:"fixed64,1,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Start                float64  `protobuf:"fixed64,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  float64  `protobuf:"fixed64,4,opt,name=end,proto3" json:"end,omitempty"`
	Id                   uint32   `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponseTime() float64 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

func (m *Response) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *Response) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Response) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Response) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Operation)(nil), "OpWire.Operation")
	proto.RegisterType((*OperationOpPut)(nil), "OpWire.Operation.op_put")
	proto.RegisterType((*OperationOpGet)(nil), "OpWire.Operation.op_get")
	proto.RegisterType((*OperationOpQuit)(nil), "OpWire.Operation.op_quit")
	proto.RegisterType((*Response)(nil), "OpWire.Response")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x9b, 0x3f, 0xcd, 0xdb, 0xcc, 0xdb, 0x88, 0x2c, 0x1e, 0xd6, 0x78, 0x29, 0xf5, 0xd2,
	0x83, 0x44, 0xd1, 0x6f, 0xd0, 0x53, 0x6f, 0x85, 0x45, 0xf0, 0x58, 0x22, 0x19, 0xc2, 0xa2, 0x49,
	0xd6, 0xdd, 0x8d, 0x90, 0x2f, 0xe1, 0x67, 0x96, 0x99, 0x8d, 0x5e, 0x14, 0x6f, 0xcf, 0xcc, 0xf3,
	0x9b, 0x7d, 0x66, 0x19, 0x28, 0x3a, 0x74, 0xae, 0x6e, 0xb1, 0x32, 0x76, 0xf0, 0x83, 0xc8, 0x8e,
	0xe6, 0x49, 0x5b, 0xdc, 0x7e, 0xc4, 0x90, 0x1f, 0x0d, 0xda, 0xda, 0xeb, 0xa1, 0x17, 0x37, 0x90,
	0x98, 0xd1, 0xcb, 0x68, 0x13, 0xed, 0xfe, 0xdf, 0xcb, 0x2a, 0x30, 0xd5, 0xb7, 0x5f, 0x0d, 0xe6,
	0x64, 0x46, 0x7f, 0x58, 0x28, 0xc2, 0x88, 0x6e, 0xd1, 0xcb, 0xf8, 0x0f, 0xba, 0x45, 0xa6, 0x5b,
	0xf4, 0xe2, 0x16, 0xd2, 0xb7, 0x51, 0x7b, 0x99, 0x30, 0x7e, 0xf9, 0x2b, 0x4e, 0xc0, 0x61, 0xa1,
	0x18, 0x2c, 0xef, 0x20, 0x0b, 0x79, 0xe2, 0x1c, 0x92, 0x17, 0x9c, 0x78, 0xad, 0x54, 0x91, 0x14,
	0x17, 0xb0, 0x7c, 0xaf, 0x5f, 0x47, 0xe4, 0xf0, 0xb5, 0x0a, 0x45, 0x59, 0xf2, 0x04, 0x85, 0xfd,
	0x98, 0x28, 0xaf, 0xe0, 0xdf, 0x1c, 0x40, 0x66, 0xe7, 0x5a, 0x36, 0x73, 0x45, 0x72, 0x9f, 0xb3,
	0xe9, 0x27, 0x83, 0xdb, 0x09, 0x56, 0x0a, 0x9d, 0x19, 0x7a, 0x87, 0xe2, 0x1a, 0x0a, 0x3b, 0xeb,
	0x93, 0xd7, 0x1d, 0xf2, 0x48, 0xa4, 0xd6, 0x5f, 0xcd, 0x47, 0xdd, 0x21, 0xbd, 0x86, 0xd6, 0xf2,
	0x22, 0xb9, 0x22, 0x49, 0xcb, 0x39, 0x5f, 0xdb, 0xf0, 0xd5, 0x48, 0x85, 0x82, 0xb9, 0xbe, 0x91,
	0x29, 0xf7, 0x48, 0x8a, 0x33, 0x88, 0x75, 0x23, 0x97, 0x9b, 0x68, 0x57, 0xa8, 0x58, 0x37, 0xfb,
	0x15, 0xcc, 0x57, 0x79, 0xce, 0xf8, 0x48, 0x0f, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x96, 0x88,
	0x82, 0xd7, 0xb5, 0x01, 0x00, 0x00,
}