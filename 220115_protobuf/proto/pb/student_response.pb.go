// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: student_response.proto

package pb

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

// 嵌套写法
type Student2Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results2 []*Student2Response_Result2 `protobuf:"bytes,1,rep,name=results2,proto3" json:"results2,omitempty"`
}

func (x *Student2Response) Reset() {
	*x = Student2Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student2Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student2Response) ProtoMessage() {}

func (x *Student2Response) ProtoReflect() protoreflect.Message {
	mi := &file_student_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student2Response.ProtoReflect.Descriptor instead.
func (*Student2Response) Descriptor() ([]byte, []int) {
	return file_student_response_proto_rawDescGZIP(), []int{0}
}

func (x *Student2Response) GetResults2() []*Student2Response_Result2 {
	if x != nil {
		return x.Results2
	}
	return nil
}

type Student2Response_Result2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets []string `protobuf:"bytes,3,rep,name=snippets,proto3" json:"snippets,omitempty"`
}

func (x *Student2Response_Result2) Reset() {
	*x = Student2Response_Result2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student2Response_Result2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student2Response_Result2) ProtoMessage() {}

func (x *Student2Response_Result2) ProtoReflect() protoreflect.Message {
	mi := &file_student_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student2Response_Result2.ProtoReflect.Descriptor instead.
func (*Student2Response_Result2) Descriptor() ([]byte, []int) {
	return file_student_response_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Student2Response_Result2) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Student2Response_Result2) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Student2Response_Result2) GetSnippets() []string {
	if x != nil {
		return x.Snippets
	}
	return nil
}

var File_student_response_proto protoreflect.FileDescriptor

var file_student_response_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x9d,
	0x01, 0x0a, 0x10, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x1a,
	0x4d, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_student_response_proto_rawDescOnce sync.Once
	file_student_response_proto_rawDescData = file_student_response_proto_rawDesc
)

func file_student_response_proto_rawDescGZIP() []byte {
	file_student_response_proto_rawDescOnce.Do(func() {
		file_student_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_response_proto_rawDescData)
	})
	return file_student_response_proto_rawDescData
}

var file_student_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_student_response_proto_goTypes = []interface{}{
	(*Student2Response)(nil),         // 0: main.Student2Response
	(*Student2Response_Result2)(nil), // 1: main.Student2Response.Result2
}
var file_student_response_proto_depIdxs = []int32{
	1, // 0: main.Student2Response.results2:type_name -> main.Student2Response.Result2
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_student_response_proto_init() }
func file_student_response_proto_init() {
	if File_student_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student2Response); i {
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
		file_student_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student2Response_Result2); i {
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
			RawDescriptor: file_student_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_student_response_proto_goTypes,
		DependencyIndexes: file_student_response_proto_depIdxs,
		MessageInfos:      file_student_response_proto_msgTypes,
	}.Build()
	File_student_response_proto = out.File
	file_student_response_proto_rawDesc = nil
	file_student_response_proto_goTypes = nil
	file_student_response_proto_depIdxs = nil
}