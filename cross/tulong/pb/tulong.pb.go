// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tulong.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	tulong.proto

It has these top-level messages:
	TuLongRankInfo
	TuLongRankListRequest
	TuLongRankListResponse
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

type TuLongRankInfo struct {
	ServerId     int32  `protobuf:"varint,1,opt,name=serverId" json:"serverId,omitempty"`
	AllianceId   int64  `protobuf:"varint,2,opt,name=allianceId" json:"allianceId,omitempty"`
	AllianceName string `protobuf:"bytes,3,opt,name=allianceName" json:"allianceName,omitempty"`
	KillNum      int32  `protobuf:"varint,4,opt,name=killNum" json:"killNum,omitempty"`
}

func (m *TuLongRankInfo) Reset()                    { *m = TuLongRankInfo{} }
func (m *TuLongRankInfo) String() string            { return proto.CompactTextString(m) }
func (*TuLongRankInfo) ProtoMessage()               {}
func (*TuLongRankInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TuLongRankInfo) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *TuLongRankInfo) GetAllianceId() int64 {
	if m != nil {
		return m.AllianceId
	}
	return 0
}

func (m *TuLongRankInfo) GetAllianceName() string {
	if m != nil {
		return m.AllianceName
	}
	return ""
}

func (m *TuLongRankInfo) GetKillNum() int32 {
	if m != nil {
		return m.KillNum
	}
	return 0
}

// 获取跨服屠龙排行榜列表
type TuLongRankListRequest struct {
}

func (m *TuLongRankListRequest) Reset()                    { *m = TuLongRankListRequest{} }
func (m *TuLongRankListRequest) String() string            { return proto.CompactTextString(m) }
func (*TuLongRankListRequest) ProtoMessage()               {}
func (*TuLongRankListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 跨服屠龙排行榜列表回复
type TuLongRankListResponse struct {
	RankInfoList []*TuLongRankInfo `protobuf:"bytes,1,rep,name=rankInfoList" json:"rankInfoList,omitempty"`
}

func (m *TuLongRankListResponse) Reset()                    { *m = TuLongRankListResponse{} }
func (m *TuLongRankListResponse) String() string            { return proto.CompactTextString(m) }
func (*TuLongRankListResponse) ProtoMessage()               {}
func (*TuLongRankListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TuLongRankListResponse) GetRankInfoList() []*TuLongRankInfo {
	if m != nil {
		return m.RankInfoList
	}
	return nil
}

func init() {
	proto.RegisterType((*TuLongRankInfo)(nil), "pb.TuLongRankInfo")
	proto.RegisterType((*TuLongRankListRequest)(nil), "pb.TuLongRankListRequest")
	proto.RegisterType((*TuLongRankListResponse)(nil), "pb.TuLongRankListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TuLong service

type TuLongClient interface {
	// 获取
	GetTuLongRankList(ctx context.Context, in *TuLongRankListRequest, opts ...grpc.CallOption) (*TuLongRankListResponse, error)
}

type tuLongClient struct {
	cc *grpc.ClientConn
}

func NewTuLongClient(cc *grpc.ClientConn) TuLongClient {
	return &tuLongClient{cc}
}

func (c *tuLongClient) GetTuLongRankList(ctx context.Context, in *TuLongRankListRequest, opts ...grpc.CallOption) (*TuLongRankListResponse, error) {
	out := new(TuLongRankListResponse)
	err := grpc.Invoke(ctx, "/pb.TuLong/GetTuLongRankList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TuLong service

type TuLongServer interface {
	// 获取
	GetTuLongRankList(context.Context, *TuLongRankListRequest) (*TuLongRankListResponse, error)
}

func RegisterTuLongServer(s *grpc.Server, srv TuLongServer) {
	s.RegisterService(&_TuLong_serviceDesc, srv)
}

func _TuLong_GetTuLongRankList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TuLongRankListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TuLongServer).GetTuLongRankList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TuLong/GetTuLongRankList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TuLongServer).GetTuLongRankList(ctx, req.(*TuLongRankListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TuLong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TuLong",
	HandlerType: (*TuLongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTuLongRankList",
			Handler:    _TuLong_GetTuLongRankList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tulong.proto",
}

func init() { proto.RegisterFile("tulong.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x4d, 0x47, 0xab, 0x5e, 0x07, 0xc1, 0x0b, 0x6a, 0x9c, 0x85, 0x0c, 0x59, 0xcd, 0x6a,
	0x16, 0x15, 0x7c, 0x05, 0x19, 0x18, 0x8a, 0x04, 0x71, 0x9f, 0xb1, 0xd7, 0x32, 0x34, 0x4d, 0x62,
	0x7e, 0x7c, 0x08, 0x9f, 0x5a, 0x6c, 0xad, 0x36, 0xe2, 0xf2, 0x9c, 0x13, 0x3e, 0xbe, 0x5c, 0x28,
	0x63, 0xd2, 0xd6, 0x2c, 0x5b, 0xe7, 0x6d, 0xb4, 0x38, 0x71, 0x83, 0xf8, 0x60, 0x70, 0xfe, 0x94,
	0x7a, 0x6b, 0x96, 0x52, 0x99, 0x55, 0x67, 0x5e, 0x2d, 0x56, 0x70, 0x12, 0xc8, 0xbf, 0x93, 0xef,
	0x16, 0x9c, 0xd5, 0xac, 0x39, 0x92, 0x3f, 0x19, 0x6f, 0x01, 0x94, 0xd6, 0xa3, 0x32, 0x2f, 0xd4,
	0x2d, 0xf8, 0xa4, 0x66, 0x4d, 0x21, 0xf7, 0x1a, 0x14, 0x50, 0xee, 0xd2, 0x5c, 0xad, 0x89, 0x17,
	0x35, 0x6b, 0x4e, 0x65, 0xd6, 0x21, 0x87, 0xe3, 0xd5, 0xa8, 0xf5, 0x3c, 0xad, 0xf9, 0xe1, 0x06,
	0xbf, 0x8b, 0xe2, 0x1a, 0x2e, 0x7f, 0x5d, 0xfa, 0x31, 0x44, 0x49, 0x6f, 0x89, 0x42, 0x14, 0x8f,
	0x70, 0xf5, 0x77, 0x08, 0xce, 0x9a, 0x40, 0x78, 0x0f, 0xa5, 0xff, 0x16, 0xff, 0xea, 0x39, 0xab,
	0x8b, 0xe6, 0x6c, 0x86, 0xad, 0x1b, 0xda, 0xfc, 0x5b, 0x32, 0x7b, 0x37, 0x7b, 0x86, 0xe9, 0x76,
	0xc7, 0x1e, 0x2e, 0x1e, 0x28, 0xe6, 0x78, 0xbc, 0xc9, 0x01, 0x7b, 0x2e, 0x55, 0xf5, 0xdf, 0xb4,
	0xb5, 0x11, 0x07, 0xc3, 0x74, 0x73, 0xda, 0xbb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x06,
	0xb6, 0xee, 0x6a, 0x01, 0x00, 0x00,
}