// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContributionsServiceClient is the client API for ContributionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContributionsServiceClient interface {
	ListAuthorContributions(ctx context.Context, in *ListAuthorContributionsRequest, opts ...grpc.CallOption) (*AuthorContributionList, error)
	ListAuthorRanks(ctx context.Context, in *ListAuthorRanksRequest, opts ...grpc.CallOption) (*AuthorRankList, error)
	ListPendingRewards(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PendingRewardList, error)
	ListClaimedRewards(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClaimedRewardList, error)
	ClaimReward(ctx context.Context, in *ClaimRewardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListDonationRecipients(ctx context.Context, in *ListDonationRecipientsRequest, opts ...grpc.CallOption) (*DonationRecipientList, error)
	GetDonationStats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DonationStats, error)
	ListIncomingDonations(ctx context.Context, in *ListIncomingDonationsRequest, opts ...grpc.CallOption) (*IncomingDonationList, error)
}

type contributionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContributionsServiceClient(cc grpc.ClientConnInterface) ContributionsServiceClient {
	return &contributionsServiceClient{cc}
}

func (c *contributionsServiceClient) ListAuthorContributions(ctx context.Context, in *ListAuthorContributionsRequest, opts ...grpc.CallOption) (*AuthorContributionList, error) {
	out := new(AuthorContributionList)
	err := c.cc.Invoke(ctx, "/rsk.ContributionsService/ListAuthorContributions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contributionsServiceClient) ListAuthorRanks(ctx context.Context, in *ListAuthorRanksRequest, opts ...grpc.CallOption) (*AuthorRankList, error) {
	out := new(AuthorRankList)
	err := c.cc.Invoke(ctx, "/rsk.ContributionsService/ListAuthorRanks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contributionsServiceClient) ListPendingRewards(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PendingRewardList, error) {
	out := new(PendingRewardList)
	err := c.cc.Invoke(ctx, "/rsk.ContributionsService/ListPendingRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contributionsServiceClient) ListClaimedRewards(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClaimedRewardList, error) {
	out := new(ClaimedRewardList)
	err := c.cc.Invoke(ctx, "/rsk.ContributionsService/ListClaimedRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contributionsServiceClient) ClaimReward(ctx context.Context, in *ClaimRewardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/rsk.ContributionsService/ClaimReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contributionsServiceClient) ListDonationRecipients(ctx context.Context, in *ListDonationRecipientsRequest, opts ...grpc.CallOption) (*DonationRecipientList, error) {
	out := new(DonationRecipientList)
	err := c.cc.Invoke(ctx, "/rsk.ContributionsService/ListDonationRecipients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contributionsServiceClient) GetDonationStats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DonationStats, error) {
	out := new(DonationStats)
	err := c.cc.Invoke(ctx, "/rsk.ContributionsService/GetDonationStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contributionsServiceClient) ListIncomingDonations(ctx context.Context, in *ListIncomingDonationsRequest, opts ...grpc.CallOption) (*IncomingDonationList, error) {
	out := new(IncomingDonationList)
	err := c.cc.Invoke(ctx, "/rsk.ContributionsService/ListIncomingDonations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContributionsServiceServer is the server API for ContributionsService service.
// All implementations should embed UnimplementedContributionsServiceServer
// for forward compatibility
type ContributionsServiceServer interface {
	ListAuthorContributions(context.Context, *ListAuthorContributionsRequest) (*AuthorContributionList, error)
	ListAuthorRanks(context.Context, *ListAuthorRanksRequest) (*AuthorRankList, error)
	ListPendingRewards(context.Context, *emptypb.Empty) (*PendingRewardList, error)
	ListClaimedRewards(context.Context, *emptypb.Empty) (*ClaimedRewardList, error)
	ClaimReward(context.Context, *ClaimRewardRequest) (*emptypb.Empty, error)
	ListDonationRecipients(context.Context, *ListDonationRecipientsRequest) (*DonationRecipientList, error)
	GetDonationStats(context.Context, *emptypb.Empty) (*DonationStats, error)
	ListIncomingDonations(context.Context, *ListIncomingDonationsRequest) (*IncomingDonationList, error)
}

// UnimplementedContributionsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedContributionsServiceServer struct {
}

func (UnimplementedContributionsServiceServer) ListAuthorContributions(context.Context, *ListAuthorContributionsRequest) (*AuthorContributionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthorContributions not implemented")
}
func (UnimplementedContributionsServiceServer) ListAuthorRanks(context.Context, *ListAuthorRanksRequest) (*AuthorRankList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthorRanks not implemented")
}
func (UnimplementedContributionsServiceServer) ListPendingRewards(context.Context, *emptypb.Empty) (*PendingRewardList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPendingRewards not implemented")
}
func (UnimplementedContributionsServiceServer) ListClaimedRewards(context.Context, *emptypb.Empty) (*ClaimedRewardList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClaimedRewards not implemented")
}
func (UnimplementedContributionsServiceServer) ClaimReward(context.Context, *ClaimRewardRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimReward not implemented")
}
func (UnimplementedContributionsServiceServer) ListDonationRecipients(context.Context, *ListDonationRecipientsRequest) (*DonationRecipientList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDonationRecipients not implemented")
}
func (UnimplementedContributionsServiceServer) GetDonationStats(context.Context, *emptypb.Empty) (*DonationStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonationStats not implemented")
}
func (UnimplementedContributionsServiceServer) ListIncomingDonations(context.Context, *ListIncomingDonationsRequest) (*IncomingDonationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIncomingDonations not implemented")
}

// UnsafeContributionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContributionsServiceServer will
// result in compilation errors.
type UnsafeContributionsServiceServer interface {
	mustEmbedUnimplementedContributionsServiceServer()
}

func RegisterContributionsServiceServer(s grpc.ServiceRegistrar, srv ContributionsServiceServer) {
	s.RegisterService(&ContributionsService_ServiceDesc, srv)
}

func _ContributionsService_ListAuthorContributions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthorContributionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContributionsServiceServer).ListAuthorContributions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsk.ContributionsService/ListAuthorContributions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContributionsServiceServer).ListAuthorContributions(ctx, req.(*ListAuthorContributionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContributionsService_ListAuthorRanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthorRanksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContributionsServiceServer).ListAuthorRanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsk.ContributionsService/ListAuthorRanks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContributionsServiceServer).ListAuthorRanks(ctx, req.(*ListAuthorRanksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContributionsService_ListPendingRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContributionsServiceServer).ListPendingRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsk.ContributionsService/ListPendingRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContributionsServiceServer).ListPendingRewards(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContributionsService_ListClaimedRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContributionsServiceServer).ListClaimedRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsk.ContributionsService/ListClaimedRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContributionsServiceServer).ListClaimedRewards(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContributionsService_ClaimReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContributionsServiceServer).ClaimReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsk.ContributionsService/ClaimReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContributionsServiceServer).ClaimReward(ctx, req.(*ClaimRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContributionsService_ListDonationRecipients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDonationRecipientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContributionsServiceServer).ListDonationRecipients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsk.ContributionsService/ListDonationRecipients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContributionsServiceServer).ListDonationRecipients(ctx, req.(*ListDonationRecipientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContributionsService_GetDonationStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContributionsServiceServer).GetDonationStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsk.ContributionsService/GetDonationStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContributionsServiceServer).GetDonationStats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContributionsService_ListIncomingDonations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIncomingDonationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContributionsServiceServer).ListIncomingDonations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsk.ContributionsService/ListIncomingDonations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContributionsServiceServer).ListIncomingDonations(ctx, req.(*ListIncomingDonationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContributionsService_ServiceDesc is the grpc.ServiceDesc for ContributionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContributionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rsk.ContributionsService",
	HandlerType: (*ContributionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAuthorContributions",
			Handler:    _ContributionsService_ListAuthorContributions_Handler,
		},
		{
			MethodName: "ListAuthorRanks",
			Handler:    _ContributionsService_ListAuthorRanks_Handler,
		},
		{
			MethodName: "ListPendingRewards",
			Handler:    _ContributionsService_ListPendingRewards_Handler,
		},
		{
			MethodName: "ListClaimedRewards",
			Handler:    _ContributionsService_ListClaimedRewards_Handler,
		},
		{
			MethodName: "ClaimReward",
			Handler:    _ContributionsService_ClaimReward_Handler,
		},
		{
			MethodName: "ListDonationRecipients",
			Handler:    _ContributionsService_ListDonationRecipients_Handler,
		},
		{
			MethodName: "GetDonationStats",
			Handler:    _ContributionsService_GetDonationStats_Handler,
		},
		{
			MethodName: "ListIncomingDonations",
			Handler:    _ContributionsService_ListIncomingDonations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contribution.proto",
}
