//
// DISCLAIMER
//
// Copyright 2020-2023 ArangoDB GmbH, Cologne, Germany
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
// source: platform.proto

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
	PlatformService_GetAPIVersion_FullMethodName = "/arangodb.cloud.platform.v1.PlatformService/GetAPIVersion"
	PlatformService_ListProviders_FullMethodName = "/arangodb.cloud.platform.v1.PlatformService/ListProviders"
	PlatformService_GetProvider_FullMethodName   = "/arangodb.cloud.platform.v1.PlatformService/GetProvider"
	PlatformService_ListRegions_FullMethodName   = "/arangodb.cloud.platform.v1.PlatformService/ListRegions"
	PlatformService_GetRegion_FullMethodName     = "/arangodb.cloud.platform.v1.PlatformService/GetRegion"
)

// PlatformServiceClient is the client API for PlatformService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// PlatformService is the API used to query for cloud provider & regional info.
type PlatformServiceClient interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error)
	// Fetch all providers that are supported by the ArangoDB cloud.
	// Required permissions:
	// - None
	ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ProviderList, error)
	// Fetch a provider by its id.
	// Required permissions:
	// - None
	GetProvider(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Provider, error)
	// Fetch all regions provided by the provided identified by the given context ID.
	// If the given context identifier contains a valid organization ID,
	// the result includes all regions for that organization.
	// Otherwise only regions are returned that are available to all organizations.
	// Required permissions:
	// - None
	ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*RegionList, error)
	// Fetch a region by its id.
	// Required permissions:
	// - None
	GetRegion(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Region, error)
}

type platformServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformServiceClient(cc grpc.ClientConnInterface) PlatformServiceClient {
	return &platformServiceClient{cc}
}

func (c *platformServiceClient) GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Version)
	err := c.cc.Invoke(ctx, PlatformService_GetAPIVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ProviderList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProviderList)
	err := c.cc.Invoke(ctx, PlatformService_ListProviders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) GetProvider(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Provider, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Provider)
	err := c.cc.Invoke(ctx, PlatformService_GetProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*RegionList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegionList)
	err := c.cc.Invoke(ctx, PlatformService_ListRegions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) GetRegion(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*Region, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Region)
	err := c.cc.Invoke(ctx, PlatformService_GetRegion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServiceServer is the server API for PlatformService service.
// All implementations must embed UnimplementedPlatformServiceServer
// for forward compatibility.
//
// PlatformService is the API used to query for cloud provider & regional info.
type PlatformServiceServer interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error)
	// Fetch all providers that are supported by the ArangoDB cloud.
	// Required permissions:
	// - None
	ListProviders(context.Context, *ListProvidersRequest) (*ProviderList, error)
	// Fetch a provider by its id.
	// Required permissions:
	// - None
	GetProvider(context.Context, *v1.IDOptions) (*Provider, error)
	// Fetch all regions provided by the provided identified by the given context ID.
	// If the given context identifier contains a valid organization ID,
	// the result includes all regions for that organization.
	// Otherwise only regions are returned that are available to all organizations.
	// Required permissions:
	// - None
	ListRegions(context.Context, *ListRegionsRequest) (*RegionList, error)
	// Fetch a region by its id.
	// Required permissions:
	// - None
	GetRegion(context.Context, *v1.IDOptions) (*Region, error)
	mustEmbedUnimplementedPlatformServiceServer()
}

// UnimplementedPlatformServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPlatformServiceServer struct{}

func (UnimplementedPlatformServiceServer) GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIVersion not implemented")
}
func (UnimplementedPlatformServiceServer) ListProviders(context.Context, *ListProvidersRequest) (*ProviderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (UnimplementedPlatformServiceServer) GetProvider(context.Context, *v1.IDOptions) (*Provider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProvider not implemented")
}
func (UnimplementedPlatformServiceServer) ListRegions(context.Context, *ListRegionsRequest) (*RegionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegions not implemented")
}
func (UnimplementedPlatformServiceServer) GetRegion(context.Context, *v1.IDOptions) (*Region, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegion not implemented")
}
func (UnimplementedPlatformServiceServer) mustEmbedUnimplementedPlatformServiceServer() {}
func (UnimplementedPlatformServiceServer) testEmbeddedByValue()                         {}

// UnsafePlatformServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlatformServiceServer will
// result in compilation errors.
type UnsafePlatformServiceServer interface {
	mustEmbedUnimplementedPlatformServiceServer()
}

func RegisterPlatformServiceServer(s grpc.ServiceRegistrar, srv PlatformServiceServer) {
	// If the following call pancis, it indicates UnimplementedPlatformServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PlatformService_ServiceDesc, srv)
}

func _PlatformService_GetAPIVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).GetAPIVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_GetAPIVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).GetAPIVersion(ctx, req.(*v1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_ListProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).ListProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_ListProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).ListProviders(ctx, req.(*ListProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_GetProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).GetProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_GetProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).GetProvider(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_ListRegions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).ListRegions(ctx, req.(*ListRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_GetRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).GetRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_GetRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).GetRegion(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// PlatformService_ServiceDesc is the grpc.ServiceDesc for PlatformService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlatformService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.platform.v1.PlatformService",
	HandlerType: (*PlatformServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAPIVersion",
			Handler:    _PlatformService_GetAPIVersion_Handler,
		},
		{
			MethodName: "ListProviders",
			Handler:    _PlatformService_ListProviders_Handler,
		},
		{
			MethodName: "GetProvider",
			Handler:    _PlatformService_GetProvider_Handler,
		},
		{
			MethodName: "ListRegions",
			Handler:    _PlatformService_ListRegions_Handler,
		},
		{
			MethodName: "GetRegion",
			Handler:    _PlatformService_GetRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform.proto",
}
