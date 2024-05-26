// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: char.proto

package gen

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

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname     string `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Job         string `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
	YearOfBirth int32  `protobuf:"varint,4,opt,name=year_of_birth,json=yearOfBirth,proto3" json:"year_of_birth,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_char_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_char_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_char_proto_rawDescGZIP(), []int{0}
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Character) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *Character) GetYearOfBirth() int32 {
	if x != nil {
		return x.YearOfBirth
	}
	return 0
}

var File_char_proto protoreflect.FileDescriptor

var file_char_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x65,
	0x6e, 0x22, 0x6f, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x22,
	0x0a, 0x0d, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x79, 0x65, 0x61, 0x72, 0x4f, 0x66, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x62, 0x69, 0x73, 0x63, 0x75, 0x6d, 0x2f, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2d, 0x4f,
	0x6e, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x6d, 0x69, 0x6e, 0x67, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d, 0x47, 0x6f, 0x2f, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x31, 0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_char_proto_rawDescOnce sync.Once
	file_char_proto_rawDescData = file_char_proto_rawDesc
)

func file_char_proto_rawDescGZIP() []byte {
	file_char_proto_rawDescOnce.Do(func() {
		file_char_proto_rawDescData = protoimpl.X.CompressGZIP(file_char_proto_rawDescData)
	})
	return file_char_proto_rawDescData
}

var file_char_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_char_proto_goTypes = []interface{}{
	(*Character)(nil), // 0: gen.Character
}
var file_char_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_char_proto_init() }
func file_char_proto_init() {
	if File_char_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_char_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
			RawDescriptor: file_char_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_char_proto_goTypes,
		DependencyIndexes: file_char_proto_depIdxs,
		MessageInfos:      file_char_proto_msgTypes,
	}.Build()
	File_char_proto = out.File
	file_char_proto_rawDesc = nil
	file_char_proto_goTypes = nil
	file_char_proto_depIdxs = nil
}
