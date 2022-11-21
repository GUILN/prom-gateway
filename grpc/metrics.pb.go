// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: grpc/metrics.proto

package protocol

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

type CreateCounterResponse_CreateCounterResult int32

const (
	CreateCounterResponse_SUCCESS CreateCounterResponse_CreateCounterResult = 0
	CreateCounterResponse_FAILED  CreateCounterResponse_CreateCounterResult = 1
)

// Enum value maps for CreateCounterResponse_CreateCounterResult.
var (
	CreateCounterResponse_CreateCounterResult_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
	}
	CreateCounterResponse_CreateCounterResult_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
	}
)

func (x CreateCounterResponse_CreateCounterResult) Enum() *CreateCounterResponse_CreateCounterResult {
	p := new(CreateCounterResponse_CreateCounterResult)
	*p = x
	return p
}

func (x CreateCounterResponse_CreateCounterResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateCounterResponse_CreateCounterResult) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_metrics_proto_enumTypes[0].Descriptor()
}

func (CreateCounterResponse_CreateCounterResult) Type() protoreflect.EnumType {
	return &file_grpc_metrics_proto_enumTypes[0]
}

func (x CreateCounterResponse_CreateCounterResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateCounterResponse_CreateCounterResult.Descriptor instead.
func (CreateCounterResponse_CreateCounterResult) EnumDescriptor() ([]byte, []int) {
	return file_grpc_metrics_proto_rawDescGZIP(), []int{1, 0}
}

type CreateCounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CounterName string  `protobuf:"bytes,1,opt,name=counter_name,json=counterName,proto3" json:"counter_name,omitempty"`
	CounterHelp *string `protobuf:"bytes,2,opt,name=counter_help,json=counterHelp,proto3,oneof" json:"counter_help,omitempty"`
}

func (x *CreateCounterRequest) Reset() {
	*x = CreateCounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCounterRequest) ProtoMessage() {}

func (x *CreateCounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCounterRequest.ProtoReflect.Descriptor instead.
func (*CreateCounterRequest) Descriptor() ([]byte, []int) {
	return file_grpc_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCounterRequest) GetCounterName() string {
	if x != nil {
		return x.CounterName
	}
	return ""
}

func (x *CreateCounterRequest) GetCounterHelp() string {
	if x != nil && x.CounterHelp != nil {
		return *x.CounterHelp
	}
	return ""
}

type CreateCounterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result CreateCounterResponse_CreateCounterResult `protobuf:"varint,1,opt,name=result,proto3,enum=CreateCounterResponse_CreateCounterResult" json:"result,omitempty"`
}

func (x *CreateCounterResponse) Reset() {
	*x = CreateCounterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCounterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCounterResponse) ProtoMessage() {}

func (x *CreateCounterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCounterResponse.ProtoReflect.Descriptor instead.
func (*CreateCounterResponse) Descriptor() ([]byte, []int) {
	return file_grpc_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCounterResponse) GetResult() CreateCounterResponse_CreateCounterResult {
	if x != nil {
		return x.Result
	}
	return CreateCounterResponse_SUCCESS
}

var File_grpc_metrics_proto protoreflect.FileDescriptor

var file_grpc_metrics_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x48, 0x65, 0x6c, 0x70, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x22, 0x8b, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x32, 0x49, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x75, 0x69, 0x6c, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_metrics_proto_rawDescOnce sync.Once
	file_grpc_metrics_proto_rawDescData = file_grpc_metrics_proto_rawDesc
)

func file_grpc_metrics_proto_rawDescGZIP() []byte {
	file_grpc_metrics_proto_rawDescOnce.Do(func() {
		file_grpc_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_metrics_proto_rawDescData)
	})
	return file_grpc_metrics_proto_rawDescData
}

var file_grpc_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_metrics_proto_goTypes = []interface{}{
	(CreateCounterResponse_CreateCounterResult)(0), // 0: CreateCounterResponse.CreateCounterResult
	(*CreateCounterRequest)(nil),                   // 1: CreateCounterRequest
	(*CreateCounterResponse)(nil),                  // 2: CreateCounterResponse
}
var file_grpc_metrics_proto_depIdxs = []int32{
	0, // 0: CreateCounterResponse.result:type_name -> CreateCounterResponse.CreateCounterResult
	1, // 1: Metrics.CreateCounter:input_type -> CreateCounterRequest
	2, // 2: Metrics.CreateCounter:output_type -> CreateCounterResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_metrics_proto_init() }
func file_grpc_metrics_proto_init() {
	if File_grpc_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCounterRequest); i {
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
		file_grpc_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCounterResponse); i {
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
	file_grpc_metrics_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_metrics_proto_goTypes,
		DependencyIndexes: file_grpc_metrics_proto_depIdxs,
		EnumInfos:         file_grpc_metrics_proto_enumTypes,
		MessageInfos:      file_grpc_metrics_proto_msgTypes,
	}.Build()
	File_grpc_metrics_proto = out.File
	file_grpc_metrics_proto_rawDesc = nil
	file_grpc_metrics_proto_goTypes = nil
	file_grpc_metrics_proto_depIdxs = nil
}
