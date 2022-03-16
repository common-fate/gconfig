// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: gconfig/okta/v1alpha1/gconfig.proto

package oktav1alpha1

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admins        []*Member      `protobuf:"bytes,1,rep,name=admins,proto3" json:"admins,omitempty"`
	Roles         []*Role        `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	AccessHandler *AccessHandler `protobuf:"bytes,3,opt,name=accessHandler,proto3" json:"accessHandler,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_gconfig_okta_v1alpha1_gconfig_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetAdmins() []*Member {
	if x != nil {
		return x.Admins
	}
	return nil
}

func (x *Config) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Config) GetAccessHandler() *AccessHandler {
	if x != nil {
		return x.AccessHandler
	}
	return nil
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_gconfig_okta_v1alpha1_gconfig_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AccessHandler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *AccessHandler) Reset() {
	*x = AccessHandler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessHandler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessHandler) ProtoMessage() {}

func (x *AccessHandler) ProtoReflect() protoreflect.Message {
	mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessHandler.ProtoReflect.Descriptor instead.
func (*AccessHandler) Descriptor() ([]byte, []int) {
	return file_gconfig_okta_v1alpha1_gconfig_proto_rawDescGZIP(), []int{2}
}

func (x *AccessHandler) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Group string  `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Rules []*Rule `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_gconfig_okta_v1alpha1_gconfig_proto_rawDescGZIP(), []int{3}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Role) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy     string `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	Group      string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Breakglass bool   `protobuf:"varint,3,opt,name=breakglass,proto3" json:"breakglass,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_gconfig_okta_v1alpha1_gconfig_proto_rawDescGZIP(), []int{4}
}

func (x *Rule) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *Rule) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Rule) GetBreakglass() bool {
	if x != nil {
		return x.Breakglass
	}
	return false
}

var File_gconfig_okta_v1alpha1_gconfig_proto protoreflect.FileDescriptor

var file_gconfig_okta_v1alpha1_gconfig_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6f, 0x6b, 0x74, 0x61, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f,
	0x6b, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0xbe, 0x01, 0x0a,
	0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x6f, 0x6b, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x12, 0x31,
	0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x6b, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x12, 0x4a, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x6f, 0x6b, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x0d,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x22, 0x1e, 0x0a,
	0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x21, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x5f, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x31,
	0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x6b, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x22, 0x54, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x65, 0x61, 0x6b,
	0x67, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x67, 0x6c, 0x61, 0x73, 0x73, 0x42, 0xe6, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x6b, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x47, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x6f, 0x6b, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x6f, 0x6b, 0x74, 0x61, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x47,
	0x4f, 0x58, 0xaa, 0x02, 0x15, 0x47, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f, 0x6b, 0x74,
	0x61, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x47, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5c, 0x4f, 0x6b, 0x74, 0x61, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x21, 0x47, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x4f, 0x6b, 0x74,
	0x61, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x47, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x3a, 0x3a, 0x4f, 0x6b, 0x74, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gconfig_okta_v1alpha1_gconfig_proto_rawDescOnce sync.Once
	file_gconfig_okta_v1alpha1_gconfig_proto_rawDescData = file_gconfig_okta_v1alpha1_gconfig_proto_rawDesc
)

func file_gconfig_okta_v1alpha1_gconfig_proto_rawDescGZIP() []byte {
	file_gconfig_okta_v1alpha1_gconfig_proto_rawDescOnce.Do(func() {
		file_gconfig_okta_v1alpha1_gconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_gconfig_okta_v1alpha1_gconfig_proto_rawDescData)
	})
	return file_gconfig_okta_v1alpha1_gconfig_proto_rawDescData
}

var file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gconfig_okta_v1alpha1_gconfig_proto_goTypes = []interface{}{
	(*Config)(nil),        // 0: gconfig.okta.v1alpha1.Config
	(*Member)(nil),        // 1: gconfig.okta.v1alpha1.Member
	(*AccessHandler)(nil), // 2: gconfig.okta.v1alpha1.AccessHandler
	(*Role)(nil),          // 3: gconfig.okta.v1alpha1.Role
	(*Rule)(nil),          // 4: gconfig.okta.v1alpha1.Rule
}
var file_gconfig_okta_v1alpha1_gconfig_proto_depIdxs = []int32{
	1, // 0: gconfig.okta.v1alpha1.Config.admins:type_name -> gconfig.okta.v1alpha1.Member
	3, // 1: gconfig.okta.v1alpha1.Config.roles:type_name -> gconfig.okta.v1alpha1.Role
	2, // 2: gconfig.okta.v1alpha1.Config.accessHandler:type_name -> gconfig.okta.v1alpha1.AccessHandler
	4, // 3: gconfig.okta.v1alpha1.Role.rules:type_name -> gconfig.okta.v1alpha1.Rule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gconfig_okta_v1alpha1_gconfig_proto_init() }
func file_gconfig_okta_v1alpha1_gconfig_proto_init() {
	if File_gconfig_okta_v1alpha1_gconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessHandler); i {
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
		file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
			RawDescriptor: file_gconfig_okta_v1alpha1_gconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gconfig_okta_v1alpha1_gconfig_proto_goTypes,
		DependencyIndexes: file_gconfig_okta_v1alpha1_gconfig_proto_depIdxs,
		MessageInfos:      file_gconfig_okta_v1alpha1_gconfig_proto_msgTypes,
	}.Build()
	File_gconfig_okta_v1alpha1_gconfig_proto = out.File
	file_gconfig_okta_v1alpha1_gconfig_proto_rawDesc = nil
	file_gconfig_okta_v1alpha1_gconfig_proto_goTypes = nil
	file_gconfig_okta_v1alpha1_gconfig_proto_depIdxs = nil
}
