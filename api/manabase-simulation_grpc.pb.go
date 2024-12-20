// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: api/manabase-simulation.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ManabaseSimulator_Echo_FullMethodName             = "/manabase_simulation.ManabaseSimulator/Echo"
	ManabaseSimulator_SimulateDeck_FullMethodName     = "/manabase_simulation.ManabaseSimulator/SimulateDeck"
	ManabaseSimulator_ValidateDeckList_FullMethodName = "/manabase_simulation.ManabaseSimulator/ValidateDeckList"
)

// ManabaseSimulatorClient is the client API for ManabaseSimulator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManabaseSimulatorClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	// SimulateDeck simulates a deck against a given objective with the provided configuration.
	SimulateDeck(ctx context.Context, in *SimulateDeckRequest, opts ...grpc.CallOption) (*SimulateDeckResponse, error)
	// ValidateDeckList validates a deckList is allowed to be played and can be simulated.
	ValidateDeckList(ctx context.Context, in *ValidateDeckListRequest, opts ...grpc.CallOption) (*ValidateDeckListResponse, error)
}

type manabaseSimulatorClient struct {
	cc grpc.ClientConnInterface
}

func NewManabaseSimulatorClient(cc grpc.ClientConnInterface) ManabaseSimulatorClient {
	return &manabaseSimulatorClient{cc}
}

func (c *manabaseSimulatorClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, ManabaseSimulator_Echo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manabaseSimulatorClient) SimulateDeck(ctx context.Context, in *SimulateDeckRequest, opts ...grpc.CallOption) (*SimulateDeckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SimulateDeckResponse)
	err := c.cc.Invoke(ctx, ManabaseSimulator_SimulateDeck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manabaseSimulatorClient) ValidateDeckList(ctx context.Context, in *ValidateDeckListRequest, opts ...grpc.CallOption) (*ValidateDeckListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateDeckListResponse)
	err := c.cc.Invoke(ctx, ManabaseSimulator_ValidateDeckList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManabaseSimulatorServer is the server API for ManabaseSimulator service.
// All implementations must embed UnimplementedManabaseSimulatorServer
// for forward compatibility.
type ManabaseSimulatorServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	// SimulateDeck simulates a deck against a given objective with the provided configuration.
	SimulateDeck(context.Context, *SimulateDeckRequest) (*SimulateDeckResponse, error)
	// ValidateDeckList validates a deckList is allowed to be played and can be simulated.
	ValidateDeckList(context.Context, *ValidateDeckListRequest) (*ValidateDeckListResponse, error)
	mustEmbedUnimplementedManabaseSimulatorServer()
}

// UnimplementedManabaseSimulatorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedManabaseSimulatorServer struct{}

func (UnimplementedManabaseSimulatorServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedManabaseSimulatorServer) SimulateDeck(context.Context, *SimulateDeckRequest) (*SimulateDeckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimulateDeck not implemented")
}
func (UnimplementedManabaseSimulatorServer) ValidateDeckList(context.Context, *ValidateDeckListRequest) (*ValidateDeckListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateDeckList not implemented")
}
func (UnimplementedManabaseSimulatorServer) mustEmbedUnimplementedManabaseSimulatorServer() {}
func (UnimplementedManabaseSimulatorServer) testEmbeddedByValue()                           {}

// UnsafeManabaseSimulatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManabaseSimulatorServer will
// result in compilation errors.
type UnsafeManabaseSimulatorServer interface {
	mustEmbedUnimplementedManabaseSimulatorServer()
}

func RegisterManabaseSimulatorServer(s grpc.ServiceRegistrar, srv ManabaseSimulatorServer) {
	// If the following call pancis, it indicates UnimplementedManabaseSimulatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ManabaseSimulator_ServiceDesc, srv)
}

func _ManabaseSimulator_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManabaseSimulatorServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManabaseSimulator_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManabaseSimulatorServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManabaseSimulator_SimulateDeck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulateDeckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManabaseSimulatorServer).SimulateDeck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManabaseSimulator_SimulateDeck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManabaseSimulatorServer).SimulateDeck(ctx, req.(*SimulateDeckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManabaseSimulator_ValidateDeckList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateDeckListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManabaseSimulatorServer).ValidateDeckList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManabaseSimulator_ValidateDeckList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManabaseSimulatorServer).ValidateDeckList(ctx, req.(*ValidateDeckListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManabaseSimulator_ServiceDesc is the grpc.ServiceDesc for ManabaseSimulator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManabaseSimulator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manabase_simulation.ManabaseSimulator",
	HandlerType: (*ManabaseSimulatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _ManabaseSimulator_Echo_Handler,
		},
		{
			MethodName: "SimulateDeck",
			Handler:    _ManabaseSimulator_SimulateDeck_Handler,
		},
		{
			MethodName: "ValidateDeckList",
			Handler:    _ManabaseSimulator_ValidateDeckList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/manabase-simulation.proto",
}
