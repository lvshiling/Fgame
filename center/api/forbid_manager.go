package api

import (
	"context"
	"fgame/fgame/center/center"
	centerpb "fgame/fgame/center/pb"
)

//配置服务器
type ForbidManagerServer struct {
	m center.ForbidManager
}

func (ps *ForbidManagerServer) GetForbidIpList(ctx context.Context, req *centerpb.ForbidIpListRequest) (res *centerpb.ForbidIpListResponse, err error) {
	res, err = ps.m.GetForbidIpList(ctx, req)
	return
}

func (ps *ForbidManagerServer) ForbidIp(ctx context.Context, req *centerpb.ForbidIpRequest) (res *centerpb.ForbidIpResponse, err error) {
	res, err = ps.m.ForbidIp(ctx, req)
	return
}

func (ps *ForbidManagerServer) ForbidSearch(ctx context.Context, req *centerpb.ForbidIpSearchRequest) (res *centerpb.ForbidIpSearchResponse, err error) {
	res, err = ps.m.ForbidSearch(ctx, req)
	return
}

func (ps *ForbidManagerServer) ForbidUser(ctx context.Context, req *centerpb.ForbidUserRequest) (res *centerpb.ForbidUserResponse, err error) {
	res, err = ps.m.ForbidUser(ctx, req)
	return
}

func (ps *ForbidManagerServer) ForbidUserSearch(ctx context.Context, req *centerpb.ForbidUserSearchRequest) (res *centerpb.ForbidUserSearchResponse, err error) {
	res, err = ps.m.ForbidUserSearch(ctx, req)
	return
}

func NewForbidManagerServer(centerServer *center.CenterServer) *ForbidManagerServer {
	ss := &ForbidManagerServer{
		m: centerServer,
	}
	return ss
}
