// Code generated by protoc-gen-go.
// source: page.proto
// DO NOT EDIT!

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContentType int32

const (
	ContentType_unknown  ContentType = 0
	ContentType_html     ContentType = 1
	ContentType_markdown ContentType = 2
)

var ContentType_name = map[int32]string{
	0: "unknown",
	1: "html",
	2: "markdown",
}
var ContentType_value = map[string]int32{
	"unknown":  0,
	"html":     1,
	"markdown": 2,
}

func (x ContentType) Enum() *ContentType {
	p := new(ContentType)
	*p = x
	return p
}
func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}
func (x *ContentType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ContentType_value, data, "ContentType")
	if err != nil {
		return err
	}
	*x = ContentType(value)
	return nil
}
func (ContentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type PageTemplate struct {
	Uuid             *string    `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name             *string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Contents         []*Content `protobuf:"bytes,10,rep,name=contents" json:"contents,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *PageTemplate) Reset()                    { *m = PageTemplate{} }
func (m *PageTemplate) String() string            { return proto.CompactTextString(m) }
func (*PageTemplate) ProtoMessage()               {}
func (*PageTemplate) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *PageTemplate) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *PageTemplate) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PageTemplate) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

type Page struct {
	Uuid             *string           `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name             *string           `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Template         *string           `protobuf:"bytes,3,opt,name=template" json:"template,omitempty"`
	TemplateEngine   *string           `protobuf:"bytes,4,opt,name=templateEngine,def=html" json:"templateEngine,omitempty"`
	Contents         []*Content        `protobuf:"bytes,10,rep,name=contents" json:"contents,omitempty"`
	Metadata         map[string]string `protobuf:"bytes,11,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tags             []string          `protobuf:"bytes,12,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

const Default_Page_TemplateEngine string = "html"

func (m *Page) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *Page) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Page) GetTemplate() string {
	if m != nil && m.Template != nil {
		return *m.Template
	}
	return ""
}

func (m *Page) GetTemplateEngine() string {
	if m != nil && m.TemplateEngine != nil {
		return *m.TemplateEngine
	}
	return Default_Page_TemplateEngine
}

func (m *Page) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Page) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Page) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// versioned?
type Content struct {
	Uuid             *string      `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	ContentType      *ContentType `protobuf:"varint,2,opt,name=contentType,enum=press.models.ContentType" json:"contentType,omitempty"`
	Key              *string      `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value            *string      `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Content) Reset()                    { *m = Content{} }
func (m *Content) String() string            { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()               {}
func (*Content) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Content) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *Content) GetContentType() ContentType {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ContentType_unknown
}

func (m *Content) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Content) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*PageTemplate)(nil), "press.models.PageTemplate")
	proto.RegisterType((*Page)(nil), "press.models.Page")
	proto.RegisterType((*Content)(nil), "press.models.Content")
	proto.RegisterEnum("press.models.ContentType", ContentType_name, ContentType_value)
}

var fileDescriptor2 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x50, 0x4d, 0x4b, 0x84, 0x50,
	0x14, 0xcd, 0x8f, 0x18, 0xbb, 0xda, 0x20, 0x97, 0x82, 0xd7, 0xac, 0xc4, 0xd5, 0x10, 0x21, 0xe4,
	0x2a, 0x66, 0xda, 0xc5, 0x2c, 0x83, 0x90, 0xf9, 0x03, 0x8f, 0x7c, 0x98, 0xa8, 0x4f, 0xd1, 0x67,
	0xe1, 0xb6, 0x9f, 0xd5, 0xaf, 0xcb, 0xe7, 0x17, 0x4e, 0xb8, 0x88, 0xd9, 0x9d, 0x77, 0xef, 0x39,
	0xef, 0x9e, 0x73, 0x00, 0x0a, 0x1a, 0x31, 0xaf, 0x28, 0x73, 0x91, 0xa3, 0x55, 0x94, 0xac, 0xaa,
	0xbc, 0x2c, 0x0f, 0x59, 0x5a, 0xb9, 0x31, 0x58, 0x6f, 0xed, 0xee, 0xc8, 0xb2, 0x22, 0xa5, 0x82,
	0x21, 0x82, 0x5e, 0xd7, 0x71, 0x48, 0x14, 0x47, 0xd9, 0x5e, 0x05, 0x1d, 0x96, 0x33, 0x4e, 0x33,
	0x46, 0xd4, 0x7e, 0x26, 0x31, 0x3e, 0x82, 0xf1, 0x9e, 0x73, 0xc1, 0xb8, 0xa8, 0x08, 0x38, 0xda,
	0xd6, 0xf4, 0x6f, 0xbd, 0xf9, 0xc7, 0xde, 0x4b, 0xbf, 0x0d, 0x26, 0x9a, 0xfb, 0xa3, 0x82, 0x2e,
	0x6f, 0xfd, 0xfb, 0xc6, 0x06, 0x0c, 0x31, 0xf8, 0x22, 0x5a, 0x37, 0x9f, 0xde, 0xf8, 0x00, 0xeb,
	0x11, 0x1f, 0x78, 0x14, 0x73, 0x46, 0x74, 0xc9, 0xd8, 0xe9, 0x1f, 0x22, 0x4b, 0x83, 0x3f, 0xbb,
	0x33, 0xdc, 0xe2, 0x33, 0x18, 0x19, 0x13, 0x34, 0xa4, 0x82, 0x12, 0xb3, 0x93, 0x38, 0xa7, 0x12,
	0x19, 0xc5, 0x7b, 0x1d, 0x28, 0x07, 0x2e, 0xca, 0x26, 0x98, 0x14, 0x32, 0x8e, 0xa0, 0x51, 0x45,
	0xac, 0x56, 0xd9, 0xc6, 0x91, 0x78, 0xb3, 0x87, 0xeb, 0x13, 0x3a, 0xda, 0xa0, 0x25, 0xac, 0x19,
	0x6a, 0x90, 0x10, 0x6f, 0xe0, 0xf2, 0x93, 0xa6, 0xf5, 0x58, 0x43, 0xff, 0xd8, 0xa9, 0x4f, 0x8a,
	0xfb, 0xad, 0xc0, 0x6a, 0x30, 0xb9, 0xd8, 0xdf, 0x1e, 0xcc, 0xc1, 0xfa, 0xb1, 0x29, 0x7a, 0xfd,
	0xda, 0xbf, 0x5b, 0x0c, 0x29, 0x09, 0xc1, 0x9c, 0x3d, 0x1a, 0xd1, 0x16, 0x8c, 0xe8, 0x33, 0x23,
	0xf7, 0x3e, 0x98, 0xb3, 0x3f, 0xd0, 0x84, 0x55, 0xcd, 0x13, 0x9e, 0x7f, 0x71, 0xfb, 0x02, 0x0d,
	0xe8, 0xaa, 0xb7, 0x15, 0xb4, 0xda, 0xe6, 0x68, 0x99, 0x84, 0x72, 0xae, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x20, 0xf6, 0x19, 0xcc, 0x7b, 0x02, 0x00, 0x00,
}
