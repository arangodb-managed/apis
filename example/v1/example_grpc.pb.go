//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
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
// Author Gergely Brautigam
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: example.proto

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
	ExampleDatasetService_GetAPIVersion_FullMethodName                    = "/arangodb.cloud.example.v1.ExampleDatasetService/GetAPIVersion"
	ExampleDatasetService_ListExampleDatasets_FullMethodName              = "/arangodb.cloud.example.v1.ExampleDatasetService/ListExampleDatasets"
	ExampleDatasetService_GetExampleDataset_FullMethodName                = "/arangodb.cloud.example.v1.ExampleDatasetService/GetExampleDataset"
	ExampleDatasetService_ListExampleDatasetInstallations_FullMethodName  = "/arangodb.cloud.example.v1.ExampleDatasetService/ListExampleDatasetInstallations"
	ExampleDatasetService_GetExampleDatasetInstallation_FullMethodName    = "/arangodb.cloud.example.v1.ExampleDatasetService/GetExampleDatasetInstallation"
	ExampleDatasetService_CreateExampleDatasetInstallation_FullMethodName = "/arangodb.cloud.example.v1.ExampleDatasetService/CreateExampleDatasetInstallation"
	ExampleDatasetService_UpdateExampleDatasetInstallation_FullMethodName = "/arangodb.cloud.example.v1.ExampleDatasetService/UpdateExampleDatasetInstallation"
	ExampleDatasetService_DeleteExampleDatasetInstallation_FullMethodName = "/arangodb.cloud.example.v1.ExampleDatasetService/DeleteExampleDatasetInstallation"
)

// ExampleDatasetServiceClient is the client API for ExampleDatasetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ExampleDatasetService is the API used to query for example datasets.
type ExampleDatasetServiceClient interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error)
	// Fetch all example datasets.
	// Required permissions:
	// - None: Caller must be authenticated.
	ListExampleDatasets(ctx context.Context, in *ListExampleDatasetsRequest, opts ...grpc.CallOption) (*ExampleDatasetList, error)
	// Fetch an example dataset identified by the given ID.
	// Required permissions:
	// - None: Caller must be authenticated.
	GetExampleDataset(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*ExampleDataset, error)
	// Fetch all installations for a specific deployment.
	// Required permissions:
	// - example.exampledatasetinstallation.list on the deployment that owns the installation and is identified by the given ID.
	ListExampleDatasetInstallations(ctx context.Context, in *ListExampleDatasetInstallationsRequest, opts ...grpc.CallOption) (*ExampleDatasetInstallationList, error)
	// Fetch an installations identified by the given ID.
	// Required permissions:
	// - example.exampledatasetinstallation.get on the installation identified by the given ID.
	GetExampleDatasetInstallation(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*ExampleDatasetInstallation, error)
	// Create a new example installation. This represents a request made by the user to create an installation
	// for a deployment given a specified example dataset.
	// Required permissions:
	// -  example.exampledatasetinstallation.create on the deployment that the installation is for and is identified by the given ID.
	CreateExampleDatasetInstallation(ctx context.Context, in *ExampleDatasetInstallation, opts ...grpc.CallOption) (*ExampleDatasetInstallation, error)
	// Update an installation.
	// Required permissions:
	// -  example.exampledatasetinstallation.update on the installation identified by the given ID.
	UpdateExampleDatasetInstallation(ctx context.Context, in *ExampleDatasetInstallation, opts ...grpc.CallOption) (*ExampleDatasetInstallation, error)
	// Delete an installation identified by the given ID.
	// Required permissions:
	// -  example.exampledatasetinstallation.delete on the installation identified by the given ID.
	DeleteExampleDatasetInstallation(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v1.Empty, error)
}

type exampleDatasetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleDatasetServiceClient(cc grpc.ClientConnInterface) ExampleDatasetServiceClient {
	return &exampleDatasetServiceClient{cc}
}

func (c *exampleDatasetServiceClient) GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Version)
	err := c.cc.Invoke(ctx, ExampleDatasetService_GetAPIVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleDatasetServiceClient) ListExampleDatasets(ctx context.Context, in *ListExampleDatasetsRequest, opts ...grpc.CallOption) (*ExampleDatasetList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExampleDatasetList)
	err := c.cc.Invoke(ctx, ExampleDatasetService_ListExampleDatasets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleDatasetServiceClient) GetExampleDataset(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*ExampleDataset, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExampleDataset)
	err := c.cc.Invoke(ctx, ExampleDatasetService_GetExampleDataset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleDatasetServiceClient) ListExampleDatasetInstallations(ctx context.Context, in *ListExampleDatasetInstallationsRequest, opts ...grpc.CallOption) (*ExampleDatasetInstallationList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExampleDatasetInstallationList)
	err := c.cc.Invoke(ctx, ExampleDatasetService_ListExampleDatasetInstallations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleDatasetServiceClient) GetExampleDatasetInstallation(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*ExampleDatasetInstallation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExampleDatasetInstallation)
	err := c.cc.Invoke(ctx, ExampleDatasetService_GetExampleDatasetInstallation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleDatasetServiceClient) CreateExampleDatasetInstallation(ctx context.Context, in *ExampleDatasetInstallation, opts ...grpc.CallOption) (*ExampleDatasetInstallation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExampleDatasetInstallation)
	err := c.cc.Invoke(ctx, ExampleDatasetService_CreateExampleDatasetInstallation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleDatasetServiceClient) UpdateExampleDatasetInstallation(ctx context.Context, in *ExampleDatasetInstallation, opts ...grpc.CallOption) (*ExampleDatasetInstallation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExampleDatasetInstallation)
	err := c.cc.Invoke(ctx, ExampleDatasetService_UpdateExampleDatasetInstallation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleDatasetServiceClient) DeleteExampleDatasetInstallation(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v1.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Empty)
	err := c.cc.Invoke(ctx, ExampleDatasetService_DeleteExampleDatasetInstallation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleDatasetServiceServer is the server API for ExampleDatasetService service.
// All implementations must embed UnimplementedExampleDatasetServiceServer
// for forward compatibility.
//
// ExampleDatasetService is the API used to query for example datasets.
type ExampleDatasetServiceServer interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error)
	// Fetch all example datasets.
	// Required permissions:
	// - None: Caller must be authenticated.
	ListExampleDatasets(context.Context, *ListExampleDatasetsRequest) (*ExampleDatasetList, error)
	// Fetch an example dataset identified by the given ID.
	// Required permissions:
	// - None: Caller must be authenticated.
	GetExampleDataset(context.Context, *v1.IDOptions) (*ExampleDataset, error)
	// Fetch all installations for a specific deployment.
	// Required permissions:
	// - example.exampledatasetinstallation.list on the deployment that owns the installation and is identified by the given ID.
	ListExampleDatasetInstallations(context.Context, *ListExampleDatasetInstallationsRequest) (*ExampleDatasetInstallationList, error)
	// Fetch an installations identified by the given ID.
	// Required permissions:
	// - example.exampledatasetinstallation.get on the installation identified by the given ID.
	GetExampleDatasetInstallation(context.Context, *v1.IDOptions) (*ExampleDatasetInstallation, error)
	// Create a new example installation. This represents a request made by the user to create an installation
	// for a deployment given a specified example dataset.
	// Required permissions:
	// -  example.exampledatasetinstallation.create on the deployment that the installation is for and is identified by the given ID.
	CreateExampleDatasetInstallation(context.Context, *ExampleDatasetInstallation) (*ExampleDatasetInstallation, error)
	// Update an installation.
	// Required permissions:
	// -  example.exampledatasetinstallation.update on the installation identified by the given ID.
	UpdateExampleDatasetInstallation(context.Context, *ExampleDatasetInstallation) (*ExampleDatasetInstallation, error)
	// Delete an installation identified by the given ID.
	// Required permissions:
	// -  example.exampledatasetinstallation.delete on the installation identified by the given ID.
	DeleteExampleDatasetInstallation(context.Context, *v1.IDOptions) (*v1.Empty, error)
	mustEmbedUnimplementedExampleDatasetServiceServer()
}

// UnimplementedExampleDatasetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExampleDatasetServiceServer struct{}

func (UnimplementedExampleDatasetServiceServer) GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIVersion not implemented")
}
func (UnimplementedExampleDatasetServiceServer) ListExampleDatasets(context.Context, *ListExampleDatasetsRequest) (*ExampleDatasetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExampleDatasets not implemented")
}
func (UnimplementedExampleDatasetServiceServer) GetExampleDataset(context.Context, *v1.IDOptions) (*ExampleDataset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExampleDataset not implemented")
}
func (UnimplementedExampleDatasetServiceServer) ListExampleDatasetInstallations(context.Context, *ListExampleDatasetInstallationsRequest) (*ExampleDatasetInstallationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExampleDatasetInstallations not implemented")
}
func (UnimplementedExampleDatasetServiceServer) GetExampleDatasetInstallation(context.Context, *v1.IDOptions) (*ExampleDatasetInstallation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExampleDatasetInstallation not implemented")
}
func (UnimplementedExampleDatasetServiceServer) CreateExampleDatasetInstallation(context.Context, *ExampleDatasetInstallation) (*ExampleDatasetInstallation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExampleDatasetInstallation not implemented")
}
func (UnimplementedExampleDatasetServiceServer) UpdateExampleDatasetInstallation(context.Context, *ExampleDatasetInstallation) (*ExampleDatasetInstallation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExampleDatasetInstallation not implemented")
}
func (UnimplementedExampleDatasetServiceServer) DeleteExampleDatasetInstallation(context.Context, *v1.IDOptions) (*v1.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExampleDatasetInstallation not implemented")
}
func (UnimplementedExampleDatasetServiceServer) mustEmbedUnimplementedExampleDatasetServiceServer() {}
func (UnimplementedExampleDatasetServiceServer) testEmbeddedByValue()                               {}

// UnsafeExampleDatasetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleDatasetServiceServer will
// result in compilation errors.
type UnsafeExampleDatasetServiceServer interface {
	mustEmbedUnimplementedExampleDatasetServiceServer()
}

func RegisterExampleDatasetServiceServer(s grpc.ServiceRegistrar, srv ExampleDatasetServiceServer) {
	// If the following call pancis, it indicates UnimplementedExampleDatasetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExampleDatasetService_ServiceDesc, srv)
}

func _ExampleDatasetService_GetAPIVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleDatasetServiceServer).GetAPIVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleDatasetService_GetAPIVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleDatasetServiceServer).GetAPIVersion(ctx, req.(*v1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleDatasetService_ListExampleDatasets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExampleDatasetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleDatasetServiceServer).ListExampleDatasets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleDatasetService_ListExampleDatasets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleDatasetServiceServer).ListExampleDatasets(ctx, req.(*ListExampleDatasetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleDatasetService_GetExampleDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleDatasetServiceServer).GetExampleDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleDatasetService_GetExampleDataset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleDatasetServiceServer).GetExampleDataset(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleDatasetService_ListExampleDatasetInstallations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExampleDatasetInstallationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleDatasetServiceServer).ListExampleDatasetInstallations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleDatasetService_ListExampleDatasetInstallations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleDatasetServiceServer).ListExampleDatasetInstallations(ctx, req.(*ListExampleDatasetInstallationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleDatasetService_GetExampleDatasetInstallation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleDatasetServiceServer).GetExampleDatasetInstallation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleDatasetService_GetExampleDatasetInstallation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleDatasetServiceServer).GetExampleDatasetInstallation(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleDatasetService_CreateExampleDatasetInstallation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleDatasetInstallation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleDatasetServiceServer).CreateExampleDatasetInstallation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleDatasetService_CreateExampleDatasetInstallation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleDatasetServiceServer).CreateExampleDatasetInstallation(ctx, req.(*ExampleDatasetInstallation))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleDatasetService_UpdateExampleDatasetInstallation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleDatasetInstallation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleDatasetServiceServer).UpdateExampleDatasetInstallation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleDatasetService_UpdateExampleDatasetInstallation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleDatasetServiceServer).UpdateExampleDatasetInstallation(ctx, req.(*ExampleDatasetInstallation))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleDatasetService_DeleteExampleDatasetInstallation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleDatasetServiceServer).DeleteExampleDatasetInstallation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleDatasetService_DeleteExampleDatasetInstallation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleDatasetServiceServer).DeleteExampleDatasetInstallation(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleDatasetService_ServiceDesc is the grpc.ServiceDesc for ExampleDatasetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleDatasetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.example.v1.ExampleDatasetService",
	HandlerType: (*ExampleDatasetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAPIVersion",
			Handler:    _ExampleDatasetService_GetAPIVersion_Handler,
		},
		{
			MethodName: "ListExampleDatasets",
			Handler:    _ExampleDatasetService_ListExampleDatasets_Handler,
		},
		{
			MethodName: "GetExampleDataset",
			Handler:    _ExampleDatasetService_GetExampleDataset_Handler,
		},
		{
			MethodName: "ListExampleDatasetInstallations",
			Handler:    _ExampleDatasetService_ListExampleDatasetInstallations_Handler,
		},
		{
			MethodName: "GetExampleDatasetInstallation",
			Handler:    _ExampleDatasetService_GetExampleDatasetInstallation_Handler,
		},
		{
			MethodName: "CreateExampleDatasetInstallation",
			Handler:    _ExampleDatasetService_CreateExampleDatasetInstallation_Handler,
		},
		{
			MethodName: "UpdateExampleDatasetInstallation",
			Handler:    _ExampleDatasetService_UpdateExampleDatasetInstallation_Handler,
		},
		{
			MethodName: "DeleteExampleDatasetInstallation",
			Handler:    _ExampleDatasetService_DeleteExampleDatasetInstallation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
