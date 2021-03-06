// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: question-2/transport/grpc/imdb_grpc/imdb.proto

package imdb_grpc

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

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Page   string `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *FetchRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

type SingleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SingleRequest) Reset() {
	*x = SingleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleRequest) ProtoMessage() {}

func (x *SingleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleRequest.ProtoReflect.Descriptor instead.
func (*SingleRequest) Descriptor() ([]byte, []int) {
	return file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescGZIP(), []int{1}
}

func (x *SingleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListResults []*MovieDetail `protobuf:"bytes,1,rep,name=ListResults,proto3" json:"ListResults,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescGZIP(), []int{2}
}

func (x *SearchResponse) GetListResults() []*MovieDetail {
	if x != nil {
		return x.ListResults
	}
	return nil
}

type MovieDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *MovieDetail) Reset() {
	*x = MovieDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetail) ProtoMessage() {}

func (x *MovieDetail) ProtoReflect() protoreflect.Message {
	mi := &file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetail.ProtoReflect.Descriptor instead.
func (*MovieDetail) Descriptor() ([]byte, []int) {
	return file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescGZIP(), []int{3}
}

func (x *MovieDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieDetail) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieDetail) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *MovieDetail) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieDetail) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

var File_question_2_transport_grpc_imdb_grpc_imdb_proto protoreflect.FileDescriptor

var file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x32, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6d, 0x64, 0x62,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x69, 0x6d, 0x64, 0x62, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x3a, 0x0a, 0x0c, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x69, 0x6d, 0x64, 0x62, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0x7b, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49,
	0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x32, 0x8f, 0x01, 0x0a, 0x0f, 0x49, 0x6d, 0x64, 0x62, 0x47, 0x72, 0x70, 0x63, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x17, 0x2e, 0x69, 0x6d, 0x64, 0x62, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x6d,
	0x64, 0x62, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x18, 0x2e, 0x69, 0x6d, 0x64, 0x62, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x6d,
	0x64, 0x62, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x72, 0x69, 0x66, 0x39, 0x38, 0x37, 0x38, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x62, 0x69, 0x74, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2d, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x69, 0x6d, 0x64, 0x62, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescOnce sync.Once
	file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescData = file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDesc
)

func file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescGZIP() []byte {
	file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescOnce.Do(func() {
		file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescData)
	})
	return file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDescData
}

var file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_question_2_transport_grpc_imdb_grpc_imdb_proto_goTypes = []interface{}{
	(*FetchRequest)(nil),   // 0: imdb_grpc.FetchRequest
	(*SingleRequest)(nil),  // 1: imdb_grpc.SingleRequest
	(*SearchResponse)(nil), // 2: imdb_grpc.SearchResponse
	(*MovieDetail)(nil),    // 3: imdb_grpc.MovieDetail
}
var file_question_2_transport_grpc_imdb_grpc_imdb_proto_depIdxs = []int32{
	3, // 0: imdb_grpc.SearchResponse.ListResults:type_name -> imdb_grpc.MovieDetail
	0, // 1: imdb_grpc.ImdbGrpcHandler.FetchList:input_type -> imdb_grpc.FetchRequest
	1, // 2: imdb_grpc.ImdbGrpcHandler.GetByID:input_type -> imdb_grpc.SingleRequest
	2, // 3: imdb_grpc.ImdbGrpcHandler.FetchList:output_type -> imdb_grpc.SearchResponse
	3, // 4: imdb_grpc.ImdbGrpcHandler.GetByID:output_type -> imdb_grpc.MovieDetail
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_question_2_transport_grpc_imdb_grpc_imdb_proto_init() }
func file_question_2_transport_grpc_imdb_grpc_imdb_proto_init() {
	if File_question_2_transport_grpc_imdb_grpc_imdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleRequest); i {
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
		file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetail); i {
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
			RawDescriptor: file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_question_2_transport_grpc_imdb_grpc_imdb_proto_goTypes,
		DependencyIndexes: file_question_2_transport_grpc_imdb_grpc_imdb_proto_depIdxs,
		MessageInfos:      file_question_2_transport_grpc_imdb_grpc_imdb_proto_msgTypes,
	}.Build()
	File_question_2_transport_grpc_imdb_grpc_imdb_proto = out.File
	file_question_2_transport_grpc_imdb_grpc_imdb_proto_rawDesc = nil
	file_question_2_transport_grpc_imdb_grpc_imdb_proto_goTypes = nil
	file_question_2_transport_grpc_imdb_grpc_imdb_proto_depIdxs = nil
}
