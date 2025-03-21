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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.3
// source: monitoring.proto

package v1

import (
	context "context"
	v1 "github.com/arangodb-managed/apis/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for getting deployment metrics
type GetDeploymentMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the deployment for which metrics are being requested.
	// This is a required field.
	DeploymentId string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	// Start time for the query.
	// Note: maximum allowed range is 1 week.
	// This is a required field.
	StartAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	// The end time for the query.
	// Note: maximum allowed range is 1 week.
	// This is a required field.
	EndAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	// The type of server for which metrics are being requested.
	// Should be one of the following values:
	// - Dbserver
	// - Coordinator
	// - Single
	ServerType string `protobuf:"bytes,4,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty"`
	// The type of metric being requested.
	// Should be one of the following values:
	// - cpu
	// - memory
	// - disk
	MetricType string `protobuf:"bytes,5,opt,name=metric_type,json=metricType,proto3" json:"metric_type,omitempty"`
}

func (x *GetDeploymentMetricsRequest) Reset() {
	*x = GetDeploymentMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentMetricsRequest) ProtoMessage() {}

func (x *GetDeploymentMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentMetricsRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentMetricsRequest) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_rawDescGZIP(), []int{0}
}

func (x *GetDeploymentMetricsRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *GetDeploymentMetricsRequest) GetStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *GetDeploymentMetricsRequest) GetEndAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *GetDeploymentMetricsRequest) GetServerType() string {
	if x != nil {
		return x.ServerType
	}
	return ""
}

func (x *GetDeploymentMetricsRequest) GetMetricType() string {
	if x != nil {
		return x.MetricType
	}
	return ""
}

// DeploymentMetrics contains the deployment metrics
type DeploymentMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains a list of timeseries metrics for each server (of type specified in the request).
	Metrics []*DeploymentMetrics_Timeseries `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *DeploymentMetrics) Reset() {
	*x = DeploymentMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentMetrics) ProtoMessage() {}

func (x *DeploymentMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentMetrics.ProtoReflect.Descriptor instead.
func (*DeploymentMetrics) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_rawDescGZIP(), []int{1}
}

func (x *DeploymentMetrics) GetMetrics() []*DeploymentMetrics_Timeseries {
	if x != nil {
		return x.Metrics
	}
	return nil
}

// GetDeploymentLogsRequest contains request arguments for GetDeploymentLogs.
type GetDeploymentLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of the deployment to get the logs from.
	DeploymentId string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	// If set, limit logs to servers of this role.
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	// Formatting for log messages.
	// The possible values are text and json, with default value being text.
	Format string `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	// If set limits logs to the server with this ID.
	ServerId string `protobuf:"bytes,4,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// The start time for the query. Defaults to one hour ago.
	StartAt *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	// The end time for the query. Defaults to now.
	EndAt *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	// Limit the number of log lines. Defaults to 1000.
	Limit int32 `protobuf:"varint,102,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetDeploymentLogsRequest) Reset() {
	*x = GetDeploymentLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentLogsRequest) ProtoMessage() {}

func (x *GetDeploymentLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentLogsRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentLogsRequest) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeploymentLogsRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *GetDeploymentLogsRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetDeploymentLogsRequest) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *GetDeploymentLogsRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *GetDeploymentLogsRequest) GetStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *GetDeploymentLogsRequest) GetEndAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *GetDeploymentLogsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type DeploymentLogsChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chunk of bytes from the logs
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *DeploymentLogsChunk) Reset() {
	*x = DeploymentLogsChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentLogsChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentLogsChunk) ProtoMessage() {}

func (x *DeploymentLogsChunk) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentLogsChunk.ProtoReflect.Descriptor instead.
func (*DeploymentLogsChunk) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_rawDescGZIP(), []int{3}
}

func (x *DeploymentLogsChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

// Sample defines a single data point.
// It contains the value of the sample and the timestamp at which it was captured.
type DeploymentMetrics_Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp at which this sample is captured.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Value of the given sample.
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DeploymentMetrics_Sample) Reset() {
	*x = DeploymentMetrics_Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentMetrics_Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentMetrics_Sample) ProtoMessage() {}

func (x *DeploymentMetrics_Sample) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentMetrics_Sample.ProtoReflect.Descriptor instead.
func (*DeploymentMetrics_Sample) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DeploymentMetrics_Sample) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DeploymentMetrics_Sample) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// A timeseries contains a list of samples recorded at different (regular) intervals
// for a single database server identified by server_id.
type DeploymentMetrics_Timeseries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the server.
	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// List of metric samples for the server identified by the server_id.
	Samples []*DeploymentMetrics_Sample `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *DeploymentMetrics_Timeseries) Reset() {
	*x = DeploymentMetrics_Timeseries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitoring_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentMetrics_Timeseries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentMetrics_Timeseries) ProtoMessage() {}

func (x *DeploymentMetrics_Timeseries) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentMetrics_Timeseries.ProtoReflect.Descriptor instead.
func (*DeploymentMetrics_Timeseries) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_rawDescGZIP(), []int{1, 1}
}

func (x *DeploymentMetrics_Timeseries) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *DeploymentMetrics_Timeseries) GetSamples() []*DeploymentMetrics_Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

var File_monitoring_proto protoreflect.FileDescriptor

var file_monitoring_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc0, 0x02, 0x0a, 0x11, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x54,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3a, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x1a, 0x58, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x7b,
	0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x07, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x61, 0x72, 0x61,
	0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x88, 0x02, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x61, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x31, 0x0a,
	0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x2b, 0x0a, 0x13, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x32, 0xfa, 0x03, 0x0a, 0x11, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x61, 0x72, 0x61,
	0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x61, 0x72,
	0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0xb9, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x36, 0x2e, 0x61,
	0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x3a,
	0x01, 0x2a, 0x22, 0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x6f, 0x67, 0x73,
	0x30, 0x01, 0x12, 0xab, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x39, 0x2e, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x72,
	0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitoring_proto_rawDescOnce sync.Once
	file_monitoring_proto_rawDescData = file_monitoring_proto_rawDesc
)

func file_monitoring_proto_rawDescGZIP() []byte {
	file_monitoring_proto_rawDescOnce.Do(func() {
		file_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitoring_proto_rawDescData)
	})
	return file_monitoring_proto_rawDescData
}

var file_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_monitoring_proto_goTypes = []interface{}{
	(*GetDeploymentMetricsRequest)(nil),  // 0: arangodb.cloud.monitoring.v1.GetDeploymentMetricsRequest
	(*DeploymentMetrics)(nil),            // 1: arangodb.cloud.monitoring.v1.DeploymentMetrics
	(*GetDeploymentLogsRequest)(nil),     // 2: arangodb.cloud.monitoring.v1.GetDeploymentLogsRequest
	(*DeploymentLogsChunk)(nil),          // 3: arangodb.cloud.monitoring.v1.DeploymentLogsChunk
	(*DeploymentMetrics_Sample)(nil),     // 4: arangodb.cloud.monitoring.v1.DeploymentMetrics.Sample
	(*DeploymentMetrics_Timeseries)(nil), // 5: arangodb.cloud.monitoring.v1.DeploymentMetrics.Timeseries
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
	(*v1.Empty)(nil),                     // 7: arangodb.cloud.common.v1.Empty
	(*v1.Version)(nil),                   // 8: arangodb.cloud.common.v1.Version
}
var file_monitoring_proto_depIdxs = []int32{
	6,  // 0: arangodb.cloud.monitoring.v1.GetDeploymentMetricsRequest.start_at:type_name -> google.protobuf.Timestamp
	6,  // 1: arangodb.cloud.monitoring.v1.GetDeploymentMetricsRequest.end_at:type_name -> google.protobuf.Timestamp
	5,  // 2: arangodb.cloud.monitoring.v1.DeploymentMetrics.metrics:type_name -> arangodb.cloud.monitoring.v1.DeploymentMetrics.Timeseries
	6,  // 3: arangodb.cloud.monitoring.v1.GetDeploymentLogsRequest.start_at:type_name -> google.protobuf.Timestamp
	6,  // 4: arangodb.cloud.monitoring.v1.GetDeploymentLogsRequest.end_at:type_name -> google.protobuf.Timestamp
	6,  // 5: arangodb.cloud.monitoring.v1.DeploymentMetrics.Sample.timestamp:type_name -> google.protobuf.Timestamp
	4,  // 6: arangodb.cloud.monitoring.v1.DeploymentMetrics.Timeseries.samples:type_name -> arangodb.cloud.monitoring.v1.DeploymentMetrics.Sample
	7,  // 7: arangodb.cloud.monitoring.v1.MonitoringService.GetAPIVersion:input_type -> arangodb.cloud.common.v1.Empty
	2,  // 8: arangodb.cloud.monitoring.v1.MonitoringService.GetDeploymentLogs:input_type -> arangodb.cloud.monitoring.v1.GetDeploymentLogsRequest
	0,  // 9: arangodb.cloud.monitoring.v1.MonitoringService.GetDeploymentUsageMetrics:input_type -> arangodb.cloud.monitoring.v1.GetDeploymentMetricsRequest
	8,  // 10: arangodb.cloud.monitoring.v1.MonitoringService.GetAPIVersion:output_type -> arangodb.cloud.common.v1.Version
	3,  // 11: arangodb.cloud.monitoring.v1.MonitoringService.GetDeploymentLogs:output_type -> arangodb.cloud.monitoring.v1.DeploymentLogsChunk
	1,  // 12: arangodb.cloud.monitoring.v1.MonitoringService.GetDeploymentUsageMetrics:output_type -> arangodb.cloud.monitoring.v1.DeploymentMetrics
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_monitoring_proto_init() }
func file_monitoring_proto_init() {
	if File_monitoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentMetricsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentMetrics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitoring_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentLogsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitoring_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentLogsChunk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitoring_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentMetrics_Sample); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitoring_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentMetrics_Timeseries); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_monitoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitoring_proto_goTypes,
		DependencyIndexes: file_monitoring_proto_depIdxs,
		MessageInfos:      file_monitoring_proto_msgTypes,
	}.Build()
	File_monitoring_proto = out.File
	file_monitoring_proto_rawDesc = nil
	file_monitoring_proto_goTypes = nil
	file_monitoring_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MonitoringServiceClient is the client API for MonitoringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitoringServiceClient interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error)
	// Fetch all logs of the deployment that matches the given request.
	// Required permissions:
	// - monitoring.logs.get on the deployment identified by the given deployment ID.
	GetDeploymentLogs(ctx context.Context, in *GetDeploymentLogsRequest, opts ...grpc.CallOption) (MonitoringService_GetDeploymentLogsClient, error)
	// Get the usage metrics for the deployment based on the given request.
	// Note: Only at most 1 week worth of data may be requested.
	// Required permissions:
	// - monitoring.metrics.get on the deployment identified by the given deployment ID.
	GetDeploymentUsageMetrics(ctx context.Context, in *GetDeploymentMetricsRequest, opts ...grpc.CallOption) (*DeploymentMetrics, error)
}

type monitoringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitoringServiceClient(cc grpc.ClientConnInterface) MonitoringServiceClient {
	return &monitoringServiceClient{cc}
}

func (c *monitoringServiceClient) GetAPIVersion(ctx context.Context, in *v1.Empty, opts ...grpc.CallOption) (*v1.Version, error) {
	out := new(v1.Version)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.monitoring.v1.MonitoringService/GetAPIVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringServiceClient) GetDeploymentLogs(ctx context.Context, in *GetDeploymentLogsRequest, opts ...grpc.CallOption) (MonitoringService_GetDeploymentLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MonitoringService_serviceDesc.Streams[0], "/arangodb.cloud.monitoring.v1.MonitoringService/GetDeploymentLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitoringServiceGetDeploymentLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MonitoringService_GetDeploymentLogsClient interface {
	Recv() (*DeploymentLogsChunk, error)
	grpc.ClientStream
}

type monitoringServiceGetDeploymentLogsClient struct {
	grpc.ClientStream
}

func (x *monitoringServiceGetDeploymentLogsClient) Recv() (*DeploymentLogsChunk, error) {
	m := new(DeploymentLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *monitoringServiceClient) GetDeploymentUsageMetrics(ctx context.Context, in *GetDeploymentMetricsRequest, opts ...grpc.CallOption) (*DeploymentMetrics, error) {
	out := new(DeploymentMetrics)
	err := c.cc.Invoke(ctx, "/arangodb.cloud.monitoring.v1.MonitoringService/GetDeploymentUsageMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitoringServiceServer is the server API for MonitoringService service.
type MonitoringServiceServer interface {
	// Get the current API version of this service.
	// Required permissions:
	// - None
	GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error)
	// Fetch all logs of the deployment that matches the given request.
	// Required permissions:
	// - monitoring.logs.get on the deployment identified by the given deployment ID.
	GetDeploymentLogs(*GetDeploymentLogsRequest, MonitoringService_GetDeploymentLogsServer) error
	// Get the usage metrics for the deployment based on the given request.
	// Note: Only at most 1 week worth of data may be requested.
	// Required permissions:
	// - monitoring.metrics.get on the deployment identified by the given deployment ID.
	GetDeploymentUsageMetrics(context.Context, *GetDeploymentMetricsRequest) (*DeploymentMetrics, error)
}

// UnimplementedMonitoringServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMonitoringServiceServer struct {
}

func (*UnimplementedMonitoringServiceServer) GetAPIVersion(context.Context, *v1.Empty) (*v1.Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIVersion not implemented")
}
func (*UnimplementedMonitoringServiceServer) GetDeploymentLogs(*GetDeploymentLogsRequest, MonitoringService_GetDeploymentLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDeploymentLogs not implemented")
}
func (*UnimplementedMonitoringServiceServer) GetDeploymentUsageMetrics(context.Context, *GetDeploymentMetricsRequest) (*DeploymentMetrics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeploymentUsageMetrics not implemented")
}

func RegisterMonitoringServiceServer(s *grpc.Server, srv MonitoringServiceServer) {
	s.RegisterService(&_MonitoringService_serviceDesc, srv)
}

func _MonitoringService_GetAPIVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServiceServer).GetAPIVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.monitoring.v1.MonitoringService/GetAPIVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServiceServer).GetAPIVersion(ctx, req.(*v1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitoringService_GetDeploymentLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDeploymentLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitoringServiceServer).GetDeploymentLogs(m, &monitoringServiceGetDeploymentLogsServer{stream})
}

type MonitoringService_GetDeploymentLogsServer interface {
	Send(*DeploymentLogsChunk) error
	grpc.ServerStream
}

type monitoringServiceGetDeploymentLogsServer struct {
	grpc.ServerStream
}

func (x *monitoringServiceGetDeploymentLogsServer) Send(m *DeploymentLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _MonitoringService_GetDeploymentUsageMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServiceServer).GetDeploymentUsageMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arangodb.cloud.monitoring.v1.MonitoringService/GetDeploymentUsageMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServiceServer).GetDeploymentUsageMetrics(ctx, req.(*GetDeploymentMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitoringService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arangodb.cloud.monitoring.v1.MonitoringService",
	HandlerType: (*MonitoringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAPIVersion",
			Handler:    _MonitoringService_GetAPIVersion_Handler,
		},
		{
			MethodName: "GetDeploymentUsageMetrics",
			Handler:    _MonitoringService_GetDeploymentUsageMetrics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDeploymentLogs",
			Handler:       _MonitoringService_GetDeploymentLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "monitoring.proto",
}
