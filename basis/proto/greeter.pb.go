// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	greeter.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto1.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto1.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto1.RegisterType((*HelloRequest)(nil), "proto.HelloRequest")
	proto1.RegisterType((*HelloResponse)(nil), "proto.HelloResponse")
}

func init() { proto1.RegisterFile("greeter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x4a, 0x5c,
	0x3c, 0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c,
	0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x36,
	0x17, 0x2f, 0x54, 0x4d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17, 0x07, 0xd8, 0xb0,
	0xcc, 0xbc, 0x74, 0x09, 0x26, 0xb0, 0x42, 0x38, 0xdf, 0xc8, 0x96, 0x8b, 0xdd, 0x1d, 0x62, 0x91,
	0x90, 0x11, 0x17, 0x2b, 0x58, 0x9f, 0x90, 0x30, 0xc4, 0x4e, 0x3d, 0x64, 0x9b, 0xa4, 0x44, 0x50,
	0x05, 0x21, 0x46, 0x27, 0xb1, 0x81, 0x05, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xa5,
	0x91, 0xd4, 0xae, 0x00, 0x00, 0x00,
}