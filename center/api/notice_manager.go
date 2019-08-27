package api

import (
	"context"
	"fgame/fgame/center/center"
	centerpb "fgame/fgame/center/pb"
)

//服务器服务
type NoticeManagerServer struct {
	m center.NoticeManager
}

func (ps *NoticeManagerServer) GetNotice(ctx context.Context, req *centerpb.NoticeRequest) (res *centerpb.NoticeResponse, err error) {
	res, err = ps.m.GetNotice(ctx, req)
	return
}

func (ps *NoticeManagerServer) RefreshNotice(ctx context.Context, req *centerpb.RefreshNoticeRequest) (res *centerpb.RefreshNoticeResponse, err error) {
	res, err = ps.m.RefreshNotice(ctx, req)
	return
}

func NewNoticeManagerServer(centerServer *center.CenterServer) *NoticeManagerServer {
	ss := &NoticeManagerServer{
		m: centerServer,
	}
	return ss
}
