package client

import "context"
import centerpb "fgame/fgame/center/pb"

type ConfigManager interface {
	RefreshClientVersion(ctx context.Context) (resp *centerpb.ClientVersionRefreshResponse, err error)
	RefreshServerConfig(ctx context.Context) (resp *centerpb.ServerConfigRefreshResponse, err error)
	RefreshPlatformConfig(ctx context.Context, platformId int32) (resp *centerpb.PlatformConfigRefreshResponse, err error)
}

type configManager struct {
	c      *Client
	remote centerpb.ConfigManageClient
}

func (m *configManager) RefreshClientVersion(ctx context.Context) (resp *centerpb.ClientVersionRefreshResponse, err error) {
	req := &centerpb.ClientVersionRefreshRequest{}
	resp, err = m.remote.RefreshClientVersion(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *configManager) RefreshServerConfig(ctx context.Context) (resp *centerpb.ServerConfigRefreshResponse, err error) {
	req := &centerpb.ServerConfigRefreshRequest{}
	resp, err = m.remote.RefreshServerConfig(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *configManager) RefreshPlatformConfig(ctx context.Context, platformId int32) (resp *centerpb.PlatformConfigRefreshResponse, err error) {
	req := &centerpb.PlatformConfigRefreshRequest{}
	req.Platform = platformId
	resp, err = m.remote.RefreshPlatformConfig(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewConfigManager(c *Client) ConfigManager {
	m := &configManager{}
	m.c = c
	m.remote = centerpb.NewConfigManageClient(c.conn)
	return m
}
