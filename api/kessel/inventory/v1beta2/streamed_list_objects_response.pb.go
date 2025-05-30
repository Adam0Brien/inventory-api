// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta2/streamed_list_objects_response.proto

package v1beta2

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

type StreamedListObjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object           *ResourceReference  `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Pagination       *ResponsePagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	ConsistencyToken *ConsistencyToken   `protobuf:"bytes,3,opt,name=consistency_token,json=consistencyToken,proto3" json:"consistency_token,omitempty"`
}

func (x *StreamedListObjectsResponse) Reset() {
	*x = StreamedListObjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamedListObjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamedListObjectsResponse) ProtoMessage() {}

func (x *StreamedListObjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamedListObjectsResponse.ProtoReflect.Descriptor instead.
func (*StreamedListObjectsResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamedListObjectsResponse) GetObject() *ResourceReference {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *StreamedListObjectsResponse) GetPagination() *ResponsePagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *StreamedListObjectsResponse) GetConsistencyToken() *ConsistencyToken {
	if x != nil {
		return x.ConsistencyToken
	}
	return nil
}

var File_kessel_inventory_v1beta2_streamed_list_objects_response_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x31, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x30, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x1b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4c, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x10, 0x63, 0x6f,
	0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x72,
	0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDescData = file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDesc
)

func file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDescData
}

var file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_goTypes = []any{
	(*StreamedListObjectsResponse)(nil), // 0: kessel.inventory.v1beta2.StreamedListObjectsResponse
	(*ResourceReference)(nil),           // 1: kessel.inventory.v1beta2.ResourceReference
	(*ResponsePagination)(nil),          // 2: kessel.inventory.v1beta2.ResponsePagination
	(*ConsistencyToken)(nil),            // 3: kessel.inventory.v1beta2.ConsistencyToken
}
var file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_depIdxs = []int32{
	1, // 0: kessel.inventory.v1beta2.StreamedListObjectsResponse.object:type_name -> kessel.inventory.v1beta2.ResourceReference
	2, // 1: kessel.inventory.v1beta2.StreamedListObjectsResponse.pagination:type_name -> kessel.inventory.v1beta2.ResponsePagination
	3, // 2: kessel.inventory.v1beta2.StreamedListObjectsResponse.consistency_token:type_name -> kessel.inventory.v1beta2.ConsistencyToken
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_init() }
func file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_init() {
	if File_kessel_inventory_v1beta2_streamed_list_objects_response_proto != nil {
		return
	}
	file_kessel_inventory_v1beta2_resource_reference_proto_init()
	file_kessel_inventory_v1beta2_response_pagination_proto_init()
	file_kessel_inventory_v1beta2_consistency_token_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StreamedListObjectsResponse); i {
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
			RawDescriptor: file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta2_streamed_list_objects_response_proto = out.File
	file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_rawDesc = nil
	file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_goTypes = nil
	file_kessel_inventory_v1beta2_streamed_list_objects_response_proto_depIdxs = nil
}
