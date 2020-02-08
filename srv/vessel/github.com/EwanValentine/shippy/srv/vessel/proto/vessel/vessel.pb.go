// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessel

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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
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

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716) }

var fileDescriptor_04ef66862bb50716 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x69, 0x9b, 0xa6, 0x23, 0x2d, 0xb2, 0x20, 0x6c, 0x8b, 0x42, 0xc8, 0x41, 0x72, 0x90,
	0x06, 0xea, 0x41, 0xf4, 0xa6, 0xa2, 0xa0, 0xc7, 0x14, 0x2a, 0x78, 0x29, 0xdb, 0x64, 0x6c, 0x07,
	0x92, 0x4d, 0x48, 0xd6, 0xb4, 0xfd, 0x1b, 0x3f, 0x55, 0xd8, 0x24, 0xd5, 0x96, 0xe2, 0x29, 0xf3,
	0xde, 0xbc, 0x3c, 0xde, 0x3e, 0x06, 0x86, 0x59, 0x9e, 0xaa, 0xd4, 0x2f, 0xb1, 0x28, 0x30, 0xae,
	0x3f, 0x63, 0xcd, 0x31, 0xab, 0x42, 0xee, 0xb7, 0x01, 0xd6, 0x4c, 0x8f, 0x6c, 0x00, 0x26, 0x45,
	0xdc, 0x70, 0x0c, 0xaf, 0x17, 0x98, 0x14, 0xb1, 0x11, 0xd8, 0xa1, 0xc8, 0x44, 0x48, 0x6a, 0xcb,
	0x4d, 0xc7, 0xf0, 0x3a, 0xc1, 0x0e, 0xb3, 0x4b, 0x80, 0x44, 0x6c, 0xe6, 0x6b, 0xa4, 0xe5, 0x4a,
	0xf1, 0x96, 0xde, 0xf6, 0x12, 0xb1, 0x79, 0xd7, 0x04, 0x63, 0xd0, 0x96, 0x22, 0x41, 0xde, 0xd6,
	0x66, 0x7a, 0x66, 0x17, 0xd0, 0x13, 0xa5, 0xa0, 0x58, 0x2c, 0x62, 0xe4, 0x1d, 0xc7, 0xf0, 0xec,
	0xe0, 0x97, 0x60, 0x43, 0xb0, 0xd3, 0xb5, 0xc4, 0x7c, 0x4e, 0x11, 0xb7, 0xf4, 0x5f, 0x5d, 0x8d,
	0x5f, 0x23, 0xf7, 0x0d, 0xfa, 0xd3, 0x0c, 0x43, 0xfa, 0xa4, 0x50, 0x28, 0x4a, 0xe5, 0x5e, 0x30,
	0xe3, 0xdf, 0x60, 0xe6, 0x41, 0x30, 0xb7, 0x04, 0x3b, 0xc0, 0x22, 0x4b, 0x65, 0x81, 0xec, 0x0a,
	0xea, 0x12, 0xb4, 0xc9, 0xe9, 0x64, 0x30, 0xae, 0x1b, 0xaa, 0xfa, 0x08, 0xea, 0x2d, 0xf3, 0xa0,
	0x5b, 0x4d, 0x05, 0x37, 0x9d, 0xd6, 0x11, 0x61, 0xb3, 0x66, 0x1c, 0xba, 0x61, 0x8e, 0x42, 0x61,
	0xa4, 0x2b, 0xb1, 0x83, 0x06, 0x4e, 0xb6, 0xd0, 0xaf, 0xc4, 0x53, 0xcc, 0x4b, 0x0a, 0x91, 0xdd,
	0x43, 0xff, 0x85, 0x64, 0xf4, 0xb0, 0x2b, 0xe0, 0xbc, 0x31, 0xdd, 0x7b, 0xeb, 0xe8, 0xac, 0xa1,
	0x9b, 0xd8, 0xee, 0x09, 0xbb, 0x06, 0xeb, 0x49, 0xfb, 0xb2, 0x83, 0x24, 0xc7, 0xd4, 0x8f, 0x77,
	0x1f, 0xb7, 0x4b, 0x52, 0xab, 0xaf, 0xc5, 0x38, 0x4c, 0x13, 0xff, 0x79, 0x2d, 0xe4, 0x4c, 0xc4,
	0x28, 0x15, 0x49, 0xf4, 0x8b, 0x15, 0x65, 0xd9, 0xd6, 0x2f, 0xf2, 0xb2, 0x39, 0x92, 0xbf, 0x17,
	0xb3, 0xb0, 0x34, 0xba, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x98, 0xf5, 0xb6, 0xd3, 0x48, 0x02,
	0x00, 0x00,
}