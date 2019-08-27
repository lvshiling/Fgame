package shareboss

import (
	"fgame/fgame/game/scene/scene"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"sync"
)

type ShareBossService interface {
	Start()
	//获取跨服世界boss列表
	GetShareBossList(bossType worldbosstypes.BossType) []scene.NPC
	//获取跨服地图世界boss列表
	GetShareBossListGroupByMap(bossType worldbosstypes.BossType, mapId int32) []scene.NPC
	//获取跨服世界boss
	GetShareBoss(bossType worldbosstypes.BossType, biologyId int32) scene.NPC
	GetGuaiJiShareBossList(bossType worldbosstypes.BossType, force int64) []scene.NPC
}

type shareBossService struct {
}

func (s *shareBossService) init() (err error) {
	return
}

func (s *shareBossService) GetShareBossList(bossType worldbosstypes.BossType) []scene.NPC {
	h := getShareBossHandler(bossType)
	if h == nil {
		return nil
	}
	return h.GetShareBossList()
}

func (s *shareBossService) GetShareBossListGroupByMap(bossType worldbosstypes.BossType, mapId int32) []scene.NPC {
	h := getShareBossHandler(bossType)
	if h == nil {
		return nil
	}
	return h.GetShareBossListGroupByMap(mapId)
}

func (s *shareBossService) GetShareBoss(bossType worldbosstypes.BossType, biologyId int32) scene.NPC {
	h := getShareBossHandler(bossType)
	if h == nil {
		return nil
	}
	return h.GetShareBoss(biologyId)
}

func (s *shareBossService) GetGuaiJiShareBossList(bossType worldbosstypes.BossType, force int64) []scene.NPC {
	h := getShareBossHandler(bossType)
	if h == nil {
		return nil
	}
	return h.GetGuaiJiShareBossList(force)
}

func (s *shareBossService) Start() {
	for _, h := range shareBossHandlerMap {
		h.Start()
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
