package handler

import (
	"fgame/fgame/common/lang"

	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/alliance/alliance"
	alliancescene "fgame/fgame/game/alliance/scene"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeAllianceScene, command.CommandHandlerFunc(handleAllianceScene))
}

func handleAllianceScene(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:九霄城战")

	err = allianceScene(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:九霄城战,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:九霄城战,完成")
	return
}

func allianceScene(p scene.Player) (err error) {
	pl := p.(player.Player)
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("gm:九霄城战,玩家不在仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}
	allActivities := activitytemplate.GetActivityTemplateService().GetActiveAll()
	now := global.GetGame().GetTimeService().Now()
	var sd alliancescene.AllianceSceneData
Loop:
	for _, act := range allActivities {
		if act.GetActivityType() != activitytypes.ActivityTypeAlliance {
			continue
		}

		for _, timeTemplate := range act.GetTimeList() {
			endTime := timeTemplate.ActivityTimes + now
			sd = alliance.GetAllianceService().CreateAllianceSceneData(act.Mapid, endTime, allianceId)
			if sd != nil {
				break Loop
			}
		}
	}

	if sd == nil {
		log.WithFields(
			log.Fields{
				"id": pl.GetId(),
			}).Warn("gm:九霄城战,未开启")
		return
	}

	chengWaiScene := sd.GetScene()
	pos := chengWaiScene.MapTemplate().GetBornPos()
	scenelogic.PlayerEnterScene(pl, chengWaiScene, pos)

	return
}
