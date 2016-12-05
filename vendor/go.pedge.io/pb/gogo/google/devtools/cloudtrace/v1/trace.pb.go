// Code generated by protoc-gen-gogo.
// source: google/devtools/cloudtrace/v1/trace.proto
// DO NOT EDIT!

package google_devtools_cloudtrace_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/protobuf"
import google_protobuf2 "go.pedge.io/pb/gogo/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Type of span. Can be used to specify additional relationships between spans
// in addition to a parent/child relationship.
type TraceSpan_SpanKind int32

const (
	// Unspecified.
	TraceSpan_SPAN_KIND_UNSPECIFIED TraceSpan_SpanKind = 0
	// Indicates that the span covers server-side handling of an RPC or other
	// remote network request.
	TraceSpan_RPC_SERVER TraceSpan_SpanKind = 1
	// Indicates that the span covers the client-side wrapper around an RPC or
	// other remote request.
	TraceSpan_RPC_CLIENT TraceSpan_SpanKind = 2
)

var TraceSpan_SpanKind_name = map[int32]string{
	0: "SPAN_KIND_UNSPECIFIED",
	1: "RPC_SERVER",
	2: "RPC_CLIENT",
}
var TraceSpan_SpanKind_value = map[string]int32{
	"SPAN_KIND_UNSPECIFIED": 0,
	"RPC_SERVER":            1,
	"RPC_CLIENT":            2,
}

func (x TraceSpan_SpanKind) String() string {
	return proto.EnumName(TraceSpan_SpanKind_name, int32(x))
}
func (TraceSpan_SpanKind) EnumDescriptor() ([]byte, []int) { return fileDescriptorTrace, []int{2, 0} }

// Type of data returned for traces in the list.
type ListTracesRequest_ViewType int32

const (
	// Default is `MINIMAL` if unspecified.
	ListTracesRequest_VIEW_TYPE_UNSPECIFIED ListTracesRequest_ViewType = 0
	// Minimal view of the trace record that contains only the project
	// and trace IDs.
	ListTracesRequest_MINIMAL ListTracesRequest_ViewType = 1
	// Root span view of the trace record that returns the root spans along
	// with the minimal trace data.
	ListTracesRequest_ROOTSPAN ListTracesRequest_ViewType = 2
	// Complete view of the trace record that contains the actual trace data.
	// This is equivalent to calling the REST `get` or RPC `GetTrace` method
	// using the ID of each listed trace.
	ListTracesRequest_COMPLETE ListTracesRequest_ViewType = 3
)

var ListTracesRequest_ViewType_name = map[int32]string{
	0: "VIEW_TYPE_UNSPECIFIED",
	1: "MINIMAL",
	2: "ROOTSPAN",
	3: "COMPLETE",
}
var ListTracesRequest_ViewType_value = map[string]int32{
	"VIEW_TYPE_UNSPECIFIED": 0,
	"MINIMAL":               1,
	"ROOTSPAN":              2,
	"COMPLETE":              3,
}

func (x ListTracesRequest_ViewType) String() string {
	return proto.EnumName(ListTracesRequest_ViewType_name, int32(x))
}
func (ListTracesRequest_ViewType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorTrace, []int{3, 0}
}

// A trace describes how long it takes for an application to perform an
// operation. It consists of a set of spans, each of which represent a single
// timed event within the operation.
type Trace struct {
	// Project ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Globally unique identifier for the trace. This identifier is a 128-bit
	// numeric value formatted as a 32-byte hex string.
	TraceId string `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Collection of spans in the trace.
	Spans []*TraceSpan `protobuf:"bytes,3,rep,name=spans" json:"spans,omitempty"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{0} }

func (m *Trace) GetSpans() []*TraceSpan {
	if m != nil {
		return m.Spans
	}
	return nil
}

// List of new or updated traces.
type Traces struct {
	// List of traces.
	Traces []*Trace `protobuf:"bytes,1,rep,name=traces" json:"traces,omitempty"`
}

func (m *Traces) Reset()                    { *m = Traces{} }
func (m *Traces) String() string            { return proto.CompactTextString(m) }
func (*Traces) ProtoMessage()               {}
func (*Traces) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{1} }

func (m *Traces) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

// A span represents a single timed event within a trace. Spans can be nested
// and form a trace tree. Often, a trace contains a root span that describes the
// end-to-end latency of an operation and, optionally, one or more subspans for
// its suboperations. Spans do not need to be contiguous. There may be gaps
// between spans in a trace.
type TraceSpan struct {
	// Identifier for the span. Must be a 64-bit integer other than 0 and
	// unique within a trace.
	SpanId uint64 `protobuf:"fixed64,1,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// Distinguishes between spans generated in a particular context. For example,
	// two spans with the same name may be distinguished using `RPC_CLIENT`
	// and `RPC_SERVER` to identify queueing latency associated with the span.
	Kind TraceSpan_SpanKind `protobuf:"varint,2,opt,name=kind,proto3,enum=google.devtools.cloudtrace.v1.TraceSpan_SpanKind" json:"kind,omitempty"`
	// Name of the trace. The trace name is sanitized and displayed in the
	// Stackdriver Trace tool in the Google Developers Console.
	// The name may be a method name or some other per-call site name.
	// For the same executable and the same call point, a best practice is
	// to use a consistent name, which makes it easier to correlate
	// cross-trace spans.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Start time of the span in nanoseconds from the UNIX epoch.
	StartTime *google_protobuf2.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// End time of the span in nanoseconds from the UNIX epoch.
	EndTime *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// ID of the parent span, if any. Optional.
	ParentSpanId uint64 `protobuf:"fixed64,6,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	// Collection of labels associated with the span.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *TraceSpan) Reset()                    { *m = TraceSpan{} }
func (m *TraceSpan) String() string            { return proto.CompactTextString(m) }
func (*TraceSpan) ProtoMessage()               {}
func (*TraceSpan) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{2} }

func (m *TraceSpan) GetStartTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TraceSpan) GetEndTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TraceSpan) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// The request message for the `ListTraces` method. All fields are required
// unless specified.
type ListTracesRequest struct {
	// ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Type of data returned for traces in the list. Optional. Default is
	// `MINIMAL`.
	View ListTracesRequest_ViewType `protobuf:"varint,2,opt,name=view,proto3,enum=google.devtools.cloudtrace.v1.ListTracesRequest_ViewType" json:"view,omitempty"`
	// Maximum number of traces to return. If not specified or <= 0, the
	// implementation selects a reasonable value.  The implementation may
	// return fewer traces than the requested page size. Optional.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token identifying the page of results to return. If provided, use the
	// value of the `next_page_token` field from a previous request. Optional.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// End of the time interval (inclusive) during which the trace data was
	// collected from the application.
	StartTime *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// Start of the time interval (inclusive) during which the trace data was
	// collected from the application.
	EndTime *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// An optional filter for the request.
	Filter string `protobuf:"bytes,7,opt,name=filter,proto3" json:"filter,omitempty"`
	// Field used to sort the returned traces. Optional.
	// Can be one of the following:
	//
	// *   `trace_id`
	// *   `name` (`name` field of root span in the trace)
	// *   `duration` (difference between `end_time` and `start_time` fields of
	//      the root span)
	// *   `start` (`start_time` field of the root span)
	//
	// Descending order can be specified by appending `desc` to the sort field
	// (for example, `name desc`).
	//
	// Only one sort field is permitted.
	OrderBy string `protobuf:"bytes,8,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (m *ListTracesRequest) Reset()                    { *m = ListTracesRequest{} }
func (m *ListTracesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTracesRequest) ProtoMessage()               {}
func (*ListTracesRequest) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{3} }

func (m *ListTracesRequest) GetStartTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ListTracesRequest) GetEndTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// The response message for the `ListTraces` method.
type ListTracesResponse struct {
	// List of trace records returned.
	Traces []*Trace `protobuf:"bytes,1,rep,name=traces" json:"traces,omitempty"`
	// If defined, indicates that there are more traces that match the request
	// and that this value should be passed to the next request to continue
	// retrieving additional traces.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListTracesResponse) Reset()                    { *m = ListTracesResponse{} }
func (m *ListTracesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTracesResponse) ProtoMessage()               {}
func (*ListTracesResponse) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{4} }

func (m *ListTracesResponse) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

// The request message for the `GetTrace` method.
type GetTraceRequest struct {
	// ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// ID of the trace to return.
	TraceId string `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (m *GetTraceRequest) Reset()                    { *m = GetTraceRequest{} }
func (m *GetTraceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTraceRequest) ProtoMessage()               {}
func (*GetTraceRequest) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{5} }

// The request message for the `PatchTraces` method.
type PatchTracesRequest struct {
	// ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The body of the message.
	Traces *Traces `protobuf:"bytes,2,opt,name=traces" json:"traces,omitempty"`
}

func (m *PatchTracesRequest) Reset()                    { *m = PatchTracesRequest{} }
func (m *PatchTracesRequest) String() string            { return proto.CompactTextString(m) }
func (*PatchTracesRequest) ProtoMessage()               {}
func (*PatchTracesRequest) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{6} }

func (m *PatchTracesRequest) GetTraces() *Traces {
	if m != nil {
		return m.Traces
	}
	return nil
}

func init() {
	proto.RegisterType((*Trace)(nil), "google.devtools.cloudtrace.v1.Trace")
	proto.RegisterType((*Traces)(nil), "google.devtools.cloudtrace.v1.Traces")
	proto.RegisterType((*TraceSpan)(nil), "google.devtools.cloudtrace.v1.TraceSpan")
	proto.RegisterType((*ListTracesRequest)(nil), "google.devtools.cloudtrace.v1.ListTracesRequest")
	proto.RegisterType((*ListTracesResponse)(nil), "google.devtools.cloudtrace.v1.ListTracesResponse")
	proto.RegisterType((*GetTraceRequest)(nil), "google.devtools.cloudtrace.v1.GetTraceRequest")
	proto.RegisterType((*PatchTracesRequest)(nil), "google.devtools.cloudtrace.v1.PatchTracesRequest")
	proto.RegisterEnum("google.devtools.cloudtrace.v1.TraceSpan_SpanKind", TraceSpan_SpanKind_name, TraceSpan_SpanKind_value)
	proto.RegisterEnum("google.devtools.cloudtrace.v1.ListTracesRequest_ViewType", ListTracesRequest_ViewType_name, ListTracesRequest_ViewType_value)
}

func init() { proto.RegisterFile("google/devtools/cloudtrace/v1/trace.proto", fileDescriptorTrace) }

var fileDescriptorTrace = []byte{
	// 840 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x71, 0xe2, 0x38, 0x27, 0xa5, 0x1b, 0x46, 0xb0, 0x78, 0xb3, 0xac, 0x08, 0xd6, 0x82,
	0x02, 0x08, 0x7b, 0x93, 0x05, 0x89, 0xae, 0x00, 0x69, 0x9b, 0xf5, 0x22, 0xab, 0x49, 0x6a, 0x39,
	0xa6, 0x88, 0x2b, 0xcb, 0x89, 0x67, 0x83, 0x69, 0x62, 0x1b, 0xcf, 0x24, 0x4b, 0x5a, 0xf5, 0x02,
	0x2e, 0xb9, 0x45, 0x5c, 0xf1, 0x16, 0xbc, 0x09, 0xe2, 0x15, 0x90, 0x78, 0x0d, 0x34, 0x33, 0x76,
	0x1b, 0xa5, 0xa2, 0x49, 0xe1, 0x26, 0x9a, 0x73, 0x66, 0xce, 0xcf, 0x77, 0xbe, 0xef, 0xc4, 0xf0,
	0xfe, 0x34, 0x49, 0xa6, 0x33, 0x6c, 0x86, 0x78, 0x49, 0x93, 0x64, 0x46, 0xcc, 0xc9, 0x2c, 0x59,
	0x84, 0x34, 0x0b, 0x26, 0xd8, 0x5c, 0x76, 0x4c, 0x7e, 0x30, 0xd2, 0x2c, 0xa1, 0x09, 0x7a, 0x20,
	0x9e, 0x1a, 0xc5, 0x53, 0xe3, 0xea, 0xa9, 0xb1, 0xec, 0x34, 0xdf, 0xca, 0x33, 0x05, 0x69, 0x64,
	0x06, 0x71, 0x9c, 0xd0, 0x80, 0x46, 0x49, 0x4c, 0x44, 0x70, 0xf3, 0x7e, 0x7e, 0xcb, 0xad, 0xf1,
	0xe2, 0x85, 0x89, 0xe7, 0x29, 0x5d, 0xe5, 0x97, 0x6f, 0x6f, 0x5e, 0xd2, 0x68, 0x8e, 0x09, 0x0d,
	0xe6, 0xa9, 0x78, 0xa0, 0xff, 0x28, 0x41, 0xc5, 0x63, 0x85, 0xd0, 0x03, 0x80, 0x34, 0x4b, 0xbe,
	0xc3, 0x13, 0xea, 0x47, 0xa1, 0x26, 0xb5, 0xa4, 0x76, 0xcd, 0xad, 0xe5, 0x1e, 0x3b, 0x44, 0xf7,
	0x40, 0xe5, 0x0d, 0xb1, 0xcb, 0x12, 0xbf, 0xac, 0x72, 0xdb, 0x0e, 0xd1, 0x17, 0x50, 0x21, 0x69,
	0x10, 0x13, 0x4d, 0x6e, 0xc9, 0xed, 0x7a, 0xb7, 0x6d, 0xdc, 0x08, 0xc7, 0xe0, 0xe5, 0x46, 0x69,
	0x10, 0xbb, 0x22, 0x4c, 0x7f, 0x0e, 0x0a, 0xf7, 0x11, 0xf4, 0x19, 0x28, 0xfc, 0x19, 0xd1, 0x24,
	0x9e, 0xea, 0xe1, 0x2e, 0xa9, 0xdc, 0x3c, 0x46, 0xff, 0x5b, 0x86, 0xda, 0x65, 0x72, 0xf4, 0x26,
	0x54, 0x59, 0xfa, 0x02, 0x8c, 0xe2, 0x2a, 0xcc, 0xb4, 0x43, 0x64, 0x41, 0xf9, 0x34, 0x8a, 0x05,
	0x8a, 0xfd, 0x6e, 0x67, 0xd7, 0x6e, 0x0d, 0xf6, 0x73, 0x14, 0xc5, 0xa1, 0xcb, 0xc3, 0x11, 0x82,
	0x72, 0x1c, 0xcc, 0xb1, 0x26, 0xf3, 0x61, 0xf0, 0x33, 0x3a, 0x00, 0x20, 0x34, 0xc8, 0xa8, 0xcf,
	0xc6, 0xac, 0x95, 0x5b, 0x52, 0xbb, 0xde, 0x6d, 0x16, 0x05, 0x0a, 0x0e, 0x0c, 0xaf, 0xe0, 0xc0,
	0xad, 0xf1, 0xd7, 0xcc, 0x46, 0x9f, 0x80, 0x8a, 0xe3, 0x50, 0x04, 0x56, 0xb6, 0x06, 0x56, 0x71,
	0x1c, 0xf2, 0xb0, 0x87, 0xb0, 0x9f, 0x06, 0x19, 0x8e, 0xa9, 0x5f, 0x80, 0x55, 0x38, 0xd8, 0x3d,
	0xe1, 0x1d, 0x09, 0xc8, 0x7d, 0x50, 0x66, 0xc1, 0x18, 0xcf, 0x88, 0x56, 0xe5, 0x73, 0xfd, 0x78,
	0x67, 0xd0, 0x7d, 0x1e, 0x66, 0xc5, 0x34, 0x5b, 0xb9, 0x79, 0x8e, 0xe6, 0x01, 0xd4, 0xd7, 0xdc,
	0xa8, 0x01, 0xf2, 0x29, 0x5e, 0xe5, 0x8a, 0x61, 0x47, 0xf4, 0x3a, 0x54, 0x96, 0xc1, 0x6c, 0x81,
	0x73, 0xa1, 0x08, 0xe3, 0x49, 0xe9, 0x53, 0x49, 0xb7, 0x40, 0x2d, 0xc6, 0x88, 0xee, 0xc1, 0x1b,
	0x23, 0xe7, 0xe9, 0xd0, 0x3f, 0xb2, 0x87, 0xcf, 0xfc, 0xaf, 0x86, 0x23, 0xc7, 0xea, 0xd9, 0xcf,
	0x6d, 0xeb, 0x59, 0xe3, 0x15, 0xb4, 0x0f, 0xe0, 0x3a, 0x3d, 0x7f, 0x64, 0xb9, 0x27, 0x96, 0xdb,
	0x90, 0x0a, 0xbb, 0xd7, 0xb7, 0xad, 0xa1, 0xd7, 0x28, 0xe9, 0xbf, 0xcb, 0xf0, 0x5a, 0x3f, 0x22,
	0x54, 0xc8, 0xc6, 0xc5, 0xdf, 0x2f, 0x30, 0xa1, 0xdb, 0x14, 0x3c, 0x80, 0xf2, 0x32, 0xc2, 0x2f,
	0x73, 0xde, 0x0f, 0xb6, 0x8c, 0xe0, 0x5a, 0x7a, 0xe3, 0x24, 0xc2, 0x2f, 0xbd, 0x55, 0x8a, 0x5d,
	0x9e, 0x06, 0xdd, 0x87, 0x5a, 0x1a, 0x4c, 0xb1, 0x4f, 0xa2, 0x33, 0x21, 0x82, 0x8a, 0xab, 0x32,
	0xc7, 0x28, 0x3a, 0x13, 0xcb, 0xc4, 0x2e, 0x69, 0x72, 0x8a, 0x63, 0x2e, 0x04, 0xd6, 0x4a, 0x30,
	0xc5, 0x1e, 0x73, 0x6c, 0xe8, 0xa4, 0xf2, 0x5f, 0x75, 0xa2, 0xec, 0xae, 0x93, 0xbb, 0xa0, 0xbc,
	0x88, 0x66, 0x14, 0x67, 0x5a, 0x95, 0x37, 0x93, 0x5b, 0x6c, 0xad, 0x93, 0x2c, 0xc4, 0x99, 0x3f,
	0x5e, 0x69, 0xaa, 0x58, 0x6b, 0x6e, 0x1f, 0xae, 0xf4, 0x21, 0xa8, 0x05, 0x64, 0xc6, 0xd5, 0x89,
	0x6d, 0x7d, 0xed, 0x7b, 0xdf, 0x38, 0xd6, 0x06, 0x57, 0x75, 0xa8, 0x0e, 0xec, 0xa1, 0x3d, 0x78,
	0xda, 0x6f, 0x48, 0x68, 0x0f, 0x54, 0xf7, 0xf8, 0xd8, 0x63, 0xbc, 0x36, 0x4a, 0xcc, 0xea, 0x1d,
	0x0f, 0x9c, 0xbe, 0xe5, 0x59, 0x0d, 0x59, 0x3f, 0x03, 0xb4, 0x3e, 0x54, 0x92, 0x26, 0x31, 0xc1,
	0xff, 0x6f, 0xe5, 0xd1, 0x7b, 0x70, 0x27, 0xc6, 0x3f, 0x50, 0x7f, 0x6d, 0xd8, 0x42, 0x73, 0xaf,
	0x32, 0xb7, 0x53, 0x0c, 0x5c, 0x3f, 0x82, 0x3b, 0x5f, 0x62, 0x51, 0x7a, 0x47, 0xb5, 0xfc, 0xfb,
	0xff, 0x9d, 0x9e, 0x01, 0x72, 0x02, 0x3a, 0xf9, 0xf6, 0x56, 0xea, 0xfb, 0xfc, 0x12, 0x67, 0x89,
	0xb3, 0xf6, 0xee, 0x2e, 0x38, 0x49, 0x01, 0xb4, 0xfb, 0x87, 0x0c, 0x7b, 0x62, 0x2b, 0x71, 0xb6,
	0x8c, 0x26, 0x18, 0xfd, 0x26, 0x01, 0x5c, 0x8d, 0x13, 0x3d, 0xba, 0xad, 0x9c, 0x9b, 0x9d, 0x5b,
	0x44, 0x08, 0xae, 0xf4, 0xf6, 0x4f, 0x7f, 0xfe, 0xf5, 0x4b, 0x49, 0x47, 0x2d, 0xf6, 0x01, 0xcb,
	0xa1, 0x11, 0xf3, 0xfc, 0x0a, 0xf6, 0x85, 0x99, 0xf3, 0xf2, 0xab, 0x04, 0x6a, 0x31, 0x70, 0x64,
	0x6c, 0xa9, 0xb4, 0xc1, 0x4c, 0x73, 0x27, 0x09, 0xe8, 0x8f, 0x79, 0x33, 0x1f, 0xa1, 0x0f, 0xb7,
	0x35, 0x63, 0x9e, 0x17, 0x44, 0x5e, 0xa0, 0x9f, 0x25, 0xa8, 0xaf, 0x71, 0x87, 0xb6, 0x0d, 0xe1,
	0x3a, 0xcf, 0xcd, 0xbb, 0xd7, 0xd6, 0xcd, 0x62, 0x1f, 0x5c, 0xfd, 0x11, 0xef, 0xe7, 0x83, 0xee,
	0xd6, 0xe1, 0x3c, 0xc9, 0x39, 0x3d, 0xec, 0xc0, 0x3b, 0x93, 0x64, 0x7e, 0x73, 0x07, 0x87, 0xc0,
	0xab, 0x3b, 0xac, 0x96, 0x23, 0x8d, 0x15, 0x5e, 0xf4, 0xf1, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4e, 0xf5, 0x29, 0xd7, 0x5d, 0x08, 0x00, 0x00,
}
