// Code generated by protoc-gen-go.
// source: google/rpc/error_details.proto
// DO NOT EDIT!

package google_rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Describes when the clients can retry a failed request. Clients could ignore
// the recommendation here or retry when this information is missing from error
// responses.
//
// It's always recommended that clients should use exponential backoff when
// retrying.
//
// Clients should wait until `retry_delay` amount of time has passed since
// receiving the error response before retrying.  If retrying requests also
// fail, clients should use an exponential backoff scheme to gradually increase
// the delay between retries based on `retry_delay`, until either a maximum
// number of retires have been reached or a maximum retry delay cap has been
// reached.
type RetryInfo struct {
	// Clients should wait at least this long between retrying the same request.
	RetryDelay *google_protobuf.Duration `protobuf:"bytes,1,opt,name=retry_delay,json=retryDelay" json:"retry_delay,omitempty"`
}

func (m *RetryInfo) Reset()                    { *m = RetryInfo{} }
func (m *RetryInfo) String() string            { return proto.CompactTextString(m) }
func (*RetryInfo) ProtoMessage()               {}
func (*RetryInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *RetryInfo) GetRetryDelay() *google_protobuf.Duration {
	if m != nil {
		return m.RetryDelay
	}
	return nil
}

// Describes additional debugging info.
type DebugInfo struct {
	// The stack trace entries indicating where the error occurred.
	StackEntries []string `protobuf:"bytes,1,rep,name=stack_entries,json=stackEntries" json:"stack_entries,omitempty"`
	// Additional debugging information provided by the server.
	Detail string `protobuf:"bytes,2,opt,name=detail" json:"detail,omitempty"`
}

func (m *DebugInfo) Reset()                    { *m = DebugInfo{} }
func (m *DebugInfo) String() string            { return proto.CompactTextString(m) }
func (*DebugInfo) ProtoMessage()               {}
func (*DebugInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *DebugInfo) GetStackEntries() []string {
	if m != nil {
		return m.StackEntries
	}
	return nil
}

func (m *DebugInfo) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

// Describes how a quota check failed.
//
// For example if a daily limit was exceeded for the calling project,
// a service could respond with a QuotaFailure detail containing the project
// id and the description of the quota limit that was exceeded.  If the
// calling project hasn't enabled the service in the developer console, then
// a service could respond with the project id and set `service_disabled`
// to true.
//
// Also see RetryDetail and Help types for other details about handling a
// quota failure.
type QuotaFailure struct {
	// Describes all quota violations.
	Violations []*QuotaFailure_Violation `protobuf:"bytes,1,rep,name=violations" json:"violations,omitempty"`
}

func (m *QuotaFailure) Reset()                    { *m = QuotaFailure{} }
func (m *QuotaFailure) String() string            { return proto.CompactTextString(m) }
func (*QuotaFailure) ProtoMessage()               {}
func (*QuotaFailure) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *QuotaFailure) GetViolations() []*QuotaFailure_Violation {
	if m != nil {
		return m.Violations
	}
	return nil
}

// A message type used to describe a single quota violation.  For example, a
// daily quota or a custom quota that was exceeded.
type QuotaFailure_Violation struct {
	// The subject on which the quota check failed.
	// For example, "clientip:<ip address of client>" or "project:<Google
	// developer project id>".
	Subject string `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	// A description of how the quota check failed. Clients can use this
	// description to find more about the quota configuration in the service's
	// public documentation, or find the relevant quota limit to adjust through
	// developer console.
	//
	// For example: "Service disabled" or "Daily Limit for read operations
	// exceeded".
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *QuotaFailure_Violation) Reset()                    { *m = QuotaFailure_Violation{} }
func (m *QuotaFailure_Violation) String() string            { return proto.CompactTextString(m) }
func (*QuotaFailure_Violation) ProtoMessage()               {}
func (*QuotaFailure_Violation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

func (m *QuotaFailure_Violation) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *QuotaFailure_Violation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Describes violations in a client request. This error type focuses on the
// syntactic aspects of the request.
type BadRequest struct {
	// Describes all violations in a client request.
	FieldViolations []*BadRequest_FieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations" json:"field_violations,omitempty"`
}

func (m *BadRequest) Reset()                    { *m = BadRequest{} }
func (m *BadRequest) String() string            { return proto.CompactTextString(m) }
func (*BadRequest) ProtoMessage()               {}
func (*BadRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BadRequest) GetFieldViolations() []*BadRequest_FieldViolation {
	if m != nil {
		return m.FieldViolations
	}
	return nil
}

// A message type used to describe a single bad request field.
type BadRequest_FieldViolation struct {
	// A path leading to a field in the request body. The value will be a
	// sequence of dot-separated identifiers that identify a protocol buffer
	// field. E.g., "field_violations.field" would identify this field.
	Field string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	// A description of why the request element is bad.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *BadRequest_FieldViolation) Reset()                    { *m = BadRequest_FieldViolation{} }
func (m *BadRequest_FieldViolation) String() string            { return proto.CompactTextString(m) }
func (*BadRequest_FieldViolation) ProtoMessage()               {}
func (*BadRequest_FieldViolation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

func (m *BadRequest_FieldViolation) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *BadRequest_FieldViolation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Contains metadata about the request that clients can attach when filing a bug
// or providing other forms of feedback.
type RequestInfo struct {
	// An opaque string that should only be interpreted by the service generating
	// it. For example, it can be used to identify requests in the service's logs.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	// Any data that was used to serve this request. For example, an encrypted
	// stack trace that can be sent back to the service provider for debugging.
	ServingData string `protobuf:"bytes,2,opt,name=serving_data,json=servingData" json:"serving_data,omitempty"`
}

func (m *RequestInfo) Reset()                    { *m = RequestInfo{} }
func (m *RequestInfo) String() string            { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()               {}
func (*RequestInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *RequestInfo) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *RequestInfo) GetServingData() string {
	if m != nil {
		return m.ServingData
	}
	return ""
}

// Describes the resource that is being accessed.
type ResourceInfo struct {
	// A name for the type of resource being accessed, e.g. "sql table",
	// "cloud storage bucket", "file", "Google calendar"; or the type URL
	// of the resource: e.g. "type.googleapis.com/google.pubsub.v1.Topic".
	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType" json:"resource_type,omitempty"`
	// The name of the resource being accessed.  For example, a shared calendar
	// name: "example.com_4fghdhgsrgh@group.calendar.google.com", if the current
	// error is [google.rpc.Code.PERMISSION_DENIED][google.rpc.Code.PERMISSION_DENIED].
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName" json:"resource_name,omitempty"`
	// The owner of the resource (optional).
	// For example, "user:<owner email>" or "project:<Google developer project
	// id>".
	Owner string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	// Describes what error is encountered when accessing this resource.
	// For example, updating a cloud project may require the `writer` permission
	// on the developer console project.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
}

func (m *ResourceInfo) Reset()                    { *m = ResourceInfo{} }
func (m *ResourceInfo) String() string            { return proto.CompactTextString(m) }
func (*ResourceInfo) ProtoMessage()               {}
func (*ResourceInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ResourceInfo) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ResourceInfo) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ResourceInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ResourceInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Provides links to documentation or for performing an out of band action.
//
// For example, if a quota check failed with an error indicating the calling
// project hasn't enabled the accessed service, this can contain a URL pointing
// directly to the right place in the developer console to flip the bit.
type Help struct {
	// URL(s) pointing to additional information on handling the current error.
	Links []*Help_Link `protobuf:"bytes,1,rep,name=links" json:"links,omitempty"`
}

func (m *Help) Reset()                    { *m = Help{} }
func (m *Help) String() string            { return proto.CompactTextString(m) }
func (*Help) ProtoMessage()               {}
func (*Help) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *Help) GetLinks() []*Help_Link {
	if m != nil {
		return m.Links
	}
	return nil
}

// Describes a URL link.
type Help_Link struct {
	// Describes what the link offers.
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// The URL of the link.
	Url string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *Help_Link) Reset()                    { *m = Help_Link{} }
func (m *Help_Link) String() string            { return proto.CompactTextString(m) }
func (*Help_Link) ProtoMessage()               {}
func (*Help_Link) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6, 0} }

func (m *Help_Link) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Help_Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// Provides a localized error message that is safe to return to the user
// which can be attached to an RPC error.
type LocalizedMessage struct {
	// The locale used following the specification defined at
	// http://www.rfc-editor.org/rfc/bcp/bcp47.txt.
	// Examples are: "en-US", "fr-CH", "es-MX"
	Locale string `protobuf:"bytes,1,opt,name=locale" json:"locale,omitempty"`
	// The localized error message in the above locale.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *LocalizedMessage) Reset()                    { *m = LocalizedMessage{} }
func (m *LocalizedMessage) String() string            { return proto.CompactTextString(m) }
func (*LocalizedMessage) ProtoMessage()               {}
func (*LocalizedMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *LocalizedMessage) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *LocalizedMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RetryInfo)(nil), "google.rpc.RetryInfo")
	proto.RegisterType((*DebugInfo)(nil), "google.rpc.DebugInfo")
	proto.RegisterType((*QuotaFailure)(nil), "google.rpc.QuotaFailure")
	proto.RegisterType((*QuotaFailure_Violation)(nil), "google.rpc.QuotaFailure.Violation")
	proto.RegisterType((*BadRequest)(nil), "google.rpc.BadRequest")
	proto.RegisterType((*BadRequest_FieldViolation)(nil), "google.rpc.BadRequest.FieldViolation")
	proto.RegisterType((*RequestInfo)(nil), "google.rpc.RequestInfo")
	proto.RegisterType((*ResourceInfo)(nil), "google.rpc.ResourceInfo")
	proto.RegisterType((*Help)(nil), "google.rpc.Help")
	proto.RegisterType((*Help_Link)(nil), "google.rpc.Help.Link")
	proto.RegisterType((*LocalizedMessage)(nil), "google.rpc.LocalizedMessage")
}

func init() { proto.RegisterFile("google/rpc/error_details.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0xc7, 0x95, 0xb5, 0xdb, 0x4f, 0x79, 0xed, 0x6f, 0x14, 0x0b, 0x50, 0xa9, 0x04, 0x2a, 0x41,
	0x48, 0x95, 0x26, 0xa5, 0xd2, 0xb8, 0xed, 0x58, 0xb2, 0xad, 0x93, 0x06, 0x94, 0x08, 0x71, 0x8d,
	0xdc, 0xe4, 0x35, 0x32, 0x4d, 0xe3, 0x60, 0x3b, 0x43, 0xe5, 0xaf, 0xe0, 0xce, 0x8d, 0x13, 0x7f,
	0x26, 0x72, 0x6c, 0xaf, 0xe9, 0x7a, 0xe1, 0x96, 0xef, 0xf3, 0x27, 0x5f, 0x7f, 0x9f, 0xed, 0x07,
	0x2f, 0x73, 0xce, 0xf3, 0x02, 0xa7, 0xa2, 0x4a, 0xa7, 0x28, 0x04, 0x17, 0x49, 0x86, 0x8a, 0xb2,
	0x42, 0x86, 0x95, 0xe0, 0x8a, 0x13, 0x30, 0xeb, 0xa1, 0xa8, 0xd2, 0x91, 0x63, 0x9b, 0x95, 0x65,
	0xbd, 0x9a, 0x66, 0xb5, 0xa0, 0x8a, 0xf1, 0xd2, 0xb0, 0xc1, 0x35, 0xf8, 0x31, 0x2a, 0xb1, 0xbd,
	0x29, 0x57, 0x9c, 0x5c, 0x40, 0x4f, 0x68, 0x91, 0x64, 0x58, 0xd0, 0xed, 0xd0, 0x1b, 0x7b, 0x93,
	0xde, 0xf9, 0xf3, 0xd0, 0xda, 0x39, 0x8b, 0x30, 0xb2, 0x16, 0x31, 0x34, 0x74, 0xa4, 0xe1, 0x60,
	0x0e, 0x7e, 0x84, 0xcb, 0x3a, 0x6f, 0x8c, 0x5e, 0xc3, 0xff, 0x52, 0xd1, 0x74, 0x9d, 0x60, 0xa9,
	0x04, 0x43, 0x39, 0xf4, 0xc6, 0x9d, 0x89, 0x1f, 0xf7, 0x9b, 0xe2, 0xa5, 0xa9, 0x91, 0x67, 0x70,
	0x62, 0x72, 0x0f, 0x8f, 0xc6, 0xde, 0xc4, 0x8f, 0xad, 0x0a, 0x7e, 0x79, 0xd0, 0xff, 0x54, 0x73,
	0x45, 0xaf, 0x28, 0x2b, 0x6a, 0x81, 0x64, 0x06, 0x70, 0xc7, 0x78, 0xd1, 0xec, 0x69, 0xac, 0x7a,
	0xe7, 0x41, 0xb8, 0x6b, 0x32, 0x6c, 0xd3, 0xe1, 0x17, 0x87, 0xc6, 0xad, 0xbf, 0x46, 0xd7, 0xe0,
	0xdf, 0x2f, 0x90, 0x21, 0xfc, 0x27, 0xeb, 0xe5, 0x57, 0x4c, 0x55, 0xd3, 0xa3, 0x1f, 0x3b, 0x49,
	0xc6, 0xd0, 0xcb, 0x50, 0xa6, 0x82, 0x55, 0x1a, 0xb4, 0xc1, 0xda, 0xa5, 0xe0, 0x8f, 0x07, 0x30,
	0xa3, 0x59, 0x8c, 0xdf, 0x6a, 0x94, 0x8a, 0x2c, 0x60, 0xb0, 0x62, 0x58, 0x64, 0xc9, 0x41, 0xc2,
	0x37, 0xed, 0x84, 0xbb, 0x3f, 0xc2, 0x2b, 0x8d, 0xef, 0x42, 0x3e, 0x5a, 0xed, 0x69, 0x39, 0x9a,
	0xc3, 0xe9, 0x3e, 0x42, 0x9e, 0xc0, 0x71, 0x03, 0xd9, 0xb0, 0x46, 0xfc, 0x43, 0xd4, 0x8f, 0xd0,
	0xb3, 0x9b, 0x36, 0x97, 0xf2, 0x02, 0x40, 0x18, 0x99, 0x30, 0xe7, 0xe5, 0xdb, 0xca, 0x4d, 0x46,
	0x5e, 0x41, 0x5f, 0xa2, 0xb8, 0x63, 0x65, 0x9e, 0x64, 0x54, 0x51, 0x67, 0x68, 0x6b, 0x11, 0x55,
	0x34, 0xf8, 0xe9, 0x41, 0x3f, 0x46, 0xc9, 0x6b, 0x91, 0xa2, 0xbb, 0x67, 0x61, 0x75, 0xa2, 0xb6,
	0x15, 0x5a, 0xd7, 0xbe, 0x2b, 0x7e, 0xde, 0x56, 0xb8, 0x07, 0x95, 0x74, 0x83, 0xd6, 0xf9, 0x1e,
	0xfa, 0x40, 0x37, 0xa8, 0x7b, 0xe4, 0xdf, 0x4b, 0x14, 0xc3, 0x8e, 0xe9, 0xb1, 0x11, 0x0f, 0x7b,
	0xec, 0x1e, 0xf6, 0xc8, 0xa1, 0x3b, 0xc7, 0xa2, 0x22, 0x67, 0x70, 0x5c, 0xb0, 0x72, 0xed, 0x0e,
	0xff, 0x69, 0xfb, 0xf0, 0x35, 0x10, 0xde, 0xb2, 0x72, 0x1d, 0x1b, 0x66, 0x74, 0x01, 0x5d, 0x2d,
	0x1f, 0xda, 0x7b, 0x07, 0xf6, 0x64, 0x00, 0x9d, 0x5a, 0xb8, 0x07, 0xaa, 0x3f, 0x83, 0x08, 0x06,
	0xb7, 0x3c, 0xa5, 0x05, 0xfb, 0x81, 0xd9, 0x7b, 0x94, 0x92, 0xe6, 0xa8, 0x5f, 0x72, 0xa1, 0x6b,
	0xae, 0x7f, 0xab, 0xf4, 0x3b, 0xdb, 0x18, 0xc4, 0x3a, 0x38, 0x39, 0x3b, 0x83, 0xd3, 0x94, 0x6f,
	0x5a, 0x21, 0x67, 0x8f, 0x2f, 0xf5, 0x24, 0x47, 0x66, 0x90, 0x17, 0x7a, 0xd4, 0x16, 0xde, 0xef,
	0xa3, 0x4e, 0xbc, 0x78, 0xb7, 0x3c, 0x69, 0x26, 0xef, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf4, 0x50, 0x88, 0x09, 0xf8, 0x03, 0x00, 0x00,
}