// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package pb_user_v1

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

type Index struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NullResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullResponse) Reset()         { *m = NullResponse{} }
func (m *NullResponse) String() string { return proto.CompactTextString(m) }
func (*NullResponse) ProtoMessage()    {}
func (*NullResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}

func (m *NullResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullResponse.Unmarshal(m, b)
}
func (m *NullResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullResponse.Marshal(b, m, deterministic)
}
func (m *NullResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullResponse.Merge(m, src)
}
func (m *NullResponse) XXX_Size() int {
	return xxx_messageInfo_NullResponse.Size(m)
}
func (m *NullResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NullResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NullResponse proto.InternalMessageInfo

func (m *NullResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*Index)(nil), "pb_user_v1.Index")
	proto.RegisterType((*NullResponse)(nil), "pb_user_v1.NullResponse")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0x48, 0x8a, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a,
	0x2f, 0x33, 0x54, 0x12, 0xe7, 0x62, 0xf5, 0xcc, 0x4b, 0x49, 0xad, 0x10, 0xe2, 0xe3, 0x62, 0xca,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0xca, 0x4c, 0x51, 0x52, 0xe2, 0xe2, 0xf1,
	0x2b, 0xcd, 0xc9, 0x09, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe2, 0x62, 0x49,
	0xce, 0x4f, 0x49, 0x05, 0xab, 0x60, 0x0d, 0x02, 0xb3, 0x93, 0xd8, 0xc0, 0xe6, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x3d, 0x47, 0xb9, 0x04, 0x5d, 0x00, 0x00, 0x00,
}
