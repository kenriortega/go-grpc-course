// Code generated by protoc-gen-go. DO NOT EDIT.
// source: handOne/greet/greetpb/greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca996a35e733899f, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca996a35e733899f, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca996a35e733899f, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetManyTimesRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetManyTimesRequest) Reset()         { *m = GreetManyTimesRequest{} }
func (m *GreetManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesRequest) ProtoMessage()    {}
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca996a35e733899f, []int{3}
}

func (m *GreetManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesRequest.Unmarshal(m, b)
}
func (m *GreetManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesRequest.Merge(m, src)
}
func (m *GreetManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesRequest.Size(m)
}
func (m *GreetManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesRequest proto.InternalMessageInfo

func (m *GreetManyTimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreeManyTimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreeManyTimesResponse) Reset()         { *m = GreeManyTimesResponse{} }
func (m *GreeManyTimesResponse) String() string { return proto.CompactTextString(m) }
func (*GreeManyTimesResponse) ProtoMessage()    {}
func (*GreeManyTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca996a35e733899f, []int{4}
}

func (m *GreeManyTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreeManyTimesResponse.Unmarshal(m, b)
}
func (m *GreeManyTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreeManyTimesResponse.Marshal(b, m, deterministic)
}
func (m *GreeManyTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreeManyTimesResponse.Merge(m, src)
}
func (m *GreeManyTimesResponse) XXX_Size() int {
	return xxx_messageInfo_GreeManyTimesResponse.Size(m)
}
func (m *GreeManyTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreeManyTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreeManyTimesResponse proto.InternalMessageInfo

func (m *GreeManyTimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type LongGreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LongGreetRequest) Reset()         { *m = LongGreetRequest{} }
func (m *LongGreetRequest) String() string { return proto.CompactTextString(m) }
func (*LongGreetRequest) ProtoMessage()    {}
func (*LongGreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca996a35e733899f, []int{5}
}

func (m *LongGreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetRequest.Unmarshal(m, b)
}
func (m *LongGreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetRequest.Marshal(b, m, deterministic)
}
func (m *LongGreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetRequest.Merge(m, src)
}
func (m *LongGreetRequest) XXX_Size() int {
	return xxx_messageInfo_LongGreetRequest.Size(m)
}
func (m *LongGreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetRequest proto.InternalMessageInfo

func (m *LongGreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type LongGreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongGreetResponse) Reset()         { *m = LongGreetResponse{} }
func (m *LongGreetResponse) String() string { return proto.CompactTextString(m) }
func (*LongGreetResponse) ProtoMessage()    {}
func (*LongGreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca996a35e733899f, []int{6}
}

func (m *LongGreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetResponse.Unmarshal(m, b)
}
func (m *LongGreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetResponse.Marshal(b, m, deterministic)
}
func (m *LongGreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetResponse.Merge(m, src)
}
func (m *LongGreetResponse) XXX_Size() int {
	return xxx_messageInfo_LongGreetResponse.Size(m)
}
func (m *LongGreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetResponse proto.InternalMessageInfo

func (m *LongGreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetEveryoneRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetEveryoneRequest) Reset()         { *m = GreetEveryoneRequest{} }
func (m *GreetEveryoneRequest) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneRequest) ProtoMessage()    {}
func (*GreetEveryoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca996a35e733899f, []int{7}
}

func (m *GreetEveryoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneRequest.Unmarshal(m, b)
}
func (m *GreetEveryoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneRequest.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneRequest.Merge(m, src)
}
func (m *GreetEveryoneRequest) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneRequest.Size(m)
}
func (m *GreetEveryoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneRequest proto.InternalMessageInfo

func (m *GreetEveryoneRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetEveryoneResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetEveryoneResponse) Reset()         { *m = GreetEveryoneResponse{} }
func (m *GreetEveryoneResponse) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneResponse) ProtoMessage()    {}
func (*GreetEveryoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca996a35e733899f, []int{8}
}

func (m *GreetEveryoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneResponse.Unmarshal(m, b)
}
func (m *GreetEveryoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneResponse.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneResponse.Merge(m, src)
}
func (m *GreetEveryoneResponse) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneResponse.Size(m)
}
func (m *GreetEveryoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneResponse proto.InternalMessageInfo

func (m *GreetEveryoneResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetManyTimesRequest)(nil), "greet.GreetManyTimesRequest")
	proto.RegisterType((*GreeManyTimesResponse)(nil), "greet.GreeManyTimesResponse")
	proto.RegisterType((*LongGreetRequest)(nil), "greet.LongGreetRequest")
	proto.RegisterType((*LongGreetResponse)(nil), "greet.LongGreetResponse")
	proto.RegisterType((*GreetEveryoneRequest)(nil), "greet.GreetEveryoneRequest")
	proto.RegisterType((*GreetEveryoneResponse)(nil), "greet.GreetEveryoneResponse")
}

func init() { proto.RegisterFile("handOne/greet/greetpb/greet.proto", fileDescriptor_ca996a35e733899f) }

var fileDescriptor_ca996a35e733899f = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0xcd, 0x92, 0xc0, 0x8f, 0xce, 0xcf, 0xbf, 0x23, 0x2a, 0x01, 0x4d, 0xb4, 0x17, 0x49, 0x48,
	0x80, 0xe0, 0xd1, 0x83, 0xc6, 0xbf, 0x17, 0xc5, 0xa4, 0x7a, 0xf2, 0x62, 0x8a, 0x8e, 0xb5, 0x09,
	0xdd, 0xd6, 0xdd, 0x42, 0xc2, 0xf7, 0xf1, 0x83, 0x9a, 0x6e, 0xb7, 0x64, 0xdb, 0x26, 0x92, 0xe0,
	0x05, 0x3a, 0xf3, 0x5e, 0xdf, 0x9b, 0x9d, 0xd7, 0x85, 0xe3, 0x4f, 0x97, 0xbf, 0x3f, 0x72, 0xea,
	0x7b, 0x82, 0x28, 0x4e, 0x7f, 0xa3, 0x71, 0xfa, 0xdf, 0x8b, 0x44, 0x18, 0x87, 0x58, 0x55, 0x85,
	0x7d, 0x0b, 0xf5, 0xbb, 0xe4, 0xc1, 0xe7, 0x1e, 0x1e, 0x02, 0x7c, 0xf8, 0x42, 0xc6, 0xaf, 0xdc,
	0x0d, 0xa8, 0xc9, 0x8e, 0x58, 0xc7, 0x72, 0x2c, 0xd5, 0x19, 0xb9, 0x01, 0x61, 0x1b, 0xac, 0x89,
	0x9b, 0xa1, 0x15, 0x85, 0xd6, 0x93, 0x46, 0x02, 0xda, 0x67, 0xb0, 0xa6, 0x74, 0x1c, 0xfa, 0x9a,
	0x92, 0x8c, 0xb1, 0x0b, 0x75, 0x4f, 0xeb, 0x2a, 0xa5, 0xff, 0xc3, 0xcd, 0x5e, 0x6a, 0x9f, 0xd9,
	0x39, 0x0b, 0x82, 0x7d, 0x02, 0xeb, 0xfa, 0x65, 0x19, 0x85, 0x5c, 0x12, 0xee, 0x41, 0x4d, 0x90,
	0x9c, 0x4e, 0x62, 0x3d, 0x85, 0xae, 0xec, 0x6b, 0xd8, 0x55, 0xc4, 0x07, 0x97, 0xcf, 0x9f, 0xfd,
	0x80, 0xe4, 0x4a, 0x76, 0xfd, 0x54, 0xc5, 0x10, 0x59, 0x62, 0x7b, 0x0e, 0x5b, 0xf7, 0x21, 0xf7,
	0x56, 0x3f, 0x60, 0x17, 0xb6, 0x0d, 0x81, 0x25, 0x6e, 0x57, 0xd0, 0x50, 0xc4, 0x9b, 0x19, 0x89,
	0x79, 0xc8, 0xe9, 0x2f, 0x67, 0x34, 0x44, 0x7e, 0x77, 0x1d, 0x7e, 0x57, 0x74, 0x82, 0x4f, 0x24,
	0x66, 0xfe, 0x1b, 0xe1, 0x10, 0xaa, 0xaa, 0xc6, 0x1d, 0xd3, 0x45, 0x0f, 0xd3, 0x6a, 0xe4, 0x9b,
	0x5a, 0x7c, 0x04, 0x1b, 0xf9, 0x7c, 0xf0, 0xc0, 0xe4, 0x15, 0x63, 0x6b, 0x99, 0x68, 0x29, 0x8e,
	0x01, 0xc3, 0x0b, 0xb0, 0x16, 0x7b, 0xc3, 0x7d, 0x4d, 0x2e, 0x46, 0xd1, 0x6a, 0x96, 0x81, 0x54,
	0xa1, 0xc3, 0x70, 0xa4, 0x3f, 0xad, 0x6c, 0x0f, 0xd8, 0x36, 0x07, 0x2a, 0xac, 0x38, 0x37, 0x4f,
	0x69, 0x75, 0x1d, 0x36, 0x60, 0x97, 0xd6, 0xcb, 0x3f, 0x7d, 0x9b, 0xc6, 0x35, 0x75, 0x91, 0x4e,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x5b, 0x82, 0xda, 0x6d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	//Unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	// Server Streaming
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
	// Client Streaming
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error)
	// BiDi streaming
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Recv() (*GreeManyTimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreeManyTimesResponse, error) {
	m := new(GreeManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[1], "/greet.GreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceLongGreetClient{stream}
	return x, nil
}

type GreetService_LongGreetClient interface {
	Send(*LongGreetRequest) error
	CloseAndRecv() (*LongGreetResponse, error)
	grpc.ClientStream
}

type greetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceLongGreetClient) Send(m *LongGreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceLongGreetClient) CloseAndRecv() (*LongGreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongGreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[2], "/greet.GreetService/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetEveryoneClient{stream}
	return x, nil
}

type GreetService_GreetEveryoneClient interface {
	Send(*GreetEveryoneRequest) error
	Recv() (*GreetEveryoneResponse, error)
	grpc.ClientStream
}

type greetServiceGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetEveryoneClient) Send(m *GreetEveryoneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneClient) Recv() (*GreetEveryoneResponse, error) {
	m := new(GreetEveryoneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	//Unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	// Server Streaming
	GreetManyTimes(*GreetManyTimesRequest, GreetService_GreetManyTimesServer) error
	// Client Streaming
	LongGreet(GreetService_LongGreetServer) error
	// BiDi streaming
	GreetEveryone(GreetService_GreetEveryoneServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetManyTimes(req *GreetManyTimesRequest, srv GreetService_GreetManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimes not implemented")
}
func (*UnimplementedGreetServiceServer) LongGreet(srv GreetService_LongGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetEveryone(srv GreetService_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTimes(m, &greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreeManyTimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreeManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).LongGreet(&greetServiceLongGreetServer{stream})
}

type GreetService_LongGreetServer interface {
	SendAndClose(*LongGreetResponse) error
	Recv() (*LongGreetRequest, error)
	grpc.ServerStream
}

type greetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceLongGreetServer) SendAndClose(m *LongGreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceLongGreetServer) Recv() (*LongGreetRequest, error) {
	m := new(LongGreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetEveryone(&greetServiceGreetEveryoneServer{stream})
}

type GreetService_GreetEveryoneServer interface {
	Send(*GreetEveryoneResponse) error
	Recv() (*GreetEveryoneRequest, error)
	grpc.ServerStream
}

type greetServiceGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetEveryoneServer) Send(m *GreetEveryoneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneServer) Recv() (*GreetEveryoneRequest, error) {
	m := new(GreetEveryoneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongGreet",
			Handler:       _GreetService_LongGreet_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetService_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "handOne/greet/greetpb/greet.proto",
}
