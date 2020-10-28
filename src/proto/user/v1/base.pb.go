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

type Page struct {
	Total                int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	TotalPage            int64    `protobuf:"varint,2,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	PageSize             int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNum              int64    `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{2}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Page) GetTotalPage() int64 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

func (m *Page) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Page) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func init() {
	proto.RegisterType((*Index)(nil), "pb_user_v1.Index")
	proto.RegisterType((*NullResponse)(nil), "pb_user_v1.NullResponse")
	proto.RegisterType((*Page)(nil), "pb_user_v1.Page")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0xcd, 0xaa, 0xc2, 0x30,
	0x10, 0x85, 0xe9, 0xdf, 0xbd, 0xb7, 0xc3, 0xc5, 0xc5, 0x20, 0x18, 0x11, 0x41, 0xb2, 0x72, 0x25,
	0x88, 0x4f, 0xe1, 0xa6, 0x48, 0x7d, 0x80, 0x90, 0x9a, 0xa1, 0x04, 0xda, 0x26, 0x34, 0x8d, 0x48,
	0x9f, 0x5e, 0x3a, 0xba, 0xfb, 0xce, 0xf9, 0xe6, 0x2c, 0x06, 0xa0, 0xd1, 0x81, 0x4e, 0x7e, 0x74,
	0x93, 0x43, 0xf0, 0x8d, 0x8a, 0x81, 0x46, 0xf5, 0x3c, 0xcb, 0x0d, 0x14, 0xd7, 0xc1, 0xd0, 0x0b,
	0x57, 0x90, 0x5a, 0x23, 0x92, 0x43, 0x72, 0xcc, 0xea, 0xd4, 0x1a, 0x29, 0xe1, 0xbf, 0x8a, 0x5d,
	0x57, 0x53, 0xf0, 0x6e, 0x08, 0x84, 0x08, 0xf9, 0xc3, 0x19, 0xe2, 0x8b, 0xa2, 0x66, 0x96, 0x01,
	0xf2, 0x9b, 0x6e, 0x09, 0xd7, 0x50, 0x4c, 0x6e, 0xd2, 0xdd, 0x77, 0xfe, 0x09, 0xb8, 0x07, 0x60,
	0x50, 0x5e, 0xb7, 0x24, 0x52, 0x56, 0x25, 0x37, 0x3c, 0xda, 0x41, 0xb9, 0x08, 0x15, 0xec, 0x4c,
	0x22, 0x63, 0xfb, 0xb7, 0x14, 0x77, 0x3b, 0x13, 0x6e, 0x81, 0x59, 0x0d, 0xb1, 0x17, 0x39, 0xbb,
	0xdf, 0x25, 0x57, 0xb1, 0x6f, 0x7e, 0xf8, 0x89, 0xcb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x4e,
	0x13, 0x9b, 0xd2, 0x00, 0x00, 0x00,
}