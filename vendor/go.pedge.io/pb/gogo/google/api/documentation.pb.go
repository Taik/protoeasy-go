// Code generated by protoc-gen-gogo.
// source: google/api/documentation.proto
// DO NOT EDIT!

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Documentation` provides the information for describing a service.
//
// Example:
// <pre><code>documentation:
//   summary: >
//     The Google Calendar API gives access
//     to most calendar features.
//   pages:
//   - name: Overview
//     content: &#40;== include google/foo/overview.md ==&#41;
//   - name: Tutorial
//     content: &#40;== include google/foo/tutorial.md ==&#41;
//     subpages;
//     - name: Java
//       content: &#40;== include google/foo/tutorial_java.md ==&#41;
//   rules:
//   - selector: google.calendar.Calendar.Get
//     description: >
//       ...
//   - selector: google.calendar.Calendar.Put
//     description: >
//       ...
// </code></pre>
// Documentation is provided in markdown syntax. In addition to
// standard markdown features, definition lists, tables and fenced
// code blocks are supported. Section headers can be provided and are
// interpreted relative to the section nesting of the context where
// a documentation fragment is embedded.
//
// Documentation from the IDL is merged with documentation defined
// via the config at normalization time, where documentation provided
// by config rules overrides IDL provided.
//
// A number of constructs specific to the API platform are supported
// in documentation text.
//
// In order to reference a proto element, the following
// notation can be used:
// <pre><code>&#91;fully.qualified.proto.name]&#91;]</code></pre>
// To override the display text used for the link, this can be used:
// <pre><code>&#91;display text]&#91;fully.qualified.proto.name]</code></pre>
// Text can be excluded from doc using the following notation:
// <pre><code>&#40;-- internal comment --&#41;</code></pre>
// Comments can be made conditional using a visibility label. The below
// text will be only rendered if the `BETA` label is available:
// <pre><code>&#40;--BETA: comment for BETA users --&#41;</code></pre>
// A few directives are available in documentation. Note that
// directives must appear on a single line to be properly
// identified. The `include` directive includes a markdown file from
// an external source:
// <pre><code>&#40;== include path/to/file ==&#41;</code></pre>
// The `resource_for` directive marks a message to be the resource of
// a collection in REST view. If it is not specified, tools attempt
// to infer the resource from the operations in a collection:
// <pre><code>&#40;== resource_for v1.shelves.books ==&#41;</code></pre>
// The directive `suppress_warning` does not directly affect documentation
// and is documented together with service config validation.
type Documentation struct {
	// A short summary of what the service does. Can only be provided by
	// plain text.
	Summary string `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	// The top level pages for the documentation set.
	Pages []*Page `protobuf:"bytes,5,rep,name=pages" json:"pages,omitempty"`
	// A list of documentation rules that apply to individual API elements.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*DocumentationRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
	// The URL to the root of documentation.
	DocumentationRootUrl string `protobuf:"bytes,4,opt,name=documentation_root_url,json=documentationRootUrl,proto3" json:"documentation_root_url,omitempty"`
	// Declares a single overview page. For example:
	// <pre><code>documentation:
	//   summary: ...
	//   overview: &#40;== include overview.md ==&#41;
	// </code></pre>
	// This is a shortcut for the following declaration (using pages style):
	// <pre><code>documentation:
	//   summary: ...
	//   pages:
	//   - name: Overview
	//     content: &#40;== include overview.md ==&#41;
	// </code></pre>
	// Note: you cannot specify both `overview` field and `pages` field.
	Overview string `protobuf:"bytes,2,opt,name=overview,proto3" json:"overview,omitempty"`
}

func (m *Documentation) Reset()                    { *m = Documentation{} }
func (m *Documentation) String() string            { return proto.CompactTextString(m) }
func (*Documentation) ProtoMessage()               {}
func (*Documentation) Descriptor() ([]byte, []int) { return fileDescriptorDocumentation, []int{0} }

func (m *Documentation) GetPages() []*Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

func (m *Documentation) GetRules() []*DocumentationRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A documentation rule provides information about individual API elements.
type DocumentationRule struct {
	// The selector is a comma-separated list of patterns. Each pattern is a
	// qualified name of the element which may end in "*", indicating a wildcard.
	// Wildcards are only allowed at the end and for a whole component of the
	// qualified name, i.e. "foo.*" is ok, but not "foo.b*" or "foo.*.bar". To
	// specify a default for all applicable elements, the whole pattern "*"
	// is used.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Description of the selected API(s).
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Deprecation description of the selected element(s). It can be provided if an
	// element is marked as `deprecated`.
	DeprecationDescription string `protobuf:"bytes,3,opt,name=deprecation_description,json=deprecationDescription,proto3" json:"deprecation_description,omitempty"`
}

func (m *DocumentationRule) Reset()                    { *m = DocumentationRule{} }
func (m *DocumentationRule) String() string            { return proto.CompactTextString(m) }
func (*DocumentationRule) ProtoMessage()               {}
func (*DocumentationRule) Descriptor() ([]byte, []int) { return fileDescriptorDocumentation, []int{1} }

// Represents a documentation page. A page can contain subpages to represent
// nested documentation set structure.
type Page struct {
	// The name of the page. It will be used as an identity of the page to
	// generate URI of the page, text of the link to this page in navigation,
	// etc. The full page name (start from the root page name to this page
	// concatenated with `.`) can be used as reference to the page in your
	// documentation. For example:
	// <pre><code>pages:
	// - name: Tutorial
	//   content: &#40;== include tutorial.md ==&#41;
	//   subpages:
	//   - name: Java
	//     content: &#40;== include tutorial_java.md ==&#41;
	// </code></pre>
	// You can reference `Java` page using Markdown reference link syntax:
	// `[Java][Tutorial.Java]`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Markdown content of the page. You can use <code>&#40;== include {path} ==&#41;</code>
	// to include content from a Markdown file.
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// Subpages of this page. The order of subpages specified here will be
	// honored in the generated docset.
	Subpages []*Page `protobuf:"bytes,3,rep,name=subpages" json:"subpages,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptorDocumentation, []int{2} }

func (m *Page) GetSubpages() []*Page {
	if m != nil {
		return m.Subpages
	}
	return nil
}

func init() {
	proto.RegisterType((*Documentation)(nil), "google.api.Documentation")
	proto.RegisterType((*DocumentationRule)(nil), "google.api.DocumentationRule")
	proto.RegisterType((*Page)(nil), "google.api.Page")
}

func init() { proto.RegisterFile("google/api/documentation.proto", fileDescriptorDocumentation) }

var fileDescriptorDocumentation = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x93, 0x6a, 0x9d, 0xa2, 0xe8, 0x20, 0x35, 0x08, 0x4a, 0xe9, 0x41, 0x7a, 0xd0,
	0x14, 0xac, 0xe0, 0xd9, 0x52, 0x10, 0x6f, 0x21, 0xe0, 0xb9, 0x6c, 0xd3, 0x21, 0x04, 0x92, 0x4c,
	0xd8, 0xec, 0x56, 0x7c, 0x05, 0x1f, 0xc3, 0xa7, 0xf2, 0x71, 0x24, 0x9b, 0xb4, 0xdd, 0x20, 0xde,
	0xf2, 0xe7, 0xff, 0xb2, 0xff, 0xcc, 0x9f, 0x85, 0xdb, 0x84, 0x39, 0xc9, 0x68, 0x26, 0xca, 0x74,
	0xb6, 0xe1, 0x58, 0xe7, 0x54, 0x28, 0xa1, 0x52, 0x2e, 0x82, 0x52, 0xb2, 0x62, 0x84, 0xc6, 0x0f,
	0x44, 0x99, 0x4e, 0x7e, 0x1c, 0x38, 0x5d, 0xda, 0x0c, 0xfa, 0x70, 0x5c, 0xe9, 0x3c, 0x17, 0xf2,
	0xd3, 0x77, 0xc6, 0xce, 0xf4, 0x24, 0xda, 0x49, 0xbc, 0x83, 0x7e, 0x29, 0x12, 0xaa, 0xfc, 0xfe,
	0xd8, 0x9d, 0x0e, 0x1f, 0xcf, 0x83, 0xc3, 0x39, 0x41, 0x28, 0x12, 0x8a, 0x1a, 0x1b, 0xe7, 0xd0,
	0x97, 0x3a, 0xa3, 0xca, 0x77, 0x0d, 0x77, 0x63, 0x73, 0x9d, 0xac, 0x48, 0x67, 0x14, 0x35, 0x2c,
	0x3e, 0xc1, 0xa8, 0x33, 0xeb, 0x4a, 0x32, 0xab, 0x95, 0x96, 0x99, 0xef, 0x99, 0x29, 0x2e, 0x3b,
	0x6e, 0xc4, 0xac, 0xde, 0x65, 0x86, 0xd7, 0x30, 0xe0, 0x2d, 0xc9, 0x6d, 0x4a, 0x1f, 0x7e, 0xcf,
	0x70, 0x7b, 0x3d, 0xf9, 0x72, 0xe0, 0xe2, 0x4f, 0x5c, 0xfd, 0x45, 0x45, 0x19, 0xc5, 0x8a, 0x65,
	0xbb, 0xdf, 0x5e, 0xe3, 0x18, 0x86, 0x1b, 0xaa, 0x62, 0x99, 0x96, 0x35, 0xde, 0x1e, 0x68, 0xbf,
	0xc2, 0x67, 0xb8, 0xda, 0x50, 0x29, 0x29, 0x6e, 0x66, 0xb4, 0x69, 0xd7, 0xd0, 0x23, 0xcb, 0x5e,
	0x1e, 0xdc, 0xc9, 0x1a, 0xbc, 0xba, 0x22, 0x44, 0xf0, 0x0a, 0x91, 0x53, 0x1b, 0x6d, 0x9e, 0xeb,
	0xc6, 0x63, 0x2e, 0x14, 0x15, 0xaa, 0x8d, 0xdc, 0x49, 0xbc, 0x87, 0x41, 0xa5, 0xd7, 0x4d, 0xe9,
	0xee, 0x3f, 0xa5, 0xef, 0x89, 0xc5, 0x03, 0x9c, 0xc5, 0x9c, 0x5b, 0xc0, 0x02, 0x3b, 0xfb, 0x87,
	0xf5, 0xdf, 0x0f, 0x9d, 0xef, 0x9e, 0xf7, 0xfa, 0x12, 0xbe, 0xad, 0x8f, 0xcc, 0x6d, 0x98, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xcc, 0x01, 0x2d, 0x2f, 0x02, 0x00, 0x00,
}
