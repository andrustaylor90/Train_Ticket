// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/train_service.proto

package proto

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

const (
	TrainTicketService_PurchaseTicket_FullMethodName      = "/train.TrainTicketService/PurchaseTicket"
	TrainTicketService_ViewReceipt_FullMethodName         = "/train.TrainTicketService/ViewReceipt"
	TrainTicketService_ViewAllUsers_FullMethodName        = "/train.TrainTicketService/ViewAllUsers"
	TrainTicketService_RemoveUserFromTrain_FullMethodName = "/train.TrainTicketService/RemoveUserFromTrain"
	TrainTicketService_ModifyUserSeat_FullMethodName      = "/train.TrainTicketService/ModifyUserSeat"
)

// TrainTicketServiceClient is the client API for TrainTicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainTicketServiceClient interface {
	PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseReply, error)
	ViewReceipt(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*ReceiptReply, error)
	ViewAllUsers(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*UserListReply, error)
	RemoveUserFromTrain(ctx context.Context, in *ModifyUserRequest, opts ...grpc.CallOption) (*ModifyUserReply, error)
	ModifyUserSeat(ctx context.Context, in *ModifyUserRequest, opts ...grpc.CallOption) (*ModifyUserReply, error)
}

type trainTicketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainTicketServiceClient(cc grpc.ClientConnInterface) TrainTicketServiceClient {
	return &trainTicketServiceClient{cc}
}

func (c *trainTicketServiceClient) PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseReply, error) {
	out := new(PurchaseReply)
	err := c.cc.Invoke(ctx, TrainTicketService_PurchaseTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) ViewReceipt(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*ReceiptReply, error) {
	out := new(ReceiptReply)
	err := c.cc.Invoke(ctx, TrainTicketService_ViewReceipt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) ViewAllUsers(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*UserListReply, error) {
	out := new(UserListReply)
	err := c.cc.Invoke(ctx, TrainTicketService_ViewAllUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) RemoveUserFromTrain(ctx context.Context, in *ModifyUserRequest, opts ...grpc.CallOption) (*ModifyUserReply, error) {
	out := new(ModifyUserReply)
	err := c.cc.Invoke(ctx, TrainTicketService_RemoveUserFromTrain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) ModifyUserSeat(ctx context.Context, in *ModifyUserRequest, opts ...grpc.CallOption) (*ModifyUserReply, error) {
	out := new(ModifyUserReply)
	err := c.cc.Invoke(ctx, TrainTicketService_ModifyUserSeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainTicketServiceServer is the server API for TrainTicketService service.
// All implementations must embed UnimplementedTrainTicketServiceServer
// for forward compatibility
type TrainTicketServiceServer interface {
	PurchaseTicket(context.Context, *PurchaseRequest) (*PurchaseReply, error)
	ViewReceipt(context.Context, *AuthRequest) (*ReceiptReply, error)
	ViewAllUsers(context.Context, *AuthRequest) (*UserListReply, error)
	RemoveUserFromTrain(context.Context, *ModifyUserRequest) (*ModifyUserReply, error)
	ModifyUserSeat(context.Context, *ModifyUserRequest) (*ModifyUserReply, error)
	mustEmbedUnimplementedTrainTicketServiceServer()
}

// UnimplementedTrainTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrainTicketServiceServer struct {
}

func (UnimplementedTrainTicketServiceServer) PurchaseTicket(context.Context, *PurchaseRequest) (*PurchaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTrainTicketServiceServer) ViewReceipt(context.Context, *AuthRequest) (*ReceiptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewReceipt not implemented")
}
func (UnimplementedTrainTicketServiceServer) ViewAllUsers(context.Context, *AuthRequest) (*UserListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAllUsers not implemented")
}
func (UnimplementedTrainTicketServiceServer) RemoveUserFromTrain(context.Context, *ModifyUserRequest) (*ModifyUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromTrain not implemented")
}
func (UnimplementedTrainTicketServiceServer) ModifyUserSeat(context.Context, *ModifyUserRequest) (*ModifyUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserSeat not implemented")
}
func (UnimplementedTrainTicketServiceServer) mustEmbedUnimplementedTrainTicketServiceServer() {}

// UnsafeTrainTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainTicketServiceServer will
// result in compilation errors.
type UnsafeTrainTicketServiceServer interface {
	mustEmbedUnimplementedTrainTicketServiceServer()
}

func RegisterTrainTicketServiceServer(s grpc.ServiceRegistrar, srv TrainTicketServiceServer) {
	s.RegisterService(&TrainTicketService_ServiceDesc, srv)
}

func _TrainTicketService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_PurchaseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).PurchaseTicket(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_ViewReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).ViewReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_ViewReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).ViewReceipt(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_ViewAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).ViewAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_ViewAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).ViewAllUsers(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_RemoveUserFromTrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).RemoveUserFromTrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_RemoveUserFromTrain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).RemoveUserFromTrain(ctx, req.(*ModifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_ModifyUserSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).ModifyUserSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_ModifyUserSeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).ModifyUserSeat(ctx, req.(*ModifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainTicketService_ServiceDesc is the grpc.ServiceDesc for TrainTicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainTicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "train.TrainTicketService",
	HandlerType: (*TrainTicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TrainTicketService_PurchaseTicket_Handler,
		},
		{
			MethodName: "ViewReceipt",
			Handler:    _TrainTicketService_ViewReceipt_Handler,
		},
		{
			MethodName: "ViewAllUsers",
			Handler:    _TrainTicketService_ViewAllUsers_Handler,
		},
		{
			MethodName: "RemoveUserFromTrain",
			Handler:    _TrainTicketService_RemoveUserFromTrain_Handler,
		},
		{
			MethodName: "ModifyUserSeat",
			Handler:    _TrainTicketService_ModifyUserSeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/train_service.proto",
}