package shareboss

import (
	"context"
	sharebossclient "fgame/fgame/cross/shareboss/client"
	"fgame/fgame/game/center/center"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fmt"
	"sync"

	crosstypes "fgame/fgame/game/cross/types"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShareBossService interface {
	Heartbeat()
	Start()
	//获取跨服世界boss列表
	GetShareBossList(typ worldbosstypes.BossType) []*ShareBossInfo
	//获取跨服世界boss
	GetShareBoss(typ worldbosstypes.BossType, biologyId int32) *ShareBossInfo
}

type shareBossService struct {
	rwm                sync.RWMutex
	shareBossClientMap map[crosstypes.CrossType]sharebossclient.ShareBossClient

	shareBostListOfMap map[worldbosstypes.BossType][]*ShareBossInfo
}

func (s *shareBossService) init() (err error) {
	s.shareBossClientMap = make(map[crosstypes.CrossType]sharebossclient.ShareBossClient)
	s.shareBostListOfMap = make(map[worldbosstypes.BossType][]*ShareBossInfo)
	err = s.syncBossListMap()
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("shareboss:跨服boss同步列表")
		return err
	}
	return
}

//定时同步boss列表
func (s *shareBossService) syncBossListMap() (err error) {

	for bossType := worldbosstypes.MinBossType; bossType <= worldbosstypes.MaxBossType; bossType++ {
		crossType := bossType.CrossType()
		if crossType == crosstypes.CrossTypeNone {
			continue
		}
		c := s.getClient(crossType)
		if c == nil {
			//重置
			c, err = s.resetClient(crossType)
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("shareboss:同步跨服boss,重新获取客户端")
				continue
			}

		}
		ctx := context.TODO()
		resp, err := c.GetShareBossList(ctx, int32(bossType))
		if err != nil {
			sta := status.Convert(err)
			if sta.Code() == codes.Canceled {
				//重新获取
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("shareboss:同步跨服boss,重新获取客户端")
				_, err = s.resetClient(crossType)
				if err != nil {
					log.WithFields(
						log.Fields{
							"err": err,
						}).Warn("shareboss:同步跨服boss,失败")
					continue
				}
				continue
			}
			log.WithFields(
				log.Fields{
					"err": err,
				}).Warn("shareboss:同步跨服boss,获取跨服列表错误")
			continue
		}
		shareBossList := convertFromBossInfoList(resp.BossInfoList)
		s.syncBossList(bossType, shareBossList)
	}
	return nil
}

func (s *shareBossService) syncBossList(typ worldbosstypes.BossType, bossList []*ShareBossInfo) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.shareBostListOfMap[typ] = bossList
}

//定时同步boss列表
func (s *shareBossService) getClient(t crosstypes.CrossType) sharebossclient.ShareBossClient {
	c := s.shareBossClientMap[t]
	if c != nil {
		return c
	}
	return nil
}

func (s *shareBossService) GetShareBossList(typ worldbosstypes.BossType) []*ShareBossInfo {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	tempBossList, ok := s.shareBostListOfMap[typ]
	if !ok {
		return nil
	}
	return tempBossList
}

func (s *shareBossService) GetShareBoss(typ worldbosstypes.BossType, biologyId int32) *ShareBossInfo {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	tempBossList, ok := s.shareBostListOfMap[typ]
	if !ok {
		return nil
	}
	for _, boss := range tempBossList {
		bossBiologyId := boss.GetBiologyId()
		if bossBiologyId == biologyId {
			return boss
		}
	}
	return nil
}

func (s *shareBossService) Start() {

}

func (s *shareBossService) resetClient(crossType crosstypes.CrossType) (c sharebossclient.ShareBossClient, err error) {
	serverType := crossType.GetServerType()
	conn := center.GetCenterService().GetCross(serverType)
	if conn == nil {
		return nil, fmt.Errorf("shareboss:跨服连接不存在")
	}

	shareBossClient := sharebossclient.NewShareBossClient(conn)
	s.shareBossClientMap[crossType] = shareBossClient
	return shareBossClient, nil
}

func (s *shareBossService) Heartbeat() {
	err := s.syncBossListMap()
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("shareboss:同步跨服boss,失败")
	}
}

var (
	once sync.Once
	ws   *shareBossService
)

func Init() (err error) {
	once.Do(func() {
		ws = &shareBossService{}
		err = ws.init()
	})
	return err
}

func GetShareBossService() ShareBossService {
	return ws
}
