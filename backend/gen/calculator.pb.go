// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: calculator.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CalculatorRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             float64                `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"` // 第一个操作数
	B             float64                `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"` // 第二个操作数
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	mi := &file_calculator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculatorRequest) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *CalculatorRequest) GetB() float64 {
	if x != nil {
		return x.B
	}
	return 0
}

type CalculatorResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float64                `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"` // 运算结果
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	mi := &file_calculator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_proto protoreflect.FileDescriptor

const file_calculator_proto_rawDesc = "" +
	"\n" +
	"\x10calculator.proto\x12\n" +
	"calculator\"/\n" +
	"\x11CalculatorRequest\x12\f\n" +
	"\x01a\x18\x01 \x01(\x01R\x01a\x12\f\n" +
	"\x01b\x18\x02 \x01(\x01R\x01b\",\n" +
	"\x12CalculatorResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\x01R\x06result2\xb8\x02\n" +
	"\x11CalculatorService\x12D\n" +
	"\x03Add\x12\x1d.calculator.CalculatorRequest\x1a\x1e.calculator.CalculatorResponse\x12I\n" +
	"\bSubtract\x12\x1d.calculator.CalculatorRequest\x1a\x1e.calculator.CalculatorResponse\x12I\n" +
	"\bMultiply\x12\x1d.calculator.CalculatorRequest\x1a\x1e.calculator.CalculatorResponse\x12G\n" +
	"\x06Divide\x12\x1d.calculator.CalculatorRequest\x1a\x1e.calculator.CalculatorResponseB\x10Z\x0ecalculator/genb\x06proto3"

var (
	file_calculator_proto_rawDescOnce sync.Once
	file_calculator_proto_rawDescData []byte
)

func file_calculator_proto_rawDescGZIP() []byte {
	file_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_calculator_proto_rawDesc), len(file_calculator_proto_rawDesc)))
	})
	return file_calculator_proto_rawDescData
}

var file_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculator_proto_goTypes = []any{
	(*CalculatorRequest)(nil),  // 0: calculator.CalculatorRequest
	(*CalculatorResponse)(nil), // 1: calculator.CalculatorResponse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorService.Add:input_type -> calculator.CalculatorRequest
	0, // 1: calculator.CalculatorService.Subtract:input_type -> calculator.CalculatorRequest
	0, // 2: calculator.CalculatorService.Multiply:input_type -> calculator.CalculatorRequest
	0, // 3: calculator.CalculatorService.Divide:input_type -> calculator.CalculatorRequest
	1, // 4: calculator.CalculatorService.Add:output_type -> calculator.CalculatorResponse
	1, // 5: calculator.CalculatorService.Subtract:output_type -> calculator.CalculatorResponse
	1, // 6: calculator.CalculatorService.Multiply:output_type -> calculator.CalculatorResponse
	1, // 7: calculator.CalculatorService.Divide:output_type -> calculator.CalculatorResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calculator_proto_rawDesc), len(file_calculator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_proto_msgTypes,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}
