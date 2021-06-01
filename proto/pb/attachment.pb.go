// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: attachment.proto

package pb

import (
	empb "github.com/Etpmls/Etpmls-Micro/v3/proto/empb"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AttachmentGetOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service   string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	OwnerId   uint32 `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerType string `protobuf:"bytes,3,opt,name=owner_type,json=ownerType,proto3" json:"owner_type,omitempty"`
}

func (x *AttachmentGetOne) Reset() {
	*x = AttachmentGetOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentGetOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentGetOne) ProtoMessage() {}

func (x *AttachmentGetOne) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentGetOne.ProtoReflect.Descriptor instead.
func (*AttachmentGetOne) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{0}
}

func (x *AttachmentGetOne) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AttachmentGetOne) GetOwnerId() uint32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *AttachmentGetOne) GetOwnerType() string {
	if x != nil {
		return x.OwnerType
	}
	return ""
}

type AttachmentGetMany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service   string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	OwnerIds  []uint32 `protobuf:"varint,2,rep,packed,name=owner_ids,json=ownerIds,proto3" json:"owner_ids,omitempty"`
	OwnerType string   `protobuf:"bytes,3,opt,name=owner_type,json=ownerType,proto3" json:"owner_type,omitempty"`
}

func (x *AttachmentGetMany) Reset() {
	*x = AttachmentGetMany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentGetMany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentGetMany) ProtoMessage() {}

func (x *AttachmentGetMany) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentGetMany.ProtoReflect.Descriptor instead.
func (*AttachmentGetMany) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{1}
}

func (x *AttachmentGetMany) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AttachmentGetMany) GetOwnerIds() []uint32 {
	if x != nil {
		return x.OwnerIds
	}
	return nil
}

func (x *AttachmentGetMany) GetOwnerType() string {
	if x != nil {
		return x.OwnerType
	}
	return ""
}

type AttachmentCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service       string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	StorageMethod string `protobuf:"bytes,2,opt,name=storage_method,json=storageMethod,proto3" json:"storage_method,omitempty"`
	Path          string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	OwnerId       uint32 `protobuf:"varint,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerType     string `protobuf:"bytes,5,opt,name=owner_type,json=ownerType,proto3" json:"owner_type,omitempty"`
}

func (x *AttachmentCreate) Reset() {
	*x = AttachmentCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentCreate) ProtoMessage() {}

func (x *AttachmentCreate) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentCreate.ProtoReflect.Descriptor instead.
func (*AttachmentCreate) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{2}
}

func (x *AttachmentCreate) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AttachmentCreate) GetStorageMethod() string {
	if x != nil {
		return x.StorageMethod
	}
	return ""
}

func (x *AttachmentCreate) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AttachmentCreate) GetOwnerId() uint32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *AttachmentCreate) GetOwnerType() string {
	if x != nil {
		return x.OwnerType
	}
	return ""
}

type AttachmentCreateMany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service       string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	StorageMethod string   `protobuf:"bytes,2,opt,name=storage_method,json=storageMethod,proto3" json:"storage_method,omitempty"`
	Paths         []string `protobuf:"bytes,3,rep,name=paths,proto3" json:"paths,omitempty"`
	OwnerId       uint32   `protobuf:"varint,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerType     string   `protobuf:"bytes,5,opt,name=owner_type,json=ownerType,proto3" json:"owner_type,omitempty"`
}

func (x *AttachmentCreateMany) Reset() {
	*x = AttachmentCreateMany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentCreateMany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentCreateMany) ProtoMessage() {}

func (x *AttachmentCreateMany) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentCreateMany.ProtoReflect.Descriptor instead.
func (*AttachmentCreateMany) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{3}
}

func (x *AttachmentCreateMany) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AttachmentCreateMany) GetStorageMethod() string {
	if x != nil {
		return x.StorageMethod
	}
	return ""
}

func (x *AttachmentCreateMany) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *AttachmentCreateMany) GetOwnerId() uint32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *AttachmentCreateMany) GetOwnerType() string {
	if x != nil {
		return x.OwnerType
	}
	return ""
}

type AttachmentAppend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service       string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	StorageMethod string   `protobuf:"bytes,2,opt,name=storage_method,json=storageMethod,proto3" json:"storage_method,omitempty"`
	Paths         []string `protobuf:"bytes,3,rep,name=paths,proto3" json:"paths,omitempty"`
	OwnerId       uint32   `protobuf:"varint,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerType     string   `protobuf:"bytes,5,opt,name=owner_type,json=ownerType,proto3" json:"owner_type,omitempty"`
}

func (x *AttachmentAppend) Reset() {
	*x = AttachmentAppend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentAppend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentAppend) ProtoMessage() {}

func (x *AttachmentAppend) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentAppend.ProtoReflect.Descriptor instead.
func (*AttachmentAppend) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{4}
}

func (x *AttachmentAppend) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AttachmentAppend) GetStorageMethod() string {
	if x != nil {
		return x.StorageMethod
	}
	return ""
}

func (x *AttachmentAppend) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *AttachmentAppend) GetOwnerId() uint32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *AttachmentAppend) GetOwnerType() string {
	if x != nil {
		return x.OwnerType
	}
	return ""
}

type AttachmentDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service   string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	OwnerIds  []uint32 `protobuf:"varint,2,rep,packed,name=owner_ids,json=ownerIds,proto3" json:"owner_ids,omitempty"`
	OwnerType string   `protobuf:"bytes,3,opt,name=owner_type,json=ownerType,proto3" json:"owner_type,omitempty"`
}

func (x *AttachmentDelete) Reset() {
	*x = AttachmentDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentDelete) ProtoMessage() {}

func (x *AttachmentDelete) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentDelete.ProtoReflect.Descriptor instead.
func (*AttachmentDelete) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{5}
}

func (x *AttachmentDelete) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AttachmentDelete) GetOwnerIds() []uint32 {
	if x != nil {
		return x.OwnerIds
	}
	return nil
}

func (x *AttachmentDelete) GetOwnerType() string {
	if x != nil {
		return x.OwnerType
	}
	return ""
}

type AttachmentDiskCleanUp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *AttachmentDiskCleanUp) Reset() {
	*x = AttachmentDiskCleanUp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentDiskCleanUp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentDiskCleanUp) ProtoMessage() {}

func (x *AttachmentDiskCleanUp) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentDiskCleanUp.ProtoReflect.Descriptor instead.
func (*AttachmentDiskCleanUp) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{6}
}

func (x *AttachmentDiskCleanUp) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

var File_attachment_proto protoreflect.FileDescriptor

var file_attachment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x66, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x69, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x14, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x22, 0xa3, 0x01, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x68, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x31, 0x0a, 0x15, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x69, 0x73, 0x6b, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x55, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x32, 0xa8, 0x03, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x1c, 0x2e,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x1a, 0x0c, 0x2e, 0x65, 0x6d,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x79, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e,
	0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x20, 0x2e, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x1a, 0x0c, 0x2e,
	0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x06, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x0c, 0x2e,
	0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x0b, 0x44, 0x69, 0x73, 0x6b, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x55, 0x70, 0x12, 0x21, 0x2e, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x55, 0x70, 0x1a,
	0x0c, 0x2e, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attachment_proto_rawDescOnce sync.Once
	file_attachment_proto_rawDescData = file_attachment_proto_rawDesc
)

func file_attachment_proto_rawDescGZIP() []byte {
	file_attachment_proto_rawDescOnce.Do(func() {
		file_attachment_proto_rawDescData = protoimpl.X.CompressGZIP(file_attachment_proto_rawDescData)
	})
	return file_attachment_proto_rawDescData
}

var file_attachment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_attachment_proto_goTypes = []interface{}{
	(*AttachmentGetOne)(nil),      // 0: attachment.AttachmentGetOne
	(*AttachmentGetMany)(nil),     // 1: attachment.AttachmentGetMany
	(*AttachmentCreate)(nil),      // 2: attachment.AttachmentCreate
	(*AttachmentCreateMany)(nil),  // 3: attachment.AttachmentCreateMany
	(*AttachmentAppend)(nil),      // 4: attachment.AttachmentAppend
	(*AttachmentDelete)(nil),      // 5: attachment.AttachmentDelete
	(*AttachmentDiskCleanUp)(nil), // 6: attachment.AttachmentDiskCleanUp
	(*empb.Response)(nil),         // 7: em.Response
}
var file_attachment_proto_depIdxs = []int32{
	0, // 0: attachment.Attachment.GetOne:input_type -> attachment.AttachmentGetOne
	1, // 1: attachment.Attachment.GetMany:input_type -> attachment.AttachmentGetMany
	2, // 2: attachment.Attachment.Create:input_type -> attachment.AttachmentCreate
	3, // 3: attachment.Attachment.CreateMany:input_type -> attachment.AttachmentCreateMany
	4, // 4: attachment.Attachment.Append:input_type -> attachment.AttachmentAppend
	5, // 5: attachment.Attachment.Delete:input_type -> attachment.AttachmentDelete
	6, // 6: attachment.Attachment.DiskCleanUp:input_type -> attachment.AttachmentDiskCleanUp
	7, // 7: attachment.Attachment.GetOne:output_type -> em.Response
	7, // 8: attachment.Attachment.GetMany:output_type -> em.Response
	7, // 9: attachment.Attachment.Create:output_type -> em.Response
	7, // 10: attachment.Attachment.CreateMany:output_type -> em.Response
	7, // 11: attachment.Attachment.Append:output_type -> em.Response
	7, // 12: attachment.Attachment.Delete:output_type -> em.Response
	7, // 13: attachment.Attachment.DiskCleanUp:output_type -> em.Response
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_attachment_proto_init() }
func file_attachment_proto_init() {
	if File_attachment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attachment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentGetOne); i {
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
		file_attachment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentGetMany); i {
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
		file_attachment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentCreate); i {
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
		file_attachment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentCreateMany); i {
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
		file_attachment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentAppend); i {
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
		file_attachment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentDelete); i {
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
		file_attachment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentDiskCleanUp); i {
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
			RawDescriptor: file_attachment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attachment_proto_goTypes,
		DependencyIndexes: file_attachment_proto_depIdxs,
		MessageInfos:      file_attachment_proto_msgTypes,
	}.Build()
	File_attachment_proto = out.File
	file_attachment_proto_rawDesc = nil
	file_attachment_proto_goTypes = nil
	file_attachment_proto_depIdxs = nil
}