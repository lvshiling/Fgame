package battle

import (
	actvitytypes "fgame/fgame/game/activity/types"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//活动采集管理器
type PlayerActivityCollectManager struct {
	p                scene.Player
	activityCountMap map[actvitytypes.ActivityType]*scene.PlayerActvitiyCollectData
}

func (m *PlayerActivityCollectManager) enterActivity(activityType actvitytypes.ActivityType, endTime int64) {
	collectData := m.getCollectData(activityType)
	if collectData == nil {
		defaultMap := make(map[int32]int32)
		collectData = scene.CreatePlayerActvitiyCollectData(activityType, defaultMap, endTime)
		m.activityCountMap[activityType] = collectData 
		goto Changed
	}

	collectData.RefreshData(endTime)

Changed:
	//发送事件更新数据
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerActivityCollectDataRefresh, m.p, collectData)
}

func (m *PlayerActivityCollectManager) GetActivityCollectCountMap(activityType actvitytypes.ActivityType) map[int32]int32 {
	collectData := m.getCollectData(activityType)
	if collectData == nil {
		return nil
	}
	return collectData.GetCountMap()
}

func (m *PlayerActivityCollectManager) GetActivityTotalCollectCount(activityType actvitytypes.ActivityType) int32 {
	return 0
}

func (m *PlayerActivityCollectManager) UpdateActivityCollect(activityType actvitytypes.ActivityType, biologyId int32) {
	collectData := m.getCollectData(activityType)
	if collectData == nil {
		return
	}

	collectData.UpdateData(biologyId)
	//发送事件更新数据
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerActivityCollectDataChanged, m.p, collectData)
}

func (m *PlayerActivityCollectManager) getCollectData(activityType actvitytypes.ActivityType) *scene.PlayerActvitiyCollectData {
	data, ok := m.activityCountMap[activityType]
	if !ok {
		return nil
	}

	return data
}

func CreatePlayerActivityCollectManager(p scene.Player, activityCollectDataList []*scene.PlayerActvitiyCollectData) *PlayerActivityCollectManager {
	m := &PlayerActivityCollectManager{}
	m.p = p
	m.activityCountMap = make(map[actvitytypes.ActivityType]*scene.PlayerActvitiyCollectData)
	for _, collectData := range activityCollectDataList {
		m.activityCountMap[collectData.GetActivityType()] = collectData
	}
	return m
}
