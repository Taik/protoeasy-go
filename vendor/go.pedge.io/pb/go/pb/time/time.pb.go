// Code generated by protoc-gen-go.
// source: pb/time/time.proto
// DO NOT EDIT!

package pbtime

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TimestampRange struct {
	Start *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *TimestampRange) Reset()                    { *m = TimestampRange{} }
func (m *TimestampRange) String() string            { return proto.CompactTextString(m) }
func (*TimestampRange) ProtoMessage()               {}
func (*TimestampRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TimestampRange) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *TimestampRange) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*TimestampRange)(nil), "pb.time.TimestampRange")
}

func init() { proto.RegisterFile("pb/time/time.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x48, 0xd2, 0x2f,
	0xc9, 0xcc, 0x4d, 0x05, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x05, 0x49, 0x7a,
	0x20, 0xae, 0x94, 0x7c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x38, 0xa9, 0x34, 0x0d,
	0xac, 0xa8, 0xb8, 0x24, 0x31, 0xb7, 0x00, 0xa2, 0x52, 0xa9, 0x80, 0x8b, 0x2f, 0x04, 0x26, 0x14,
	0x94, 0x98, 0x97, 0x9e, 0x2a, 0x64, 0xc0, 0xc5, 0x5a, 0x5c, 0x92, 0x58, 0x54, 0x22, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0x31, 0x42, 0x0f, 0x66, 0x84, 0x1e, 0x42, 0x3d, 0x44,
	0xa1, 0x90, 0x0e, 0x17, 0x73, 0x6a, 0x5e, 0x8a, 0x04, 0x13, 0x41, 0xf5, 0x20, 0x65, 0x4e, 0x1c,
	0x51, 0x6c, 0x05, 0x49, 0x20, 0x67, 0x24, 0xb1, 0x81, 0x95, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0x4c, 0x7a, 0x99, 0xc2, 0x00, 0x00, 0x00,
}