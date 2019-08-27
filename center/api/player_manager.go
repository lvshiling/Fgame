package api

import (
	"context"
	"fgame/fgame/center/center"
	centerpb "fgame/fgame/center/pb"
)

//服务器服务
type PlayerManagerServer struct {
	m center.PlayerManager
}

func (ps *PlayerManagerServer) SyncPlayerServerInfo(ctx context.Context, req *centerpb.PlayerInfoSyncRequest) (res *centerpb.PlayerInfoSyncResponse, err error) {
	res, err = ps.m.SyncPlayerServerInfo(ctx, req)
	return
}

func (ps *PlayerManagerServer) GetPlayerServerList(ctx context.Context, req *centerpb.PlayerListRequest) (res *centerpb.PlayerListResponse, err error) {
	res, err = ps.m.GetPlayerServerList(ctx, req)
	return
}

func NewPlayerManagerServer(centerServer *center.CenterServer) *PlayerManagerServer {
	ss := &PlayerManagerServer{
		m: centerServer,
	}
	return ss
}
