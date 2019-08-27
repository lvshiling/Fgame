package serverlist

import (
	"context"
	"fgame/fgame/account/center/center"
	centerclient "fgame/fgame/center/client"
)

type ServerService interface {
	Heartbeat()
	GetServerList(ctx context.Context, platform int32, gm int32) (infoList []*ServerInfo, err error)
}

type serverService struct {
	c *centerclient.Client
}

func (s *serverService) init() (err error) {
	s.c = center.GetCenterService().GetServerClient()
	return
}

func (s *serverService) Heartbeat() {

}

func (s *serverService) GetServerList(ctx context.Context, platform int32, gm int32) (infoList []*ServerInfo, err error) {
	resp, err := s.c.GetServerList(ctx, platform, gm)
	if err != nil {
		return
	}
	infoList = convertFromServerInfoList(resp.GetServerInfoList())
	return
}

func newServerService() *serverService {
	s := &serverService{}
	return s
}

var (
	s *serverService
)

func GetServerService() ServerService {
	return s
}

func Init() (err error) {
	s = newServerService()
	err = s.init()
	if err != nil {
		return
	}
	return
}
