// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: unused.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Logs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Logs) Reset() {
	*x = Logs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unused_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logs) ProtoMessage() {}

func (x *Logs) ProtoReflect() protoreflect.Message {
	mi := &file_unused_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logs.ProtoReflect.Descriptor instead.
func (*Logs) Descriptor() ([]byte, []int) {
	return file_unused_proto_rawDescGZIP(), []int{0}
}

type Logs_Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stream_id is the stream ID to connect to to get access to the
	// LogViewer service.
	StreamId uint32 `protobuf:"varint,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
}

func (x *Logs_Resp) Reset() {
	*x = Logs_Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unused_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logs_Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logs_Resp) ProtoMessage() {}

func (x *Logs_Resp) ProtoReflect() protoreflect.Message {
	mi := &file_unused_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logs_Resp.ProtoReflect.Descriptor instead.
func (*Logs_Resp) Descriptor() ([]byte, []int) {
	return file_unused_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Logs_Resp) GetStreamId() uint32 {
	if x != nil {
		return x.StreamId
	}
	return 0
}

type Logs_NextBatchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Logs_Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Logs_NextBatchResp) Reset() {
	*x = Logs_NextBatchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unused_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logs_NextBatchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logs_NextBatchResp) ProtoMessage() {}

func (x *Logs_NextBatchResp) ProtoReflect() protoreflect.Message {
	mi := &file_unused_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logs_NextBatchResp.ProtoReflect.Descriptor instead.
func (*Logs_NextBatchResp) Descriptor() ([]byte, []int) {
	return file_unused_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Logs_NextBatchResp) GetEvents() []*Logs_Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Logs_Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partition string               `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Contents  string               `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *Logs_Event) Reset() {
	*x = Logs_Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unused_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logs_Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logs_Event) ProtoMessage() {}

func (x *Logs_Event) ProtoReflect() protoreflect.Message {
	mi := &file_unused_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logs_Event.ProtoReflect.Descriptor instead.
func (*Logs_Event) Descriptor() ([]byte, []int) {
	return file_unused_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Logs_Event) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *Logs_Event) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Logs_Event) GetContents() string {
	if x != nil {
		return x.Contents
	}
	return ""
}

var File_unused_proto protoreflect.FileDescriptor

var file_unused_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x23, 0x0a, 0x04, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x1a, 0x4b, 0x0a, 0x0d, 0x4e, 0x65, 0x78, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3a, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61,
	0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x7b, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x0b, 0x4c,
	0x6f, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x44, 0x0a, 0x08, 0x4c, 0x6f,
	0x67, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x50, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x25, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x41, 0x72, 0x67, 0x73, 0x1a,
	0x21, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x32, 0x5f, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x12,
	0x52, 0x0a, 0x0c, 0x4e, 0x65, 0x78, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_unused_proto_rawDescOnce sync.Once
	file_unused_proto_rawDescData = file_unused_proto_rawDesc
)

func file_unused_proto_rawDescGZIP() []byte {
	file_unused_proto_rawDescOnce.Do(func() {
		file_unused_proto_rawDescData = protoimpl.X.CompressGZIP(file_unused_proto_rawDescData)
	})
	return file_unused_proto_rawDescData
}

var file_unused_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_unused_proto_goTypes = []interface{}{
	(*Logs)(nil),                // 0: hashicorp.waypoint.sdk.Logs
	(*Logs_Resp)(nil),           // 1: hashicorp.waypoint.sdk.Logs.Resp
	(*Logs_NextBatchResp)(nil),  // 2: hashicorp.waypoint.sdk.Logs.NextBatchResp
	(*Logs_Event)(nil),          // 3: hashicorp.waypoint.sdk.Logs.Event
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 5: google.protobuf.Empty
	(*FuncSpec_Args)(nil),       // 6: hashicorp.waypoint.sdk.FuncSpec.Args
	(*FuncSpec)(nil),            // 7: hashicorp.waypoint.sdk.FuncSpec
}
var file_unused_proto_depIdxs = []int32{
	3, // 0: hashicorp.waypoint.sdk.Logs.NextBatchResp.events:type_name -> hashicorp.waypoint.sdk.Logs.Event
	4, // 1: hashicorp.waypoint.sdk.Logs.Event.timestamp:type_name -> google.protobuf.Timestamp
	5, // 2: hashicorp.waypoint.sdk.LogPlatform.LogsSpec:input_type -> google.protobuf.Empty
	6, // 3: hashicorp.waypoint.sdk.LogPlatform.Logs:input_type -> hashicorp.waypoint.sdk.FuncSpec.Args
	5, // 4: hashicorp.waypoint.sdk.LogViewer.NextLogBatch:input_type -> google.protobuf.Empty
	7, // 5: hashicorp.waypoint.sdk.LogPlatform.LogsSpec:output_type -> hashicorp.waypoint.sdk.FuncSpec
	1, // 6: hashicorp.waypoint.sdk.LogPlatform.Logs:output_type -> hashicorp.waypoint.sdk.Logs.Resp
	2, // 7: hashicorp.waypoint.sdk.LogViewer.NextLogBatch:output_type -> hashicorp.waypoint.sdk.Logs.NextBatchResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_unused_proto_init() }
func file_unused_proto_init() {
	if File_unused_proto != nil {
		return
	}
	file_plugin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_unused_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logs); i {
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
		file_unused_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logs_Resp); i {
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
		file_unused_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logs_NextBatchResp); i {
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
		file_unused_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logs_Event); i {
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
			RawDescriptor: file_unused_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_unused_proto_goTypes,
		DependencyIndexes: file_unused_proto_depIdxs,
		MessageInfos:      file_unused_proto_msgTypes,
	}.Build()
	File_unused_proto = out.File
	file_unused_proto_rawDesc = nil
	file_unused_proto_goTypes = nil
	file_unused_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogPlatformClient is the client API for LogPlatform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogPlatformClient interface {
	LogsSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	Logs(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Logs_Resp, error)
}

type logPlatformClient struct {
	cc grpc.ClientConnInterface
}

func NewLogPlatformClient(cc grpc.ClientConnInterface) LogPlatformClient {
	return &logPlatformClient{cc}
}

func (c *logPlatformClient) LogsSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint.sdk.LogPlatform/LogsSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logPlatformClient) Logs(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Logs_Resp, error) {
	out := new(Logs_Resp)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint.sdk.LogPlatform/Logs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogPlatformServer is the server API for LogPlatform service.
type LogPlatformServer interface {
	LogsSpec(context.Context, *empty.Empty) (*FuncSpec, error)
	Logs(context.Context, *FuncSpec_Args) (*Logs_Resp, error)
}

// UnimplementedLogPlatformServer can be embedded to have forward compatible implementations.
type UnimplementedLogPlatformServer struct {
}

func (*UnimplementedLogPlatformServer) LogsSpec(context.Context, *empty.Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogsSpec not implemented")
}
func (*UnimplementedLogPlatformServer) Logs(context.Context, *FuncSpec_Args) (*Logs_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logs not implemented")
}

func RegisterLogPlatformServer(s *grpc.Server, srv LogPlatformServer) {
	s.RegisterService(&_LogPlatform_serviceDesc, srv)
}

func _LogPlatform_LogsSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogPlatformServer).LogsSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint.sdk.LogPlatform/LogsSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogPlatformServer).LogsSpec(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogPlatform_Logs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogPlatformServer).Logs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint.sdk.LogPlatform/Logs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogPlatformServer).Logs(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogPlatform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.waypoint.sdk.LogPlatform",
	HandlerType: (*LogPlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogsSpec",
			Handler:    _LogPlatform_LogsSpec_Handler,
		},
		{
			MethodName: "Logs",
			Handler:    _LogPlatform_Logs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unused.proto",
}

// LogViewerClient is the client API for LogViewer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogViewerClient interface {
	NextLogBatch(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Logs_NextBatchResp, error)
}

type logViewerClient struct {
	cc grpc.ClientConnInterface
}

func NewLogViewerClient(cc grpc.ClientConnInterface) LogViewerClient {
	return &logViewerClient{cc}
}

func (c *logViewerClient) NextLogBatch(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Logs_NextBatchResp, error) {
	out := new(Logs_NextBatchResp)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint.sdk.LogViewer/NextLogBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogViewerServer is the server API for LogViewer service.
type LogViewerServer interface {
	NextLogBatch(context.Context, *empty.Empty) (*Logs_NextBatchResp, error)
}

// UnimplementedLogViewerServer can be embedded to have forward compatible implementations.
type UnimplementedLogViewerServer struct {
}

func (*UnimplementedLogViewerServer) NextLogBatch(context.Context, *empty.Empty) (*Logs_NextBatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextLogBatch not implemented")
}

func RegisterLogViewerServer(s *grpc.Server, srv LogViewerServer) {
	s.RegisterService(&_LogViewer_serviceDesc, srv)
}

func _LogViewer_NextLogBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogViewerServer).NextLogBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint.sdk.LogViewer/NextLogBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogViewerServer).NextLogBatch(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogViewer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.waypoint.sdk.LogViewer",
	HandlerType: (*LogViewerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NextLogBatch",
			Handler:    _LogViewer_NextLogBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unused.proto",
}