// Code generated by copy-client.sh. DO NOT EDIT.
// version: v1.0.0
// source: https://github.com/cybercryptio/d1-service-generic.git
// commit: 88afaccef27c4ea1feb29dccf1d2a5c3866db309

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authz

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

// AuthzClient is the client API for Authz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthzClient interface {
	//*
	// Returns a list of groups with access to the specified object.
	// This call can fail if the auth storage cannot be reached, in which case an error is returned.
	// The calling user has to be authenticated and authorized to access the object in order to get the object permissions.
	// Requires the scope `OBJECTPERMISSIONS`.
	GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error)
	//*
	// Adds a group to the access list of the specified object.
	// This call can fail if the caller does not have access to the object, if the target group does not exist, or if the auth storage cannot be reached.
	// In these cases, an error is returned.
	// Requires the scope `OBJECTPERMISSIONS`.
	AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*AddPermissionResponse, error)
	//*
	// Removes a group from the access list of the specified object.
	// This call can fail if the caller does not have access to the object or if the auth storage cannot reached.
	// In these cases, an error is returned.
	// Requires the scope `OBJECTPERMISSIONS`.
	RemovePermission(ctx context.Context, in *RemovePermissionRequest, opts ...grpc.CallOption) (*RemovePermissionResponse, error)
}

type authzClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzClient(cc grpc.ClientConnInterface) AuthzClient {
	return &authzClient{cc}
}

func (c *authzClient) GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error) {
	out := new(GetPermissionsResponse)
	err := c.cc.Invoke(ctx, "/d1.authz.Authz/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authzClient) AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*AddPermissionResponse, error) {
	out := new(AddPermissionResponse)
	err := c.cc.Invoke(ctx, "/d1.authz.Authz/AddPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authzClient) RemovePermission(ctx context.Context, in *RemovePermissionRequest, opts ...grpc.CallOption) (*RemovePermissionResponse, error) {
	out := new(RemovePermissionResponse)
	err := c.cc.Invoke(ctx, "/d1.authz.Authz/RemovePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzServer is the server API for Authz service.
// All implementations must embed UnimplementedAuthzServer
// for forward compatibility
type AuthzServer interface {
	//*
	// Returns a list of groups with access to the specified object.
	// This call can fail if the auth storage cannot be reached, in which case an error is returned.
	// The calling user has to be authenticated and authorized to access the object in order to get the object permissions.
	// Requires the scope `OBJECTPERMISSIONS`.
	GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error)
	//*
	// Adds a group to the access list of the specified object.
	// This call can fail if the caller does not have access to the object, if the target group does not exist, or if the auth storage cannot be reached.
	// In these cases, an error is returned.
	// Requires the scope `OBJECTPERMISSIONS`.
	AddPermission(context.Context, *AddPermissionRequest) (*AddPermissionResponse, error)
	//*
	// Removes a group from the access list of the specified object.
	// This call can fail if the caller does not have access to the object or if the auth storage cannot reached.
	// In these cases, an error is returned.
	// Requires the scope `OBJECTPERMISSIONS`.
	RemovePermission(context.Context, *RemovePermissionRequest) (*RemovePermissionResponse, error)
	mustEmbedUnimplementedAuthzServer()
}

// UnimplementedAuthzServer must be embedded to have forward compatible implementations.
type UnimplementedAuthzServer struct {
}

func (UnimplementedAuthzServer) GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (UnimplementedAuthzServer) AddPermission(context.Context, *AddPermissionRequest) (*AddPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermission not implemented")
}
func (UnimplementedAuthzServer) RemovePermission(context.Context, *RemovePermissionRequest) (*RemovePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermission not implemented")
}
func (UnimplementedAuthzServer) mustEmbedUnimplementedAuthzServer() {}

// UnsafeAuthzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthzServer will
// result in compilation errors.
type UnsafeAuthzServer interface {
	mustEmbedUnimplementedAuthzServer()
}

func RegisterAuthzServer(s grpc.ServiceRegistrar, srv AuthzServer) {
	s.RegisterService(&Authz_ServiceDesc, srv)
}

func _Authz_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.authz.Authz/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServer).GetPermissions(ctx, req.(*GetPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authz_AddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServer).AddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.authz.Authz/AddPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServer).AddPermission(ctx, req.(*AddPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authz_RemovePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServer).RemovePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.authz.Authz/RemovePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServer).RemovePermission(ctx, req.(*RemovePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authz_ServiceDesc is the grpc.ServiceDesc for Authz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "d1.authz.Authz",
	HandlerType: (*AuthzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPermissions",
			Handler:    _Authz_GetPermissions_Handler,
		},
		{
			MethodName: "AddPermission",
			Handler:    _Authz_AddPermission_Handler,
		},
		{
			MethodName: "RemovePermission",
			Handler:    _Authz_RemovePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authz.proto",
}
