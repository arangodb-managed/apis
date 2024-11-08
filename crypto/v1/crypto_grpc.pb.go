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
// Author Ewout Prangsma
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: crypto.proto

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
	CryptoService_GetAPIVersion_FullMethodName                = "/arangodb.cloud.crypto.v1.CryptoService/GetAPIVersion"
	CryptoService_ListCACertificates_FullMethodName           = "/arangodb.cloud.crypto.v1.CryptoService/ListCACertificates"
	CryptoService_ListCACertificatesWithFilter_FullMethodName = "/arangodb.cloud.crypto.v1.CryptoService/ListCACertificatesWithFilter"
	CryptoService_GetCACertificate_FullMethodName             = "/arangodb.cloud.crypto.v1.CryptoService/GetCACertificate"
	CryptoService_GetCACertificateInstructions_FullMethodName = "/arangodb.cloud.crypto.v1.CryptoService/GetCACertificateInstructions"
	CryptoService_CreateCACertificate_FullMethodName          = "/arangodb.cloud.crypto.v1.CryptoService/CreateCACertificate"
	CryptoService_CloneCACertificate_FullMethodName           = "/arangodb.cloud.crypto.v1.CryptoService/CloneCACertificate"
	CryptoService_UpdateCACertificate_FullMethodName          = "/arangodb.cloud.crypto.v1.CryptoService/UpdateCACertificate"
	CryptoService_DeleteCACertificate_FullMethodName          = "/arangodb.cloud.crypto.v1.CryptoService/DeleteCACertificate"
	CryptoService_SetDefaultCACertificate_FullMethodName      = "/arangodb.cloud.crypto.v1.CryptoService/SetDefaultCACertificate"
)

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// CryptoService is the API used to configure various crypto objects.
type CryptoServiceClient interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error)
	// Fetch all CA certificates in the project identified by the given context ID.
	// Required permissions:
	// - crypto.cacertificate.list on the project identified by the given context ID
	ListCACertificates(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*CACertificateList, error)
	// Fetch all CA certificates in the project identified by the given project ID
	// that match the given filter.
	// Required permissions:
	// - crypto.cacertificate.list on the project identified by the given context ID
	ListCACertificatesWithFilter(ctx context.Context, in *ListCACertificatesRequest, opts ...grpc.CallOption) (*CACertificateList, error)
	// Fetch a CA certificate by its id.
	// Required permissions:
	// - crypto.cacertificate.get on the CA certificate identified by the given ID
	GetCACertificate(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*CACertificate, error)
	// Fetch instructions for installing & unistalling a CA certificate identified by its id
	// on various platforms.
	// Required permissions:
	// - crypto.cacertificate.get on the CA certificate identified by the given ID
	GetCACertificateInstructions(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*CACertificateInstructions, error)
	// Create a new CA certificate
	// Required permissions:
	// - crypto.cacertificate.create on the project that owns the CA certificate
	CreateCACertificate(ctx context.Context, in *CACertificate, opts ...grpc.CallOption) (*CACertificate, error)
	// Clone a CA certificate identified by given id.
	// Required permissions:
	// - crypto.cacertificate.clone on the project that owns the CA certificate identified by the given ID
	CloneCACertificate(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*CACertificate, error)
	// Update a CA certificate
	// Required permissions:
	// - crypto.cacertificate.update on the CA certificate
	UpdateCACertificate(ctx context.Context, in *CACertificate, opts ...grpc.CallOption) (*CACertificate, error)
	// Delete a CA certificate
	// Note that CA certificate are initially only marked for deleted.
	// Once all the resources that depend on it are removed the CA certificate itself is deleted
	// and cannot be restored.
	// Required permissions:
	// - crypto.cacertificate.delete on the CA certificate
	DeleteCACertificate(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v1.Empty, error)
	// Mark the given CA certificate as default for its containing project.
	// Required permissions:
	// - crypto.cacertificate.set-default on the project that owns the certificate.
	SetDefaultCACertificate(ctx context.Context, in *CACertificate, opts ...grpc.CallOption) (*v1.Empty, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Version)
	err := c.cc.Invoke(ctx, CryptoService_GetAPIVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) ListCACertificates(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*CACertificateList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CACertificateList)
	err := c.cc.Invoke(ctx, CryptoService_ListCACertificates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) ListCACertificatesWithFilter(ctx context.Context, in *ListCACertificatesRequest, opts ...grpc.CallOption) (*CACertificateList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CACertificateList)
	err := c.cc.Invoke(ctx, CryptoService_ListCACertificatesWithFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetCACertificate(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*CACertificate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CACertificate)
	err := c.cc.Invoke(ctx, CryptoService_GetCACertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetCACertificateInstructions(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*CACertificateInstructions, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CACertificateInstructions)
	err := c.cc.Invoke(ctx, CryptoService_GetCACertificateInstructions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CreateCACertificate(ctx context.Context, in *CACertificate, opts ...grpc.CallOption) (*CACertificate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CACertificate)
	err := c.cc.Invoke(ctx, CryptoService_CreateCACertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CloneCACertificate(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*CACertificate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CACertificate)
	err := c.cc.Invoke(ctx, CryptoService_CloneCACertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) UpdateCACertificate(ctx context.Context, in *CACertificate, opts ...grpc.CallOption) (*CACertificate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CACertificate)
	err := c.cc.Invoke(ctx, CryptoService_UpdateCACertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DeleteCACertificate(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v1.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Empty)
	err := c.cc.Invoke(ctx, CryptoService_DeleteCACertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) SetDefaultCACertificate(ctx context.Context, in *CACertificate, opts ...grpc.CallOption) (*v1.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Empty)
	err := c.cc.Invoke(ctx, CryptoService_SetDefaultCACertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
// All implementations must embed UnimplementedCryptoServiceServer
// for forward compatibility.
//
// CryptoService is the API used to configure various crypto objects.
type CryptoServiceServer interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error)
	// Fetch all CA certificates in the project identified by the given context ID.
	// Required permissions:
	// - crypto.cacertificate.list on the project identified by the given context ID
	ListCACertificates(context.Context, *v1.ListOptions) (*CACertificateList, error)
	// Fetch all CA certificates in the project identified by the given project ID
	// that match the given filter.
	// Required permissions:
	// - crypto.cacertificate.list on the project identified by the given context ID
	ListCACertificatesWithFilter(context.Context, *ListCACertificatesRequest) (*CACertificateList, error)
	// Fetch a CA certificate by its id.
	// Required permissions:
	// - crypto.cacertificate.get on the CA certificate identified by the given ID
	GetCACertificate(context.Context, *v1.IDOptions) (*CACertificate, error)
	// Fetch instructions for installing & unistalling a CA certificate identified by its id
	// on various platforms.
	// Required permissions:
	// - crypto.cacertificate.get on the CA certificate identified by the given ID
	GetCACertificateInstructions(context.Context, *v1.IDOptions) (*CACertificateInstructions, error)
	// Create a new CA certificate
	// Required permissions:
	// - crypto.cacertificate.create on the project that owns the CA certificate
	CreateCACertificate(context.Context, *CACertificate) (*CACertificate, error)
	// Clone a CA certificate identified by given id.
	// Required permissions:
	// - crypto.cacertificate.clone on the project that owns the CA certificate identified by the given ID
	CloneCACertificate(context.Context, *v1.IDOptions) (*CACertificate, error)
	// Update a CA certificate
	// Required permissions:
	// - crypto.cacertificate.update on the CA certificate
	UpdateCACertificate(context.Context, *CACertificate) (*CACertificate, error)
	// Delete a CA certificate
	// Note that CA certificate are initially only marked for deleted.
	// Once all the resources that depend on it are removed the CA certificate itself is deleted
	// and cannot be restored.
	// Required permissions:
	// - crypto.cacertificate.delete on the CA certificate
	DeleteCACertificate(context.Context, *v1.IDOptions) (*v1.Empty, error)
	// Mark the given CA certificate as default for its containing project.
	// Required permissions:
	// - crypto.cacertificate.set-default on the project that owns the certificate.
	SetDefaultCACertificate(context.Context, *CACertificate) (*v1.Empty, error)
	mustEmbedUnimplementedCryptoServiceServer()
}

// UnimplementedCryptoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCryptoServiceServer struct{}

func (UnimplementedCryptoServiceServer) GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIVersion not implemented")
}
func (UnimplementedCryptoServiceServer) ListCACertificates(context.Context, *v1.ListOptions) (*CACertificateList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCACertificates not implemented")
}
func (UnimplementedCryptoServiceServer) ListCACertificatesWithFilter(context.Context, *ListCACertificatesRequest) (*CACertificateList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCACertificatesWithFilter not implemented")
}
func (UnimplementedCryptoServiceServer) GetCACertificate(context.Context, *v1.IDOptions) (*CACertificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCACertificate not implemented")
}
func (UnimplementedCryptoServiceServer) GetCACertificateInstructions(context.Context, *v1.IDOptions) (*CACertificateInstructions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCACertificateInstructions not implemented")
}
func (UnimplementedCryptoServiceServer) CreateCACertificate(context.Context, *CACertificate) (*CACertificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCACertificate not implemented")
}
func (UnimplementedCryptoServiceServer) CloneCACertificate(context.Context, *v1.IDOptions) (*CACertificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneCACertificate not implemented")
}
func (UnimplementedCryptoServiceServer) UpdateCACertificate(context.Context, *CACertificate) (*CACertificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCACertificate not implemented")
}
func (UnimplementedCryptoServiceServer) DeleteCACertificate(context.Context, *v1.IDOptions) (*v1.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCACertificate not implemented")
}
func (UnimplementedCryptoServiceServer) SetDefaultCACertificate(context.Context, *CACertificate) (*v1.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultCACertificate not implemented")
}
func (UnimplementedCryptoServiceServer) mustEmbedUnimplementedCryptoServiceServer() {}
func (UnimplementedCryptoServiceServer) testEmbeddedByValue()                       {}

// UnsafeCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoServiceServer will
// result in compilation errors.
type UnsafeCryptoServiceServer interface {
	mustEmbedUnimplementedCryptoServiceServer()
}

func RegisterCryptoServiceServer(s grpc.ServiceRegistrar, srv CryptoServiceServer) {
	// If the following call pancis, it indicates UnimplementedCryptoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CryptoService_ServiceDesc, srv)
}

func _CryptoService_GetAPIVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetAPIVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_GetAPIVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetAPIVersion(ctx, req.(*v1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_ListCACertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).ListCACertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_ListCACertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).ListCACertificates(ctx, req.(*v1.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_ListCACertificatesWithFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCACertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).ListCACertificatesWithFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_ListCACertificatesWithFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).ListCACertificatesWithFilter(ctx, req.(*ListCACertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetCACertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetCACertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_GetCACertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetCACertificate(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetCACertificateInstructions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetCACertificateInstructions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_GetCACertificateInstructions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetCACertificateInstructions(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CreateCACertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CACertificate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CreateCACertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_CreateCACertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CreateCACertificate(ctx, req.(*CACertificate))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CloneCACertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CloneCACertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_CloneCACertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CloneCACertificate(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_UpdateCACertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CACertificate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).UpdateCACertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_UpdateCACertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).UpdateCACertificate(ctx, req.(*CACertificate))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DeleteCACertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DeleteCACertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_DeleteCACertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DeleteCACertificate(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_SetDefaultCACertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CACertificate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).SetDefaultCACertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_SetDefaultCACertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).SetDefaultCACertificate(ctx, req.(*CACertificate))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptoService_ServiceDesc is the grpc.ServiceDesc for CryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.crypto.v1.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAPIVersion",
			Handler:    _CryptoService_GetAPIVersion_Handler,
		},
		{
			MethodName: "ListCACertificates",
			Handler:    _CryptoService_ListCACertificates_Handler,
		},
		{
			MethodName: "ListCACertificatesWithFilter",
			Handler:    _CryptoService_ListCACertificatesWithFilter_Handler,
		},
		{
			MethodName: "GetCACertificate",
			Handler:    _CryptoService_GetCACertificate_Handler,
		},
		{
			MethodName: "GetCACertificateInstructions",
			Handler:    _CryptoService_GetCACertificateInstructions_Handler,
		},
		{
			MethodName: "CreateCACertificate",
			Handler:    _CryptoService_CreateCACertificate_Handler,
		},
		{
			MethodName: "CloneCACertificate",
			Handler:    _CryptoService_CloneCACertificate_Handler,
		},
		{
			MethodName: "UpdateCACertificate",
			Handler:    _CryptoService_UpdateCACertificate_Handler,
		},
		{
			MethodName: "DeleteCACertificate",
			Handler:    _CryptoService_DeleteCACertificate_Handler,
		},
		{
			MethodName: "SetDefaultCACertificate",
			Handler:    _CryptoService_SetDefaultCACertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto.proto",
}
