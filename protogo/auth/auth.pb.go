// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: seekret/proto/auth/auth.proto

package auth

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

type ClientCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *ClientCredentials) Reset() {
	*x = ClientCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seekret_proto_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientCredentials) ProtoMessage() {}

func (x *ClientCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_seekret_proto_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientCredentials.ProtoReflect.Descriptor instead.
func (*ClientCredentials) Descriptor() ([]byte, []int) {
	return file_seekret_proto_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ClientCredentials) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ClientCredentials) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type JWTDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *JWTDetails) Reset() {
	*x = JWTDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seekret_proto_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTDetails) ProtoMessage() {}

func (x *JWTDetails) ProtoReflect() protoreflect.Message {
	mi := &file_seekret_proto_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTDetails.ProtoReflect.Descriptor instead.
func (*JWTDetails) Descriptor() ([]byte, []int) {
	return file_seekret_proto_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *JWTDetails) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *JWTDetails) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type ClientAuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientCredentials *ClientCredentials `protobuf:"bytes,1,opt,name=client_credentials,json=clientCredentials,proto3" json:"client_credentials,omitempty"`
}

func (x *ClientAuthenticateRequest) Reset() {
	*x = ClientAuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seekret_proto_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAuthenticateRequest) ProtoMessage() {}

func (x *ClientAuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_seekret_proto_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAuthenticateRequest.ProtoReflect.Descriptor instead.
func (*ClientAuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_seekret_proto_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *ClientAuthenticateRequest) GetClientCredentials() *ClientCredentials {
	if x != nil {
		return x.ClientCredentials
	}
	return nil
}

type ClientAuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details *JWTDetails `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *ClientAuthenticateResponse) Reset() {
	*x = ClientAuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seekret_proto_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAuthenticateResponse) ProtoMessage() {}

func (x *ClientAuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seekret_proto_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAuthenticateResponse.ProtoReflect.Descriptor instead.
func (*ClientAuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_seekret_proto_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ClientAuthenticateResponse) GetDetails() *JWTDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_seekret_proto_auth_auth_proto protoreflect.FileDescriptor

var file_seekret_proto_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x65, 0x6b, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x73, 0x65, 0x65, 0x6b, 0x72, 0x65, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x22, 0x55, 0x0a,
	0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0x54, 0x0a, 0x0a, 0x4a, 0x57, 0x54, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6b, 0x0a, 0x19, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x65, 0x6b, 0x72, 0x65, 0x74, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x50, 0x0a, 0x1a, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x65, 0x6b, 0x72, 0x65, 0x74,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x57, 0x54, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32, 0x71, 0x0a, 0x04, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x69, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x65, 0x6b, 0x72, 0x65,
	0x74, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x73, 0x65, 0x65, 0x6b, 0x72, 0x65, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x65, 0x6b, 0x2d,
	0x72, 0x65, 0x74, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_seekret_proto_auth_auth_proto_rawDescOnce sync.Once
	file_seekret_proto_auth_auth_proto_rawDescData = file_seekret_proto_auth_auth_proto_rawDesc
)

func file_seekret_proto_auth_auth_proto_rawDescGZIP() []byte {
	file_seekret_proto_auth_auth_proto_rawDescOnce.Do(func() {
		file_seekret_proto_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_seekret_proto_auth_auth_proto_rawDescData)
	})
	return file_seekret_proto_auth_auth_proto_rawDescData
}

var file_seekret_proto_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_seekret_proto_auth_auth_proto_goTypes = []interface{}{
	(*ClientCredentials)(nil),          // 0: seekret.auth.ClientCredentials
	(*JWTDetails)(nil),                 // 1: seekret.auth.JWTDetails
	(*ClientAuthenticateRequest)(nil),  // 2: seekret.auth.ClientAuthenticateRequest
	(*ClientAuthenticateResponse)(nil), // 3: seekret.auth.ClientAuthenticateResponse
}
var file_seekret_proto_auth_auth_proto_depIdxs = []int32{
	0, // 0: seekret.auth.ClientAuthenticateRequest.client_credentials:type_name -> seekret.auth.ClientCredentials
	1, // 1: seekret.auth.ClientAuthenticateResponse.details:type_name -> seekret.auth.JWTDetails
	2, // 2: seekret.auth.Auth.ClientAuthenticate:input_type -> seekret.auth.ClientAuthenticateRequest
	3, // 3: seekret.auth.Auth.ClientAuthenticate:output_type -> seekret.auth.ClientAuthenticateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_seekret_proto_auth_auth_proto_init() }
func file_seekret_proto_auth_auth_proto_init() {
	if File_seekret_proto_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seekret_proto_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientCredentials); i {
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
		file_seekret_proto_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTDetails); i {
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
		file_seekret_proto_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAuthenticateRequest); i {
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
		file_seekret_proto_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAuthenticateResponse); i {
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
			RawDescriptor: file_seekret_proto_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seekret_proto_auth_auth_proto_goTypes,
		DependencyIndexes: file_seekret_proto_auth_auth_proto_depIdxs,
		MessageInfos:      file_seekret_proto_auth_auth_proto_msgTypes,
	}.Build()
	File_seekret_proto_auth_auth_proto = out.File
	file_seekret_proto_auth_auth_proto_rawDesc = nil
	file_seekret_proto_auth_auth_proto_goTypes = nil
	file_seekret_proto_auth_auth_proto_depIdxs = nil
}
