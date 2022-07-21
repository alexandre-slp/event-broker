// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"

	v1 "github.com/alexandre-slp/event-broker/infra/api/v1/serializer"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HealthcheckClient is the client API for Healthcheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthcheckClient interface {
	GetHealthCheck(ctx context.Context, in *v1.HealthCheckRequest, opts ...grpc.CallOption) (*v1.HealthCheckResponse, error)
}

type healthcheckClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthcheckClient(cc grpc.ClientConnInterface) HealthcheckClient {
	return &healthcheckClient{cc}
}

func (c *healthcheckClient) GetHealthCheck(ctx context.Context, in *v1.HealthCheckRequest, opts ...grpc.CallOption) (*v1.HealthCheckResponse, error) {
	out := new(v1.HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/healthcheck.Healthcheck/GetHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthcheckServer is the server API for Healthcheck service.
// All implementations must embed UnimplementedHealthcheckServer
// for forward compatibility
type HealthcheckServer interface {
	GetHealthCheck(context.Context, *v1.HealthCheckRequest) (*v1.HealthCheckResponse, error)
	mustEmbedUnimplementedHealthcheckServer()
}

// UnimplementedHealthcheckServer must be embedded to have forward compatible implementations.
type UnimplementedHealthcheckServer struct {
}

func (UnimplementedHealthcheckServer) GetHealthCheck(context.Context, *v1.HealthCheckRequest) (*v1.HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealthCheck not implemented")
}
func (UnimplementedHealthcheckServer) mustEmbedUnimplementedHealthcheckServer() {}

// UnsafeHealthcheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthcheckServer will
// result in compilation errors.
type UnsafeHealthcheckServer interface {
	mustEmbedUnimplementedHealthcheckServer()
}

func RegisterHealthcheckServer(s grpc.ServiceRegistrar, srv HealthcheckServer) {
	s.RegisterService(&Healthcheck_ServiceDesc, srv)
}

func _Healthcheck_GetHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthcheckServer).GetHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/healthcheck.Healthcheck/GetHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthcheckServer).GetHealthCheck(ctx, req.(*v1.HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Healthcheck_ServiceDesc is the grpc.ServiceDesc for Healthcheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Healthcheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "healthcheck.Healthcheck",
	HandlerType: (*HealthcheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealthCheck",
			Handler:    _Healthcheck_GetHealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "healthcheck.proto",
}