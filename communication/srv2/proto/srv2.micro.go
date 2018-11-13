// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv2.proto

/*
Package go_micro_srv_srv2 is a generated protocol buffer package.

It is generated from these files:
	srv2.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_srv2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Srv2 service

type Srv2Service interface {
	Square(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type srv2Service struct {
	c    client.Client
	name string
}

func NewSrv2Service(name string, c client.Client) Srv2Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.srv2"
	}
	return &srv2Service{
		c:    c,
		name: name,
	}
}

func (c *srv2Service) Square(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Srv2.Square", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Srv2 service

type Srv2Handler interface {
	Square(context.Context, *Request, *Response) error
}

func RegisterSrv2Handler(s server.Server, hdlr Srv2Handler, opts ...server.HandlerOption) error {
	type srv2 interface {
		Square(ctx context.Context, in *Request, out *Response) error
	}
	type Srv2 struct {
		srv2
	}
	h := &srv2Handler{hdlr}
	return s.Handle(s.NewHandler(&Srv2{h}, opts...))
}

type srv2Handler struct {
	Srv2Handler
}

func (h *srv2Handler) Square(ctx context.Context, in *Request, out *Response) error {
	return h.Srv2Handler.Square(ctx, in, out)
}
