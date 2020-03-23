// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/credentials/v1/iamcredentials.proto

package credentials

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("google/iam/credentials/v1/iamcredentials.proto", fileDescriptor_ad032f4dbfa7437c)
}

var fileDescriptor_ad032f4dbfa7437c = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xb1, 0x6e, 0xd4, 0x4c,
	0x10, 0xc7, 0xb5, 0x5f, 0xf1, 0x05, 0xb9, 0x00, 0xc9, 0x91, 0x40, 0x9c, 0x48, 0x63, 0x04, 0x12,
	0xe6, 0xf0, 0x72, 0x81, 0x80, 0x74, 0x01, 0x09, 0x1f, 0x48, 0x28, 0x27, 0x21, 0x05, 0x92, 0x8a,
	0x06, 0x6d, 0xd6, 0x73, 0xce, 0x86, 0xdd, 0x1d, 0xe3, 0x5d, 0x9f, 0x85, 0x10, 0x0d, 0x12, 0x4f,
	0xc0, 0x73, 0x50, 0x52, 0xf0, 0x0a, 0x94, 0xd0, 0x5d, 0x9d, 0x92, 0x07, 0xa0, 0x44, 0xeb, 0x5b,
	0x2b, 0xc7, 0x71, 0x09, 0x77, 0xa5, 0x77, 0xfe, 0x33, 0xf3, 0xfb, 0x15, 0xe3, 0x20, 0xc9, 0x11,
	0x73, 0x09, 0x54, 0x30, 0x45, 0x79, 0x09, 0x19, 0x68, 0x2b, 0x98, 0x34, 0x74, 0xdc, 0x73, 0x4f,
	0x33, 0x2f, 0x49, 0x51, 0xa2, 0xc5, 0xf0, 0xf2, 0x34, 0x9f, 0x08, 0xa6, 0x92, 0xd9, 0xea, 0xb8,
	0xd7, 0xb9, 0xe2, 0x47, 0xb1, 0x42, 0x50, 0xa6, 0x35, 0x5a, 0x66, 0x05, 0x6a, 0xdf, 0xd8, 0xb9,
	0x34, 0x53, 0xe5, 0x52, 0x80, 0xb6, 0xbe, 0x70, 0xfd, 0x74, 0x02, 0x8e, 0x4a, 0xa1, 0x9e, 0xe6,
	0x36, 0x3f, 0xaf, 0x05, 0xe7, 0x77, 0xd2, 0x67, 0x8f, 0x4f, 0x22, 0xe1, 0x4f, 0x12, 0xac, 0x3f,
	0x05, 0x0d, 0x25, 0xb3, 0x90, 0x72, 0x0e, 0xc6, 0xec, 0xe3, 0x6b, 0xd0, 0xe1, 0x56, 0x72, 0x2a,
	0x65, 0xb2, 0x20, 0xff, 0x02, 0xde, 0x54, 0x60, 0x6c, 0xe7, 0xde, 0xaa, 0x6d, 0xa6, 0x40, 0x6d,
	0x20, 0x1a, 0x4d, 0xd2, 0x0d, 0xcd, 0x14, 0x74, 0x33, 0x90, 0x90, 0x33, 0x0b, 0xa6, 0x6b, 0x38,
	0x16, 0xd0, 0x95, 0x62, 0x04, 0x56, 0x28, 0xf8, 0xf0, 0xe3, 0xf8, 0xd3, 0x7f, 0x8f, 0xa2, 0x6d,
	0xe7, 0xf4, 0xce, 0x25, 0x1f, 0x16, 0x25, 0x1e, 0x01, 0xb7, 0x86, 0xc6, 0xd4, 0x40, 0x39, 0x16,
	0xdc, 0x0d, 0xc6, 0x4a, 0xbb, 0x97, 0xf7, 0xfd, 0xfc, 0xef, 0x65, 0x7d, 0x12, 0x87, 0xc7, 0x24,
	0xb8, 0xd0, 0x72, 0xec, 0x64, 0x53, 0xd5, 0xde, 0x12, 0xcc, 0x3e, 0xdb, 0x6a, 0x6e, 0xae, 0xd2,
	0xe2, 0x15, 0x8f, 0x26, 0xe9, 0xb5, 0x39, 0x45, 0x56, 0x65, 0x02, 0x34, 0x87, 0xae, 0xd0, 0x5c,
	0x56, 0x19, 0xbc, 0x02, 0xc5, 0x84, 0x6c, 0x54, 0x1f, 0x44, 0xf7, 0x57, 0x55, 0xf5, 0x0b, 0x9d,
	0xe6, 0x57, 0x12, 0x9c, 0xdb, 0x13, 0xb9, 0x1e, 0x48, 0x3c, 0x08, 0xe3, 0x33, 0x60, 0xdb, 0x50,
	0x2b, 0x76, 0x73, 0xa9, 0xac, 0x37, 0xda, 0x9f, 0xa4, 0x17, 0xe7, 0x8c, 0x0a, 0xf6, 0x56, 0x22,
	0xcb, 0x1a, 0x85, 0xad, 0xe8, 0xf6, 0xb2, 0x0a, 0xc6, 0x8f, 0x76, 0xec, 0x5f, 0x48, 0xb0, 0xe6,
	0x56, 0x0d, 0x6b, 0x1b, 0xde, 0xf8, 0x07, 0xce, 0xb0, 0xb6, 0x2d, 0x79, 0xbc, 0x4c, 0xd4, 0x83,
	0xef, 0x9d, 0x0d, 0x7e, 0x37, 0xa2, 0xab, 0x80, 0x0f, 0x6b, 0xdb, 0x27, 0x71, 0xe7, 0xf9, 0xb7,
	0x74, 0x63, 0xee, 0xde, 0xa7, 0x44, 0xac, 0x10, 0x26, 0xe1, 0xa8, 0xbe, 0xa7, 0xc9, 0xa1, 0xb5,
	0x85, 0xe9, 0x53, 0x5a, 0xd7, 0xf5, 0x5c, 0x91, 0xb2, 0xca, 0x1e, 0x52, 0x2e, 0xb1, 0xca, 0x6e,
	0x15, 0x92, 0xd9, 0x11, 0x96, 0x6a, 0xf0, 0x91, 0x04, 0x57, 0x39, 0xaa, 0xd6, 0xac, 0xa9, 0x2e,
	0xf0, 0x1b, 0xac, 0xff, 0x79, 0xd4, 0xbb, 0xee, 0xd8, 0x77, 0xc9, 0xcb, 0x27, 0xbe, 0x2f, 0x47,
	0xc9, 0x74, 0x9e, 0x60, 0x99, 0xd3, 0x1c, 0x74, 0xf3, 0x2b, 0xa0, 0x27, 0xdb, 0x17, 0xfc, 0x35,
	0xb6, 0x67, 0x3e, 0x7f, 0x11, 0x72, 0xf0, 0x7f, 0xd3, 0x73, 0xe7, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x25, 0x4f, 0x2f, 0xdf, 0xea, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IAMCredentialsClient is the client API for IAMCredentials service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IAMCredentialsClient interface {
	// Generates an OAuth 2.0 access token for a service account.
	GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error)
	// Generates an OpenID Connect ID token for a service account.
	GenerateIdToken(ctx context.Context, in *GenerateIdTokenRequest, opts ...grpc.CallOption) (*GenerateIdTokenResponse, error)
	// Signs a blob using a service account's system-managed private key.
	SignBlob(ctx context.Context, in *SignBlobRequest, opts ...grpc.CallOption) (*SignBlobResponse, error)
	// Signs a JWT using a service account's system-managed private key.
	SignJwt(ctx context.Context, in *SignJwtRequest, opts ...grpc.CallOption) (*SignJwtResponse, error)
}

type iAMCredentialsClient struct {
	cc grpc.ClientConnInterface
}

func NewIAMCredentialsClient(cc grpc.ClientConnInterface) IAMCredentialsClient {
	return &iAMCredentialsClient{cc}
}

func (c *iAMCredentialsClient) GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error) {
	out := new(GenerateAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/GenerateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) GenerateIdToken(ctx context.Context, in *GenerateIdTokenRequest, opts ...grpc.CallOption) (*GenerateIdTokenResponse, error) {
	out := new(GenerateIdTokenResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/GenerateIdToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) SignBlob(ctx context.Context, in *SignBlobRequest, opts ...grpc.CallOption) (*SignBlobResponse, error) {
	out := new(SignBlobResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/SignBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMCredentialsClient) SignJwt(ctx context.Context, in *SignJwtRequest, opts ...grpc.CallOption) (*SignJwtResponse, error) {
	out := new(SignJwtResponse)
	err := c.cc.Invoke(ctx, "/google.iam.credentials.v1.IAMCredentials/SignJwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMCredentialsServer is the server API for IAMCredentials service.
type IAMCredentialsServer interface {
	// Generates an OAuth 2.0 access token for a service account.
	GenerateAccessToken(context.Context, *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error)
	// Generates an OpenID Connect ID token for a service account.
	GenerateIdToken(context.Context, *GenerateIdTokenRequest) (*GenerateIdTokenResponse, error)
	// Signs a blob using a service account's system-managed private key.
	SignBlob(context.Context, *SignBlobRequest) (*SignBlobResponse, error)
	// Signs a JWT using a service account's system-managed private key.
	SignJwt(context.Context, *SignJwtRequest) (*SignJwtResponse, error)
}

// UnimplementedIAMCredentialsServer can be embedded to have forward compatible implementations.
type UnimplementedIAMCredentialsServer struct {
}

func (*UnimplementedIAMCredentialsServer) GenerateAccessToken(ctx context.Context, req *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAccessToken not implemented")
}
func (*UnimplementedIAMCredentialsServer) GenerateIdToken(ctx context.Context, req *GenerateIdTokenRequest) (*GenerateIdTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateIdToken not implemented")
}
func (*UnimplementedIAMCredentialsServer) SignBlob(ctx context.Context, req *SignBlobRequest) (*SignBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignBlob not implemented")
}
func (*UnimplementedIAMCredentialsServer) SignJwt(ctx context.Context, req *SignJwtRequest) (*SignJwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignJwt not implemented")
}

func RegisterIAMCredentialsServer(s *grpc.Server, srv IAMCredentialsServer) {
	s.RegisterService(&_IAMCredentials_serviceDesc, srv)
}

func _IAMCredentials_GenerateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).GenerateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/GenerateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).GenerateAccessToken(ctx, req.(*GenerateAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_GenerateIdToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateIdTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).GenerateIdToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/GenerateIdToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).GenerateIdToken(ctx, req.(*GenerateIdTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_SignBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).SignBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/SignBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).SignBlob(ctx, req.(*SignBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMCredentials_SignJwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignJwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMCredentialsServer).SignJwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.credentials.v1.IAMCredentials/SignJwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMCredentialsServer).SignJwt(ctx, req.(*SignJwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IAMCredentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.iam.credentials.v1.IAMCredentials",
	HandlerType: (*IAMCredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateAccessToken",
			Handler:    _IAMCredentials_GenerateAccessToken_Handler,
		},
		{
			MethodName: "GenerateIdToken",
			Handler:    _IAMCredentials_GenerateIdToken_Handler,
		},
		{
			MethodName: "SignBlob",
			Handler:    _IAMCredentials_SignBlob_Handler,
		},
		{
			MethodName: "SignJwt",
			Handler:    _IAMCredentials_SignJwt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/iam/credentials/v1/iamcredentials.proto",
}