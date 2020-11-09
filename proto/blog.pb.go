// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: proto/blog.proto

package blogpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Blog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Blog) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// create, read and update will return a blog message
type CreateBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"` // with blank id , that is to be filled by MongoBD logic
}

func (x *CreateBlogReq) Reset() {
	*x = CreateBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogReq) ProtoMessage() {}

func (x *CreateBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogReq.ProtoReflect.Descriptor instead.
func (*CreateBlogReq) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBlogReq) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type CreateBlogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"` // id will get filled
}

func (x *CreateBlogRes) Reset() {
	*x = CreateBlogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogRes) ProtoMessage() {}

func (x *CreateBlogRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogRes.ProtoReflect.Descriptor instead.
func (*CreateBlogRes) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBlogRes) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

// blog will be searched using an id
type ReadBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadBlogReq) Reset() {
	*x = ReadBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlogReq) ProtoMessage() {}

func (x *ReadBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlogReq.ProtoReflect.Descriptor instead.
func (*ReadBlogReq) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{3}
}

func (x *ReadBlogReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadBlogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *ReadBlogRes) Reset() {
	*x = ReadBlogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlogRes) ProtoMessage() {}

func (x *ReadBlogRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlogRes.ProtoReflect.Descriptor instead.
func (*ReadBlogRes) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{4}
}

func (x *ReadBlogRes) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

// same as create, but with id filled already
type UpdateBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *UpdateBlogReq) Reset() {
	*x = UpdateBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogReq) ProtoMessage() {}

func (x *UpdateBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogReq.ProtoReflect.Descriptor instead.
func (*UpdateBlogReq) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBlogReq) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type UpdateBlogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *UpdateBlogRes) Reset() {
	*x = UpdateBlogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogRes) ProtoMessage() {}

func (x *UpdateBlogRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogRes.ProtoReflect.Descriptor instead.
func (*UpdateBlogRes) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBlogRes) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type DeleteBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBlogReq) Reset() {
	*x = DeleteBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogReq) ProtoMessage() {}

func (x *DeleteBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogReq.ProtoReflect.Descriptor instead.
func (*DeleteBlogReq) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBlogReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBlogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteBlogRes) Reset() {
	*x = DeleteBlogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogRes) ProtoMessage() {}

func (x *DeleteBlogRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogRes.ProtoReflect.Descriptor instead.
func (*DeleteBlogRes) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBlogRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// ListBlogs will use server-streaming
type ListBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBlogReq) Reset() {
	*x = ListBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogReq) ProtoMessage() {}

func (x *ListBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogReq.ProtoReflect.Descriptor instead.
func (*ListBlogReq) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{9}
}

type ListBlogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *ListBlogRes) Reset() {
	*x = ListBlogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blog_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogRes) ProtoMessage() {}

func (x *ListBlogRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogRes.ProtoReflect.Descriptor instead.
func (*ListBlogRes) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{10}
}

func (x *ListBlogRes) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

var File_proto_blog_proto protoreflect.FileDescriptor

var file_proto_blog_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x63, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2f, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x2f,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22,
	0x1d, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d,
	0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x2f, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x2f,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22,
	0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x29, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x22, 0x2d, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6c, 0x6f,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x32, 0xa3, 0x02, 0x0a, 0x0b, 0x42, 0x6c,
	0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x12,
	0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12,
	0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x11, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_blog_proto_rawDescOnce sync.Once
	file_proto_blog_proto_rawDescData = file_proto_blog_proto_rawDesc
)

func file_proto_blog_proto_rawDescGZIP() []byte {
	file_proto_blog_proto_rawDescOnce.Do(func() {
		file_proto_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_blog_proto_rawDescData)
	})
	return file_proto_blog_proto_rawDescData
}

var file_proto_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_blog_proto_goTypes = []interface{}{
	(*Blog)(nil),          // 0: blog.Blog
	(*CreateBlogReq)(nil), // 1: blog.CreateBlogReq
	(*CreateBlogRes)(nil), // 2: blog.CreateBlogRes
	(*ReadBlogReq)(nil),   // 3: blog.ReadBlogReq
	(*ReadBlogRes)(nil),   // 4: blog.ReadBlogRes
	(*UpdateBlogReq)(nil), // 5: blog.UpdateBlogReq
	(*UpdateBlogRes)(nil), // 6: blog.UpdateBlogRes
	(*DeleteBlogReq)(nil), // 7: blog.DeleteBlogReq
	(*DeleteBlogRes)(nil), // 8: blog.DeleteBlogRes
	(*ListBlogReq)(nil),   // 9: blog.ListBlogReq
	(*ListBlogRes)(nil),   // 10: blog.ListBlogRes
}
var file_proto_blog_proto_depIdxs = []int32{
	0,  // 0: blog.CreateBlogReq.blog:type_name -> blog.Blog
	0,  // 1: blog.CreateBlogRes.blog:type_name -> blog.Blog
	0,  // 2: blog.ReadBlogRes.blog:type_name -> blog.Blog
	0,  // 3: blog.UpdateBlogReq.blog:type_name -> blog.Blog
	0,  // 4: blog.UpdateBlogRes.blog:type_name -> blog.Blog
	0,  // 5: blog.ListBlogRes.blog:type_name -> blog.Blog
	1,  // 6: blog.BlogService.CreateBlog:input_type -> blog.CreateBlogReq
	3,  // 7: blog.BlogService.ReadBlog:input_type -> blog.ReadBlogReq
	5,  // 8: blog.BlogService.UpdateBlog:input_type -> blog.UpdateBlogReq
	7,  // 9: blog.BlogService.DeleteBlog:input_type -> blog.DeleteBlogReq
	9,  // 10: blog.BlogService.ListBlog:input_type -> blog.ListBlogReq
	2,  // 11: blog.BlogService.CreateBlog:output_type -> blog.CreateBlogRes
	4,  // 12: blog.BlogService.ReadBlog:output_type -> blog.ReadBlogRes
	6,  // 13: blog.BlogService.UpdateBlog:output_type -> blog.UpdateBlogRes
	8,  // 14: blog.BlogService.DeleteBlog:output_type -> blog.DeleteBlogRes
	9,  // 15: blog.BlogService.ListBlog:output_type -> blog.ListBlogReq
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_blog_proto_init() }
func file_proto_blog_proto_init() {
	if File_proto_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
		file_proto_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogReq); i {
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
		file_proto_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogRes); i {
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
		file_proto_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlogReq); i {
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
		file_proto_blog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlogRes); i {
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
		file_proto_blog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogReq); i {
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
		file_proto_blog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogRes); i {
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
		file_proto_blog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlogReq); i {
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
		file_proto_blog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlogRes); i {
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
		file_proto_blog_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogReq); i {
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
		file_proto_blog_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogRes); i {
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
			RawDescriptor: file_proto_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_blog_proto_goTypes,
		DependencyIndexes: file_proto_blog_proto_depIdxs,
		MessageInfos:      file_proto_blog_proto_msgTypes,
	}.Build()
	File_proto_blog_proto = out.File
	file_proto_blog_proto_rawDesc = nil
	file_proto_blog_proto_goTypes = nil
	file_proto_blog_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	// unary service
	CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error)
	ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error)
	// server streaming - for one request message the server will send back multiple blog messages.
	ListBlog(ctx context.Context, in *ListBlogReq, opts ...grpc.CallOption) (*ListBlogReq, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error) {
	out := new(CreateBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error) {
	out := new(ReadBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error) {
	out := new(UpdateBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error) {
	out := new(DeleteBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListBlog(ctx context.Context, in *ListBlogReq, opts ...grpc.CallOption) (*ListBlogReq, error) {
	out := new(ListBlogReq)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ListBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	// unary service
	CreateBlog(context.Context, *CreateBlogReq) (*CreateBlogRes, error)
	ReadBlog(context.Context, *ReadBlogReq) (*ReadBlogRes, error)
	UpdateBlog(context.Context, *UpdateBlogReq) (*UpdateBlogRes, error)
	DeleteBlog(context.Context, *DeleteBlogReq) (*DeleteBlogRes, error)
	// server streaming - for one request message the server will send back multiple blog messages.
	ListBlog(context.Context, *ListBlogReq) (*ListBlogReq, error)
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreateBlog(context.Context, *CreateBlogReq) (*CreateBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ReadBlog(context.Context, *ReadBlogReq) (*ReadBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (*UnimplementedBlogServiceServer) UpdateBlog(context.Context, *UpdateBlogReq) (*UpdateBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) DeleteBlog(context.Context, *DeleteBlogReq) (*DeleteBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ListBlog(context.Context, *ListBlogReq) (*ListBlogReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlog not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlog(ctx, req.(*DeleteBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ListBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ListBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ListBlog(ctx, req.(*ListBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogService_DeleteBlog_Handler,
		},
		{
			MethodName: "ListBlog",
			Handler:    _BlogService_ListBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blog.proto",
}
