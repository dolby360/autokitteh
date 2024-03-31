// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: autokitteh/events/v1/event.proto

package eventsv1

import (
	v1 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/values/v1"
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

type EventState int32

const (
	EventState_EVENT_STATE_UNSPECIFIED EventState = 0
	EventState_EVENT_STATE_SAVED       EventState = 1
	EventState_EVENT_STATE_PROCESSING  EventState = 2
	EventState_EVENT_STATE_COMPLETED   EventState = 3
	EventState_EVENT_STATE_FAILED      EventState = 4
)

// Enum value maps for EventState.
var (
	EventState_name = map[int32]string{
		0: "EVENT_STATE_UNSPECIFIED",
		1: "EVENT_STATE_SAVED",
		2: "EVENT_STATE_PROCESSING",
		3: "EVENT_STATE_COMPLETED",
		4: "EVENT_STATE_FAILED",
	}
	EventState_value = map[string]int32{
		"EVENT_STATE_UNSPECIFIED": 0,
		"EVENT_STATE_SAVED":       1,
		"EVENT_STATE_PROCESSING":  2,
		"EVENT_STATE_COMPLETED":   3,
		"EVENT_STATE_FAILED":      4,
	}
)

func (x EventState) Enum() *EventState {
	p := new(EventState)
	*p = x
	return p
}

func (x EventState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventState) Descriptor() protoreflect.EnumDescriptor {
	return file_autokitteh_events_v1_event_proto_enumTypes[0].Descriptor()
}

func (EventState) Type() protoreflect.EnumType {
	return &file_autokitteh_events_v1_event_proto_enumTypes[0]
}

func (x EventState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventState.Descriptor instead.
func (EventState) EnumDescriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_event_proto_rawDescGZIP(), []int{0}
}

type EventRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq       uint32                 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	EventId   string                 `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	State     EventState             `protobuf:"varint,3,opt,name=state,proto3,enum=autokitteh.events.v1.EventState" json:"state,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *EventRecord) Reset() {
	*x = EventRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRecord) ProtoMessage() {}

func (x *EventRecord) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRecord.ProtoReflect.Descriptor instead.
func (*EventRecord) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventRecord) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *EventRecord) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *EventRecord) GetState() EventState {
	if x != nil {
		return x.State
	}
	return EventState_EVENT_STATE_UNSPECIFIED
}

func (x *EventRecord) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId          string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	IntegrationId    string                 `protobuf:"bytes,2,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	IntegrationToken string                 `protobuf:"bytes,3,opt,name=integration_token,json=integrationToken,proto3" json:"integration_token,omitempty"` // TODO: think of a name that does not hint authn.
	EventType        string                 `protobuf:"bytes,4,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Data             map[string]*v1.Value   `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Memo             map[string]string      `protobuf:"bytes,6,rep,name=memo,proto3" json:"memo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Seq              uint64                 `protobuf:"varint,8,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Event) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *Event) GetIntegrationToken() string {
	if x != nil {
		return x.IntegrationToken
	}
	return ""
}

func (x *Event) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Event) GetData() map[string]*v1.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Event) GetMemo() map[string]string {
	if x != nil {
		return x.Memo
	}
	return nil
}

func (x *Event) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Event) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

var File_autokitteh_events_v1_event_proto protoreflect.FileDescriptor

var file_autokitteh_events_v1_event_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69,
	0x74, 0x74, 0x65, 0x68, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a,
	0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b,
	0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe7, 0x03, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x39, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x73, 0x65, 0x71, 0x1a, 0x54, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a,
	0x09, 0x4d, 0x65, 0x6d, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x8f, 0x01, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x53, 0x41, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x16, 0x0a, 0x12, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x42, 0xe1, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41,
	0x45, 0x58, 0xaa, 0x02, 0x14, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x41, 0x75, 0x74, 0x6f,
	0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x20, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68,
	0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autokitteh_events_v1_event_proto_rawDescOnce sync.Once
	file_autokitteh_events_v1_event_proto_rawDescData = file_autokitteh_events_v1_event_proto_rawDesc
)

func file_autokitteh_events_v1_event_proto_rawDescGZIP() []byte {
	file_autokitteh_events_v1_event_proto_rawDescOnce.Do(func() {
		file_autokitteh_events_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_autokitteh_events_v1_event_proto_rawDescData)
	})
	return file_autokitteh_events_v1_event_proto_rawDescData
}

var file_autokitteh_events_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_autokitteh_events_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_autokitteh_events_v1_event_proto_goTypes = []interface{}{
	(EventState)(0),               // 0: autokitteh.events.v1.EventState
	(*EventRecord)(nil),           // 1: autokitteh.events.v1.EventRecord
	(*Event)(nil),                 // 2: autokitteh.events.v1.Event
	nil,                           // 3: autokitteh.events.v1.Event.DataEntry
	nil,                           // 4: autokitteh.events.v1.Event.MemoEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*v1.Value)(nil),              // 6: autokitteh.values.v1.Value
}
var file_autokitteh_events_v1_event_proto_depIdxs = []int32{
	0, // 0: autokitteh.events.v1.EventRecord.state:type_name -> autokitteh.events.v1.EventState
	5, // 1: autokitteh.events.v1.EventRecord.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: autokitteh.events.v1.Event.data:type_name -> autokitteh.events.v1.Event.DataEntry
	4, // 3: autokitteh.events.v1.Event.memo:type_name -> autokitteh.events.v1.Event.MemoEntry
	5, // 4: autokitteh.events.v1.Event.created_at:type_name -> google.protobuf.Timestamp
	6, // 5: autokitteh.events.v1.Event.DataEntry.value:type_name -> autokitteh.values.v1.Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_autokitteh_events_v1_event_proto_init() }
func file_autokitteh_events_v1_event_proto_init() {
	if File_autokitteh_events_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autokitteh_events_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRecord); i {
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
		file_autokitteh_events_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_autokitteh_events_v1_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_autokitteh_events_v1_event_proto_goTypes,
		DependencyIndexes: file_autokitteh_events_v1_event_proto_depIdxs,
		EnumInfos:         file_autokitteh_events_v1_event_proto_enumTypes,
		MessageInfos:      file_autokitteh_events_v1_event_proto_msgTypes,
	}.Build()
	File_autokitteh_events_v1_event_proto = out.File
	file_autokitteh_events_v1_event_proto_rawDesc = nil
	file_autokitteh_events_v1_event_proto_goTypes = nil
	file_autokitteh_events_v1_event_proto_depIdxs = nil
}
