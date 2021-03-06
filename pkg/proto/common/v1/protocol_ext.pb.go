// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol_ext.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Type of HTTP 1.1 Verbs
// RFC 7231: https://tools.ietf.org/html/rfc7231#page-24
type HTTPExtensions_HTTPVerb int32

const (
	HTTPExtensions_GET     HTTPExtensions_HTTPVerb = 0
	HTTPExtensions_HEAD    HTTPExtensions_HTTPVerb = 1
	HTTPExtensions_POST    HTTPExtensions_HTTPVerb = 2
	HTTPExtensions_PUT     HTTPExtensions_HTTPVerb = 3
	HTTPExtensions_DELETE  HTTPExtensions_HTTPVerb = 4
	HTTPExtensions_CONNECT HTTPExtensions_HTTPVerb = 5
	HTTPExtensions_OPTIONS HTTPExtensions_HTTPVerb = 6
	HTTPExtensions_TRACE   HTTPExtensions_HTTPVerb = 7
)

var HTTPExtensions_HTTPVerb_name = map[int32]string{
	0: "GET",
	1: "HEAD",
	2: "POST",
	3: "PUT",
	4: "DELETE",
	5: "CONNECT",
	6: "OPTIONS",
	7: "TRACE",
}

var HTTPExtensions_HTTPVerb_value = map[string]int32{
	"GET":     0,
	"HEAD":    1,
	"POST":    2,
	"PUT":     3,
	"DELETE":  4,
	"CONNECT": 5,
	"OPTIONS": 6,
	"TRACE":   7,
}

func (x HTTPExtensions_HTTPVerb) String() string {
	return proto.EnumName(HTTPExtensions_HTTPVerb_name, int32(x))
}

func (HTTPExtensions_HTTPVerb) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6cc522cf428f2150, []int{0, 0}
}

// HTTPExtensions includes HTTP verb and querystring
// when Dapr runtime delivers HTTP content.
//
// For example, when callers calls http invoke api
// POST http://localhost:3500/v1.0/invoke/<app_id>/method/<method>?query1=value1&query2=value2
//
// Dapr runtime will parse POST as a verb and extract querystring to quersytring map.
type HTTPExtensions struct {
	// verb is HTTP verb.
	//
	// This is required.
	Verb HTTPExtensions_HTTPVerb `protobuf:"varint,1,opt,name=verb,proto3,enum=dapr.pkg.proto.common.v1.HTTPExtensions_HTTPVerb" json:"verb,omitempty"`
	// querystring includes HTTP querystring.
	Querystring          map[string]string `protobuf:"bytes,2,rep,name=querystring,proto3" json:"querystring,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HTTPExtensions) Reset()         { *m = HTTPExtensions{} }
func (m *HTTPExtensions) String() string { return proto.CompactTextString(m) }
func (*HTTPExtensions) ProtoMessage()    {}
func (*HTTPExtensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc522cf428f2150, []int{0}
}

func (m *HTTPExtensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPExtensions.Unmarshal(m, b)
}
func (m *HTTPExtensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPExtensions.Marshal(b, m, deterministic)
}
func (m *HTTPExtensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPExtensions.Merge(m, src)
}
func (m *HTTPExtensions) XXX_Size() int {
	return xxx_messageInfo_HTTPExtensions.Size(m)
}
func (m *HTTPExtensions) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPExtensions.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPExtensions proto.InternalMessageInfo

func (m *HTTPExtensions) GetVerb() HTTPExtensions_HTTPVerb {
	if m != nil {
		return m.Verb
	}
	return HTTPExtensions_GET
}

func (m *HTTPExtensions) GetQuerystring() map[string]string {
	if m != nil {
		return m.Querystring
	}
	return nil
}

func init() {
	proto.RegisterEnum("dapr.pkg.proto.common.v1.HTTPExtensions_HTTPVerb", HTTPExtensions_HTTPVerb_name, HTTPExtensions_HTTPVerb_value)
	proto.RegisterType((*HTTPExtensions)(nil), "dapr.pkg.proto.common.v1.HTTPExtensions")
	proto.RegisterMapType((map[string]string)(nil), "dapr.pkg.proto.common.v1.HTTPExtensions.QuerystringEntry")
}

func init() { proto.RegisterFile("protocol_ext.proto", fileDescriptor_6cc522cf428f2150) }

var fileDescriptor_6cc522cf428f2150 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcd, 0x6a, 0xe3, 0x30,
	0x10, 0x80, 0xd7, 0x76, 0x7e, 0x27, 0x10, 0x84, 0xd8, 0x83, 0xd9, 0xbd, 0x84, 0x9c, 0x72, 0x92,
	0x71, 0xf6, 0xb2, 0xbb, 0x87, 0x85, 0xc4, 0x11, 0xc9, 0x42, 0x89, 0x5d, 0x47, 0xed, 0xa1, 0x3d,
	0xb4, 0x76, 0x2a, 0x5c, 0x93, 0xc4, 0x72, 0x65, 0xd9, 0x24, 0xaf, 0xd4, 0x17, 0xea, 0xeb, 0x14,
	0xd9, 0x94, 0xfe, 0x40, 0xa1, 0x17, 0x31, 0x9f, 0x66, 0xbe, 0x61, 0x66, 0x00, 0xe7, 0x52, 0x28,
	0xb1, 0x15, 0xfb, 0x1b, 0x7e, 0x54, 0xa4, 0x06, 0x6c, 0xdf, 0x45, 0xb9, 0x24, 0xf9, 0x2e, 0x69,
	0x98, 0x6c, 0xc5, 0xe1, 0x20, 0x32, 0x52, 0xb9, 0xe3, 0x27, 0x13, 0x86, 0x2b, 0xc6, 0x02, 0x7a,
	0x54, 0x3c, 0x2b, 0x52, 0x91, 0x15, 0x98, 0x42, 0xab, 0xe2, 0x32, 0xb6, 0x8d, 0x91, 0x31, 0x19,
	0x4e, 0x5d, 0xf2, 0x99, 0x4b, 0xde, 0x7b, 0x35, 0x5e, 0x72, 0x19, 0x87, 0xb5, 0x8e, 0xaf, 0x61,
	0xf0, 0x50, 0x72, 0x79, 0x2a, 0x94, 0x4c, 0xb3, 0xc4, 0x36, 0x47, 0xd6, 0x64, 0x30, 0xfd, 0xf3,
	0xe5, 0x6e, 0xe7, 0xaf, 0x2e, 0xcd, 0x94, 0x3c, 0x85, 0x6f, 0xbb, 0xfd, 0xf8, 0x07, 0xe8, 0x63,
	0x01, 0x46, 0x60, 0xed, 0xf8, 0xa9, 0x1e, 0xbb, 0x1f, 0xea, 0x10, 0x7f, 0x87, 0x76, 0x15, 0xed,
	0x4b, 0x6e, 0x9b, 0xf5, 0x5f, 0x03, 0x7f, 0xcd, 0xdf, 0xc6, 0x38, 0x82, 0xde, 0xcb, 0xb8, 0xb8,
	0x0b, 0xd6, 0x92, 0x32, 0xf4, 0x0d, 0xf7, 0xa0, 0xb5, 0xa2, 0xb3, 0x05, 0x32, 0x74, 0x14, 0xf8,
	0x1b, 0x86, 0x4c, 0x9d, 0x0c, 0x2e, 0x18, 0xb2, 0x30, 0x40, 0x67, 0x41, 0xcf, 0x28, 0xa3, 0xa8,
	0x85, 0x07, 0xd0, 0xf5, 0xfc, 0xf5, 0x9a, 0x7a, 0x0c, 0xb5, 0x35, 0xf8, 0x01, 0xfb, 0xef, 0xaf,
	0x37, 0xa8, 0x83, 0xfb, 0xd0, 0x66, 0xe1, 0xcc, 0xa3, 0xa8, 0x3b, 0xbf, 0x05, 0x48, 0x45, 0xb3,
	0x6e, 0xe5, 0xce, 0x61, 0x11, 0xe5, 0x32, 0xd0, 0x2b, 0x17, 0x57, 0x93, 0x24, 0x55, 0xf7, 0x65,
	0xac, 0x77, 0x77, 0x74, 0xbe, 0x79, 0xf2, 0x5d, 0xe2, 0xd4, 0x37, 0x71, 0x9a, 0x9b, 0x38, 0x95,
	0xfb, 0x68, 0xfe, 0xd4, 0x1a, 0xf1, 0xf6, 0x29, 0xcf, 0x14, 0x99, 0x95, 0x4a, 0x24, 0x3c, 0x23,
	0x4b, 0x99, 0x6f, 0x49, 0xe5, 0xc6, 0x9d, 0xba, 0xfc, 0xd7, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0x7d, 0xad, 0x0e, 0xf2, 0x01, 0x00, 0x00,
}
