//
// DISCLAIMER
//
// Copyright 2020-2024 ArangoDB GmbH, Cologne, Germany
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
// source: replication.proto

package v1

import (
	context "context"
	v1 "github.com/arangodb-managed/apis/common/v1"
	v11 "github.com/arangodb-managed/apis/data/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ReplicationService_GetAPIVersion_FullMethodName               = "/arangodb.cloud.replication.v1.ReplicationService/GetAPIVersion"
	ReplicationService_CloneDeploymentFromBackup_FullMethodName   = "/arangodb.cloud.replication.v1.ReplicationService/CloneDeploymentFromBackup"
	ReplicationService_GetDeploymentReplication_FullMethodName    = "/arangodb.cloud.replication.v1.ReplicationService/GetDeploymentReplication"
	ReplicationService_UpdateDeploymentReplication_FullMethodName = "/arangodb.cloud.replication.v1.ReplicationService/UpdateDeploymentReplication"
	ReplicationService_CreateDeploymentMigration_FullMethodName   = "/arangodb.cloud.replication.v1.ReplicationService/CreateDeploymentMigration"
	ReplicationService_GetDeploymentMigration_FullMethodName      = "/arangodb.cloud.replication.v1.ReplicationService/GetDeploymentMigration"
	ReplicationService_DeleteDeploymentMigration_FullMethodName   = "/arangodb.cloud.replication.v1.ReplicationService/DeleteDeploymentMigration"
)

// ReplicationServiceClient is the client API for ReplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ReplicationService is the API used to replicate a deployment.
type ReplicationServiceClient interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error)
	// Takes a backup and creates a deployment from it. For all intents and purposes this new deployment
	// will be the same as the deployment at that exact moment when the backup was taken from it. This means that
	// the new deployment will be in the same project and use the same provider as the old deployment did. Optionally
	// a different region can be provided using the region id field on the request. Furthermore, the new deployment
	// will have the same server settings ( count, mode, replication factor ) as the old deployment did at the time
	// of taking the backup. After the new deployment successfully started, the backup will be used to restore the
	// data into the new deployment. The new deployment will have a different endpoint, and the password will also
	// be reset for it. All other user settings will remain the same.
	// The old deployment will not be touched.
	// Required permissions:
	// if project_id is specified
	// - backup.backup.get on the backup specified by backup_id in request
	// - replication.deployment.clone-from-backup on the project specified in request
	// if project_id is not specified
	// - replication.deployment.clone-from-backup on the backup specified by backup_id
	CloneDeploymentFromBackup(ctx context.Context, in *CloneDeploymentFromBackupRequest, opts ...grpc.CallOption) (*v11.Deployment, error)
	// Get an existing DeploymentReplication using its deployment ID
	// Required permissions:
	// - replication.deploymentreplication.get
	// [Deprecated] This method shouldn't be used anymore, the permission is removed from the system already to prevent usage.
	GetDeploymentReplication(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*DeploymentReplication, error)
	// Update an existing DeploymentReplication spec. If does not exist, this will create a new one.
	// This call expects the complete entity with the updated fields.
	// Required permissions:
	// - replication.deploymentreplication.update
	// [Deprecated] This method shouldn't be used anymore, the permission is removed from the system already to prevent usage.
	UpdateDeploymentReplication(ctx context.Context, in *DeploymentReplication, opts ...grpc.CallOption) (*DeploymentReplication, error)
	// Create a new deployment migration.
	// Note: currently migration is supported only for Deployments with 'free' model.
	// Required permissions:
	// - replication.deploymentmigration.create on the specified deployment ID
	CreateDeploymentMigration(ctx context.Context, in *DeploymentMigration, opts ...grpc.CallOption) (*DeploymentMigration, error)
	// Get info about the deployment migration for a deployment identified by the given ID.
	// Required permissions:
	// - replication.deploymentmigration.get on the specified deployment ID
	GetDeploymentMigration(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*DeploymentMigration, error)
	// Delete an existing DeploymentMigration.
	// A DeploymentMigration may be deleted only if it is in COMPLETE or FAILED state.
	// Required permissions:
	// - replication.deploymentmigration.delete on the specified deployment ID.
	DeleteDeploymentMigration(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v1.Empty, error)
}

type replicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationServiceClient(cc grpc.ClientConnInterface) ReplicationServiceClient {
	return &replicationServiceClient{cc}
}

func (c *replicationServiceClient) GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Version)
	err := c.cc.Invoke(ctx, ReplicationService_GetAPIVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) CloneDeploymentFromBackup(ctx context.Context, in *CloneDeploymentFromBackupRequest, opts ...grpc.CallOption) (*v11.Deployment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.Deployment)
	err := c.cc.Invoke(ctx, ReplicationService_CloneDeploymentFromBackup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) GetDeploymentReplication(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*DeploymentReplication, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeploymentReplication)
	err := c.cc.Invoke(ctx, ReplicationService_GetDeploymentReplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) UpdateDeploymentReplication(ctx context.Context, in *DeploymentReplication, opts ...grpc.CallOption) (*DeploymentReplication, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeploymentReplication)
	err := c.cc.Invoke(ctx, ReplicationService_UpdateDeploymentReplication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) CreateDeploymentMigration(ctx context.Context, in *DeploymentMigration, opts ...grpc.CallOption) (*DeploymentMigration, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeploymentMigration)
	err := c.cc.Invoke(ctx, ReplicationService_CreateDeploymentMigration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) GetDeploymentMigration(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*DeploymentMigration, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeploymentMigration)
	err := c.cc.Invoke(ctx, ReplicationService_GetDeploymentMigration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) DeleteDeploymentMigration(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v1.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.Empty)
	err := c.cc.Invoke(ctx, ReplicationService_DeleteDeploymentMigration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServiceServer is the server API for ReplicationService service.
// All implementations must embed UnimplementedReplicationServiceServer
// for forward compatibility.
//
// ReplicationService is the API used to replicate a deployment.
type ReplicationServiceServer interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error)
	// Takes a backup and creates a deployment from it. For all intents and purposes this new deployment
	// will be the same as the deployment at that exact moment when the backup was taken from it. This means that
	// the new deployment will be in the same project and use the same provider as the old deployment did. Optionally
	// a different region can be provided using the region id field on the request. Furthermore, the new deployment
	// will have the same server settings ( count, mode, replication factor ) as the old deployment did at the time
	// of taking the backup. After the new deployment successfully started, the backup will be used to restore the
	// data into the new deployment. The new deployment will have a different endpoint, and the password will also
	// be reset for it. All other user settings will remain the same.
	// The old deployment will not be touched.
	// Required permissions:
	// if project_id is specified
	// - backup.backup.get on the backup specified by backup_id in request
	// - replication.deployment.clone-from-backup on the project specified in request
	// if project_id is not specified
	// - replication.deployment.clone-from-backup on the backup specified by backup_id
	CloneDeploymentFromBackup(context.Context, *CloneDeploymentFromBackupRequest) (*v11.Deployment, error)
	// Get an existing DeploymentReplication using its deployment ID
	// Required permissions:
	// - replication.deploymentreplication.get
	// [Deprecated] This method shouldn't be used anymore, the permission is removed from the system already to prevent usage.
	GetDeploymentReplication(context.Context, *v1.IDOptions) (*DeploymentReplication, error)
	// Update an existing DeploymentReplication spec. If does not exist, this will create a new one.
	// This call expects the complete entity with the updated fields.
	// Required permissions:
	// - replication.deploymentreplication.update
	// [Deprecated] This method shouldn't be used anymore, the permission is removed from the system already to prevent usage.
	UpdateDeploymentReplication(context.Context, *DeploymentReplication) (*DeploymentReplication, error)
	// Create a new deployment migration.
	// Note: currently migration is supported only for Deployments with 'free' model.
	// Required permissions:
	// - replication.deploymentmigration.create on the specified deployment ID
	CreateDeploymentMigration(context.Context, *DeploymentMigration) (*DeploymentMigration, error)
	// Get info about the deployment migration for a deployment identified by the given ID.
	// Required permissions:
	// - replication.deploymentmigration.get on the specified deployment ID
	GetDeploymentMigration(context.Context, *v1.IDOptions) (*DeploymentMigration, error)
	// Delete an existing DeploymentMigration.
	// A DeploymentMigration may be deleted only if it is in COMPLETE or FAILED state.
	// Required permissions:
	// - replication.deploymentmigration.delete on the specified deployment ID.
	DeleteDeploymentMigration(context.Context, *v1.IDOptions) (*v1.Empty, error)
	mustEmbedUnimplementedReplicationServiceServer()
}

// UnimplementedReplicationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReplicationServiceServer struct{}

func (UnimplementedReplicationServiceServer) GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIVersion not implemented")
}
func (UnimplementedReplicationServiceServer) CloneDeploymentFromBackup(context.Context, *CloneDeploymentFromBackupRequest) (*v11.Deployment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneDeploymentFromBackup not implemented")
}
func (UnimplementedReplicationServiceServer) GetDeploymentReplication(context.Context, *v1.IDOptions) (*DeploymentReplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeploymentReplication not implemented")
}
func (UnimplementedReplicationServiceServer) UpdateDeploymentReplication(context.Context, *DeploymentReplication) (*DeploymentReplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeploymentReplication not implemented")
}
func (UnimplementedReplicationServiceServer) CreateDeploymentMigration(context.Context, *DeploymentMigration) (*DeploymentMigration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeploymentMigration not implemented")
}
func (UnimplementedReplicationServiceServer) GetDeploymentMigration(context.Context, *v1.IDOptions) (*DeploymentMigration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeploymentMigration not implemented")
}
func (UnimplementedReplicationServiceServer) DeleteDeploymentMigration(context.Context, *v1.IDOptions) (*v1.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeploymentMigration not implemented")
}
func (UnimplementedReplicationServiceServer) mustEmbedUnimplementedReplicationServiceServer() {}
func (UnimplementedReplicationServiceServer) testEmbeddedByValue()                            {}

// UnsafeReplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationServiceServer will
// result in compilation errors.
type UnsafeReplicationServiceServer interface {
	mustEmbedUnimplementedReplicationServiceServer()
}

func RegisterReplicationServiceServer(s grpc.ServiceRegistrar, srv ReplicationServiceServer) {
	// If the following call pancis, it indicates UnimplementedReplicationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReplicationService_ServiceDesc, srv)
}

func _ReplicationService_GetAPIVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).GetAPIVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_GetAPIVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).GetAPIVersion(ctx, req.(*v1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_CloneDeploymentFromBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneDeploymentFromBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).CloneDeploymentFromBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_CloneDeploymentFromBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).CloneDeploymentFromBackup(ctx, req.(*CloneDeploymentFromBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_GetDeploymentReplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).GetDeploymentReplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_GetDeploymentReplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).GetDeploymentReplication(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_UpdateDeploymentReplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentReplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).UpdateDeploymentReplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_UpdateDeploymentReplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).UpdateDeploymentReplication(ctx, req.(*DeploymentReplication))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_CreateDeploymentMigration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentMigration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).CreateDeploymentMigration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_CreateDeploymentMigration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).CreateDeploymentMigration(ctx, req.(*DeploymentMigration))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_GetDeploymentMigration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).GetDeploymentMigration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_GetDeploymentMigration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).GetDeploymentMigration(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_DeleteDeploymentMigration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).DeleteDeploymentMigration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_DeleteDeploymentMigration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).DeleteDeploymentMigration(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// ReplicationService_ServiceDesc is the grpc.ServiceDesc for ReplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.replication.v1.ReplicationService",
	HandlerType: (*ReplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAPIVersion",
			Handler:    _ReplicationService_GetAPIVersion_Handler,
		},
		{
			MethodName: "CloneDeploymentFromBackup",
			Handler:    _ReplicationService_CloneDeploymentFromBackup_Handler,
		},
		{
			MethodName: "GetDeploymentReplication",
			Handler:    _ReplicationService_GetDeploymentReplication_Handler,
		},
		{
			MethodName: "UpdateDeploymentReplication",
			Handler:    _ReplicationService_UpdateDeploymentReplication_Handler,
		},
		{
			MethodName: "CreateDeploymentMigration",
			Handler:    _ReplicationService_CreateDeploymentMigration_Handler,
		},
		{
			MethodName: "GetDeploymentMigration",
			Handler:    _ReplicationService_GetDeploymentMigration_Handler,
		},
		{
			MethodName: "DeleteDeploymentMigration",
			Handler:    _ReplicationService_DeleteDeploymentMigration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "replication.proto",
}
