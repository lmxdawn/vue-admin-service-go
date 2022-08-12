// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: v1/rbac.proto

package v1

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

type CreateAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CreateAdminRequest) Reset() {
	*x = CreateAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminRequest) ProtoMessage() {}

func (x *CreateAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminRequest.ProtoReflect.Descriptor instead.
func (*CreateAdminRequest) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAdminRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin *CreateAdminReply_Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *CreateAdminReply) Reset() {
	*x = CreateAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminReply) ProtoMessage() {}

func (x *CreateAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminReply.ProtoReflect.Descriptor instead.
func (*CreateAdminReply) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAdminReply) GetAdmin() *CreateAdminReply_Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type UpdateAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAdminRequest) Reset() {
	*x = UpdateAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminRequest) ProtoMessage() {}

func (x *UpdateAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdminRequest) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{2}
}

type UpdateAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAdminReply) Reset() {
	*x = UpdateAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminReply) ProtoMessage() {}

func (x *UpdateAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminReply.ProtoReflect.Descriptor instead.
func (*UpdateAdminReply) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{3}
}

type DeleteAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAdminRequest) Reset() {
	*x = DeleteAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminRequest) ProtoMessage() {}

func (x *DeleteAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminRequest.ProtoReflect.Descriptor instead.
func (*DeleteAdminRequest) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{4}
}

type DeleteAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAdminReply) Reset() {
	*x = DeleteAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminReply) ProtoMessage() {}

func (x *DeleteAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminReply.ProtoReflect.Descriptor instead.
func (*DeleteAdminReply) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{5}
}

type GetAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAdminRequest) Reset() {
	*x = GetAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminRequest) ProtoMessage() {}

func (x *GetAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminRequest.ProtoReflect.Descriptor instead.
func (*GetAdminRequest) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{6}
}

type GetAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAdminReply) Reset() {
	*x = GetAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminReply) ProtoMessage() {}

func (x *GetAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminReply.ProtoReflect.Descriptor instead.
func (*GetAdminReply) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{7}
}

type ListAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAdminRequest) Reset() {
	*x = ListAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdminRequest) ProtoMessage() {}

func (x *ListAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdminRequest.ProtoReflect.Descriptor instead.
func (*ListAdminRequest) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{8}
}

type ListAdminReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAdminReply) Reset() {
	*x = ListAdminReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAdminReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdminReply) ProtoMessage() {}

func (x *ListAdminReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdminReply.ProtoReflect.Descriptor instead.
func (*ListAdminReply) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{9}
}

type CreateAdminReply_Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CreateAdminReply_Admin) Reset() {
	*x = CreateAdminReply_Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rbac_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminReply_Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminReply_Admin) ProtoMessage() {}

func (x *CreateAdminReply_Admin) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rbac_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminReply_Admin.ProtoReflect.Descriptor instead.
func (*CreateAdminReply_Admin) Descriptor() ([]byte, []int) {
	return file_v1_rbac_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreateAdminReply_Admin) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateAdminReply_Admin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_v1_rbac_proto protoreflect.FileDescriptor

var file_v1_rbac_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x30, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x33,
	0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0xd2, 0x03, 0x0a, 0x04, 0x52, 0x62, 0x61, 0x63, 0x12, 0x5d, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5d, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62,
	0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x57, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x22, 0x5a, 0x20, 0x76, 0x75, 0x65, 0x2d, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_rbac_proto_rawDescOnce sync.Once
	file_v1_rbac_proto_rawDescData = file_v1_rbac_proto_rawDesc
)

func file_v1_rbac_proto_rawDescGZIP() []byte {
	file_v1_rbac_proto_rawDescOnce.Do(func() {
		file_v1_rbac_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_rbac_proto_rawDescData)
	})
	return file_v1_rbac_proto_rawDescData
}

var file_v1_rbac_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_v1_rbac_proto_goTypes = []interface{}{
	(*CreateAdminRequest)(nil),     // 0: api.rbac.service.v1.CreateAdminRequest
	(*CreateAdminReply)(nil),       // 1: api.rbac.service.v1.CreateAdminReply
	(*UpdateAdminRequest)(nil),     // 2: api.rbac.service.v1.UpdateAdminRequest
	(*UpdateAdminReply)(nil),       // 3: api.rbac.service.v1.UpdateAdminReply
	(*DeleteAdminRequest)(nil),     // 4: api.rbac.service.v1.DeleteAdminRequest
	(*DeleteAdminReply)(nil),       // 5: api.rbac.service.v1.DeleteAdminReply
	(*GetAdminRequest)(nil),        // 6: api.rbac.service.v1.GetAdminRequest
	(*GetAdminReply)(nil),          // 7: api.rbac.service.v1.GetAdminReply
	(*ListAdminRequest)(nil),       // 8: api.rbac.service.v1.ListAdminRequest
	(*ListAdminReply)(nil),         // 9: api.rbac.service.v1.ListAdminReply
	(*CreateAdminReply_Admin)(nil), // 10: api.rbac.service.v1.CreateAdminReply.Admin
}
var file_v1_rbac_proto_depIdxs = []int32{
	10, // 0: api.rbac.service.v1.CreateAdminReply.admin:type_name -> api.rbac.service.v1.CreateAdminReply.Admin
	0,  // 1: api.rbac.service.v1.Rbac.CreateAdmin:input_type -> api.rbac.service.v1.CreateAdminRequest
	2,  // 2: api.rbac.service.v1.Rbac.UpdateAdmin:input_type -> api.rbac.service.v1.UpdateAdminRequest
	4,  // 3: api.rbac.service.v1.Rbac.DeleteAdmin:input_type -> api.rbac.service.v1.DeleteAdminRequest
	6,  // 4: api.rbac.service.v1.Rbac.GetAdmin:input_type -> api.rbac.service.v1.GetAdminRequest
	8,  // 5: api.rbac.service.v1.Rbac.ListAdmin:input_type -> api.rbac.service.v1.ListAdminRequest
	1,  // 6: api.rbac.service.v1.Rbac.CreateAdmin:output_type -> api.rbac.service.v1.CreateAdminReply
	3,  // 7: api.rbac.service.v1.Rbac.UpdateAdmin:output_type -> api.rbac.service.v1.UpdateAdminReply
	5,  // 8: api.rbac.service.v1.Rbac.DeleteAdmin:output_type -> api.rbac.service.v1.DeleteAdminReply
	7,  // 9: api.rbac.service.v1.Rbac.GetAdmin:output_type -> api.rbac.service.v1.GetAdminReply
	9,  // 10: api.rbac.service.v1.Rbac.ListAdmin:output_type -> api.rbac.service.v1.ListAdminReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_v1_rbac_proto_init() }
func file_v1_rbac_proto_init() {
	if File_v1_rbac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_rbac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdminRequest); i {
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
		file_v1_rbac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdminReply); i {
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
		file_v1_rbac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminRequest); i {
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
		file_v1_rbac_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminReply); i {
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
		file_v1_rbac_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdminRequest); i {
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
		file_v1_rbac_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdminReply); i {
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
		file_v1_rbac_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminRequest); i {
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
		file_v1_rbac_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminReply); i {
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
		file_v1_rbac_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAdminRequest); i {
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
		file_v1_rbac_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAdminReply); i {
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
		file_v1_rbac_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdminReply_Admin); i {
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
			RawDescriptor: file_v1_rbac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_rbac_proto_goTypes,
		DependencyIndexes: file_v1_rbac_proto_depIdxs,
		MessageInfos:      file_v1_rbac_proto_msgTypes,
	}.Build()
	File_v1_rbac_proto = out.File
	file_v1_rbac_proto_rawDesc = nil
	file_v1_rbac_proto_goTypes = nil
	file_v1_rbac_proto_depIdxs = nil
}