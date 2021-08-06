// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/ScheduleSign.proto

package proto

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

//
//Adds zero or more signing keys to a schedule. If the resulting set of signing keys satisfy the scheduled
//transaction's signing requirements, it will be executed immediately after the triggering <tt>ScheduleSign</tt>.
//
//Upon <tt>SUCCESS</tt>, the receipt includes the <tt>scheduledTransactionID</tt> to use to query for the
//record of the scheduled transaction's execution (if it occurs).
//
//Other notable response codes include <tt>INVALID_SCHEDULE_ID</tt>, <tt>SCHEDULE_WAS_DELETD</tt>,
//<tt>INVALID_ACCOUNT_ID</tt>, <tt>UNRESOLVABLE_REQUIRED_SIGNERS</tt>, <tt>SOME_SIGNATURES_WERE_INVALID</tt>,
//and <tt>NO_NEW_VALID_SIGNATURES</tt>. For more information please see the section of this
//documentation on the <tt>ResponseCode</tt> enum.
type ScheduleSignTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleID *ScheduleID `protobuf:"bytes,1,opt,name=scheduleID,proto3" json:"scheduleID,omitempty"` // The id of the schedule to add signing keys to
}

func (x *ScheduleSignTransactionBody) Reset() {
	*x = ScheduleSignTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ScheduleSign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleSignTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleSignTransactionBody) ProtoMessage() {}

func (x *ScheduleSignTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ScheduleSign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleSignTransactionBody.ProtoReflect.Descriptor instead.
func (*ScheduleSignTransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_ScheduleSign_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleSignTransactionBody) GetScheduleID() *ScheduleID {
	if x != nil {
		return x.ScheduleID
	}
	return nil
}

var File_proto_ScheduleSign_proto protoreflect.FileDescriptor

var file_proto_ScheduleSign_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x1b, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x52,
	0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x42, 0x4b, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ScheduleSign_proto_rawDescOnce sync.Once
	file_proto_ScheduleSign_proto_rawDescData = file_proto_ScheduleSign_proto_rawDesc
)

func file_proto_ScheduleSign_proto_rawDescGZIP() []byte {
	file_proto_ScheduleSign_proto_rawDescOnce.Do(func() {
		file_proto_ScheduleSign_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ScheduleSign_proto_rawDescData)
	})
	return file_proto_ScheduleSign_proto_rawDescData
}

var file_proto_ScheduleSign_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_ScheduleSign_proto_goTypes = []interface{}{
	(*ScheduleSignTransactionBody)(nil), // 0: proto.ScheduleSignTransactionBody
	(*ScheduleID)(nil),                  // 1: proto.ScheduleID
}
var file_proto_ScheduleSign_proto_depIdxs = []int32{
	1, // 0: proto.ScheduleSignTransactionBody.scheduleID:type_name -> proto.ScheduleID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_ScheduleSign_proto_init() }
func file_proto_ScheduleSign_proto_init() {
	if File_proto_ScheduleSign_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_ScheduleSign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleSignTransactionBody); i {
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
			RawDescriptor: file_proto_ScheduleSign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ScheduleSign_proto_goTypes,
		DependencyIndexes: file_proto_ScheduleSign_proto_depIdxs,
		MessageInfos:      file_proto_ScheduleSign_proto_msgTypes,
	}.Build()
	File_proto_ScheduleSign_proto = out.File
	file_proto_ScheduleSign_proto_rawDesc = nil
	file_proto_ScheduleSign_proto_goTypes = nil
	file_proto_ScheduleSign_proto_depIdxs = nil
}