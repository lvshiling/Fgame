package battle

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
)

// 外域boss
type PlayerOutlandBossManager struct {
	p                    scene.Player
	zhuoqiNum            int32
	lastAttackNoticeTime int64
}

const (
	attackNoticeCd = int64(10 * common.SECOND)
)

func (m *PlayerOutlandBossManager) IsZhuoQiLimit() bool {
	return m.zhuoqiNum >= constant.GetConstantService().GetConstant(constanttypes.ConstantTypeOutlandBossZhuoQiLimit)
}

func (m *PlayerOutlandBossManager) IsZhuoQiNoticeCd() bool {
	now := global.GetGame().GetTimeService().Now()
	if now-m.lastAttackNoticeTime <= attackNoticeCd {
		return true
	}

	m.lastAttackNoticeTime = now
	return false
}

func (m *PlayerOutlandBossManager) SynZhuoQiNum(zhuoqi int32) {
	m.zhuoqiNum = zhuoqi
}

func CreatePlayerOutlandBossManagerWithData(p scene.Player, zhuoqi int32) *PlayerOutlandBossManager {
	m := &PlayerOutlandBossManager{
		p:         p,
		zhuoqiNum: zhuoqi,
	}
	return m
}

func CreatePlayerOutlandBossManager(p scene.Player) *PlayerOutlandBossManager {
	m := &PlayerOutlandBossManager{
		p: p,
	}
	return m
}
