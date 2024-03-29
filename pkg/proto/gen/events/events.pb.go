// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: events/events.proto

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

// Event container. At the time may be fired more than one event.
type Events struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTime uint64   `protobuf:"varint,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	CompanyId  string   `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Events     []*Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Events) Reset() {
	*x = Events{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Events) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events) ProtoMessage() {}

func (x *Events) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events.ProtoReflect.Descriptor instead.
func (*Events) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{0}
}

func (x *Events) GetServerTime() uint64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *Events) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *Events) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdempotencyKey string `protobuf:"bytes,1,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*Event_Login
	//	*Event_SignUp
	Payload isEvent_Payload `protobuf_oneof:"payload"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[1]
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
	return file_events_events_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetIdempotencyKey() string {
	if x != nil {
		return x.IdempotencyKey
	}
	return ""
}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Event) GetLogin() *EventUserLogin {
	if x, ok := x.GetPayload().(*Event_Login); ok {
		return x.Login
	}
	return nil
}

func (x *Event) GetSignUp() *EventUserSignedUp {
	if x, ok := x.GetPayload().(*Event_SignUp); ok {
		return x.SignUp
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_Login struct {
	Login *EventUserLogin `protobuf:"bytes,2,opt,name=login,proto3,oneof"`
}

type Event_SignUp struct {
	SignUp *EventUserSignedUp `protobuf:"bytes,3,opt,name=sign_up,json=signUp,proto3,oneof"`
}

func (*Event_Login) isEvent_Payload() {}

func (*Event_SignUp) isEvent_Payload() {}

// Event thrown by Warden service. Indicates that user logged in.
// This event is handeled by LogstashService (TODO rename).
type EventUserLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Time      uint64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *EventUserLogin) Reset() {
	*x = EventUserLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventUserLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventUserLogin) ProtoMessage() {}

func (x *EventUserLogin) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventUserLogin.ProtoReflect.Descriptor instead.
func (*EventUserLogin) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{2}
}

func (x *EventUserLogin) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EventUserLogin) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *EventUserLogin) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type EventUserSignedUp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Time      uint64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *EventUserSignedUp) Reset() {
	*x = EventUserSignedUp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventUserSignedUp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventUserSignedUp) ProtoMessage() {}

func (x *EventUserSignedUp) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventUserSignedUp.ProtoReflect.Descriptor instead.
func (*EventUserSignedUp) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{3}
}

func (x *EventUserSignedUp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EventUserSignedUp) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *EventUserSignedUp) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_events_events_proto protoreflect.FileDescriptor

var file_events_events_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x22, 0x72, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x69,
	0x67, 0x6e, 0x5f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x70, 0x48, 0x00, 0x52, 0x06, 0x73, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5c,
	0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x11,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55,
	0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x63,
	0x6c, 0x62, 0x6c, 0x61, 0x73, 0x74, 0x2f, 0x62, 0x69, 0x6f, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_events_proto_rawDescOnce sync.Once
	file_events_events_proto_rawDescData = file_events_events_proto_rawDesc
)

func file_events_events_proto_rawDescGZIP() []byte {
	file_events_events_proto_rawDescOnce.Do(func() {
		file_events_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_events_proto_rawDescData)
	})
	return file_events_events_proto_rawDescData
}

var file_events_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_events_events_proto_goTypes = []interface{}{
	(*Events)(nil),            // 0: events.v1.Events
	(*Event)(nil),             // 1: events.v1.Event
	(*EventUserLogin)(nil),    // 2: events.v1.EventUserLogin
	(*EventUserSignedUp)(nil), // 3: events.v1.EventUserSignedUp
}
var file_events_events_proto_depIdxs = []int32{
	1, // 0: events.v1.Events.events:type_name -> events.v1.Event
	2, // 1: events.v1.Event.login:type_name -> events.v1.EventUserLogin
	3, // 2: events.v1.Event.sign_up:type_name -> events.v1.EventUserSignedUp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_events_events_proto_init() }
func file_events_events_proto_init() {
	if File_events_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Events); i {
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
		file_events_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_events_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventUserLogin); i {
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
		file_events_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventUserSignedUp); i {
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
	file_events_events_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Event_Login)(nil),
		(*Event_SignUp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_events_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_events_proto_goTypes,
		DependencyIndexes: file_events_events_proto_depIdxs,
		MessageInfos:      file_events_events_proto_msgTypes,
	}.Build()
	File_events_events_proto = out.File
	file_events_events_proto_rawDesc = nil
	file_events_events_proto_goTypes = nil
	file_events_events_proto_depIdxs = nil
}
