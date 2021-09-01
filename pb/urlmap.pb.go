// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: urlmap.proto

package urlmap

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type RedirectPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	NotifyTo string `protobuf:"bytes,2,opt,name=notify_to,json=notifyTo,proto3" json:"notify_to,omitempty"`
}

func (x *RedirectPath) Reset() {
	*x = RedirectPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlmap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectPath) ProtoMessage() {}

func (x *RedirectPath) ProtoReflect() protoreflect.Message {
	mi := &file_urlmap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectPath.ProtoReflect.Descriptor instead.
func (*RedirectPath) Descriptor() ([]byte, []int) {
	return file_urlmap_proto_rawDescGZIP(), []int{0}
}

func (x *RedirectPath) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RedirectPath) GetNotifyTo() string {
	if x != nil {
		return x.NotifyTo
	}
	return ""
}

type OrgUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Org      string `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
	NotifyTo string `protobuf:"bytes,2,opt,name=notify_to,json=notifyTo,proto3" json:"notify_to,omitempty"`
}

func (x *OrgUrl) Reset() {
	*x = OrgUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlmap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgUrl) ProtoMessage() {}

func (x *OrgUrl) ProtoReflect() protoreflect.Message {
	mi := &file_urlmap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgUrl.ProtoReflect.Descriptor instead.
func (*OrgUrl) Descriptor() ([]byte, []int) {
	return file_urlmap_proto_rawDescGZIP(), []int{1}
}

func (x *OrgUrl) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

func (x *OrgUrl) GetNotifyTo() string {
	if x != nil {
		return x.NotifyTo
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	NotifyTo string `protobuf:"bytes,2,opt,name=notify_to,json=notifyTo,proto3" json:"notify_to,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlmap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_urlmap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_urlmap_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *User) GetNotifyTo() string {
	if x != nil {
		return x.NotifyTo
	}
	return ""
}

type ArrayRedirectData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Redirects []*RedirectData `protobuf:"bytes,1,rep,name=redirects,proto3" json:"redirects,omitempty"`
}

func (x *ArrayRedirectData) Reset() {
	*x = ArrayRedirectData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlmap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayRedirectData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayRedirectData) ProtoMessage() {}

func (x *ArrayRedirectData) ProtoReflect() protoreflect.Message {
	mi := &file_urlmap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayRedirectData.ProtoReflect.Descriptor instead.
func (*ArrayRedirectData) Descriptor() ([]byte, []int) {
	return file_urlmap_proto_rawDescGZIP(), []int{3}
}

func (x *ArrayRedirectData) GetRedirects() []*RedirectData {
	if x != nil {
		return x.Redirects
	}
	return nil
}

type RedirectData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Redirect *RedirectInfo `protobuf:"bytes,1,opt,name=redirect,proto3" json:"redirect,omitempty"`
}

func (x *RedirectData) Reset() {
	*x = RedirectData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlmap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectData) ProtoMessage() {}

func (x *RedirectData) ProtoReflect() protoreflect.Message {
	mi := &file_urlmap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectData.ProtoReflect.Descriptor instead.
func (*RedirectData) Descriptor() ([]byte, []int) {
	return file_urlmap_proto_rawDescGZIP(), []int{4}
}

func (x *RedirectData) GetRedirect() *RedirectInfo {
	if x != nil {
		return x.Redirect
	}
	return nil
}

type RedirectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	RedirectPath string `protobuf:"bytes,2,opt,name=redirectPath,proto3" json:"redirectPath,omitempty"`
	Org          string `protobuf:"bytes,3,opt,name=org,proto3" json:"org,omitempty"`
	Host         string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Comment      string `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	Active       int32  `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *RedirectInfo) Reset() {
	*x = RedirectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlmap_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectInfo) ProtoMessage() {}

func (x *RedirectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_urlmap_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectInfo.ProtoReflect.Descriptor instead.
func (*RedirectInfo) Descriptor() ([]byte, []int) {
	return file_urlmap_proto_rawDescGZIP(), []int{5}
}

func (x *RedirectInfo) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RedirectInfo) GetRedirectPath() string {
	if x != nil {
		return x.RedirectPath
	}
	return ""
}

func (x *RedirectInfo) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

func (x *RedirectInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RedirectInfo) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *RedirectInfo) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

type RedirectData_ValidDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Begin string `protobuf:"bytes,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *RedirectData_ValidDate) Reset() {
	*x = RedirectData_ValidDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlmap_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectData_ValidDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectData_ValidDate) ProtoMessage() {}

func (x *RedirectData_ValidDate) ProtoReflect() protoreflect.Message {
	mi := &file_urlmap_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectData_ValidDate.ProtoReflect.Descriptor instead.
func (*RedirectData_ValidDate) Descriptor() ([]byte, []int) {
	return file_urlmap_proto_rawDescGZIP(), []int{4, 0}
}

func (x *RedirectData_ValidDate) GetBegin() string {
	if x != nil {
		return x.Begin
	}
	return ""
}

func (x *RedirectData_ValidDate) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

var File_urlmap_proto protoreflect.FileDescriptor

var file_urlmap_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0c, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x54, 0x6f, 0x22, 0x37, 0x0a, 0x06, 0x4f, 0x72, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x22, 0x37, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x22, 0x47, 0x0a, 0x11, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x09, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x75, 0x0a, 0x0c, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x30, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x1a, 0x33, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x32, 0x89, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x67, 0x42, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x2e, 0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70,
	0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x0e, 0x2e,
	0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x4f, 0x72, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x38, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c,
	0x2e, 0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x75,
	0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x2e, 0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x75, 0x72, 0x6c, 0x6d, 0x61,
	0x70, 0x2e, 0x4f, 0x72, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x0c, 0x2e, 0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x32, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x2e,
	0x75, 0x72, 0x6c, 0x6d, 0x61, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_urlmap_proto_rawDescOnce sync.Once
	file_urlmap_proto_rawDescData = file_urlmap_proto_rawDesc
)

func file_urlmap_proto_rawDescGZIP() []byte {
	file_urlmap_proto_rawDescOnce.Do(func() {
		file_urlmap_proto_rawDescData = protoimpl.X.CompressGZIP(file_urlmap_proto_rawDescData)
	})
	return file_urlmap_proto_rawDescData
}

var file_urlmap_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_urlmap_proto_goTypes = []interface{}{
	(*RedirectPath)(nil),           // 0: urlmap.RedirectPath
	(*OrgUrl)(nil),                 // 1: urlmap.OrgUrl
	(*User)(nil),                   // 2: urlmap.User
	(*ArrayRedirectData)(nil),      // 3: urlmap.ArrayRedirectData
	(*RedirectData)(nil),           // 4: urlmap.RedirectData
	(*RedirectInfo)(nil),           // 5: urlmap.RedirectInfo
	(*RedirectData_ValidDate)(nil), // 6: urlmap.RedirectData.ValidDate
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_urlmap_proto_depIdxs = []int32{
	4, // 0: urlmap.ArrayRedirectData.redirects:type_name -> urlmap.RedirectData
	5, // 1: urlmap.RedirectData.redirect:type_name -> urlmap.RedirectInfo
	0, // 2: urlmap.Redirection.GetOrgByPath:input_type -> urlmap.RedirectPath
	2, // 3: urlmap.Redirection.GetInfoByUser:input_type -> urlmap.User
	4, // 4: urlmap.Redirection.SetInfo:input_type -> urlmap.RedirectData
	2, // 5: urlmap.Redirection.SetUser:input_type -> urlmap.User
	2, // 6: urlmap.Redirection.RemoveUser:input_type -> urlmap.User
	1, // 7: urlmap.Redirection.GetOrgByPath:output_type -> urlmap.OrgUrl
	3, // 8: urlmap.Redirection.GetInfoByUser:output_type -> urlmap.ArrayRedirectData
	1, // 9: urlmap.Redirection.SetInfo:output_type -> urlmap.OrgUrl
	2, // 10: urlmap.Redirection.SetUser:output_type -> urlmap.User
	7, // 11: urlmap.Redirection.RemoveUser:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_urlmap_proto_init() }
func file_urlmap_proto_init() {
	if File_urlmap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_urlmap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectPath); i {
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
		file_urlmap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgUrl); i {
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
		file_urlmap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_urlmap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayRedirectData); i {
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
		file_urlmap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectData); i {
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
		file_urlmap_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectInfo); i {
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
		file_urlmap_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectData_ValidDate); i {
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
			RawDescriptor: file_urlmap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_urlmap_proto_goTypes,
		DependencyIndexes: file_urlmap_proto_depIdxs,
		MessageInfos:      file_urlmap_proto_msgTypes,
	}.Build()
	File_urlmap_proto = out.File
	file_urlmap_proto_rawDesc = nil
	file_urlmap_proto_goTypes = nil
	file_urlmap_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RedirectionClient is the client API for Redirection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RedirectionClient interface {
	GetOrgByPath(ctx context.Context, in *RedirectPath, opts ...grpc.CallOption) (*OrgUrl, error)
	GetInfoByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ArrayRedirectData, error)
	SetInfo(ctx context.Context, in *RedirectData, opts ...grpc.CallOption) (*OrgUrl, error)
	SetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	RemoveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type redirectionClient struct {
	cc grpc.ClientConnInterface
}

func NewRedirectionClient(cc grpc.ClientConnInterface) RedirectionClient {
	return &redirectionClient{cc}
}

func (c *redirectionClient) GetOrgByPath(ctx context.Context, in *RedirectPath, opts ...grpc.CallOption) (*OrgUrl, error) {
	out := new(OrgUrl)
	err := c.cc.Invoke(ctx, "/urlmap.Redirection/GetOrgByPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redirectionClient) GetInfoByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ArrayRedirectData, error) {
	out := new(ArrayRedirectData)
	err := c.cc.Invoke(ctx, "/urlmap.Redirection/GetInfoByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redirectionClient) SetInfo(ctx context.Context, in *RedirectData, opts ...grpc.CallOption) (*OrgUrl, error) {
	out := new(OrgUrl)
	err := c.cc.Invoke(ctx, "/urlmap.Redirection/SetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redirectionClient) SetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/urlmap.Redirection/SetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redirectionClient) RemoveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/urlmap.Redirection/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedirectionServer is the server API for Redirection service.
type RedirectionServer interface {
	GetOrgByPath(context.Context, *RedirectPath) (*OrgUrl, error)
	GetInfoByUser(context.Context, *User) (*ArrayRedirectData, error)
	SetInfo(context.Context, *RedirectData) (*OrgUrl, error)
	SetUser(context.Context, *User) (*User, error)
	RemoveUser(context.Context, *User) (*emptypb.Empty, error)
}

// UnimplementedRedirectionServer can be embedded to have forward compatible implementations.
type UnimplementedRedirectionServer struct {
}

func (*UnimplementedRedirectionServer) GetOrgByPath(context.Context, *RedirectPath) (*OrgUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgByPath not implemented")
}
func (*UnimplementedRedirectionServer) GetInfoByUser(context.Context, *User) (*ArrayRedirectData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByUser not implemented")
}
func (*UnimplementedRedirectionServer) SetInfo(context.Context, *RedirectData) (*OrgUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInfo not implemented")
}
func (*UnimplementedRedirectionServer) SetUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUser not implemented")
}
func (*UnimplementedRedirectionServer) RemoveUser(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}

func RegisterRedirectionServer(s *grpc.Server, srv RedirectionServer) {
	s.RegisterService(&_Redirection_serviceDesc, srv)
}

func _Redirection_GetOrgByPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedirectPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedirectionServer).GetOrgByPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlmap.Redirection/GetOrgByPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedirectionServer).GetOrgByPath(ctx, req.(*RedirectPath))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redirection_GetInfoByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedirectionServer).GetInfoByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlmap.Redirection/GetInfoByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedirectionServer).GetInfoByUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redirection_SetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedirectData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedirectionServer).SetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlmap.Redirection/SetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedirectionServer).SetInfo(ctx, req.(*RedirectData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redirection_SetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedirectionServer).SetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlmap.Redirection/SetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedirectionServer).SetUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Redirection_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedirectionServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urlmap.Redirection/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedirectionServer).RemoveUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Redirection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "urlmap.Redirection",
	HandlerType: (*RedirectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrgByPath",
			Handler:    _Redirection_GetOrgByPath_Handler,
		},
		{
			MethodName: "GetInfoByUser",
			Handler:    _Redirection_GetInfoByUser_Handler,
		},
		{
			MethodName: "SetInfo",
			Handler:    _Redirection_SetInfo_Handler,
		},
		{
			MethodName: "SetUser",
			Handler:    _Redirection_SetUser_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _Redirection_RemoveUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "urlmap.proto",
}
