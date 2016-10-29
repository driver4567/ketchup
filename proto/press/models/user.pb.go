// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	Uuid             *string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Email            *string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	HashedPassword   *string `protobuf:"bytes,3,opt,name=hashed_password,json=hashedPassword" json:"hashed_password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *User) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *User) GetHashedPassword() string {
	if m != nil && m.HashedPassword != nil {
		return *m.HashedPassword
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "press.models.User")
}

var fileDescriptor4 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x28, 0x4a, 0x2d, 0x2e, 0xd6, 0xcb, 0xcd,
	0x4f, 0x49, 0xcd, 0x29, 0x56, 0x8a, 0xe4, 0x62, 0x09, 0x05, 0xca, 0x09, 0x09, 0x71, 0xb1, 0x94,
	0x96, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x22, 0x5c, 0xac,
	0xa9, 0xb9, 0x89, 0x99, 0x39, 0x12, 0x4c, 0x60, 0x41, 0x08, 0x47, 0x48, 0x9d, 0x8b, 0x3f, 0x23,
	0xb1, 0x38, 0x23, 0x35, 0x25, 0xbe, 0x20, 0xb1, 0xb8, 0xb8, 0x3c, 0xbf, 0x28, 0x45, 0x82, 0x19,
	0x2c, 0xcf, 0x07, 0x11, 0x0e, 0x80, 0x8a, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xca, 0x07,
	0x92, 0x75, 0x00, 0x00, 0x00,
}
