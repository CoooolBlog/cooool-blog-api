/*
 * Copyright 2022 Kristian Huang <krishuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package v1

import (
	"context"
	"reflect"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ListSecretsRequest defines ListSecrets request struct.
type ListSecretsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset *int64 `protobuf:"varint,1,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit  *int64 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *ListSecretsRequest) Reset() {
	*x = ListSecretsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apiserver_v1_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSecretsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSecretsRequest) ProtoMessage() {}

func (x *ListSecretsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apiserver_v1_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSecretsRequest.ProtoReflect.Descriptor instead.
func (*ListSecretsRequest) Descriptor() ([]byte, []int) {
	return file_proto_apiserver_v1_cache_proto_rawDescGZIP(), []int{0}
}

func (x *ListSecretsRequest) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *ListSecretsRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

// SecretInfo contains secret details.
type SecretInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SecretId    string `protobuf:"bytes,2,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	Username    string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	SecretKey   string `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Expires     int64  `protobuf:"varint,5,opt,name=expires,proto3" json:"expires,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SecretInfo) Reset() {
	*x = SecretInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apiserver_v1_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretInfo) ProtoMessage() {}

func (x *SecretInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apiserver_v1_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretInfo.ProtoReflect.Descriptor instead.
func (*SecretInfo) Descriptor() ([]byte, []int) {
	return file_proto_apiserver_v1_cache_proto_rawDescGZIP(), []int{1}
}

func (x *SecretInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecretInfo) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *SecretInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SecretInfo) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *SecretInfo) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

func (x *SecretInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SecretInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SecretInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// ListSecretsResponse defines ListSecrets response struct.
type ListSecretsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64         `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Items      []*SecretInfo `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListSecretsResponse) Reset() {
	*x = ListSecretsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apiserver_v1_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSecretsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSecretsResponse) ProtoMessage() {}

func (x *ListSecretsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apiserver_v1_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSecretsResponse.ProtoReflect.Descriptor instead.
func (*ListSecretsResponse) Descriptor() ([]byte, []int) {
	return file_proto_apiserver_v1_cache_proto_rawDescGZIP(), []int{2}
}

func (x *ListSecretsResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListSecretsResponse) GetItems() []*SecretInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

// ListPoliciesRequest defines ListPolicies request struct.
type ListPoliciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset *int64 `protobuf:"varint,1,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit  *int64 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *ListPoliciesRequest) Reset() {
	*x = ListPoliciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apiserver_v1_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesRequest) ProtoMessage() {}

func (x *ListPoliciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apiserver_v1_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesRequest.ProtoReflect.Descriptor instead.
func (*ListPoliciesRequest) Descriptor() ([]byte, []int) {
	return file_proto_apiserver_v1_cache_proto_rawDescGZIP(), []int{3}
}

func (x *ListPoliciesRequest) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *ListPoliciesRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

// PolicyInfo contains policy details.
type PolicyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	PolicyStr    string `protobuf:"bytes,3,opt,name=policy_str,json=policyStr,proto3" json:"policy_str,omitempty"`
	PolicyShadow string `protobuf:"bytes,4,opt,name=policy_shadow,json=policyShadow,proto3" json:"policy_shadow,omitempty"`
	CreatedAt    string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PolicyInfo) Reset() {
	*x = PolicyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apiserver_v1_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyInfo) ProtoMessage() {}

func (x *PolicyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apiserver_v1_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyInfo.ProtoReflect.Descriptor instead.
func (*PolicyInfo) Descriptor() ([]byte, []int) {
	return file_proto_apiserver_v1_cache_proto_rawDescGZIP(), []int{4}
}

func (x *PolicyInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PolicyInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PolicyInfo) GetPolicyStr() string {
	if x != nil {
		return x.PolicyStr
	}
	return ""
}

func (x *PolicyInfo) GetPolicyShadow() string {
	if x != nil {
		return x.PolicyShadow
	}
	return ""
}

func (x *PolicyInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// ListPoliciesResponse defines ListPolicies response struct.
type ListPoliciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64         `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Items      []*PolicyInfo `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPoliciesResponse) Reset() {
	*x = ListPoliciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apiserver_v1_cache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesResponse) ProtoMessage() {}

func (x *ListPoliciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apiserver_v1_cache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesResponse.ProtoReflect.Descriptor instead.
func (*ListPoliciesResponse) Descriptor() ([]byte, []int) {
	return file_proto_apiserver_v1_cache_proto_rawDescGZIP(), []int{5}
}

func (x *ListPoliciesResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListPoliciesResponse) GetItems() []*PolicyInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_apiserver_v1_cache_proto protoreflect.FileDescriptor

var file_proto_apiserver_v1_cache_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xf2, 0x01, 0x0a, 0x0a, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x5f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x62, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x74,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53,
	0x74, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x60, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x05, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x75, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_apiserver_v1_cache_proto_rawDescOnce sync.Once
	file_proto_apiserver_v1_cache_proto_rawDescData = file_proto_apiserver_v1_cache_proto_rawDesc
)

func file_proto_apiserver_v1_cache_proto_rawDescGZIP() []byte {
	file_proto_apiserver_v1_cache_proto_rawDescOnce.Do(func() {
		file_proto_apiserver_v1_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_apiserver_v1_cache_proto_rawDescData)
	})
	return file_proto_apiserver_v1_cache_proto_rawDescData
}

var file_proto_apiserver_v1_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_apiserver_v1_cache_proto_goTypes = []interface{}{
	(*ListSecretsRequest)(nil),   // 0: proto.ListSecretsRequest
	(*SecretInfo)(nil),           // 1: proto.SecretInfo
	(*ListSecretsResponse)(nil),  // 2: proto.ListSecretsResponse
	(*ListPoliciesRequest)(nil),  // 3: proto.ListPoliciesRequest
	(*PolicyInfo)(nil),           // 4: proto.PolicyInfo
	(*ListPoliciesResponse)(nil), // 5: proto.ListPoliciesResponse
}
var file_proto_apiserver_v1_cache_proto_depIdxs = []int32{
	1, // 0: proto.ListSecretsResponse.items:type_name -> proto.SecretInfo
	4, // 1: proto.ListPoliciesResponse.items:type_name -> proto.PolicyInfo
	0, // 2: proto.Cache.ListSecrets:input_type -> proto.ListSecretsRequest
	3, // 3: proto.Cache.ListPolicies:input_type -> proto.ListPoliciesRequest
	2, // 4: proto.Cache.ListSecrets:output_type -> proto.ListSecretsResponse
	5, // 5: proto.Cache.ListPolicies:output_type -> proto.ListPoliciesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_apiserver_v1_cache_proto_init() }
func file_proto_apiserver_v1_cache_proto_init() {
	if File_proto_apiserver_v1_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_apiserver_v1_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSecretsRequest); i {
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
		file_proto_apiserver_v1_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretInfo); i {
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
		file_proto_apiserver_v1_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSecretsResponse); i {
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
		file_proto_apiserver_v1_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesRequest); i {
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
		file_proto_apiserver_v1_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyInfo); i {
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
		file_proto_apiserver_v1_cache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesResponse); i {
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
	file_proto_apiserver_v1_cache_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_apiserver_v1_cache_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_apiserver_v1_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_apiserver_v1_cache_proto_goTypes,
		DependencyIndexes: file_proto_apiserver_v1_cache_proto_depIdxs,
		MessageInfos:      file_proto_apiserver_v1_cache_proto_msgTypes,
	}.Build()
	File_proto_apiserver_v1_cache_proto = out.File
	file_proto_apiserver_v1_cache_proto_rawDesc = nil
	file_proto_apiserver_v1_cache_proto_goTypes = nil
	file_proto_apiserver_v1_cache_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheClient interface {
	ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (*ListSecretsResponse, error)
	ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error)
}

type cacheClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheClient(cc grpc.ClientConnInterface) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (*ListSecretsResponse, error) {
	out := new(ListSecretsResponse)
	err := c.cc.Invoke(ctx, "/proto.Cache/ListSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error) {
	out := new(ListPoliciesResponse)
	err := c.cc.Invoke(ctx, "/proto.Cache/ListPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
type CacheServer interface {
	ListSecrets(context.Context, *ListSecretsRequest) (*ListSecretsResponse, error)
	ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error)
}

// UnimplementedCacheServer can be embedded to have forward compatible implementations.
type UnimplementedCacheServer struct {
}

func (*UnimplementedCacheServer) ListSecrets(context.Context, *ListSecretsRequest) (*ListSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecrets not implemented")
}
func (*UnimplementedCacheServer) ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicies not implemented")
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_ListSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).ListSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cache/ListSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).ListSecrets(ctx, req.(*ListSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_ListPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).ListPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cache/ListPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).ListPolicies(ctx, req.(*ListPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSecrets",
			Handler:    _Cache_ListSecrets_Handler,
		},
		{
			MethodName: "ListPolicies",
			Handler:    _Cache_ListPolicies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/apiserver/v1/cache.proto",
}