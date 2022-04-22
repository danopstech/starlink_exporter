// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spacex/api/device/command.proto

package device

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

type Capability int32

const (
	Capability_READ             Capability = 0
	Capability_READ_INTERNAL    Capability = 13
	Capability_READ_PRIVATE     Capability = 7
	Capability_WRITE            Capability = 1
	Capability_WRITE_PERSISTENT Capability = 11
	Capability_DEBUG            Capability = 2
	Capability_ADMIN            Capability = 3
	Capability_SETUP            Capability = 4
	Capability_SET_SKU          Capability = 5
	Capability_REFRESH          Capability = 6
	Capability_FUSE             Capability = 8
	Capability_RESET            Capability = 9
	Capability_TEST             Capability = 10
	Capability_SSH              Capability = 12
)

var Capability_name = map[int32]string{
	0:  "READ",
	13: "READ_INTERNAL",
	7:  "READ_PRIVATE",
	1:  "WRITE",
	11: "WRITE_PERSISTENT",
	2:  "DEBUG",
	3:  "ADMIN",
	4:  "SETUP",
	5:  "SET_SKU",
	6:  "REFRESH",
	8:  "FUSE",
	9:  "RESET",
	10: "TEST",
	12: "SSH",
}

var Capability_value = map[string]int32{
	"READ":             0,
	"READ_INTERNAL":    13,
	"READ_PRIVATE":     7,
	"WRITE":            1,
	"WRITE_PERSISTENT": 11,
	"DEBUG":            2,
	"ADMIN":            3,
	"SETUP":            4,
	"SET_SKU":          5,
	"REFRESH":          6,
	"FUSE":             8,
	"RESET":            9,
	"TEST":             10,
	"SSH":              12,
}

func (x Capability) String() string {
	return proto.EnumName(Capability_name, int32(x))
}

func (Capability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c24c48853ef12d9, []int{0}
}

type PublicKey struct {
	Key                  string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Capabilities         []Capability `protobuf:"varint,2,rep,packed,name=capabilities,proto3,enum=SpaceX.API.Device.Capability" json:"capabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c24c48853ef12d9, []int{0}
}

func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (m *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(m, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PublicKey) GetCapabilities() []Capability {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

func init() {
	proto.RegisterEnum("SpaceX.API.Device.Capability", Capability_name, Capability_value)
	proto.RegisterType((*PublicKey)(nil), "SpaceX.API.Device.PublicKey")
}

func init() { proto.RegisterFile("spacex/api/device/command.proto", fileDescriptor_3c24c48853ef12d9) }

var fileDescriptor_3c24c48853ef12d9 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x2d, 0xe5, 0x5f, 0x07, 0x30, 0xc3, 0x46, 0x23, 0x17, 0x23, 0xf1, 0x44, 0x3c, 0x94,
	0x44, 0x9f, 0xa0, 0xd8, 0x41, 0x1a, 0xb4, 0x69, 0x76, 0xb6, 0x6a, 0xbc, 0x60, 0x29, 0x3d, 0x6c,
	0x04, 0xdb, 0x08, 0x1a, 0xfb, 0x74, 0xbe, 0x9a, 0xd9, 0x1a, 0xa3, 0xc6, 0xdb, 0x6f, 0xbf, 0xfd,
	0xcd, 0x64, 0xf2, 0xc1, 0xc9, 0xb6, 0x48, 0xd2, 0xec, 0x7d, 0x9c, 0x14, 0x7a, 0xbc, 0xca, 0xde,
	0x74, 0x9a, 0x8d, 0xd3, 0x7c, 0xb3, 0x49, 0x9e, 0x57, 0x6e, 0xf1, 0x92, 0xef, 0x72, 0xd1, 0x67,
	0x23, 0xdc, 0xbb, 0x5e, 0x14, 0xb8, 0x7e, 0x25, 0x9c, 0x3e, 0x82, 0x13, 0xbd, 0x2e, 0xd7, 0x3a,
	0x9d, 0x67, 0xa5, 0x40, 0xb0, 0x9f, 0xb2, 0x72, 0x60, 0x0d, 0xad, 0x91, 0x23, 0x0d, 0x0a, 0x0f,
	0xba, 0x69, 0x52, 0x24, 0x4b, 0xbd, 0xd6, 0x3b, 0x9d, 0x6d, 0x07, 0xb5, 0xa1, 0x3d, 0xda, 0x3f,
	0x3f, 0x76, 0xff, 0x2d, 0x72, 0x2f, 0xbf, 0xb5, 0x52, 0xfe, 0x19, 0x39, 0xfb, 0xb0, 0x00, 0x7e,
	0x3e, 0x45, 0x1b, 0xea, 0x92, 0x3c, 0x1f, 0xf7, 0x44, 0x1f, 0x7a, 0x86, 0x16, 0x41, 0xa8, 0x48,
	0x86, 0xde, 0x35, 0xf6, 0x04, 0x42, 0xb7, 0x8a, 0x22, 0x19, 0xdc, 0x7a, 0x8a, 0xb0, 0x25, 0x1c,
	0x68, 0xdc, 0xc9, 0x40, 0x11, 0x5a, 0xe2, 0x00, 0xb0, 0xc2, 0x45, 0x44, 0x92, 0x03, 0x56, 0x14,
	0x2a, 0xec, 0x18, 0xc1, 0xa7, 0x49, 0x7c, 0x85, 0x35, 0x83, 0x9e, 0x7f, 0x13, 0x84, 0x68, 0x1b,
	0x64, 0x52, 0x71, 0x84, 0x75, 0xd1, 0x81, 0x16, 0x93, 0x5a, 0xf0, 0x3c, 0xc6, 0x86, 0x79, 0x48,
	0x9a, 0x4a, 0xe2, 0x19, 0x36, 0xcd, 0x29, 0xd3, 0x98, 0x09, 0xdb, 0x46, 0x97, 0xc4, 0xa4, 0xd0,
	0x31, 0xa1, 0x22, 0x56, 0x08, 0xa2, 0x05, 0x36, 0xf3, 0x0c, 0xbb, 0x93, 0xa3, 0x87, 0xc3, 0xaf,
	0x66, 0xdd, 0x34, 0xdf, 0xfc, 0x6a, 0x77, 0xd9, 0xac, 0x6a, 0xbd, 0xf8, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x24, 0x6f, 0x61, 0x4b, 0x79, 0x01, 0x00, 0x00,
}
