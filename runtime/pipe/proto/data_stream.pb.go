// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: data_stream.proto

package proto

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

type PutAddrMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func (x *PutAddrMsg) Reset() {
	*x = PutAddrMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutAddrMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutAddrMsg) ProtoMessage() {}

func (x *PutAddrMsg) ProtoReflect() protoreflect.Message {
	mi := &file_data_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutAddrMsg.ProtoReflect.Descriptor instead.
func (*PutAddrMsg) Descriptor() ([]byte, []int) {
	return file_data_stream_proto_rawDescGZIP(), []int{0}
}

func (x *PutAddrMsg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutAddrMsg) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type AddrReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *AddrReq) Reset() {
	*x = AddrReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddrReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddrReq) ProtoMessage() {}

func (x *AddrReq) ProtoReflect() protoreflect.Message {
	mi := &file_data_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddrReq.ProtoReflect.Descriptor instead.
func (*AddrReq) Descriptor() ([]byte, []int) {
	return file_data_stream_proto_rawDescGZIP(), []int{1}
}

func (x *AddrReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAddrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func (x *GetAddrReply) Reset() {
	*x = GetAddrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAddrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddrReply) ProtoMessage() {}

func (x *GetAddrReply) ProtoReflect() protoreflect.Message {
	mi := &file_data_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddrReply.ProtoReflect.Descriptor instead.
func (*GetAddrReply) Descriptor() ([]byte, []int) {
	return file_data_stream_proto_rawDescGZIP(), []int{2}
}

func (x *GetAddrReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetAddrReply) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_data_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_data_stream_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buffer []byte `protobuf:"bytes,1,opt,name=buffer,proto3" json:"buffer,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"` // only sent with first message
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_data_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_data_stream_proto_rawDescGZIP(), []int{4}
}

func (x *Data) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *Data) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_data_stream_proto protoreflect.FileDescriptor

var file_data_stream_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x41, 0x64, 0x64, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x41, 0x64, 0x64, 0x72, 0x22, 0x19, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x3c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x22, 0x22,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x32, 0x98, 0x02, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x12, 0x21, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x0b, 0x2e, 0x50, 0x75,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x4d, 0x73, 0x67, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x08,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0a, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0a,
	0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x08, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x05, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x21, 0x0a, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x05,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x2a, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x41, 0x64, 0x64, 0x72, 0x4f, 0x70, 0x74,
	0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x50, 0x75, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x4d, 0x73, 0x67, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2d,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a,
	0x65, 0x64, 0x12, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x14, 0x5a,
	0x12, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_stream_proto_rawDescOnce sync.Once
	file_data_stream_proto_rawDescData = file_data_stream_proto_rawDesc
)

func file_data_stream_proto_rawDescGZIP() []byte {
	file_data_stream_proto_rawDescOnce.Do(func() {
		file_data_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_stream_proto_rawDescData)
	})
	return file_data_stream_proto_rawDescData
}

var file_data_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_data_stream_proto_goTypes = []interface{}{
	(*PutAddrMsg)(nil),   // 0: PutAddrMsg
	(*AddrReq)(nil),      // 1: AddrReq
	(*GetAddrReply)(nil), // 2: GetAddrReply
	(*Status)(nil),       // 3: Status
	(*Data)(nil),         // 4: Data
}
var file_data_stream_proto_depIdxs = []int32{
	0, // 0: Discovery.PutAddr:input_type -> PutAddrMsg
	1, // 1: Discovery.GetAddr:input_type -> AddrReq
	1, // 2: Discovery.RemoveAddr:input_type -> AddrReq
	1, // 3: Discovery.readStream:input_type -> AddrReq
	4, // 4: Discovery.writeStream:input_type -> Data
	0, // 5: Discovery.PutAddrOptimized:input_type -> PutAddrMsg
	1, // 6: Discovery.GetAddrOptimized:input_type -> AddrReq
	3, // 7: Discovery.PutAddr:output_type -> Status
	2, // 8: Discovery.GetAddr:output_type -> GetAddrReply
	3, // 9: Discovery.RemoveAddr:output_type -> Status
	4, // 10: Discovery.readStream:output_type -> Data
	3, // 11: Discovery.writeStream:output_type -> Status
	3, // 12: Discovery.PutAddrOptimized:output_type -> Status
	2, // 13: Discovery.GetAddrOptimized:output_type -> GetAddrReply
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_stream_proto_init() }
func file_data_stream_proto_init() {
	if File_data_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutAddrMsg); i {
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
		file_data_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddrReq); i {
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
		file_data_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAddrReply); i {
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
		file_data_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_data_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
			RawDescriptor: file_data_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_stream_proto_goTypes,
		DependencyIndexes: file_data_stream_proto_depIdxs,
		MessageInfos:      file_data_stream_proto_msgTypes,
	}.Build()
	File_data_stream_proto = out.File
	file_data_stream_proto_rawDesc = nil
	file_data_stream_proto_goTypes = nil
	file_data_stream_proto_depIdxs = nil
}
