// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: city_message.proto

package cities

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

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_city_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_city_message_proto_rawDescGZIP(), []int{0}
}

func (x *City) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *City) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_city_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_city_message_proto_rawDescGZIP(), []int{1}
}

type CitiesStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City *City `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *CitiesStream) Reset() {
	*x = CitiesStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CitiesStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitiesStream) ProtoMessage() {}

func (x *CitiesStream) ProtoReflect() protoreflect.Message {
	mi := &file_city_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitiesStream.ProtoReflect.Descriptor instead.
func (*CitiesStream) Descriptor() ([]byte, []int) {
	return file_city_message_proto_rawDescGZIP(), []int{2}
}

func (x *CitiesStream) GetCity() *City {
	if x != nil {
		return x.City
	}
	return nil
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_city_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_city_message_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CityInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CityInput) Reset() {
	*x = CityInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityInput) ProtoMessage() {}

func (x *CityInput) ProtoReflect() protoreflect.Message {
	mi := &file_city_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityInput.ProtoReflect.Descriptor instead.
func (*CityInput) Descriptor() ([]byte, []int) {
	return file_city_message_proto_rawDescGZIP(), []int{4}
}

func (x *CityInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MyBoolean struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boolean bool `protobuf:"varint,1,opt,name=boolean,proto3" json:"boolean,omitempty"`
}

func (x *MyBoolean) Reset() {
	*x = MyBoolean{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyBoolean) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyBoolean) ProtoMessage() {}

func (x *MyBoolean) ProtoReflect() protoreflect.Message {
	mi := &file_city_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyBoolean.ProtoReflect.Descriptor instead.
func (*MyBoolean) Descriptor() ([]byte, []int) {
	return file_city_message_proto_rawDescGZIP(), []int{5}
}

func (x *MyBoolean) GetBoolean() bool {
	if x != nil {
		return x.Boolean
	}
	return false
}

var File_city_message_proto protoreflect.FileDescriptor

var file_city_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x2a, 0x0a, 0x04,
	0x43, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x0c, 0x43, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1f, 0x0a, 0x09, 0x43, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x25, 0x0a, 0x09, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x62, 0x2f, 0x63,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x3b, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_message_proto_rawDescOnce sync.Once
	file_city_message_proto_rawDescData = file_city_message_proto_rawDesc
)

func file_city_message_proto_rawDescGZIP() []byte {
	file_city_message_proto_rawDescOnce.Do(func() {
		file_city_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_message_proto_rawDescData)
	})
	return file_city_message_proto_rawDescData
}

var file_city_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_city_message_proto_goTypes = []any{
	(*City)(nil),         // 0: cities.City
	(*EmptyMessage)(nil), // 1: cities.EmptyMessage
	(*CitiesStream)(nil), // 2: cities.CitiesStream
	(*Id)(nil),           // 3: cities.Id
	(*CityInput)(nil),    // 4: cities.CityInput
	(*MyBoolean)(nil),    // 5: cities.MyBoolean
}
var file_city_message_proto_depIdxs = []int32{
	0, // 0: cities.CitiesStream.city:type_name -> cities.City
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_city_message_proto_init() }
func file_city_message_proto_init() {
	if File_city_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*City); i {
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
		file_city_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EmptyMessage); i {
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
		file_city_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CitiesStream); i {
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
		file_city_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Id); i {
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
		file_city_message_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CityInput); i {
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
		file_city_message_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*MyBoolean); i {
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
			RawDescriptor: file_city_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_city_message_proto_goTypes,
		DependencyIndexes: file_city_message_proto_depIdxs,
		MessageInfos:      file_city_message_proto_msgTypes,
	}.Build()
	File_city_message_proto = out.File
	file_city_message_proto_rawDesc = nil
	file_city_message_proto_goTypes = nil
	file_city_message_proto_depIdxs = nil
}
