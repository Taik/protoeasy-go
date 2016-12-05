// Code generated by protoc-gen-go.
// source: pb/money/money.proto
// DO NOT EDIT!

package pbmoney

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Currency struct {
	Name        string       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Code        CurrencyCode `protobuf:"varint,2,opt,name=code,enum=pb.money.CurrencyCode" json:"code,omitempty"`
	NumericCode uint32       `protobuf:"varint,3,opt,name=numeric_code,json=numericCode" json:"numeric_code,omitempty"`
	MinorUnit   uint32       `protobuf:"varint,4,opt,name=minor_unit,json=minorUnit" json:"minor_unit,omitempty"`
}

func (m *Currency) Reset()                    { *m = Currency{} }
func (m *Currency) String() string            { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()               {}
func (*Currency) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Currency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Currency) GetCode() CurrencyCode {
	if m != nil {
		return m.Code
	}
	return CurrencyCode_CURRENCY_CODE_NONE
}

func (m *Currency) GetNumericCode() uint32 {
	if m != nil {
		return m.NumericCode
	}
	return 0
}

func (m *Currency) GetMinorUnit() uint32 {
	if m != nil {
		return m.MinorUnit
	}
	return 0
}

type Money struct {
	CurrencyCode CurrencyCode `protobuf:"varint,1,opt,name=currency_code,json=currencyCode,enum=pb.money.CurrencyCode" json:"currency_code,omitempty"`
	ValueMicros  int64        `protobuf:"varint,2,opt,name=value_micros,json=valueMicros" json:"value_micros,omitempty"`
}

func (m *Money) Reset()                    { *m = Money{} }
func (m *Money) String() string            { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()               {}
func (*Money) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Money) GetCurrencyCode() CurrencyCode {
	if m != nil {
		return m.CurrencyCode
	}
	return CurrencyCode_CURRENCY_CODE_NONE
}

func (m *Money) GetValueMicros() int64 {
	if m != nil {
		return m.ValueMicros
	}
	return 0
}

type ExchangeRate struct {
	From        CurrencyCode `protobuf:"varint,1,opt,name=from,enum=pb.money.CurrencyCode" json:"from,omitempty"`
	To          CurrencyCode `protobuf:"varint,2,opt,name=to,enum=pb.money.CurrencyCode" json:"to,omitempty"`
	ValueMicros int64        `protobuf:"varint,3,opt,name=value_micros,json=valueMicros" json:"value_micros,omitempty"`
}

func (m *ExchangeRate) Reset()                    { *m = ExchangeRate{} }
func (m *ExchangeRate) String() string            { return proto.CompactTextString(m) }
func (*ExchangeRate) ProtoMessage()               {}
func (*ExchangeRate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ExchangeRate) GetFrom() CurrencyCode {
	if m != nil {
		return m.From
	}
	return CurrencyCode_CURRENCY_CODE_NONE
}

func (m *ExchangeRate) GetTo() CurrencyCode {
	if m != nil {
		return m.To
	}
	return CurrencyCode_CURRENCY_CODE_NONE
}

func (m *ExchangeRate) GetValueMicros() int64 {
	if m != nil {
		return m.ValueMicros
	}
	return 0
}

func init() {
	proto.RegisterType((*Currency)(nil), "pb.money.Currency")
	proto.RegisterType((*Money)(nil), "pb.money.Money")
	proto.RegisterType((*ExchangeRate)(nil), "pb.money.ExchangeRate")
}

func init() { proto.RegisterFile("pb/money/money.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x51, 0x3f, 0x4b, 0xc4, 0x30,
	0x14, 0x27, 0x6d, 0xd5, 0xeb, 0xbb, 0x9e, 0x43, 0x10, 0x29, 0x82, 0x50, 0x3b, 0x48, 0x71, 0xa8,
	0xa0, 0xa3, 0x9b, 0x87, 0xe3, 0x2d, 0x01, 0x17, 0x97, 0xd2, 0xe6, 0x62, 0x0d, 0x98, 0xbc, 0x12,
	0x53, 0xf1, 0x3e, 0x80, 0x9b, 0x1f, 0x5a, 0xfa, 0x7a, 0x07, 0xe2, 0xc1, 0x75, 0x09, 0x79, 0xbf,
	0xfc, 0xfe, 0x25, 0x81, 0xb3, 0xae, 0xb9, 0x35, 0x68, 0xd5, 0x66, 0x5c, 0xcb, 0xce, 0xa1, 0x47,
	0x3e, 0xeb, 0x9a, 0x92, 0xe6, 0x8b, 0xf4, 0xdf, 0x79, 0xab, 0xec, 0xc8, 0xc9, 0x7f, 0x18, 0xcc,
	0x96, 0xbd, 0x73, 0xca, 0xca, 0x0d, 0xe7, 0x10, 0xd9, 0xda, 0xa8, 0x94, 0x65, 0xac, 0x88, 0x05,
	0xed, 0xf9, 0x0d, 0x44, 0x12, 0xd7, 0x2a, 0x0d, 0x32, 0x56, 0x9c, 0xde, 0x9d, 0x97, 0x3b, 0xcf,
	0x72, 0xa7, 0x5a, 0xe2, 0x5a, 0x09, 0xe2, 0xf0, 0x2b, 0x48, 0x6c, 0x6f, 0x94, 0xd3, 0xb2, 0x22,
	0x4d, 0x98, 0xb1, 0x62, 0x21, 0xe6, 0x5b, 0x6c, 0x20, 0xf2, 0x4b, 0x00, 0xa3, 0x2d, 0xba, 0xaa,
	0xb7, 0xda, 0xa7, 0x11, 0x11, 0x62, 0x42, 0x9e, 0xad, 0xf6, 0x79, 0x0b, 0x47, 0xab, 0xc1, 0x9d,
	0x3f, 0xc0, 0x42, 0x6e, 0x03, 0x46, 0x2f, 0x76, 0x30, 0x3f, 0x91, 0x7f, 0xa6, 0xa1, 0xc7, 0x67,
	0xfd, 0xde, 0xab, 0xca, 0x68, 0xe9, 0xf0, 0x83, 0xba, 0x87, 0x62, 0x4e, 0xd8, 0x8a, 0xa0, 0xfc,
	0x9b, 0x41, 0xf2, 0xf4, 0x25, 0xdf, 0x6a, 0xdb, 0x2a, 0x51, 0x7b, 0xba, 0xe7, 0xab, 0x43, 0x33,
	0x91, 0x43, 0x1c, 0x7e, 0x0d, 0x81, 0xc7, 0x89, 0x17, 0x09, 0x3c, 0xee, 0xf5, 0x08, 0xf7, 0x7a,
	0x3c, 0xc6, 0x2f, 0x27, 0x5d, 0x43, 0xf2, 0xe6, 0x98, 0x7e, 0xe4, 0xfe, 0x37, 0x00, 0x00, 0xff,
	0xff, 0xd1, 0xdc, 0x6f, 0x56, 0xcd, 0x01, 0x00, 0x00,
}
