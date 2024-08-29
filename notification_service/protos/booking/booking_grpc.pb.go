// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: booking.proto

package booking

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BookHotel_Create_FullMethodName         = "/BookHotel/Create"
	BookHotel_Get_FullMethodName            = "/BookHotel/Get"
	BookHotel_Update_FullMethodName         = "/BookHotel/Update"
	BookHotel_Delete_FullMethodName         = "/BookHotel/Delete"
	BookHotel_CreateWaiting_FullMethodName  = "/BookHotel/CreateWaiting"
	BookHotel_GetWaitinglist_FullMethodName = "/BookHotel/GetWaitinglist"
	BookHotel_Getall_FullMethodName         = "/BookHotel/Getall"
	BookHotel_UpdateWaiting_FullMethodName  = "/BookHotel/UpdateWaiting"
	BookHotel_CancelWaiting_FullMethodName  = "/BookHotel/CancelWaiting"
)

// BookHotelClient is the client API for BookHotel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookHotelClient interface {
	Create(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error)
	Get(ctx context.Context, in *GetUsersBookRequest, opts ...grpc.CallOption) (*GetUsersBookResponse, error)
	Update(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error)
	Delete(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error)
	CreateWaiting(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error)
	GetWaitinglist(ctx context.Context, in *GetWaitinglistRequest, opts ...grpc.CallOption) (*GetWaitinglistResponse, error)
	Getall(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UpdateWaiting(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error)
	CancelWaiting(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error)
}

type bookHotelClient struct {
	cc grpc.ClientConnInterface
}

func NewBookHotelClient(cc grpc.ClientConnInterface) BookHotelClient {
	return &bookHotelClient{cc}
}

func (c *bookHotelClient) Create(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, BookHotel_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookHotelClient) Get(ctx context.Context, in *GetUsersBookRequest, opts ...grpc.CallOption) (*GetUsersBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUsersBookResponse)
	err := c.cc.Invoke(ctx, BookHotel_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookHotelClient) Update(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, BookHotel_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookHotelClient) Delete(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, BookHotel_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookHotelClient) CreateWaiting(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, BookHotel_CreateWaiting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookHotelClient) GetWaitinglist(ctx context.Context, in *GetWaitinglistRequest, opts ...grpc.CallOption) (*GetWaitinglistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWaitinglistResponse)
	err := c.cc.Invoke(ctx, BookHotel_GetWaitinglist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookHotelClient) Getall(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, BookHotel_Getall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookHotelClient) UpdateWaiting(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, BookHotel_UpdateWaiting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookHotelClient) CancelWaiting(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*GeneralResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, BookHotel_CancelWaiting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookHotelServer is the server API for BookHotel service.
// All implementations must embed UnimplementedBookHotelServer
// for forward compatibility.
type BookHotelServer interface {
	Create(context.Context, *Bytes) (*GeneralResponse, error)
	Get(context.Context, *GetUsersBookRequest) (*GetUsersBookResponse, error)
	Update(context.Context, *Bytes) (*GeneralResponse, error)
	Delete(context.Context, *Bytes) (*GeneralResponse, error)
	CreateWaiting(context.Context, *Bytes) (*GeneralResponse, error)
	GetWaitinglist(context.Context, *GetWaitinglistRequest) (*GetWaitinglistResponse, error)
	Getall(context.Context, *Request) (*Response, error)
	UpdateWaiting(context.Context, *Bytes) (*GeneralResponse, error)
	CancelWaiting(context.Context, *Bytes) (*GeneralResponse, error)
	mustEmbedUnimplementedBookHotelServer()
}

// UnimplementedBookHotelServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookHotelServer struct{}

func (UnimplementedBookHotelServer) Create(context.Context, *Bytes) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBookHotelServer) Get(context.Context, *GetUsersBookRequest) (*GetUsersBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBookHotelServer) Update(context.Context, *Bytes) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBookHotelServer) Delete(context.Context, *Bytes) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBookHotelServer) CreateWaiting(context.Context, *Bytes) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWaiting not implemented")
}
func (UnimplementedBookHotelServer) GetWaitinglist(context.Context, *GetWaitinglistRequest) (*GetWaitinglistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaitinglist not implemented")
}
func (UnimplementedBookHotelServer) Getall(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getall not implemented")
}
func (UnimplementedBookHotelServer) UpdateWaiting(context.Context, *Bytes) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWaiting not implemented")
}
func (UnimplementedBookHotelServer) CancelWaiting(context.Context, *Bytes) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelWaiting not implemented")
}
func (UnimplementedBookHotelServer) mustEmbedUnimplementedBookHotelServer() {}
func (UnimplementedBookHotelServer) testEmbeddedByValue()                   {}

// UnsafeBookHotelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookHotelServer will
// result in compilation errors.
type UnsafeBookHotelServer interface {
	mustEmbedUnimplementedBookHotelServer()
}

func RegisterBookHotelServer(s grpc.ServiceRegistrar, srv BookHotelServer) {
	// If the following call pancis, it indicates UnimplementedBookHotelServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookHotel_ServiceDesc, srv)
}

func _BookHotel_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookHotelServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookHotel_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookHotelServer).Create(ctx, req.(*Bytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookHotel_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookHotelServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookHotel_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookHotelServer).Get(ctx, req.(*GetUsersBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookHotel_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookHotelServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookHotel_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookHotelServer).Update(ctx, req.(*Bytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookHotel_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookHotelServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookHotel_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookHotelServer).Delete(ctx, req.(*Bytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookHotel_CreateWaiting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookHotelServer).CreateWaiting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookHotel_CreateWaiting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookHotelServer).CreateWaiting(ctx, req.(*Bytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookHotel_GetWaitinglist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWaitinglistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookHotelServer).GetWaitinglist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookHotel_GetWaitinglist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookHotelServer).GetWaitinglist(ctx, req.(*GetWaitinglistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookHotel_Getall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookHotelServer).Getall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookHotel_Getall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookHotelServer).Getall(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookHotel_UpdateWaiting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookHotelServer).UpdateWaiting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookHotel_UpdateWaiting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookHotelServer).UpdateWaiting(ctx, req.(*Bytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookHotel_CancelWaiting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookHotelServer).CancelWaiting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookHotel_CancelWaiting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookHotelServer).CancelWaiting(ctx, req.(*Bytes))
	}
	return interceptor(ctx, in, info, handler)
}

// BookHotel_ServiceDesc is the grpc.ServiceDesc for BookHotel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookHotel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BookHotel",
	HandlerType: (*BookHotelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BookHotel_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BookHotel_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BookHotel_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookHotel_Delete_Handler,
		},
		{
			MethodName: "CreateWaiting",
			Handler:    _BookHotel_CreateWaiting_Handler,
		},
		{
			MethodName: "GetWaitinglist",
			Handler:    _BookHotel_GetWaitinglist_Handler,
		},
		{
			MethodName: "Getall",
			Handler:    _BookHotel_Getall_Handler,
		},
		{
			MethodName: "UpdateWaiting",
			Handler:    _BookHotel_UpdateWaiting_Handler,
		},
		{
			MethodName: "CancelWaiting",
			Handler:    _BookHotel_CancelWaiting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}
