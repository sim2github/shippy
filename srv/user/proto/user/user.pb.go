// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "shippy.user.User")
	proto.RegisterType((*Request)(nil), "shippy.user.Request")
	proto.RegisterType((*Response)(nil), "shippy.user.Response")
	proto.RegisterType((*Token)(nil), "shippy.user.Token")
	proto.RegisterType((*Error)(nil), "shippy.user.Error")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xb5, 0xf5, 0x65, 0x7b, 0x44, 0x0b, 0x1d, 0x6c, 0x58, 0x7c, 0x32, 0x82, 0xd2, 0xd2, 0x83,
	0x4c, 0xdd, 0x42, 0xe9, 0xa1, 0x07, 0xb7, 0x14, 0xdf, 0xd5, 0x8f, 0x43, 0x2f, 0x45, 0x91, 0x06,
	0xbc, 0x44, 0xd6, 0x2a, 0xbb, 0x6b, 0x07, 0xff, 0x81, 0xfc, 0x89, 0xfc, 0xd9, 0xb0, 0x23, 0x3b,
	0xd8, 0x89, 0x72, 0xf0, 0x45, 0x9a, 0xf7, 0xe6, 0xcd, 0xee, 0xcc, 0xdb, 0x81, 0x49, 0xa3, 0x95,
	0x55, 0xf3, 0xad, 0x21, 0xcd, 0x9f, 0x94, 0x31, 0xc6, 0x66, 0x2d, 0x9b, 0x66, 0x9f, 0x3a, 0x2a,
	0xd9, 0x41, 0xf0, 0xc7, 0x90, 0xc6, 0xd7, 0xe0, 0xc9, 0x52, 0xf4, 0x67, 0xfd, 0xf7, 0xa3, 0xcc,
	0x93, 0x25, 0x22, 0x04, 0x75, 0xbe, 0x21, 0xe1, 0x31, 0xc3, 0x31, 0x0a, 0x18, 0x14, 0x6a, 0xd3,
	0xe4, 0xf5, 0x5e, 0xf8, 0x4c, 0x1f, 0x21, 0x8e, 0x21, 0xa4, 0x4d, 0x2e, 0x2b, 0x11, 0x30, 0xdf,
	0x02, 0x9c, 0xc2, 0xb0, 0xc9, 0x8d, 0xb9, 0x55, 0xba, 0x14, 0x21, 0x27, 0x1e, 0x71, 0x32, 0x82,
	0x41, 0x46, 0x37, 0x5b, 0x32, 0x36, 0xb9, 0xeb, 0xc3, 0x30, 0x23, 0xd3, 0xa8, 0xda, 0x10, 0xbe,
	0x85, 0xc0, 0xf5, 0xc5, 0x9d, 0xc4, 0x8b, 0x37, 0xe9, 0x49, 0xaf, 0xa9, 0x6b, 0x34, 0xe3, 0x34,
	0xbe, 0x83, 0xd0, 0xfd, 0x8d, 0xf0, 0x66, 0x7e, 0xb7, 0xae, 0xcd, 0xe3, 0x07, 0x88, 0x48, 0x6b,
	0xa5, 0x8d, 0xf0, 0x59, 0x89, 0x67, 0xca, 0x9f, 0x2e, 0x95, 0x1d, 0x14, 0xc9, 0x7f, 0x08, 0x7f,
	0xab, 0x6b, 0xaa, 0xdd, 0x38, 0xd6, 0x05, 0x07, 0x3f, 0x5a, 0xe0, 0xd8, 0x5d, 0x5e, 0xc9, 0x92,
	0x3d, 0x19, 0x66, 0x2d, 0xb8, 0xe8, 0x82, 0x6f, 0x10, 0x32, 0xe1, 0xdc, 0x2d, 0x54, 0x49, 0x7c,
	0x7e, 0x98, 0x71, 0x8c, 0x33, 0x88, 0x4b, 0x32, 0x85, 0x96, 0x8d, 0x95, 0xaa, 0x3e, 0x18, 0x7f,
	0x4a, 0x2d, 0xee, 0x3d, 0x88, 0xdd, 0x6c, 0xbf, 0x48, 0xef, 0x64, 0x41, 0xf8, 0x19, 0xa2, 0x1f,
	0x9a, 0x72, 0x4b, 0xf8, 0x7c, 0xfe, 0xe9, 0xe4, 0x8c, 0x3a, 0xfa, 0x9b, 0xf4, 0xf0, 0x23, 0xf8,
	0x2b, 0xb2, 0x17, 0x95, 0x7c, 0x81, 0x68, 0x45, 0x76, 0x59, 0x55, 0x38, 0x7e, 0x22, 0xe1, 0x17,
	0x7c, 0xb9, 0x70, 0x0e, 0xc1, 0x72, 0x6b, 0xd7, 0x5d, 0x97, 0x9d, 0xfb, 0xc4, 0xbe, 0x27, 0x3d,
	0xfc, 0x0a, 0xaf, 0xfe, 0x3a, 0x5b, 0x73, 0x4b, 0xed, 0x53, 0x74, 0xc8, 0xba, 0x4b, 0xbf, 0x47,
	0xff, 0x78, 0x35, 0xae, 0x22, 0xde, 0xf2, 0x4f, 0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x22, 0xc2,
	0xa0, 0xd5, 0xfe, 0x02, 0x00, 0x00,
}