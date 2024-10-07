// protoc -I proto/ proto/plugin.proto --go-grpc_out=proto/ --proto_path=$GOPATH/bin

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: plugin.proto

package proto

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
	Driver_SetConfig_FullMethodName = "/proto.Driver/SetConfig"
	Driver_Start_FullMethodName     = "/proto.Driver/Start"
	Driver_Restart_FullMethodName   = "/proto.Driver/Restart"
	Driver_Stop_FullMethodName      = "/proto.Driver/Stop"
	Driver_Get_FullMethodName       = "/proto.Driver/Get"
	Driver_Set_FullMethodName       = "/proto.Driver/Set"
)

// DriverClient is the client API for Driver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverClient interface {
	// 宿主（client） --> 驱动（server）
	SetConfig(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error)
	Start(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error)
	Restart(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error)
	Stop(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error)
	Get(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error)
	Set(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error)
}

type driverClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverClient(cc grpc.ClientConnInterface) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) SetConfig(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseResult)
	err := c.cc.Invoke(ctx, Driver_SetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Start(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseResult)
	err := c.cc.Invoke(ctx, Driver_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Restart(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseResult)
	err := c.cc.Invoke(ctx, Driver_Restart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Stop(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseResult)
	err := c.cc.Invoke(ctx, Driver_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Get(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseResult)
	err := c.cc.Invoke(ctx, Driver_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Set(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseResult)
	err := c.cc.Invoke(ctx, Driver_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServer is the server API for Driver service.
// All implementations must embed UnimplementedDriverServer
// for forward compatibility.
type DriverServer interface {
	// 宿主（client） --> 驱动（server）
	SetConfig(context.Context, *RequestArgs) (*ResponseResult, error)
	Start(context.Context, *RequestArgs) (*ResponseResult, error)
	Restart(context.Context, *RequestArgs) (*ResponseResult, error)
	Stop(context.Context, *RequestArgs) (*ResponseResult, error)
	Get(context.Context, *RequestArgs) (*ResponseResult, error)
	Set(context.Context, *RequestArgs) (*ResponseResult, error)
}

// UnimplementedDriverServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDriverServer struct{}

func (UnimplementedDriverServer) SetConfig(context.Context, *RequestArgs) (*ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedDriverServer) Start(context.Context, *RequestArgs) (*ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedDriverServer) Restart(context.Context, *RequestArgs) (*ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedDriverServer) Stop(context.Context, *RequestArgs) (*ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedDriverServer) Get(context.Context, *RequestArgs) (*ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDriverServer) Set(context.Context, *RequestArgs) (*ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedDriverServer) testEmbeddedByValue()                {}

func RegisterDriverServer(s grpc.ServiceRegistrar, srv DriverServer) {
	// If the following call pancis, it indicates UnimplementedDriverServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Driver_ServiceDesc, srv)
}

func _Driver_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Driver_SetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).SetConfig(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Driver_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Start(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Driver_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Restart(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Driver_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Stop(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Driver_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Get(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Driver_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Set(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// Driver_ServiceDesc is the grpc.ServiceDesc for Driver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Driver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetConfig",
			Handler:    _Driver_SetConfig_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Driver_Start_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Driver_Restart_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Driver_Stop_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Driver_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Driver_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}

const (
	Report_Post_FullMethodName  = "/proto.Report/Post"
	Report_State_FullMethodName = "/proto.Report/State"
)

// ReportClient is the client API for Report service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportClient interface {
	// 驱动（client） --> 宿主（server）
	Post(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error)
	State(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error)
}

type reportClient struct {
	cc grpc.ClientConnInterface
}

func NewReportClient(cc grpc.ClientConnInterface) ReportClient {
	return &reportClient{cc}
}

func (c *reportClient) Post(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseResult)
	err := c.cc.Invoke(ctx, Report_Post_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportClient) State(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseResult)
	err := c.cc.Invoke(ctx, Report_State_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServer is the server API for Report service.
// All implementations must embed UnimplementedReportServer
// for forward compatibility.
type ReportServer interface {
	// 驱动（client） --> 宿主（server）
	Post(context.Context, *RequestArgs) (*ResponseResult, error)
	State(context.Context, *RequestArgs) (*ResponseResult, error)
	mustEmbedUnimplementedReportServer()
}

// UnimplementedReportServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReportServer struct{}

func (UnimplementedReportServer) Post(context.Context, *RequestArgs) (*ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedReportServer) State(context.Context, *RequestArgs) (*ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (UnimplementedReportServer) mustEmbedUnimplementedReportServer() {}
func (UnimplementedReportServer) testEmbeddedByValue()                {}

// UnsafeReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServer will
// result in compilation errors.
type UnsafeReportServer interface {
	mustEmbedUnimplementedReportServer()
}

func RegisterReportServer(s grpc.ServiceRegistrar, srv ReportServer) {
	// If the following call pancis, it indicates UnimplementedReportServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Report_ServiceDesc, srv)
}

func _Report_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Report_Post_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).Post(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Report_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Report_State_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).State(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// Report_ServiceDesc is the grpc.ServiceDesc for Report service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Report_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Report",
	HandlerType: (*ReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _Report_Post_Handler,
		},
		{
			MethodName: "State",
			Handler:    _Report_State_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}
