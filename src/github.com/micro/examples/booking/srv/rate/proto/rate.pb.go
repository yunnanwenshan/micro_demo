// Code generated by protoc-gen-go.
// source: srv/rate/proto/rate.proto
// DO NOT EDIT!

/*
Package rate is a generated protocol buffer package.

It is generated from these files:
	srv/rate/proto/rate.proto

It has these top-level messages:
	Request
	Result
	RatePlan
	RoomType
*/
package rate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Request struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	InDate   string   `protobuf:"bytes,2,opt,name=inDate" json:"inDate,omitempty"`
	OutDate  string   `protobuf:"bytes,3,opt,name=outDate" json:"outDate,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
	RatePlans []*RatePlan `protobuf:"bytes,1,rep,name=ratePlans" json:"ratePlans,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetRatePlans() []*RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

type RatePlan struct {
	HotelId  string    `protobuf:"bytes,1,opt,name=hotelId" json:"hotelId,omitempty"`
	Code     string    `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	InDate   string    `protobuf:"bytes,3,opt,name=inDate" json:"inDate,omitempty"`
	OutDate  string    `protobuf:"bytes,4,opt,name=outDate" json:"outDate,omitempty"`
	RoomType *RoomType `protobuf:"bytes,5,opt,name=roomType" json:"roomType,omitempty"`
}

func (m *RatePlan) Reset()                    { *m = RatePlan{} }
func (m *RatePlan) String() string            { return proto.CompactTextString(m) }
func (*RatePlan) ProtoMessage()               {}
func (*RatePlan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RatePlan) GetRoomType() *RoomType {
	if m != nil {
		return m.RoomType
	}
	return nil
}

type RoomType struct {
	BookableRate       float64 `protobuf:"fixed64,1,opt,name=bookableRate" json:"bookableRate,omitempty"`
	TotalRate          float64 `protobuf:"fixed64,2,opt,name=totalRate" json:"totalRate,omitempty"`
	TotalRateInclusive float64 `protobuf:"fixed64,3,opt,name=totalRateInclusive" json:"totalRateInclusive,omitempty"`
	Code               string  `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
	Currency           string  `protobuf:"bytes,5,opt,name=currency" json:"currency,omitempty"`
	RoomDescription    string  `protobuf:"bytes,6,opt,name=roomDescription" json:"roomDescription,omitempty"`
}

func (m *RoomType) Reset()                    { *m = RoomType{} }
func (m *RoomType) String() string            { return proto.CompactTextString(m) }
func (*RoomType) ProtoMessage()               {}
func (*RoomType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Request)(nil), "rate.Request")
	proto.RegisterType((*Result)(nil), "rate.Result")
	proto.RegisterType((*RatePlan)(nil), "rate.RatePlan")
	proto.RegisterType((*RoomType)(nil), "rate.RoomType")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Rate service

type RateClient interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error)
}

type rateClient struct {
	c           client.Client
	serviceName string
}

func NewRateClient(serviceName string, c client.Client) RateClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "rate"
	}
	return &rateClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *rateClient) GetRates(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.serviceName, "Rate.GetRates", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rate service

type RateHandler interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(context.Context, *Request, *Result) error
}

func RegisterRateHandler(s server.Server, hdlr RateHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Rate{hdlr}, opts...))
}

type Rate struct {
	RateHandler
}

func (h *Rate) GetRates(ctx context.Context, in *Request, out *Result) error {
	return h.RateHandler.GetRates(ctx, in, out)
}

func init() { proto.RegisterFile("srv/rate/proto/rate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0xc7, 0x89, 0xab, 0x5d, 0x7b, 0x9c, 0x0a, 0xe7, 0x42, 0xe2, 0xf0, 0x62, 0xf4, 0xc6, 0x21,
	0xb2, 0xc1, 0x04, 0x9f, 0x60, 0x20, 0xbb, 0x93, 0x83, 0xe0, 0x75, 0xd7, 0x05, 0x2c, 0xc6, 0x66,
	0x26, 0xe9, 0x60, 0x2f, 0xe2, 0xa3, 0xf9, 0x3c, 0x92, 0x34, 0xed, 0x3a, 0xd1, 0xbb, 0xff, 0x47,
	0x38, 0xfd, 0x25, 0x3d, 0x70, 0x6d, 0xf4, 0x6e, 0xae, 0x73, 0x2b, 0xe6, 0x5b, 0xad, 0xac, 0xf2,
	0x72, 0xe6, 0x25, 0x46, 0x4e, 0x67, 0xaf, 0x30, 0x24, 0xf1, 0x59, 0x0b, 0x63, 0x71, 0x0c, 0xc9,
	0x9b, 0xb2, 0x42, 0xae, 0x36, 0x86, 0xb3, 0xc9, 0x60, 0x9a, 0x52, 0xe7, 0xf1, 0x0a, 0xe2, 0xb2,
	0x5a, 0xe6, 0x56, 0xf0, 0x93, 0x09, 0x9b, 0xa6, 0x14, 0x1c, 0x72, 0x18, 0xaa, 0xda, 0xfa, 0x62,
	0xe0, 0x8b, 0xd6, 0x66, 0x8f, 0x10, 0x93, 0x30, 0xb5, 0xb4, 0x78, 0x0f, 0xa9, 0xfb, 0xd4, 0xb3,
	0xcc, 0xab, 0x66, 0xf0, 0xd9, 0xe2, 0x62, 0xe6, 0x41, 0x28, 0xc4, 0x74, 0x38, 0x90, 0x7d, 0x31,
	0x48, 0xda, 0xdc, 0x8d, 0x0f, 0x08, 0x9c, 0x35, 0xe3, 0x83, 0x45, 0x84, 0xa8, 0x50, 0x9b, 0x16,
	0xc7, 0xeb, 0x1e, 0xe4, 0xe0, 0x3f, 0xc8, 0xe8, 0x08, 0x12, 0xef, 0x20, 0xd1, 0x4a, 0x7d, 0xbc,
	0xec, 0xb7, 0x82, 0x9f, 0x4e, 0x58, 0x8f, 0x2c, 0xa4, 0xd4, 0xf5, 0xd9, 0xb7, 0x03, 0x0b, 0x06,
	0x33, 0x18, 0xad, 0x95, 0x7a, 0xcf, 0xd7, 0x52, 0x38, 0x58, 0x4f, 0xc7, 0xe8, 0x28, 0xc3, 0x1b,
	0x48, 0xad, 0xb2, 0xb9, 0xa4, 0xf6, 0xd9, 0x18, 0x1d, 0x02, 0x9c, 0x01, 0x76, 0x66, 0x55, 0x15,
	0xb2, 0x36, 0xe5, 0xae, 0x01, 0x67, 0xf4, 0x47, 0xd3, 0x5d, 0x38, 0xea, 0x5d, 0x78, 0x0c, 0x49,
	0x51, 0x6b, 0x2d, 0xaa, 0x62, 0xef, 0xf1, 0x53, 0xea, 0x3c, 0x4e, 0xe1, 0xd2, 0xa1, 0x2f, 0x85,
	0x29, 0x74, 0xb9, 0xb5, 0xa5, 0xaa, 0x78, 0xec, 0x8f, 0xfc, 0x8e, 0x17, 0x73, 0x88, 0x3c, 0xd1,
	0x2d, 0x24, 0x4f, 0xc2, 0x3a, 0x69, 0xf0, 0x3c, 0x3c, 0x43, 0xb3, 0x1a, 0xe3, 0x51, 0x6b, 0xdd,
	0x0f, 0x5d, 0xc7, 0x7e, 0x81, 0x1e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xf4, 0x9f, 0xd8,
	0x5d, 0x02, 0x00, 0x00,
}
