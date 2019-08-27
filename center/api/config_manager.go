package api

import (
	"context"
	"fgame/fgame/center/center"
	centerpb "fgame/fgame/center/pb"
)

//配置服务器
type ConfigManagerServer struct {
	m center.ConfigManager
}

func (ps *ConfigManagerServer) RefreshClientVersion(ctx context.Context, req *centerpb.ClientVersionRefreshRequest) (res *centerpb.ClientVersionRefreshResponse, err error) {
	res, err = ps.m.RefreshClientVersion(ctx, req)
	return
}

func (ps *ConfigManagerServer) RefreshServerConfig(ctx context.Context, req *centerpb.ServerConfigRefreshRequest) (res *centerpb.ServerConfigRefreshResponse, err error) {
	res, err = ps.m.RefreshServerConfig(ctx, req)
	return
}

func (ps *ConfigManagerServer) RefreshPlatformConfig(ctx context.Context, req *centerpb.PlatformConfigRefreshRequest) (res *centerpb.PlatformConfigRefreshResponse, err error) {
	res, err = ps.m.RefreshPlatformConfig(ctx, req)
	return
}

func NewConfigManagerServer(centerServer *center.CenterServer) *ConfigManagerServer {
	ss := &ConfigManagerServer{
		m: centerServer,
	}
	return ss
}
