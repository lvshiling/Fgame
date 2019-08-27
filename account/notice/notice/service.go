package notice

import (
	"context"
	"fgame/fgame/account/center/center"
	centerclient "fgame/fgame/center/client"
)

type NoticeService interface {
	GetNotice(ctx context.Context, platform int32) (notice string, err error)
}

type noticeService struct {
	c *centerclient.Client
}

func (s *noticeService) init() (err error) {
	s.c = center.GetCenterService().GetNoticeClient()
	return
}

func (s *noticeService) GetNotice(ctx context.Context, platform int32) (notice string, err error) {
	resp, err := s.c.GetNotice(ctx, platform)
	if err != nil {
		return
	}
	notice = resp.GetNotice()
	return
}

func newNoticeService() *noticeService {
	ps := &noticeService{}
	return ps
}

var (
	s *noticeService
)

func GetNoticeService() NoticeService {
	return s
}

func Init() (err error) {
	s = newNoticeService()
	err = s.init()
	if err != nil {
		return
	}
	return
}
