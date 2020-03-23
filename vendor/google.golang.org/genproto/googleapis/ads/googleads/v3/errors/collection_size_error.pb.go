// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/collection_size_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible collection size errors.
type CollectionSizeErrorEnum_CollectionSizeError int32

const (
	// Enum unspecified.
	CollectionSizeErrorEnum_UNSPECIFIED CollectionSizeErrorEnum_CollectionSizeError = 0
	// The received error code is not known in this version.
	CollectionSizeErrorEnum_UNKNOWN CollectionSizeErrorEnum_CollectionSizeError = 1
	// Too few.
	CollectionSizeErrorEnum_TOO_FEW CollectionSizeErrorEnum_CollectionSizeError = 2
	// Too many.
	CollectionSizeErrorEnum_TOO_MANY CollectionSizeErrorEnum_CollectionSizeError = 3
)

var CollectionSizeErrorEnum_CollectionSizeError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "TOO_FEW",
	3: "TOO_MANY",
}

var CollectionSizeErrorEnum_CollectionSizeError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"TOO_FEW":     2,
	"TOO_MANY":    3,
}

func (x CollectionSizeErrorEnum_CollectionSizeError) String() string {
	return proto.EnumName(CollectionSizeErrorEnum_CollectionSizeError_name, int32(x))
}

func (CollectionSizeErrorEnum_CollectionSizeError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_04810df5586fcd04, []int{0, 0}
}

// Container for enum describing possible collection size errors.
type CollectionSizeErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionSizeErrorEnum) Reset()         { *m = CollectionSizeErrorEnum{} }
func (m *CollectionSizeErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CollectionSizeErrorEnum) ProtoMessage()    {}
func (*CollectionSizeErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_04810df5586fcd04, []int{0}
}

func (m *CollectionSizeErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionSizeErrorEnum.Unmarshal(m, b)
}
func (m *CollectionSizeErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionSizeErrorEnum.Marshal(b, m, deterministic)
}
func (m *CollectionSizeErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionSizeErrorEnum.Merge(m, src)
}
func (m *CollectionSizeErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CollectionSizeErrorEnum.Size(m)
}
func (m *CollectionSizeErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionSizeErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionSizeErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.CollectionSizeErrorEnum_CollectionSizeError", CollectionSizeErrorEnum_CollectionSizeError_name, CollectionSizeErrorEnum_CollectionSizeError_value)
	proto.RegisterType((*CollectionSizeErrorEnum)(nil), "google.ads.googleads.v3.errors.CollectionSizeErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/collection_size_error.proto", fileDescriptor_04810df5586fcd04)
}

var fileDescriptor_04810df5586fcd04 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x5d, 0x07, 0x2a, 0x99, 0x60, 0xa9, 0x07, 0x45, 0x64, 0x87, 0x3e, 0x40, 0x72, 0xc8,
	0x2d, 0x9e, 0xb2, 0xad, 0x1b, 0x43, 0xcc, 0x06, 0x73, 0x1b, 0x4a, 0xa1, 0xd4, 0x36, 0x84, 0x40,
	0x97, 0x94, 0xa6, 0xee, 0xb0, 0xc7, 0xf1, 0xe8, 0xa3, 0xf8, 0x28, 0x5e, 0x7d, 0x01, 0x49, 0x63,
	0x7b, 0x9a, 0x9e, 0xfa, 0xff, 0xbe, 0xfe, 0x7f, 0xff, 0xef, 0xcb, 0x07, 0x88, 0xd0, 0x5a, 0x14,
	0x1c, 0xa5, 0xb9, 0x41, 0x4e, 0x5a, 0xb5, 0xc7, 0x88, 0x57, 0x95, 0xae, 0x0c, 0xca, 0x74, 0x51,
	0xf0, 0xac, 0x96, 0x5a, 0x25, 0x46, 0x1e, 0x78, 0xd2, 0xb4, 0x61, 0x59, 0xe9, 0x5a, 0x07, 0x43,
	0x07, 0xc0, 0x34, 0x37, 0xb0, 0x63, 0xe1, 0x1e, 0x43, 0xc7, 0xde, 0xde, 0xb5, 0xd9, 0xa5, 0x44,
	0xa9, 0x52, 0xba, 0x4e, 0x6d, 0x8e, 0x71, 0x74, 0x28, 0xc1, 0xf5, 0xb8, 0x0b, 0x5f, 0xc9, 0x03,
	0x8f, 0x2c, 0x15, 0xa9, 0xb7, 0x5d, 0xc8, 0xc0, 0xd5, 0x91, 0x5f, 0xc1, 0x25, 0x18, 0xac, 0xd9,
	0x6a, 0x19, 0x8d, 0xe7, 0xd3, 0x79, 0x34, 0xf1, 0x4f, 0x82, 0x01, 0x38, 0x5b, 0xb3, 0x07, 0xb6,
	0xd8, 0x32, 0xbf, 0x67, 0x8b, 0xa7, 0xc5, 0x22, 0x99, 0x46, 0x5b, 0xdf, 0x0b, 0x2e, 0xc0, 0xb9,
	0x2d, 0x1e, 0x29, 0x7b, 0xf6, 0xfb, 0xa3, 0xef, 0x1e, 0x08, 0x33, 0xbd, 0x83, 0xff, 0xef, 0x3b,
	0xba, 0x39, 0x32, 0x74, 0x69, 0x77, 0x5d, 0xf6, 0x5e, 0x26, 0xbf, 0xac, 0xd0, 0x45, 0xaa, 0x04,
	0xd4, 0x95, 0x40, 0x82, 0xab, 0xe6, 0x25, 0xed, 0xdd, 0x4a, 0x69, 0xfe, 0x3a, 0xe3, 0xbd, 0xfb,
	0xbc, 0x7b, 0xfd, 0x19, 0xa5, 0x1f, 0xde, 0x70, 0xe6, 0xc2, 0x68, 0x6e, 0xa0, 0x93, 0x56, 0x6d,
	0x30, 0x6c, 0x46, 0x9a, 0xcf, 0xd6, 0x10, 0xd3, 0xdc, 0xc4, 0x9d, 0x21, 0xde, 0xe0, 0xd8, 0x19,
	0xbe, 0xbc, 0xd0, 0x75, 0x09, 0xa1, 0xb9, 0x21, 0xa4, 0xb3, 0x10, 0xb2, 0xc1, 0x84, 0x38, 0xd3,
	0xeb, 0x69, 0xb3, 0x1d, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xed, 0xe7, 0xdb, 0x35, 0xe3, 0x01,
	0x00, 0x00,
}