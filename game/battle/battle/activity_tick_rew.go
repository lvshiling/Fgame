package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//活动定时奖励管理器
type PlayerActivityTickRewManager struct {
	p               scene.Player
	activityTickRew *scene.PlayerActvitiyTickRewData
}

func (m *PlayerActivityTickRewManager) enterActivity() {
	tickRewData := m.activityTickRew
	if tickRewData == nil {
		tickRewData = scene.CreatePlayerActvitiyTickRewData()
		m.activityTickRew = tickRewData
		return
	}

	tickRewData.RefreshData()
}

func (m *PlayerActivityTickRewManager) GetActivityTickRewData() *scene.PlayerActvitiyTickRewData {
	return m.activityTickRew
}

func (m *PlayerActivityTickRewManager) AddActivityTickRew(resMap, specailResMap map[int32]int32) {
	tickRewData := m.activityTickRew
	if tickRewData == nil {
		return
	}

	tickRewData.UpdateTickRewData(resMap, specailResMap)

	//发送事件更新数据
	eventData := CreateBattlePlayerActivityTickRewDataChangedEventData(resMap, specailResMap)
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerActivityTickRewDataChanged, m.p, eventData)
}

func CreatePlayerActivityTickRewManager(p scene.Player, activityTickRewData *scene.PlayerActvitiyTickRewData) *PlayerActivityTickRewManager {
	m := &PlayerActivityTickRewManager{}
	m.p = p
	m.activityTickRew = activityTickRewData
	return m
}

//定时数据变化
type BattlePlayerActivityTickRewDataChangedEventData struct {
	addResMap     map[int32]int32
	specailResMap map[int32]int32
}

func (d *BattlePlayerActivityTickRewDataChangedEventData) GetAddResMap() map[int32]int32 {
	return d.addResMap
}

func (d *BattlePlayerActivityTickRewDataChangedEventData) GetSpecialResMap() map[int32]int32 {
	return d.specailResMap
}

func CreateBattlePlayerActivityTickRewDataChangedEventData(addResMap, specailResMap map[int32]int32) *BattlePlayerActivityTickRewDataChangedEventData {
	d := &BattlePlayerActivityTickRewDataChangedEventData{}
	d.addResMap = addResMap
	d.specailResMap = specailResMap
	return d
}
