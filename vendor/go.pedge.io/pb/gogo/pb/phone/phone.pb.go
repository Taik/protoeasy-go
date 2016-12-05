// Code generated by protoc-gen-gogo.
// source: pb/phone/phone.proto
// DO NOT EDIT!

package pbphone

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TelephoneNumber struct {
	CountryCode             uint32 `protobuf:"varint,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	NationalDestinationCode uint64 `protobuf:"varint,2,opt,name=national_destination_code,json=nationalDestinationCode,proto3" json:"national_destination_code,omitempty"`
	SubscriberNumber        uint64 `protobuf:"varint,3,opt,name=subscriber_number,json=subscriberNumber,proto3" json:"subscriber_number,omitempty"`
	Extension               uint64 `protobuf:"varint,4,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *TelephoneNumber) Reset()                    { *m = TelephoneNumber{} }
func (m *TelephoneNumber) String() string            { return proto.CompactTextString(m) }
func (*TelephoneNumber) ProtoMessage()               {}
func (*TelephoneNumber) Descriptor() ([]byte, []int) { return fileDescriptorPhone, []int{0} }

func init() {
	proto.RegisterType((*TelephoneNumber)(nil), "pb.phone.TelephoneNumber")
}

func init() { proto.RegisterFile("pb/phone/phone.proto", fileDescriptorPhone) }

var fileDescriptorPhone = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x48, 0xd2, 0x2f,
	0xc8, 0xc8, 0xcf, 0x4b, 0x85, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x1c, 0x05, 0x49,
	0x7a, 0x60, 0xbe, 0xd2, 0x6e, 0x46, 0x2e, 0xfe, 0x90, 0xd4, 0x9c, 0x54, 0x30, 0xcf, 0xaf, 0x34,
	0x37, 0x29, 0xb5, 0x48, 0x48, 0x91, 0x8b, 0x27, 0x39, 0xbf, 0x34, 0xaf, 0xa4, 0xa8, 0x32, 0x3e,
	0x39, 0x3f, 0x25, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x37, 0x88, 0x1b, 0x2a, 0xe6, 0x9c, 0x9f,
	0x92, 0x2a, 0x64, 0xc5, 0x25, 0x99, 0x97, 0x58, 0x92, 0x99, 0x9f, 0x97, 0x98, 0x13, 0x9f, 0x92,
	0x5a, 0x5c, 0x92, 0x09, 0xe1, 0x41, 0xd4, 0x33, 0x29, 0x30, 0x6a, 0xb0, 0x04, 0x89, 0xc3, 0x14,
	0xb8, 0x20, 0xe4, 0xc1, 0x7a, 0xb5, 0xb9, 0x04, 0x8b, 0x4b, 0x93, 0x8a, 0x93, 0x8b, 0x32, 0x93,
	0x52, 0x8b, 0xe2, 0xf3, 0xc0, 0x76, 0x4a, 0x30, 0x83, 0xf5, 0x08, 0x20, 0x24, 0xa0, 0x6e, 0x91,
	0xe1, 0xe2, 0x4c, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x93, 0x60, 0x01, 0x2b, 0x42,
	0x08, 0x38, 0x71, 0x46, 0xb1, 0x17, 0x24, 0x81, 0x9d, 0x9e, 0xc4, 0x06, 0xf6, 0x99, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xb7, 0xbc, 0x6d, 0x17, 0xf1, 0x00, 0x00, 0x00,
}
