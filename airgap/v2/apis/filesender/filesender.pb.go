// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: apis/filesender/filesender.proto

package filesender

import (
	v1 "github.com/redhat-marketplace/redhat-marketplace-operator/airgap/v2/apis/model/v1"
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

// streams a file identified by
type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UploadFileRequest_Info
	//	*UploadFileRequest_ChunkData
	Data isUploadFileRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_filesender_filesender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_filesender_filesender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_apis_filesender_filesender_proto_rawDescGZIP(), []int{0}
}

func (m *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadFileRequest) GetInfo() *v1.FileInfo {
	if x, ok := x.GetData().(*UploadFileRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadFileRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UploadFileRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_Info struct {
	Info *v1.FileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadFileRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadFileRequest_Info) isUploadFileRequest_Data() {}

func (*UploadFileRequest_ChunkData) isUploadFileRequest_Data() {}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId *v1.FileID `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Size   uint32     `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_filesender_filesender_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_filesender_filesender_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_apis_filesender_filesender_proto_rawDescGZIP(), []int{1}
}

func (x *UploadFileResponse) GetFileId() *v1.FileID {
	if x != nil {
		return x.FileId
	}
	return nil
}

func (x *UploadFileResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type UpdateFileMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId *v1.FileID `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *UpdateFileMetadataRequest) Reset() {
	*x = UpdateFileMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_filesender_filesender_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileMetadataRequest) ProtoMessage() {}

func (x *UpdateFileMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_filesender_filesender_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileMetadataRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileMetadataRequest) Descriptor() ([]byte, []int) {
	return file_apis_filesender_filesender_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFileMetadataRequest) GetFileId() *v1.FileID {
	if x != nil {
		return x.FileId
	}
	return nil
}

type UpdateFileMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId *v1.FileID `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *UpdateFileMetadataResponse) Reset() {
	*x = UpdateFileMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_filesender_filesender_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileMetadataResponse) ProtoMessage() {}

func (x *UpdateFileMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_filesender_filesender_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileMetadataResponse.ProtoReflect.Descriptor instead.
func (*UpdateFileMetadataResponse) Descriptor() ([]byte, []int) {
	return file_apis_filesender_filesender_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFileMetadataResponse) GetFileId() *v1.FileID {
	if x != nil {
		return x.FileId
	}
	return nil
}

var File_apis_filesender_filesender_proto protoreflect.FileDescriptor

var file_apis_filesender_filesender_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66,
	0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a,
	0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x53, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x46, 0x0a, 0x19, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x32, 0xd0, 0x01, 0x0a,
	0x0a, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x0a, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x6b, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x8d, 0x01, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x64, 0x68, 0x61, 0x74, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x69, 0x72, 0x67, 0x61,
	0x70, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x0a, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x68, 0x61, 0x74, 0x2d, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x64, 0x68, 0x61, 0x74,
	0x2d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x69, 0x72, 0x67, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_filesender_filesender_proto_rawDescOnce sync.Once
	file_apis_filesender_filesender_proto_rawDescData = file_apis_filesender_filesender_proto_rawDesc
)

func file_apis_filesender_filesender_proto_rawDescGZIP() []byte {
	file_apis_filesender_filesender_proto_rawDescOnce.Do(func() {
		file_apis_filesender_filesender_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_filesender_filesender_proto_rawDescData)
	})
	return file_apis_filesender_filesender_proto_rawDescData
}

var file_apis_filesender_filesender_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apis_filesender_filesender_proto_goTypes = []interface{}{
	(*UploadFileRequest)(nil),          // 0: filesender.v1.UploadFileRequest
	(*UploadFileResponse)(nil),         // 1: filesender.v1.UploadFileResponse
	(*UpdateFileMetadataRequest)(nil),  // 2: filesender.v1.UpdateFileMetadataRequest
	(*UpdateFileMetadataResponse)(nil), // 3: filesender.v1.UpdateFileMetadataResponse
	(*v1.FileInfo)(nil),                // 4: model.v1.FileInfo
	(*v1.FileID)(nil),                  // 5: model.v1.FileID
}
var file_apis_filesender_filesender_proto_depIdxs = []int32{
	4, // 0: filesender.v1.UploadFileRequest.info:type_name -> model.v1.FileInfo
	5, // 1: filesender.v1.UploadFileResponse.file_id:type_name -> model.v1.FileID
	5, // 2: filesender.v1.UpdateFileMetadataRequest.file_id:type_name -> model.v1.FileID
	5, // 3: filesender.v1.UpdateFileMetadataResponse.file_id:type_name -> model.v1.FileID
	0, // 4: filesender.v1.FileSender.UploadFile:input_type -> filesender.v1.UploadFileRequest
	2, // 5: filesender.v1.FileSender.UpdateFileMetadata:input_type -> filesender.v1.UpdateFileMetadataRequest
	1, // 6: filesender.v1.FileSender.UploadFile:output_type -> filesender.v1.UploadFileResponse
	3, // 7: filesender.v1.FileSender.UpdateFileMetadata:output_type -> filesender.v1.UpdateFileMetadataResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apis_filesender_filesender_proto_init() }
func file_apis_filesender_filesender_proto_init() {
	if File_apis_filesender_filesender_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_filesender_filesender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_apis_filesender_filesender_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
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
		file_apis_filesender_filesender_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileMetadataRequest); i {
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
		file_apis_filesender_filesender_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileMetadataResponse); i {
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
	file_apis_filesender_filesender_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UploadFileRequest_Info)(nil),
		(*UploadFileRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_filesender_filesender_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_filesender_filesender_proto_goTypes,
		DependencyIndexes: file_apis_filesender_filesender_proto_depIdxs,
		MessageInfos:      file_apis_filesender_filesender_proto_msgTypes,
	}.Build()
	File_apis_filesender_filesender_proto = out.File
	file_apis_filesender_filesender_proto_rawDesc = nil
	file_apis_filesender_filesender_proto_goTypes = nil
	file_apis_filesender_filesender_proto_depIdxs = nil
}
