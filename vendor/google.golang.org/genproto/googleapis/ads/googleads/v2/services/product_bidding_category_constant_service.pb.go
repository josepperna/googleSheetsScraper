// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/product_bidding_category_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for
// [ProductBiddingCategoryService.GetProductBiddingCategory][].
type GetProductBiddingCategoryConstantRequest struct {
	// Required. Resource name of the Product Bidding Category to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductBiddingCategoryConstantRequest) Reset() {
	*m = GetProductBiddingCategoryConstantRequest{}
}
func (m *GetProductBiddingCategoryConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductBiddingCategoryConstantRequest) ProtoMessage()    {}
func (*GetProductBiddingCategoryConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d5e5b75993b3b91, []int{0}
}

func (m *GetProductBiddingCategoryConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Unmarshal(m, b)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Merge(m, src)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Size(m)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductBiddingCategoryConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductBiddingCategoryConstantRequest proto.InternalMessageInfo

func (m *GetProductBiddingCategoryConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductBiddingCategoryConstantRequest)(nil), "google.ads.googleads.v2.services.GetProductBiddingCategoryConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/product_bidding_category_constant_service.proto", fileDescriptor_8d5e5b75993b3b91)
}

var fileDescriptor_8d5e5b75993b3b91 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x99, 0x2c, 0x08, 0x16, 0xbd, 0xf4, 0xe2, 0x52, 0x05, 0xc7, 0x65, 0x85, 0x61, 0x0f,
	0x09, 0x54, 0x44, 0x88, 0x78, 0x48, 0xf7, 0x30, 0x2a, 0x28, 0x65, 0x95, 0x39, 0x48, 0xa1, 0x64,
	0x9a, 0x18, 0x03, 0x6d, 0x52, 0x93, 0x4c, 0x41, 0xc4, 0x83, 0x7e, 0x03, 0xf1, 0x1b, 0x78, 0xf4,
	0x93, 0xc8, 0x5e, 0xbd, 0x79, 0xf2, 0xe0, 0xc9, 0x4f, 0x21, 0xdd, 0x34, 0xdd, 0x9d, 0xc3, 0xfc,
	0xb9, 0x3d, 0xcc, 0xfb, 0xcc, 0xef, 0x7d, 0xfa, 0xbc, 0x6d, 0x94, 0x0b, 0xad, 0x45, 0xcd, 0x11,
	0x65, 0x16, 0x79, 0xd9, 0xab, 0x2e, 0x45, 0x96, 0x9b, 0x4e, 0x56, 0xdc, 0xa2, 0xd6, 0x68, 0xb6,
	0xaa, 0x5c, 0xb9, 0x94, 0x8c, 0x49, 0x25, 0xca, 0x8a, 0x3a, 0x2e, 0xb4, 0xf9, 0x50, 0x56, 0x5a,
	0x59, 0x47, 0x95, 0x2b, 0x07, 0x2b, 0x6c, 0x8d, 0x76, 0x3a, 0x9e, 0x7a, 0x0c, 0xa4, 0xcc, 0xc2,
	0x91, 0x08, 0xbb, 0x14, 0x06, 0x62, 0xf2, 0x6c, 0xd3, 0x4e, 0xc3, 0xad, 0x5e, 0x99, 0xbd, 0x96,
	0xfa, 0x65, 0xc9, 0x9d, 0x80, 0x6a, 0x25, 0xa2, 0x4a, 0x69, 0x47, 0x9d, 0xd4, 0xca, 0x0e, 0xd3,
	0x5b, 0x57, 0xa6, 0x55, 0x2d, 0xf9, 0xf8, 0xb7, 0xbb, 0x57, 0x06, 0x6f, 0x25, 0xaf, 0x59, 0xb9,
	0xe4, 0xef, 0x68, 0x27, 0xb5, 0xf1, 0x86, 0xa3, 0xd7, 0xd1, 0x6c, 0xce, 0x5d, 0xee, 0x53, 0x64,
	0x3e, 0xc4, 0xe9, 0x90, 0xe1, 0x74, 0x88, 0x70, 0xc6, 0xdf, 0xaf, 0xb8, 0x75, 0xf1, 0x2c, 0xba,
	0x19, 0x82, 0x97, 0x8a, 0x36, 0xfc, 0x70, 0x32, 0x9d, 0xcc, 0xae, 0x67, 0x07, 0x7f, 0x08, 0x38,
	0xbb, 0x11, 0x26, 0x2f, 0x69, 0xc3, 0xd3, 0x9f, 0x20, 0xba, 0xbf, 0x9d, 0xf9, 0xca, 0x77, 0x14,
	0x7f, 0x06, 0xd1, 0xbd, 0x9d, 0x01, 0xe2, 0xe7, 0x70, 0x57, 0xd7, 0x70, 0xdf, 0xa7, 0x48, 0xc8,
	0x46, 0xd6, 0x78, 0x15, 0xb8, 0x9d, 0x74, 0xf4, 0xe2, 0x37, 0x59, 0x6f, 0xe2, 0xcb, 0xaf, 0xbf,
	0xdf, 0xc0, 0xa3, 0xf8, 0x61, 0x7f, 0xdb, 0x8f, 0x6b, 0x93, 0x27, 0xed, 0x56, 0x94, 0x45, 0x27,
	0x9f, 0x92, 0xdb, 0xe7, 0xe4, 0xf0, 0x32, 0xc8, 0xa0, 0x5a, 0x69, 0x61, 0xa5, 0x9b, 0xec, 0x2b,
	0x88, 0x8e, 0x2b, 0xdd, 0xec, 0x2c, 0x20, 0x3b, 0xd9, 0xab, 0xf0, 0xbc, 0xbf, 0x7a, 0x3e, 0x79,
	0xf3, 0x74, 0xe0, 0x09, 0x5d, 0x53, 0x25, 0xa0, 0x36, 0x02, 0x09, 0xae, 0x2e, 0xde, 0x09, 0x74,
	0x99, 0x60, 0xf3, 0xd7, 0xf2, 0x38, 0x88, 0xef, 0xe0, 0x60, 0x4e, 0xc8, 0x0f, 0x30, 0x9d, 0x7b,
	0x20, 0x61, 0x16, 0x7a, 0xd9, 0xab, 0x45, 0x0a, 0x87, 0xc5, 0xf6, 0x3c, 0x58, 0x0a, 0xc2, 0x6c,
	0x31, 0x5a, 0x8a, 0x45, 0x5a, 0x04, 0xcb, 0x3f, 0x70, 0xec, 0x7f, 0xc7, 0x98, 0x30, 0x8b, 0xf1,
	0x68, 0xc2, 0x78, 0x91, 0x62, 0x1c, 0x6c, 0xcb, 0x6b, 0x17, 0x39, 0x1f, 0xfc, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x25, 0x46, 0xd1, 0x28, 0xd4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductBiddingCategoryConstantServiceClient is the client API for ProductBiddingCategoryConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductBiddingCategoryConstantServiceClient interface {
	// Returns the requested Product Bidding Category in full detail.
	GetProductBiddingCategoryConstant(ctx context.Context, in *GetProductBiddingCategoryConstantRequest, opts ...grpc.CallOption) (*resources.ProductBiddingCategoryConstant, error)
}

type productBiddingCategoryConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductBiddingCategoryConstantServiceClient(cc grpc.ClientConnInterface) ProductBiddingCategoryConstantServiceClient {
	return &productBiddingCategoryConstantServiceClient{cc}
}

func (c *productBiddingCategoryConstantServiceClient) GetProductBiddingCategoryConstant(ctx context.Context, in *GetProductBiddingCategoryConstantRequest, opts ...grpc.CallOption) (*resources.ProductBiddingCategoryConstant, error) {
	out := new(resources.ProductBiddingCategoryConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.ProductBiddingCategoryConstantService/GetProductBiddingCategoryConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductBiddingCategoryConstantServiceServer is the server API for ProductBiddingCategoryConstantService service.
type ProductBiddingCategoryConstantServiceServer interface {
	// Returns the requested Product Bidding Category in full detail.
	GetProductBiddingCategoryConstant(context.Context, *GetProductBiddingCategoryConstantRequest) (*resources.ProductBiddingCategoryConstant, error)
}

// UnimplementedProductBiddingCategoryConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductBiddingCategoryConstantServiceServer struct {
}

func (*UnimplementedProductBiddingCategoryConstantServiceServer) GetProductBiddingCategoryConstant(ctx context.Context, req *GetProductBiddingCategoryConstantRequest) (*resources.ProductBiddingCategoryConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductBiddingCategoryConstant not implemented")
}

func RegisterProductBiddingCategoryConstantServiceServer(s *grpc.Server, srv ProductBiddingCategoryConstantServiceServer) {
	s.RegisterService(&_ProductBiddingCategoryConstantService_serviceDesc, srv)
}

func _ProductBiddingCategoryConstantService_GetProductBiddingCategoryConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductBiddingCategoryConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductBiddingCategoryConstantServiceServer).GetProductBiddingCategoryConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.ProductBiddingCategoryConstantService/GetProductBiddingCategoryConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductBiddingCategoryConstantServiceServer).GetProductBiddingCategoryConstant(ctx, req.(*GetProductBiddingCategoryConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductBiddingCategoryConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.ProductBiddingCategoryConstantService",
	HandlerType: (*ProductBiddingCategoryConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductBiddingCategoryConstant",
			Handler:    _ProductBiddingCategoryConstantService_GetProductBiddingCategoryConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/product_bidding_category_constant_service.proto",
}