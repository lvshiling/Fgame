package battle

import (
	activitytemplate "fgame/fgame/game/activity/template"
	actvitytypes "fgame/fgame/game/activity/types"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
)

//pk管理器
type PlayerActivityPKManager struct {
	p           scene.Player
	killDataMap map[actvitytypes.ActivityType]*scene.PlayerActvitiyKillData
}

func (m *PlayerActivityPKManager) enterActivity(activityType actvitytypes.ActivityType, endTime int64) {

}

func (m *PlayerActivityPKManager) GetPlayerActvitityKillMap() map[actvitytypes.ActivityType]*scene.PlayerActvitiyKillData {
	return m.killDataMap
}

func (m *PlayerActivityPKManager) IfCanKilledInActivity(activityType actvitytypes.ActivityType) (bool, int32) {
	pkRewardCdTemplate := activitytemplate.GetActivityTemplateService().GetPkRewardCdByType(activityType)
	if pkRewardCdTemplate == nil {
		return true, 0
	}
	killData, ok := m.killDataMap[activityType]
	if !ok {
		return true, 0
	}
	if killData.GetKilledNum() < pkRewardCdTemplate.KillCount {
		return true, 0
	}
	now := global.GetGame().GetTimeService().Now()
	elapse := now - killData.GetLastKilledTime()
	if elapse >= int64(pkRewardCdTemplate.KillCd) {
		return true, 0
	}

	return false, int32(int64(pkRewardCdTemplate.KillCd) - elapse)
}

func (m *PlayerActivityPKManager) KilledInActivity(activityType actvitytypes.ActivityType) (flag bool) {
	flag, _ = m.IfCanKilledInActivity(activityType)
	if !flag {
		return false
	}
	pkRewardCdTemplate := activitytemplate.GetActivityTemplateService().GetPkRewardCdByType(activityType)
	if pkRewardCdTemplate == nil {
		return true
	}
	killData, ok := m.killDataMap[activityType]
	if !ok {
		killData = scene.CreatePlayerActvitiyKillData(activityType, 0, 0)
		m.killDataMap[activityType] = killData
	}

	now := global.GetGame().GetTimeService().Now()
	if killData.GetKilledNum() > pkRewardCdTemplate.KillCount {
		return false
	}
	elapse := now - killData.GetLastKilledTime()
	if elapse >= int64(pkRewardCdTemplate.KillCd) {
		killData.Reset()
	}
	killData.Kill(now)
	//发送事件
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerActivityPkDataChanged, m.p, killData)
	return true
}

func (m *PlayerActivityPKManager) SyncKillData(killData *scene.PlayerActvitiyKillData) {
	m.killDataMap[killData.GetActivityType()] = killData
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerActivityPkDataSync, m.p, killData)
}

func CreatePlayerActivityPKManager(p scene.Player, activityKillDataList []*scene.PlayerActvitiyKillData) *PlayerActivityPKManager {
	m := &PlayerActivityPKManager{}
	m.p = p
	m.killDataMap = make(map[actvitytypes.ActivityType]*scene.PlayerActvitiyKillData)
	for _, activityKillData := range activityKillDataList {
		m.killDataMap[activityKillData.GetActivityType()] = activityKillData
	}
	return m
}
