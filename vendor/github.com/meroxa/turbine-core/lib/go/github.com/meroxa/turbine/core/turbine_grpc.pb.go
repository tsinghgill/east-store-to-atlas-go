// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: turbine/v1/turbine.proto

package core

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TurbineServiceClient is the client API for TurbineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TurbineServiceClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error)
	ReadCollection(ctx context.Context, in *ReadCollectionRequest, opts ...grpc.CallOption) (*Collection, error)
	WriteCollectionToResource(ctx context.Context, in *WriteCollectionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddProcessToCollection(ctx context.Context, in *ProcessCollectionRequest, opts ...grpc.CallOption) (*Collection, error)
	RegisterSecret(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HasFunctions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	ListResources(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListResourcesResponse, error)
	GetSpec(ctx context.Context, in *GetSpecRequest, opts ...grpc.CallOption) (*GetSpecResponse, error)
}

type turbineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTurbineServiceClient(cc grpc.ClientConnInterface) TurbineServiceClient {
	return &turbineServiceClient{cc}
}

func (c *turbineServiceClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/turbine_core.TurbineService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turbineServiceClient) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/turbine_core.TurbineService/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turbineServiceClient) ReadCollection(ctx context.Context, in *ReadCollectionRequest, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, "/turbine_core.TurbineService/ReadCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turbineServiceClient) WriteCollectionToResource(ctx context.Context, in *WriteCollectionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/turbine_core.TurbineService/WriteCollectionToResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turbineServiceClient) AddProcessToCollection(ctx context.Context, in *ProcessCollectionRequest, opts ...grpc.CallOption) (*Collection, error) {
	out := new(Collection)
	err := c.cc.Invoke(ctx, "/turbine_core.TurbineService/AddProcessToCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turbineServiceClient) RegisterSecret(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/turbine_core.TurbineService/RegisterSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turbineServiceClient) HasFunctions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/turbine_core.TurbineService/HasFunctions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turbineServiceClient) ListResources(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	out := new(ListResourcesResponse)
	err := c.cc.Invoke(ctx, "/turbine_core.TurbineService/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turbineServiceClient) GetSpec(ctx context.Context, in *GetSpecRequest, opts ...grpc.CallOption) (*GetSpecResponse, error) {
	out := new(GetSpecResponse)
	err := c.cc.Invoke(ctx, "/turbine_core.TurbineService/GetSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TurbineServiceServer is the server API for TurbineService service.
// All implementations should embed UnimplementedTurbineServiceServer
// for forward compatibility
type TurbineServiceServer interface {
	Init(context.Context, *InitRequest) (*emptypb.Empty, error)
	GetResource(context.Context, *GetResourceRequest) (*Resource, error)
	ReadCollection(context.Context, *ReadCollectionRequest) (*Collection, error)
	WriteCollectionToResource(context.Context, *WriteCollectionRequest) (*emptypb.Empty, error)
	AddProcessToCollection(context.Context, *ProcessCollectionRequest) (*Collection, error)
	RegisterSecret(context.Context, *Secret) (*emptypb.Empty, error)
	HasFunctions(context.Context, *emptypb.Empty) (*wrapperspb.BoolValue, error)
	ListResources(context.Context, *emptypb.Empty) (*ListResourcesResponse, error)
	GetSpec(context.Context, *GetSpecRequest) (*GetSpecResponse, error)
}

// UnimplementedTurbineServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTurbineServiceServer struct {
}

func (UnimplementedTurbineServiceServer) Init(context.Context, *InitRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedTurbineServiceServer) GetResource(context.Context, *GetResourceRequest) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResource not implemented")
}
func (UnimplementedTurbineServiceServer) ReadCollection(context.Context, *ReadCollectionRequest) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCollection not implemented")
}
func (UnimplementedTurbineServiceServer) WriteCollectionToResource(context.Context, *WriteCollectionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteCollectionToResource not implemented")
}
func (UnimplementedTurbineServiceServer) AddProcessToCollection(context.Context, *ProcessCollectionRequest) (*Collection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProcessToCollection not implemented")
}
func (UnimplementedTurbineServiceServer) RegisterSecret(context.Context, *Secret) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSecret not implemented")
}
func (UnimplementedTurbineServiceServer) HasFunctions(context.Context, *emptypb.Empty) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasFunctions not implemented")
}
func (UnimplementedTurbineServiceServer) ListResources(context.Context, *emptypb.Empty) (*ListResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}
func (UnimplementedTurbineServiceServer) GetSpec(context.Context, *GetSpecRequest) (*GetSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpec not implemented")
}

// UnsafeTurbineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TurbineServiceServer will
// result in compilation errors.
type UnsafeTurbineServiceServer interface {
	mustEmbedUnimplementedTurbineServiceServer()
}

func RegisterTurbineServiceServer(s grpc.ServiceRegistrar, srv TurbineServiceServer) {
	s.RegisterService(&TurbineService_ServiceDesc, srv)
}

func _TurbineService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurbineServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turbine_core.TurbineService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurbineServiceServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TurbineService_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurbineServiceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turbine_core.TurbineService/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurbineServiceServer).GetResource(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TurbineService_ReadCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurbineServiceServer).ReadCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turbine_core.TurbineService/ReadCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurbineServiceServer).ReadCollection(ctx, req.(*ReadCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TurbineService_WriteCollectionToResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurbineServiceServer).WriteCollectionToResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turbine_core.TurbineService/WriteCollectionToResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurbineServiceServer).WriteCollectionToResource(ctx, req.(*WriteCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TurbineService_AddProcessToCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurbineServiceServer).AddProcessToCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turbine_core.TurbineService/AddProcessToCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurbineServiceServer).AddProcessToCollection(ctx, req.(*ProcessCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TurbineService_RegisterSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurbineServiceServer).RegisterSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turbine_core.TurbineService/RegisterSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurbineServiceServer).RegisterSecret(ctx, req.(*Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _TurbineService_HasFunctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurbineServiceServer).HasFunctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turbine_core.TurbineService/HasFunctions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurbineServiceServer).HasFunctions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TurbineService_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurbineServiceServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turbine_core.TurbineService/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurbineServiceServer).ListResources(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TurbineService_GetSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurbineServiceServer).GetSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/turbine_core.TurbineService/GetSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurbineServiceServer).GetSpec(ctx, req.(*GetSpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TurbineService_ServiceDesc is the grpc.ServiceDesc for TurbineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TurbineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "turbine_core.TurbineService",
	HandlerType: (*TurbineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _TurbineService_Init_Handler,
		},
		{
			MethodName: "GetResource",
			Handler:    _TurbineService_GetResource_Handler,
		},
		{
			MethodName: "ReadCollection",
			Handler:    _TurbineService_ReadCollection_Handler,
		},
		{
			MethodName: "WriteCollectionToResource",
			Handler:    _TurbineService_WriteCollectionToResource_Handler,
		},
		{
			MethodName: "AddProcessToCollection",
			Handler:    _TurbineService_AddProcessToCollection_Handler,
		},
		{
			MethodName: "RegisterSecret",
			Handler:    _TurbineService_RegisterSecret_Handler,
		},
		{
			MethodName: "HasFunctions",
			Handler:    _TurbineService_HasFunctions_Handler,
		},
		{
			MethodName: "ListResources",
			Handler:    _TurbineService_ListResources_Handler,
		},
		{
			MethodName: "GetSpec",
			Handler:    _TurbineService_GetSpec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "turbine/v1/turbine.proto",
}
