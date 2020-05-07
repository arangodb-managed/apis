// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: replication.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/arangodb-managed/apis/common/v1"
	v11 "github.com/arangodb-managed/apis/data/v1"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("replication.proto", fileDescriptor_ed0454e9e09fb71a) }

var fileDescriptor_ed0454e9e09fb71a = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xbf, 0xf4, 0xf0, 0x1d, 0x02, 0x1e, 0xdc, 0x83, 0x60, 0xd0, 0x94, 0x46, 0x50, 0x50,
	0x9a, 0x25, 0x7a, 0xf3, 0xd6, 0x5a, 0x95, 0x9e, 0x14, 0x05, 0x0f, 0xde, 0x26, 0xc9, 0x10, 0x17,
	0x93, 0x9d, 0x65, 0xbb, 0x0d, 0x14, 0xf5, 0xe2, 0x2b, 0x78, 0xd0, 0x87, 0xf0, 0x41, 0x3c, 0x0a,
	0xbe, 0x80, 0x54, 0x1f, 0x44, 0xb2, 0x8d, 0xb4, 0x04, 0x7b, 0xda, 0x65, 0x76, 0xf6, 0xf7, 0xff,
	0xcd, 0xb8, 0xab, 0x1a, 0x55, 0x2e, 0x12, 0x30, 0x82, 0x64, 0xa8, 0x34, 0x19, 0x62, 0x9b, 0xa0,
	0x41, 0x66, 0x94, 0xc6, 0x61, 0x92, 0xd3, 0x38, 0x0d, 0x17, 0x3b, 0xca, 0xc8, 0x5b, 0x4b, 0xa8,
	0x28, 0x48, 0xf2, 0x32, 0xe2, 0xb3, 0xdb, 0xec, 0x9b, 0xc7, 0x52, 0x30, 0x50, 0x55, 0xab, 0xb3,
	0xae, 0x6d, 0x64, 0x44, 0x59, 0x8e, 0x1c, 0x94, 0xe0, 0x20, 0x25, 0x19, 0x4b, 0x19, 0xcd, 0x5e,
	0xf7, 0x5f, 0x5b, 0x2e, 0xbb, 0x98, 0xc3, 0x2f, 0x51, 0x97, 0x22, 0x41, 0x76, 0xef, 0xae, 0x9c,
	0xa2, 0xe9, 0x9d, 0x0f, 0xaf, 0x50, 0x8f, 0x04, 0x49, 0xd6, 0x0e, 0x1b, 0x46, 0x75, 0x6e, 0x19,
	0x85, 0xc7, 0x85, 0x32, 0x13, 0xaf, 0xb3, 0xbc, 0xa1, 0x66, 0x04, 0x3b, 0x8f, 0x1f, 0xdf, 0x4f,
	0xad, 0x0e, 0x6b, 0x5b, 0x97, 0x85, 0x91, 0x2a, 0x63, 0x50, 0xa2, 0x5b, 0xd6, 0x61, 0xcf, 0x8e,
	0xbb, 0x7e, 0x94, 0x93, 0xc4, 0x01, 0xaa, 0x9c, 0x26, 0x05, 0x4a, 0x73, 0xa2, 0xa9, 0xe8, 0x43,
	0x72, 0x3b, 0x56, 0x6c, 0x6b, 0x79, 0xd2, 0x70, 0x70, 0xa6, 0xec, 0x74, 0x5e, 0xd0, 0x6c, 0xb2,
	0x1b, 0x29, 0xa3, 0x70, 0x8e, 0x0c, 0x22, 0xeb, 0xb3, 0x17, 0x6c, 0xff, 0xe5, 0x13, 0xdb, 0x30,
	0x7e, 0x27, 0xd2, 0x07, 0x9e, 0x68, 0x04, 0x83, 0x87, 0xce, 0x6e, 0xbf, 0xf7, 0x36, 0xf5, 0x9d,
	0xf7, 0xa9, 0xef, 0x7c, 0x4e, 0x7d, 0xe7, 0xe5, 0xcb, 0xff, 0x77, 0xcd, 0x33, 0x61, 0x6e, 0xc6,
	0x71, 0x25, 0xc1, 0x7f, 0x23, 0xbb, 0x05, 0x48, 0xc8, 0x30, 0xad, 0xb8, 0xa3, 0x06, 0x38, 0xfe,
	0x6f, 0x17, 0x7f, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xee, 0xaa, 0xa2, 0xf6, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReplicationServiceClient is the client API for ReplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReplicationServiceClient interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error)
	// Takes a backup and creates a deployment from it. For all intents and purposes this new deployment
	// will be the same as the deployment at that exact moment when the backup was taken from it. This means that
	// the new deployment will be in the same region and in the same project and use the same provider as the old
	// deployment did. Furthermore, the new deployment will have the same server settings ( count, mode, replication
	// factor ) as the old deployment did at the time of taking the backup. After the new deployment successfuly
	// started, the backup will be used to restore the data into the new deployment. The new deployment will have
	// a different endpoint, and the password will also be reset for it. All other user settings will remain the same.
	// The old deployment will not be touched.
	// Required permissions:
	// - replication.deployment.create
	CloneDeploymentFromBackup(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v11.Deployment, error)
}

type replicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewReplicationServiceClient(cc *grpc.ClientConn) ReplicationServiceClient {
	return &replicationServiceClient{cc}
}

func (c *replicationServiceClient) GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error) {
	out := new(v1.Version)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.replication.v1.ReplicationService/GetAPIVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) CloneDeploymentFromBackup(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v11.Deployment, error) {
	out := new(v11.Deployment)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.replication.v1.ReplicationService/CloneDeploymentFromBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServiceServer is the server API for ReplicationService service.
type ReplicationServiceServer interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error)
	// Takes a backup and creates a deployment from it. For all intents and purposes this new deployment
	// will be the same as the deployment at that exact moment when the backup was taken from it. This means that
	// the new deployment will be in the same region and in the same project and use the same provider as the old
	// deployment did. Furthermore, the new deployment will have the same server settings ( count, mode, replication
	// factor ) as the old deployment did at the time of taking the backup. After the new deployment successfuly
	// started, the backup will be used to restore the data into the new deployment. The new deployment will have
	// a different endpoint, and the password will also be reset for it. All other user settings will remain the same.
	// The old deployment will not be touched.
	// Required permissions:
	// - replication.deployment.create
	CloneDeploymentFromBackup(context.Context, *v1.IDOptions) (*v11.Deployment, error)
}

// UnimplementedReplicationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReplicationServiceServer struct {
}

func (*UnimplementedReplicationServiceServer) GetAPIVersion(ctx context.Context, req *v1.Empty) (*v1.Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIVersion not implemented")
}
func (*UnimplementedReplicationServiceServer) CloneDeploymentFromBackup(ctx context.Context, req *v1.IDOptions) (*v11.Deployment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneDeploymentFromBackup not implemented")
}

func RegisterReplicationServiceServer(s *grpc.Server, srv ReplicationServiceServer) {
	s.RegisterService(&_ReplicationService_serviceDesc, srv)
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
		FullMethod: "/arangodb.cloud.replication.v1.ReplicationService/GetAPIVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).GetAPIVersion(ctx, req.(*v1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_CloneDeploymentFromBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).CloneDeploymentFromBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.replication.v1.ReplicationService/CloneDeploymentFromBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).CloneDeploymentFromBackup(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReplicationService_serviceDesc = grpc.ServiceDesc{
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "replication.proto",
}
