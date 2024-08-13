// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: lifestyle_data.proto

package health_sync

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
	LifestyleService_Create_FullMethodName = "/health_sync.LifestyleService/Create"
	LifestyleService_Update_FullMethodName = "/health_sync.LifestyleService/Update"
	LifestyleService_Delete_FullMethodName = "/health_sync.LifestyleService/Delete"
	LifestyleService_Get_FullMethodName    = "/health_sync.LifestyleService/Get"
	LifestyleService_List_FullMethodName   = "/health_sync.LifestyleService/List"
)

// LifestyleServiceClient is the client API for LifestyleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LifestyleServiceClient interface {
	Create(ctx context.Context, in *LifestyleCreate, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *LifestyleUpdate, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*Lifestyle, error)
	List(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*LifestyleList, error)
}

type lifestyleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLifestyleServiceClient(cc grpc.ClientConnInterface) LifestyleServiceClient {
	return &lifestyleServiceClient{cc}
}

func (c *lifestyleServiceClient) Create(ctx context.Context, in *LifestyleCreate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, LifestyleService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifestyleServiceClient) Update(ctx context.Context, in *LifestyleUpdate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, LifestyleService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifestyleServiceClient) Delete(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, LifestyleService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifestyleServiceClient) Get(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*Lifestyle, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Lifestyle)
	err := c.cc.Invoke(ctx, LifestyleService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifestyleServiceClient) List(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*LifestyleList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LifestyleList)
	err := c.cc.Invoke(ctx, LifestyleService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LifestyleServiceServer is the server API for LifestyleService service.
// All implementations must embed UnimplementedLifestyleServiceServer
// for forward compatibility
type LifestyleServiceServer interface {
	Create(context.Context, *LifestyleCreate) (*Void, error)
	Update(context.Context, *LifestyleUpdate) (*Void, error)
	Delete(context.Context, *GetById) (*Void, error)
	Get(context.Context, *GetById) (*Lifestyle, error)
	List(context.Context, *GetById) (*LifestyleList, error)
	mustEmbedUnimplementedLifestyleServiceServer()
}

// UnimplementedLifestyleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLifestyleServiceServer struct {
}

func (UnimplementedLifestyleServiceServer) Create(context.Context, *LifestyleCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLifestyleServiceServer) Update(context.Context, *LifestyleUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLifestyleServiceServer) Delete(context.Context, *GetById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLifestyleServiceServer) Get(context.Context, *GetById) (*Lifestyle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLifestyleServiceServer) List(context.Context, *GetById) (*LifestyleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLifestyleServiceServer) mustEmbedUnimplementedLifestyleServiceServer() {}

// UnsafeLifestyleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LifestyleServiceServer will
// result in compilation errors.
type UnsafeLifestyleServiceServer interface {
	mustEmbedUnimplementedLifestyleServiceServer()
}

func RegisterLifestyleServiceServer(s grpc.ServiceRegistrar, srv LifestyleServiceServer) {
	s.RegisterService(&LifestyleService_ServiceDesc, srv)
}

func _LifestyleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LifestyleCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifestyleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifestyleService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifestyleServiceServer).Create(ctx, req.(*LifestyleCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifestyleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LifestyleUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifestyleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifestyleService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifestyleServiceServer).Update(ctx, req.(*LifestyleUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifestyleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifestyleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifestyleService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifestyleServiceServer).Delete(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifestyleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifestyleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifestyleService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifestyleServiceServer).Get(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifestyleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifestyleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifestyleService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifestyleServiceServer).List(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

// LifestyleService_ServiceDesc is the grpc.ServiceDesc for LifestyleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LifestyleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "health_sync.LifestyleService",
	HandlerType: (*LifestyleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LifestyleService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LifestyleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LifestyleService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LifestyleService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LifestyleService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lifestyle_data.proto",
}
