// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user-msg.proto

package pb_user_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AddUserRequest 添加用户请求
type UserProto struct {
	Id                   *Index   `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	LoginName            string   `protobuf:"bytes,2,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Mobile               int64    `protobuf:"varint,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	GroupId              int64    `protobuf:"varint,7,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	UserType             int64    `protobuf:"varint,8,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	RoleIds              []*Index `protobuf:"bytes,5,rep,name=role_ids,json=roleIds,proto3" json:"role_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProto) Reset()         { *m = UserProto{} }
func (m *UserProto) String() string { return proto.CompactTextString(m) }
func (*UserProto) ProtoMessage()    {}
func (*UserProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_msg_8cbbf50112d9cc7d, []int{0}
}
func (m *UserProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProto.Unmarshal(m, b)
}
func (m *UserProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProto.Marshal(b, m, deterministic)
}
func (dst *UserProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProto.Merge(dst, src)
}
func (m *UserProto) XXX_Size() int {
	return xxx_messageInfo_UserProto.Size(m)
}
func (m *UserProto) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProto.DiscardUnknown(m)
}

var xxx_messageInfo_UserProto proto.InternalMessageInfo

func (m *UserProto) GetId() *Index {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UserProto) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserProto) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *UserProto) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserProto) GetMobile() int64 {
	if m != nil {
		return m.Mobile
	}
	return 0
}

func (m *UserProto) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *UserProto) GetUserType() int64 {
	if m != nil {
		return m.UserType
	}
	return 0
}

func (m *UserProto) GetRoleIds() []*Index {
	if m != nil {
		return m.RoleIds
	}
	return nil
}

type Users struct {
	Users                []*Users `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_msg_8cbbf50112d9cc7d, []int{1}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*Users {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*UserProto)(nil), "pb_user_v1.UserProto")
	proto.RegisterType((*Users)(nil), "pb_user_v1.Users")
}

func init() { proto.RegisterFile("user-msg.proto", fileDescriptor_user_msg_8cbbf50112d9cc7d) }

var fileDescriptor_user_msg_8cbbf50112d9cc7d = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0xf4, 0x30,
	0x14, 0x85, 0x69, 0xfb, 0xb6, 0x4d, 0xef, 0x0b, 0x82, 0x59, 0x48, 0x1c, 0x11, 0xea, 0x6c, 0xec,
	0x42, 0x8b, 0x1f, 0xbf, 0xa2, 0x1b, 0x91, 0xa2, 0xeb, 0xd2, 0x92, 0x4b, 0x09, 0xb4, 0x4d, 0xc8,
	0x9d, 0x51, 0x67, 0xef, 0x0f, 0x97, 0xdc, 0x2a, 0xb3, 0xd0, 0xe5, 0x79, 0x9e, 0x73, 0x20, 0xb9,
	0x70, 0xb2, 0x27, 0xf4, 0xb7, 0x33, 0x8d, 0xb5, 0xf3, 0x76, 0x67, 0x25, 0xb8, 0xa1, 0x0b, 0xa8,
	0x7b, 0xbb, 0xdf, 0xc0, 0xd0, 0x13, 0xae, 0x7c, 0xfb, 0x19, 0x43, 0xf1, 0x4a, 0xe8, 0x9f, 0xb9,
	0x75, 0x05, 0xb1, 0xd1, 0x2a, 0x2b, 0xa3, 0xea, 0xff, 0xc3, 0x69, 0x7d, 0x9c, 0xd4, 0xcd, 0xa2,
	0xf1, 0xa3, 0x8d, 0x8d, 0x96, 0x17, 0x50, 0x30, 0x5c, 0xfa, 0x19, 0x55, 0x54, 0x46, 0x55, 0xd1,
	0x8a, 0x00, 0x9e, 0xfa, 0x19, 0xe5, 0x25, 0xc0, 0x64, 0x47, 0xb3, 0xac, 0x36, 0x66, 0x5b, 0x30,
	0x61, 0xbd, 0x01, 0xe1, 0x7a, 0xa2, 0x77, 0xeb, 0xb5, 0x4a, 0xd6, 0xe9, 0x4f, 0x96, 0x67, 0x90,
	0xcd, 0x76, 0x30, 0x13, 0xaa, 0x7f, 0x65, 0x54, 0x25, 0xed, 0x77, 0x92, 0xe7, 0x20, 0x46, 0x6f,
	0xf7, 0xae, 0x33, 0x5a, 0xe5, 0x6c, 0x72, 0xce, 0xcd, 0xf1, 0x29, 0xbb, 0x83, 0x43, 0x25, 0xd8,
	0xf1, 0x53, 0x5e, 0x0e, 0x0e, 0xe5, 0x0d, 0x08, 0x6f, 0x27, 0xec, 0x8c, 0x26, 0x95, 0x96, 0xc9,
	0xdf, 0x1f, 0xca, 0x43, 0xa5, 0xd1, 0xb4, 0xbd, 0x83, 0x34, 0x5c, 0x81, 0xe4, 0x35, 0xa4, 0xa1,
	0x42, 0x2a, 0xfa, 0xbd, 0xe1, 0x46, 0xbb, 0xfa, 0x21, 0xe3, 0xfb, 0x3d, 0x7e, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x56, 0xc2, 0xa0, 0x6e, 0x69, 0x01, 0x00, 0x00,
}
