// Code generated by copy-client.sh. DO NOT EDIT.
// version: v1.0.0
// source: https://github.com/cybercryptio/d1-service-generic.git
// commit: 88afaccef27c4ea1feb29dccf1d2a5c3866db309

// Copyright 2020-2022 CYBERCRYPT

//
// > NOTE: This API is only exposed when the D1 service is configured to use the standalone ID provider.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: authn.proto

package authn

import (
	scopes "github.com/cybercryptio/d1-client-go/d1-generic/protobuf/scopes"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// / Represents creation of a new user.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// The scopes the user should be assigned.
	Scopes []scopes.Scope `protobuf:"varint,1,rep,packed,name=scopes,proto3,enum=d1.scopes.Scope" json:"scopes,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetScopes() []scopes.Scope {
	if x != nil {
		return x.Scopes
	}
	return nil
}

// / Represents the result of a user creation.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// ID of the new user.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	/// The users password.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateUserResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// / Represents the user login.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type LoginUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// ID of the user.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	/// The users password.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginUserRequest) Reset() {
	*x = LoginUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserRequest) ProtoMessage() {}

func (x *LoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserRequest.ProtoReflect.Descriptor instead.
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{2}
}

func (x *LoginUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// / Represents the result of a user login.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type LoginUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Access token to be used in subsequent requests.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	/// The expiration time of the access token.
	ExpiryTime int64 `protobuf:"varint,2,opt,name=expiry_time,json=expiryTime,proto3" json:"expiry_time,omitempty"`
}

func (x *LoginUserResponse) Reset() {
	*x = LoginUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserResponse) ProtoMessage() {}

func (x *LoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserResponse.ProtoReflect.Descriptor instead.
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{3}
}

func (x *LoginUserResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginUserResponse) GetExpiryTime() int64 {
	if x != nil {
		return x.ExpiryTime
	}
	return 0
}

// / Represents the removal of a user.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type RemoveUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// ID of the user to remove.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RemoveUserRequest) Reset() {
	*x = RemoveUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserRequest) ProtoMessage() {}

func (x *RemoveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserRequest) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// / Represents the result of a user removal.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type RemoveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserResponse) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{5}
}

// / Represents creation of a new group.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// The scopes the group should be assigned.
	Scopes []scopes.Scope `protobuf:"varint,1,rep,packed,name=scopes,proto3,enum=d1.scopes.Scope" json:"scopes,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{6}
}

func (x *CreateGroupRequest) GetScopes() []scopes.Scope {
	if x != nil {
		return x.Scopes
	}
	return nil
}

// / Represents the result of a group creation.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// ID of the new group.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{7}
}

func (x *CreateGroupResponse) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// / Represents the addition of a user to a group.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type AddUserToGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// ID of the user to add.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	/// ID of the group the user should be added to.
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *AddUserToGroupRequest) Reset() {
	*x = AddUserToGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToGroupRequest) ProtoMessage() {}

func (x *AddUserToGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToGroupRequest.ProtoReflect.Descriptor instead.
func (*AddUserToGroupRequest) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{8}
}

func (x *AddUserToGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddUserToGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// / Represents the result of a user addition to a group.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type AddUserToGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddUserToGroupResponse) Reset() {
	*x = AddUserToGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToGroupResponse) ProtoMessage() {}

func (x *AddUserToGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToGroupResponse.ProtoReflect.Descriptor instead.
func (*AddUserToGroupResponse) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{9}
}

// / Represents the removal of a user from a group.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type RemoveUserFromGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// ID of the user to remove.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	/// ID of the group the user should be removed from.
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *RemoveUserFromGroupRequest) Reset() {
	*x = RemoveUserFromGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromGroupRequest) ProtoMessage() {}

func (x *RemoveUserFromGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromGroupRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserFromGroupRequest) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveUserFromGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveUserFromGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// / Represents the result of a user removal from a group.
// > NOTE: This is only exposed when the D1 service is configured to use the standalone ID provider.
type RemoveUserFromGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveUserFromGroupResponse) Reset() {
	*x = RemoveUserFromGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserFromGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserFromGroupResponse) ProtoMessage() {}

func (x *RemoveUserFromGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authn_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserFromGroupResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserFromGroupResponse) Descriptor() ([]byte, []int) {
	return file_authn_proto_rawDescGZIP(), []int{11}
}

var File_authn_proto protoreflect.FileDescriptor

var file_authn_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64,
	0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x1a, 0x0c, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x64, 0x31, 0x2e,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x06, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x47, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x57, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x64, 0x31,
	0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x06, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50,
	0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x22, 0x1d, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72,
	0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xf0, 0x03, 0x0a, 0x05, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1a, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x64, 0x31, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x13,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x24, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x31, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x8e, 0x01, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x2e, 0x64, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x42, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x79, 0x62, 0x65, 0x72, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x2f, 0x64, 0x31, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0xaa, 0x02, 0x23,
	0x43, 0x79, 0x62, 0x65, 0x72, 0x43, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x44, 0x31, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authn_proto_rawDescOnce sync.Once
	file_authn_proto_rawDescData = file_authn_proto_rawDesc
)

func file_authn_proto_rawDescGZIP() []byte {
	file_authn_proto_rawDescOnce.Do(func() {
		file_authn_proto_rawDescData = protoimpl.X.CompressGZIP(file_authn_proto_rawDescData)
	})
	return file_authn_proto_rawDescData
}

var file_authn_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_authn_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),           // 0: d1.authn.CreateUserRequest
	(*CreateUserResponse)(nil),          // 1: d1.authn.CreateUserResponse
	(*LoginUserRequest)(nil),            // 2: d1.authn.LoginUserRequest
	(*LoginUserResponse)(nil),           // 3: d1.authn.LoginUserResponse
	(*RemoveUserRequest)(nil),           // 4: d1.authn.RemoveUserRequest
	(*RemoveUserResponse)(nil),          // 5: d1.authn.RemoveUserResponse
	(*CreateGroupRequest)(nil),          // 6: d1.authn.CreateGroupRequest
	(*CreateGroupResponse)(nil),         // 7: d1.authn.CreateGroupResponse
	(*AddUserToGroupRequest)(nil),       // 8: d1.authn.AddUserToGroupRequest
	(*AddUserToGroupResponse)(nil),      // 9: d1.authn.AddUserToGroupResponse
	(*RemoveUserFromGroupRequest)(nil),  // 10: d1.authn.RemoveUserFromGroupRequest
	(*RemoveUserFromGroupResponse)(nil), // 11: d1.authn.RemoveUserFromGroupResponse
	(scopes.Scope)(0),                   // 12: d1.scopes.Scope
}
var file_authn_proto_depIdxs = []int32{
	12, // 0: d1.authn.CreateUserRequest.scopes:type_name -> d1.scopes.Scope
	12, // 1: d1.authn.CreateGroupRequest.scopes:type_name -> d1.scopes.Scope
	0,  // 2: d1.authn.Authn.CreateUser:input_type -> d1.authn.CreateUserRequest
	2,  // 3: d1.authn.Authn.LoginUser:input_type -> d1.authn.LoginUserRequest
	4,  // 4: d1.authn.Authn.RemoveUser:input_type -> d1.authn.RemoveUserRequest
	6,  // 5: d1.authn.Authn.CreateGroup:input_type -> d1.authn.CreateGroupRequest
	8,  // 6: d1.authn.Authn.AddUserToGroup:input_type -> d1.authn.AddUserToGroupRequest
	10, // 7: d1.authn.Authn.RemoveUserFromGroup:input_type -> d1.authn.RemoveUserFromGroupRequest
	1,  // 8: d1.authn.Authn.CreateUser:output_type -> d1.authn.CreateUserResponse
	3,  // 9: d1.authn.Authn.LoginUser:output_type -> d1.authn.LoginUserResponse
	5,  // 10: d1.authn.Authn.RemoveUser:output_type -> d1.authn.RemoveUserResponse
	7,  // 11: d1.authn.Authn.CreateGroup:output_type -> d1.authn.CreateGroupResponse
	9,  // 12: d1.authn.Authn.AddUserToGroup:output_type -> d1.authn.AddUserToGroupResponse
	11, // 13: d1.authn.Authn.RemoveUserFromGroup:output_type -> d1.authn.RemoveUserFromGroupResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_authn_proto_init() }
func file_authn_proto_init() {
	if File_authn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_authn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_authn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserRequest); i {
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
		file_authn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserResponse); i {
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
		file_authn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserRequest); i {
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
		file_authn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserResponse); i {
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
		file_authn_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_authn_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_authn_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToGroupRequest); i {
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
		file_authn_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToGroupResponse); i {
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
		file_authn_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromGroupRequest); i {
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
		file_authn_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserFromGroupResponse); i {
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
			RawDescriptor: file_authn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authn_proto_goTypes,
		DependencyIndexes: file_authn_proto_depIdxs,
		MessageInfos:      file_authn_proto_msgTypes,
	}.Build()
	File_authn_proto = out.File
	file_authn_proto_rawDesc = nil
	file_authn_proto_goTypes = nil
	file_authn_proto_depIdxs = nil
}
