// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.6
// source: straingauge/straingauge.proto

package staingauge

import (
	shared "Bionic-Web-Control/proto/shared"
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

// датчики давления на кончиках пальцев
type StrainGuage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Finger        shared.Finger `protobuf:"varint,1,opt,name=finger,proto3,enum=Shared.Finger" json:"finger,omitempty"` //палец, на котором расположен датчик
	Pressure      uint32        `protobuf:"varint,2,opt,name=pressure,proto3" json:"pressure,omitempty"`                //измеренное давление
	ConnectionPin string        `protobuf:"bytes,3,opt,name=connectionPin,proto3" json:"connectionPin,omitempty"`       // Пин (порт на плате) к которому будет крепиться тензодатчик
}

func (x *StrainGuage) Reset() {
	*x = StrainGuage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_straingauge_straingauge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrainGuage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrainGuage) ProtoMessage() {}

func (x *StrainGuage) ProtoReflect() protoreflect.Message {
	mi := &file_straingauge_straingauge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrainGuage.ProtoReflect.Descriptor instead.
func (*StrainGuage) Descriptor() ([]byte, []int) {
	return file_straingauge_straingauge_proto_rawDescGZIP(), []int{0}
}

func (x *StrainGuage) GetFinger() shared.Finger {
	if x != nil {
		return x.Finger
	}
	return shared.Finger(0)
}

func (x *StrainGuage) GetPressure() uint32 {
	if x != nil {
		return x.Pressure
	}
	return 0
}

func (x *StrainGuage) GetConnectionPin() string {
	if x != nil {
		return x.ConnectionPin
	}
	return ""
}

var File_straingauge_straingauge_proto protoreflect.FileDescriptor

var file_straingauge_straingauge_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2f, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x53, 0x74, 0x61, 0x69, 0x6e, 0x67, 0x61, 0x75, 0x67, 0x65, 0x1a, 0x13, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x77, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x47, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x52,
	0x06, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x6e, 0x42, 0x25, 0x5a, 0x23, 0x42, 0x69, 0x6f,
	0x6e, 0x69, 0x63, 0x2d, 0x57, 0x65, 0x62, 0x2d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x67, 0x61, 0x75, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_straingauge_straingauge_proto_rawDescOnce sync.Once
	file_straingauge_straingauge_proto_rawDescData = file_straingauge_straingauge_proto_rawDesc
)

func file_straingauge_straingauge_proto_rawDescGZIP() []byte {
	file_straingauge_straingauge_proto_rawDescOnce.Do(func() {
		file_straingauge_straingauge_proto_rawDescData = protoimpl.X.CompressGZIP(file_straingauge_straingauge_proto_rawDescData)
	})
	return file_straingauge_straingauge_proto_rawDescData
}

var file_straingauge_straingauge_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_straingauge_straingauge_proto_goTypes = []interface{}{
	(*StrainGuage)(nil), // 0: Staingauge.StrainGuage
	(shared.Finger)(0),  // 1: Shared.Finger
}
var file_straingauge_straingauge_proto_depIdxs = []int32{
	1, // 0: Staingauge.StrainGuage.finger:type_name -> Shared.Finger
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_straingauge_straingauge_proto_init() }
func file_straingauge_straingauge_proto_init() {
	if File_straingauge_straingauge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_straingauge_straingauge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrainGuage); i {
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
			RawDescriptor: file_straingauge_straingauge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_straingauge_straingauge_proto_goTypes,
		DependencyIndexes: file_straingauge_straingauge_proto_depIdxs,
		MessageInfos:      file_straingauge_straingauge_proto_msgTypes,
	}.Build()
	File_straingauge_straingauge_proto = out.File
	file_straingauge_straingauge_proto_rawDesc = nil
	file_straingauge_straingauge_proto_goTypes = nil
	file_straingauge_straingauge_proto_depIdxs = nil
}