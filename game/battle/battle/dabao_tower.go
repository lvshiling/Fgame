package battle

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
)

// 打宝塔数据管理器
type PlayerTowerManager struct {
	p                   scene.Player
	countExp            int64
	countItemMap        map[int32]int32
	dabaoFlag           bool
	lastDaBaoNoticeTime int64
}

const (
	noticeCD = int64(10 * common.SECOND)
)

func (m *PlayerTowerManager) GetCountTowerExp() int64 {
	return m.countExp
}

func (m *PlayerTowerManager) GetCountTowerItemMap() map[int32]int32 {
	return m.countItemMap
}

func (m *PlayerTowerManager) ResetCountTower() {
	m.countExp = 0
	m.countItemMap = map[int32]int32{}
}

func (m *PlayerTowerManager) CountTowerExp(exp int64) {
	m.countExp += exp
}

func (m *PlayerTowerManager) CountTowerItemMap(itemId, num int32) {
	_, ok := m.countItemMap[itemId]
	if !ok {
		m.countItemMap[itemId] = num
	} else {
		m.countItemMap[itemId] += num
	}
}

func (m *PlayerTowerManager) StartDaBao() {
	m.dabaoFlag = true
}

func (m *PlayerTowerManager) EndDaBao() {
	m.dabaoFlag = false
}

func (m *PlayerTowerManager) IsOnDabao() bool {
	return m.dabaoFlag
}

func (m *PlayerTowerManager) IfNotDaBaoNotice() bool {
	now := global.GetGame().GetTimeService().Now()
	if m.lastDaBaoNoticeTime == 0 {
		m.lastDaBaoNoticeTime = now
		return true
	}

	diff := now - m.lastDaBaoNoticeTime
	if diff > noticeCD {
		m.lastDaBaoNoticeTime = now
		return true
	}

	return false
}

func CreatePlayerTowerManager(p scene.Player) *PlayerTowerManager {
	m := &PlayerTowerManager{
		p: p,
	}
	m.countItemMap = make(map[int32]int32)

	return m
}
