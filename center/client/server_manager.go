package client

import "context"
import centerpb "fgame/fgame/center/pb"
import centertypes "fgame/fgame/center/types"

type ServerManager interface {
	Register(ctx context.Context, serverType centertypes.GameServerType, platform int32, serverId int32, serverIp string, serverPort int32) (resp *centerpb.ServerRegisterResponse, err error)
	Unregister(ctx context.Context, serverId int32) (resp *centerpb.ServerUnregisterResponse, err error)
	GetCrossList(ctx context.Context, serverId int32) (resp *centerpb.ServerCrossListResponse, err error)
	GetServerList(ctx context.Context, platform int32, gm int32) (resp *centerpb.ServerInfoListResponse, err error)
	Refresh(ctx context.Context, platform int32) (resp *centerpb.RefreshServerInfoListResponse, err error)
	RefreshSDK(ctx context.Context) (resp *centerpb.RefreshSDKListResponse, err error)
	GetServerInfo(ctx context.Context, sdk int32, serverId int32) (resp *centerpb.ServerInfoResponse, err error)
	GetServerInfoByPlatform(ctx context.Context, platform int32, serverId int32) (resp *centerpb.ServerInfoByPlatformResponse, err error)
	RefreshMarryPrice(ctx context.Context, platform int32) (resp *centerpb.RefreshMarryPriceListResponse, err error)
}

type serverManager struct {
	c      *Client
	remote centerpb.ServerManageClient
}

func (m *serverManager) Register(ctx context.Context, serverType centertypes.GameServerType, platform int32, serverId int32, serverIp string, serverPort int32) (resp *centerpb.ServerRegisterResponse, err error) {
	req := &centerpb.ServerRegisterRequest{}
	req.ServerType = int32(serverType)
	req.ServerId = serverId
	req.ServerIp = serverIp
	req.ServerPort = serverPort
	req.Platform = platform
	resp, err = m.remote.Register(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *serverManager) Unregister(ctx context.Context, serverId int32) (resp *centerpb.ServerUnregisterResponse, err error) {
	req := &centerpb.ServerUnregisterRequest{}
	req.ServerId = serverId
	resp, err = m.remote.Unregister(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *serverManager) GetCrossList(ctx context.Context, serverId int32) (resp *centerpb.ServerCrossListResponse, err error) {
	req := &centerpb.ServerCrossListRequest{}
	req.ServerId = serverId
	resp, err = m.remote.GetCrossList(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *serverManager) GetServerList(ctx context.Context, platform int32, gm int32) (resp *centerpb.ServerInfoListResponse, err error) {
	req := &centerpb.ServerInfoListRequest{}
	req.Platform = platform
	req.Gm = gm
	resp, err = m.remote.GetServerList(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *serverManager) Refresh(ctx context.Context, platform int32) (resp *centerpb.RefreshServerInfoListResponse, err error) {
	req := &centerpb.RefreshServerInfoListRequest{}
	req.Platform = platform
	resp, err = m.remote.Refresh(ctx, req)
	if err != nil {
		return
	}
	return
}

func (m *serverManager) RefreshSDK(ctx context.Context) (resp *centerpb.RefreshSDKListResponse, err error) {
	req := &centerpb.RefreshSDKListRequest{}
	resp, err = m.remote.RefreshSDK(ctx, req)
	if err != nil {
		return
	}
	return
}

func (m *serverManager) GetServerInfo(ctx context.Context, sdk int32, serverId int32) (resp *centerpb.ServerInfoResponse, err error) {
	req := &centerpb.ServerInfoRequest{}
	req.Platform = sdk
	req.ServerId = serverId
	resp, err = m.remote.GetServerInfo(ctx, req)
	if err != nil {
		return
	}
	return
}

func (m *serverManager) GetServerInfoByPlatform(ctx context.Context, platform int32, serverId int32) (resp *centerpb.ServerInfoByPlatformResponse, err error) {
	req := &centerpb.ServerInfoByPlatformRequest{}
	req.Platform = platform
	req.ServerId = serverId
	resp, err = m.remote.GetServerInfoByPlatform(ctx, req)
	if err != nil {
		return
	}
	return
}

func (m *serverManager) RefreshMarryPrice(ctx context.Context, platform int32) (resp *centerpb.RefreshMarryPriceListResponse, err error) {
	req := &centerpb.RefreshMarryPriceListRequest{}
	req.Platform = platform

	resp, err = m.remote.RefreshMarryPrice(ctx, req)
	if err != nil {
		return
	}
	return
}

func NewServerManager(c *Client) ServerManager {
	m := &serverManager{}
	m.c = c
	m.remote = centerpb.NewServerManageClient(c.conn)
	return m
}
