// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/train.proto

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
	TrainService_PurchaseTicket_FullMethodName = "/train.TrainService/PurchaseTicket"
	TrainService_GetReceipt_FullMethodName     = "/train.TrainService/GetReceipt"
	TrainService_ViewSeats_FullMethodName      = "/train.TrainService/ViewSeats"
	TrainService_RemoveUser_FullMethodName     = "/train.TrainService/RemoveUser"
	TrainService_ModifySeat_FullMethodName     = "/train.TrainService/ModifySeat"
)

// TrainServiceClient is the client API for TrainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainServiceClient interface {
	PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Receipt, error)
	GetReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Receipt, error)
	ViewSeats(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SeatResponse, error)
	RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*GenericResponse, error)
}

type trainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainServiceClient(cc grpc.ClientConnInterface) TrainServiceClient {
	return &trainServiceClient{cc}
}

func (c *trainServiceClient) PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, TrainService_PurchaseTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) GetReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, TrainService_GetReceipt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ViewSeats(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SeatResponse, error) {
	out := new(SeatResponse)
	err := c.cc.Invoke(ctx, TrainService_ViewSeats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, TrainService_RemoveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, TrainService_ModifySeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainServiceServer is the server API for TrainService service.
// All implementations must embed UnimplementedTrainServiceServer
// for forward compatibility
type TrainServiceServer interface {
	PurchaseTicket(context.Context, *PurchaseRequest) (*Receipt, error)
	GetReceipt(context.Context, *UserRequest) (*Receipt, error)
	ViewSeats(context.Context, *SectionRequest) (*SeatResponse, error)
	RemoveUser(context.Context, *UserRequest) (*GenericResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*GenericResponse, error)
	mustEmbedUnimplementedTrainServiceServer()
}

// UnimplementedTrainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrainServiceServer struct {
}

func (UnimplementedTrainServiceServer) PurchaseTicket(context.Context, *PurchaseRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTrainServiceServer) GetReceipt(context.Context, *UserRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTrainServiceServer) ViewSeats(context.Context, *SectionRequest) (*SeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewSeats not implemented")
}
func (UnimplementedTrainServiceServer) RemoveUser(context.Context, *UserRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTrainServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTrainServiceServer) mustEmbedUnimplementedTrainServiceServer() {}

// UnsafeTrainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainServiceServer will
// result in compilation errors.
type UnsafeTrainServiceServer interface {
	mustEmbedUnimplementedTrainServiceServer()
}

func RegisterTrainServiceServer(s grpc.ServiceRegistrar, srv TrainServiceServer) {
	s.RegisterService(&TrainService_ServiceDesc, srv)
}

func _TrainService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_PurchaseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_GetReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).GetReceipt(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ViewSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ViewSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_ViewSeats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ViewSeats(ctx, req.(*SectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).RemoveUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_ModifySeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainService_ServiceDesc is the grpc.ServiceDesc for TrainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "train.TrainService",
	HandlerType: (*TrainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TrainService_PurchaseTicket_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TrainService_GetReceipt_Handler,
		},
		{
			MethodName: "ViewSeats",
			Handler:    _TrainService_ViewSeats_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TrainService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TrainService_ModifySeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/train.proto",
}
