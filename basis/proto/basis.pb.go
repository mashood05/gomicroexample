// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basis.proto

/*
Package go_micro_srv_basis is a generated protocol buffer package.

It is generated from these files:
	basis.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_basis

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.basis.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.basis.Response")
}

func init() { proto.RegisterFile("basis.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4a, 0x2c, 0xce,
	0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c, 0x2e,
	0xca, 0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x03, 0xcb, 0x28, 0xc9, 0x72, 0xb1, 0x07, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0xd9, 0x4a, 0x32, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9,
	0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x50, 0x69, 0x10, 0xd3, 0xc8, 0x93, 0x8b, 0x39, 0x38,
	0xb1, 0x52, 0xc8, 0x89, 0x8b, 0xd5, 0x23, 0x35, 0x27, 0x27, 0x5f, 0x48, 0x5a, 0x0f, 0xd3, 0x06,
	0x3d, 0xa8, 0xf1, 0x52, 0x32, 0xd8, 0x25, 0x21, 0x86, 0x27, 0xb1, 0x81, 0x9d, 0x68, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0xf6, 0x87, 0x58, 0xb1, 0x00, 0x00, 0x00,
}