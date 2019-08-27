package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lianyu/lianyu"
	lianyulogic "fgame/fgame/game/lianyu/logic"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeLianYu, activity.ActivityAttendHandlerFunc(lianyulogic.PlayerEnterLianYuScene))
	activity.RegisterActivityHandler(activitytypes.ActivityTypeLocalLianYu, activity.ActivityAttendHandlerFunc(AttendLocalLianYu))
}

// 参加无间炼狱（本服）
func AttendLocalLianYu(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (success bool, err error) {

	s := lianyu.GetLianYuService().GetLianYuScene()
	if s == nil {
		now := global.GetGame().GetTimeService().Now()
		openTime := global.GetGame().GetServerTime()
		mergeTime := merge.GetMergeService().GetMergeTime()
		activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
		if err != nil {
			return false, err
		}
		if activityTimeTemplate == nil {
			return false, err
		}
		endTime, err := activityTimeTemplate.GetEndTime(now)
		if err != nil {
			return false, err
		}
		s = lianyu.GetLianYuService().CreateLianYuScene(activityTemplate.Mapid, endTime)
		if s == nil {
			return false, err
		}
	}

	// beforeNum, isLineUp := lianyu.GetLianYuService().Attend(pl.GetId())
	// if isLineUp {
	// 	pl.LianYuLineUp(isLineUp)
	// 	scMsg := pbutil.BuildSCLianYuLineUp(beforeNum)
	// 	pl.SendMsg(scMsg)
	// 	return
	// }

	pos, flag := lianyu.GetLianYuService().GetRebornPos(pl.GetId())
	if !flag {
		return
	}

	if !scenelogic.PlayerEnterScene(pl, s, pos) {
		return
	}

	success = true
	return
}
