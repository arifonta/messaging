// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: common/communication/mesaging.proto

package communication

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagingServiceClient is the client API for MessagingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingServiceClient interface {
	SendEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
	SendWhatsApp(ctx context.Context, in *WhatsAppRequest, opts ...grpc.CallOption) (*WhatsAppResponse, error)
}

type messagingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingServiceClient(cc grpc.ClientConnInterface) MessagingServiceClient {
	return &messagingServiceClient{cc}
}

func (c *messagingServiceClient) SendEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, "/communication.MessagingService/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) SendWhatsApp(ctx context.Context, in *WhatsAppRequest, opts ...grpc.CallOption) (*WhatsAppResponse, error) {
	out := new(WhatsAppResponse)
	err := c.cc.Invoke(ctx, "/communication.MessagingService/SendWhatsApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServiceServer is the server API for MessagingService service.
// All implementations must embed UnimplementedMessagingServiceServer
// for forward compatibility
type MessagingServiceServer interface {
	SendEmail(context.Context, *EmailRequest) (*EmailResponse, error)
	SendWhatsApp(context.Context, *WhatsAppRequest) (*WhatsAppResponse, error)
	mustEmbedUnimplementedMessagingServiceServer()
}

// UnimplementedMessagingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessagingServiceServer struct {
}

func (UnimplementedMessagingServiceServer) SendEmail(context.Context, *EmailRequest) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedMessagingServiceServer) SendWhatsApp(context.Context, *WhatsAppRequest) (*WhatsAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendWhatsApp not implemented")
}
func (UnimplementedMessagingServiceServer) mustEmbedUnimplementedMessagingServiceServer() {}

// UnsafeMessagingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServiceServer will
// result in compilation errors.
type UnsafeMessagingServiceServer interface {
	mustEmbedUnimplementedMessagingServiceServer()
}

func RegisterMessagingServiceServer(s grpc.ServiceRegistrar, srv MessagingServiceServer) {
	s.RegisterService(&MessagingService_ServiceDesc, srv)
}

func _MessagingService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.MessagingService/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).SendEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_SendWhatsApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhatsAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).SendWhatsApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.MessagingService/SendWhatsApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).SendWhatsApp(ctx, req.(*WhatsAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagingService_ServiceDesc is the grpc.ServiceDesc for MessagingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "communication.MessagingService",
	HandlerType: (*MessagingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _MessagingService_SendEmail_Handler,
		},
		{
			MethodName: "SendWhatsApp",
			Handler:    _MessagingService_SendWhatsApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/communication/mesaging.proto",
}