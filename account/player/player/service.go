package player

import (
	"context"
	"fgame/fgame/account/center/center"
	centerclient "fgame/fgame/center/client"
)

type PlayerInfo struct {
	UserId   int64
	ServerId int32
	PlayerId int64
	Role     int32
	Sex      int32
	Level    int32
	ZhuanShu int32
}

type PlayerService interface {
	GetPlayerList(ctx context.Context, userId int64) (playerInfoList []*PlayerInfo, err error)
}

type playerService struct {
	c *centerclient.Client
}

func (s *playerService) init() (err error) {
	s.c = center.GetCenterService().GetLoginClient()
	return
}

func (s *playerService) GetPlayerList(ctx context.Context, userId int64) (playerInfoList []*PlayerInfo, err error) {
	resp, err := s.c.GetPlayerServerList(ctx, userId)
	if err != nil {
		return
	}
	//通过缓存获取
	playerInfoList = convertFromPlayerInfoList(resp.PlayerList)
	return
}

func newPlayerService() *playerService {
	ps := &playerService{}
	return ps
}

var (
	s *playerService
)

func GetPlayerService() PlayerService {
	return s
}

func Init() (err error) {
	s = newPlayerService()
	err = s.init()
	if err != nil {
		return
	}
	return
}
