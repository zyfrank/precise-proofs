// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proofs/proto/salt.proto

package proofspb

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

type Salt struct {
	Compact              []byte   `protobuf:"bytes,1,opt,name=compact,proto3" json:"compact,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Salt) Reset()         { *m = Salt{} }
func (m *Salt) String() string { return proto.CompactTextString(m) }
func (*Salt) ProtoMessage()    {}
func (*Salt) Descriptor() ([]byte, []int) {
	return fileDescriptor_276674408a4daf93, []int{0}
}

func (m *Salt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Salt.Unmarshal(m, b)
}
func (m *Salt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Salt.Marshal(b, m, deterministic)
}
func (m *Salt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Salt.Merge(m, src)
}
func (m *Salt) XXX_Size() int {
	return xxx_messageInfo_Salt.Size(m)
}
func (m *Salt) XXX_DiscardUnknown() {
	xxx_messageInfo_Salt.DiscardUnknown(m)
}

var xxx_messageInfo_Salt proto.InternalMessageInfo

func (m *Salt) GetCompact() []byte {
	if m != nil {
		return m.Compact
	}
	return nil
}

func (m *Salt) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Salt)(nil), "proofs.Salt")
}

func init() { proto.RegisterFile("proofs/proto/salt.proto", fileDescriptor_276674408a4daf93) }

var fileDescriptor_276674408a4daf93 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0xcf,
	0x4f, 0x2b, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x4e, 0xcc, 0x29, 0xd1, 0x03, 0x33,
	0x85, 0xd8, 0x20, 0x12, 0x4a, 0x66, 0x5c, 0x2c, 0xc1, 0x89, 0x39, 0x25, 0x42, 0x12, 0x5c, 0xec,
	0xc9, 0xf9, 0xb9, 0x05, 0x89, 0xc9, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x30, 0xae,
	0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x13, 0x58, 0x1c, 0xc2, 0x71, 0x52,
	0xe1, 0xe2, 0x4a, 0xce, 0xcf, 0xd5, 0x83, 0x98, 0xe2, 0xc4, 0x15, 0x00, 0xa2, 0x03, 0x40, 0x26,
	0x07, 0x30, 0x46, 0x71, 0x40, 0x44, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0x96, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xc5, 0xb4, 0x75, 0x72, 0x87, 0x00, 0x00, 0x00,
}
