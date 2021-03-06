//
//protoc --go_out=. --go_opt=paths=source_relative \
//--go-grpc_out=. --go-grpc_opt=paths=source_relative \
//grpc/cache.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: grpc/cache.proto

package __

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

type GetKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetKeyRequest) Reset() {
	*x = GetKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyRequest) ProtoMessage() {}

func (x *GetKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyRequest.ProtoReflect.Descriptor instead.
func (*GetKeyRequest) Descriptor() ([]byte, []int) {
	return file_grpc_cache_proto_rawDescGZIP(), []int{0}
}

func (x *GetKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetKeyReply) Reset() {
	*x = GetKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyReply) ProtoMessage() {}

func (x *GetKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyReply.ProtoReflect.Descriptor instead.
func (*GetKeyReply) Descriptor() ([]byte, []int) {
	return file_grpc_cache_proto_rawDescGZIP(), []int{1}
}

func (x *GetKeyReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetKeyRequest) Reset() {
	*x = SetKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetKeyRequest) ProtoMessage() {}

func (x *SetKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetKeyRequest.ProtoReflect.Descriptor instead.
func (*SetKeyRequest) Descriptor() ([]byte, []int) {
	return file_grpc_cache_proto_rawDescGZIP(), []int{2}
}

func (x *SetKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetKeyRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetKeyReply) Reset() {
	*x = SetKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetKeyReply) ProtoMessage() {}

func (x *SetKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetKeyReply.ProtoReflect.Descriptor instead.
func (*SetKeyReply) Descriptor() ([]byte, []int) {
	return file_grpc_cache_proto_rawDescGZIP(), []int{3}
}

type ClearRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearRequest) Reset() {
	*x = ClearRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearRequest) ProtoMessage() {}

func (x *ClearRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearRequest.ProtoReflect.Descriptor instead.
func (*ClearRequest) Descriptor() ([]byte, []int) {
	return file_grpc_cache_proto_rawDescGZIP(), []int{4}
}

type ClearReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearReply) Reset() {
	*x = ClearReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_cache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearReply) ProtoMessage() {}

func (x *ClearReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_cache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearReply.ProtoReflect.Descriptor instead.
func (*ClearReply) Descriptor() ([]byte, []int) {
	return file_grpc_cache_proto_rawDescGZIP(), []int{5}
}

type RemoveKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RemoveKeyRequest) Reset() {
	*x = RemoveKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_cache_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveKeyRequest) ProtoMessage() {}

func (x *RemoveKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_cache_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveKeyRequest.ProtoReflect.Descriptor instead.
func (*RemoveKeyRequest) Descriptor() ([]byte, []int) {
	return file_grpc_cache_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RemoveKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveKeyReply) Reset() {
	*x = RemoveKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_cache_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveKeyReply) ProtoMessage() {}

func (x *RemoveKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_cache_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveKeyReply.ProtoReflect.Descriptor instead.
func (*RemoveKeyReply) Descriptor() ([]byte, []int) {
	return file_grpc_cache_proto_rawDescGZIP(), []int{7}
}

var File_grpc_cache_proto protoreflect.FileDescriptor

var file_grpc_cache_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x23, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x37, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x10, 0x0a,
	0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0xe9, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x14, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x53,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x43, 0x6c,
	0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_cache_proto_rawDescOnce sync.Once
	file_grpc_cache_proto_rawDescData = file_grpc_cache_proto_rawDesc
)

func file_grpc_cache_proto_rawDescGZIP() []byte {
	file_grpc_cache_proto_rawDescOnce.Do(func() {
		file_grpc_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_cache_proto_rawDescData)
	})
	return file_grpc_cache_proto_rawDescData
}

var file_grpc_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_cache_proto_goTypes = []interface{}{
	(*GetKeyRequest)(nil),    // 0: cache.GetKeyRequest
	(*GetKeyReply)(nil),      // 1: cache.GetKeyReply
	(*SetKeyRequest)(nil),    // 2: cache.SetKeyRequest
	(*SetKeyReply)(nil),      // 3: cache.SetKeyReply
	(*ClearRequest)(nil),     // 4: cache.ClearRequest
	(*ClearReply)(nil),       // 5: cache.ClearReply
	(*RemoveKeyRequest)(nil), // 6: cache.RemoveKeyRequest
	(*RemoveKeyReply)(nil),   // 7: cache.RemoveKeyReply
}
var file_grpc_cache_proto_depIdxs = []int32{
	0, // 0: cache.CacheHandler.GetKey:input_type -> cache.GetKeyRequest
	2, // 1: cache.CacheHandler.SetKey:input_type -> cache.SetKeyRequest
	4, // 2: cache.CacheHandler.Clear:input_type -> cache.ClearRequest
	6, // 3: cache.CacheHandler.Remove:input_type -> cache.RemoveKeyRequest
	1, // 4: cache.CacheHandler.GetKey:output_type -> cache.GetKeyReply
	3, // 5: cache.CacheHandler.SetKey:output_type -> cache.SetKeyReply
	5, // 6: cache.CacheHandler.Clear:output_type -> cache.ClearReply
	7, // 7: cache.CacheHandler.Remove:output_type -> cache.RemoveKeyReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_cache_proto_init() }
func file_grpc_cache_proto_init() {
	if File_grpc_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyRequest); i {
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
		file_grpc_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyReply); i {
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
		file_grpc_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetKeyRequest); i {
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
		file_grpc_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetKeyReply); i {
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
		file_grpc_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearRequest); i {
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
		file_grpc_cache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearReply); i {
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
		file_grpc_cache_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveKeyRequest); i {
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
		file_grpc_cache_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveKeyReply); i {
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
			RawDescriptor: file_grpc_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_cache_proto_goTypes,
		DependencyIndexes: file_grpc_cache_proto_depIdxs,
		MessageInfos:      file_grpc_cache_proto_msgTypes,
	}.Build()
	File_grpc_cache_proto = out.File
	file_grpc_cache_proto_rawDesc = nil
	file_grpc_cache_proto_goTypes = nil
	file_grpc_cache_proto_depIdxs = nil
}
