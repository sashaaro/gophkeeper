// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: pkg/gophkeeper/gophkeeper.proto

package gophkeeper

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Keeper_Auth_FullMethodName              = "/gophkeeper.v1.Keeper/Auth"
	Keeper_Register_FullMethodName          = "/gophkeeper.v1.Keeper/Register"
	Keeper_GetAll_FullMethodName            = "/gophkeeper.v1.Keeper/GetAll"
	Keeper_CreateCredentials_FullMethodName = "/gophkeeper.v1.Keeper/CreateCredentials"
	Keeper_CreateCreditCard_FullMethodName  = "/gophkeeper.v1.Keeper/CreateCreditCard"
	Keeper_CreateText_FullMethodName        = "/gophkeeper.v1.Keeper/CreateText"
	Keeper_CreateBinary_FullMethodName      = "/gophkeeper.v1.Keeper/CreateBinary"
	Keeper_SendData_FullMethodName          = "/gophkeeper.v1.Keeper/SendData"
	Keeper_ReceiveData_FullMethodName       = "/gophkeeper.v1.Keeper/ReceiveData"
	Keeper_Ping_FullMethodName              = "/gophkeeper.v1.Keeper/Ping"
)

// KeeperClient is the client API for Keeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeeperClient interface {
	Auth(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Empty, error)
	Register(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Empty, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EntityList, error)
	CreateCredentials(ctx context.Context, in *CreateCreadentialsReq, opts ...grpc.CallOption) (*Entity, error)
	CreateCreditCard(ctx context.Context, in *CreateCreditCardReq, opts ...grpc.CallOption) (*Entity, error)
	CreateText(ctx context.Context, in *CreateTextReq, opts ...grpc.CallOption) (*Entity, error)
	CreateBinary(ctx context.Context, in *CreateBinaryReq, opts ...grpc.CallOption) (*Entity, error)
	SendData(ctx context.Context, opts ...grpc.CallOption) (Keeper_SendDataClient, error)
	ReceiveData(ctx context.Context, in *Entity, opts ...grpc.CallOption) (Keeper_ReceiveDataClient, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type keeperClient struct {
	cc grpc.ClientConnInterface
}

func NewKeeperClient(cc grpc.ClientConnInterface) KeeperClient {
	return &keeperClient{cc}
}

func (c *keeperClient) Auth(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Keeper_Auth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) Register(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Keeper_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EntityList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EntityList)
	err := c.cc.Invoke(ctx, Keeper_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) CreateCredentials(ctx context.Context, in *CreateCreadentialsReq, opts ...grpc.CallOption) (*Entity, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Entity)
	err := c.cc.Invoke(ctx, Keeper_CreateCredentials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) CreateCreditCard(ctx context.Context, in *CreateCreditCardReq, opts ...grpc.CallOption) (*Entity, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Entity)
	err := c.cc.Invoke(ctx, Keeper_CreateCreditCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) CreateText(ctx context.Context, in *CreateTextReq, opts ...grpc.CallOption) (*Entity, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Entity)
	err := c.cc.Invoke(ctx, Keeper_CreateText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) CreateBinary(ctx context.Context, in *CreateBinaryReq, opts ...grpc.CallOption) (*Entity, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Entity)
	err := c.cc.Invoke(ctx, Keeper_CreateBinary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperClient) SendData(ctx context.Context, opts ...grpc.CallOption) (Keeper_SendDataClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Keeper_ServiceDesc.Streams[0], Keeper_SendData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &keeperSendDataClient{ClientStream: stream}
	return x, nil
}

type Keeper_SendDataClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type keeperSendDataClient struct {
	grpc.ClientStream
}

func (x *keeperSendDataClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *keeperSendDataClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keeperClient) ReceiveData(ctx context.Context, in *Entity, opts ...grpc.CallOption) (Keeper_ReceiveDataClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Keeper_ServiceDesc.Streams[1], Keeper_ReceiveData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &keeperReceiveDataClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Keeper_ReceiveDataClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type keeperReceiveDataClient struct {
	grpc.ClientStream
}

func (x *keeperReceiveDataClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keeperClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Keeper_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeeperServer is the server API for Keeper service.
// All implementations must embed UnimplementedKeeperServer
// for forward compatibility
type KeeperServer interface {
	Auth(context.Context, *Credentials) (*Empty, error)
	Register(context.Context, *Credentials) (*Empty, error)
	GetAll(context.Context, *Empty) (*EntityList, error)
	CreateCredentials(context.Context, *CreateCreadentialsReq) (*Entity, error)
	CreateCreditCard(context.Context, *CreateCreditCardReq) (*Entity, error)
	CreateText(context.Context, *CreateTextReq) (*Entity, error)
	CreateBinary(context.Context, *CreateBinaryReq) (*Entity, error)
	SendData(Keeper_SendDataServer) error
	ReceiveData(*Entity, Keeper_ReceiveDataServer) error
	Ping(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedKeeperServer()
}

// UnimplementedKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedKeeperServer struct {
}

func (UnimplementedKeeperServer) Auth(context.Context, *Credentials) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedKeeperServer) Register(context.Context, *Credentials) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedKeeperServer) GetAll(context.Context, *Empty) (*EntityList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedKeeperServer) CreateCredentials(context.Context, *CreateCreadentialsReq) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCredentials not implemented")
}
func (UnimplementedKeeperServer) CreateCreditCard(context.Context, *CreateCreditCardReq) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCreditCard not implemented")
}
func (UnimplementedKeeperServer) CreateText(context.Context, *CreateTextReq) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateText not implemented")
}
func (UnimplementedKeeperServer) CreateBinary(context.Context, *CreateBinaryReq) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBinary not implemented")
}
func (UnimplementedKeeperServer) SendData(Keeper_SendDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SendData not implemented")
}
func (UnimplementedKeeperServer) ReceiveData(*Entity, Keeper_ReceiveDataServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveData not implemented")
}
func (UnimplementedKeeperServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedKeeperServer) mustEmbedUnimplementedKeeperServer() {}

// UnsafeKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeeperServer will
// result in compilation errors.
type UnsafeKeeperServer interface {
	mustEmbedUnimplementedKeeperServer()
}

func RegisterKeeperServer(s grpc.ServiceRegistrar, srv KeeperServer) {
	s.RegisterService(&Keeper_ServiceDesc, srv)
}

func _Keeper_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).Auth(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).Register(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_CreateCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCreadentialsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).CreateCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_CreateCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).CreateCredentials(ctx, req.(*CreateCreadentialsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_CreateCreditCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCreditCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).CreateCreditCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_CreateCreditCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).CreateCreditCard(ctx, req.(*CreateCreditCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_CreateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTextReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).CreateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_CreateText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).CreateText(ctx, req.(*CreateTextReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_CreateBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBinaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).CreateBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_CreateBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).CreateBinary(ctx, req.(*CreateBinaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keeper_SendData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KeeperServer).SendData(&keeperSendDataServer{ServerStream: stream})
}

type Keeper_SendDataServer interface {
	SendAndClose(*Empty) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type keeperSendDataServer struct {
	grpc.ServerStream
}

func (x *keeperSendDataServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *keeperSendDataServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Keeper_ReceiveData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Entity)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeeperServer).ReceiveData(m, &keeperReceiveDataServer{ServerStream: stream})
}

type Keeper_ReceiveDataServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type keeperReceiveDataServer struct {
	grpc.ServerStream
}

func (x *keeperReceiveDataServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Keeper_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Keeper_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Keeper_ServiceDesc is the grpc.ServiceDesc for Keeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Keeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.v1.Keeper",
	HandlerType: (*KeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Keeper_Auth_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Keeper_Register_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Keeper_GetAll_Handler,
		},
		{
			MethodName: "CreateCredentials",
			Handler:    _Keeper_CreateCredentials_Handler,
		},
		{
			MethodName: "CreateCreditCard",
			Handler:    _Keeper_CreateCreditCard_Handler,
		},
		{
			MethodName: "CreateText",
			Handler:    _Keeper_CreateText_Handler,
		},
		{
			MethodName: "CreateBinary",
			Handler:    _Keeper_CreateBinary_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Keeper_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendData",
			Handler:       _Keeper_SendData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveData",
			Handler:       _Keeper_ReceiveData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/gophkeeper/gophkeeper.proto",
}