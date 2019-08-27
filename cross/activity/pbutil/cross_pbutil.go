package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosspbutil "fgame/fgame/game/cross/pbutil"
	"fgame/fgame/game/scene/scene"
)

func BuildISPlayerActivityPkDataChanged(killData *scene.PlayerActvitiyKillData) *crosspb.ISPlayerActivityPkDataChanged {
	isPlayerActivityPkDataChanged := &crosspb.ISPlayerActivityPkDataChanged{}
	pkData := &crosspb.ActivityPkData{}
	activityType := int32(killData.GetActivityType())
	killedNum := killData.GetKilledNum()
	lastKilledTime := killData.GetLastKilledTime()
	pkData.KilledNum = &killedNum
	pkData.LastKillTime = &lastKilledTime
	pkData.ActivityType = &activityType
	isPlayerActivityPkDataChanged.ActivityPkData = pkData
	return isPlayerActivityPkDataChanged
}

func BuildISPlayerActivityRankDataChanged(activityType, rankType int32, val int64) *crosspb.ISPlayerActivityRankDataChanged {
	isMsg := &crosspb.ISPlayerActivityRankDataChanged{}
	isMsg.ActivityType = &activityType
	isMsg.RankType = &rankType
	isMsg.Val = &val
	return isMsg
}

func BuildISPlayerActivityTickRewDataChanged(resMap, specialResMap map[int32]int32) *crosspb.ISPlayerActivityTickRewDataChanged {
	isMsg := &crosspb.ISPlayerActivityTickRewDataChanged{}

	for key, val := range resMap {
		p := crosspbutil.BuildProperty(key, int64(val))
		isMsg.PropertyList = append(isMsg.PropertyList, p)
	}

	for key, val := range specialResMap {
		p := crosspbutil.BuildProperty(key, int64(val))
		isMsg.SpecialPropertyList = append(isMsg.SpecialPropertyList, p)
	}
	return isMsg
}
