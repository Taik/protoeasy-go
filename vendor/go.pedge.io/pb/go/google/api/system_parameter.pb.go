// Code generated by protoc-gen-go.
// source: google/api/system_parameter.proto
// DO NOT EDIT!

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ### System parameter configuration
//
// A system parameter is a special kind of parameter defined by the API
// system, not by an individual API. It is typically mapped to an HTTP header
// and/or a URL query parameter. This configuration specifies which methods
// change the names of the system parameters.
type SystemParameters struct {
	// Define system parameters.
	//
	// The parameters defined here will override the default parameters
	// implemented by the system. If this field is missing from the service
	// config, default system parameters will be used. Default system parameters
	// and names is implementation-dependent.
	//
	// Example: define api key and alt name for all methods
	//
	// system_parameters
	//   rules:
	//     - selector: "*"
	//       parameters:
	//         - name: api_key
	//           url_query_parameter: api_key
	//         - name: alt
	//           http_header: Response-Content-Type
	//
	// Example: define 2 api key names for a specific method.
	//
	// system_parameters
	//   rules:
	//     - selector: "/ListShelves"
	//       parameters:
	//         - name: api_key
	//           http_header: Api-Key1
	//         - name: api_key
	//           http_header: Api-Key2
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*SystemParameterRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *SystemParameters) Reset()                    { *m = SystemParameters{} }
func (m *SystemParameters) String() string            { return proto.CompactTextString(m) }
func (*SystemParameters) ProtoMessage()               {}
func (*SystemParameters) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

func (m *SystemParameters) GetRules() []*SystemParameterRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Define a system parameter rule mapping system parameter definitions to
// methods.
type SystemParameterRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// Define parameters. Multiple names may be defined for a parameter.
	// For a given method call, only one of them should be used. If multiple
	// names are used the behavior is implementation-dependent.
	// If none of the specified names are present the behavior is
	// parameter-dependent.
	Parameters []*SystemParameter `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty"`
}

func (m *SystemParameterRule) Reset()                    { *m = SystemParameterRule{} }
func (m *SystemParameterRule) String() string            { return proto.CompactTextString(m) }
func (*SystemParameterRule) ProtoMessage()               {}
func (*SystemParameterRule) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{1} }

func (m *SystemParameterRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *SystemParameterRule) GetParameters() []*SystemParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// Define a parameter's name and location. The parameter may be passed as either
// an HTTP header or a URL query parameter, and if both are passed the behavior
// is implementation-dependent.
type SystemParameter struct {
	// Define the name of the parameter, such as "api_key", "alt", "callback",
	// and etc. It is case sensitive.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Define the HTTP header name to use for the parameter. It is case
	// insensitive.
	HttpHeader string `protobuf:"bytes,2,opt,name=http_header,json=httpHeader" json:"http_header,omitempty"`
	// Define the URL query parameter name to use for the parameter. It is case
	// sensitive.
	UrlQueryParameter string `protobuf:"bytes,3,opt,name=url_query_parameter,json=urlQueryParameter" json:"url_query_parameter,omitempty"`
}

func (m *SystemParameter) Reset()                    { *m = SystemParameter{} }
func (m *SystemParameter) String() string            { return proto.CompactTextString(m) }
func (*SystemParameter) ProtoMessage()               {}
func (*SystemParameter) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{2} }

func (m *SystemParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SystemParameter) GetHttpHeader() string {
	if m != nil {
		return m.HttpHeader
	}
	return ""
}

func (m *SystemParameter) GetUrlQueryParameter() string {
	if m != nil {
		return m.UrlQueryParameter
	}
	return ""
}

func init() {
	proto.RegisterType((*SystemParameters)(nil), "google.api.SystemParameters")
	proto.RegisterType((*SystemParameterRule)(nil), "google.api.SystemParameterRule")
	proto.RegisterType((*SystemParameter)(nil), "google.api.SystemParameter")
}

func init() { proto.RegisterFile("google/api/system_parameter.proto", fileDescriptor20) }

var fileDescriptor20 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x99, 0x69, 0x15, 0xbd, 0x05, 0x7f, 0x6e, 0x5d, 0x04, 0x5d, 0xb4, 0xce, 0xaa, 0xab,
	0x0c, 0x28, 0xae, 0x5c, 0xd9, 0x8d, 0x76, 0x37, 0x8e, 0x0f, 0x30, 0xc4, 0x7a, 0x69, 0x0b, 0x99,
	0x26, 0xde, 0x24, 0x42, 0x5f, 0xc7, 0x27, 0x95, 0x49, 0x75, 0x5a, 0x06, 0xe9, 0x2e, 0xb9, 0xe7,
	0x83, 0xef, 0x70, 0xe0, 0x76, 0x61, 0xcc, 0x42, 0x53, 0xae, 0xec, 0x2a, 0x77, 0x1b, 0xe7, 0xa9,
	0xae, 0xac, 0x62, 0x55, 0x93, 0x27, 0x96, 0x96, 0x8d, 0x37, 0x08, 0x5b, 0x44, 0x2a, 0xbb, 0xca,
	0x66, 0x70, 0xf1, 0x16, 0xa9, 0xe2, 0x0f, 0x72, 0xf8, 0x00, 0x47, 0x1c, 0x34, 0x39, 0x91, 0x8c,
	0x7b, 0x93, 0xc1, 0xdd, 0x48, 0xee, 0x78, 0xd9, 0x81, 0xcb, 0xa0, 0xa9, 0xdc, 0xd2, 0xd9, 0x1a,
	0x86, 0xff, 0xa4, 0x78, 0x0d, 0x27, 0x8e, 0x34, 0xcd, 0xbd, 0x61, 0x91, 0x8c, 0x93, 0xc9, 0x69,
	0xd9, 0xfe, 0xf1, 0x11, 0xa0, 0x2d, 0xe7, 0x44, 0x1a, 0x75, 0x37, 0x87, 0x74, 0x7b, 0x78, 0xf6,
	0x05, 0xe7, 0x9d, 0x18, 0x11, 0xfa, 0x6b, 0x55, 0xd3, 0xaf, 0x27, 0xbe, 0x71, 0x04, 0x83, 0xa5,
	0xf7, 0xb6, 0x5a, 0x92, 0xfa, 0x20, 0x16, 0x69, 0x8c, 0xa0, 0x39, 0xbd, 0xc4, 0x0b, 0x4a, 0x18,
	0x06, 0xd6, 0xd5, 0x67, 0x20, 0xde, 0xec, 0xb6, 0x12, 0xbd, 0x08, 0x5e, 0x06, 0xd6, 0xaf, 0x4d,
	0xd2, 0x4a, 0xa6, 0x39, 0x9c, 0xcd, 0x4d, 0xbd, 0xd7, 0x72, 0x7a, 0xd5, 0xe9, 0x51, 0x34, 0x33,
	0x17, 0xc9, 0x77, 0xda, 0x7f, 0x7e, 0x2a, 0x66, 0xef, 0xc7, 0x71, 0xf6, 0xfb, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc9, 0x8a, 0x6c, 0x74, 0x9b, 0x01, 0x00, 0x00,
}
