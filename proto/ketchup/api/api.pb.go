// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto
	content.proto
	data.proto
	export/export.proto
	metadata.proto
	page.proto
	route.proto
	theme.proto
	user.proto
	packages.proto
	struct.proto

It has these top-level messages:
	Error
	FieldError
	ListOptions
	UpdatePageRoutesRequest
	GetRenderedPageRequest
	GetRenderedPageResponse
	ListPageRequest
	ListPageResponse
	ListRouteRequest
	ListRouteResponse
	ListThemeResponse
	UpdateDataRequest
	ListDataResponse
	TLSSettingsReponse
	EnableTLSRequest
	InstallThemeRequest
	ContentMultiple
	ContentText
	ContentString
	Data
	Export
	Metadata
	Timestamp
	Author
	Page
	Content
	Route
	Theme
	ThemeTemplate
	ThemePlaceholder
	ThemeAsset
	User
	PackageAuthor
	Package
	Registry
	Struct
	Value
	ListValue
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ketchup_models "github.com/ketchuphq/ketchup/proto/ketchup/models"
import structpb "github.com/ketchuphq/ketchup/proto/structpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_INTERNAL_SERVER_ERROR ErrorCode = 1
	ErrorCode_NOT_FOUND             ErrorCode = 2
	ErrorCode_FORBIDDEN             ErrorCode = 3
)

var ErrorCode_name = map[int32]string{
	1: "INTERNAL_SERVER_ERROR",
	2: "NOT_FOUND",
	3: "FORBIDDEN",
}
var ErrorCode_value = map[string]int32{
	"INTERNAL_SERVER_ERROR": 1,
	"NOT_FOUND":             2,
	"FORBIDDEN":             3,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}
func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetRenderedPageRequest_RenderedPageFilter int32

const (
	GetRenderedPageRequest_all       GetRenderedPageRequest_RenderedPageFilter = 1
	GetRenderedPageRequest_published GetRenderedPageRequest_RenderedPageFilter = 2
	GetRenderedPageRequest_draft     GetRenderedPageRequest_RenderedPageFilter = 3
)

var GetRenderedPageRequest_RenderedPageFilter_name = map[int32]string{
	1: "all",
	2: "published",
	3: "draft",
}
var GetRenderedPageRequest_RenderedPageFilter_value = map[string]int32{
	"all":       1,
	"published": 2,
	"draft":     3,
}

func (x GetRenderedPageRequest_RenderedPageFilter) Enum() *GetRenderedPageRequest_RenderedPageFilter {
	p := new(GetRenderedPageRequest_RenderedPageFilter)
	*p = x
	return p
}
func (x GetRenderedPageRequest_RenderedPageFilter) String() string {
	return proto.EnumName(GetRenderedPageRequest_RenderedPageFilter_name, int32(x))
}
func (x *GetRenderedPageRequest_RenderedPageFilter) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GetRenderedPageRequest_RenderedPageFilter_value, data, "GetRenderedPageRequest_RenderedPageFilter")
	if err != nil {
		return err
	}
	*x = GetRenderedPageRequest_RenderedPageFilter(value)
	return nil
}
func (GetRenderedPageRequest_RenderedPageFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type ListPageRequest_ListPageFilter int32

const (
	ListPageRequest_all       ListPageRequest_ListPageFilter = 1
	ListPageRequest_published ListPageRequest_ListPageFilter = 2
	ListPageRequest_draft     ListPageRequest_ListPageFilter = 3
)

var ListPageRequest_ListPageFilter_name = map[int32]string{
	1: "all",
	2: "published",
	3: "draft",
}
var ListPageRequest_ListPageFilter_value = map[string]int32{
	"all":       1,
	"published": 2,
	"draft":     3,
}

func (x ListPageRequest_ListPageFilter) Enum() *ListPageRequest_ListPageFilter {
	p := new(ListPageRequest_ListPageFilter)
	*p = x
	return p
}
func (x ListPageRequest_ListPageFilter) String() string {
	return proto.EnumName(ListPageRequest_ListPageFilter_name, int32(x))
}
func (x *ListPageRequest_ListPageFilter) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ListPageRequest_ListPageFilter_value, data, "ListPageRequest_ListPageFilter")
	if err != nil {
		return err
	}
	*x = ListPageRequest_ListPageFilter(value)
	return nil
}
func (ListPageRequest_ListPageFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type Error struct {
	Code             *ErrorCode    `protobuf:"varint,1,opt,name=code,enum=ketchup.api.ErrorCode" json:"code,omitempty"`
	Title            *string       `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Detail           *string       `protobuf:"bytes,3,opt,name=detail" json:"detail,omitempty"`
	Fields           []*FieldError `protobuf:"bytes,4,rep,name=fields" json:"fields,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetCode() ErrorCode {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ErrorCode_INTERNAL_SERVER_ERROR
}

func (m *Error) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Error) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

func (m *Error) GetFields() []*FieldError {
	if m != nil {
		return m.Fields
	}
	return nil
}

type FieldError struct {
	Field            *string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	Code             *string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	Title            *string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Detail           *string `protobuf:"bytes,4,opt,name=detail" json:"detail,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FieldError) Reset()                    { *m = FieldError{} }
func (m *FieldError) String() string            { return proto.CompactTextString(m) }
func (*FieldError) ProtoMessage()               {}
func (*FieldError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FieldError) GetField() string {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return ""
}

func (m *FieldError) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *FieldError) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *FieldError) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

type ListOptions struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListOptions) Reset()                    { *m = ListOptions{} }
func (m *ListOptions) String() string            { return proto.CompactTextString(m) }
func (*ListOptions) ProtoMessage()               {}
func (*ListOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type UpdatePageRoutesRequest struct {
	Routes           []*ketchup_models.Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdatePageRoutesRequest) Reset()                    { *m = UpdatePageRoutesRequest{} }
func (m *UpdatePageRoutesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePageRoutesRequest) ProtoMessage()               {}
func (*UpdatePageRoutesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdatePageRoutesRequest) GetRoutes() []*ketchup_models.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type GetRenderedPageRequest struct {
	Options          *GetRenderedPageRequest_RenderedPageOptions `protobuf:"bytes,1,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *GetRenderedPageRequest) Reset()                    { *m = GetRenderedPageRequest{} }
func (m *GetRenderedPageRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRenderedPageRequest) ProtoMessage()               {}
func (*GetRenderedPageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetRenderedPageRequest) GetOptions() *GetRenderedPageRequest_RenderedPageOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetRenderedPageRequest_RenderedPageOptions struct {
	Filter           *GetRenderedPageRequest_RenderedPageFilter `protobuf:"varint,1,opt,name=filter,enum=ketchup.api.GetRenderedPageRequest_RenderedPageFilter" json:"filter,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *GetRenderedPageRequest_RenderedPageOptions) Reset() {
	*m = GetRenderedPageRequest_RenderedPageOptions{}
}
func (m *GetRenderedPageRequest_RenderedPageOptions) String() string {
	return proto.CompactTextString(m)
}
func (*GetRenderedPageRequest_RenderedPageOptions) ProtoMessage() {}
func (*GetRenderedPageRequest_RenderedPageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

func (m *GetRenderedPageRequest_RenderedPageOptions) GetFilter() GetRenderedPageRequest_RenderedPageFilter {
	if m != nil && m.Filter != nil {
		return *m.Filter
	}
	return GetRenderedPageRequest_all
}

type GetRenderedPageResponse struct {
	Data             *structpb.Struct `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GetRenderedPageResponse) Reset()                    { *m = GetRenderedPageResponse{} }
func (m *GetRenderedPageResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRenderedPageResponse) ProtoMessage()               {}
func (*GetRenderedPageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetRenderedPageResponse) GetData() *structpb.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListPageRequest struct {
	List             *ListOptions                     `protobuf:"bytes,1,opt,name=list" json:"list,omitempty"`
	Options          *ListPageRequest_ListPageOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *ListPageRequest) Reset()                    { *m = ListPageRequest{} }
func (m *ListPageRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPageRequest) ProtoMessage()               {}
func (*ListPageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListPageRequest) GetList() *ListOptions {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ListPageRequest) GetOptions() *ListPageRequest_ListPageOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ListPageRequest_ListPageOptions struct {
	Filter           *ListPageRequest_ListPageFilter `protobuf:"varint,1,opt,name=filter,enum=ketchup.api.ListPageRequest_ListPageFilter" json:"filter,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *ListPageRequest_ListPageOptions) Reset()         { *m = ListPageRequest_ListPageOptions{} }
func (m *ListPageRequest_ListPageOptions) String() string { return proto.CompactTextString(m) }
func (*ListPageRequest_ListPageOptions) ProtoMessage()    {}
func (*ListPageRequest_ListPageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

func (m *ListPageRequest_ListPageOptions) GetFilter() ListPageRequest_ListPageFilter {
	if m != nil && m.Filter != nil {
		return *m.Filter
	}
	return ListPageRequest_all
}

type ListPageResponse struct {
	Pages            []*ketchup_models.Page `protobuf:"bytes,1,rep,name=pages" json:"pages,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ListPageResponse) Reset()                    { *m = ListPageResponse{} }
func (m *ListPageResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPageResponse) ProtoMessage()               {}
func (*ListPageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListPageResponse) GetPages() []*ketchup_models.Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

type ListRouteRequest struct {
	List             *ListOptions                       `protobuf:"bytes,1,opt,name=list" json:"list,omitempty"`
	Options          *ListRouteRequest_ListRouteOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *ListRouteRequest) Reset()                    { *m = ListRouteRequest{} }
func (m *ListRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRouteRequest) ProtoMessage()               {}
func (*ListRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListRouteRequest) GetList() *ListOptions {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ListRouteRequest) GetOptions() *ListRouteRequest_ListRouteOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ListRouteRequest_ListRouteOptions struct {
	PageUuid         *string `protobuf:"bytes,1,opt,name=page_uuid,json=pageUuid" json:"page_uuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ListRouteRequest_ListRouteOptions) Reset()         { *m = ListRouteRequest_ListRouteOptions{} }
func (m *ListRouteRequest_ListRouteOptions) String() string { return proto.CompactTextString(m) }
func (*ListRouteRequest_ListRouteOptions) ProtoMessage()    {}
func (*ListRouteRequest_ListRouteOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

func (m *ListRouteRequest_ListRouteOptions) GetPageUuid() string {
	if m != nil && m.PageUuid != nil {
		return *m.PageUuid
	}
	return ""
}

type ListRouteResponse struct {
	Routes           []*ketchup_models.Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ListRouteResponse) Reset()                    { *m = ListRouteResponse{} }
func (m *ListRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*ListRouteResponse) ProtoMessage()               {}
func (*ListRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListRouteResponse) GetRoutes() []*ketchup_models.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type ListThemeResponse struct {
	Themes           []*ketchup_models.Theme `protobuf:"bytes,1,rep,name=themes" json:"themes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ListThemeResponse) Reset()                    { *m = ListThemeResponse{} }
func (m *ListThemeResponse) String() string            { return proto.CompactTextString(m) }
func (*ListThemeResponse) ProtoMessage()               {}
func (*ListThemeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListThemeResponse) GetThemes() []*ketchup_models.Theme {
	if m != nil {
		return m.Themes
	}
	return nil
}

type UpdateDataRequest struct {
	Data             []*ketchup_models.Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *UpdateDataRequest) Reset()                    { *m = UpdateDataRequest{} }
func (m *UpdateDataRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDataRequest) ProtoMessage()               {}
func (*UpdateDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateDataRequest) GetData() []*ketchup_models.Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListDataResponse struct {
	Data             []*ketchup_models.Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ListDataResponse) Reset()                    { *m = ListDataResponse{} }
func (m *ListDataResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDataResponse) ProtoMessage()               {}
func (*ListDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ListDataResponse) GetData() []*ketchup_models.Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type TLSSettingsReponse struct {
	TlsEmail         *string `protobuf:"bytes,1,opt,name=tls_email,json=tlsEmail" json:"tls_email,omitempty"`
	TlsDomain        *string `protobuf:"bytes,2,opt,name=tls_domain,json=tlsDomain" json:"tls_domain,omitempty"`
	AgreedOn         *string `protobuf:"bytes,3,opt,name=agreed_on,json=agreedOn" json:"agreed_on,omitempty"`
	TermsOfService   *string `protobuf:"bytes,4,opt,name=terms_of_service,json=termsOfService" json:"terms_of_service,omitempty"`
	HasCertificate   *bool   `protobuf:"varint,5,opt,name=has_certificate,json=hasCertificate" json:"has_certificate,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TLSSettingsReponse) Reset()                    { *m = TLSSettingsReponse{} }
func (m *TLSSettingsReponse) String() string            { return proto.CompactTextString(m) }
func (*TLSSettingsReponse) ProtoMessage()               {}
func (*TLSSettingsReponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TLSSettingsReponse) GetTlsEmail() string {
	if m != nil && m.TlsEmail != nil {
		return *m.TlsEmail
	}
	return ""
}

func (m *TLSSettingsReponse) GetTlsDomain() string {
	if m != nil && m.TlsDomain != nil {
		return *m.TlsDomain
	}
	return ""
}

func (m *TLSSettingsReponse) GetAgreedOn() string {
	if m != nil && m.AgreedOn != nil {
		return *m.AgreedOn
	}
	return ""
}

func (m *TLSSettingsReponse) GetTermsOfService() string {
	if m != nil && m.TermsOfService != nil {
		return *m.TermsOfService
	}
	return ""
}

func (m *TLSSettingsReponse) GetHasCertificate() bool {
	if m != nil && m.HasCertificate != nil {
		return *m.HasCertificate
	}
	return false
}

type EnableTLSRequest struct {
	TlsEmail         *string `protobuf:"bytes,1,opt,name=tls_email,json=tlsEmail" json:"tls_email,omitempty"`
	TlsDomain        *string `protobuf:"bytes,2,opt,name=tls_domain,json=tlsDomain" json:"tls_domain,omitempty"`
	Agreed           *bool   `protobuf:"varint,3,opt,name=agreed" json:"agreed,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EnableTLSRequest) Reset()                    { *m = EnableTLSRequest{} }
func (m *EnableTLSRequest) String() string            { return proto.CompactTextString(m) }
func (*EnableTLSRequest) ProtoMessage()               {}
func (*EnableTLSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *EnableTLSRequest) GetTlsEmail() string {
	if m != nil && m.TlsEmail != nil {
		return *m.TlsEmail
	}
	return ""
}

func (m *EnableTLSRequest) GetTlsDomain() string {
	if m != nil && m.TlsDomain != nil {
		return *m.TlsDomain
	}
	return ""
}

func (m *EnableTLSRequest) GetAgreed() bool {
	if m != nil && m.Agreed != nil {
		return *m.Agreed
	}
	return false
}

type InstallThemeRequest struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	VcsUrl           *string `protobuf:"bytes,2,opt,name=vcs_url,json=vcsUrl" json:"vcs_url,omitempty"`
	RegistryUrl      *string `protobuf:"bytes,3,opt,name=registry_url,json=registryUrl" json:"registry_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InstallThemeRequest) Reset()                    { *m = InstallThemeRequest{} }
func (m *InstallThemeRequest) String() string            { return proto.CompactTextString(m) }
func (*InstallThemeRequest) ProtoMessage()               {}
func (*InstallThemeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *InstallThemeRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *InstallThemeRequest) GetVcsUrl() string {
	if m != nil && m.VcsUrl != nil {
		return *m.VcsUrl
	}
	return ""
}

func (m *InstallThemeRequest) GetRegistryUrl() string {
	if m != nil && m.RegistryUrl != nil {
		return *m.RegistryUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "ketchup.api.Error")
	proto.RegisterType((*FieldError)(nil), "ketchup.api.FieldError")
	proto.RegisterType((*ListOptions)(nil), "ketchup.api.ListOptions")
	proto.RegisterType((*UpdatePageRoutesRequest)(nil), "ketchup.api.UpdatePageRoutesRequest")
	proto.RegisterType((*GetRenderedPageRequest)(nil), "ketchup.api.GetRenderedPageRequest")
	proto.RegisterType((*GetRenderedPageRequest_RenderedPageOptions)(nil), "ketchup.api.GetRenderedPageRequest.RenderedPageOptions")
	proto.RegisterType((*GetRenderedPageResponse)(nil), "ketchup.api.GetRenderedPageResponse")
	proto.RegisterType((*ListPageRequest)(nil), "ketchup.api.ListPageRequest")
	proto.RegisterType((*ListPageRequest_ListPageOptions)(nil), "ketchup.api.ListPageRequest.ListPageOptions")
	proto.RegisterType((*ListPageResponse)(nil), "ketchup.api.ListPageResponse")
	proto.RegisterType((*ListRouteRequest)(nil), "ketchup.api.ListRouteRequest")
	proto.RegisterType((*ListRouteRequest_ListRouteOptions)(nil), "ketchup.api.ListRouteRequest.ListRouteOptions")
	proto.RegisterType((*ListRouteResponse)(nil), "ketchup.api.ListRouteResponse")
	proto.RegisterType((*ListThemeResponse)(nil), "ketchup.api.ListThemeResponse")
	proto.RegisterType((*UpdateDataRequest)(nil), "ketchup.api.UpdateDataRequest")
	proto.RegisterType((*ListDataResponse)(nil), "ketchup.api.ListDataResponse")
	proto.RegisterType((*TLSSettingsReponse)(nil), "ketchup.api.TLSSettingsReponse")
	proto.RegisterType((*EnableTLSRequest)(nil), "ketchup.api.EnableTLSRequest")
	proto.RegisterType((*InstallThemeRequest)(nil), "ketchup.api.InstallThemeRequest")
	proto.RegisterEnum("ketchup.api.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("ketchup.api.GetRenderedPageRequest_RenderedPageFilter", GetRenderedPageRequest_RenderedPageFilter_name, GetRenderedPageRequest_RenderedPageFilter_value)
	proto.RegisterEnum("ketchup.api.ListPageRequest_ListPageFilter", ListPageRequest_ListPageFilter_name, ListPageRequest_ListPageFilter_value)
}
func (m *Error) SetCode(v *ErrorCode) {
	m.Code = v
}

func (m *Error) SetTitle(v *string) {
	m.Title = v
}

func (m *Error) SetDetail(v *string) {
	m.Detail = v
}

func (m *Error) SetFields(v []*FieldError) {
	m.Fields = v
}

func (m *FieldError) SetField(v *string) {
	m.Field = v
}

func (m *FieldError) SetCode(v *string) {
	m.Code = v
}

func (m *FieldError) SetTitle(v *string) {
	m.Title = v
}

func (m *FieldError) SetDetail(v *string) {
	m.Detail = v
}

func (m *UpdatePageRoutesRequest) SetRoutes(v []*ketchup_models.Route) {
	m.Routes = v
}

func (m *GetRenderedPageRequest) SetOptions(v *GetRenderedPageRequest_RenderedPageOptions) {
	m.Options = v
}

func (m *GetRenderedPageResponse) SetData(v *structpb.Struct) {
	m.Data = v
}

func (m *ListPageRequest) SetList(v *ListOptions) {
	m.List = v
}

func (m *ListPageRequest) SetOptions(v *ListPageRequest_ListPageOptions) {
	m.Options = v
}

func (m *ListPageResponse) SetPages(v []*ketchup_models.Page) {
	m.Pages = v
}

func (m *ListRouteRequest) SetList(v *ListOptions) {
	m.List = v
}

func (m *ListRouteRequest) SetOptions(v *ListRouteRequest_ListRouteOptions) {
	m.Options = v
}

func (m *ListRouteResponse) SetRoutes(v []*ketchup_models.Route) {
	m.Routes = v
}

func (m *ListThemeResponse) SetThemes(v []*ketchup_models.Theme) {
	m.Themes = v
}

func (m *UpdateDataRequest) SetData(v []*ketchup_models.Data) {
	m.Data = v
}

func (m *ListDataResponse) SetData(v []*ketchup_models.Data) {
	m.Data = v
}

func (m *TLSSettingsReponse) SetTlsEmail(v *string) {
	m.TlsEmail = v
}

func (m *TLSSettingsReponse) SetTlsDomain(v *string) {
	m.TlsDomain = v
}

func (m *TLSSettingsReponse) SetAgreedOn(v *string) {
	m.AgreedOn = v
}

func (m *TLSSettingsReponse) SetTermsOfService(v *string) {
	m.TermsOfService = v
}

func (m *TLSSettingsReponse) SetHasCertificate(v *bool) {
	m.HasCertificate = v
}

func (m *EnableTLSRequest) SetTlsEmail(v *string) {
	m.TlsEmail = v
}

func (m *EnableTLSRequest) SetTlsDomain(v *string) {
	m.TlsDomain = v
}

func (m *EnableTLSRequest) SetAgreed(v *bool) {
	m.Agreed = v
}

func (m *InstallThemeRequest) SetName(v *string) {
	m.Name = v
}

func (m *InstallThemeRequest) SetVcsUrl(v *string) {
	m.VcsUrl = v
}

func (m *InstallThemeRequest) SetRegistryUrl(v *string) {
	m.RegistryUrl = v
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 832 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdf, 0x6f, 0xdb, 0x36,
	0x10, 0x86, 0xfc, 0x2b, 0xf1, 0xb9, 0x4d, 0x5d, 0xb6, 0x8d, 0xbd, 0x14, 0x03, 0x32, 0xbd, 0xcc,
	0x48, 0x3b, 0x15, 0xc8, 0x80, 0xf5, 0x65, 0x1b, 0xb0, 0xc6, 0xf6, 0x1a, 0x20, 0xb0, 0x37, 0xda,
	0xee, 0xab, 0xc0, 0x48, 0x67, 0x9b, 0x18, 0x2d, 0x69, 0x24, 0x15, 0x60, 0x7f, 0xc3, 0xde, 0x86,
	0xfd, 0x37, 0x7b, 0xd9, 0x9f, 0x36, 0x88, 0xa4, 0x62, 0xc5, 0xce, 0x80, 0x64, 0x7b, 0xf3, 0x1d,
	0xbf, 0xfb, 0xee, 0xe3, 0xc7, 0xd3, 0x19, 0xda, 0x2c, 0xe3, 0x41, 0x26, 0x53, 0x9d, 0x92, 0xce,
	0x2f, 0xa8, 0xa3, 0x75, 0x9e, 0x05, 0x2c, 0xe3, 0x27, 0x90, 0xb1, 0x15, 0xda, 0x83, 0x93, 0x8e,
	0x4c, 0x73, 0x7d, 0x1b, 0xe8, 0x35, 0x6e, 0xca, 0xe0, 0x89, 0xd2, 0x32, 0x8f, 0xb4, 0x8b, 0x20,
	0x66, 0x9a, 0xd9, 0xdf, 0xfe, 0x1f, 0x1e, 0x34, 0x47, 0x52, 0xa6, 0x92, 0x9c, 0x41, 0x23, 0x4a,
	0x63, 0xec, 0x7b, 0xa7, 0xde, 0xe0, 0xe8, 0xfc, 0x38, 0xa8, 0x74, 0x09, 0x0c, 0xe2, 0x22, 0x8d,
	0x91, 0x1a, 0x0c, 0x79, 0x09, 0x4d, 0xcd, 0xb5, 0xc0, 0x7e, 0xed, 0xd4, 0x1b, 0xb4, 0xa9, 0x0d,
	0xc8, 0x31, 0xb4, 0x62, 0xd4, 0x8c, 0x8b, 0x7e, 0xdd, 0xa4, 0x5d, 0x44, 0xde, 0x41, 0x6b, 0xc9,
	0x51, 0xc4, 0xaa, 0xdf, 0x38, 0xad, 0x0f, 0x3a, 0xe7, 0xbd, 0x3b, 0xdc, 0xe3, 0xe2, 0xc8, 0x34,
	0xa0, 0x0e, 0xe6, 0xc7, 0x00, 0xdb, 0x6c, 0xd1, 0xcc, 0xe4, 0x8d, 0xb2, 0x36, 0xb5, 0x01, 0x21,
	0x4e, 0xae, 0x55, 0xb0, 0x23, 0xab, 0x7e, 0xbf, 0xac, 0x46, 0x55, 0x96, 0xff, 0x14, 0x3a, 0x57,
	0x5c, 0xe9, 0x69, 0xa6, 0x79, 0x9a, 0x28, 0xff, 0x23, 0xf4, 0x16, 0x59, 0xcc, 0x34, 0xfe, 0xc4,
	0x56, 0x48, 0x0b, 0x27, 0x15, 0xc5, 0x5f, 0x73, 0x54, 0x9a, 0x7c, 0x05, 0x2d, 0x63, 0xad, 0xea,
	0x7b, 0xe6, 0x02, 0xaf, 0x6e, 0x2f, 0xb0, 0x49, 0x63, 0x14, 0x2a, 0x30, 0x70, 0xea, 0x40, 0xfe,
	0xef, 0x35, 0x38, 0xfe, 0x11, 0x35, 0xc5, 0x24, 0x46, 0x89, 0xb1, 0xe1, 0x73, 0x4c, 0x3f, 0xc3,
	0x41, 0x6a, 0xfb, 0x99, 0xdb, 0x74, 0xce, 0xdf, 0xdf, 0xf1, 0xe2, 0xfe, 0xaa, 0xa0, 0x9a, 0x73,
	0x72, 0x69, 0xc9, 0x73, 0x82, 0xf0, 0xe2, 0x9e, 0x73, 0x32, 0x29, 0x4c, 0x17, 0x1a, 0xa5, 0x7b,
	0xd0, 0x6f, 0x1e, 0xdb, 0x68, 0x6c, 0xaa, 0xa9, 0x63, 0xf1, 0xdf, 0x03, 0xd9, 0x3f, 0x25, 0x07,
	0x50, 0x67, 0x42, 0x74, 0x3d, 0xf2, 0x14, 0xda, 0x59, 0x7e, 0x2d, 0xb8, 0x5a, 0x63, 0xdc, 0xad,
	0x91, 0x36, 0x34, 0x63, 0xc9, 0x96, 0xba, 0x5b, 0xf7, 0xc7, 0xd0, 0xdb, 0xeb, 0xa6, 0xb2, 0x34,
	0x51, 0x48, 0xde, 0x40, 0xa3, 0x18, 0x45, 0x67, 0x45, 0x2f, 0x58, 0xa5, 0xe9, 0x4a, 0xb8, 0x99,
	0xbd, 0xce, 0x97, 0xc1, 0xcc, 0x4c, 0x2d, 0x35, 0x20, 0xff, 0xcf, 0x1a, 0x3c, 0x2b, 0xde, 0xab,
	0x6a, 0xe7, 0x5b, 0x68, 0x08, 0xae, 0xb4, 0x23, 0xe8, 0xdf, 0xb9, 0x62, 0xe5, 0x6d, 0xa9, 0x41,
	0x91, 0xf1, 0xd6, 0xfc, 0x9a, 0x29, 0x78, 0xbb, 0x57, 0x50, 0x35, 0xa3, 0x8c, 0xf7, 0x1c, 0xff,
	0xb4, 0x15, 0x52, 0xba, 0x7d, 0xb1, 0xe3, 0xf6, 0x9b, 0x07, 0x31, 0xef, 0x58, 0xfc, 0x35, 0x1c,
	0xdd, 0x3d, 0x79, 0x88, 0xbd, 0xdf, 0x43, 0x77, 0x4b, 0xef, 0x7c, 0x3d, 0x83, 0x66, 0xb1, 0x16,
	0xca, 0x71, 0x7d, 0xb9, 0x3b, 0xae, 0x06, 0x6c, 0x21, 0xfe, 0x5f, 0x9e, 0x25, 0xb0, 0x23, 0xfc,
	0x9f, 0x7c, 0xfd, 0xb8, 0xeb, 0x6b, 0xb0, 0x57, 0x50, 0x65, 0xdf, 0x26, 0xf6, 0x9c, 0x7d, 0x57,
	0xd1, 0x52, 0x5a, 0xfb, 0x1a, 0xda, 0x85, 0xd2, 0x30, 0xcf, 0x79, 0xb9, 0x02, 0x0e, 0x8b, 0xc4,
	0x22, 0xe7, 0xb1, 0xff, 0x01, 0x9e, 0x57, 0xe8, 0xdd, 0xf5, 0x1f, 0xf9, 0xb9, 0x3a, 0x8e, 0x79,
	0xb1, 0x2f, 0xab, 0x1c, 0x66, 0x81, 0xfe, 0x2b, 0x87, 0x85, 0x3b, 0x90, 0xff, 0x1d, 0x3c, 0xb7,
	0xcb, 0x63, 0xc8, 0x34, 0x2b, 0x5d, 0x1c, 0xdc, 0x8e, 0xf7, 0xbd, 0xaf, 0x60, 0xa0, 0x76, 0xb6,
	0xbf, 0xb5, 0xf7, 0xb6, 0xc5, 0x4e, 0xc1, 0xc3, 0xab, 0xff, 0xf6, 0x80, 0xcc, 0xaf, 0x66, 0x33,
	0xd4, 0x9a, 0x27, 0x2b, 0x45, 0xd1, 0x12, 0xbc, 0x86, 0xb6, 0x16, 0x2a, 0xc4, 0x4d, 0xb1, 0xfa,
	0x9c, 0x71, 0x5a, 0xa8, 0x51, 0x11, 0x93, 0xcf, 0x01, 0x8a, 0xc3, 0x38, 0xdd, 0x30, 0x9e, 0xb8,
	0x25, 0x5a, 0xc0, 0x87, 0x26, 0x51, 0xd4, 0xb2, 0x95, 0x44, 0x8c, 0xc3, 0x34, 0x71, 0xdb, 0xf4,
	0xd0, 0x26, 0xa6, 0x09, 0x19, 0x40, 0x57, 0xa3, 0xdc, 0xa8, 0x30, 0x5d, 0x86, 0x0a, 0xe5, 0x0d,
	0x8f, 0xd0, 0xad, 0xd6, 0x23, 0x93, 0x9f, 0x2e, 0x67, 0x36, 0x4b, 0xbe, 0x84, 0x67, 0x6b, 0xa6,
	0xc2, 0x08, 0xa5, 0xe6, 0x4b, 0x1e, 0x31, 0x8d, 0xfd, 0xe6, 0xa9, 0x37, 0x38, 0xa4, 0x47, 0x6b,
	0xa6, 0x2e, 0xb6, 0x59, 0x7f, 0x09, 0xdd, 0x51, 0xc2, 0xae, 0x05, 0xce, 0xaf, 0x66, 0xa5, 0x7d,
	0xff, 0x47, 0xff, 0x31, 0xb4, 0xac, 0x5c, 0x23, 0xfe, 0x90, 0xba, 0xc8, 0x47, 0x78, 0x71, 0x99,
	0x28, 0xcd, 0x84, 0x70, 0xcf, 0x6d, 0x5b, 0x11, 0x68, 0x24, 0x6c, 0x83, 0xae, 0x8b, 0xf9, 0x4d,
	0x7a, 0x70, 0x70, 0x13, 0xa9, 0x30, 0x97, 0xc2, 0xd1, 0xb7, 0x6e, 0x22, 0xb5, 0x90, 0x82, 0x7c,
	0x01, 0x4f, 0x24, 0xae, 0xb8, 0xd2, 0xf2, 0x37, 0x73, 0x6a, 0xed, 0xe9, 0x94, 0xb9, 0x85, 0x14,
	0x67, 0x43, 0x68, 0xdf, 0xfe, 0x65, 0x92, 0xcf, 0xe0, 0xd5, 0xe5, 0x64, 0x3e, 0xa2, 0x93, 0x1f,
	0xae, 0xc2, 0xd9, 0x88, 0x7e, 0x1a, 0xd1, 0x70, 0x44, 0xe9, 0x94, 0xda, 0xcf, 0x7a, 0x32, 0x9d,
	0x87, 0xe3, 0xe9, 0x62, 0x32, 0xec, 0xd6, 0x8a, 0x70, 0x3c, 0xa5, 0x1f, 0x2e, 0x87, 0xc3, 0xd1,
	0xa4, 0x5b, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xcc, 0x8e, 0x04, 0xf4, 0x07, 0x00, 0x00,
}
