// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ChangeSet struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Checksum             string   `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Format               string   `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	Source               string   `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeSet) Reset()         { *m = ChangeSet{} }
func (m *ChangeSet) String() string { return proto.CompactTextString(m) }
func (*ChangeSet) ProtoMessage()    {}
func (*ChangeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ChangeSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeSet.Unmarshal(m, b)
}
func (m *ChangeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeSet.Marshal(b, m, deterministic)
}
func (m *ChangeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeSet.Merge(m, src)
}
func (m *ChangeSet) XXX_Size() int {
	return xxx_messageInfo_ChangeSet.Size(m)
}
func (m *ChangeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeSet.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeSet proto.InternalMessageInfo

func (m *ChangeSet) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ChangeSet) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *ChangeSet) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *ChangeSet) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ChangeSet) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Change struct {
	Namespace            string     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path                 string     `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	ChangeSet            *ChangeSet `protobuf:"bytes,3,opt,name=changeSet,proto3" json:"changeSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Change) Reset()         { *m = Change{} }
func (m *Change) String() string { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()    {}
func (*Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Change) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Change.Unmarshal(m, b)
}
func (m *Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Change.Marshal(b, m, deterministic)
}
func (m *Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Change.Merge(m, src)
}
func (m *Change) XXX_Size() int {
	return xxx_messageInfo_Change.Size(m)
}
func (m *Change) XXX_DiscardUnknown() {
	xxx_messageInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_Change proto.InternalMessageInfo

func (m *Change) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Change) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Change) GetChangeSet() *ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

type CreateRequest struct {
	Change               *Change  `protobuf:"bytes,1,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type UpdateRequest struct {
	Change               *Change  `protobuf:"bytes,1,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Change               *Change  `protobuf:"bytes,1,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ListRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Suffix               string   `protobuf:"bytes,2,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Substr               string   `protobuf:"bytes,3,opt,name=substr,proto3" json:"substr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListRequest) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

func (m *ListRequest) GetSubstr() string {
	if m != nil {
		return m.Substr
	}
	return ""
}

type ListResponse struct {
	Values               []*Change `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetValues() []*Change {
	if m != nil {
		return m.Values
	}
	return nil
}

type ReadRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{10}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ReadRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ReadResponse struct {
	Change               *Change  `protobuf:"bytes,1,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{11}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type WatchRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{12}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *WatchRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type WatchResponse struct {
	Namespace            string     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path                 string     `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	ChangeSet            *ChangeSet `protobuf:"bytes,2,opt,name=changeSet,proto3" json:"changeSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{13}
}

func (m *WatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchResponse.Unmarshal(m, b)
}
func (m *WatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchResponse.Marshal(b, m, deterministic)
}
func (m *WatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchResponse.Merge(m, src)
}
func (m *WatchResponse) XXX_Size() int {
	return xxx_messageInfo_WatchResponse.Size(m)
}
func (m *WatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchResponse proto.InternalMessageInfo

func (m *WatchResponse) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *WatchResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *WatchResponse) GetChangeSet() *ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeSet)(nil), "ChangeSet")
	proto.RegisterType((*Change)(nil), "Change")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "ListRequest")
	proto.RegisterType((*ListResponse)(nil), "ListResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*WatchRequest)(nil), "WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "WatchResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xb5, 0x6c, 0x47, 0xad, 0xc7, 0x96, 0x5c, 0xe6, 0x50, 0x84, 0x28, 0xd4, 0x2c, 0x14, 0x4c,
	0x0b, 0xdb, 0xe0, 0xfe, 0x80, 0x16, 0xdc, 0x63, 0x4f, 0x2a, 0xa5, 0x87, 0x9e, 0x36, 0xf2, 0x38,
	0x12, 0x89, 0x3e, 0xaa, 0x5d, 0x85, 0xfe, 0x85, 0xfc, 0xeb, 0xb2, 0x1f, 0x8a, 0x24, 0xd3, 0x9a,
	0x24, 0x37, 0xcd, 0xdb, 0x9d, 0x37, 0x6f, 0xe7, 0x3d, 0x04, 0x81, 0xa4, 0xe6, 0x2e, 0x4f, 0x89,
	0xd7, 0x4d, 0xa5, 0x2a, 0x76, 0xef, 0xc1, 0x62, 0x9f, 0x89, 0xf2, 0x9a, 0xbe, 0x93, 0x42, 0x84,
	0xf9, 0x41, 0x28, 0x11, 0x79, 0x1b, 0x6f, 0xbb, 0x48, 0xcc, 0x37, 0xc6, 0xf0, 0x32, 0xcd, 0x28,
	0xbd, 0x91, 0x6d, 0x11, 0x4d, 0x0d, 0xfe, 0x50, 0xe3, 0x6b, 0xf0, 0x8f, 0x55, 0x53, 0x08, 0x15,
	0xcd, 0xcc, 0x89, 0xab, 0x34, 0x2e, 0xab, 0xb6, 0x49, 0x29, 0x9a, 0x5b, 0xdc, 0x56, 0xf8, 0x06,
	0x16, 0x2a, 0x2f, 0x48, 0x2a, 0x51, 0xd4, 0xd1, 0xc5, 0xc6, 0xdb, 0xce, 0x92, 0x1e, 0x60, 0x07,
	0xf0, 0xad, 0x14, 0x7d, 0xaf, 0x14, 0x05, 0xc9, 0x5a, 0xa4, 0xe4, 0xc4, 0xf4, 0x80, 0x56, 0x59,
	0x0b, 0x95, 0x39, 0x35, 0xe6, 0x1b, 0xb7, 0xb0, 0x48, 0xbb, 0x67, 0x18, 0x31, 0xcb, 0x1d, 0xf0,
	0x87, 0x87, 0x25, 0xfd, 0x21, 0xbb, 0x84, 0x60, 0xdf, 0x90, 0x50, 0x94, 0xd0, 0xef, 0x96, 0xa4,
	0xc2, 0xb7, 0xe0, 0xdb, 0x53, 0x33, 0x69, 0xb9, 0x7b, 0xe1, 0xfa, 0x12, 0x07, 0xb3, 0x57, 0x10,
	0x76, 0x1d, 0xb2, 0xae, 0x4a, 0x49, 0x9a, 0xe3, 0x47, 0x7d, 0x78, 0x22, 0x47, 0xd7, 0xd1, 0x73,
	0x7c, 0xa5, 0x5b, 0x7a, 0x1a, 0x47, 0xd7, 0xe1, 0x38, 0x7e, 0xc1, 0xf2, 0x5b, 0x2e, 0x55, 0xc7,
	0x70, 0x7e, 0x6d, 0xda, 0x94, 0xf6, 0x78, 0xcc, 0xff, 0xb8, 0xc5, 0xb9, 0xca, 0xe2, 0x57, 0x52,
	0x35, 0x9d, 0x89, 0xb6, 0x62, 0x1f, 0x61, 0x65, 0xc9, 0xed, 0x30, 0xad, 0xef, 0x4e, 0xdc, 0xb6,
	0x24, 0x23, 0x6f, 0x33, 0x1b, 0xe9, 0xb3, 0x30, 0xfb, 0x0c, 0xcb, 0x84, 0xc4, 0xe1, 0x71, 0x6a,
	0xfe, 0x61, 0xa2, 0x9e, 0x68, 0x09, 0xfa, 0x89, 0xe7, 0x37, 0xf2, 0x05, 0x56, 0x3f, 0x85, 0x4a,
	0xb3, 0xe7, 0x8f, 0xbc, 0x81, 0xc0, 0x31, 0xb8, 0x99, 0x8f, 0xa3, 0x98, 0xfd, 0x2f, 0x7a, 0xd3,
	0x33, 0xd1, 0xdb, 0xdd, 0x4f, 0xc1, 0xdf, 0x57, 0xe5, 0x31, 0xbf, 0xc6, 0x0f, 0xe0, 0xdb, 0x4c,
	0x61, 0xc8, 0x47, 0x71, 0x8c, 0xd7, 0xfc, 0x24, 0x6c, 0x13, 0x7d, 0xd9, 0x86, 0x07, 0x43, 0x3e,
	0xca, 0x5d, 0xbc, 0xe6, 0x27, 0xa9, 0x32, 0x97, 0x6d, 0x4a, 0x30, 0xe4, 0xa3, 0x80, 0xc5, 0x6b,
	0x7e, 0x12, 0x9f, 0x09, 0xbe, 0x83, 0xb9, 0xf6, 0x18, 0x57, 0x7c, 0x90, 0xa3, 0x38, 0xe0, 0x43,
	0xe3, 0xed, 0x35, 0x6d, 0x0c, 0xae, 0xf8, 0xc0, 0xe0, 0x38, 0xe0, 0x43, 0xb7, 0xd8, 0x04, 0xdf,
	0xc3, 0x85, 0x59, 0x26, 0x06, 0x7c, 0x68, 0x4b, 0x1c, 0xf2, 0xd1, 0x8e, 0xd9, 0xe4, 0xd2, 0xbb,
	0xf2, 0xcd, 0xff, 0xe7, 0xd3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xd1, 0x07, 0x74, 0x90,
	0x04, 0x00, 0x00,
}
