// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/ports.proto

package shared

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PortBlueprint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Alias       []string  `protobuf:"bytes,3,rep,name=alias,proto3" json:"alias,omitempty"`
	City        string    `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Code        string    `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	Coordinates []float64 `protobuf:"fixed64,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Country     string    `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	Province    string    `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Regions     []string  `protobuf:"bytes,9,rep,name=regions,proto3" json:"regions,omitempty"`
	Timezone    string    `protobuf:"bytes,10,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,11,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
}

func (x *PortBlueprint) Reset() {
	*x = PortBlueprint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortBlueprint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortBlueprint) ProtoMessage() {}

func (x *PortBlueprint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortBlueprint.ProtoReflect.Descriptor instead.
func (*PortBlueprint) Descriptor() ([]byte, []int) {
	return file_proto_ports_proto_rawDescGZIP(), []int{0}
}

func (x *PortBlueprint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PortBlueprint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PortBlueprint) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *PortBlueprint) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *PortBlueprint) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PortBlueprint) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *PortBlueprint) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PortBlueprint) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *PortBlueprint) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *PortBlueprint) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *PortBlueprint) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submitted uint64 `protobuf:"varint,1,opt,name=submitted,proto3" json:"submitted,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_proto_ports_proto_rawDescGZIP(), []int{1}
}

func (x *Summary) GetSubmitted() uint64 {
	if x != nil {
		return x.Submitted
	}
	return 0
}

var File_proto_ports_proto protoreflect.FileDescriptor

var file_proto_ports_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x22, 0x97, 0x02, 0x0a, 0x0d,
	0x50, 0x6f, 0x72, 0x74, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x22, 0x27, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x32, 0x3e,
	0x0a, 0x08, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x04, 0x53, 0x65,
	0x6e, 0x64, 0x12, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0c,
	0x5a, 0x0a, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ports_proto_rawDescOnce sync.Once
	file_proto_ports_proto_rawDescData = file_proto_ports_proto_rawDesc
)

func file_proto_ports_proto_rawDescGZIP() []byte {
	file_proto_ports_proto_rawDescOnce.Do(func() {
		file_proto_ports_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ports_proto_rawDescData)
	})
	return file_proto_ports_proto_rawDescData
}

var file_proto_ports_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_ports_proto_goTypes = []interface{}{
	(*PortBlueprint)(nil), // 0: shared.PortBlueprint
	(*Summary)(nil),       // 1: shared.Summary
}
var file_proto_ports_proto_depIdxs = []int32{
	0, // 0: shared.Streamer.Send:input_type -> shared.PortBlueprint
	1, // 1: shared.Streamer.Send:output_type -> shared.Summary
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ports_proto_init() }
func file_proto_ports_proto_init() {
	if File_proto_ports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortBlueprint); i {
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
		file_proto_ports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary); i {
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
			RawDescriptor: file_proto_ports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ports_proto_goTypes,
		DependencyIndexes: file_proto_ports_proto_depIdxs,
		MessageInfos:      file_proto_ports_proto_msgTypes,
	}.Build()
	File_proto_ports_proto = out.File
	file_proto_ports_proto_rawDesc = nil
	file_proto_ports_proto_goTypes = nil
	file_proto_ports_proto_depIdxs = nil
}
