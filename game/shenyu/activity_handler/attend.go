package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/shenyu/shenyu"
	shenyutemplate "fgame/fgame/game/shenyu/template"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeShenYu, activity.ActivityAttendHandlerFunc(playerEnterShenYu))
}

func playerEnterShenYu(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	activityTimeTemplate, _ := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
	if activityTimeTemplate == nil {
		return
	}
	startTime, _ := activityTimeTemplate.GetBeginTime(now)
	endTime, _ := activityTimeTemplate.GetEndTime(now)
	shenYuTemp := shenyutemplate.GetShenYuTemplateService().GetShenYuInitRoundTemplate()
	sceneEndTime := startTime + int64(shenYuTemp.RoundTime)
	if now > sceneEndTime {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"Mapid":        activityTemplate.Mapid,
				"sceneEndTime": sceneEndTime,
			}).Warn("shenyu:第一轮神域之战已经结束")
		return
	}

	s := shenyu.GetShenYuService().GetShenYuScene()
	if s == nil {
		//创建神域场景
		s = shenyu.GetShenYuService().CreateShenYuScene(activityTemplate.Mapid, endTime, sceneEndTime)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"Mapid":    activityTemplate.Mapid,
					"endTime":  endTime,
				}).Warn("shenyu:神域之战场景不存在")
			return
		}
	}

	bornPos := s.MapTemplate().GetMap().RandomPosition()
	if !scenelogic.PlayerEnterScene(pl, s, bornPos) {
		return
	}
	flag = true
	return
}
