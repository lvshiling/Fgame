package handler

import (
	"fgame/fgame/common/lang"

	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	shengtanscene "fgame/fgame/game/shengtan/scene"
	"fgame/fgame/game/shengtan/shengtan"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeShengTan, command.CommandHandlerFunc(handleShengTan))
}

func handleShengTan(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:仙盟圣坛")

	err = shengTan(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:仙盟圣坛,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:仙盟圣坛,完成")
	return
}

func shengTan(p scene.Player) (err error) {
	pl := p.(player.Player)
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("gm:仙盟圣坛,玩家不在仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}
	allActivities := activitytemplate.GetActivityTemplateService().GetActiveAll()
	now := global.GetGame().GetTimeService().Now()
	var sd shengtanscene.ShengTanSceneData
Loop:
	for _, act := range allActivities {
		if act.GetActivityType() != activitytypes.ActivityTypeAllianceShengTan {
			continue
		}

		for _, timeTemplate := range act.GetTimeList() {
			endTime := timeTemplate.ActivityTimes + now
			sd = shengtan.GetShengTanService().CreateShengTanScene(allianceId, act.Mapid, endTime)
			if sd != nil {
				break Loop
			}
		}
	}

	if sd == nil {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("gm:仙盟圣坛,未开启")
		return
	}

	s := sd.GetScene()
	pos := s.MapTemplate().GetBornPos()
	scenelogic.PlayerEnterScene(pl, s, pos)

	return
}
