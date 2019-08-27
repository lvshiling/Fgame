package worldboss

import (
	"fgame/fgame/game/player"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fmt"
)

type KillBossHandler interface {
	KillBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32)
}

type KillBossHandlerFunc func(pl player.Player, typ worldbosstypes.BossType, biologyId int32)

func (f KillBossHandlerFunc) KillBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32) {
	f(pl, typ, biologyId)
}

var (
	killHandlerMap = make(map[worldbosstypes.BossType]KillBossHandler)
)

func RegistKillBossHandler(typ worldbosstypes.BossType, h KillBossHandler) {
	_, ok := killHandlerMap[typ]
	if ok {
		panic(fmt.Errorf("worldBoss:击杀处理器重复注册，类型:%d", typ))
	}

	killHandlerMap[typ] = h
}

func GetKillBossHandler(typ worldbosstypes.BossType) KillBossHandler {
	h, ok := killHandlerMap[typ]
	if !ok {
		return nil
	}

	return h
}
