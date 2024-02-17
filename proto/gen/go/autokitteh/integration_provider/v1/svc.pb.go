// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: autokitteh/integration_provider/v1/svc.proto

package integration_providerv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v11 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/program/v1"
	v1 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/values/v1"
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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId string `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	// TODO(ENG-113): Is this needed / available in a remote integration?
	ExecutorId string `protobuf:"bytes,2,opt,name=executor_id,json=executorId,proto3" json:"executor_id,omitempty"`
	// TODO(ENG-113): Is this needed?
	ConnectionToken string `protobuf:"bytes,3,opt,name=connection_token,json=connectionToken,proto3" json:"connection_token,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_integration_provider_v1_svc_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *GetRequest) GetExecutorId() string {
	if x != nil {
		return x.ExecutorId
	}
	return ""
}

func (x *GetRequest) GetConnectionToken() string {
	if x != nil {
		return x.ConnectionToken
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Integration *Integration `protobuf:"bytes,1,opt,name=integration,proto3" json:"integration,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_integration_provider_v1_svc_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetIntegration() *Integration {
	if x != nil {
		return x.Integration
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_autokitteh_integration_provider_v1_svc_proto_rawDescGZIP(), []int{2}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Integrations []*Integration `protobuf:"bytes,1,rep,name=integrations,proto3" json:"integrations,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_autokitteh_integration_provider_v1_svc_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetIntegrations() []*Integration {
	if x != nil {
		return x.Integrations
	}
	return nil
}

// TODO(ENG-112): This part of the API is still being designed.
type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId string               `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	Function      *v1.Value            `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
	Args          []*v1.Value          `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	Kwargs        map[string]*v1.Value `protobuf:"bytes,4,rep,name=kwargs,proto3" json:"kwargs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_integration_provider_v1_svc_proto_rawDescGZIP(), []int{4}
}

func (x *CallRequest) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *CallRequest) GetFunction() *v1.Value {
	if x != nil {
		return x.Function
	}
	return nil
}

func (x *CallRequest) GetArgs() []*v1.Value {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *CallRequest) GetKwargs() map[string]*v1.Value {
	if x != nil {
		return x.Kwargs
	}
	return nil
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *v1.Value  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Error *v11.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_integration_provider_v1_svc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_integration_provider_v1_svc_proto_rawDescGZIP(), []int{5}
}

func (x *CallResponse) GetValue() *v1.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *CallResponse) GetError() *v11.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_autokitteh_integration_provider_v1_svc_proto protoreflect.FileDescriptor

var file_autokitteh_integration_provider_v1_svc_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x34, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69,
	0x74, 0x74, 0x65, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x61,
	0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0d,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a,
	0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x60, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69,
	0x74, 0x74, 0x65, 0x68, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x80, 0x03, 0x0a, 0x0b, 0x43, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0xfa, 0xf7, 0x18, 0x08, 0x92, 0x01, 0x05,
	0x22, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x67, 0x0a, 0x06, 0x6b,
	0x77, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4b, 0x77, 0x61,
	0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x12, 0xfa, 0xf7, 0x18, 0x0e, 0x9a, 0x01,
	0x0b, 0x22, 0x04, 0x72, 0x02, 0x10, 0x01, 0x2a, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x6b, 0x77,
	0x61, 0x72, 0x67, 0x73, 0x1a, 0x56, 0x0a, 0x0b, 0x4b, 0x77, 0x61, 0x72, 0x67, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x75, 0x0a, 0x0c,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x32, 0xda, 0x02, 0x0a, 0x1a, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x66, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2e, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2f, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x2f, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xbd, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74,
	0x74, 0x65, 0x68, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x53, 0x76, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b,
	0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41,
	0x49, 0x58, 0xaa, 0x02, 0x21, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x21, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74,
	0x74, 0x65, 0x68, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2d, 0x41, 0x75, 0x74,
	0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23, 0x41, 0x75, 0x74,
	0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autokitteh_integration_provider_v1_svc_proto_rawDescOnce sync.Once
	file_autokitteh_integration_provider_v1_svc_proto_rawDescData = file_autokitteh_integration_provider_v1_svc_proto_rawDesc
)

func file_autokitteh_integration_provider_v1_svc_proto_rawDescGZIP() []byte {
	file_autokitteh_integration_provider_v1_svc_proto_rawDescOnce.Do(func() {
		file_autokitteh_integration_provider_v1_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_autokitteh_integration_provider_v1_svc_proto_rawDescData)
	})
	return file_autokitteh_integration_provider_v1_svc_proto_rawDescData
}

var file_autokitteh_integration_provider_v1_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_autokitteh_integration_provider_v1_svc_proto_goTypes = []interface{}{
	(*GetRequest)(nil),   // 0: autokitteh.integration_provider.v1.GetRequest
	(*GetResponse)(nil),  // 1: autokitteh.integration_provider.v1.GetResponse
	(*ListRequest)(nil),  // 2: autokitteh.integration_provider.v1.ListRequest
	(*ListResponse)(nil), // 3: autokitteh.integration_provider.v1.ListResponse
	(*CallRequest)(nil),  // 4: autokitteh.integration_provider.v1.CallRequest
	(*CallResponse)(nil), // 5: autokitteh.integration_provider.v1.CallResponse
	nil,                  // 6: autokitteh.integration_provider.v1.CallRequest.KwargsEntry
	(*Integration)(nil),  // 7: autokitteh.integration_provider.v1.Integration
	(*v1.Value)(nil),     // 8: autokitteh.values.v1.Value
	(*v11.Error)(nil),    // 9: autokitteh.program.v1.Error
}
var file_autokitteh_integration_provider_v1_svc_proto_depIdxs = []int32{
	7,  // 0: autokitteh.integration_provider.v1.GetResponse.integration:type_name -> autokitteh.integration_provider.v1.Integration
	7,  // 1: autokitteh.integration_provider.v1.ListResponse.integrations:type_name -> autokitteh.integration_provider.v1.Integration
	8,  // 2: autokitteh.integration_provider.v1.CallRequest.function:type_name -> autokitteh.values.v1.Value
	8,  // 3: autokitteh.integration_provider.v1.CallRequest.args:type_name -> autokitteh.values.v1.Value
	6,  // 4: autokitteh.integration_provider.v1.CallRequest.kwargs:type_name -> autokitteh.integration_provider.v1.CallRequest.KwargsEntry
	8,  // 5: autokitteh.integration_provider.v1.CallResponse.value:type_name -> autokitteh.values.v1.Value
	9,  // 6: autokitteh.integration_provider.v1.CallResponse.error:type_name -> autokitteh.program.v1.Error
	8,  // 7: autokitteh.integration_provider.v1.CallRequest.KwargsEntry.value:type_name -> autokitteh.values.v1.Value
	0,  // 8: autokitteh.integration_provider.v1.IntegrationProviderService.Get:input_type -> autokitteh.integration_provider.v1.GetRequest
	2,  // 9: autokitteh.integration_provider.v1.IntegrationProviderService.List:input_type -> autokitteh.integration_provider.v1.ListRequest
	4,  // 10: autokitteh.integration_provider.v1.IntegrationProviderService.Call:input_type -> autokitteh.integration_provider.v1.CallRequest
	1,  // 11: autokitteh.integration_provider.v1.IntegrationProviderService.Get:output_type -> autokitteh.integration_provider.v1.GetResponse
	3,  // 12: autokitteh.integration_provider.v1.IntegrationProviderService.List:output_type -> autokitteh.integration_provider.v1.ListResponse
	5,  // 13: autokitteh.integration_provider.v1.IntegrationProviderService.Call:output_type -> autokitteh.integration_provider.v1.CallResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_autokitteh_integration_provider_v1_svc_proto_init() }
func file_autokitteh_integration_provider_v1_svc_proto_init() {
	if File_autokitteh_integration_provider_v1_svc_proto != nil {
		return
	}
	file_autokitteh_integration_provider_v1_integration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_autokitteh_integration_provider_v1_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_autokitteh_integration_provider_v1_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_autokitteh_integration_provider_v1_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_autokitteh_integration_provider_v1_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_autokitteh_integration_provider_v1_svc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_autokitteh_integration_provider_v1_svc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
			RawDescriptor: file_autokitteh_integration_provider_v1_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_autokitteh_integration_provider_v1_svc_proto_goTypes,
		DependencyIndexes: file_autokitteh_integration_provider_v1_svc_proto_depIdxs,
		MessageInfos:      file_autokitteh_integration_provider_v1_svc_proto_msgTypes,
	}.Build()
	File_autokitteh_integration_provider_v1_svc_proto = out.File
	file_autokitteh_integration_provider_v1_svc_proto_rawDesc = nil
	file_autokitteh_integration_provider_v1_svc_proto_goTypes = nil
	file_autokitteh_integration_provider_v1_svc_proto_depIdxs = nil
}
