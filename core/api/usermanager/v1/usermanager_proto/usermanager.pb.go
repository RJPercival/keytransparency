// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermanager/v1/usermanager_proto/usermanager.proto

/*
Package usermanager_proto is a generated protocol buffer package.

It is generated from these files:
	usermanager/v1/usermanager_proto/usermanager.proto

It has these top-level messages:
	GetKeySetRequest
	CreateUserRequest
	UpdateUserRequest
	BatchCreateUserRequest
	BatchCreateUserResponse
*/
package usermanager_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "google.golang.org/genproto/protobuf/field_mask"
import google_keytransparency_type "github.com/google/keytransparency/core/api/type/type_proto"
import google_keytransparency_type1 "github.com/google/keytransparency/core/api/type/type_proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// GetKeySetsRequest requests the keyset of a domain_id/app_id
type GetKeySetRequest struct {
	// domain_id identifies the domain.
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// app_id identifies the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *GetKeySetRequest) Reset()                    { *m = GetKeySetRequest{} }
func (m *GetKeySetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetKeySetRequest) ProtoMessage()               {}
func (*GetKeySetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetKeySetRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *GetKeySetRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

// CreateUserRequest specifies the information with which a new user should be initialized.
// New users will be signed with the current active key.
// It is the responsibility of authorized callers to verify that domain_id/app_id/user_id is correct.
type CreateUserRequest struct {
	// user is the user to create.
	User *google_keytransparency_type.User `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	// add_signing_keys specifies whether to add this service's signing keys to the set of authorized_keys.
	// This must be set to true if any further operations from this API are meant to succeed.
	// If set to false, there must be at least one key in authorized_keys.
	AddSigningKeys bool `protobuf:"varint,7,opt,name=add_signing_keys,json=addSigningKeys" json:"add_signing_keys,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateUserRequest) GetUser() *google_keytransparency_type.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRequest) GetAddSigningKeys() bool {
	if m != nil {
		return m.AddSigningKeys
	}
	return false
}

// UpdateUserRequest sets the data field for the user.
// The user must have the service's current signing key in its list of
// authorized_keys in order to succeed.
type UpdateUserRequest struct {
	// user contains data which will be applied to the user.
	User *google_keytransparency_type.User `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	// update_mask specifies which fields of user to update.
	// For example: "data" or "authorized_keys"
	UpdateMask *google_protobuf1.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateUserRequest) Reset()                    { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()               {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateUserRequest) GetUser() *google_keytransparency_type.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetUpdateMask() *google_protobuf1.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// BatchCreateUserRequest creates multiple users all at once.
type BatchCreateUserRequest struct {
	// domain_id identifies the domain.
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// app_id identifies the application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// users is the set of users to create.
	Users []*google_keytransparency_type.User `protobuf:"bytes,3,rep,name=users" json:"users,omitempty"`
	// add_signing_keys specifies whether to add this service's signing_keys to the set of authorized_keys.
	// This must be set to true if any further operations from this API are meant to succeed.
	// If set to false, there must be at least one key in authorized_keys.
	AddSigningKeys bool `protobuf:"varint,4,opt,name=add_signing_keys,json=addSigningKeys" json:"add_signing_keys,omitempty"`
}

func (m *BatchCreateUserRequest) Reset()                    { *m = BatchCreateUserRequest{} }
func (m *BatchCreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchCreateUserRequest) ProtoMessage()               {}
func (*BatchCreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BatchCreateUserRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *BatchCreateUserRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *BatchCreateUserRequest) GetUsers() []*google_keytransparency_type.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *BatchCreateUserRequest) GetAddSigningKeys() bool {
	if m != nil {
		return m.AddSigningKeys
	}
	return false
}

// BatchCreateUserResponse creates multiple users at once.
type BatchCreateUserResponse struct {
	// users returns the list of created users.
	Users []*google_keytransparency_type.User `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
}

func (m *BatchCreateUserResponse) Reset()                    { *m = BatchCreateUserResponse{} }
func (m *BatchCreateUserResponse) String() string            { return proto.CompactTextString(m) }
func (*BatchCreateUserResponse) ProtoMessage()               {}
func (*BatchCreateUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BatchCreateUserResponse) GetUsers() []*google_keytransparency_type.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*GetKeySetRequest)(nil), "google.keytransparency.usermanager.v1.GetKeySetRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "google.keytransparency.usermanager.v1.CreateUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "google.keytransparency.usermanager.v1.UpdateUserRequest")
	proto.RegisterType((*BatchCreateUserRequest)(nil), "google.keytransparency.usermanager.v1.BatchCreateUserRequest")
	proto.RegisterType((*BatchCreateUserResponse)(nil), "google.keytransparency.usermanager.v1.BatchCreateUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserManager service

type UserManagerClient interface {
	// GetKeySet returns a list of public keys (a keyset) that corresponds to the signing keys
	// this service has for a given domain and app.
	GetKeySet(ctx context.Context, in *GetKeySetRequest, opts ...grpc.CallOption) (*google_keytransparency_type1.KeySet, error)
	// CreateUser creates a new user and initializes it.
	// If the user already exists, this operation will fail.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*google_keytransparency_type.User, error)
	// UpdateUserData sets the public key for an user.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*google_keytransparency_type.User, error)
	// BatchCreateUser creates a set of new users.
	BatchCreateUser(ctx context.Context, in *BatchCreateUserRequest, opts ...grpc.CallOption) (*google_keytransparency_type.User, error)
}

type userManagerClient struct {
	cc *grpc.ClientConn
}

func NewUserManagerClient(cc *grpc.ClientConn) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) GetKeySet(ctx context.Context, in *GetKeySetRequest, opts ...grpc.CallOption) (*google_keytransparency_type1.KeySet, error) {
	out := new(google_keytransparency_type1.KeySet)
	err := grpc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/GetKeySet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*google_keytransparency_type.User, error) {
	out := new(google_keytransparency_type.User)
	err := grpc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*google_keytransparency_type.User, error) {
	out := new(google_keytransparency_type.User)
	err := grpc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) BatchCreateUser(ctx context.Context, in *BatchCreateUserRequest, opts ...grpc.CallOption) (*google_keytransparency_type.User, error) {
	out := new(google_keytransparency_type.User)
	err := grpc.Invoke(ctx, "/google.keytransparency.usermanager.v1.UserManager/BatchCreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserManager service

type UserManagerServer interface {
	// GetKeySet returns a list of public keys (a keyset) that corresponds to the signing keys
	// this service has for a given domain and app.
	GetKeySet(context.Context, *GetKeySetRequest) (*google_keytransparency_type1.KeySet, error)
	// CreateUser creates a new user and initializes it.
	// If the user already exists, this operation will fail.
	CreateUser(context.Context, *CreateUserRequest) (*google_keytransparency_type.User, error)
	// UpdateUserData sets the public key for an user.
	UpdateUser(context.Context, *UpdateUserRequest) (*google_keytransparency_type.User, error)
	// BatchCreateUser creates a set of new users.
	BatchCreateUser(context.Context, *BatchCreateUserRequest) (*google_keytransparency_type.User, error)
}

func RegisterUserManagerServer(s *grpc.Server, srv UserManagerServer) {
	s.RegisterService(&_UserManager_serviceDesc, srv)
}

func _UserManager_GetKeySet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).GetKeySet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/GetKeySet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).GetKeySet(ctx, req.(*GetKeySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_BatchCreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).BatchCreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.usermanager.v1.UserManager/BatchCreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).BatchCreateUser(ctx, req.(*BatchCreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.usermanager.v1.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeySet",
			Handler:    _UserManager_GetKeySet_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserManager_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserManager_UpdateUser_Handler,
		},
		{
			MethodName: "BatchCreateUser",
			Handler:    _UserManager_BatchCreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermanager/v1/usermanager_proto/usermanager.proto",
}

func init() { proto.RegisterFile("usermanager/v1/usermanager_proto/usermanager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0x65, 0x93, 0x36, 0x36, 0x13, 0xd0, 0x76, 0x40, 0x0d, 0xab, 0x60, 0x5c, 0x11, 0x82, 0x0f,
	0x3b, 0x34, 0x22, 0x2d, 0x91, 0x82, 0x44, 0xa8, 0x96, 0x10, 0x90, 0x94, 0x82, 0xf8, 0x60, 0x98,
	0x64, 0x6e, 0xb7, 0x4b, 0x9a, 0x99, 0x71, 0x67, 0x52, 0x58, 0x4a, 0x5f, 0x7c, 0xf1, 0x03, 0xfa,
	0x19, 0x3e, 0xf8, 0x05, 0x7e, 0x80, 0xcf, 0xfe, 0x82, 0x1f, 0x22, 0x33, 0xb3, 0x69, 0x42, 0x92,
	0xda, 0x8d, 0xe8, 0xcb, 0x66, 0xe7, 0xde, 0x33, 0x73, 0xcf, 0x39, 0x7b, 0x6f, 0x06, 0x35, 0xc6,
	0x0a, 0x92, 0x11, 0xe5, 0x34, 0x82, 0x84, 0x9c, 0x6d, 0x93, 0x99, 0x65, 0x4f, 0x26, 0x42, 0x8b,
	0xd9, 0x48, 0x68, 0x23, 0xf8, 0x69, 0x24, 0x44, 0x74, 0x0a, 0xe1, 0x10, 0x52, 0x9d, 0x50, 0xae,
	0x24, 0x4d, 0x80, 0x0f, 0xd2, 0x70, 0x16, 0x79, 0xb6, 0xed, 0x3f, 0x74, 0x30, 0x42, 0x65, 0x4c,
	0x28, 0xe7, 0x42, 0x53, 0x1d, 0x0b, 0xae, 0xdc, 0x21, 0x7e, 0x2d, 0xcb, 0xda, 0x55, 0x7f, 0x7c,
	0x4c, 0x8e, 0x63, 0x38, 0x65, 0xbd, 0x11, 0x55, 0xc3, 0x0c, 0xe1, 0xeb, 0x54, 0x02, 0x31, 0x8f,
	0x8c, 0x89, 0x79, 0xcd, 0x72, 0x8f, 0xe6, 0x73, 0x43, 0x48, 0x47, 0x54, 0xe9, 0x09, 0xc7, 0x60,
	0x1f, 0x6d, 0xbe, 0x01, 0xdd, 0x86, 0xf4, 0x10, 0x74, 0x17, 0x3e, 0x8d, 0x41, 0x69, 0xfc, 0x00,
	0x95, 0x99, 0x18, 0xd1, 0x98, 0xf7, 0x62, 0x56, 0xf5, 0x6a, 0x5e, 0xbd, 0xdc, 0xdd, 0x70, 0x81,
	0x03, 0x86, 0xef, 0xa2, 0x12, 0x95, 0xd2, 0x64, 0x0a, 0x36, 0xb3, 0x4e, 0xa5, 0x3c, 0x60, 0x81,
	0x46, 0x5b, 0xaf, 0x13, 0xa0, 0x1a, 0x8e, 0x14, 0x24, 0x93, 0x83, 0x5e, 0xa0, 0x35, 0xa3, 0xd5,
	0x22, 0x2b, 0x8d, 0xc7, 0xe1, 0x35, 0x7e, 0x58, 0xbe, 0x76, 0x9f, 0x85, 0xe3, 0x3a, 0xda, 0xa4,
	0x8c, 0xf5, 0x54, 0x1c, 0xf1, 0x98, 0x47, 0xbd, 0x21, 0xa4, 0xaa, 0x7a, 0xab, 0xe6, 0xd5, 0x37,
	0xba, 0xb7, 0x29, 0x63, 0x87, 0x2e, 0xdc, 0x86, 0x54, 0x05, 0x5f, 0x3c, 0xb4, 0x75, 0x24, 0xd9,
	0xbf, 0x29, 0xfb, 0x12, 0x55, 0xc6, 0xf6, 0x2c, 0x6b, 0x6e, 0xb5, 0x68, 0x77, 0xfb, 0x93, 0xdd,
	0x13, 0xff, 0xc3, 0x7d, 0xe3, 0x7f, 0x87, 0xaa, 0x61, 0x17, 0x39, 0xb8, 0x79, 0x0f, 0xbe, 0x79,
	0xe8, 0x5e, 0x8b, 0xea, 0xc1, 0xc9, 0xa2, 0x0b, 0x7f, 0x61, 0x27, 0xde, 0x41, 0xeb, 0x86, 0x93,
	0xaa, 0x16, 0x6b, 0xc5, 0x7c, 0x1a, 0x1c, 0x7e, 0xa9, 0x77, 0x6b, 0x4b, 0xbd, 0xeb, 0xa2, 0xfb,
	0x0b, 0x84, 0x95, 0x14, 0x5c, 0xc1, 0xb4, 0x7a, 0x61, 0xb5, 0xea, 0x8d, 0xcb, 0x12, 0xaa, 0x98,
	0x75, 0xc7, 0x75, 0x37, 0xfe, 0xea, 0xa1, 0xf2, 0x55, 0x7b, 0xe1, 0x9d, 0x30, 0xd7, 0x40, 0x84,
	0xf3, 0x0d, 0xe9, 0x3f, 0xf9, 0x23, 0x01, 0x87, 0x0d, 0x5e, 0x7d, 0xfe, 0xf9, 0xeb, 0xb2, 0xd0,
	0xc4, 0xbb, 0x64, 0x6e, 0x54, 0x9d, 0xd7, 0x8a, 0x9c, 0x5f, 0x7d, 0x85, 0x0b, 0x42, 0xa5, 0x54,
	0xe4, 0xdc, 0x39, 0x7f, 0x61, 0x86, 0x42, 0x81, 0xc6, 0x3f, 0x3c, 0x84, 0xa6, 0x6e, 0xe0, 0xdd,
	0x9c, 0x74, 0x17, 0xbe, 0xb8, 0x7f, 0xb3, 0x61, 0xc1, 0x47, 0xcb, 0xf6, 0x7d, 0xf0, 0xee, 0x5a,
	0xb6, 0x26, 0x1e, 0x2e, 0x50, 0xb6, 0xd1, 0x09, 0x6f, 0x6b, 0x7a, 0x16, 0x33, 0x0f, 0x13, 0x6c,
	0xba, 0x66, 0x36, 0x5a, 0xa6, 0x93, 0x91, 0x5b, 0xcb, 0xc2, 0x30, 0xad, 0xa0, 0xa5, 0xf1, 0xbf,
	0xb4, 0x7c, 0xf7, 0xd0, 0x9d, 0xb9, 0x56, 0xc5, 0x7b, 0x39, 0x05, 0x2d, 0x9f, 0xc9, 0x3c, 0xaa,
	0xde, 0x5a, 0x55, 0xad, 0x60, 0x6f, 0xe5, 0x7e, 0x6a, 0xce, 0x14, 0x6d, 0x7a, 0xcf, 0x5a, 0x9d,
	0x0f, 0xed, 0x28, 0xd6, 0x27, 0xe3, 0x7e, 0x38, 0x10, 0x23, 0x92, 0xfd, 0x9f, 0xcf, 0x15, 0x26,
	0x03, 0x91, 0xb8, 0x2b, 0xe0, 0xa6, 0x8b, 0xa6, 0x5f, 0xb2, 0x3f, 0xcf, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x5d, 0x05, 0x8d, 0xee, 0x93, 0x06, 0x00, 0x00,
}
