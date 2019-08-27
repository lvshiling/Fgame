package activity_handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/activity/activity"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/shengtan/shengtan"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeAllianceShengTan, activity.ActivityAttendHandlerFunc(playerEnterAllianceShengTanScene))
}

func playerEnterAllianceShengTanScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("shengtan:仙盟圣坛,玩家不在仙盟")
		playerlogic.SendSystemMessage(pl, lang.ShengTanAllianceUserNotInAlliance)
		return
	}

	sd := shengtan.GetShengTanService().GetShengTanScene(allianceId)
	if sd == nil {
		act := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeAllianceShengTan)
		now := global.GetGame().GetTimeService().Now()
		openTime := global.GetGame().GetServerTime()
		mergeTime := merge.GetMergeService().GetMergeTime()
		timeTemplate, err := act.GetActivityTimeTemplate(now, openTime, mergeTime)
		if err != nil {
			return false, err
		}
		if timeTemplate == nil {
			log.WithFields(
				log.Fields{
					"id": pl.GetId(),
				}).Warn("shengtan:仙盟圣坛,城战未开始")
			playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
			return false, nil
		}
		endTime, err := timeTemplate.GetEndTime(now)
		if err != nil {
			return false, err
		}

		sd = shengtan.GetShengTanService().CreateShengTanScene(allianceId, act.Mapid, endTime)
		if sd == nil {
			log.WithFields(
				log.Fields{
					"id": pl.GetId(),
				}).Warn("shengtan:仙盟圣坛,未开启")
			playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
			return false, nil
		}
	}
	s := sd.GetScene()
	pos := s.MapTemplate().GetBornPos()

	if !scenelogic.PlayerEnterScene(pl, s, pos) {
		return
	}
	flag = true
	return
}
