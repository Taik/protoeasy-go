// Code generated by protoc-gen-go.
// source: google/api/label.proto
// DO NOT EDIT!

package google_api

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

// Value types that can be used as label values.
type LabelDescriptor_ValueType int32

const (
	// A variable-length string. This is the default.
	LabelDescriptor_STRING LabelDescriptor_ValueType = 0
	// Boolean; true or false.
	LabelDescriptor_BOOL LabelDescriptor_ValueType = 1
	// A 64-bit signed integer.
	LabelDescriptor_INT64 LabelDescriptor_ValueType = 2
)

var LabelDescriptor_ValueType_name = map[int32]string{
	0: "STRING",
	1: "BOOL",
	2: "INT64",
}
var LabelDescriptor_ValueType_value = map[string]int32{
	"STRING": 0,
	"BOOL":   1,
	"INT64":  2,
}

func (x LabelDescriptor_ValueType) String() string {
	return proto.EnumName(LabelDescriptor_ValueType_name, int32(x))
}
func (LabelDescriptor_ValueType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// A description of a label.
type LabelDescriptor struct {
	// The label key.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The type of data that can be assigned to the label.
	ValueType LabelDescriptor_ValueType `protobuf:"varint,2,opt,name=value_type,json=valueType,enum=google.api.LabelDescriptor_ValueType" json:"value_type,omitempty"`
	// A human-readable description for the label.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *LabelDescriptor) Reset()                    { *m = LabelDescriptor{} }
func (m *LabelDescriptor) String() string            { return proto.CompactTextString(m) }
func (*LabelDescriptor) ProtoMessage()               {}
func (*LabelDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*LabelDescriptor)(nil), "google.api.LabelDescriptor")
	proto.RegisterEnum("google.api.LabelDescriptor_ValueType", LabelDescriptor_ValueType_name, LabelDescriptor_ValueType_value)
}

func init() { proto.RegisterFile("google/api/label.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x49, 0x4c, 0x4a, 0xcd, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x88, 0xeb, 0x01, 0xc5, 0x95, 0x76, 0x32, 0x72, 0xf1, 0xfb, 0x80,
	0xe4, 0x5c, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0xf2, 0x8b, 0x84, 0x04, 0xb8, 0x98, 0xb3,
	0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x17, 0x2e, 0xae, 0xb2,
	0xc4, 0x9c, 0xd2, 0xd4, 0xf8, 0x92, 0xca, 0x82, 0x54, 0x09, 0x26, 0xa0, 0x04, 0x9f, 0x91, 0xaa,
	0x1e, 0xc2, 0x18, 0x3d, 0x34, 0x23, 0xf4, 0xc2, 0x40, 0xaa, 0x43, 0x80, 0x8a, 0x83, 0x38, 0xcb,
	0x60, 0x4c, 0x21, 0x05, 0x2e, 0xee, 0x14, 0xa8, 0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x66, 0xb0, 0xf9,
	0xc8, 0x42, 0x4a, 0x3a, 0x5c, 0x9c, 0x70, 0x9d, 0x42, 0x5c, 0x5c, 0x6c, 0xc1, 0x21, 0x41, 0x9e,
	0x7e, 0xee, 0x02, 0x0c, 0x42, 0x1c, 0x5c, 0x2c, 0x4e, 0xfe, 0xfe, 0x3e, 0x02, 0x8c, 0x42, 0x9c,
	0x5c, 0xac, 0x9e, 0x7e, 0x21, 0x66, 0x26, 0x02, 0x4c, 0x4e, 0x72, 0x5c, 0x7c, 0xc9, 0xf9, 0xb9,
	0x48, 0xce, 0x70, 0xe2, 0x02, 0xbb, 0x23, 0x00, 0xe4, 0xcb, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0x77,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x4a, 0xb5, 0x2c, 0x08, 0x01, 0x00, 0x00,
}
