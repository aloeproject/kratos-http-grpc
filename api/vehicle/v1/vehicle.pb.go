// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: vehicle.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetVehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetVehicleRequest) Reset() {
	*x = GetVehicleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVehicleRequest) ProtoMessage() {}

func (x *GetVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVehicleRequest.ProtoReflect.Descriptor instead.
func (*GetVehicleRequest) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{0}
}

func (x *GetVehicleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetVehicleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetVehicleReply) Reset() {
	*x = GetVehicleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVehicleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVehicleReply) ProtoMessage() {}

func (x *GetVehicleReply) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVehicleReply.ProtoReflect.Descriptor instead.
func (*GetVehicleReply) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{1}
}

func (x *GetVehicleReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_vehicle_proto protoreflect.FileDescriptor

var file_vehicle_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x7c, 0x0a, 0x07, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x71, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x14,
	0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0x36, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x22, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2d, 0x68, 0x74, 0x74, 0x70, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vehicle_proto_rawDescOnce sync.Once
	file_vehicle_proto_rawDescData = file_vehicle_proto_rawDesc
)

func file_vehicle_proto_rawDescGZIP() []byte {
	file_vehicle_proto_rawDescOnce.Do(func() {
		file_vehicle_proto_rawDescData = protoimpl.X.CompressGZIP(file_vehicle_proto_rawDescData)
	})
	return file_vehicle_proto_rawDescData
}

var file_vehicle_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vehicle_proto_goTypes = []interface{}{
	(*GetVehicleRequest)(nil), // 0: api.vehicle.v1.GetVehicleRequest
	(*GetVehicleReply)(nil),   // 1: api.vehicle.v1.GetVehicleReply
}
var file_vehicle_proto_depIdxs = []int32{
	0, // 0: api.vehicle.v1.Vehicle.GetVehicle:input_type -> api.vehicle.v1.GetVehicleRequest
	1, // 1: api.vehicle.v1.Vehicle.GetVehicle:output_type -> api.vehicle.v1.GetVehicleReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vehicle_proto_init() }
func file_vehicle_proto_init() {
	if File_vehicle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vehicle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVehicleRequest); i {
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
		file_vehicle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVehicleReply); i {
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
			RawDescriptor: file_vehicle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vehicle_proto_goTypes,
		DependencyIndexes: file_vehicle_proto_depIdxs,
		MessageInfos:      file_vehicle_proto_msgTypes,
	}.Build()
	File_vehicle_proto = out.File
	file_vehicle_proto_rawDesc = nil
	file_vehicle_proto_goTypes = nil
	file_vehicle_proto_depIdxs = nil
}