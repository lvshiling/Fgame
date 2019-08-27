package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/godsiege/godsiege"
	godsiegelogic "fgame/fgame/game/godsiege/logic"
	"fgame/fgame/game/godsiege/pbutil"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeGodSiegeQiLin, activity.ActivityAttendHandlerFunc(godsiegelogic.PlayerEnterGodSiegeQiLinScene))
	activity.RegisterActivityHandler(activitytypes.ActivityTypeGodSiegeHuoFeng, activity.ActivityAttendHandlerFunc(godsiegelogic.PlayerEnterGodSiegeHuoFengScene))
	activity.RegisterActivityHandler(activitytypes.ActivityTypeGodSiegeDuLong, activity.ActivityAttendHandlerFunc(godsiegelogic.PlayerEnterGodSiegeDuLongScene))
	activity.RegisterActivityHandler(activitytypes.ActivityTypeLocalGodSiegeQiLin, activity.ActivityAttendHandlerFunc(AttendLocalGodSiegeQiLin))
}

// 参加神兽攻城-麒麟来袭（本服）
func AttendLocalGodSiegeQiLin(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (success bool, err error) {
	godType, ok := godsiegetypes.GetGodSiegeType(activityTemplate.GetActivityType())
	if !ok {
		return
	}

	s := godsiege.GetGodSiegeService().GetGodSiegeScene(godType)
	if s == nil {
		now := global.GetGame().GetTimeService().Now()
		openTime := global.GetGame().GetServerTime()
		mergeTime := merge.GetMergeService().GetMergeTime()
		activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
		if err != nil {
			return false, nil
		}
		if activityTimeTemplate == nil {
			return false, nil
		}
		endTime, err := activityTimeTemplate.GetEndTime(now)
		if err != nil {
			return false, nil
		}
		s = godsiege.GetGodSiegeService().CreateGodSiegeScene(godType, activityTemplate.Mapid, endTime)
		if s == nil {
			return false, nil
		}
	}

	beforeNum, isLineUp, _ := godsiege.GetGodSiegeService().Attend(godType, pl.GetId())
	if isLineUp {
		pl.GodSiegeLineUp(godType)
		scGodSiegeLineUp := pbutil.BuildSCGodSiegeLineUp(int32(godType), beforeNum)
		pl.SendMsg(scGodSiegeLineUp)
		return
	}

	pos, flag := godsiege.GetGodSiegeService().GetRebornPos(godType, pl.GetId())
	if !flag {
		return
	}

	if !scenelogic.PlayerEnterScene(pl, s, pos) {
		return
	}

	success = true
	return
}
