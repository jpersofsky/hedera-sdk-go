// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/CryptoGetLiveHash.proto

package hedera_sdk_go

import (
	proto "github.com/golang/protobuf/proto"
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

// Requests a livehash associated to an account.
type CryptoGetLiveHashQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard info sent from client to node, including the signed payment, and what kind of response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The account to which the livehash is associated
	AccountID *AccountID `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	// The SHA-384 data in the livehash
	Hash []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *CryptoGetLiveHashQuery) Reset() {
	*x = CryptoGetLiveHashQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CryptoGetLiveHash_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoGetLiveHashQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoGetLiveHashQuery) ProtoMessage() {}

func (x *CryptoGetLiveHashQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CryptoGetLiveHash_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoGetLiveHashQuery.ProtoReflect.Descriptor instead.
func (*CryptoGetLiveHashQuery) Descriptor() ([]byte, []int) {
	return file_proto_CryptoGetLiveHash_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoGetLiveHashQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CryptoGetLiveHashQuery) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *CryptoGetLiveHashQuery) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

// Returns the full livehash associated to an account, if it is present. Note that the only way to obtain a state proof exhibiting the absence of a livehash from an account is to retrieve a state proof of the entire account with its list of livehashes.
type CryptoGetLiveHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The livehash, if present
	LiveHash *LiveHash `protobuf:"bytes,2,opt,name=liveHash,proto3" json:"liveHash,omitempty"`
}

func (x *CryptoGetLiveHashResponse) Reset() {
	*x = CryptoGetLiveHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CryptoGetLiveHash_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoGetLiveHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoGetLiveHashResponse) ProtoMessage() {}

func (x *CryptoGetLiveHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CryptoGetLiveHash_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoGetLiveHashResponse.ProtoReflect.Descriptor instead.
func (*CryptoGetLiveHashResponse) Descriptor() ([]byte, []int) {
	return file_proto_CryptoGetLiveHash_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoGetLiveHashResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CryptoGetLiveHashResponse) GetLiveHash() *LiveHash {
	if x != nil {
		return x.LiveHash
	}
	return nil
}

var File_proto_CryptoGetLiveHash_proto protoreflect.FileDescriptor

var file_proto_CryptoGetLiveHash_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x77, 0x0a,
	0x19, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x69, 0x76,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x08, 0x6c, 0x69,
	0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x42, 0x4a, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d,
	0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_CryptoGetLiveHash_proto_rawDescOnce sync.Once
	file_proto_CryptoGetLiveHash_proto_rawDescData = file_proto_CryptoGetLiveHash_proto_rawDesc
)

func file_proto_CryptoGetLiveHash_proto_rawDescGZIP() []byte {
	file_proto_CryptoGetLiveHash_proto_rawDescOnce.Do(func() {
		file_proto_CryptoGetLiveHash_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_CryptoGetLiveHash_proto_rawDescData)
	})
	return file_proto_CryptoGetLiveHash_proto_rawDescData
}

var file_proto_CryptoGetLiveHash_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_CryptoGetLiveHash_proto_goTypes = []interface{}{
	(*CryptoGetLiveHashQuery)(nil),    // 0: proto.CryptoGetLiveHashQuery
	(*CryptoGetLiveHashResponse)(nil), // 1: proto.CryptoGetLiveHashResponse
	(*QueryHeader)(nil),               // 2: proto.QueryHeader
	(*AccountID)(nil),                 // 3: proto.AccountID
	(*ResponseHeader)(nil),            // 4: proto.ResponseHeader
	(*LiveHash)(nil),                  // 5: proto.LiveHash
}
var file_proto_CryptoGetLiveHash_proto_depIdxs = []int32{
	2, // 0: proto.CryptoGetLiveHashQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.CryptoGetLiveHashQuery.accountID:type_name -> proto.AccountID
	4, // 2: proto.CryptoGetLiveHashResponse.header:type_name -> proto.ResponseHeader
	5, // 3: proto.CryptoGetLiveHashResponse.liveHash:type_name -> proto.LiveHash
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_CryptoGetLiveHash_proto_init() }
func file_proto_CryptoGetLiveHash_proto_init() {
	if File_proto_CryptoGetLiveHash_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	file_proto_QueryHeader_proto_init()
	file_proto_ResponseHeader_proto_init()
	file_proto_CryptoAddLiveHash_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_CryptoGetLiveHash_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoGetLiveHashQuery); i {
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
		file_proto_CryptoGetLiveHash_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoGetLiveHashResponse); i {
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
			RawDescriptor: file_proto_CryptoGetLiveHash_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_CryptoGetLiveHash_proto_goTypes,
		DependencyIndexes: file_proto_CryptoGetLiveHash_proto_depIdxs,
		MessageInfos:      file_proto_CryptoGetLiveHash_proto_msgTypes,
	}.Build()
	File_proto_CryptoGetLiveHash_proto = out.File
	file_proto_CryptoGetLiveHash_proto_rawDesc = nil
	file_proto_CryptoGetLiveHash_proto_goTypes = nil
	file_proto_CryptoGetLiveHash_proto_depIdxs = nil
}
