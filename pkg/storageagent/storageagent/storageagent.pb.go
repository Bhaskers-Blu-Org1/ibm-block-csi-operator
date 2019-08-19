// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storageagent.proto

package storageagent

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Information about a specific host.
type Host struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Array                string   `protobuf:"bytes,5,opt,name=array,proto3" json:"array,omitempty"`
	Iqns                 []string `protobuf:"bytes,6,rep,name=iqns,proto3" json:"iqns,omitempty"`
	Wwpns                []string `protobuf:"bytes,7,rep,name=wwpns,proto3" json:"wwpns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{0}
}

func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Host) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Host) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Host) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Host) GetArray() string {
	if m != nil {
		return m.Array
	}
	return ""
}

func (m *Host) GetIqns() []string {
	if m != nil {
		return m.Iqns
	}
	return nil
}

func (m *Host) GetWwpns() []string {
	if m != nil {
		return m.Wwpns
	}
	return nil
}

// Information about a specific iscsi target, including ip and iqn,
// but only ip is mandatory since iqn can be discovered by iscsiadm.
type IscsiTarget struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Iqn                  string   `protobuf:"bytes,2,opt,name=iqn,proto3" json:"iqn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IscsiTarget) Reset()         { *m = IscsiTarget{} }
func (m *IscsiTarget) String() string { return proto.CompactTextString(m) }
func (*IscsiTarget) ProtoMessage()    {}
func (*IscsiTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{1}
}

func (m *IscsiTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IscsiTarget.Unmarshal(m, b)
}
func (m *IscsiTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IscsiTarget.Marshal(b, m, deterministic)
}
func (m *IscsiTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IscsiTarget.Merge(m, src)
}
func (m *IscsiTarget) XXX_Size() int {
	return xxx_messageInfo_IscsiTarget.Size(m)
}
func (m *IscsiTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_IscsiTarget.DiscardUnknown(m)
}

var xxx_messageInfo_IscsiTarget proto.InternalMessageInfo

func (m *IscsiTarget) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IscsiTarget) GetIqn() string {
	if m != nil {
		return m.Iqn
	}
	return ""
}

// type is optional
type CreateHostRequest struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Iqns                 []string          `protobuf:"bytes,3,rep,name=iqns,proto3" json:"iqns,omitempty"`
	Wwpns                []string          `protobuf:"bytes,4,rep,name=wwpns,proto3" json:"wwpns,omitempty"`
	Secrets              map[string]string `protobuf:"bytes,5,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateHostRequest) Reset()         { *m = CreateHostRequest{} }
func (m *CreateHostRequest) String() string { return proto.CompactTextString(m) }
func (*CreateHostRequest) ProtoMessage()    {}
func (*CreateHostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{2}
}

func (m *CreateHostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateHostRequest.Unmarshal(m, b)
}
func (m *CreateHostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateHostRequest.Marshal(b, m, deterministic)
}
func (m *CreateHostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateHostRequest.Merge(m, src)
}
func (m *CreateHostRequest) XXX_Size() int {
	return xxx_messageInfo_CreateHostRequest.Size(m)
}
func (m *CreateHostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateHostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateHostRequest proto.InternalMessageInfo

func (m *CreateHostRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateHostRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateHostRequest) GetIqns() []string {
	if m != nil {
		return m.Iqns
	}
	return nil
}

func (m *CreateHostRequest) GetWwpns() []string {
	if m != nil {
		return m.Wwpns
	}
	return nil
}

func (m *CreateHostRequest) GetSecrets() map[string]string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type CreateHostReply struct {
	Host                 *Host    `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateHostReply) Reset()         { *m = CreateHostReply{} }
func (m *CreateHostReply) String() string { return proto.CompactTextString(m) }
func (*CreateHostReply) ProtoMessage()    {}
func (*CreateHostReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{3}
}

func (m *CreateHostReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateHostReply.Unmarshal(m, b)
}
func (m *CreateHostReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateHostReply.Marshal(b, m, deterministic)
}
func (m *CreateHostReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateHostReply.Merge(m, src)
}
func (m *CreateHostReply) XXX_Size() int {
	return xxx_messageInfo_CreateHostReply.Size(m)
}
func (m *CreateHostReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateHostReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateHostReply proto.InternalMessageInfo

func (m *CreateHostReply) GetHost() *Host {
	if m != nil {
		return m.Host
	}
	return nil
}

// Either name or identifier is required
type DeleteHostRequest struct {
	// Types that are valid to be assigned to UniqueKey:
	//	*DeleteHostRequest_Identifier
	//	*DeleteHostRequest_Name
	UniqueKey            isDeleteHostRequest_UniqueKey `protobuf_oneof:"unique_key"`
	Secrets              map[string]string             `protobuf:"bytes,3,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *DeleteHostRequest) Reset()         { *m = DeleteHostRequest{} }
func (m *DeleteHostRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteHostRequest) ProtoMessage()    {}
func (*DeleteHostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{4}
}

func (m *DeleteHostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteHostRequest.Unmarshal(m, b)
}
func (m *DeleteHostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteHostRequest.Marshal(b, m, deterministic)
}
func (m *DeleteHostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteHostRequest.Merge(m, src)
}
func (m *DeleteHostRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteHostRequest.Size(m)
}
func (m *DeleteHostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteHostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteHostRequest proto.InternalMessageInfo

type isDeleteHostRequest_UniqueKey interface {
	isDeleteHostRequest_UniqueKey()
}

type DeleteHostRequest_Identifier struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3,oneof"`
}

type DeleteHostRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*DeleteHostRequest_Identifier) isDeleteHostRequest_UniqueKey() {}

func (*DeleteHostRequest_Name) isDeleteHostRequest_UniqueKey() {}

func (m *DeleteHostRequest) GetUniqueKey() isDeleteHostRequest_UniqueKey {
	if m != nil {
		return m.UniqueKey
	}
	return nil
}

func (m *DeleteHostRequest) GetIdentifier() string {
	if x, ok := m.GetUniqueKey().(*DeleteHostRequest_Identifier); ok {
		return x.Identifier
	}
	return ""
}

func (m *DeleteHostRequest) GetName() string {
	if x, ok := m.GetUniqueKey().(*DeleteHostRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (m *DeleteHostRequest) GetSecrets() map[string]string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DeleteHostRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DeleteHostRequest_Identifier)(nil),
		(*DeleteHostRequest_Name)(nil),
	}
}

type DeleteHostReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteHostReply) Reset()         { *m = DeleteHostReply{} }
func (m *DeleteHostReply) String() string { return proto.CompactTextString(m) }
func (*DeleteHostReply) ProtoMessage()    {}
func (*DeleteHostReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{5}
}

func (m *DeleteHostReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteHostReply.Unmarshal(m, b)
}
func (m *DeleteHostReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteHostReply.Marshal(b, m, deterministic)
}
func (m *DeleteHostReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteHostReply.Merge(m, src)
}
func (m *DeleteHostReply) XXX_Size() int {
	return xxx_messageInfo_DeleteHostReply.Size(m)
}
func (m *DeleteHostReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteHostReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteHostReply proto.InternalMessageInfo

type ListHostsRequest struct {
	// Either name or identifier should be specified.
	// The identifier and name are OPTIONAL, if not specified, return all.
	//
	// Types that are valid to be assigned to UniqueKey:
	//	*ListHostsRequest_Identifier
	//	*ListHostsRequest_Name
	UniqueKey            isListHostsRequest_UniqueKey `protobuf_oneof:"unique_key"`
	Secrets              map[string]string            `protobuf:"bytes,3,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListHostsRequest) Reset()         { *m = ListHostsRequest{} }
func (m *ListHostsRequest) String() string { return proto.CompactTextString(m) }
func (*ListHostsRequest) ProtoMessage()    {}
func (*ListHostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{6}
}

func (m *ListHostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHostsRequest.Unmarshal(m, b)
}
func (m *ListHostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHostsRequest.Marshal(b, m, deterministic)
}
func (m *ListHostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHostsRequest.Merge(m, src)
}
func (m *ListHostsRequest) XXX_Size() int {
	return xxx_messageInfo_ListHostsRequest.Size(m)
}
func (m *ListHostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListHostsRequest proto.InternalMessageInfo

type isListHostsRequest_UniqueKey interface {
	isListHostsRequest_UniqueKey()
}

type ListHostsRequest_Identifier struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3,oneof"`
}

type ListHostsRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*ListHostsRequest_Identifier) isListHostsRequest_UniqueKey() {}

func (*ListHostsRequest_Name) isListHostsRequest_UniqueKey() {}

func (m *ListHostsRequest) GetUniqueKey() isListHostsRequest_UniqueKey {
	if m != nil {
		return m.UniqueKey
	}
	return nil
}

func (m *ListHostsRequest) GetIdentifier() string {
	if x, ok := m.GetUniqueKey().(*ListHostsRequest_Identifier); ok {
		return x.Identifier
	}
	return ""
}

func (m *ListHostsRequest) GetName() string {
	if x, ok := m.GetUniqueKey().(*ListHostsRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (m *ListHostsRequest) GetSecrets() map[string]string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListHostsRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListHostsRequest_Identifier)(nil),
		(*ListHostsRequest_Name)(nil),
	}
}

type ListHostsReply struct {
	Hosts                []*Host  `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListHostsReply) Reset()         { *m = ListHostsReply{} }
func (m *ListHostsReply) String() string { return proto.CompactTextString(m) }
func (*ListHostsReply) ProtoMessage()    {}
func (*ListHostsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{7}
}

func (m *ListHostsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHostsReply.Unmarshal(m, b)
}
func (m *ListHostsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHostsReply.Marshal(b, m, deterministic)
}
func (m *ListHostsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHostsReply.Merge(m, src)
}
func (m *ListHostsReply) XXX_Size() int {
	return xxx_messageInfo_ListHostsReply.Size(m)
}
func (m *ListHostsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHostsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListHostsReply proto.InternalMessageInfo

func (m *ListHostsReply) GetHosts() []*Host {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type ListIscsiTargetsRequest struct {
	Secrets              map[string]string `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListIscsiTargetsRequest) Reset()         { *m = ListIscsiTargetsRequest{} }
func (m *ListIscsiTargetsRequest) String() string { return proto.CompactTextString(m) }
func (*ListIscsiTargetsRequest) ProtoMessage()    {}
func (*ListIscsiTargetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{8}
}

func (m *ListIscsiTargetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIscsiTargetsRequest.Unmarshal(m, b)
}
func (m *ListIscsiTargetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIscsiTargetsRequest.Marshal(b, m, deterministic)
}
func (m *ListIscsiTargetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIscsiTargetsRequest.Merge(m, src)
}
func (m *ListIscsiTargetsRequest) XXX_Size() int {
	return xxx_messageInfo_ListIscsiTargetsRequest.Size(m)
}
func (m *ListIscsiTargetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIscsiTargetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListIscsiTargetsRequest proto.InternalMessageInfo

func (m *ListIscsiTargetsRequest) GetSecrets() map[string]string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type ListIscsiTargetsReply struct {
	Targets              []*IscsiTarget `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListIscsiTargetsReply) Reset()         { *m = ListIscsiTargetsReply{} }
func (m *ListIscsiTargetsReply) String() string { return proto.CompactTextString(m) }
func (*ListIscsiTargetsReply) ProtoMessage()    {}
func (*ListIscsiTargetsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9012731fe56366, []int{9}
}

func (m *ListIscsiTargetsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIscsiTargetsReply.Unmarshal(m, b)
}
func (m *ListIscsiTargetsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIscsiTargetsReply.Marshal(b, m, deterministic)
}
func (m *ListIscsiTargetsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIscsiTargetsReply.Merge(m, src)
}
func (m *ListIscsiTargetsReply) XXX_Size() int {
	return xxx_messageInfo_ListIscsiTargetsReply.Size(m)
}
func (m *ListIscsiTargetsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIscsiTargetsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListIscsiTargetsReply proto.InternalMessageInfo

func (m *ListIscsiTargetsReply) GetTargets() []*IscsiTarget {
	if m != nil {
		return m.Targets
	}
	return nil
}

func init() {
	proto.RegisterType((*Host)(nil), "storageagent.Host")
	proto.RegisterType((*IscsiTarget)(nil), "storageagent.IscsiTarget")
	proto.RegisterType((*CreateHostRequest)(nil), "storageagent.CreateHostRequest")
	proto.RegisterMapType((map[string]string)(nil), "storageagent.CreateHostRequest.SecretsEntry")
	proto.RegisterType((*CreateHostReply)(nil), "storageagent.CreateHostReply")
	proto.RegisterType((*DeleteHostRequest)(nil), "storageagent.DeleteHostRequest")
	proto.RegisterMapType((map[string]string)(nil), "storageagent.DeleteHostRequest.SecretsEntry")
	proto.RegisterType((*DeleteHostReply)(nil), "storageagent.DeleteHostReply")
	proto.RegisterType((*ListHostsRequest)(nil), "storageagent.ListHostsRequest")
	proto.RegisterMapType((map[string]string)(nil), "storageagent.ListHostsRequest.SecretsEntry")
	proto.RegisterType((*ListHostsReply)(nil), "storageagent.ListHostsReply")
	proto.RegisterType((*ListIscsiTargetsRequest)(nil), "storageagent.ListIscsiTargetsRequest")
	proto.RegisterMapType((map[string]string)(nil), "storageagent.ListIscsiTargetsRequest.SecretsEntry")
	proto.RegisterType((*ListIscsiTargetsReply)(nil), "storageagent.ListIscsiTargetsReply")
}

func init() { proto.RegisterFile("storageagent.proto", fileDescriptor_ed9012731fe56366) }

var fileDescriptor_ed9012731fe56366 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x6f, 0x6f, 0x12, 0x4f,
	0x10, 0x66, 0x39, 0xfe, 0xa4, 0x03, 0xf9, 0xb5, 0x6c, 0xf8, 0xe9, 0x49, 0xb4, 0x92, 0x33, 0x1a,
	0x12, 0x0d, 0x89, 0xf4, 0x8d, 0xe5, 0x9d, 0x68, 0x0d, 0x46, 0x62, 0x1a, 0x6a, 0xe2, 0x4b, 0xdd,
	0xc2, 0x8a, 0x97, 0xc2, 0x1d, 0xec, 0x2e, 0x36, 0xf7, 0x75, 0x34, 0xf1, 0x4b, 0x99, 0x7e, 0x00,
	0xbf, 0x85, 0x99, 0xdd, 0x03, 0xf6, 0xee, 0xc0, 0x9a, 0x68, 0xdf, 0xed, 0xcc, 0x3e, 0x3b, 0xf3,
	0x3c, 0x73, 0xcf, 0xe4, 0x80, 0x4a, 0x15, 0x0a, 0x36, 0xe1, 0x6c, 0xc2, 0x03, 0xd5, 0x9e, 0x8b,
	0x50, 0x85, 0xb4, 0x6a, 0xe7, 0xbc, 0xef, 0x04, 0x0a, 0xfd, 0x50, 0x2a, 0x7a, 0x08, 0xe0, 0x8f,
	0x79, 0xa0, 0xfc, 0x4f, 0x3e, 0x17, 0x2e, 0x69, 0x92, 0xd6, 0xde, 0xd0, 0xca, 0x50, 0x0a, 0x85,
	0x80, 0xcd, 0xb8, 0x9b, 0xd7, 0x37, 0xfa, 0x8c, 0x39, 0x15, 0xcd, 0xb9, 0xeb, 0x98, 0x1c, 0x9e,
	0xe9, 0x2d, 0x28, 0x49, 0xc5, 0xd4, 0x52, 0xba, 0x05, 0x9d, 0x8d, 0x23, 0x5a, 0x87, 0x22, 0x13,
	0x82, 0x45, 0x6e, 0x51, 0xa7, 0x4d, 0x80, 0x15, 0xfc, 0x45, 0x20, 0xdd, 0x52, 0xd3, 0xc1, 0x0a,
	0x78, 0x46, 0xe4, 0xe5, 0xe5, 0x3c, 0x90, 0x6e, 0x59, 0x27, 0x4d, 0xe0, 0x1d, 0x43, 0xe5, 0xb5,
	0x1c, 0x49, 0xff, 0x1d, 0x13, 0x13, 0xae, 0xa8, 0x0b, 0x65, 0x36, 0x1e, 0x0b, 0x2e, 0x65, 0xcc,
	0x75, 0x15, 0xd2, 0x03, 0x70, 0xfc, 0x45, 0x10, 0xf3, 0xc4, 0xa3, 0xf7, 0x93, 0x40, 0xed, 0x85,
	0xe0, 0x4c, 0x71, 0x54, 0x3a, 0xe4, 0x8b, 0x25, 0x97, 0x6a, 0x2d, 0x88, 0x6c, 0x11, 0x94, 0xb7,
	0x04, 0xad, 0x28, 0x3a, 0xdb, 0x28, 0x16, 0x2c, 0x8a, 0xf4, 0x15, 0x94, 0x25, 0x1f, 0x09, 0xae,
	0xa4, 0x5b, 0x6c, 0x3a, 0xad, 0x4a, 0xe7, 0x49, 0x3b, 0x31, 0xff, 0x0c, 0x87, 0xf6, 0x99, 0x81,
	0x9f, 0x04, 0x4a, 0x44, 0xc3, 0xd5, 0xe3, 0x46, 0x17, 0xaa, 0xf6, 0x05, 0x2a, 0xba, 0xe0, 0x51,
	0x4c, 0x14, 0x8f, 0xd8, 0xff, 0x0b, 0x9b, 0x2e, 0x57, 0x44, 0x4d, 0xd0, 0xcd, 0x3f, 0x23, 0xde,
	0x31, 0xec, 0xdb, 0x6d, 0xe6, 0xd3, 0x88, 0x3e, 0x82, 0xc2, 0xe7, 0x50, 0x2a, 0xfd, 0xbe, 0xd2,
	0xa1, 0x49, 0x4e, 0x1a, 0xa6, 0xef, 0xbd, 0x2b, 0x02, 0xb5, 0x97, 0x7c, 0xca, 0x93, 0x63, 0x6a,
	0x66, 0x7d, 0xd1, 0xcf, 0x25, 0x9c, 0x51, 0xb7, 0x9d, 0xd1, 0xcf, 0xc5, 0xa3, 0xb4, 0x86, 0xe1,
	0x6c, 0x1b, 0x46, 0xa6, 0xd3, 0xbf, 0x1f, 0x46, 0xaf, 0x0a, 0xb0, 0x0c, 0xfc, 0xc5, 0x92, 0x7f,
	0xb8, 0xe0, 0x91, 0x57, 0x83, 0x7d, 0xbb, 0xe9, 0x7c, 0x1a, 0x79, 0x3f, 0x08, 0x1c, 0x0c, 0x7c,
	0xa9, 0x30, 0x23, 0xff, 0x56, 0xf1, 0x49, 0x5a, 0xf1, 0xe3, 0xa4, 0xe2, 0x74, 0xa3, 0x1b, 0x17,
	0xdc, 0x85, 0xff, 0xac, 0x9e, 0x68, 0x85, 0x16, 0x14, 0xf1, 0x53, 0xe3, 0xce, 0x38, 0x3b, 0xbc,
	0x60, 0x00, 0xde, 0x37, 0x02, 0xb7, 0xf1, 0xb1, 0xb5, 0x73, 0xeb, 0x01, 0x0d, 0x36, 0x42, 0x4d,
	0x9d, 0x4e, 0x56, 0xe8, 0x96, 0x77, 0x37, 0xe0, 0xf6, 0x01, 0xfc, 0x9f, 0x6d, 0x86, 0x42, 0x8f,
	0xa0, 0xac, 0x4c, 0x1c, 0x53, 0xbc, 0x93, 0xa4, 0x68, 0xbd, 0x18, 0xae, 0x90, 0x9d, 0xab, 0x3c,
	0x54, 0xcf, 0x0c, 0xea, 0x39, 0xa2, 0xe8, 0x5b, 0x80, 0xcd, 0x32, 0xd1, 0xfb, 0xd7, 0x6c, 0x73,
	0xe3, 0xde, 0x6e, 0x00, 0x9a, 0x2d, 0x87, 0xf5, 0x36, 0x0e, 0x4c, 0xd7, 0xcb, 0x2c, 0x44, 0xba,
	0x5e, 0xda, 0xbc, 0x39, 0xfa, 0x06, 0xf6, 0xd6, 0x1f, 0x98, 0x1e, 0xfe, 0xde, 0x6d, 0x8d, 0xbb,
	0x3b, 0xef, 0x4d, 0xb1, 0x8f, 0x66, 0x15, 0xec, 0x59, 0xd2, 0x87, 0x7f, 0xf4, 0x61, 0x1b, 0x0f,
	0xae, 0x83, 0xe9, 0x0e, 0xbd, 0xa7, 0x50, 0x1f, 0x85, 0xb3, 0xb6, 0x7f, 0x3e, 0x4b, 0xe0, 0x7b,
	0x35, 0x7b, 0xe8, 0xa7, 0xf8, 0x93, 0x3a, 0x25, 0x5f, 0xf3, 0x4e, 0x7f, 0xf0, 0xfe, 0xbc, 0xa4,
	0xff, 0x59, 0x47, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x27, 0x03, 0x04, 0xc9, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageAgentClient is the client API for StorageAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageAgentClient interface {
	// Define a host on storage side.
	CreateHost(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*CreateHostReply, error)
	DeleteHost(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*DeleteHostReply, error)
	ListHosts(ctx context.Context, in *ListHostsRequest, opts ...grpc.CallOption) (*ListHostsReply, error)
	ListIscsiTargets(ctx context.Context, in *ListIscsiTargetsRequest, opts ...grpc.CallOption) (*ListIscsiTargetsReply, error)
}

type storageAgentClient struct {
	cc *grpc.ClientConn
}

func NewStorageAgentClient(cc *grpc.ClientConn) StorageAgentClient {
	return &storageAgentClient{cc}
}

func (c *storageAgentClient) CreateHost(ctx context.Context, in *CreateHostRequest, opts ...grpc.CallOption) (*CreateHostReply, error) {
	out := new(CreateHostReply)
	err := c.cc.Invoke(ctx, "/storageagent.StorageAgent/CreateHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageAgentClient) DeleteHost(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*DeleteHostReply, error) {
	out := new(DeleteHostReply)
	err := c.cc.Invoke(ctx, "/storageagent.StorageAgent/DeleteHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageAgentClient) ListHosts(ctx context.Context, in *ListHostsRequest, opts ...grpc.CallOption) (*ListHostsReply, error) {
	out := new(ListHostsReply)
	err := c.cc.Invoke(ctx, "/storageagent.StorageAgent/ListHosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageAgentClient) ListIscsiTargets(ctx context.Context, in *ListIscsiTargetsRequest, opts ...grpc.CallOption) (*ListIscsiTargetsReply, error) {
	out := new(ListIscsiTargetsReply)
	err := c.cc.Invoke(ctx, "/storageagent.StorageAgent/ListIscsiTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageAgentServer is the server API for StorageAgent service.
type StorageAgentServer interface {
	// Define a host on storage side.
	CreateHost(context.Context, *CreateHostRequest) (*CreateHostReply, error)
	DeleteHost(context.Context, *DeleteHostRequest) (*DeleteHostReply, error)
	ListHosts(context.Context, *ListHostsRequest) (*ListHostsReply, error)
	ListIscsiTargets(context.Context, *ListIscsiTargetsRequest) (*ListIscsiTargetsReply, error)
}

// UnimplementedStorageAgentServer can be embedded to have forward compatible implementations.
type UnimplementedStorageAgentServer struct {
}

func (*UnimplementedStorageAgentServer) CreateHost(ctx context.Context, req *CreateHostRequest) (*CreateHostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHost not implemented")
}
func (*UnimplementedStorageAgentServer) DeleteHost(ctx context.Context, req *DeleteHostRequest) (*DeleteHostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHost not implemented")
}
func (*UnimplementedStorageAgentServer) ListHosts(ctx context.Context, req *ListHostsRequest) (*ListHostsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHosts not implemented")
}
func (*UnimplementedStorageAgentServer) ListIscsiTargets(ctx context.Context, req *ListIscsiTargetsRequest) (*ListIscsiTargetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIscsiTargets not implemented")
}

func RegisterStorageAgentServer(s *grpc.Server, srv StorageAgentServer) {
	s.RegisterService(&_StorageAgent_serviceDesc, srv)
}

func _StorageAgent_CreateHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageAgentServer).CreateHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageagent.StorageAgent/CreateHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageAgentServer).CreateHost(ctx, req.(*CreateHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageAgent_DeleteHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageAgentServer).DeleteHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageagent.StorageAgent/DeleteHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageAgentServer).DeleteHost(ctx, req.(*DeleteHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageAgent_ListHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageAgentServer).ListHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageagent.StorageAgent/ListHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageAgentServer).ListHosts(ctx, req.(*ListHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageAgent_ListIscsiTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIscsiTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageAgentServer).ListIscsiTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageagent.StorageAgent/ListIscsiTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageAgentServer).ListIscsiTargets(ctx, req.(*ListIscsiTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storageagent.StorageAgent",
	HandlerType: (*StorageAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHost",
			Handler:    _StorageAgent_CreateHost_Handler,
		},
		{
			MethodName: "DeleteHost",
			Handler:    _StorageAgent_DeleteHost_Handler,
		},
		{
			MethodName: "ListHosts",
			Handler:    _StorageAgent_ListHosts_Handler,
		},
		{
			MethodName: "ListIscsiTargets",
			Handler:    _StorageAgent_ListIscsiTargets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storageagent.proto",
}
