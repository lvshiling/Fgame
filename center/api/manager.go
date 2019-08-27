package api

import (
	"context"
	"fgame/fgame/center/center"
	centerpb "fgame/fgame/center/pb"
)

//服务器服务
type ServerManagerServer struct {
	m center.ServerManager
}

func (ss *ServerManagerServer) Register(ctx context.Context, req *centerpb.ServerRegisterRequest) (res *centerpb.ServerRegisterResponse, err error) {
	res, err = ss.m.Register(ctx, req)
	return
}

func (ss *ServerManagerServer) Unregister(ctx context.Context, req *centerpb.ServerUnregisterRequest) (res *centerpb.ServerUnregisterResponse, err error) {
	res, err = ss.m.Unregister(ctx, req)
	return
}

func (ss *ServerManagerServer) GetCrossList(ctx context.Context, req *centerpb.ServerCrossListRequest) (res *centerpb.ServerCrossListResponse, err error) {
	res, err = ss.m.GetCrossList(ctx, req)
	return
}

func (ss *ServerManagerServer) GetServerList(ctx context.Context, req *centerpb.ServerInfoListRequest) (res *centerpb.ServerInfoListResponse, err error) {
	res, err = ss.m.GetServerList(ctx, req)
	return
}

func (ss *ServerManagerServer) Refresh(ctx context.Context, req *centerpb.RefreshServerInfoListRequest) (res *centerpb.RefreshServerInfoListResponse, err error) {
	res, err = ss.m.Refresh(ctx, req)
	return
}

func (ss *ServerManagerServer) RefreshSDK(ctx context.Context, req *centerpb.RefreshSDKListRequest) (res *centerpb.RefreshSDKListResponse, err error) {
	res, err = ss.m.RefreshSDK(ctx, req)
	return
}

func (ss *ServerManagerServer) GetServerInfo(ctx context.Context, req *centerpb.ServerInfoRequest) (res *centerpb.ServerInfoResponse, err error) {
	res, err = ss.m.GetServerInfo(ctx, req)
	return
}

func (ss *ServerManagerServer) GetServerInfoByPlatform(ctx context.Context, req *centerpb.ServerInfoByPlatformRequest) (res *centerpb.ServerInfoByPlatformResponse, err error) {
	res, err = ss.m.GetServerInfoByPlatform(ctx, req)
	return
}

func (ss *ServerManagerServer) RefreshMarryPrice(ctx context.Context, req *centerpb.RefreshMarryPriceListRequest) (res *centerpb.RefreshMarryPriceListResponse, err error) {
	res, err = ss.m.RefreshMarryPrice(ctx, req)
	return
}

func NewServerManagerServer(centerServer *center.CenterServer) *ServerManagerServer {
	ss := &ServerManagerServer{
		m: centerServer,
	}
	return ss
}
