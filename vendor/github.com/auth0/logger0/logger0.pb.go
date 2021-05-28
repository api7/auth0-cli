// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: api/logger0.proto

package logger0

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

type LogRecord_Type int32

const (
	LogRecord_TYPE_UNSPECIFIED LogRecord_Type = 0
	LogRecord_TYPE_ACTIONS     LogRecord_Type = 1
)

// Enum value maps for LogRecord_Type.
var (
	LogRecord_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_ACTIONS",
	}
	LogRecord_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_ACTIONS":     1,
	}
)

func (x LogRecord_Type) Enum() *LogRecord_Type {
	p := new(LogRecord_Type)
	*p = x
	return p
}

func (x LogRecord_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogRecord_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_logger0_proto_enumTypes[0].Descriptor()
}

func (LogRecord_Type) Type() protoreflect.EnumType {
	return &file_api_logger0_proto_enumTypes[0]
}

func (x LogRecord_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogRecord_Type.Descriptor instead.
func (LogRecord_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_logger0_proto_rawDescGZIP(), []int{0, 0}
}

type LogRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     LogRecord_Type    `protobuf:"varint,1,opt,name=type,proto3,enum=auth0.logger0.LogRecord_Type" json:"type,omitempty"`
	Tenant   string            `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Messages [][]byte          `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *LogRecord) Reset() {
	*x = LogRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logger0_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRecord) ProtoMessage() {}

func (x *LogRecord) ProtoReflect() protoreflect.Message {
	mi := &file_api_logger0_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRecord.ProtoReflect.Descriptor instead.
func (*LogRecord) Descriptor() ([]byte, []int) {
	return file_api_logger0_proto_rawDescGZIP(), []int{0}
}

func (x *LogRecord) GetType() LogRecord_Type {
	if x != nil {
		return x.Type
	}
	return LogRecord_TYPE_UNSPECIFIED
}

func (x *LogRecord) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *LogRecord) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *LogRecord) GetMessages() [][]byte {
	if x != nil {
		return x.Messages
	}
	return nil
}

type Drain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant  string    `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Sink    *Sink     `protobuf:"bytes,3,opt,name=sink,proto3" json:"sink,omitempty"`
	Filters []*Filter `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *Drain) Reset() {
	*x = Drain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logger0_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Drain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Drain) ProtoMessage() {}

func (x *Drain) ProtoReflect() protoreflect.Message {
	mi := &file_api_logger0_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Drain.ProtoReflect.Descriptor instead.
func (*Drain) Descriptor() ([]byte, []int) {
	return file_api_logger0_proto_rawDescGZIP(), []int{1}
}

func (x *Drain) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Drain) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *Drain) GetSink() *Sink {
	if x != nil {
		return x.Sink
	}
	return nil
}

func (x *Drain) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant  string    `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Filters []*Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logger0_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_api_logger0_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_api_logger0_proto_rawDescGZIP(), []int{2}
}

func (x *Session) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *Session) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logger0_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_api_logger0_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_api_logger0_proto_rawDescGZIP(), []int{3}
}

func (x *Filter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Filter) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type Sink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//	*Sink_Url
	Target isSink_Target `protobuf_oneof:"target"`
}

func (x *Sink) Reset() {
	*x = Sink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logger0_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink) ProtoMessage() {}

func (x *Sink) ProtoReflect() protoreflect.Message {
	mi := &file_api_logger0_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sink.ProtoReflect.Descriptor instead.
func (*Sink) Descriptor() ([]byte, []int) {
	return file_api_logger0_proto_rawDescGZIP(), []int{4}
}

func (m *Sink) GetTarget() isSink_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *Sink) GetUrl() string {
	if x, ok := x.GetTarget().(*Sink_Url); ok {
		return x.Url
	}
	return ""
}

type isSink_Target interface {
	isSink_Target()
}

type Sink_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}

func (*Sink_Url) isSink_Target() {}

var File_api_logger0_proto protoreflect.FileDescriptor

var file_api_logger0_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x30, 0x22, 0xa3, 0x02, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x01, 0x22, 0x89, 0x01, 0x0a, 0x05, 0x44, 0x72, 0x61,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x73, 0x69,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x30,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x73,
	0x69, 0x6e, 0x6b, 0x12, 0x2f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x30, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x22, 0x52, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x30,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x24, 0x0a, 0x04, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x1a, 0x5a, 0x18,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x30,
	0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_logger0_proto_rawDescOnce sync.Once
	file_api_logger0_proto_rawDescData = file_api_logger0_proto_rawDesc
)

func file_api_logger0_proto_rawDescGZIP() []byte {
	file_api_logger0_proto_rawDescOnce.Do(func() {
		file_api_logger0_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_logger0_proto_rawDescData)
	})
	return file_api_logger0_proto_rawDescData
}

var file_api_logger0_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_logger0_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_logger0_proto_goTypes = []interface{}{
	(LogRecord_Type)(0), // 0: auth0.logger0.LogRecord.Type
	(*LogRecord)(nil),   // 1: auth0.logger0.LogRecord
	(*Drain)(nil),       // 2: auth0.logger0.Drain
	(*Session)(nil),     // 3: auth0.logger0.Session
	(*Filter)(nil),      // 4: auth0.logger0.Filter
	(*Sink)(nil),        // 5: auth0.logger0.Sink
	nil,                 // 6: auth0.logger0.LogRecord.MetadataEntry
}
var file_api_logger0_proto_depIdxs = []int32{
	0, // 0: auth0.logger0.LogRecord.type:type_name -> auth0.logger0.LogRecord.Type
	6, // 1: auth0.logger0.LogRecord.metadata:type_name -> auth0.logger0.LogRecord.MetadataEntry
	5, // 2: auth0.logger0.Drain.sink:type_name -> auth0.logger0.Sink
	4, // 3: auth0.logger0.Drain.filters:type_name -> auth0.logger0.Filter
	4, // 4: auth0.logger0.Session.filters:type_name -> auth0.logger0.Filter
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_logger0_proto_init() }
func file_api_logger0_proto_init() {
	if File_api_logger0_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_logger0_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRecord); i {
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
		file_api_logger0_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Drain); i {
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
		file_api_logger0_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_api_logger0_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_api_logger0_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sink); i {
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
	file_api_logger0_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Sink_Url)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_logger0_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_logger0_proto_goTypes,
		DependencyIndexes: file_api_logger0_proto_depIdxs,
		EnumInfos:         file_api_logger0_proto_enumTypes,
		MessageInfos:      file_api_logger0_proto_msgTypes,
	}.Build()
	File_api_logger0_proto = out.File
	file_api_logger0_proto_rawDesc = nil
	file_api_logger0_proto_goTypes = nil
	file_api_logger0_proto_depIdxs = nil
}
