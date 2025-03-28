// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.21.12
// source: controllerapi/controllerapi.proto

package controllerapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WatchAction int32

const (
	WatchAction_CREATE WatchAction = 0
	WatchAction_UPDATE WatchAction = 1
	WatchAction_DELETE WatchAction = 2
	WatchAction_ALL    WatchAction = 3
)

// Enum value maps for WatchAction.
var (
	WatchAction_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
		3: "ALL",
	}
	WatchAction_value = map[string]int32{
		"CREATE": 0,
		"UPDATE": 1,
		"DELETE": 2,
		"ALL":    3,
	}
)

func (x WatchAction) Enum() *WatchAction {
	p := new(WatchAction)
	*p = x
	return p
}

func (x WatchAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WatchAction) Descriptor() protoreflect.EnumDescriptor {
	return file_controllerapi_controllerapi_proto_enumTypes[0].Descriptor()
}

func (WatchAction) Type() protoreflect.EnumType {
	return &file_controllerapi_controllerapi_proto_enumTypes[0]
}

func (x WatchAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WatchAction.Descriptor instead.
func (WatchAction) EnumDescriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{0}
}

// TODO: Add filtering, now all resources returned are those
//
//	who are newly changed or created. The client has to
//	filter based on the Resource.Status.Phase
type WatchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Id            *string                `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Owner         *Owner                 `protobuf:"bytes,3,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
	Action        WatchAction            `protobuf:"varint,4,opt,name=action,proto3,enum=controllerapi.WatchAction" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{0}
}

func (x *WatchRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *WatchRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *WatchRequest) GetOwner() *Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *WatchRequest) GetAction() WatchAction {
	if x != nil {
		return x.Action
	}
	return WatchAction_CREATE
}

type WatchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resource      *Resource              `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	PrevResource  *Resource              `protobuf:"bytes,2,opt,name=prev_resource,json=prevResource,proto3,oneof" json:"prev_resource,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{1}
}

func (x *WatchResponse) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *WatchResponse) GetPrevResource() *Resource {
	if x != nil {
		return x.PrevResource
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Id            *string                `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Owner         *Owner                 `protobuf:"bytes,3,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ListRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *ListRequest) GetOwner() *Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resources     []*Resource            `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type PatchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Id            *string                `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Owner         *Owner                 `protobuf:"bytes,3,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
	Patches       []byte                 `protobuf:"bytes,4,opt,name=patches,proto3" json:"patches,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PatchRequest) Reset() {
	*x = PatchRequest{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchRequest) ProtoMessage() {}

func (x *PatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchRequest.ProtoReflect.Descriptor instead.
func (*PatchRequest) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{4}
}

func (x *PatchRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *PatchRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *PatchRequest) GetOwner() *Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *PatchRequest) GetPatches() []byte {
	if x != nil {
		return x.Patches
	}
	return nil
}

type PatchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PatchResponse) Reset() {
	*x = PatchResponse{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchResponse) ProtoMessage() {}

func (x *PatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchResponse.ProtoReflect.Descriptor instead.
func (*PatchResponse) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{5}
}

func (x *PatchResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type CreateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This allows clients to fill in their own ids for example. And saves a lot of pain.
	// The controller will only validate the resource, append neccesary fields if not set,
	// such as id.
	Resource      *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{6}
}

func (x *CreateRequest) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{7}
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{9}
}

type Owner struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Owner) Reset() {
	*x = Owner{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{10}
}

func (x *Owner) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Owner) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind          string                 `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Annotations   map[string]string      `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Owner         *Owner                 `protobuf:"bytes,4,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
	Spec          *structpb.Struct       `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	Status        *structpb.Struct       `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Phase         string                 `protobuf:"bytes,7,opt,name=phase,proto3" json:"phase,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Resource) Reset() {
	*x = Resource{}
	mi := &file_controllerapi_controllerapi_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_controllerapi_controllerapi_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_controllerapi_controllerapi_proto_rawDescGZIP(), []int{11}
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Resource) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Resource) GetOwner() *Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Resource) GetSpec() *structpb.Struct {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Resource) GetStatus() *structpb.Struct {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Resource) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

var File_controllerapi_controllerapi_proto protoreflect.FileDescriptor

var file_controllerapi_controllerapi_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xad, 0x01, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x48,
	0x01, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x22, 0x99, 0x01, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70,
	0x72, 0x65, 0x76, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x78, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x48, 0x01, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x93, 0x01,
	0x0a, 0x0c, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x48, 0x01, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x22, 0x44, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xe9, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2b,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x2f, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x61,
	0x73, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2a, 0x3a, 0x0a, 0x0b,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x03, 0x32, 0xe8, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x41, 0x70, 0x69, 0x12, 0x44, 0x0a, 0x05, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x3f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x73, 0x6b, 0x70, 0x69, 0x6c, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x66, 0x65, 0x72,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_controllerapi_controllerapi_proto_rawDescOnce sync.Once
	file_controllerapi_controllerapi_proto_rawDescData []byte
)

func file_controllerapi_controllerapi_proto_rawDescGZIP() []byte {
	file_controllerapi_controllerapi_proto_rawDescOnce.Do(func() {
		file_controllerapi_controllerapi_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controllerapi_controllerapi_proto_rawDesc), len(file_controllerapi_controllerapi_proto_rawDesc)))
	})
	return file_controllerapi_controllerapi_proto_rawDescData
}

var file_controllerapi_controllerapi_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controllerapi_controllerapi_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_controllerapi_controllerapi_proto_goTypes = []any{
	(WatchAction)(0),        // 0: controllerapi.WatchAction
	(*WatchRequest)(nil),    // 1: controllerapi.WatchRequest
	(*WatchResponse)(nil),   // 2: controllerapi.WatchResponse
	(*ListRequest)(nil),     // 3: controllerapi.ListRequest
	(*ListResponse)(nil),    // 4: controllerapi.ListResponse
	(*PatchRequest)(nil),    // 5: controllerapi.PatchRequest
	(*PatchResponse)(nil),   // 6: controllerapi.PatchResponse
	(*CreateRequest)(nil),   // 7: controllerapi.CreateRequest
	(*CreateResponse)(nil),  // 8: controllerapi.CreateResponse
	(*DeleteRequest)(nil),   // 9: controllerapi.DeleteRequest
	(*DeleteResponse)(nil),  // 10: controllerapi.DeleteResponse
	(*Owner)(nil),           // 11: controllerapi.Owner
	(*Resource)(nil),        // 12: controllerapi.Resource
	nil,                     // 13: controllerapi.Resource.AnnotationsEntry
	(*structpb.Struct)(nil), // 14: google.protobuf.Struct
}
var file_controllerapi_controllerapi_proto_depIdxs = []int32{
	11, // 0: controllerapi.WatchRequest.owner:type_name -> controllerapi.Owner
	0,  // 1: controllerapi.WatchRequest.action:type_name -> controllerapi.WatchAction
	12, // 2: controllerapi.WatchResponse.resource:type_name -> controllerapi.Resource
	12, // 3: controllerapi.WatchResponse.prev_resource:type_name -> controllerapi.Resource
	11, // 4: controllerapi.ListRequest.owner:type_name -> controllerapi.Owner
	12, // 5: controllerapi.ListResponse.resources:type_name -> controllerapi.Resource
	11, // 6: controllerapi.PatchRequest.owner:type_name -> controllerapi.Owner
	12, // 7: controllerapi.CreateRequest.resource:type_name -> controllerapi.Resource
	13, // 8: controllerapi.Resource.annotations:type_name -> controllerapi.Resource.AnnotationsEntry
	11, // 9: controllerapi.Resource.owner:type_name -> controllerapi.Owner
	14, // 10: controllerapi.Resource.spec:type_name -> google.protobuf.Struct
	14, // 11: controllerapi.Resource.status:type_name -> google.protobuf.Struct
	1,  // 12: controllerapi.ControllerApi.Watch:input_type -> controllerapi.WatchRequest
	3,  // 13: controllerapi.ControllerApi.List:input_type -> controllerapi.ListRequest
	7,  // 14: controllerapi.ControllerApi.Create:input_type -> controllerapi.CreateRequest
	5,  // 15: controllerapi.ControllerApi.Patch:input_type -> controllerapi.PatchRequest
	9,  // 16: controllerapi.ControllerApi.Delete:input_type -> controllerapi.DeleteRequest
	2,  // 17: controllerapi.ControllerApi.Watch:output_type -> controllerapi.WatchResponse
	4,  // 18: controllerapi.ControllerApi.List:output_type -> controllerapi.ListResponse
	8,  // 19: controllerapi.ControllerApi.Create:output_type -> controllerapi.CreateResponse
	6,  // 20: controllerapi.ControllerApi.Patch:output_type -> controllerapi.PatchResponse
	10, // 21: controllerapi.ControllerApi.Delete:output_type -> controllerapi.DeleteResponse
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_controllerapi_controllerapi_proto_init() }
func file_controllerapi_controllerapi_proto_init() {
	if File_controllerapi_controllerapi_proto != nil {
		return
	}
	file_controllerapi_controllerapi_proto_msgTypes[0].OneofWrappers = []any{}
	file_controllerapi_controllerapi_proto_msgTypes[1].OneofWrappers = []any{}
	file_controllerapi_controllerapi_proto_msgTypes[2].OneofWrappers = []any{}
	file_controllerapi_controllerapi_proto_msgTypes[4].OneofWrappers = []any{}
	file_controllerapi_controllerapi_proto_msgTypes[11].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controllerapi_controllerapi_proto_rawDesc), len(file_controllerapi_controllerapi_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controllerapi_controllerapi_proto_goTypes,
		DependencyIndexes: file_controllerapi_controllerapi_proto_depIdxs,
		EnumInfos:         file_controllerapi_controllerapi_proto_enumTypes,
		MessageInfos:      file_controllerapi_controllerapi_proto_msgTypes,
	}.Build()
	File_controllerapi_controllerapi_proto = out.File
	file_controllerapi_controllerapi_proto_goTypes = nil
	file_controllerapi_controllerapi_proto_depIdxs = nil
}
