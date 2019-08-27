package battle

import (
	actvitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/scene/scene"
)

//pk管理器
type PlayerActivityManager struct {
	*PlayerActivityRankManager
	*PlayerActivityPKManager
	*PlayerActivityCollectManager
	*PlayerActivityTickRewManager
	p scene.Player
}

func (m *PlayerActivityManager) EnterActivity(activityType actvitytypes.ActivityType, endTime int64) {
	//进入活动
	m.PlayerActivityPKManager.enterActivity(activityType, endTime)
	m.PlayerActivityRankManager.enterActivity(activityType, endTime)
	m.PlayerActivityCollectManager.enterActivity(activityType, endTime)
	m.PlayerActivityTickRewManager.enterActivity()
}

func CreatePlayerActivityManager(
	p scene.Player,
	activityKillDataList []*scene.PlayerActvitiyKillData,
	rankDataList []*scene.PlayerActvitiyRankData,
	collectDataList []*scene.PlayerActvitiyCollectData,
	tickRewData *scene.PlayerActvitiyTickRewData,
) *PlayerActivityManager {
	m := &PlayerActivityManager{}
	m.p = p
	m.PlayerActivityPKManager = CreatePlayerActivityPKManager(p, activityKillDataList)
	m.PlayerActivityRankManager = CreatePlayerActivityRankManager(p, rankDataList)
	m.PlayerActivityCollectManager = CreatePlayerActivityCollectManager(p, collectDataList)
	m.PlayerActivityTickRewManager = CreatePlayerActivityTickRewManager(p, tickRewData)
	return m
}
