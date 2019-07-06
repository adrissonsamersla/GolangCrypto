// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto.proto

package crypto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_crypto_4408140520a15d80, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_crypto_4408140520a15d80, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "crypto.Request")
	proto.RegisterType((*Response)(nil), "crypto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EncryptClient is the client API for Encrypt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EncryptClient interface {
	Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type encryptClient struct {
	cc *grpc.ClientConn
}

func NewEncryptClient(cc *grpc.ClientConn) EncryptClient {
	return &encryptClient{cc}
}

func (c *encryptClient) Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/crypto.Encrypt/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncryptServer is the server API for Encrypt service.
type EncryptServer interface {
	Do(context.Context, *Request) (*Response, error)
}

func RegisterEncryptServer(s *grpc.Server, srv EncryptServer) {
	s.RegisterService(&_Encrypt_serviceDesc, srv)
}

func _Encrypt_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.Encrypt/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptServer).Do(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Encrypt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crypto.Encrypt",
	HandlerType: (*EncryptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _Encrypt_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto.proto",
}

func init() { proto.RegisterFile("crypto.proto", fileDescriptor_crypto_4408140520a15d80) }

var fileDescriptor_crypto_4408140520a15d80 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2e, 0xaa, 0x2c,
	0x28, 0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x94, 0xb9, 0xd8,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13,
	0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x15, 0x2e, 0x8e, 0xa0,
	0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xdc, 0xaa, 0x8c, 0x4c, 0xb8, 0xd8, 0x5d, 0xf3, 0xc0,
	0xc6, 0x0a, 0x69, 0x72, 0x31, 0xb9, 0xe4, 0x0b, 0xf1, 0xeb, 0x41, 0xad, 0x84, 0xda, 0x20, 0x25,
	0x80, 0x10, 0x80, 0x98, 0xa6, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0x8f, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0xec, 0x69, 0x74, 0x9f, 0x00, 0x00, 0x00,
}
