package client

import "context"
import centerpb "fgame/fgame/center/pb"

type PlayerManager interface {
	SyncPlayerServerInfo(ctx context.Context, req *centerpb.PlayerInfoSyncRequest) (resp *centerpb.PlayerInfoSyncResponse, err error)
	GetPlayerServerList(ctx context.Context, userId int64) (resp *centerpb.PlayerListResponse, err error)
}

type playerManager struct {
	c      *Client
	remote centerpb.PlayerManageClient
}

func (m *playerManager) SyncPlayerServerInfo(ctx context.Context, req *centerpb.PlayerInfoSyncRequest) (resp *centerpb.PlayerInfoSyncResponse, err error) {
	resp, err = m.remote.SyncPlayerServerInfo(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *playerManager) GetPlayerServerList(ctx context.Context, userId int64) (resp *centerpb.PlayerListResponse, err error) {
	req := &centerpb.PlayerListRequest{}
	req.UserId = userId
	resp, err = m.remote.GetPlayerServerList(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewPlayerManager(c *Client) PlayerManager {
	m := &playerManager{}
	m.c = c
	m.remote = centerpb.NewPlayerManageClient(c.conn)
	return m
}
