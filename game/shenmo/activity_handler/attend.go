package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	shenmologic "fgame/fgame/game/shenmo/logic"
	"fgame/fgame/game/shenmo/pbutil"
	"fgame/fgame/game/shenmo/shenmo"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeShenMoWar, activity.ActivityAttendHandlerFunc(shenmologic.PlayerEnterShenMoScene))
	activity.RegisterActivityHandler(activitytypes.ActivityTypeLocalShenMoWar, activity.ActivityAttendHandlerFunc(PlayerEnterLocalShenMoScene))
}

//神魔战场-本服
func PlayerEnterLocalShenMoScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (success bool, err error) {
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

	s := shenmo.GetShenMoService().GetShenMoScene()
	if s == nil {
		s = shenmo.GetShenMoService().CreateShenMoScene(activityTemplate.Mapid, endTime)
		if s == nil {
			return false, err
		}
	}

	shenMoEndTime := pl.GetShenMoEndTime()
	if endTime != shenMoEndTime {
		pl.SetShenMoEndTime(endTime)
	}

	beforeNum, isLineUp := shenmo.GetShenMoService().Attend(pl.GetId())
	if isLineUp {
		pl.ShenMoLineUp(isLineUp)
		scMsg := pbutil.BuildSCShenMoLineUp(beforeNum)
		pl.SendMsg(scMsg)
		return
	}

	pos := s.MapTemplate().GetBornPos()
	if !scenelogic.PlayerEnterScene(pl, s, pos) {
		return
	}
	success = true
	return
}
