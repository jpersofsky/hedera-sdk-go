// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/SystemUndelete.proto

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

//
//Undelete a file or smart contract that was deleted by SystemDelete; requires a Hedera administrative multisignature.
type SystemUndeleteTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Id:
	//	*SystemUndeleteTransactionBody_FileID
	//	*SystemUndeleteTransactionBody_ContractID
	Id isSystemUndeleteTransactionBody_Id `protobuf_oneof:"id"`
}

func (x *SystemUndeleteTransactionBody) Reset() {
	*x = SystemUndeleteTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_SystemUndelete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUndeleteTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUndeleteTransactionBody) ProtoMessage() {}

func (x *SystemUndeleteTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_SystemUndelete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUndeleteTransactionBody.ProtoReflect.Descriptor instead.
func (*SystemUndeleteTransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_SystemUndelete_proto_rawDescGZIP(), []int{0}
}

func (m *SystemUndeleteTransactionBody) GetId() isSystemUndeleteTransactionBody_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *SystemUndeleteTransactionBody) GetFileID() *FileID {
	if x, ok := x.GetId().(*SystemUndeleteTransactionBody_FileID); ok {
		return x.FileID
	}
	return nil
}

func (x *SystemUndeleteTransactionBody) GetContractID() *ContractID {
	if x, ok := x.GetId().(*SystemUndeleteTransactionBody_ContractID); ok {
		return x.ContractID
	}
	return nil
}

type isSystemUndeleteTransactionBody_Id interface {
	isSystemUndeleteTransactionBody_Id()
}

type SystemUndeleteTransactionBody_FileID struct {
	FileID *FileID `protobuf:"bytes,1,opt,name=fileID,proto3,oneof"` // The file ID to undelete, in the format used in transactions
}

type SystemUndeleteTransactionBody_ContractID struct {
	ContractID *ContractID `protobuf:"bytes,2,opt,name=contractID,proto3,oneof"` // The contract ID instance to undelete, in the format used in transactions
}

func (*SystemUndeleteTransactionBody_FileID) isSystemUndeleteTransactionBody_Id() {}

func (*SystemUndeleteTransactionBody_ContractID) isSystemUndeleteTransactionBody_Id() {}

var File_proto_SystemUndelete_proto protoreflect.FileDescriptor

var file_proto_SystemUndelete_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x6e,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x1d,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x27, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x48, 0x00, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x48, 0x00, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x42, 0x04, 0x0a, 0x02, 0x69,
	0x64, 0x42, 0x4a, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68,
	0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_SystemUndelete_proto_rawDescOnce sync.Once
	file_proto_SystemUndelete_proto_rawDescData = file_proto_SystemUndelete_proto_rawDesc
)

func file_proto_SystemUndelete_proto_rawDescGZIP() []byte {
	file_proto_SystemUndelete_proto_rawDescOnce.Do(func() {
		file_proto_SystemUndelete_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_SystemUndelete_proto_rawDescData)
	})
	return file_proto_SystemUndelete_proto_rawDescData
}

var file_proto_SystemUndelete_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_SystemUndelete_proto_goTypes = []interface{}{
	(*SystemUndeleteTransactionBody)(nil), // 0: proto.SystemUndeleteTransactionBody
	(*FileID)(nil),                        // 1: proto.FileID
	(*ContractID)(nil),                    // 2: proto.ContractID
}
var file_proto_SystemUndelete_proto_depIdxs = []int32{
	1, // 0: proto.SystemUndeleteTransactionBody.fileID:type_name -> proto.FileID
	2, // 1: proto.SystemUndeleteTransactionBody.contractID:type_name -> proto.ContractID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_SystemUndelete_proto_init() }
func file_proto_SystemUndelete_proto_init() {
	if File_proto_SystemUndelete_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_SystemUndelete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUndeleteTransactionBody); i {
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
	file_proto_SystemUndelete_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SystemUndeleteTransactionBody_FileID)(nil),
		(*SystemUndeleteTransactionBody_ContractID)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_SystemUndelete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_SystemUndelete_proto_goTypes,
		DependencyIndexes: file_proto_SystemUndelete_proto_depIdxs,
		MessageInfos:      file_proto_SystemUndelete_proto_msgTypes,
	}.Build()
	File_proto_SystemUndelete_proto = out.File
	file_proto_SystemUndelete_proto_rawDesc = nil
	file_proto_SystemUndelete_proto_goTypes = nil
	file_proto_SystemUndelete_proto_depIdxs = nil
}
