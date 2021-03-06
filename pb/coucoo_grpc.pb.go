// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CouponClient is the client API for Coupon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CouponClient interface {
	RequestCoupon(ctx context.Context, in *User, opts ...grpc.CallOption) (*Success, error)
}

type couponClient struct {
	cc grpc.ClientConnInterface
}

func NewCouponClient(cc grpc.ClientConnInterface) CouponClient {
	return &couponClient{cc}
}

func (c *couponClient) RequestCoupon(ctx context.Context, in *User, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Coupon/RequestCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponServer is the server API for Coupon service.
// All implementations must embed UnimplementedCouponServer
// for forward compatibility
type CouponServer interface {
	RequestCoupon(context.Context, *User) (*Success, error)
	mustEmbedUnimplementedCouponServer()
}

// UnimplementedCouponServer must be embedded to have forward compatible implementations.
type UnimplementedCouponServer struct {
}

func (UnimplementedCouponServer) RequestCoupon(context.Context, *User) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestCoupon not implemented")
}
func (UnimplementedCouponServer) mustEmbedUnimplementedCouponServer() {}

// UnsafeCouponServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CouponServer will
// result in compilation errors.
type UnsafeCouponServer interface {
	mustEmbedUnimplementedCouponServer()
}

func RegisterCouponServer(s grpc.ServiceRegistrar, srv CouponServer) {
	s.RegisterService(&_Coupon_serviceDesc, srv)
}

func _Coupon_RequestCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).RequestCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Coupon/RequestCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).RequestCoupon(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coupon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Coupon",
	HandlerType: (*CouponServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestCoupon",
			Handler:    _Coupon_RequestCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coucoo.proto",
}
