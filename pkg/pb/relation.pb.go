// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: relation.proto

package pb

import (
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

// 关注关系方法
type RelationActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelfUserId   int64 `protobuf:"varint,1,opt,name=self_user_id,json=selfUserId,proto3" json:"self_user_id,omitempty"`
	TargetUserId int64 `protobuf:"varint,2,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	ActionType   int64 `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
}

func (x *RelationActionReq) Reset() {
	*x = RelationActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionReq) ProtoMessage() {}

func (x *RelationActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionReq.ProtoReflect.Descriptor instead.
func (*RelationActionReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{0}
}

func (x *RelationActionReq) GetSelfUserId() int64 {
	if x != nil {
		return x.SelfUserId
	}
	return 0
}

func (x *RelationActionReq) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

func (x *RelationActionReq) GetActionType() int64 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type RelationActionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResponse *CommonResponse `protobuf:"bytes,1,opt,name=common_response,json=commonResponse,proto3" json:"common_response,omitempty"`
}

func (x *RelationActionResp) Reset() {
	*x = RelationActionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionResp) ProtoMessage() {}

func (x *RelationActionResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionResp.ProtoReflect.Descriptor instead.
func (*RelationActionResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{1}
}

func (x *RelationActionResp) GetCommonResponse() *CommonResponse {
	if x != nil {
		return x.CommonResponse
	}
	return nil
}

// 获取我的关注列表
type GetFollowListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFollowListReq) Reset() {
	*x = GetFollowListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowListReq) ProtoMessage() {}

func (x *GetFollowListReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowListReq.ProtoReflect.Descriptor instead.
func (*GetFollowListReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{2}
}

func (x *GetFollowListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFollowListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []int64 `protobuf:"varint,1,rep,packed,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *GetFollowListResp) Reset() {
	*x = GetFollowListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowListResp) ProtoMessage() {}

func (x *GetFollowListResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowListResp.ProtoReflect.Descriptor instead.
func (*GetFollowListResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{3}
}

func (x *GetFollowListResp) GetUserList() []int64 {
	if x != nil {
		return x.UserList
	}
	return nil
}

// 获取我的粉丝列表
type GetFansListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFansListReq) Reset() {
	*x = GetFansListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFansListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFansListReq) ProtoMessage() {}

func (x *GetFansListReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFansListReq.ProtoReflect.Descriptor instead.
func (*GetFansListReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{4}
}

func (x *GetFansListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFansListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerList []int64 `protobuf:"varint,1,rep,packed,name=follower_list,json=followerList,proto3" json:"follower_list,omitempty"`
}

func (x *GetFansListResp) Reset() {
	*x = GetFansListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFansListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFansListResp) ProtoMessage() {}

func (x *GetFansListResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFansListResp.ProtoReflect.Descriptor instead.
func (*GetFansListResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{5}
}

func (x *GetFansListResp) GetFollowerList() []int64 {
	if x != nil {
		return x.FollowerList
	}
	return nil
}

type FollowUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelfUserId   int64 `protobuf:"varint,1,opt,name=self_user_id,json=selfUserId,proto3" json:"self_user_id,omitempty"`
	TargetUserId int64 `protobuf:"varint,2,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
}

func (x *FollowUnit) Reset() {
	*x = FollowUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowUnit) ProtoMessage() {}

func (x *FollowUnit) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowUnit.ProtoReflect.Descriptor instead.
func (*FollowUnit) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{6}
}

func (x *FollowUnit) GetSelfUserId() int64 {
	if x != nil {
		return x.SelfUserId
	}
	return 0
}

func (x *FollowUnit) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

type IsFollowDictReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowUnit []*FollowUnit `protobuf:"bytes,1,rep,name=follow_unit,json=followUnit,proto3" json:"follow_unit,omitempty"`
}

func (x *IsFollowDictReq) Reset() {
	*x = IsFollowDictReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFollowDictReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFollowDictReq) ProtoMessage() {}

func (x *IsFollowDictReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFollowDictReq.ProtoReflect.Descriptor instead.
func (*IsFollowDictReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{7}
}

func (x *IsFollowDictReq) GetFollowUnit() []*FollowUnit {
	if x != nil {
		return x.FollowUnit
	}
	return nil
}

type IsFollowDictResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFollow map[string]bool `protobuf:"bytes,1,rep,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *IsFollowDictResp) Reset() {
	*x = IsFollowDictResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFollowDictResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFollowDictResp) ProtoMessage() {}

func (x *IsFollowDictResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFollowDictResp.ProtoReflect.Descriptor instead.
func (*IsFollowDictResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{8}
}

func (x *IsFollowDictResp) GetIsFollow() map[string]bool {
	if x != nil {
		return x.IsFollow
	}
	return nil
}

var File_relation_proto protoreflect.FileDescriptor

var file_relation_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c,
	0x0a, 0x11, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6c, 0x66, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4e, 0x0a, 0x12,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x38, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x54,
	0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0c,
	0x73, 0x65, 0x6c, 0x66, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6c, 0x66, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x2e, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x1a, 0x3b, 0x0a, 0x0d, 0x49, 0x73, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xe7, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f,
	0x0a, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x10, 0x2e, 0x49, 0x73, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x49,
	0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_relation_proto_rawDescOnce sync.Once
	file_relation_proto_rawDescData = file_relation_proto_rawDesc
)

func file_relation_proto_rawDescGZIP() []byte {
	file_relation_proto_rawDescOnce.Do(func() {
		file_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_relation_proto_rawDescData)
	})
	return file_relation_proto_rawDescData
}

var file_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_relation_proto_goTypes = []interface{}{
	(*RelationActionReq)(nil),  // 0: RelationActionReq
	(*RelationActionResp)(nil), // 1: RelationActionResp
	(*GetFollowListReq)(nil),   // 2: GetFollowListReq
	(*GetFollowListResp)(nil),  // 3: GetFollowListResp
	(*GetFansListReq)(nil),     // 4: GetFansListReq
	(*GetFansListResp)(nil),    // 5: GetFansListResp
	(*FollowUnit)(nil),         // 6: FollowUnit
	(*IsFollowDictReq)(nil),    // 7: IsFollowDictReq
	(*IsFollowDictResp)(nil),   // 8: IsFollowDictResp
	nil,                        // 9: IsFollowDictResp.IsFollowEntry
	(*CommonResponse)(nil),     // 10: CommonResponse
}
var file_relation_proto_depIdxs = []int32{
	10, // 0: RelationActionResp.common_response:type_name -> CommonResponse
	6,  // 1: IsFollowDictReq.follow_unit:type_name -> FollowUnit
	9,  // 2: IsFollowDictResp.is_follow:type_name -> IsFollowDictResp.IsFollowEntry
	0,  // 3: RelationService.RelationAction:input_type -> RelationActionReq
	2,  // 4: RelationService.GetFollowList:input_type -> GetFollowListReq
	4,  // 5: RelationService.GetFansList:input_type -> GetFansListReq
	7,  // 6: RelationService.IsFollow:input_type -> IsFollowDictReq
	1,  // 7: RelationService.RelationAction:output_type -> RelationActionResp
	3,  // 8: RelationService.GetFollowList:output_type -> GetFollowListResp
	5,  // 9: RelationService.GetFansList:output_type -> GetFansListResp
	8,  // 10: RelationService.IsFollow:output_type -> IsFollowDictResp
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_relation_proto_init() }
func file_relation_proto_init() {
	if File_relation_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionReq); i {
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
		file_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionResp); i {
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
		file_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowListReq); i {
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
		file_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowListResp); i {
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
		file_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFansListReq); i {
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
		file_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFansListResp); i {
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
		file_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowUnit); i {
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
		file_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFollowDictReq); i {
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
		file_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFollowDictResp); i {
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
			RawDescriptor: file_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relation_proto_goTypes,
		DependencyIndexes: file_relation_proto_depIdxs,
		MessageInfos:      file_relation_proto_msgTypes,
	}.Build()
	File_relation_proto = out.File
	file_relation_proto_rawDesc = nil
	file_relation_proto_goTypes = nil
	file_relation_proto_depIdxs = nil
}
