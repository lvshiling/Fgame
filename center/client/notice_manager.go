package client

import "context"
import centerpb "fgame/fgame/center/pb"

type NoticeManager interface {
	GetNotice(ctx context.Context, platform int32) (resp *centerpb.NoticeResponse, err error)
	RefreshNotice(ctx context.Context) (resp *centerpb.RefreshNoticeResponse, err error)
}

type noticeManager struct {
	c      *Client
	remote centerpb.NoticeManageClient
}

func (m *noticeManager) GetNotice(ctx context.Context, platform int32) (resp *centerpb.NoticeResponse, err error) {
	req := &centerpb.NoticeRequest{}

	req.Platform = platform
	resp, err = m.remote.GetNotice(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func (m *noticeManager) RefreshNotice(ctx context.Context) (resp *centerpb.RefreshNoticeResponse, err error) {
	req := &centerpb.RefreshNoticeRequest{}
	resp, err = m.remote.RefreshNotice(ctx, req)
	// err = toErr(ctx, err)
	if err != nil {
		return
	}
	return
}

func NewNoticeManager(c *Client) NoticeManager {
	m := &noticeManager{}
	m.c = c
	m.remote = centerpb.NewNoticeManageClient(c.conn)
	return m
}
