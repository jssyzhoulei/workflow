// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role-msg.proto

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

type RoleProto struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	DataPermit           int32    `protobuf:"varint,3,opt,name=data_permit,json=dataPermit,proto3" json:"data_permit,omitempty"`
	Status               int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Id                   int64    `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Ids                  string   `protobuf:"bytes,7,opt,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleProto) Reset()         { *m = RoleProto{} }
func (m *RoleProto) String() string { return proto.CompactTextString(m) }
func (*RoleProto) ProtoMessage()    {}
func (*RoleProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a27634faf0e2c2b, []int{0}
}

func (m *RoleProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleProto.Unmarshal(m, b)
}
func (m *RoleProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleProto.Marshal(b, m, deterministic)
}
func (m *RoleProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleProto.Merge(m, src)
}
func (m *RoleProto) XXX_Size() int {
	return xxx_messageInfo_RoleProto.Size(m)
}
func (m *RoleProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleProto.DiscardUnknown(m)
}

var xxx_messageInfo_RoleProto proto.InternalMessageInfo

func (m *RoleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleProto) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *RoleProto) GetDataPermit() int32 {
	if m != nil {
		return m.DataPermit
	}
	return 0
}

func (m *RoleProto) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RoleProto) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RoleProto) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *RoleProto) GetIds() string {
	if m != nil {
		return m.Ids
	}
	return ""
}

type RoleMenuPermissionProto struct {
	RoleId               int64    `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	MenuId               int64    `protobuf:"varint,2,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	PermissionId         int64    `protobuf:"varint,3,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleMenuPermissionProto) Reset()         { *m = RoleMenuPermissionProto{} }
func (m *RoleMenuPermissionProto) String() string { return proto.CompactTextString(m) }
func (*RoleMenuPermissionProto) ProtoMessage()    {}
func (*RoleMenuPermissionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a27634faf0e2c2b, []int{1}
}

func (m *RoleMenuPermissionProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleMenuPermissionProto.Unmarshal(m, b)
}
func (m *RoleMenuPermissionProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleMenuPermissionProto.Marshal(b, m, deterministic)
}
func (m *RoleMenuPermissionProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleMenuPermissionProto.Merge(m, src)
}
func (m *RoleMenuPermissionProto) XXX_Size() int {
	return xxx_messageInfo_RoleMenuPermissionProto.Size(m)
}
func (m *RoleMenuPermissionProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleMenuPermissionProto.DiscardUnknown(m)
}

var xxx_messageInfo_RoleMenuPermissionProto proto.InternalMessageInfo

func (m *RoleMenuPermissionProto) GetRoleId() int64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *RoleMenuPermissionProto) GetMenuId() int64 {
	if m != nil {
		return m.MenuId
	}
	return 0
}

func (m *RoleMenuPermissionProto) GetPermissionId() int64 {
	if m != nil {
		return m.PermissionId
	}
	return 0
}

type CreateMenuPermRequestProto struct {
	Id                   int64                      `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string                     `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	DataPermit           int32                      `protobuf:"varint,3,opt,name=data_permit,json=dataPermit,proto3" json:"data_permit,omitempty"`
	Status               int32                      `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	RoleMenuPermissions  []*RoleMenuPermissionProto `protobuf:"bytes,5,rep,name=role_menu_permissions,json=roleMenuPermissions,proto3" json:"role_menu_permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CreateMenuPermRequestProto) Reset()         { *m = CreateMenuPermRequestProto{} }
func (m *CreateMenuPermRequestProto) String() string { return proto.CompactTextString(m) }
func (*CreateMenuPermRequestProto) ProtoMessage()    {}
func (*CreateMenuPermRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a27634faf0e2c2b, []int{2}
}

func (m *CreateMenuPermRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMenuPermRequestProto.Unmarshal(m, b)
}
func (m *CreateMenuPermRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMenuPermRequestProto.Marshal(b, m, deterministic)
}
func (m *CreateMenuPermRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMenuPermRequestProto.Merge(m, src)
}
func (m *CreateMenuPermRequestProto) XXX_Size() int {
	return xxx_messageInfo_CreateMenuPermRequestProto.Size(m)
}
func (m *CreateMenuPermRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMenuPermRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMenuPermRequestProto proto.InternalMessageInfo

func (m *CreateMenuPermRequestProto) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateMenuPermRequestProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateMenuPermRequestProto) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *CreateMenuPermRequestProto) GetDataPermit() int32 {
	if m != nil {
		return m.DataPermit
	}
	return 0
}

func (m *CreateMenuPermRequestProto) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateMenuPermRequestProto) GetRoleMenuPermissions() []*RoleMenuPermissionProto {
	if m != nil {
		return m.RoleMenuPermissions
	}
	return nil
}

type RolePageRequestProto struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Page                 *Page        `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Roles                []*RoleProto `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RolePageRequestProto) Reset()         { *m = RolePageRequestProto{} }
func (m *RolePageRequestProto) String() string { return proto.CompactTextString(m) }
func (*RolePageRequestProto) ProtoMessage()    {}
func (*RolePageRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a27634faf0e2c2b, []int{3}
}

func (m *RolePageRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolePageRequestProto.Unmarshal(m, b)
}
func (m *RolePageRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolePageRequestProto.Marshal(b, m, deterministic)
}
func (m *RolePageRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolePageRequestProto.Merge(m, src)
}
func (m *RolePageRequestProto) XXX_Size() int {
	return xxx_messageInfo_RolePageRequestProto.Size(m)
}
func (m *RolePageRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RolePageRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_RolePageRequestProto proto.InternalMessageInfo

func (m *RolePageRequestProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RolePageRequestProto) GetPage() *Page {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *RolePageRequestProto) GetRoles() []*RoleProto {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*RoleProto)(nil), "pb_user_v1.RoleProto")
	proto.RegisterType((*RoleMenuPermissionProto)(nil), "pb_user_v1.RoleMenuPermissionProto")
	proto.RegisterType((*CreateMenuPermRequestProto)(nil), "pb_user_v1.CreateMenuPermRequestProto")
	proto.RegisterType((*RolePageRequestProto)(nil), "pb_user_v1.RolePageRequestProto")
}

func init() { proto.RegisterFile("role-msg.proto", fileDescriptor_3a27634faf0e2c2b) }

var fileDescriptor_3a27634faf0e2c2b = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x49, 0xd2, 0xa4, 0xf4, 0x56, 0x4b, 0x19, 0xad, 0x1d, 0x0a, 0x62, 0x48, 0x5d, 0x04,
	0xc4, 0x82, 0xf5, 0x09, 0xc4, 0x55, 0x17, 0x42, 0xc9, 0xc6, 0x65, 0x98, 0x3a, 0x97, 0x12, 0x6c,
	0x7e, 0x9c, 0x99, 0xb8, 0xf3, 0xa1, 0x7c, 0x2d, 0x9f, 0x42, 0xe6, 0x4e, 0x6d, 0xfd, 0xdd, 0xba,
	0x9b, 0x7b, 0xce, 0x4c, 0xce, 0xf9, 0x26, 0x03, 0x03, 0x55, 0x6f, 0xf0, 0xb2, 0xd4, 0xeb, 0x59,
	0xa3, 0x6a, 0x53, 0x33, 0x68, 0x56, 0x79, 0xab, 0x51, 0xe5, 0xcf, 0x57, 0x13, 0x58, 0x09, 0x8d,
	0x4e, 0x4f, 0x5e, 0x3d, 0xe8, 0x65, 0xf5, 0x06, 0x97, 0xb4, 0x8b, 0x41, 0xa7, 0x12, 0x25, 0x72,
	0x2f, 0xf6, 0xd2, 0x5e, 0x46, 0x6b, 0x76, 0x02, 0x91, 0xc2, 0x52, 0xa8, 0x47, 0xee, 0x93, 0xba,
	0x9d, 0xd8, 0x19, 0xf4, 0xa5, 0x30, 0x22, 0x6f, 0x50, 0x95, 0x85, 0xe1, 0x41, 0xec, 0xa5, 0x61,
	0x06, 0x56, 0x5a, 0x92, 0x62, 0x0f, 0x6a, 0x23, 0x4c, 0xab, 0x79, 0x87, 0xbc, 0xed, 0xc4, 0x06,
	0xe0, 0x17, 0x92, 0x87, 0xb1, 0x97, 0x06, 0x99, 0x5f, 0x48, 0x76, 0x0a, 0xf0, 0xa0, 0x50, 0x18,
	0x94, 0xb9, 0x30, 0x3c, 0xa2, 0x90, 0xde, 0x56, 0xb9, 0x31, 0x6c, 0x08, 0x41, 0x21, 0x35, 0xef,
	0x92, 0x6e, 0x97, 0x49, 0x03, 0x63, 0x5b, 0xf9, 0x0e, 0xab, 0x96, 0xa2, 0xb4, 0x2e, 0xea, 0xca,
	0x01, 0x8c, 0xa1, 0x6b, 0xc1, 0xf3, 0x42, 0x12, 0x43, 0x90, 0x45, 0x76, 0x5c, 0x48, 0x6b, 0x94,
	0x58, 0xb5, 0xd6, 0xf0, 0x9d, 0x61, 0xc7, 0x85, 0x64, 0x53, 0x38, 0x6c, 0x76, 0x1f, 0xb1, 0x76,
	0x40, 0xf6, 0xc1, 0x5e, 0x5c, 0xc8, 0xe4, 0xcd, 0x83, 0xc9, 0x2d, 0x35, 0xfa, 0x08, 0xcd, 0xf0,
	0xa9, 0x45, 0x6d, 0x5c, 0xaa, 0x23, 0x8a, 0x76, 0x44, 0xff, 0x72, 0x8d, 0xf7, 0x30, 0x22, 0x54,
	0xc2, 0xda, 0xb7, 0xd5, 0x3c, 0x8c, 0x83, 0xb4, 0x3f, 0x9f, 0xce, 0xf6, 0x7f, 0x7c, 0xf6, 0xc7,
	0x75, 0x65, 0x47, 0xea, 0x87, 0xa1, 0x93, 0x17, 0x38, 0xa6, 0x17, 0x21, 0xd6, 0xf8, 0x85, 0xf2,
	0x37, 0xaa, 0x73, 0xe8, 0x34, 0x62, 0x8d, 0xc4, 0xd4, 0x9f, 0x0f, 0x3f, 0x67, 0xd2, 0x79, 0x72,
	0xd9, 0x05, 0x84, 0x36, 0x48, 0xf3, 0x80, 0xaa, 0x8d, 0xbe, 0x57, 0x73, 0x65, 0xdc, 0x9e, 0x55,
	0x44, 0x0f, 0xf3, 0xfa, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x06, 0xe7, 0x57, 0x79, 0xc2, 0x02, 0x00,
	0x00,
}
