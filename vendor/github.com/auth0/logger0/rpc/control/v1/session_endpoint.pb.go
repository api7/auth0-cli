// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: rpc/control/v1/session_endpoint.proto

package controlv1

import (
	logger0 "github.com/auth0/logger0"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant  string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Type    logger0.LogRecord_Type `protobuf:"varint,2,opt,name=type,proto3,enum=auth0.logger0.LogRecord_Type" json:"type,omitempty"`
	Filters map[string]string      `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_control_v1_session_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_control_v1_session_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_rpc_control_v1_session_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSessionRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CreateSessionRequest) GetType() logger0.LogRecord_Type {
	if x != nil {
		return x.Type
	}
	return logger0.LogRecord_TYPE_UNSPECIFIED
}

func (x *CreateSessionRequest) GetFilters() map[string]string {
	if x != nil {
		return x.Filters
	}
	return nil
}

var File_rpc_control_v1_session_endpoint_proto protoreflect.FileDescriptor

var file_rpc_control_v1_session_endpoint_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x30, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x87, 0x01, 0x0a, 0x0f, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x74, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x01,
	0x2a, 0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x30, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x30, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_control_v1_session_endpoint_proto_rawDescOnce sync.Once
	file_rpc_control_v1_session_endpoint_proto_rawDescData = file_rpc_control_v1_session_endpoint_proto_rawDesc
)

func file_rpc_control_v1_session_endpoint_proto_rawDescGZIP() []byte {
	file_rpc_control_v1_session_endpoint_proto_rawDescOnce.Do(func() {
		file_rpc_control_v1_session_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_control_v1_session_endpoint_proto_rawDescData)
	})
	return file_rpc_control_v1_session_endpoint_proto_rawDescData
}

var file_rpc_control_v1_session_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_control_v1_session_endpoint_proto_goTypes = []interface{}{
	(*CreateSessionRequest)(nil), // 0: auth0.logger0.control.v1.CreateSessionRequest
	nil,                          // 1: auth0.logger0.control.v1.CreateSessionRequest.FiltersEntry
	(logger0.LogRecord_Type)(0),  // 2: auth0.logger0.LogRecord.Type
	(*logger0.LogRecord)(nil),    // 3: auth0.logger0.LogRecord
}
var file_rpc_control_v1_session_endpoint_proto_depIdxs = []int32{
	2, // 0: auth0.logger0.control.v1.CreateSessionRequest.type:type_name -> auth0.logger0.LogRecord.Type
	1, // 1: auth0.logger0.control.v1.CreateSessionRequest.filters:type_name -> auth0.logger0.control.v1.CreateSessionRequest.FiltersEntry
	0, // 2: auth0.logger0.control.v1.SessionEndpoint.CreateSession:input_type -> auth0.logger0.control.v1.CreateSessionRequest
	3, // 3: auth0.logger0.control.v1.SessionEndpoint.CreateSession:output_type -> auth0.logger0.LogRecord
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_control_v1_session_endpoint_proto_init() }
func file_rpc_control_v1_session_endpoint_proto_init() {
	if File_rpc_control_v1_session_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_control_v1_session_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
			RawDescriptor: file_rpc_control_v1_session_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_control_v1_session_endpoint_proto_goTypes,
		DependencyIndexes: file_rpc_control_v1_session_endpoint_proto_depIdxs,
		MessageInfos:      file_rpc_control_v1_session_endpoint_proto_msgTypes,
	}.Build()
	File_rpc_control_v1_session_endpoint_proto = out.File
	file_rpc_control_v1_session_endpoint_proto_rawDesc = nil
	file_rpc_control_v1_session_endpoint_proto_goTypes = nil
	file_rpc_control_v1_session_endpoint_proto_depIdxs = nil
}