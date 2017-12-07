// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/examples/form/api/proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/form/api/proto/api.proto

It has these top-level messages:
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_api "github.com/micro/go-api/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

// Client API for Form service

type FormClient interface {
	// regular form
	Submit(ctx context.Context, in *go_micro_api.Request, opts ...client.CallOption) (*go_micro_api.Response, error)
	// multipart form
	Multipart(ctx context.Context, in *go_micro_api.Request, opts ...client.CallOption) (*go_micro_api.Response, error)
}

type formClient struct {
	c           client.Client
	serviceName string
}

func NewFormClient(serviceName string, c client.Client) FormClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "form"
	}
	return &formClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *formClient) Submit(ctx context.Context, in *go_micro_api.Request, opts ...client.CallOption) (*go_micro_api.Response, error) {
	req := c.c.NewRequest(c.serviceName, "Form.Submit", in)
	out := new(go_micro_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formClient) Multipart(ctx context.Context, in *go_micro_api.Request, opts ...client.CallOption) (*go_micro_api.Response, error) {
	req := c.c.NewRequest(c.serviceName, "Form.Multipart", in)
	out := new(go_micro_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Form service

type FormHandler interface {
	// regular form
	Submit(context.Context, *go_micro_api.Request, *go_micro_api.Response) error
	// multipart form
	Multipart(context.Context, *go_micro_api.Request, *go_micro_api.Response) error
}

func RegisterFormHandler(s server.Server, hdlr FormHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Form{hdlr}, opts...))
}

type Form struct {
	FormHandler
}

func (h *Form) Submit(ctx context.Context, in *go_micro_api.Request, out *go_micro_api.Response) error {
	return h.FormHandler.Submit(ctx, in, out)
}

func (h *Form) Multipart(ctx context.Context, in *go_micro_api.Request, out *go_micro_api.Response) error {
	return h.FormHandler.Multipart(ctx, in, out)
}

func init() { proto.RegisterFile("github.com/micro/examples/form/api/proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0xad, 0x48,
	0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0xcb, 0x2f, 0xca, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0x07, 0xb1, 0xf4, 0xc0, 0x2c, 0x29, 0x75, 0x0c, 0x3d, 0xe9, 0xf9, 0xba,
	0x58, 0x14, 0x1a, 0xd5, 0x73, 0xb1, 0xb8, 0xe5, 0x17, 0xe5, 0x0a, 0x59, 0x72, 0xb1, 0x05, 0x97,
	0x26, 0xe5, 0x66, 0x96, 0x08, 0x89, 0xea, 0xa5, 0xe7, 0xeb, 0x81, 0xf5, 0xe8, 0x81, 0x94, 0x05,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0x89, 0xa1, 0x0b, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x2a, 0x31, 0x08, 0xd9, 0x70, 0x71, 0xfa, 0x96, 0xe6, 0x94, 0x64, 0x16, 0x24, 0x16, 0x91, 0xae,
	0x3b, 0x89, 0x0d, 0xec, 0x0e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xf9, 0xce, 0x80,
	0xe6, 0x00, 0x00, 0x00,
}
