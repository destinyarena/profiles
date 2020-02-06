// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts.proto

package accounts

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

type IdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ProfileRequest struct {
	Discord              string   `protobuf:"bytes,1,opt,name=discord,proto3" json:"discord,omitempty"`
	Bungie               string   `protobuf:"bytes,2,opt,name=bungie,proto3" json:"bungie,omitempty"`
	Faceit               string   `protobuf:"bytes,3,opt,name=faceit,proto3" json:"faceit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileRequest) Reset()         { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()    {}
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{2}
}

func (m *ProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileRequest.Unmarshal(m, b)
}
func (m *ProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileRequest.Marshal(b, m, deterministic)
}
func (m *ProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRequest.Merge(m, src)
}
func (m *ProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ProfileRequest.Size(m)
}
func (m *ProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRequest proto.InternalMessageInfo

func (m *ProfileRequest) GetDiscord() string {
	if m != nil {
		return m.Discord
	}
	return ""
}

func (m *ProfileRequest) GetBungie() string {
	if m != nil {
		return m.Bungie
	}
	return ""
}

func (m *ProfileRequest) GetFaceit() string {
	if m != nil {
		return m.Faceit
	}
	return ""
}

type ProfileReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Discord              string   `protobuf:"bytes,2,opt,name=discord,proto3" json:"discord,omitempty"`
	Bungie               string   `protobuf:"bytes,3,opt,name=bungie,proto3" json:"bungie,omitempty"`
	Faceit               string   `protobuf:"bytes,4,opt,name=faceit,proto3" json:"faceit,omitempty"`
	Banned               bool     `protobuf:"varint,5,opt,name=banned,proto3" json:"banned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileReply) Reset()         { *m = ProfileReply{} }
func (m *ProfileReply) String() string { return proto.CompactTextString(m) }
func (*ProfileReply) ProtoMessage()    {}
func (*ProfileReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{3}
}

func (m *ProfileReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileReply.Unmarshal(m, b)
}
func (m *ProfileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileReply.Marshal(b, m, deterministic)
}
func (m *ProfileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileReply.Merge(m, src)
}
func (m *ProfileReply) XXX_Size() int {
	return xxx_messageInfo_ProfileReply.Size(m)
}
func (m *ProfileReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileReply proto.InternalMessageInfo

func (m *ProfileReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProfileReply) GetDiscord() string {
	if m != nil {
		return m.Discord
	}
	return ""
}

func (m *ProfileReply) GetBungie() string {
	if m != nil {
		return m.Bungie
	}
	return ""
}

func (m *ProfileReply) GetFaceit() string {
	if m != nil {
		return m.Faceit
	}
	return ""
}

func (m *ProfileReply) GetBanned() bool {
	if m != nil {
		return m.Banned
	}
	return false
}

func init() {
	proto.RegisterType((*IdRequest)(nil), "accounts.IdRequest")
	proto.RegisterType((*Empty)(nil), "accounts.Empty")
	proto.RegisterType((*ProfileRequest)(nil), "accounts.ProfileRequest")
	proto.RegisterType((*ProfileReply)(nil), "accounts.ProfileReply")
}

func init() { proto.RegisterFile("accounts.proto", fileDescriptor_e1e7723af4c007b7) }

var fileDescriptor_e1e7723af4c007b7 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xa4,
	0xb9, 0x38, 0x3d, 0x53, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32,
	0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x94, 0xd8, 0xb9, 0x58, 0x5d,
	0x73, 0x0b, 0x4a, 0x2a, 0x95, 0xa2, 0xb8, 0xf8, 0x02, 0x8a, 0xf2, 0xd3, 0x32, 0x73, 0x52, 0x61,
	0x4a, 0x25, 0xb8, 0xd8, 0x53, 0x32, 0x8b, 0x93, 0xf3, 0x8b, 0x60, 0xea, 0x61, 0x5c, 0x21, 0x31,
	0x2e, 0xb6, 0xa4, 0xd2, 0xbc, 0xf4, 0xcc, 0x54, 0x09, 0x26, 0xb0, 0x04, 0x94, 0x07, 0x12, 0x4f,
	0x4b, 0x4c, 0x4e, 0xcd, 0x2c, 0x91, 0x60, 0x86, 0x88, 0x43, 0x78, 0x4a, 0x0d, 0x8c, 0x5c, 0x3c,
	0x70, 0xc3, 0x0b, 0x72, 0x2a, 0xd1, 0x5d, 0x81, 0x6c, 0x15, 0x13, 0x2e, 0xab, 0x98, 0x71, 0x58,
	0xc5, 0x82, 0x6c, 0x15, 0x58, 0x7d, 0x62, 0x5e, 0x5e, 0x6a, 0x8a, 0x04, 0xab, 0x02, 0xa3, 0x06,
	0x47, 0x10, 0x94, 0x67, 0xb4, 0x9c, 0x89, 0x8b, 0xcb, 0x11, 0x12, 0x22, 0x99, 0x79, 0xe9, 0x42,
	0xd6, 0x5c, 0x5c, 0xee, 0xa9, 0x25, 0x50, 0x37, 0x09, 0x09, 0xeb, 0xc1, 0x03, 0x0f, 0x1e, 0x52,
	0x52, 0x62, 0x08, 0x41, 0x64, 0xb7, 0x2b, 0x31, 0x08, 0x39, 0x73, 0xf1, 0x3a, 0x17, 0xa5, 0x26,
	0x96, 0xa4, 0xc2, 0xf4, 0x4b, 0x60, 0x51, 0x4a, 0xc8, 0x10, 0x73, 0x2e, 0xde, 0xa0, 0xd4, 0xdc,
	0xfc, 0xb2, 0x54, 0xbc, 0x8e, 0xe0, 0x47, 0x08, 0x42, 0xa2, 0x89, 0x41, 0x48, 0x97, 0x8b, 0xd9,
	0x29, 0x31, 0x8f, 0x68, 0xe5, 0xfa, 0x5c, 0xac, 0xa1, 0x79, 0x49, 0xc4, 0x6b, 0x48, 0x62, 0x03,
	0xa7, 0x1f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x4a, 0x92, 0x2b, 0x51, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountingClient is the client API for Accounting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountingClient interface {
	GetProfile(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	CreateProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	RemoveProfile(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error)
	Ban(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error)
	Unban(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error)
}

type accountingClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingClient(cc grpc.ClientConnInterface) AccountingClient {
	return &accountingClient{cc}
}

func (c *accountingClient) GetProfile(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/accounts.Accounting/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) CreateProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/accounts.Accounting/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) RemoveProfile(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/accounts.Accounting/RemoveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Ban(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/accounts.Accounting/Ban", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Unban(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/accounts.Accounting/Unban", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServer is the server API for Accounting service.
type AccountingServer interface {
	GetProfile(context.Context, *IdRequest) (*ProfileReply, error)
	CreateProfile(context.Context, *ProfileRequest) (*ProfileReply, error)
	RemoveProfile(context.Context, *IdRequest) (*Empty, error)
	Ban(context.Context, *IdRequest) (*Empty, error)
	Unban(context.Context, *IdRequest) (*Empty, error)
}

// UnimplementedAccountingServer can be embedded to have forward compatible implementations.
type UnimplementedAccountingServer struct {
}

func (*UnimplementedAccountingServer) GetProfile(ctx context.Context, req *IdRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (*UnimplementedAccountingServer) CreateProfile(ctx context.Context, req *ProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (*UnimplementedAccountingServer) RemoveProfile(ctx context.Context, req *IdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProfile not implemented")
}
func (*UnimplementedAccountingServer) Ban(ctx context.Context, req *IdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ban not implemented")
}
func (*UnimplementedAccountingServer) Unban(ctx context.Context, req *IdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unban not implemented")
}

func RegisterAccountingServer(s *grpc.Server, srv AccountingServer) {
	s.RegisterService(&_Accounting_serviceDesc, srv)
}

func _Accounting_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounting/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).GetProfile(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounting/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).CreateProfile(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_RemoveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).RemoveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounting/RemoveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).RemoveProfile(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Ban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Ban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounting/Ban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Ban(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Unban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Unban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounting/Unban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Unban(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.Accounting",
	HandlerType: (*AccountingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _Accounting_GetProfile_Handler,
		},
		{
			MethodName: "CreateProfile",
			Handler:    _Accounting_CreateProfile_Handler,
		},
		{
			MethodName: "RemoveProfile",
			Handler:    _Accounting_RemoveProfile_Handler,
		},
		{
			MethodName: "Ban",
			Handler:    _Accounting_Ban_Handler,
		},
		{
			MethodName: "Unban",
			Handler:    _Accounting_Unban_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}