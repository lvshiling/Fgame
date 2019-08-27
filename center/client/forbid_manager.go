package client

import "context"
import centerpb "fgame/fgame/center/pb"

type ForbidManager interface {
	ForbidIp(ctx context.Context, ip string, flag bool) (resp *centerpb.ForbidIpResponse, err error)
	ForbidIpSearch(ctx context.Context, ip string) (resp *centerpb.ForbidIpSearchResponse, err error)
	GetForbidIpList(ctx context.Context) (resp *centerpb.ForbidIpListResponse, err error)
	ForbidUser(ctx context.Context, userId int64, flag bool) (resp *centerpb.ForbidUserResponse, err error)
	ForbidUserSearch(ctx context.Context, userId int64) (resp *centerpb.ForbidUserSearchResponse, err error)
}

type forbidManager struct {
	c      *Client
	remote centerpb.ForbidManageClient
}

func (m *forbidManager) ForbidIp(ctx context.Context, ip string, flag bool) (resp *centerpb.ForbidIpResponse, err error) {
	req := &centerpb.ForbidIpRequest{}
	req.Ip = ip
	req.Forbid = flag
	resp, err = m.remote.ForbidIp(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *forbidManager) ForbidIpSearch(ctx context.Context, ip string) (resp *centerpb.ForbidIpSearchResponse, err error) {
	req := &centerpb.ForbidIpSearchRequest{}
	req.Ip = ip
	resp, err = m.remote.ForbidSearch(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *forbidManager) ForbidUser(ctx context.Context, user int64, flag bool) (resp *centerpb.ForbidUserResponse, err error) {
	req := &centerpb.ForbidUserRequest{}
	req.UserId = user
	req.Forbid = flag
	resp, err = m.remote.ForbidUser(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *forbidManager) ForbidUserSearch(ctx context.Context, user int64) (resp *centerpb.ForbidUserSearchResponse, err error) {
	req := &centerpb.ForbidUserSearchRequest{}
	req.UserId = user
	resp, err = m.remote.ForbidUserSearch(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *forbidManager) GetForbidIpList(ctx context.Context) (resp *centerpb.ForbidIpListResponse, err error) {
	req := &centerpb.ForbidIpListRequest{}

	resp, err = m.remote.GetForbidIpList(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewForbidManager(c *Client) ForbidManager {
	m := &forbidManager{}
	m.c = c
	m.remote = centerpb.NewForbidManageClient(c.conn)
	return m
}
