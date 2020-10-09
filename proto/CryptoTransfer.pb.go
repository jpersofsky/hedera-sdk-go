// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/CryptoTransfer.proto

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

// A list of accounts and amounts to transfer out of each account (negative) or into it (positive).
type TransferList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountAmounts []*AccountAmount `protobuf:"bytes,1,rep,name=accountAmounts,proto3" json:"accountAmounts,omitempty"` // Multiple list of AccountAmount pairs, each of which has an account and an amount to transfer into it (positive) or out of it (negative)
}

func (x *TransferList) Reset() {
	*x = TransferList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CryptoTransfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferList) ProtoMessage() {}

func (x *TransferList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CryptoTransfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferList.ProtoReflect.Descriptor instead.
func (*TransferList) Descriptor() ([]byte, []int) {
	return file_proto_CryptoTransfer_proto_rawDescGZIP(), []int{0}
}

func (x *TransferList) GetAccountAmounts() []*AccountAmount {
	if x != nil {
		return x.AccountAmounts
	}
	return nil
}

// Transfer cryptocurrency from some accounts to other accounts. The accounts list can contain up to 10 accounts. The amounts list must be the same length as the accounts list. Each negative amount is withdrawn from the corresponding account (a sender), and each positive one is added to the corresponding account (a receiver). The amounts list must sum to zero. Each amount is a number of tinyBars (there are 100,000,000 tinyBars in one Hbar).
//If any sender account fails to have sufficient hbars, then the entire transaction fails, and none of those transfers occur, though the transaction fee is still charged. This transaction must be signed by the keys for all the sending accounts, and for any receiving accounts that have receiverSigRequired == true. The signatures are in the same order as the accounts, skipping those accounts that don't need a signature.
type CryptoTransferTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transfers *TransferList `protobuf:"bytes,1,opt,name=transfers,proto3" json:"transfers,omitempty"` // Accounts and amounts to transfer
}

func (x *CryptoTransferTransactionBody) Reset() {
	*x = CryptoTransferTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CryptoTransfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoTransferTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoTransferTransactionBody) ProtoMessage() {}

func (x *CryptoTransferTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CryptoTransfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoTransferTransactionBody.ProtoReflect.Descriptor instead.
func (*CryptoTransferTransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_CryptoTransfer_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoTransferTransactionBody) GetTransfers() *TransferList {
	if x != nil {
		return x.Transfers
	}
	return nil
}

var File_proto_CryptoTransfer_proto protoreflect.FileDescriptor

var file_proto_CryptoTransfer_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x52, 0x0a, 0x1d, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x31, 0x0a, 0x09, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x42, 0x4a, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a,
	0x61, 0x76, 0x61, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_CryptoTransfer_proto_rawDescOnce sync.Once
	file_proto_CryptoTransfer_proto_rawDescData = file_proto_CryptoTransfer_proto_rawDesc
)

func file_proto_CryptoTransfer_proto_rawDescGZIP() []byte {
	file_proto_CryptoTransfer_proto_rawDescOnce.Do(func() {
		file_proto_CryptoTransfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_CryptoTransfer_proto_rawDescData)
	})
	return file_proto_CryptoTransfer_proto_rawDescData
}

var file_proto_CryptoTransfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_CryptoTransfer_proto_goTypes = []interface{}{
	(*TransferList)(nil),                  // 0: proto.TransferList
	(*CryptoTransferTransactionBody)(nil), // 1: proto.CryptoTransferTransactionBody
	(*AccountAmount)(nil),                 // 2: proto.AccountAmount
}
var file_proto_CryptoTransfer_proto_depIdxs = []int32{
	2, // 0: proto.TransferList.accountAmounts:type_name -> proto.AccountAmount
	0, // 1: proto.CryptoTransferTransactionBody.transfers:type_name -> proto.TransferList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_CryptoTransfer_proto_init() }
func file_proto_CryptoTransfer_proto_init() {
	if File_proto_CryptoTransfer_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_CryptoTransfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferList); i {
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
		file_proto_CryptoTransfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoTransferTransactionBody); i {
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
			RawDescriptor: file_proto_CryptoTransfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_CryptoTransfer_proto_goTypes,
		DependencyIndexes: file_proto_CryptoTransfer_proto_depIdxs,
		MessageInfos:      file_proto_CryptoTransfer_proto_msgTypes,
	}.Build()
	File_proto_CryptoTransfer_proto = out.File
	file_proto_CryptoTransfer_proto_rawDesc = nil
	file_proto_CryptoTransfer_proto_goTypes = nil
	file_proto_CryptoTransfer_proto_depIdxs = nil
}
