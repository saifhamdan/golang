// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: main/protodata.proto

package main

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

type Animal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AnimalType string `protobuf:"bytes,2,opt,name=animal_type,json=animalType,proto3" json:"animal_type,omitempty"`
	Nickname   string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Zone       int32  `protobuf:"varint,4,opt,name=zone,proto3" json:"zone,omitempty"`
	Age        int32  `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Animal) Reset() {
	*x = Animal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_protodata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Animal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Animal) ProtoMessage() {}

func (x *Animal) ProtoReflect() protoreflect.Message {
	mi := &file_main_protodata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Animal.ProtoReflect.Descriptor instead.
func (*Animal) Descriptor() ([]byte, []int) {
	return file_main_protodata_proto_rawDescGZIP(), []int{0}
}

func (x *Animal) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Animal) GetAnimalType() string {
	if x != nil {
		return x.AnimalType
	}
	return ""
}

func (x *Animal) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Animal) GetZone() int32 {
	if x != nil {
		return x.Zone
	}
	return 0
}

func (x *Animal) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_protodata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_main_protodata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_main_protodata_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

var File_main_protodata_proto protoreflect.FileDescriptor

var file_main_protodata_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x06, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x7a, 0x6f, 0x6e,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x53, 0x0a, 0x0b, 0x44, 0x69,
	0x6e, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x08, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x07, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x08, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x30, 0x01, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_main_protodata_proto_rawDescOnce sync.Once
	file_main_protodata_proto_rawDescData = file_main_protodata_proto_rawDesc
)

func file_main_protodata_proto_rawDescGZIP() []byte {
	file_main_protodata_proto_rawDescOnce.Do(func() {
		file_main_protodata_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_protodata_proto_rawDescData)
	})
	return file_main_protodata_proto_rawDescData
}

var file_main_protodata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_main_protodata_proto_goTypes = []interface{}{
	(*Animal)(nil),  // 0: animal
	(*Request)(nil), // 1: request
}
var file_main_protodata_proto_depIdxs = []int32{
	1, // 0: DinoService.GetAnimal:input_type -> request
	1, // 1: DinoService.GetAllAnimals:input_type -> request
	0, // 2: DinoService.GetAnimal:output_type -> animal
	0, // 3: DinoService.GetAllAnimals:output_type -> animal
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_main_protodata_proto_init() }
func file_main_protodata_proto_init() {
	if File_main_protodata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_protodata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Animal); i {
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
		file_main_protodata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_main_protodata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_protodata_proto_goTypes,
		DependencyIndexes: file_main_protodata_proto_depIdxs,
		MessageInfos:      file_main_protodata_proto_msgTypes,
	}.Build()
	File_main_protodata_proto = out.File
	file_main_protodata_proto_rawDesc = nil
	file_main_protodata_proto_goTypes = nil
	file_main_protodata_proto_depIdxs = nil
}
