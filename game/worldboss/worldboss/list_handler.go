package worldboss

import (
	"fgame/fgame/game/player"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fmt"
)

type BossListHandler interface {
	BossList(pl player.Player, t worldbosstypes.BossType)
}

type BossListHandlerFunc func(pl player.Player, t worldbosstypes.BossType)

func (f BossListHandlerFunc) BossList(pl player.Player, t worldbosstypes.BossType) {
	f(pl, t)
}

var (
	listHandlerMap = make(map[worldbosstypes.BossType]BossListHandler)
)

func RegistBossListHandler(typ worldbosstypes.BossType, h BossListHandler) {
	_, ok := listHandlerMap[typ]
	if ok {
		panic(fmt.Errorf("worldBoss:boss列表处理器重复注册，类型:%d", typ))
	}

	listHandlerMap[typ] = h
}

func GetBossListHandler(typ worldbosstypes.BossType) BossListHandler {
	h, ok := listHandlerMap[typ]
	if !ok {
		return nil
	}

	return h
}
