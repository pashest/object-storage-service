// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.2
// source: api/storage/storage.proto

package storage

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

type UploadChunksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId string `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadChunksRequest) Reset() {
	*x = UploadChunksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadChunksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadChunksRequest) ProtoMessage() {}

func (x *UploadChunksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadChunksRequest.ProtoReflect.Descriptor instead.
func (*UploadChunksRequest) Descriptor() ([]byte, []int) {
	return file_api_storage_storage_proto_rawDescGZIP(), []int{0}
}

func (x *UploadChunksRequest) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *UploadChunksRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadChunksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UploadChunksResponse) Reset() {
	*x = UploadChunksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadChunksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadChunksResponse) ProtoMessage() {}

func (x *UploadChunksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadChunksResponse.ProtoReflect.Descriptor instead.
func (*UploadChunksResponse) Descriptor() ([]byte, []int) {
	return file_api_storage_storage_proto_rawDescGZIP(), []int{1}
}

func (x *UploadChunksResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UploadChunksResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetChunksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkIds []string `protobuf:"bytes,1,rep,name=chunk_ids,json=chunkIds,proto3" json:"chunk_ids,omitempty"`
}

func (x *GetChunksRequest) Reset() {
	*x = GetChunksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunksRequest) ProtoMessage() {}

func (x *GetChunksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunksRequest.ProtoReflect.Descriptor instead.
func (*GetChunksRequest) Descriptor() ([]byte, []int) {
	return file_api_storage_storage_proto_rawDescGZIP(), []int{2}
}

func (x *GetChunksRequest) GetChunkIds() []string {
	if x != nil {
		return x.ChunkIds
	}
	return nil
}

type GetChunksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkId string `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetChunksResponse) Reset() {
	*x = GetChunksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunksResponse) ProtoMessage() {}

func (x *GetChunksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunksResponse.ProtoReflect.Descriptor instead.
func (*GetChunksResponse) Descriptor() ([]byte, []int) {
	return file_api_storage_storage_proto_rawDescGZIP(), []int{3}
}

func (x *GetChunksResponse) GetChunkId() string {
	if x != nil {
		return x.ChunkId
	}
	return ""
}

func (x *GetChunksResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_storage_storage_proto protoreflect.FileDescriptor

var file_api_storage_storage_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x22, 0x44, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a, 0x14, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x73, 0x22, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xa5, 0x01, 0x0a, 0x0e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d,
	0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1c,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x44, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x61, 0x73, 0x68, 0x65, 0x73, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_storage_storage_proto_rawDescOnce sync.Once
	file_api_storage_storage_proto_rawDescData = file_api_storage_storage_proto_rawDesc
)

func file_api_storage_storage_proto_rawDescGZIP() []byte {
	file_api_storage_storage_proto_rawDescOnce.Do(func() {
		file_api_storage_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_storage_storage_proto_rawDescData)
	})
	return file_api_storage_storage_proto_rawDescData
}

var file_api_storage_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_storage_storage_proto_goTypes = []any{
	(*UploadChunksRequest)(nil),  // 0: storage.UploadChunksRequest
	(*UploadChunksResponse)(nil), // 1: storage.UploadChunksResponse
	(*GetChunksRequest)(nil),     // 2: storage.GetChunksRequest
	(*GetChunksResponse)(nil),    // 3: storage.GetChunksResponse
}
var file_api_storage_storage_proto_depIdxs = []int32{
	0, // 0: storage.StorageService.UploadChunks:input_type -> storage.UploadChunksRequest
	2, // 1: storage.StorageService.GetChunks:input_type -> storage.GetChunksRequest
	1, // 2: storage.StorageService.UploadChunks:output_type -> storage.UploadChunksResponse
	3, // 3: storage.StorageService.GetChunks:output_type -> storage.GetChunksResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_storage_storage_proto_init() }
func file_api_storage_storage_proto_init() {
	if File_api_storage_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_storage_storage_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UploadChunksRequest); i {
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
		file_api_storage_storage_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UploadChunksResponse); i {
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
		file_api_storage_storage_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetChunksRequest); i {
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
		file_api_storage_storage_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetChunksResponse); i {
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
			RawDescriptor: file_api_storage_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_storage_storage_proto_goTypes,
		DependencyIndexes: file_api_storage_storage_proto_depIdxs,
		MessageInfos:      file_api_storage_storage_proto_msgTypes,
	}.Build()
	File_api_storage_storage_proto = out.File
	file_api_storage_storage_proto_rawDesc = nil
	file_api_storage_storage_proto_goTypes = nil
	file_api_storage_storage_proto_depIdxs = nil
}
