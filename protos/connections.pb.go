// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.6
// source: connections.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Actions int32

const (
	Actions_ADD_DEALER    Actions = 0
	Actions_REMOVE_DEALER Actions = 1
	Actions_ADD_CLIENT    Actions = 2
	Actions_REMOVE_CLIENT Actions = 3
)

// Enum value maps for Actions.
var (
	Actions_name = map[int32]string{
		0: "ADD_DEALER",
		1: "REMOVE_DEALER",
		2: "ADD_CLIENT",
		3: "REMOVE_CLIENT",
	}
	Actions_value = map[string]int32{
		"ADD_DEALER":    0,
		"REMOVE_DEALER": 1,
		"ADD_CLIENT":    2,
		"REMOVE_CLIENT": 3,
	}
)

func (x Actions) Enum() *Actions {
	p := new(Actions)
	*p = x
	return p
}

func (x Actions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Actions) Descriptor() protoreflect.EnumDescriptor {
	return file_connections_proto_enumTypes[0].Descriptor()
}

func (Actions) Type() protoreflect.EnumType {
	return &file_connections_proto_enumTypes[0]
}

func (x Actions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Actions.Descriptor instead.
func (Actions) EnumDescriptor() ([]byte, []int) {
	return file_connections_proto_rawDescGZIP(), []int{0}
}

// *
// Every message sent must have use Evelope
type Envelope struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SrcId         string                 `protobuf:"bytes,1,opt,name=src_id,json=srcId,proto3" json:"src_id,omitempty"`
	Src           string                 `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
	SrcPort       int32                  `protobuf:"varint,3,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstId         string                 `protobuf:"bytes,4,opt,name=dst_id,json=dstId,proto3" json:"dst_id,omitempty"`
	Dst           string                 `protobuf:"bytes,5,opt,name=dst,proto3" json:"dst,omitempty"`
	DstPort       int32                  `protobuf:"varint,6,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	Actions       Actions                `protobuf:"varint,7,opt,name=actions,proto3,enum=envelope.Actions" json:"actions,omitempty"`
	MsgId         string                 `protobuf:"bytes,8,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	SentAt        *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	mi := &file_connections_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_connections_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_connections_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetSrcId() string {
	if x != nil {
		return x.SrcId
	}
	return ""
}

func (x *Envelope) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *Envelope) GetSrcPort() int32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *Envelope) GetDstId() string {
	if x != nil {
		return x.DstId
	}
	return ""
}

func (x *Envelope) GetDst() string {
	if x != nil {
		return x.Dst
	}
	return ""
}

func (x *Envelope) GetDstPort() int32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

func (x *Envelope) GetActions() Actions {
	if x != nil {
		return x.Actions
	}
	return Actions_ADD_DEALER
}

func (x *Envelope) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *Envelope) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

var File_connections_proto protoreflect.FileDescriptor

var file_connections_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b,
	0x02, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x72, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x72, 0x63,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x72, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x64, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x64, 0x73, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x2a, 0x4f, 0x0a, 0x07,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x44, 0x44, 0x5f, 0x44,
	0x45, 0x41, 0x4c, 0x45, 0x52, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x4d, 0x4f, 0x56,
	0x45, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x44,
	0x44, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45,
	0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_connections_proto_rawDescOnce sync.Once
	file_connections_proto_rawDescData []byte
)

func file_connections_proto_rawDescGZIP() []byte {
	file_connections_proto_rawDescOnce.Do(func() {
		file_connections_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_connections_proto_rawDesc), len(file_connections_proto_rawDesc)))
	})
	return file_connections_proto_rawDescData
}

var file_connections_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_connections_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_connections_proto_goTypes = []any{
	(Actions)(0),                  // 0: envelope.Actions
	(*Envelope)(nil),              // 1: envelope.Envelope
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_connections_proto_depIdxs = []int32{
	0, // 0: envelope.Envelope.actions:type_name -> envelope.Actions
	2, // 1: envelope.Envelope.sent_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_connections_proto_init() }
func file_connections_proto_init() {
	if File_connections_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_connections_proto_rawDesc), len(file_connections_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connections_proto_goTypes,
		DependencyIndexes: file_connections_proto_depIdxs,
		EnumInfos:         file_connections_proto_enumTypes,
		MessageInfos:      file_connections_proto_msgTypes,
	}.Build()
	File_connections_proto = out.File
	file_connections_proto_goTypes = nil
	file_connections_proto_depIdxs = nil
}
