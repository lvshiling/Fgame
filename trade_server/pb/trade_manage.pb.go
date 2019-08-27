// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trade_manage.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	trade_manage.proto

It has these top-level messages:
	GlobalTradeItem
	TradeUploadItemRequest
	TradeUploadItemResponse
	TradeWithdrawItemRequest
	TradeWithdrawItemResponse
	TradeItemListRequest
	TradeItemListResponse
	TradeItemRequest
	TradeItemResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// 全局商品列表
type GlobalTradeItem struct {
	GlobalTradeId     int64  `protobuf:"varint,1,opt,name=globalTradeId" json:"globalTradeId,omitempty"`
	Platform          int32  `protobuf:"varint,2,opt,name=platform" json:"platform,omitempty"`
	ServerId          int32  `protobuf:"varint,3,opt,name=serverId" json:"serverId,omitempty"`
	TradeId           int64  `protobuf:"varint,4,opt,name=tradeId" json:"tradeId,omitempty"`
	PlayerId          int64  `protobuf:"varint,5,opt,name=playerId" json:"playerId,omitempty"`
	PlayerName        string `protobuf:"bytes,6,opt,name=playerName" json:"playerName,omitempty"`
	ItemId            int32  `protobuf:"varint,7,opt,name=itemId" json:"itemId,omitempty"`
	ItemNum           int32  `protobuf:"varint,8,opt,name=itemNum" json:"itemNum,omitempty"`
	Gold              int32  `protobuf:"varint,9,opt,name=gold" json:"gold,omitempty"`
	PropertyData      string `protobuf:"bytes,10,opt,name=propertyData" json:"propertyData,omitempty"`
	Level             int32  `protobuf:"varint,11,opt,name=level" json:"level,omitempty"`
	BuyPlayerPlatform int32  `protobuf:"varint,12,opt,name=buyPlayerPlatform" json:"buyPlayerPlatform,omitempty"`
	BuyPlayerServerId int32  `protobuf:"varint,13,opt,name=buyPlayerServerId" json:"buyPlayerServerId,omitempty"`
	BuyPlayerId       int64  `protobuf:"varint,14,opt,name=buyPlayerId" json:"buyPlayerId,omitempty"`
	BuyPlayerName     string `protobuf:"bytes,15,opt,name=buyPlayerName" json:"buyPlayerName,omitempty"`
	UpdateTime        int64  `protobuf:"varint,16,opt,name=updateTime" json:"updateTime,omitempty"`
	CreateTime        int64  `protobuf:"varint,17,opt,name=createTime" json:"createTime,omitempty"`
	DeleteTime        int64  `protobuf:"varint,18,opt,name=deleteTime" json:"deleteTime,omitempty"`
}

func (m *GlobalTradeItem) Reset()                    { *m = GlobalTradeItem{} }
func (m *GlobalTradeItem) String() string            { return proto.CompactTextString(m) }
func (*GlobalTradeItem) ProtoMessage()               {}
func (*GlobalTradeItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GlobalTradeItem) GetGlobalTradeId() int64 {
	if m != nil {
		return m.GlobalTradeId
	}
	return 0
}

func (m *GlobalTradeItem) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *GlobalTradeItem) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *GlobalTradeItem) GetTradeId() int64 {
	if m != nil {
		return m.TradeId
	}
	return 0
}

func (m *GlobalTradeItem) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *GlobalTradeItem) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *GlobalTradeItem) GetItemId() int32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *GlobalTradeItem) GetItemNum() int32 {
	if m != nil {
		return m.ItemNum
	}
	return 0
}

func (m *GlobalTradeItem) GetGold() int32 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *GlobalTradeItem) GetPropertyData() string {
	if m != nil {
		return m.PropertyData
	}
	return ""
}

func (m *GlobalTradeItem) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *GlobalTradeItem) GetBuyPlayerPlatform() int32 {
	if m != nil {
		return m.BuyPlayerPlatform
	}
	return 0
}

func (m *GlobalTradeItem) GetBuyPlayerServerId() int32 {
	if m != nil {
		return m.BuyPlayerServerId
	}
	return 0
}

func (m *GlobalTradeItem) GetBuyPlayerId() int64 {
	if m != nil {
		return m.BuyPlayerId
	}
	return 0
}

func (m *GlobalTradeItem) GetBuyPlayerName() string {
	if m != nil {
		return m.BuyPlayerName
	}
	return ""
}

func (m *GlobalTradeItem) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *GlobalTradeItem) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *GlobalTradeItem) GetDeleteTime() int64 {
	if m != nil {
		return m.DeleteTime
	}
	return 0
}

// 上架物品
type TradeUploadItemRequest struct {
	Platform     int32  `protobuf:"varint,1,opt,name=platform" json:"platform,omitempty"`
	ServerId     int32  `protobuf:"varint,2,opt,name=serverId" json:"serverId,omitempty"`
	TradeId      int64  `protobuf:"varint,3,opt,name=tradeId" json:"tradeId,omitempty"`
	PlayerId     int64  `protobuf:"varint,4,opt,name=playerId" json:"playerId,omitempty"`
	PlayerName   string `protobuf:"bytes,5,opt,name=playerName" json:"playerName,omitempty"`
	ItemId       int32  `protobuf:"varint,6,opt,name=itemId" json:"itemId,omitempty"`
	ItemNum      int32  `protobuf:"varint,7,opt,name=itemNum" json:"itemNum,omitempty"`
	Gold         int32  `protobuf:"varint,8,opt,name=gold" json:"gold,omitempty"`
	PropertyData string `protobuf:"bytes,9,opt,name=propertyData" json:"propertyData,omitempty"`
	Level        int32  `protobuf:"varint,10,opt,name=level" json:"level,omitempty"`
}

func (m *TradeUploadItemRequest) Reset()                    { *m = TradeUploadItemRequest{} }
func (m *TradeUploadItemRequest) String() string            { return proto.CompactTextString(m) }
func (*TradeUploadItemRequest) ProtoMessage()               {}
func (*TradeUploadItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TradeUploadItemRequest) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *TradeUploadItemRequest) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *TradeUploadItemRequest) GetTradeId() int64 {
	if m != nil {
		return m.TradeId
	}
	return 0
}

func (m *TradeUploadItemRequest) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *TradeUploadItemRequest) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *TradeUploadItemRequest) GetItemId() int32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *TradeUploadItemRequest) GetItemNum() int32 {
	if m != nil {
		return m.ItemNum
	}
	return 0
}

func (m *TradeUploadItemRequest) GetGold() int32 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *TradeUploadItemRequest) GetPropertyData() string {
	if m != nil {
		return m.PropertyData
	}
	return ""
}

func (m *TradeUploadItemRequest) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

// 上架物品回复
type TradeUploadItemResponse struct {
	TradeItem *GlobalTradeItem `protobuf:"bytes,1,opt,name=tradeItem" json:"tradeItem,omitempty"`
}

func (m *TradeUploadItemResponse) Reset()                    { *m = TradeUploadItemResponse{} }
func (m *TradeUploadItemResponse) String() string            { return proto.CompactTextString(m) }
func (*TradeUploadItemResponse) ProtoMessage()               {}
func (*TradeUploadItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TradeUploadItemResponse) GetTradeItem() *GlobalTradeItem {
	if m != nil {
		return m.TradeItem
	}
	return nil
}

// 下架物品
type TradeWithdrawItemRequest struct {
	Platform      int32 `protobuf:"varint,1,opt,name=platform" json:"platform,omitempty"`
	ServerId      int32 `protobuf:"varint,2,opt,name=serverId" json:"serverId,omitempty"`
	GlobalTradeId int64 `protobuf:"varint,3,opt,name=globalTradeId" json:"globalTradeId,omitempty"`
}

func (m *TradeWithdrawItemRequest) Reset()                    { *m = TradeWithdrawItemRequest{} }
func (m *TradeWithdrawItemRequest) String() string            { return proto.CompactTextString(m) }
func (*TradeWithdrawItemRequest) ProtoMessage()               {}
func (*TradeWithdrawItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TradeWithdrawItemRequest) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *TradeWithdrawItemRequest) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *TradeWithdrawItemRequest) GetGlobalTradeId() int64 {
	if m != nil {
		return m.GlobalTradeId
	}
	return 0
}

// 下架物品回复
type TradeWithdrawItemResponse struct {
	Platform      int32 `protobuf:"varint,1,opt,name=platform" json:"platform,omitempty"`
	ServerId      int32 `protobuf:"varint,2,opt,name=serverId" json:"serverId,omitempty"`
	GlobalTradeId int64 `protobuf:"varint,3,opt,name=globalTradeId" json:"globalTradeId,omitempty"`
}

func (m *TradeWithdrawItemResponse) Reset()                    { *m = TradeWithdrawItemResponse{} }
func (m *TradeWithdrawItemResponse) String() string            { return proto.CompactTextString(m) }
func (*TradeWithdrawItemResponse) ProtoMessage()               {}
func (*TradeWithdrawItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TradeWithdrawItemResponse) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *TradeWithdrawItemResponse) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *TradeWithdrawItemResponse) GetGlobalTradeId() int64 {
	if m != nil {
		return m.GlobalTradeId
	}
	return 0
}

// 获取交易列表
type TradeItemListRequest struct {
	Platform int32 `protobuf:"varint,1,opt,name=platform" json:"platform,omitempty"`
	ServerId int32 `protobuf:"varint,2,opt,name=serverId" json:"serverId,omitempty"`
}

func (m *TradeItemListRequest) Reset()                    { *m = TradeItemListRequest{} }
func (m *TradeItemListRequest) String() string            { return proto.CompactTextString(m) }
func (*TradeItemListRequest) ProtoMessage()               {}
func (*TradeItemListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TradeItemListRequest) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *TradeItemListRequest) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

// 全局商品列表
type TradeItemListResponse struct {
	TradeItemList []*GlobalTradeItem `protobuf:"bytes,1,rep,name=tradeItemList" json:"tradeItemList,omitempty"`
}

func (m *TradeItemListResponse) Reset()                    { *m = TradeItemListResponse{} }
func (m *TradeItemListResponse) String() string            { return proto.CompactTextString(m) }
func (*TradeItemListResponse) ProtoMessage()               {}
func (*TradeItemListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TradeItemListResponse) GetTradeItemList() []*GlobalTradeItem {
	if m != nil {
		return m.TradeItemList
	}
	return nil
}

// 交易物品
type TradeItemRequest struct {
	Platform      int32  `protobuf:"varint,1,opt,name=platform" json:"platform,omitempty"`
	ServerId      int32  `protobuf:"varint,2,opt,name=serverId" json:"serverId,omitempty"`
	BuyPlayerId   int64  `protobuf:"varint,4,opt,name=buyPlayerId" json:"buyPlayerId,omitempty"`
	BuyPlayerName string `protobuf:"bytes,5,opt,name=buyPlayerName" json:"buyPlayerName,omitempty"`
	GlobalTradeId int64  `protobuf:"varint,3,opt,name=globalTradeId" json:"globalTradeId,omitempty"`
}

func (m *TradeItemRequest) Reset()                    { *m = TradeItemRequest{} }
func (m *TradeItemRequest) String() string            { return proto.CompactTextString(m) }
func (*TradeItemRequest) ProtoMessage()               {}
func (*TradeItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TradeItemRequest) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *TradeItemRequest) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *TradeItemRequest) GetBuyPlayerId() int64 {
	if m != nil {
		return m.BuyPlayerId
	}
	return 0
}

func (m *TradeItemRequest) GetBuyPlayerName() string {
	if m != nil {
		return m.BuyPlayerName
	}
	return ""
}

func (m *TradeItemRequest) GetGlobalTradeId() int64 {
	if m != nil {
		return m.GlobalTradeId
	}
	return 0
}

// 交易商品回复
type TradeItemResponse struct {
	TradeItem *GlobalTradeItem `protobuf:"bytes,6,opt,name=tradeItem" json:"tradeItem,omitempty"`
}

func (m *TradeItemResponse) Reset()                    { *m = TradeItemResponse{} }
func (m *TradeItemResponse) String() string            { return proto.CompactTextString(m) }
func (*TradeItemResponse) ProtoMessage()               {}
func (*TradeItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TradeItemResponse) GetTradeItem() *GlobalTradeItem {
	if m != nil {
		return m.TradeItem
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalTradeItem)(nil), "pb.GlobalTradeItem")
	proto.RegisterType((*TradeUploadItemRequest)(nil), "pb.TradeUploadItemRequest")
	proto.RegisterType((*TradeUploadItemResponse)(nil), "pb.TradeUploadItemResponse")
	proto.RegisterType((*TradeWithdrawItemRequest)(nil), "pb.TradeWithdrawItemRequest")
	proto.RegisterType((*TradeWithdrawItemResponse)(nil), "pb.TradeWithdrawItemResponse")
	proto.RegisterType((*TradeItemListRequest)(nil), "pb.TradeItemListRequest")
	proto.RegisterType((*TradeItemListResponse)(nil), "pb.TradeItemListResponse")
	proto.RegisterType((*TradeItemRequest)(nil), "pb.TradeItemRequest")
	proto.RegisterType((*TradeItemResponse)(nil), "pb.TradeItemResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TradeManage service

type TradeManageClient interface {
	// 获取商品列表
	GetTradeList(ctx context.Context, in *TradeItemListRequest, opts ...grpc.CallOption) (*TradeItemListResponse, error)
	// 上架
	Upload(ctx context.Context, in *TradeUploadItemRequest, opts ...grpc.CallOption) (*TradeUploadItemResponse, error)
	// 下架
	Withdraw(ctx context.Context, in *TradeWithdrawItemRequest, opts ...grpc.CallOption) (*TradeWithdrawItemResponse, error)
	// 交易
	Trade(ctx context.Context, in *TradeItemRequest, opts ...grpc.CallOption) (*TradeItemResponse, error)
}

type tradeManageClient struct {
	cc *grpc.ClientConn
}

func NewTradeManageClient(cc *grpc.ClientConn) TradeManageClient {
	return &tradeManageClient{cc}
}

func (c *tradeManageClient) GetTradeList(ctx context.Context, in *TradeItemListRequest, opts ...grpc.CallOption) (*TradeItemListResponse, error) {
	out := new(TradeItemListResponse)
	err := grpc.Invoke(ctx, "/pb.TradeManage/GetTradeList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeManageClient) Upload(ctx context.Context, in *TradeUploadItemRequest, opts ...grpc.CallOption) (*TradeUploadItemResponse, error) {
	out := new(TradeUploadItemResponse)
	err := grpc.Invoke(ctx, "/pb.TradeManage/Upload", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeManageClient) Withdraw(ctx context.Context, in *TradeWithdrawItemRequest, opts ...grpc.CallOption) (*TradeWithdrawItemResponse, error) {
	out := new(TradeWithdrawItemResponse)
	err := grpc.Invoke(ctx, "/pb.TradeManage/Withdraw", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeManageClient) Trade(ctx context.Context, in *TradeItemRequest, opts ...grpc.CallOption) (*TradeItemResponse, error) {
	out := new(TradeItemResponse)
	err := grpc.Invoke(ctx, "/pb.TradeManage/Trade", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TradeManage service

type TradeManageServer interface {
	// 获取商品列表
	GetTradeList(context.Context, *TradeItemListRequest) (*TradeItemListResponse, error)
	// 上架
	Upload(context.Context, *TradeUploadItemRequest) (*TradeUploadItemResponse, error)
	// 下架
	Withdraw(context.Context, *TradeWithdrawItemRequest) (*TradeWithdrawItemResponse, error)
	// 交易
	Trade(context.Context, *TradeItemRequest) (*TradeItemResponse, error)
}

func RegisterTradeManageServer(s *grpc.Server, srv TradeManageServer) {
	s.RegisterService(&_TradeManage_serviceDesc, srv)
}

func _TradeManage_GetTradeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeItemListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeManageServer).GetTradeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeManage/GetTradeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeManageServer).GetTradeList(ctx, req.(*TradeItemListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeManage_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeUploadItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeManageServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeManage/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeManageServer).Upload(ctx, req.(*TradeUploadItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeManage_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeWithdrawItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeManageServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeManage/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeManageServer).Withdraw(ctx, req.(*TradeWithdrawItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeManage_Trade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeManageServer).Trade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TradeManage/Trade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeManageServer).Trade(ctx, req.(*TradeItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TradeManage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TradeManage",
	HandlerType: (*TradeManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTradeList",
			Handler:    _TradeManage_GetTradeList_Handler,
		},
		{
			MethodName: "Upload",
			Handler:    _TradeManage_Upload_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _TradeManage_Withdraw_Handler,
		},
		{
			MethodName: "Trade",
			Handler:    _TradeManage_Trade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trade_manage.proto",
}

func init() { proto.RegisterFile("trade_manage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x8e, 0xd2, 0x40,
	0x14, 0xb6, 0x40, 0xbb, 0x70, 0x00, 0x77, 0x39, 0xb2, 0xeb, 0x2c, 0xfe, 0x84, 0x34, 0x5e, 0x70,
	0x61, 0x48, 0x5c, 0x13, 0x13, 0xaf, 0xfd, 0xd9, 0x90, 0xac, 0x64, 0x53, 0xd7, 0x78, 0x69, 0x06,
	0x67, 0x44, 0x92, 0x76, 0x5b, 0xcb, 0xb0, 0xca, 0xbb, 0x78, 0xed, 0x33, 0xf8, 0x0a, 0xbe, 0x95,
	0x99, 0xd3, 0x4e, 0x69, 0xa1, 0x10, 0x92, 0xd5, 0xbb, 0x9e, 0xef, 0x3b, 0x73, 0xe6, 0xcc, 0x37,
	0xdf, 0x9c, 0x02, 0xaa, 0x98, 0x0b, 0xf9, 0x29, 0xe0, 0xd7, 0x7c, 0x2a, 0x87, 0x51, 0x1c, 0xaa,
	0x10, 0x2b, 0xd1, 0xc4, 0xfd, 0x53, 0x83, 0xc3, 0x73, 0x3f, 0x9c, 0x70, 0xff, 0x4a, 0x27, 0x8c,
	0x94, 0x0c, 0xf0, 0x09, 0xb4, 0xa7, 0x39, 0x48, 0x30, 0xab, 0x6f, 0x0d, 0xaa, 0x5e, 0x11, 0xc4,
	0x1e, 0xd4, 0x23, 0x9f, 0xab, 0x2f, 0x61, 0x1c, 0xb0, 0x4a, 0xdf, 0x1a, 0xd8, 0x5e, 0x16, 0x6b,
	0x6e, 0x2e, 0xe3, 0x1b, 0x19, 0x8f, 0x04, 0xab, 0x26, 0x9c, 0x89, 0x91, 0xc1, 0x81, 0x4a, 0xeb,
	0xd6, 0xa8, 0xae, 0x09, 0xd3, 0x8a, 0x4b, 0x5a, 0x65, 0x13, 0x95, 0xc5, 0xf8, 0x18, 0x20, 0xf9,
	0x1e, 0xf3, 0x40, 0x32, 0xa7, 0x6f, 0x0d, 0x1a, 0x5e, 0x0e, 0xc1, 0x13, 0x70, 0x66, 0x4a, 0x06,
	0x23, 0xc1, 0x0e, 0x68, 0xbf, 0x34, 0xd2, 0xbb, 0xe9, 0xaf, 0xf1, 0x22, 0x60, 0x75, 0x22, 0x4c,
	0x88, 0x08, 0xb5, 0x69, 0xe8, 0x0b, 0xd6, 0x20, 0x98, 0xbe, 0xd1, 0x85, 0x56, 0x14, 0x87, 0x91,
	0x8c, 0xd5, 0xf2, 0x35, 0x57, 0x9c, 0x01, 0xed, 0x53, 0xc0, 0xb0, 0x0b, 0xb6, 0x2f, 0x6f, 0xa4,
	0xcf, 0x9a, 0xb4, 0x30, 0x09, 0xf0, 0x29, 0x74, 0x26, 0x8b, 0xe5, 0x25, 0x35, 0x74, 0x69, 0x64,
	0x69, 0x51, 0xc6, 0x26, 0x51, 0xc8, 0x7e, 0x6f, 0x84, 0x6a, 0xaf, 0x65, 0x1b, 0x02, 0xfb, 0xd0,
	0xcc, 0xc0, 0x91, 0x60, 0x77, 0x49, 0x9a, 0x3c, 0xa4, 0x6f, 0x2c, 0x0b, 0x49, 0xa0, 0x43, 0x6a,
	0xbc, 0x08, 0x6a, 0x0d, 0x17, 0x91, 0xe0, 0x4a, 0x5e, 0xcd, 0x02, 0xc9, 0x8e, 0xa8, 0x4c, 0x0e,
	0xd1, 0xfc, 0xe7, 0x58, 0x1a, 0xbe, 0x93, 0xf0, 0x2b, 0x44, 0xf3, 0x42, 0xfa, 0x32, 0xe5, 0x31,
	0xe1, 0x57, 0x88, 0xfb, 0xab, 0x02, 0x27, 0xe4, 0x8e, 0x0f, 0x91, 0x1f, 0x72, 0xa1, 0xbd, 0xe4,
	0xc9, 0x6f, 0x0b, 0x39, 0x57, 0x05, 0xb3, 0x58, 0x3b, 0xcc, 0x52, 0xd9, 0x6e, 0x96, 0xea, 0x76,
	0xb3, 0xd4, 0x76, 0x9a, 0xc5, 0xde, 0x61, 0x16, 0x67, 0x9b, 0x59, 0x0e, 0xca, 0xcd, 0x52, 0xdf,
	0x61, 0x96, 0xc6, 0x2e, 0xb3, 0x40, 0xce, 0x2c, 0xee, 0x05, 0xdc, 0xdf, 0xd0, 0x69, 0x1e, 0x85,
	0xd7, 0x73, 0x89, 0xcf, 0xa0, 0xa1, 0xcc, 0x43, 0x24, 0xa5, 0x9a, 0x67, 0xf7, 0x86, 0xd1, 0x64,
	0xb8, 0xf6, 0x46, 0xbd, 0x55, 0x96, 0xfb, 0x03, 0x18, 0xe1, 0x1f, 0x67, 0xea, 0xab, 0x88, 0xf9,
	0xf7, 0x7f, 0xa1, 0xfb, 0xc6, 0x08, 0xa8, 0x96, 0x8c, 0x00, 0x77, 0x09, 0xa7, 0x25, 0x3b, 0xa7,
	0x27, 0xf9, 0xbf, 0x5b, 0x8f, 0xa1, 0x9b, 0x89, 0x71, 0x31, 0x9b, 0xab, 0x5b, 0x1e, 0xd8, 0xf5,
	0xe0, 0x78, 0xad, 0x5e, 0x7a, 0x8c, 0x97, 0xd0, 0x56, 0x79, 0x82, 0x59, 0xfd, 0xea, 0xb6, 0x4b,
	0x29, 0x66, 0xba, 0xbf, 0x2d, 0x38, 0x5a, 0x91, 0xb7, 0xbc, 0x91, 0xb5, 0x21, 0x50, 0xdb, 0x63,
	0x08, 0xd8, 0x65, 0x43, 0x60, 0x3f, 0x79, 0xdf, 0x42, 0x27, 0xd7, 0x79, 0x99, 0x37, 0x9d, 0x7d,
	0xbc, 0x79, 0xf6, 0xb3, 0x02, 0x4d, 0x22, 0xde, 0xd1, 0x8f, 0x07, 0xdf, 0x40, 0xeb, 0x5c, 0x2a,
	0x42, 0xb4, 0x44, 0xc8, 0xf4, 0xfa, 0xb2, 0x8b, 0xec, 0x9d, 0x96, 0x30, 0x49, 0x1f, 0xee, 0x1d,
	0x7c, 0x05, 0x4e, 0xf2, 0x76, 0xb0, 0x97, 0xa5, 0x6d, 0x0c, 0x9d, 0xde, 0x83, 0x52, 0x2e, 0x2b,
	0x32, 0x82, 0xba, 0x31, 0x2e, 0x3e, 0xcc, 0x52, 0x4b, 0x5e, 0x51, 0xef, 0xd1, 0x16, 0x36, 0x2b,
	0xf5, 0x02, 0x6c, 0xa2, 0xb1, 0x5b, 0xe8, 0xda, 0xac, 0x3f, 0x5e, 0x43, 0xcd, 0xba, 0x89, 0x43,
	0x3f, 0xe2, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa8, 0xdd, 0x99, 0x9e, 0x07, 0x00,
	0x00,
}