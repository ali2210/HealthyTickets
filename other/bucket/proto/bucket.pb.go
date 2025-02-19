// This codebase desgin according to mozilla open source license.
//Redistribution , contribution and improve codebase under license
//convensions. @contact Ali Hassan AliMatrixCode@protonmail.com

// protos version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: other/bucket/proto/bucket.proto

// require package

package proto

import (
	binary "github.com/ali2210/wizdwarf/other/proteins/binary"
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

// either object return ok_trait or object return error
type Object_Status int32

const (
	Object_Status_OK    Object_Status = 0
	Object_Status_ERROR Object_Status = 1
)

// Enum value maps for Object_Status.
var (
	Object_Status_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	Object_Status_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x Object_Status) Enum() *Object_Status {
	p := new(Object_Status)
	*p = x
	return p
}

func (x Object_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Object_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_other_bucket_proto_bucket_proto_enumTypes[0].Descriptor()
}

func (Object_Status) Type() protoreflect.EnumType {
	return &file_other_bucket_proto_bucket_proto_enumTypes[0]
}

func (x Object_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Object_Status.Descriptor instead.
func (Object_Status) EnumDescriptor() ([]byte, []int) {
	return file_other_bucket_proto_bucket_proto_rawDescGZIP(), []int{0}
}

type QStatus int32

const (
	QStatus_Ok  QStatus = 0
	QStatus_Err QStatus = 1
)

// Enum value maps for QStatus.
var (
	QStatus_name = map[int32]string{
		0: "Ok",
		1: "Err",
	}
	QStatus_value = map[string]int32{
		"Ok":  0,
		"Err": 1,
	}
)

func (x QStatus) Enum() *QStatus {
	p := new(QStatus)
	*p = x
	return p
}

func (x QStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_other_bucket_proto_bucket_proto_enumTypes[1].Descriptor()
}

func (QStatus) Type() protoreflect.EnumType {
	return &file_other_bucket_proto_bucket_proto_enumTypes[1]
}

func (x QStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QStatus.Descriptor instead.
func (QStatus) EnumDescriptor() ([]byte, []int) {
	return file_other_bucket_proto_bucket_proto_rawDescGZIP(), []int{1}
}

// message Object hold File_Object or File_Descriptor
type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Types   string                  `protobuf:"bytes,2,opt,name=types,proto3" json:"types,omitempty"`
	Content []*binary.Micromolecule `protobuf:"bytes,3,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_bucket_proto_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_other_bucket_proto_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_other_bucket_proto_bucket_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Object) GetTypes() string {
	if x != nil {
		return x.Types
	}
	return ""
}

func (x *Object) GetContent() []*binary.Micromolecule {
	if x != nil {
		return x.Content
	}
	return nil
}

type IObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iobject *Object       `protobuf:"bytes,1,opt,name=iobject,proto3" json:"iobject,omitempty"`
	Istatus Object_Status `protobuf:"varint,2,opt,name=istatus,proto3,enum=proto.Object_Status" json:"istatus,omitempty"`
}

func (x *IObject) Reset() {
	*x = IObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_bucket_proto_bucket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IObject) ProtoMessage() {}

func (x *IObject) ProtoReflect() protoreflect.Message {
	mi := &file_other_bucket_proto_bucket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IObject.ProtoReflect.Descriptor instead.
func (*IObject) Descriptor() ([]byte, []int) {
	return file_other_bucket_proto_bucket_proto_rawDescGZIP(), []int{1}
}

func (x *IObject) GetIobject() *Object {
	if x != nil {
		return x.Iobject
	}
	return nil
}

func (x *IObject) GetIstatus() Object_Status {
	if x != nil {
		return x.Istatus
	}
	return Object_Status_OK
}

// byname is the common waay to iterator document
type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ByName string `protobuf:"bytes,1,opt,name=ByName,proto3" json:"ByName,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_bucket_proto_bucket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_other_bucket_proto_bucket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_other_bucket_proto_bucket_proto_rawDescGZIP(), []int{2}
}

func (x *Query) GetByName() string {
	if x != nil {
		return x.ByName
	}
	return ""
}

// either query return something or none
type QState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qstatus QStatus `protobuf:"varint,1,opt,name=qstatus,proto3,enum=proto.QStatus" json:"qstatus,omitempty"`
}

func (x *QState) Reset() {
	*x = QState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_bucket_proto_bucket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QState) ProtoMessage() {}

func (x *QState) ProtoReflect() protoreflect.Message {
	mi := &file_other_bucket_proto_bucket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QState.ProtoReflect.Descriptor instead.
func (*QState) Descriptor() ([]byte, []int) {
	return file_other_bucket_proto_bucket_proto_rawDescGZIP(), []int{3}
}

func (x *QState) GetQstatus() QStatus {
	if x != nil {
		return x.Qstatus
	}
	return QStatus_Ok
}

var File_other_bucket_proto_bucket_proto protoreflect.FileDescriptor

var file_other_bucket_proto_bucket_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2f,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x06,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f,
	0x6d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x62, 0x0a, 0x07, 0x49, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x07,
	0x69, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x69, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x69, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x69, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1f, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x06, 0x51, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x28, 0x0a, 0x07, 0x71, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x07, 0x71, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x22, 0x0a, 0x0d, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x2a, 0x1a,
	0x0a, 0x07, 0x51, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x10, 0x01, 0x32, 0x87, 0x01, 0x0a, 0x07, 0x47,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x5f, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x69, 0x32, 0x32, 0x31, 0x30, 0x2f, 0x77, 0x69, 0x7a, 0x64, 0x77,
	0x61, 0x72, 0x66, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_other_bucket_proto_bucket_proto_rawDescOnce sync.Once
	file_other_bucket_proto_bucket_proto_rawDescData = file_other_bucket_proto_bucket_proto_rawDesc
)

func file_other_bucket_proto_bucket_proto_rawDescGZIP() []byte {
	file_other_bucket_proto_bucket_proto_rawDescOnce.Do(func() {
		file_other_bucket_proto_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_other_bucket_proto_bucket_proto_rawDescData)
	})
	return file_other_bucket_proto_bucket_proto_rawDescData
}

var file_other_bucket_proto_bucket_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_other_bucket_proto_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_other_bucket_proto_bucket_proto_goTypes = []interface{}{
	(Object_Status)(0),           // 0: proto.Object_Status
	(QStatus)(0),                 // 1: proto.QStatus
	(*Object)(nil),               // 2: proto.Object
	(*IObject)(nil),              // 3: proto.IObject
	(*Query)(nil),                // 4: proto.Query
	(*QState)(nil),               // 5: proto.QState
	(*binary.Micromolecule)(nil), // 6: binary.Micromolecule
}
var file_other_bucket_proto_bucket_proto_depIdxs = []int32{
	6, // 0: proto.Object.content:type_name -> binary.Micromolecule
	2, // 1: proto.IObject.iobject:type_name -> proto.Object
	0, // 2: proto.IObject.istatus:type_name -> proto.Object_Status
	1, // 3: proto.QState.qstatus:type_name -> proto.QStatus
	2, // 4: proto.GBucket.New_Bucket:input_type -> proto.Object
	4, // 5: proto.GBucket.Preview:input_type -> proto.Query
	4, // 6: proto.GBucket.Download:input_type -> proto.Query
	3, // 7: proto.GBucket.New_Bucket:output_type -> proto.IObject
	5, // 8: proto.GBucket.Preview:output_type -> proto.QState
	5, // 9: proto.GBucket.Download:output_type -> proto.QState
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_other_bucket_proto_bucket_proto_init() }
func file_other_bucket_proto_bucket_proto_init() {
	if File_other_bucket_proto_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_other_bucket_proto_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_other_bucket_proto_bucket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IObject); i {
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
		file_other_bucket_proto_bucket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_other_bucket_proto_bucket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QState); i {
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
			RawDescriptor: file_other_bucket_proto_bucket_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_other_bucket_proto_bucket_proto_goTypes,
		DependencyIndexes: file_other_bucket_proto_bucket_proto_depIdxs,
		EnumInfos:         file_other_bucket_proto_bucket_proto_enumTypes,
		MessageInfos:      file_other_bucket_proto_bucket_proto_msgTypes,
	}.Build()
	File_other_bucket_proto_bucket_proto = out.File
	file_other_bucket_proto_bucket_proto_rawDesc = nil
	file_other_bucket_proto_bucket_proto_goTypes = nil
	file_other_bucket_proto_bucket_proto_depIdxs = nil
}
