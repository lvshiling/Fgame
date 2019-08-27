package battle

import (
	actvitytypes "fgame/fgame/game/activity/types"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/scene/types"
)

//排行数据变化
type BattlePlayerActivityRankDataChangedEventData struct {
	rankType types.SceneRankType
	rankData *scene.PlayerActvitiyRankData
}

func (d *BattlePlayerActivityRankDataChangedEventData) GetRankType() types.SceneRankType {
	return d.rankType
}

func (d *BattlePlayerActivityRankDataChangedEventData) GetRankData() *scene.PlayerActvitiyRankData {
	return d.rankData
}

func CreateBattlePlayerActivityRankDataChangedEventData(rankType types.SceneRankType, rankData *scene.PlayerActvitiyRankData) *BattlePlayerActivityRankDataChangedEventData {
	d := &BattlePlayerActivityRankDataChangedEventData{}
	d.rankType = rankType
	d.rankData = rankData
	return d
}

//活动排行数据
type PlayerActivityRankManager struct {
	p           scene.Player
	rankDataMap map[actvitytypes.ActivityType]*scene.PlayerActvitiyRankData
}

func (m *PlayerActivityRankManager) enterActivity(activityType actvitytypes.ActivityType, endTime int64) {
	rankData, ok := m.rankDataMap[activityType]
	if !ok {
		defaultRankValueMap := make(map[int32]int64)
		rankData = scene.CreatePlayerActvitiyRankData(activityType, defaultRankValueMap, endTime)
		m.rankDataMap[activityType] = rankData
		goto Change
	}

	if !rankData.RefreshEndTime(endTime) {
		return
	}
Change:
	//发送事件更新数据
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerActivityRankDataRefresh, m.p, rankData)
}

func (m *PlayerActivityRankManager) GetActivityRankValue(activityType actvitytypes.ActivityType, rankType scene.ActivityRankType) int64 {
	rankData, ok := m.rankDataMap[activityType]
	if !ok {
		return 0
	}
	return rankData.GetRankValue(rankType)
}

func (m *PlayerActivityRankManager) GetActivityRankMap() map[actvitytypes.ActivityType]*scene.PlayerActvitiyRankData {
	return m.rankDataMap
}

func (m *PlayerActivityRankManager) UpdateActivityRankValue(activityType actvitytypes.ActivityType, rankType scene.ActivityRankType, val int64) {
	rankData, ok := m.rankDataMap[activityType]
	if !ok {
		return
	}
	rankData.UpdateRankValue(rankType, val)
	eventData := CreateBattlePlayerActivityRankDataChangedEventData(rankType, rankData)
	//发送事件更新数据
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerActivityRankDataChanged, m.p, eventData)
}

func CreatePlayerActivityRankManager(p scene.Player, activityRankDataList []*scene.PlayerActvitiyRankData) *PlayerActivityRankManager {
	m := &PlayerActivityRankManager{}
	m.p = p
	m.rankDataMap = make(map[actvitytypes.ActivityType]*scene.PlayerActvitiyRankData)
	for _, activityRankData := range activityRankDataList {
		m.rankDataMap[activityRankData.GetActivityType()] = activityRankData
	}

	return m
}
