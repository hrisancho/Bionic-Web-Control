// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.6
// source: notifications/notifications.proto

package notifications

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

type NotificationType int32

const (
	NotificationType_connected    NotificationType = 0
	NotificationType_disconnected NotificationType = 1
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "connected",
		1: "disconnected",
	}
	NotificationType_value = map[string]int32{
		"connected":    0,
		"disconnected": 1,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_notifications_notifications_proto_enumTypes[0].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_notifications_notifications_proto_enumTypes[0]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{0}
}

var File_notifications_notifications_proto protoreflect.FileDescriptor

var file_notifications_notifications_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2a, 0x33, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x42, 0x69, 0x6f, 0x6e, 0x69,
	0x63, 0x2d, 0x57, 0x65, 0x62, 0x2d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifications_notifications_proto_rawDescOnce sync.Once
	file_notifications_notifications_proto_rawDescData = file_notifications_notifications_proto_rawDesc
)

func file_notifications_notifications_proto_rawDescGZIP() []byte {
	file_notifications_notifications_proto_rawDescOnce.Do(func() {
		file_notifications_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifications_notifications_proto_rawDescData)
	})
	return file_notifications_notifications_proto_rawDescData
}

var file_notifications_notifications_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notifications_notifications_proto_goTypes = []interface{}{
	(NotificationType)(0), // 0: Notifications.NotificationType
}
var file_notifications_notifications_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notifications_notifications_proto_init() }
func file_notifications_notifications_proto_init() {
	if File_notifications_notifications_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notifications_notifications_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notifications_notifications_proto_goTypes,
		DependencyIndexes: file_notifications_notifications_proto_depIdxs,
		EnumInfos:         file_notifications_notifications_proto_enumTypes,
	}.Build()
	File_notifications_notifications_proto = out.File
	file_notifications_notifications_proto_rawDesc = nil
	file_notifications_notifications_proto_goTypes = nil
	file_notifications_notifications_proto_depIdxs = nil
}
