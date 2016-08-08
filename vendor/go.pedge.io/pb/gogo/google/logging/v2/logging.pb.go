// Code generated by protoc-gen-gogo.
// source: google/logging/v2/logging.proto
// DO NOT EDIT!

package google_logging_v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_api3 "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/protobuf"
import _ "go.pedge.io/pb/gogo/google/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The parameters to DeleteLog.
type DeleteLogRequest struct {
	// Required. The resource name of the log to delete.  Example:
	// `"projects/my-project/logs/syslog"`.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
}

func (m *DeleteLogRequest) Reset()                    { *m = DeleteLogRequest{} }
func (m *DeleteLogRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogRequest) ProtoMessage()               {}
func (*DeleteLogRequest) Descriptor() ([]byte, []int) { return fileDescriptorLogging, []int{0} }

// The parameters to WriteLogEntries.
type WriteLogEntriesRequest struct {
	// Optional. A default log resource name for those log entries in `entries`
	// that do not specify their own `logName`.  Example:
	// `"projects/my-project/logs/syslog"`.  See
	// [LogEntry][google.logging.v2.LogEntry].
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// Optional. A default monitored resource for those log entries in `entries`
	// that do not specify their own `resource`.
	Resource *google_api3.MonitoredResource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	// Optional. User-defined `key:value` items that are added to
	// the `labels` field of each log entry in `entries`, except when a log
	// entry specifies its own `key:value` item with the same key.
	// Example: `{ "size": "large", "color":"red" }`
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. The log entries to write. The log entries must have values for
	// all required fields.
	Entries []*LogEntry `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
}

func (m *WriteLogEntriesRequest) Reset()                    { *m = WriteLogEntriesRequest{} }
func (m *WriteLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesRequest) ProtoMessage()               {}
func (*WriteLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptorLogging, []int{1} }

func (m *WriteLogEntriesRequest) GetResource() *google_api3.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Result returned from WriteLogEntries.
type WriteLogEntriesResponse struct {
}

func (m *WriteLogEntriesResponse) Reset()                    { *m = WriteLogEntriesResponse{} }
func (m *WriteLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesResponse) ProtoMessage()               {}
func (*WriteLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptorLogging, []int{2} }

// The parameters to `ListLogEntries`.
type ListLogEntriesRequest struct {
	// Required. One or more project IDs or project numbers from which to retrieve
	// log entries.  Examples of a project ID: `"my-project-1A"`, `"1234567890"`.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds" json:"project_ids,omitempty"`
	// Optional. An [advanced logs filter](/logging/docs/view/advanced_filters).
	// The filter is compared against all log entries in the projects specified by
	// `projectIds`.  Only entries that match the filter are retrieved.  An empty
	// filter matches all log entries.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. How the results should be sorted.  Presently, the only permitted
	// values are `"timestamp"` (default) and `"timestamp desc"`.  The first
	// option returns entries in order of increasing values of
	// `LogEntry.timestamp` (oldest first), and the second option returns entries
	// in order of decreasing timestamps (newest first).  Entries with equal
	// timestamps are returned in order of `LogEntry.insertId`.
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Optional. The maximum number of results to return from this request.  Fewer
	// results might be returned. You must check for the `nextPageToken` result to
	// determine if additional results are available, which you can retrieve by
	// passing the `nextPageToken` value in the `pageToken` parameter to the next
	// request.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. If the `pageToken` request parameter is supplied, then the next
	// page of results in the set are retrieved.  The `pageToken` parameter must
	// be set with the value of the `nextPageToken` result parameter from the
	// previous request.  The values of `projectIds`, `filter`, and `orderBy` must
	// be the same as in the previous request.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (m *ListLogEntriesRequest) Reset()                    { *m = ListLogEntriesRequest{} }
func (m *ListLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesRequest) ProtoMessage()               {}
func (*ListLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptorLogging, []int{3} }

// Result returned from `ListLogEntries`.
type ListLogEntriesResponse struct {
	// A list of log entries.
	Entries []*LogEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is
	// given a value in the response.  To get the next batch of results, call
	// this method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListLogEntriesResponse) Reset()                    { *m = ListLogEntriesResponse{} }
func (m *ListLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesResponse) ProtoMessage()               {}
func (*ListLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptorLogging, []int{4} }

func (m *ListLogEntriesResponse) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// The parameters to ListMonitoredResourceDescriptors
type ListMonitoredResourceDescriptorsRequest struct {
	// Optional. The maximum number of results to return from this request.  Fewer
	// results might be returned. You must check for the `nextPageToken` result to
	// determine if additional results are available, which you can retrieve by
	// passing the `nextPageToken` value in the `pageToken` parameter to the next
	// request.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. If the `pageToken` request parameter is supplied, then the next
	// page of results in the set are retrieved.  The `pageToken` parameter must
	// be set with the value of the `nextPageToken` result parameter from the
	// previous request.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsRequest) Reset() {
	*m = ListMonitoredResourceDescriptorsRequest{}
}
func (m *ListMonitoredResourceDescriptorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsRequest) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorLogging, []int{5}
}

// Result returned from ListMonitoredResourceDescriptors.
type ListMonitoredResourceDescriptorsResponse struct {
	// A list of resource descriptors.
	ResourceDescriptors []*google_api3.MonitoredResourceDescriptor `protobuf:"bytes,1,rep,name=resource_descriptors,json=resourceDescriptors" json:"resource_descriptors,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is
	// returned in the response.  To get the next batch of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsResponse) Reset() {
	*m = ListMonitoredResourceDescriptorsResponse{}
}
func (m *ListMonitoredResourceDescriptorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsResponse) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorLogging, []int{6}
}

func (m *ListMonitoredResourceDescriptorsResponse) GetResourceDescriptors() []*google_api3.MonitoredResourceDescriptor {
	if m != nil {
		return m.ResourceDescriptors
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteLogRequest)(nil), "google.logging.v2.DeleteLogRequest")
	proto.RegisterType((*WriteLogEntriesRequest)(nil), "google.logging.v2.WriteLogEntriesRequest")
	proto.RegisterType((*WriteLogEntriesResponse)(nil), "google.logging.v2.WriteLogEntriesResponse")
	proto.RegisterType((*ListLogEntriesRequest)(nil), "google.logging.v2.ListLogEntriesRequest")
	proto.RegisterType((*ListLogEntriesResponse)(nil), "google.logging.v2.ListLogEntriesResponse")
	proto.RegisterType((*ListMonitoredResourceDescriptorsRequest)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsRequest")
	proto.RegisterType((*ListMonitoredResourceDescriptorsResponse)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsResponse")
}

var fileDescriptorLogging = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0x26, 0xfd, 0xcb, 0x14, 0x68, 0x59, 0xda, 0x34, 0x75, 0xa8, 0x9a, 0xba, 0x82, 0xa6,
	0x91, 0xea, 0x88, 0xa0, 0x4a, 0x34, 0x88, 0x4b, 0xd5, 0x1e, 0x90, 0x52, 0x54, 0xb9, 0x08, 0x24,
	0x2e, 0x96, 0x93, 0x4c, 0x2d, 0x53, 0xc7, 0x6b, 0xec, 0x4d, 0x4a, 0x40, 0x5c, 0xb8, 0x70, 0xe0,
	0xc8, 0x43, 0x70, 0x83, 0xf7, 0xe0, 0xca, 0x2b, 0xf0, 0x00, 0x1c, 0x39, 0xb2, 0xb6, 0xd7, 0x69,
	0x1a, 0x87, 0x36, 0xe2, 0x64, 0xcf, 0xec, 0xcc, 0x7c, 0xf3, 0xcd, 0x7c, 0xbb, 0xb0, 0x6e, 0x31,
	0x66, 0x39, 0x58, 0x75, 0x98, 0x65, 0xd9, 0xae, 0x55, 0xed, 0xd5, 0x92, 0x5f, 0xcd, 0xf3, 0x19,
	0x67, 0xf4, 0x76, 0x1c, 0xa0, 0x25, 0xde, 0x5e, 0x4d, 0xb9, 0x2b, 0x73, 0x4c, 0xcf, 0xae, 0x9a,
	0xae, 0xcb, 0xb8, 0xc9, 0x6d, 0xe6, 0x06, 0x71, 0x82, 0xb2, 0x39, 0x74, 0xda, 0x61, 0xae, 0xcd,
	0x99, 0x8f, 0x6d, 0xc3, 0xc7, 0x80, 0x75, 0xfd, 0x16, 0xca, 0xa0, 0x8d, 0xb1, 0xb0, 0x06, 0xba,
	0xdc, 0xef, 0xcb, 0x90, 0xa2, 0x0c, 0x89, 0xac, 0x66, 0xf7, 0xb4, 0x8a, 0x1d, 0x8f, 0x27, 0x87,
	0x2b, 0xf2, 0xd0, 0xf7, 0x5a, 0xd5, 0x40, 0xe0, 0x77, 0x25, 0xba, 0xba, 0x03, 0x8b, 0x07, 0xe8,
	0x20, 0xc7, 0x06, 0xb3, 0x74, 0x7c, 0xd3, 0xc5, 0x80, 0xd3, 0x55, 0x98, 0x0b, 0x8b, 0xbb, 0x66,
	0x07, 0x0b, 0xa4, 0x44, 0xca, 0x39, 0x7d, 0x56, 0xd8, 0xcf, 0x84, 0xa9, 0x7e, 0xcf, 0x40, 0xfe,
	0xa5, 0x6f, 0x47, 0xe1, 0x87, 0x02, 0xdc, 0xc6, 0xe0, 0xfa, 0x2c, 0xba, 0x07, 0x73, 0x09, 0x9f,
	0x42, 0x46, 0x1c, 0xcd, 0xd7, 0xd6, 0x34, 0x39, 0x26, 0xc1, 0x5a, 0x3b, 0x4a, 0x58, 0xeb, 0x32,
	0x48, 0x1f, 0x84, 0xd3, 0x23, 0x98, 0x71, 0xcc, 0x26, 0x3a, 0x41, 0x21, 0x5b, 0xca, 0x8a, 0xc4,
	0x5d, 0x2d, 0x35, 0x5f, 0x6d, 0x7c, 0x43, 0x5a, 0x23, 0xca, 0x0b, 0x9d, 0x7d, 0x5d, 0x16, 0xa1,
	0xbb, 0x30, 0x8b, 0x71, 0x54, 0x61, 0x2a, 0xaa, 0x57, 0x1c, 0x53, 0x4f, 0x96, 0xea, 0xeb, 0x49,
	0xac, 0xb2, 0x07, 0xf3, 0x43, 0xd5, 0xe8, 0x22, 0x64, 0xcf, 0xb0, 0x2f, 0x59, 0x86, 0xbf, 0x74,
	0x09, 0xa6, 0x7b, 0xa6, 0xd3, 0x8d, 0xe9, 0xe5, 0xf4, 0xd8, 0xa8, 0x67, 0x1e, 0x11, 0x75, 0x15,
	0x56, 0x52, 0xfd, 0x05, 0x9e, 0x58, 0x3f, 0xaa, 0x5f, 0x09, 0x2c, 0x37, 0xec, 0x80, 0xa7, 0x67,
	0xb9, 0x0e, 0xf3, 0x62, 0x3d, 0xaf, 0xb1, 0xc5, 0x0d, 0xbb, 0x1d, 0x08, 0xa0, 0xac, 0x28, 0x0a,
	0xd2, 0xf5, 0xb4, 0x1d, 0xd0, 0x3c, 0xcc, 0x9c, 0xda, 0x0e, 0x47, 0x5f, 0x02, 0x4a, 0x2b, 0x5c,
	0x02, 0xf3, 0xdb, 0xe8, 0x1b, 0xcd, 0xbe, 0x18, 0x58, 0xb4, 0x84, 0xc8, 0xde, 0xef, 0xd3, 0x22,
	0xe4, 0x3c, 0xd3, 0x42, 0x23, 0xb0, 0xdf, 0xa1, 0x20, 0x4f, 0xca, 0xd3, 0xfa, 0x5c, 0xe8, 0x38,
	0x11, 0x36, 0x5d, 0x03, 0x88, 0x0e, 0x39, 0x3b, 0x43, 0xb7, 0x30, 0x1d, 0x65, 0x46, 0xe1, 0xcf,
	0x43, 0x87, 0x7a, 0x0e, 0xf9, 0xd1, 0x46, 0x63, 0x0e, 0xc3, 0x03, 0x25, 0x93, 0x0f, 0x94, 0xde,
	0x87, 0x05, 0x17, 0xdf, 0x72, 0x63, 0x08, 0x34, 0x26, 0x72, 0x33, 0x74, 0x1f, 0x0f, 0x80, 0x11,
	0xb6, 0x42, 0xe0, 0x94, 0x42, 0x0e, 0x30, 0x68, 0xf9, 0xb6, 0x27, 0x7c, 0x83, 0x99, 0x5d, 0xe2,
	0x47, 0xae, 0xe4, 0x97, 0x19, 0xe5, 0xf7, 0x8d, 0x40, 0xf9, 0x7a, 0x1c, 0x49, 0xf9, 0x15, 0x2c,
	0x25, 0xf2, 0x34, 0xda, 0x17, 0xe7, 0x92, 0xff, 0xd6, 0x95, 0xca, 0xbe, 0xa8, 0xa7, 0xdf, 0xf1,
	0xd3, 0x18, 0x93, 0xce, 0xa5, 0xf6, 0x7b, 0x0a, 0x16, 0x1b, 0xf1, 0x80, 0x4f, 0xd0, 0xef, 0xd9,
	0x2d, 0x7c, 0x51, 0xa3, 0xe7, 0x90, 0x1b, 0xdc, 0x65, 0xba, 0x39, 0x66, 0x0f, 0xa3, 0x37, 0x5d,
	0xc9, 0x27, 0x41, 0xc9, 0xa3, 0xa1, 0x1d, 0x86, 0x8f, 0x86, 0xba, 0xf3, 0xf1, 0xe7, 0xaf, 0x2f,
	0x99, 0xad, 0xca, 0x3d, 0xf1, 0xd0, 0x34, 0x91, 0x9b, 0x0f, 0xaa, 0xef, 0x93, 0xbb, 0xfd, 0x44,
	0xaa, 0x30, 0xa8, 0x56, 0xc2, 0x27, 0x48, 0x7c, 0x3e, 0xd0, 0xcf, 0x04, 0x16, 0x46, 0x44, 0x4e,
	0xb7, 0x27, 0xbe, 0xa8, 0x4a, 0x65, 0x92, 0x50, 0x79, 0x67, 0x36, 0xa2, 0xce, 0x8a, 0x6a, 0x7e,
	0xd0, 0x99, 0x94, 0x54, 0xfd, 0x3c, 0xcc, 0xa8, 0x93, 0x0a, 0xfd, 0x44, 0xe0, 0xd6, 0x65, 0xb5,
	0xd2, 0xf2, 0x38, 0x51, 0x8e, 0xbb, 0x79, 0xca, 0xf6, 0x04, 0x91, 0xb2, 0x95, 0x52, 0xd4, 0x8a,
	0xa2, 0x2e, 0xa7, 0x5a, 0x71, 0x44, 0x42, 0xd8, 0xc9, 0x0f, 0x02, 0xa5, 0xeb, 0x64, 0x45, 0xeb,
	0xff, 0x40, 0x9c, 0x40, 0xf3, 0xca, 0xe3, 0xff, 0xca, 0x95, 0xfd, 0xcb, 0x25, 0xd3, 0x8b, 0x25,
	0x77, 0xae, 0x48, 0xdb, 0xaf, 0xc0, 0x72, 0x8b, 0x75, 0xd2, 0x80, 0xfb, 0x37, 0xa4, 0x10, 0x8f,
	0x43, 0x0d, 0x1d, 0x93, 0x3f, 0x84, 0x34, 0x67, 0x22, 0x3d, 0x3d, 0xfc, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x6d, 0xab, 0xb8, 0x62, 0x2e, 0x07, 0x00, 0x00,
}
