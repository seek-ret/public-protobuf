// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: seekret/proto/api_gateway/api_gateway.proto

package api_gateway

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

type WorkspaceEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseDir       string   `protobuf:"bytes,1,opt,name=base_dir,json=baseDir,proto3" json:"base_dir,omitempty"`
	BucketName    string   `protobuf:"bytes,2,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	Endpoint      string   `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AccessKey     string   `protobuf:"bytes,4,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey     string   `protobuf:"bytes,5,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Insecure      bool     `protobuf:"varint,6,opt,name=insecure,proto3" json:"insecure,omitempty"`
	HostWhitelist []string `protobuf:"bytes,7,rep,name=host_whitelist,json=hostWhitelist,proto3" json:"host_whitelist,omitempty"`
}

func (x *WorkspaceEntry) Reset() {
	*x = WorkspaceEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seekret_proto_api_gateway_api_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceEntry) ProtoMessage() {}

func (x *WorkspaceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_seekret_proto_api_gateway_api_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceEntry.ProtoReflect.Descriptor instead.
func (*WorkspaceEntry) Descriptor() ([]byte, []int) {
	return file_seekret_proto_api_gateway_api_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *WorkspaceEntry) GetBaseDir() string {
	if x != nil {
		return x.BaseDir
	}
	return ""
}

func (x *WorkspaceEntry) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *WorkspaceEntry) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *WorkspaceEntry) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *WorkspaceEntry) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *WorkspaceEntry) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *WorkspaceEntry) GetHostWhitelist() []string {
	if x != nil {
		return x.HostWhitelist
	}
	return nil
}

type FetchWorkspaceConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workspace string `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *FetchWorkspaceConfigurationRequest) Reset() {
	*x = FetchWorkspaceConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seekret_proto_api_gateway_api_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchWorkspaceConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchWorkspaceConfigurationRequest) ProtoMessage() {}

func (x *FetchWorkspaceConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seekret_proto_api_gateway_api_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchWorkspaceConfigurationRequest.ProtoReflect.Descriptor instead.
func (*FetchWorkspaceConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_seekret_proto_api_gateway_api_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *FetchWorkspaceConfigurationRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

type FetchWorkspaceConfigurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *WorkspaceEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *FetchWorkspaceConfigurationResponse) Reset() {
	*x = FetchWorkspaceConfigurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seekret_proto_api_gateway_api_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchWorkspaceConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchWorkspaceConfigurationResponse) ProtoMessage() {}

func (x *FetchWorkspaceConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seekret_proto_api_gateway_api_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchWorkspaceConfigurationResponse.ProtoReflect.Descriptor instead.
func (*FetchWorkspaceConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_seekret_proto_api_gateway_api_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *FetchWorkspaceConfigurationResponse) GetEntry() *WorkspaceEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

var File_seekret_proto_api_gateway_api_gateway_proto protoreflect.FileDescriptor

var file_seekret_proto_api_gateway_api_gateway_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x65, 0x65, 0x6b, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73,
	0x65, 0x65, 0x6b, 0x72, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x22, 0xe9, 0x01, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x69,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x44, 0x69, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x68, 0x6f, 0x73, 0x74, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x42,
	0x0a, 0x22, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0x60, 0x0a, 0x23, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x65, 0x6b, 0x72,
	0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x32, 0x9f, 0x01, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x12, 0x90, 0x01, 0x0a, 0x1b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x37, 0x2e, 0x73, 0x65, 0x65, 0x6b, 0x72, 0x65, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73,
	0x65, 0x65, 0x6b, 0x72, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x65, 0x6b, 0x2d, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_seekret_proto_api_gateway_api_gateway_proto_rawDescOnce sync.Once
	file_seekret_proto_api_gateway_api_gateway_proto_rawDescData = file_seekret_proto_api_gateway_api_gateway_proto_rawDesc
)

func file_seekret_proto_api_gateway_api_gateway_proto_rawDescGZIP() []byte {
	file_seekret_proto_api_gateway_api_gateway_proto_rawDescOnce.Do(func() {
		file_seekret_proto_api_gateway_api_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_seekret_proto_api_gateway_api_gateway_proto_rawDescData)
	})
	return file_seekret_proto_api_gateway_api_gateway_proto_rawDescData
}

var file_seekret_proto_api_gateway_api_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_seekret_proto_api_gateway_api_gateway_proto_goTypes = []interface{}{
	(*WorkspaceEntry)(nil),                      // 0: seekret.api_gateway.WorkspaceEntry
	(*FetchWorkspaceConfigurationRequest)(nil),  // 1: seekret.api_gateway.FetchWorkspaceConfigurationRequest
	(*FetchWorkspaceConfigurationResponse)(nil), // 2: seekret.api_gateway.FetchWorkspaceConfigurationResponse
}
var file_seekret_proto_api_gateway_api_gateway_proto_depIdxs = []int32{
	0, // 0: seekret.api_gateway.FetchWorkspaceConfigurationResponse.entry:type_name -> seekret.api_gateway.WorkspaceEntry
	1, // 1: seekret.api_gateway.ApiGateway.FetchWorkspaceConfiguration:input_type -> seekret.api_gateway.FetchWorkspaceConfigurationRequest
	2, // 2: seekret.api_gateway.ApiGateway.FetchWorkspaceConfiguration:output_type -> seekret.api_gateway.FetchWorkspaceConfigurationResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_seekret_proto_api_gateway_api_gateway_proto_init() }
func file_seekret_proto_api_gateway_api_gateway_proto_init() {
	if File_seekret_proto_api_gateway_api_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seekret_proto_api_gateway_api_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceEntry); i {
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
		file_seekret_proto_api_gateway_api_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchWorkspaceConfigurationRequest); i {
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
		file_seekret_proto_api_gateway_api_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchWorkspaceConfigurationResponse); i {
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
			RawDescriptor: file_seekret_proto_api_gateway_api_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seekret_proto_api_gateway_api_gateway_proto_goTypes,
		DependencyIndexes: file_seekret_proto_api_gateway_api_gateway_proto_depIdxs,
		MessageInfos:      file_seekret_proto_api_gateway_api_gateway_proto_msgTypes,
	}.Build()
	File_seekret_proto_api_gateway_api_gateway_proto = out.File
	file_seekret_proto_api_gateway_api_gateway_proto_rawDesc = nil
	file_seekret_proto_api_gateway_api_gateway_proto_goTypes = nil
	file_seekret_proto_api_gateway_api_gateway_proto_depIdxs = nil
}
