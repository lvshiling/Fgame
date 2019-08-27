package shareboss

import (
	"fgame/fgame/game/scene/scene"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fmt"
)

type ShareBossHandler interface {
	Start()
	GetShareBossList() []scene.NPC
	//获取跨服地图世界boss列表
	GetShareBossListGroupByMap(mapId int32) []scene.NPC
	//获取跨服世界boss
	GetShareBoss(biologyId int32) scene.NPC
	GetGuaiJiShareBossList(force int64) []scene.NPC
}

var (
	shareBossHandlerMap = map[worldbosstypes.BossType]ShareBossHandler{}
)

func RegisterShareBossHandler(t worldbosstypes.BossType, h ShareBossHandler) {
	_, ok := shareBossHandlerMap[t]
	if ok {
		panic(fmt.Errorf("重复注册[%s]boss处理器", t.String()))
	}
	shareBossHandlerMap[t] = h
}

func getShareBossHandler(t worldbosstypes.BossType) ShareBossHandler {
	h, ok := shareBossHandlerMap[t]
	if !ok {
		return nil
	}
	return h
}

// func GetShareBossList(bossType worldbosstypes.BossType) []scene.NPC {
// 	h := getShareBossHandler(bossType)
// 	if h == nil {
// 		return nil
// 	}
// 	return h.GetShareBossList()
// }

// func GetShareBossListGroupByMap(bossType worldbosstypes.BossType, mapId int32) []scene.NPC {
// 	h := getShareBossHandler(bossType)
// 	if h == nil {
// 		return nil
// 	}
// 	return h.GetShareBossListGroupByMap(mapId)
// }

// func GetShareBoss(bossType worldbosstypes.BossType, biologyId int32) scene.NPC {
// 	h := getShareBossHandler(bossType)
// 	if h == nil {
// 		return nil
// 	}
// 	return h.GetShareBoss(biologyId)
// }

// func GetGuaiJiShareBossList(bossType worldbosstypes.BossType, force int64) []scene.NPC {
// 	h := getShareBossHandler(bossType)
// 	if h == nil {
// 		return nil
// 	}
// 	return h.GetGuaiJiShareBossList(force)
// }
