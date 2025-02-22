// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.9
// source: factorization.proto

package factorization_pb

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

type NumberRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int64                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NumberRequest) Reset() {
	*x = NumberRequest{}
	mi := &file_factorization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberRequest) ProtoMessage() {}

func (x *NumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_factorization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberRequest.ProtoReflect.Descriptor instead.
func (*NumberRequest) Descriptor() ([]byte, []int) {
	return file_factorization_proto_rawDescGZIP(), []int{0}
}

func (x *NumberRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PrimeFactorsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Factors       []int64                `protobuf:"varint,1,rep,packed,name=factors,proto3" json:"factors,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrimeFactorsResponse) Reset() {
	*x = PrimeFactorsResponse{}
	mi := &file_factorization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrimeFactorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeFactorsResponse) ProtoMessage() {}

func (x *PrimeFactorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_factorization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeFactorsResponse.ProtoReflect.Descriptor instead.
func (*PrimeFactorsResponse) Descriptor() ([]byte, []int) {
	return file_factorization_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeFactorsResponse) GetFactors() []int64 {
	if x != nil {
		return x.Factors
	}
	return nil
}

var File_factorization_proto protoreflect.FileDescriptor

var file_factorization_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x30,
	0x0a, 0x14, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x32, 0x99, 0x01, 0x0a, 0x19, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0x0e, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x0e, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x12,
	0x2e, 0x2f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_factorization_proto_rawDescOnce sync.Once
	file_factorization_proto_rawDescData []byte
)

func file_factorization_proto_rawDescGZIP() []byte {
	file_factorization_proto_rawDescOnce.Do(func() {
		file_factorization_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_factorization_proto_rawDesc), len(file_factorization_proto_rawDesc)))
	})
	return file_factorization_proto_rawDescData
}

var file_factorization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_factorization_proto_goTypes = []any{
	(*NumberRequest)(nil),        // 0: NumberRequest
	(*PrimeFactorsResponse)(nil), // 1: PrimeFactorsResponse
}
var file_factorization_proto_depIdxs = []int32{
	0, // 0: PrimeFactorizationService.GetPrimeFactors:input_type -> NumberRequest
	0, // 1: PrimeFactorizationService.GetPrimeFactorsStream:input_type -> NumberRequest
	1, // 2: PrimeFactorizationService.GetPrimeFactors:output_type -> PrimeFactorsResponse
	1, // 3: PrimeFactorizationService.GetPrimeFactorsStream:output_type -> PrimeFactorsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_factorization_proto_init() }
func file_factorization_proto_init() {
	if File_factorization_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_factorization_proto_rawDesc), len(file_factorization_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_factorization_proto_goTypes,
		DependencyIndexes: file_factorization_proto_depIdxs,
		MessageInfos:      file_factorization_proto_msgTypes,
	}.Build()
	File_factorization_proto = out.File
	file_factorization_proto_goTypes = nil
	file_factorization_proto_depIdxs = nil
}
