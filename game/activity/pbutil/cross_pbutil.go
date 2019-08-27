package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/scene/scene"
)

func ConvertToPlayerKillData(pkData *crosspb.ActivityPkData) *scene.PlayerActvitiyKillData {
	actvityType := activitytypes.ActivityType(pkData.GetActivityType())
	killNum := pkData.GetKilledNum()
	lastKilledTime := pkData.GetLastKillTime()
	killData := scene.CreatePlayerActvitiyKillData(actvityType, killNum, lastKilledTime)
	return killData
}

func BuildSIPlayerActivityPkDataChanged(killData *scene.PlayerActvitiyKillData) *crosspb.SIPlayerActivityPkDataChanged {
	siPlayerActivityPkDataChanged := &crosspb.SIPlayerActivityPkDataChanged{}
	pkData := &crosspb.ActivityPkData{}
	activityType := int32(killData.GetActivityType())
	killedNum := killData.GetKilledNum()
	lastKilledTime := killData.GetLastKilledTime()
	pkData.KilledNum = &killedNum
	pkData.LastKillTime = &lastKilledTime
	pkData.ActivityType = &activityType
	siPlayerActivityPkDataChanged.ActivityPkData = pkData
	return siPlayerActivityPkDataChanged
}
