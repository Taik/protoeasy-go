// Code generated by protoc-gen-go.
// source: google/logging/v2/logging_metrics.proto
// DO NOT EDIT!

package google_logging_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Describes a logs-based metric.  The value of the metric is the
// number of log entries that match a logs filter.
type LogMetric struct {
	// Required. The client-assigned metric identifier. Example:
	// `"severe_errors"`.  Metric identifiers are limited to 1000
	// characters and can include only the following characters: `A-Z`,
	// `a-z`, `0-9`, and the special characters `_-.,+!*',()%/\`.  The
	// forward-slash character (`/`) denotes a hierarchy of name pieces,
	// and it cannot be the first character of the name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this metric, which is used in documentation.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// An [advanced logs filter](/logging/docs/view/advanced_filters).
	// Example: `"logName:syslog AND severity>=ERROR"`.
	Filter string `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
}

func (m *LogMetric) Reset()                    { *m = LogMetric{} }
func (m *LogMetric) String() string            { return proto.CompactTextString(m) }
func (*LogMetric) ProtoMessage()               {}
func (*LogMetric) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// The parameters to ListLogMetrics.
type ListLogMetricsRequest struct {
	// Required. The resource name of the project containing the metrics.
	// Example: `"projects/my-project-id"`.
	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	// Optional. If the `pageToken` request parameter is supplied, then the next
	// page of results in the set are retrieved.  The `pageToken` parameter must
	// be set with the value of the `nextPageToken` result parameter from the
	// previous request.  The value of `projectName` must
	// be the same as in the previous request.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// Optional. The maximum number of results to return from this request.  Fewer
	// results might be returned. You must check for the `nextPageToken` result to
	// determine if additional results are available, which you can retrieve by
	// passing the `nextPageToken` value in the `pageToken` parameter to the next
	// request.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListLogMetricsRequest) Reset()                    { *m = ListLogMetricsRequest{} }
func (m *ListLogMetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogMetricsRequest) ProtoMessage()               {}
func (*ListLogMetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

// Result returned from ListLogMetrics.
type ListLogMetricsResponse struct {
	// A list of logs-based metrics.
	Metrics []*LogMetric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is given
	// a value in the response.  To get the next batch of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogMetricsResponse) Reset()                    { *m = ListLogMetricsResponse{} }
func (m *ListLogMetricsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogMetricsResponse) ProtoMessage()               {}
func (*ListLogMetricsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ListLogMetricsResponse) GetMetrics() []*LogMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// The parameters to GetLogMetric.
type GetLogMetricRequest struct {
	// The resource name of the desired metric.
	// Example: `"projects/my-project-id/metrics/my-metric-id"`.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
}

func (m *GetLogMetricRequest) Reset()                    { *m = GetLogMetricRequest{} }
func (m *GetLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLogMetricRequest) ProtoMessage()               {}
func (*GetLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

// The parameters to CreateLogMetric.
type CreateLogMetricRequest struct {
	// The resource name of the project in which to create the metric.
	// Example: `"projects/my-project-id"`.
	//
	// The new metric must be provided in the request.
	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	// The new logs-based metric, which must not have an identifier that
	// already exists.
	Metric *LogMetric `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
}

func (m *CreateLogMetricRequest) Reset()                    { *m = CreateLogMetricRequest{} }
func (m *CreateLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateLogMetricRequest) ProtoMessage()               {}
func (*CreateLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *CreateLogMetricRequest) GetMetric() *LogMetric {
	if m != nil {
		return m.Metric
	}
	return nil
}

// The parameters to UpdateLogMetric.
//
type UpdateLogMetricRequest struct {
	// The resource name of the metric to update.
	// Example: `"projects/my-project-id/metrics/my-metric-id"`.
	//
	// The updated metric must be provided in the request and have the
	// same identifier that is specified in `metricName`.
	// If the metric does not exist, it is created.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
	// The updated metric, whose name must be the same as the
	// metric identifier in `metricName`. If `metricName` does not
	// exist, then a new metric is created.
	Metric *LogMetric `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
}

func (m *UpdateLogMetricRequest) Reset()                    { *m = UpdateLogMetricRequest{} }
func (m *UpdateLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateLogMetricRequest) ProtoMessage()               {}
func (*UpdateLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *UpdateLogMetricRequest) GetMetric() *LogMetric {
	if m != nil {
		return m.Metric
	}
	return nil
}

// The parameters to DeleteLogMetric.
type DeleteLogMetricRequest struct {
	// The resource name of the metric to delete.
	// Example: `"projects/my-project-id/metrics/my-metric-id"`.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
}

func (m *DeleteLogMetricRequest) Reset()                    { *m = DeleteLogMetricRequest{} }
func (m *DeleteLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogMetricRequest) ProtoMessage()               {}
func (*DeleteLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func init() {
	proto.RegisterType((*LogMetric)(nil), "google.logging.v2.LogMetric")
	proto.RegisterType((*ListLogMetricsRequest)(nil), "google.logging.v2.ListLogMetricsRequest")
	proto.RegisterType((*ListLogMetricsResponse)(nil), "google.logging.v2.ListLogMetricsResponse")
	proto.RegisterType((*GetLogMetricRequest)(nil), "google.logging.v2.GetLogMetricRequest")
	proto.RegisterType((*CreateLogMetricRequest)(nil), "google.logging.v2.CreateLogMetricRequest")
	proto.RegisterType((*UpdateLogMetricRequest)(nil), "google.logging.v2.UpdateLogMetricRequest")
	proto.RegisterType((*DeleteLogMetricRequest)(nil), "google.logging.v2.DeleteLogMetricRequest")
}

func init() { proto.RegisterFile("google/logging/v2/logging_metrics.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0x96, 0x06, 0x32, 0x29, 0x04, 0x16, 0xd5, 0x0a, 0x6e, 0x11, 0xc5, 0x87, 0x12,
	0x02, 0xd8, 0xc2, 0x2d, 0x95, 0x28, 0xe2, 0xc2, 0x1f, 0x71, 0x29, 0xa8, 0x4a, 0x01, 0x89, 0x93,
	0xe5, 0xb8, 0x53, 0x6b, 0x21, 0xf1, 0xba, 0xf6, 0x36, 0x2a, 0x20, 0x2e, 0xdc, 0x38, 0x23, 0x81,
	0x78, 0x2e, 0x5e, 0x81, 0x97, 0xe0, 0x86, 0xbd, 0xde, 0x4d, 0x4d, 0xb2, 0x6a, 0x9a, 0xde, 0xb2,
	0xb3, 0xb3, 0xf3, 0xfd, 0x66, 0xe6, 0x8b, 0xe1, 0x56, 0xc4, 0x58, 0xd4, 0x47, 0xb7, 0xcf, 0xa2,
	0x88, 0xc6, 0x91, 0x3b, 0xf4, 0xd4, 0x4f, 0x7f, 0x80, 0x3c, 0xa5, 0x61, 0xe6, 0x24, 0x29, 0xe3,
	0x8c, 0x5c, 0x29, 0x13, 0x1d, 0x79, 0xeb, 0x0c, 0x3d, 0x6b, 0x45, 0xbe, 0x0d, 0x12, 0xea, 0x06,
	0x71, 0xcc, 0x78, 0xc0, 0x29, 0x8b, 0xe5, 0x03, 0x6b, 0x59, 0xde, 0x8a, 0x53, 0xef, 0x70, 0xdf,
	0xc5, 0x41, 0xc2, 0x3f, 0x96, 0x97, 0xf6, 0x3b, 0xa8, 0x6f, 0xb3, 0xe8, 0xa5, 0x50, 0x20, 0x04,
	0xce, 0xc5, 0xc1, 0x00, 0x5b, 0xc6, 0xaa, 0xd1, 0xae, 0x77, 0xc5, 0x6f, 0xb2, 0x0a, 0x8d, 0x3d,
	0xcc, 0xc2, 0x94, 0x26, 0x45, 0xcd, 0xd6, 0x9c, 0xb8, 0xaa, 0x86, 0x88, 0x09, 0xb5, 0x7d, 0xda,
	0xe7, 0x98, 0xb6, 0xe6, 0xc5, 0xa5, 0x3c, 0xd9, 0x43, 0x58, 0xda, 0xa6, 0x19, 0x1f, 0x95, 0xcf,
	0xba, 0x78, 0x70, 0x88, 0x19, 0x27, 0x37, 0x61, 0x31, 0x17, 0x7f, 0x8f, 0x21, 0xf7, 0x2b, 0x72,
	0x0d, 0x19, 0x7b, 0x55, 0xa8, 0x5e, 0x07, 0x48, 0x82, 0x08, 0x7d, 0xce, 0x3e, 0xa0, 0x12, 0xad,
	0x17, 0x91, 0xd7, 0x45, 0x80, 0x2c, 0x83, 0x38, 0xf8, 0x19, 0xfd, 0x84, 0x42, 0x75, 0xa1, 0x7b,
	0xa1, 0x08, 0xec, 0xe6, 0x67, 0xfb, 0x08, 0xcc, 0x71, 0xdd, 0x2c, 0xc9, 0xc7, 0x81, 0x64, 0x13,
	0xce, 0xcb, 0x59, 0xe6, 0x9a, 0xf3, 0xed, 0x86, 0xb7, 0xe2, 0x4c, 0x0c, 0xd3, 0x19, 0xbd, 0xeb,
	0xaa, 0x64, 0xb2, 0x06, 0xcd, 0x18, 0x8f, 0xb8, 0x3f, 0x81, 0x74, 0xb1, 0x08, 0xef, 0x28, 0x2c,
	0x7b, 0x13, 0xae, 0xbe, 0xc0, 0x63, 0x61, 0xd5, 0xef, 0x0d, 0x68, 0x94, 0x95, 0xaa, 0xed, 0x42,
	0x19, 0x2a, 0xba, 0xb5, 0x0f, 0xc0, 0x7c, 0x9a, 0x62, 0xc0, 0x71, 0xe2, 0xe9, 0x29, 0x46, 0xb5,
	0x01, 0xb5, 0xb2, 0x94, 0x60, 0x9a, 0xd6, 0x93, 0xcc, 0xb5, 0x19, 0x98, 0x6f, 0x92, 0x3d, 0x9d,
	0xe4, 0x34, 0xda, 0x33, 0x0a, 0x3e, 0x04, 0xf3, 0x19, 0xf6, 0xf1, 0x0c, 0x82, 0xde, 0xdf, 0x05,
	0xb8, 0x2c, 0x57, 0xb9, 0x8b, 0xe9, 0x90, 0x86, 0xf8, 0xd6, 0x23, 0xbf, 0x0c, 0xb8, 0xf4, 0xff,
	0x9a, 0x49, 0x5b, 0x07, 0xa2, 0x73, 0xa0, 0x75, 0xfb, 0x14, 0x99, 0xa5, 0x67, 0x6c, 0xef, 0xeb,
	0xef, 0x3f, 0xdf, 0xe7, 0xee, 0x92, 0x4e, 0xfe, 0x8f, 0xec, 0x21, 0x0f, 0xee, 0xbb, 0x9f, 0xab,
	0x1b, 0x79, 0x2c, 0x0f, 0x99, 0xdb, 0xf9, 0xe2, 0x2a, 0xbf, 0x7c, 0x33, 0x60, 0xb1, 0x6a, 0x04,
	0xb2, 0xa6, 0xd1, 0xd3, 0x38, 0xc5, 0x3a, 0x71, 0x94, 0xf6, 0xba, 0x40, 0xb9, 0x47, 0xee, 0x1c,
	0xa3, 0x54, 0x06, 0x57, 0x21, 0x51, 0x20, 0x39, 0x13, 0xf9, 0x61, 0x40, 0x73, 0xcc, 0x5c, 0x44,
	0xd7, 0xbe, 0xde, 0x80, 0x53, 0x88, 0xb6, 0x04, 0xd1, 0x86, 0x3d, 0xc3, 0x70, 0xb6, 0xa4, 0x21,
	0xc8, 0xcf, 0x1c, 0x6c, 0xcc, 0x82, 0x5a, 0x30, 0xbd, 0x4d, 0xa7, 0x80, 0x3d, 0x12, 0x60, 0x0f,
	0xac, 0x59, 0x46, 0x35, 0x22, 0xcb, 0xd7, 0xd7, 0x1c, 0xf3, 0xaa, 0x96, 0x4c, 0xef, 0x67, 0xcb,
	0x54, 0xa9, 0xea, 0x83, 0xeb, 0x3c, 0x2f, 0x3e, 0xb8, 0x6a, 0x7d, 0x9d, 0x59, 0x98, 0x9e, 0x5c,
	0x83, 0xa5, 0x90, 0x0d, 0x26, 0xc5, 0x77, 0x8c, 0x5e, 0x4d, 0xd4, 0x5f, 0xff, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x82, 0xc7, 0xc3, 0x27, 0x3a, 0x06, 0x00, 0x00,
}
