// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/arangodb-managed/apis/common/v1"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// IPWhitelist represents a list of CIDR ranges from which a deployment is accessible.
type IPWhitelist struct {
	// System identifier of the whitelist.
	// This is a read-only value.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// URL of the whitelist.
	// This is a read-only value.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Name of the whitelist.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the whitelist.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Identifier of the project that contains this whitelist.
	ProjectId string `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// List of CIDR ranges.
	// Values must follow format as defined in RFC 4632 and RFC 4291.
	CidrRanges []string `protobuf:"bytes,6,rep,name=cidr_ranges,json=cidrRanges,proto3" json:"cidr_ranges,omitempty"`
	// The creation timestamp of this whitelist.
	// This is a read-only value.
	CreatedAt *types.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The deletion timestamp of the whitelist
	// This is a read-only value.
	DeletedAt *types.Timestamp `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	// Set when this whitelist is deleted.
	// This is a read-only value.
	IsDeleted bool `protobuf:"varint,9,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	// Identifier of the user who created this whitelist.
	// This is a read-only value.
	CreatedById          string   `protobuf:"bytes,10,opt,name=created_by_id,json=createdById,proto3" json:"created_by_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPWhitelist) Reset()         { *m = IPWhitelist{} }
func (m *IPWhitelist) String() string { return proto.CompactTextString(m) }
func (*IPWhitelist) ProtoMessage()    {}
func (*IPWhitelist) Descriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{0}
}
func (m *IPWhitelist) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IPWhitelist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IPWhitelist.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IPWhitelist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPWhitelist.Merge(m, src)
}
func (m *IPWhitelist) XXX_Size() int {
	return m.Size()
}
func (m *IPWhitelist) XXX_DiscardUnknown() {
	xxx_messageInfo_IPWhitelist.DiscardUnknown(m)
}

var xxx_messageInfo_IPWhitelist proto.InternalMessageInfo

func (m *IPWhitelist) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IPWhitelist) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *IPWhitelist) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IPWhitelist) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *IPWhitelist) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *IPWhitelist) GetCidrRanges() []string {
	if m != nil {
		return m.CidrRanges
	}
	return nil
}

func (m *IPWhitelist) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *IPWhitelist) GetDeletedAt() *types.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *IPWhitelist) GetIsDeleted() bool {
	if m != nil {
		return m.IsDeleted
	}
	return false
}

func (m *IPWhitelist) GetCreatedById() string {
	if m != nil {
		return m.CreatedById
	}
	return ""
}

// List of IP whitelists.
type IPWhitelistList struct {
	Items                []*IPWhitelist `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IPWhitelistList) Reset()         { *m = IPWhitelistList{} }
func (m *IPWhitelistList) String() string { return proto.CompactTextString(m) }
func (*IPWhitelistList) ProtoMessage()    {}
func (*IPWhitelistList) Descriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{1}
}
func (m *IPWhitelistList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IPWhitelistList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IPWhitelistList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IPWhitelistList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPWhitelistList.Merge(m, src)
}
func (m *IPWhitelistList) XXX_Size() int {
	return m.Size()
}
func (m *IPWhitelistList) XXX_DiscardUnknown() {
	xxx_messageInfo_IPWhitelistList.DiscardUnknown(m)
}

var xxx_messageInfo_IPWhitelistList proto.InternalMessageInfo

func (m *IPWhitelistList) GetItems() []*IPWhitelist {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*IPWhitelist)(nil), "arangodb.cloud.security.v1.IPWhitelist")
	proto.RegisterType((*IPWhitelistList)(nil), "arangodb.cloud.security.v1.IPWhitelistList")
}

func init() { proto.RegisterFile("security.proto", fileDescriptor_55a487c716a8b59c) }

var fileDescriptor_55a487c716a8b59c = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6b, 0xd4, 0x4e,
	0x18, 0xc7, 0x7f, 0xb3, 0xdb, 0xf6, 0xd7, 0xcc, 0x62, 0xff, 0xcc, 0x41, 0xc2, 0xa2, 0xdb, 0x10,
	0x95, 0x2e, 0xd5, 0x66, 0xe8, 0x16, 0x11, 0x2d, 0x15, 0x5a, 0x2b, 0xb2, 0x20, 0x58, 0x52, 0x45,
	0xf0, 0xb2, 0xcc, 0x66, 0xc6, 0x74, 0x24, 0xc9, 0x84, 0xcc, 0xec, 0xea, 0x52, 0x7a, 0x11, 0xbc,
	0xf4, 0xaa, 0x07, 0xef, 0xfa, 0x62, 0x3c, 0x0a, 0xde, 0x45, 0xaa, 0x2f, 0x44, 0x66, 0x92, 0xd8,
	0xa8, 0x2c, 0xc6, 0x83, 0xb7, 0xc9, 0x33, 0xcf, 0xf7, 0x99, 0xcf, 0x7c, 0x9f, 0x27, 0x03, 0x17,
	0x24, 0x0b, 0x46, 0x19, 0x57, 0x13, 0x2f, 0xcd, 0x84, 0x12, 0xa8, 0x4d, 0x32, 0x92, 0x84, 0x82,
	0x0e, 0xbd, 0x20, 0x12, 0x23, 0xea, 0xfd, 0xd8, 0x1e, 0x6f, 0xb4, 0xcf, 0x07, 0x22, 0x8e, 0x45,
	0x82, 0xc7, 0x1b, 0x38, 0x5f, 0xe5, 0x9a, 0xf6, 0x56, 0xc8, 0xd5, 0xe1, 0x68, 0xe8, 0x05, 0x22,
	0xc6, 0xa1, 0x88, 0x48, 0x12, 0x62, 0xb3, 0x31, 0x1c, 0x3d, 0xc5, 0xa9, 0x9a, 0xa4, 0x4c, 0x62,
	0xc5, 0x63, 0x26, 0x15, 0x89, 0xd3, 0xb3, 0x55, 0x21, 0xbe, 0x10, 0x0a, 0x11, 0x46, 0x0c, 0x93,
	0x94, 0x63, 0x92, 0x24, 0x42, 0x11, 0xc5, 0x45, 0x22, 0xf3, 0x5d, 0xf7, 0x73, 0x03, 0xb6, 0xfa,
	0xfb, 0x8f, 0x0f, 0xb9, 0x62, 0x11, 0x97, 0x0a, 0x2d, 0xc0, 0x06, 0xa7, 0x36, 0x70, 0x40, 0xd7,
	0xf2, 0x1b, 0x9c, 0xa2, 0x25, 0xd8, 0x1c, 0x65, 0x91, 0xdd, 0x30, 0x01, 0xbd, 0x44, 0x08, 0xce,
	0x24, 0x24, 0x66, 0x76, 0xd3, 0x84, 0xcc, 0x1a, 0x39, 0xb0, 0x45, 0x99, 0x0c, 0x32, 0x9e, 0xea,
	0xda, 0xf6, 0x8c, 0xd9, 0xaa, 0x86, 0xd0, 0x45, 0x08, 0xd3, 0x4c, 0x3c, 0x63, 0x81, 0x1a, 0x70,
	0x6a, 0xcf, 0x9a, 0x04, 0xab, 0x88, 0xf4, 0x29, 0x5a, 0x81, 0xad, 0x80, 0xd3, 0x6c, 0xa0, 0xbd,
	0x61, 0xd2, 0x9e, 0x73, 0x9a, 0x5d, 0xcb, 0x87, 0x3a, 0xe4, 0x9b, 0x08, 0xba, 0x09, 0x61, 0x90,
	0x31, 0xa2, 0x18, 0x1d, 0x10, 0x65, 0xff, 0xef, 0x80, 0x6e, 0xab, 0xd7, 0xf6, 0xf2, 0xab, 0x79,
	0xa5, 0x19, 0xde, 0xc3, 0xf2, 0xee, 0xbe, 0x55, 0x64, 0xef, 0x28, 0x2d, 0xa5, 0x2c, 0x62, 0x85,
	0x74, 0xfe, 0xcf, 0xd2, 0x22, 0x7b, 0x47, 0x69, 0x6a, 0x2e, 0x07, 0xc5, 0xb7, 0x6d, 0x39, 0xa0,
	0x3b, 0xef, 0x5b, 0x5c, 0xee, 0xe5, 0x01, 0xe4, 0xc2, 0x73, 0x25, 0xd4, 0x70, 0xa2, 0xef, 0x05,
	0xf3, 0x8b, 0x17, 0xc1, 0xdd, 0x49, 0x9f, 0xba, 0xfb, 0x70, 0xb1, 0xe2, 0xef, 0x7d, 0xed, 0xf1,
	0x36, 0x9c, 0xe5, 0x8a, 0xc5, 0xd2, 0x06, 0x4e, 0xb3, 0xdb, 0xea, 0xad, 0x7a, 0xd3, 0x47, 0xc2,
	0xab, 0x68, 0xfd, 0x5c, 0xd5, 0x3b, 0x99, 0x83, 0x8b, 0x07, 0x45, 0xca, 0x01, 0xcb, 0xc6, 0x3c,
	0x60, 0xe8, 0x1d, 0x80, 0x4b, 0xba, 0x76, 0x25, 0x5d, 0xa2, 0x2b, 0xbf, 0x16, 0x2e, 0x86, 0x6a,
	0xbc, 0xe1, 0xe9, 0xdc, 0x07, 0xa6, 0x33, 0xb2, 0x7d, 0xb5, 0xe6, 0xf9, 0x5a, 0xe3, 0x6e, 0xbd,
	0xfc, 0xf4, 0xed, 0x75, 0xe3, 0x3a, 0xda, 0x34, 0xf3, 0x54, 0x66, 0xea, 0x89, 0x2d, 0x9a, 0x29,
	0xf1, 0x51, 0x20, 0x12, 0xc5, 0x5e, 0xe8, 0x46, 0x1f, 0x63, 0x9e, 0x3e, 0x3f, 0x03, 0x3a, 0x01,
	0x70, 0xe1, 0x1e, 0xab, 0x42, 0xa2, 0x4b, 0xd3, 0x19, 0xfb, 0x7b, 0x25, 0x61, 0x5d, 0x87, 0xdc,
	0x35, 0x43, 0x77, 0x19, 0xb9, 0xbf, 0xd1, 0x55, 0x39, 0xf0, 0x11, 0xa7, 0xc7, 0xe8, 0x3d, 0x80,
	0xcb, 0x77, 0x4c, 0xa3, 0xaa, 0x3c, 0x75, 0x8f, 0xaa, 0xcf, 0xb4, 0x6d, 0x98, 0x6e, 0xb8, 0xbd,
	0x69, 0x8e, 0xe1, 0xa3, 0xb3, 0x3f, 0xe3, 0x67, 0xc3, 0x6e, 0x81, 0x35, 0xf4, 0x06, 0xc0, 0xe5,
	0x47, 0x29, 0xfd, 0xe7, 0x98, 0xeb, 0x06, 0x73, 0xb5, 0x57, 0xc3, 0x3a, 0x8d, 0xf5, 0x0a, 0xc0,
	0xe5, 0xfc, 0x37, 0xf8, 0xeb, 0x6e, 0xae, 0x4c, 0x4f, 0xba, 0x1b, 0xa7, 0x6a, 0x52, 0x76, 0x71,
	0xad, 0x06, 0xca, 0xee, 0xed, 0x0f, 0xa7, 0x1d, 0xf0, 0xf1, 0xb4, 0x03, 0xbe, 0x9c, 0x76, 0xc0,
	0xdb, 0xaf, 0x9d, 0xff, 0x9e, 0x5c, 0xab, 0x3c, 0x96, 0xe5, 0x41, 0xeb, 0x31, 0x49, 0x48, 0xc8,
	0xa8, 0x2e, 0x28, 0xab, 0x15, 0x87, 0x73, 0xe6, 0x01, 0xd8, 0xfc, 0x1e, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x32, 0x60, 0xa9, 0xa7, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecurityServiceClient is the client API for SecurityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecurityServiceClient interface {
	// Fetch all IP whitelists that belong to the project identified by the given
	// context ID.
	// Required permissions:
	// - security.ipwhitelist.list on the project identified by the given context ID.
	ListIPWhitelists(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*IPWhitelistList, error)
	// Fetch an IP whitelist by its id.
	// Required permissions:
	// - security.ipwhitelist.get on the IP whitelist
	GetIPWhitelist(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*IPWhitelist, error)
	// Create a new IP whitelist
	// Required permissions:
	// - security.ipwhitelist.create on the project that owns the IP whitelist.
	CreateIPWhitelist(ctx context.Context, in *IPWhitelist, opts ...grpc.CallOption) (*IPWhitelist, error)
	// Update an IP whitelist
	// Required permissions:
	// - security.ipwhitelist.update on the IP whitelist
	UpdateIPWhitelist(ctx context.Context, in *IPWhitelist, opts ...grpc.CallOption) (*IPWhitelist, error)
	// Delete an IP whitelist.
	// Note that IP whitelists are initially only marked for deletion.
	// Once all their dependent deployments are removed, the whitelist is removed.
	// Required permissions:
	// - security.ipwhitelist.delete on the IP whitelist
	DeleteIPWhitelist(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v1.Empty, error)
}

type securityServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecurityServiceClient(cc *grpc.ClientConn) SecurityServiceClient {
	return &securityServiceClient{cc}
}

func (c *securityServiceClient) ListIPWhitelists(ctx context.Context, in *v1.ListOptions, opts ...grpc.CallOption) (*IPWhitelistList, error) {
	out := new(IPWhitelistList)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.security.v1.SecurityService/ListIPWhitelists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) GetIPWhitelist(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*IPWhitelist, error) {
	out := new(IPWhitelist)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.security.v1.SecurityService/GetIPWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) CreateIPWhitelist(ctx context.Context, in *IPWhitelist, opts ...grpc.CallOption) (*IPWhitelist, error) {
	out := new(IPWhitelist)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.security.v1.SecurityService/CreateIPWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) UpdateIPWhitelist(ctx context.Context, in *IPWhitelist, opts ...grpc.CallOption) (*IPWhitelist, error) {
	out := new(IPWhitelist)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.security.v1.SecurityService/UpdateIPWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) DeleteIPWhitelist(ctx context.Context, in *v1.IDOptions, opts ...grpc.CallOption) (*v1.Empty, error) {
	out := new(v1.Empty)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.security.v1.SecurityService/DeleteIPWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityServiceServer is the server API for SecurityService service.
type SecurityServiceServer interface {
	// Fetch all IP whitelists that belong to the project identified by the given
	// context ID.
	// Required permissions:
	// - security.ipwhitelist.list on the project identified by the given context ID.
	ListIPWhitelists(context.Context, *v1.ListOptions) (*IPWhitelistList, error)
	// Fetch an IP whitelist by its id.
	// Required permissions:
	// - security.ipwhitelist.get on the IP whitelist
	GetIPWhitelist(context.Context, *v1.IDOptions) (*IPWhitelist, error)
	// Create a new IP whitelist
	// Required permissions:
	// - security.ipwhitelist.create on the project that owns the IP whitelist.
	CreateIPWhitelist(context.Context, *IPWhitelist) (*IPWhitelist, error)
	// Update an IP whitelist
	// Required permissions:
	// - security.ipwhitelist.update on the IP whitelist
	UpdateIPWhitelist(context.Context, *IPWhitelist) (*IPWhitelist, error)
	// Delete an IP whitelist.
	// Note that IP whitelists are initially only marked for deletion.
	// Once all their dependent deployments are removed, the whitelist is removed.
	// Required permissions:
	// - security.ipwhitelist.delete on the IP whitelist
	DeleteIPWhitelist(context.Context, *v1.IDOptions) (*v1.Empty, error)
}

// UnimplementedSecurityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecurityServiceServer struct {
}

func (*UnimplementedSecurityServiceServer) ListIPWhitelists(ctx context.Context, req *v1.ListOptions) (*IPWhitelistList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIPWhitelists not implemented")
}
func (*UnimplementedSecurityServiceServer) GetIPWhitelist(ctx context.Context, req *v1.IDOptions) (*IPWhitelist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIPWhitelist not implemented")
}
func (*UnimplementedSecurityServiceServer) CreateIPWhitelist(ctx context.Context, req *IPWhitelist) (*IPWhitelist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIPWhitelist not implemented")
}
func (*UnimplementedSecurityServiceServer) UpdateIPWhitelist(ctx context.Context, req *IPWhitelist) (*IPWhitelist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIPWhitelist not implemented")
}
func (*UnimplementedSecurityServiceServer) DeleteIPWhitelist(ctx context.Context, req *v1.IDOptions) (*v1.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIPWhitelist not implemented")
}

func RegisterSecurityServiceServer(s *grpc.Server, srv SecurityServiceServer) {
	s.RegisterService(&_SecurityService_serviceDesc, srv)
}

func _SecurityService_ListIPWhitelists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).ListIPWhitelists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.security.v1.SecurityService/ListIPWhitelists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).ListIPWhitelists(ctx, req.(*v1.ListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_GetIPWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).GetIPWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.security.v1.SecurityService/GetIPWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).GetIPWhitelist(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_CreateIPWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPWhitelist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).CreateIPWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.security.v1.SecurityService/CreateIPWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).CreateIPWhitelist(ctx, req.(*IPWhitelist))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_UpdateIPWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPWhitelist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).UpdateIPWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.security.v1.SecurityService/UpdateIPWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).UpdateIPWhitelist(ctx, req.(*IPWhitelist))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_DeleteIPWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).DeleteIPWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.security.v1.SecurityService/DeleteIPWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).DeleteIPWhitelist(ctx, req.(*v1.IDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecurityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.security.v1.SecurityService",
	HandlerType: (*SecurityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListIPWhitelists",
			Handler:    _SecurityService_ListIPWhitelists_Handler,
		},
		{
			MethodName: "GetIPWhitelist",
			Handler:    _SecurityService_GetIPWhitelist_Handler,
		},
		{
			MethodName: "CreateIPWhitelist",
			Handler:    _SecurityService_CreateIPWhitelist_Handler,
		},
		{
			MethodName: "UpdateIPWhitelist",
			Handler:    _SecurityService_UpdateIPWhitelist_Handler,
		},
		{
			MethodName: "DeleteIPWhitelist",
			Handler:    _SecurityService_DeleteIPWhitelist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security.proto",
}

func (m *IPWhitelist) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IPWhitelist) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IPWhitelist) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CreatedById) > 0 {
		i -= len(m.CreatedById)
		copy(dAtA[i:], m.CreatedById)
		i = encodeVarintSecurity(dAtA, i, uint64(len(m.CreatedById)))
		i--
		dAtA[i] = 0x52
	}
	if m.IsDeleted {
		i--
		if m.IsDeleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.DeletedAt != nil {
		{
			size, err := m.DeletedAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSecurity(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.CreatedAt != nil {
		{
			size, err := m.CreatedAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSecurity(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.CidrRanges) > 0 {
		for iNdEx := len(m.CidrRanges) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CidrRanges[iNdEx])
			copy(dAtA[i:], m.CidrRanges[iNdEx])
			i = encodeVarintSecurity(dAtA, i, uint64(len(m.CidrRanges[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ProjectId) > 0 {
		i -= len(m.ProjectId)
		copy(dAtA[i:], m.ProjectId)
		i = encodeVarintSecurity(dAtA, i, uint64(len(m.ProjectId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintSecurity(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSecurity(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintSecurity(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSecurity(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IPWhitelistList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IPWhitelistList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IPWhitelistList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSecurity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSecurity(dAtA []byte, offset int, v uint64) int {
	offset -= sovSecurity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IPWhitelist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSecurity(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovSecurity(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSecurity(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovSecurity(uint64(l))
	}
	l = len(m.ProjectId)
	if l > 0 {
		n += 1 + l + sovSecurity(uint64(l))
	}
	if len(m.CidrRanges) > 0 {
		for _, s := range m.CidrRanges {
			l = len(s)
			n += 1 + l + sovSecurity(uint64(l))
		}
	}
	if m.CreatedAt != nil {
		l = m.CreatedAt.Size()
		n += 1 + l + sovSecurity(uint64(l))
	}
	if m.DeletedAt != nil {
		l = m.DeletedAt.Size()
		n += 1 + l + sovSecurity(uint64(l))
	}
	if m.IsDeleted {
		n += 2
	}
	l = len(m.CreatedById)
	if l > 0 {
		n += 1 + l + sovSecurity(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IPWhitelistList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovSecurity(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSecurity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSecurity(x uint64) (n int) {
	return sovSecurity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IPWhitelist) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecurity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IPWhitelist: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IPWhitelist: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CidrRanges", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CidrRanges = append(m.CidrRanges, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = &types.Timestamp{}
			}
			if err := m.CreatedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeletedAt == nil {
				m.DeletedAt = &types.Timestamp{}
			}
			if err := m.DeletedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsDeleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsDeleted = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedById", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedById = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSecurity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSecurity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSecurity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IPWhitelistList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecurity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IPWhitelistList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IPWhitelistList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSecurity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSecurity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &IPWhitelist{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSecurity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSecurity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSecurity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSecurity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSecurity
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSecurity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSecurity
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSecurity
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSecurity
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSecurity(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSecurity
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSecurity = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSecurity   = fmt.Errorf("proto: integer overflow")
)
