//
// DISCLAIMER
//
// Copyright 2021-2022 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: network.proto

package v1

import (
	context "context"
	v1 "github.com/arangodb-managed/apis/common/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NetworkService_GetAPIVersion_FullMethodName                            = "/arangodb.cloud.network.v1.NetworkService/GetAPIVersion"
	NetworkService_IsPrivateEndpointServiceFeatureAvailable_FullMethodName = "/arangodb.cloud.network.v1.NetworkService/IsPrivateEndpointServiceFeatureAvailable"
	NetworkService_GetPrivateEndpointService_FullMethodName                = "/arangodb.cloud.network.v1.NetworkService/GetPrivateEndpointService"
	NetworkService_GetPrivateEndpointServiceByDeploymentID_FullMethodName  = "/arangodb.cloud.network.v1.NetworkService/GetPrivateEndpointServiceByDeploymentID"
	NetworkService_CreatePrivateEndpointService_FullMethodName             = "/arangodb.cloud.network.v1.NetworkService/CreatePrivateEndpointService"
	NetworkService_UpdatePrivateEndpointService_FullMethodName             = "/arangodb.cloud.network.v1.NetworkService/UpdatePrivateEndpointService"
)

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// NetworkService is the API used to manage the network for a deployment.
type NetworkServiceClient interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error)
	// Checks if the private endpoint service feature is enabled and available.
	// Required permissions:
	// - network.privateendpointservice.get-feature on the deployment that is identified by the given deployment ID (if specified).
	// - network.privateendpointservice.get-feature on the organization that is identified by the given organization ID (if specified).
	// - None, authenticated only (if only the region ID is specified).
	IsPrivateEndpointServiceFeatureAvailable(ctx context.Context, in *IsPrivateEndpointServiceFeatureAvailableRequest, opts ...grpc.CallOption) (*IsPrivateEndpointServiceFeatureAvailableResult, error)
	// Fetch a private endpoint service by its ID.
	// Required permissions:
	// - network.privateendpointservice.get on the private endpoint service.
	GetPrivateEndpointService(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*PrivateEndpointService, error)
	// Fetch a private endpoint service by the deployment ID.
	// Required permissions:
	// - network.privateendpointservice.get-by-deployment-id on the deployment that owns the private endpoint service.
	GetPrivateEndpointServiceByDeploymentID(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*PrivateEndpointService, error)
	// Create a new private endpoint service.
	// Required permissions:
	// - network.privateendpointservice.create on the deployment that owns the private endpoint service.
	CreatePrivateEndpointService(ctx context.Context, in *PrivateEndpointService, opts ...grpc.CallOption) (*PrivateEndpointService, error)
	// Update the private endpoint service.
	// Required permissions:
	// - network.privateendpointservice.update on the private endpoint service.
	UpdatePrivateEndpointService(ctx context.Context, in *PrivateEndpointService, opts ...grpc.CallOption) (*v1.Empty, error)
}

type networkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkServiceClient(cc grpc.ClientConnInterface) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Version)
	err := c.cc.Invoke(ctx, NetworkService_GetAPIVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) IsPrivateEndpointServiceFeatureAvailable(ctx context.Context, in *IsPrivateEndpointServiceFeatureAvailableRequest, opts ...grpc.CallOption) (*IsPrivateEndpointServiceFeatureAvailableResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsPrivateEndpointServiceFeatureAvailableResult)
	err := c.cc.Invoke(ctx, NetworkService_IsPrivateEndpointServiceFeatureAvailable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) GetPrivateEndpointService(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*PrivateEndpointService, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateEndpointService)
	err := c.cc.Invoke(ctx, NetworkService_GetPrivateEndpointService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) GetPrivateEndpointServiceByDeploymentID(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*PrivateEndpointService, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateEndpointService)
	err := c.cc.Invoke(ctx, NetworkService_GetPrivateEndpointServiceByDeploymentID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) CreatePrivateEndpointService(ctx context.Context, in *PrivateEndpointService, opts ...grpc.CallOption) (*PrivateEndpointService, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateEndpointService)
	err := c.cc.Invoke(ctx, NetworkService_CreatePrivateEndpointService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) UpdatePrivateEndpointService(ctx context.Context, in *PrivateEndpointService, opts ...grpc.CallOption) (*v1.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Empty)
	err := c.cc.Invoke(ctx, NetworkService_UpdatePrivateEndpointService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
// All implementations must embed UnimplementedNetworkServiceServer
// for forward compatibility.
//
// NetworkService is the API used to manage the network for a deployment.
type NetworkServiceServer interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error)
	// Checks if the private endpoint service feature is enabled and available.
	// Required permissions:
	// - network.privateendpointservice.get-feature on the deployment that is identified by the given deployment ID (if specified).
	// - network.privateendpointservice.get-feature on the organization that is identified by the given organization ID (if specified).
	// - None, authenticated only (if only the region ID is specified).
	IsPrivateEndpointServiceFeatureAvailable(context.Context, *IsPrivateEndpointServiceFeatureAvailableRequest) (*IsPrivateEndpointServiceFeatureAvailableResult, error)
	// Fetch a private endpoint service by its ID.
	// Required permissions:
	// - network.privateendpointservice.get on the private endpoint service.
	GetPrivateEndpointService(context.Context, *v1.IDOptions) (*PrivateEndpointService, error)
	// Fetch a private endpoint service by the deployment ID.
	// Required permissions:
	// - network.privateendpointservice.get-by-deployment-id on the deployment that owns the private endpoint service.
	GetPrivateEndpointServiceByDeploymentID(context.Context, *v1.IDOptions) (*PrivateEndpointService, error)
	// Create a new private endpoint service.
	// Required permissions:
	// - network.privateendpointservice.create on the deployment that owns the private endpoint service.
	CreatePrivateEndpointService(context.Context, *PrivateEndpointService) (*PrivateEndpointService, error)
	// Update the private endpoint service.
	// Required permissions:
	// - network.privateendpointservice.update on the private endpoint service.
	UpdatePrivateEndpointService(context.Context, *PrivateEndpointService) (*v1.Empty, error)
	mustEmbedUnimplementedNetworkServiceServer()
}

// UnimplementedNetworkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNetworkServiceServer struct{}

func (UnimplementedNetworkServiceServer) GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIVersion not implemented")
}
func (UnimplementedNetworkServiceServer) IsPrivateEndpointServiceFeatureAvailable(context.Context, *IsPrivateEndpointServiceFeatureAvailableRequest) (*IsPrivateEndpointServiceFeatureAvailableResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPrivateEndpointServiceFeatureAvailable not implemented")
}
func (UnimplementedNetworkServiceServer) GetPrivateEndpointService(context.Context, *v1.IDOptions) (*PrivateEndpointService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateEndpointService not implemented")
}
func (UnimplementedNetworkServiceServer) GetPrivateEndpointServiceByDeploymentID(context.Context, *v1.IDOptions) (*PrivateEndpointService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateEndpointServiceByDeploymentID not implemented")
}
func (UnimplementedNetworkServiceServer) CreatePrivateEndpointService(context.Context, *PrivateEndpointService) (*PrivateEndpointService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrivateEndpointService not implemented")
}
func (UnimplementedNetworkServiceServer) UpdatePrivateEndpointService(context.Context, *PrivateEndpointService) (*v1.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrivateEndpointService not implemented")
}
func (UnimplementedNetworkServiceServer) mustEmbedUnimplementedNetworkServiceServer() {}
func (UnimplementedNetworkServiceServer) testEmbeddedByValue()                        {}

// UnsafeNetworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkServiceServer will
// result in compilation errors.
type UnsafeNetworkServiceServer interface {
	mustEmbedUnimplementedNetworkServiceServer()
}

func RegisterNetworkServiceServer(s grpc.ServiceRegistrar, srv NetworkServiceServer) {
	// If the following call pancis, it indicates UnimplementedNetworkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NetworkService_ServiceDesc, srv)
}

func _NetworkService_GetAPIVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetAPIVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_GetAPIVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetAPIVersion(ctx, req.(*v1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_IsPrivateEndpointServiceFeatureAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsPrivateEndpointServiceFeatureAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).IsPrivateEndpointServiceFeatureAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_IsPrivateEndpointServiceFeatureAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).IsPrivateEndpointServiceFeatureAvailable(ctx, req.(*IsPrivateEndpointServiceFeatureAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_GetPrivateEndpointService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetPrivateEndpointService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_GetPrivateEndpointService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetPrivateEndpointService(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_GetPrivateEndpointServiceByDeploymentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetPrivateEndpointServiceByDeploymentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_GetPrivateEndpointServiceByDeploymentID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetPrivateEndpointServiceByDeploymentID(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_CreatePrivateEndpointService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateEndpointService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).CreatePrivateEndpointService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_CreatePrivateEndpointService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).CreatePrivateEndpointService(ctx, req.(*PrivateEndpointService))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_UpdatePrivateEndpointService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateEndpointService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).UpdatePrivateEndpointService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_UpdatePrivateEndpointService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).UpdatePrivateEndpointService(ctx, req.(*PrivateEndpointService))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkService_ServiceDesc is the grpc.ServiceDesc for NetworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.network.v1.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAPIVersion",
			Handler:    _NetworkService_GetAPIVersion_Handler,
		},
		{
			MethodName: "IsPrivateEndpointServiceFeatureAvailable",
			Handler:    _NetworkService_IsPrivateEndpointServiceFeatureAvailable_Handler,
		},
		{
			MethodName: "GetPrivateEndpointService",
			Handler:    _NetworkService_GetPrivateEndpointService_Handler,
		},
		{
			MethodName: "GetPrivateEndpointServiceByDeploymentID",
			Handler:    _NetworkService_GetPrivateEndpointServiceByDeploymentID_Handler,
		},
		{
			MethodName: "CreatePrivateEndpointService",
			Handler:    _NetworkService_CreatePrivateEndpointService_Handler,
		},
		{
			MethodName: "UpdatePrivateEndpointService",
			Handler:    _NetworkService_UpdatePrivateEndpointService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network.proto",
}
