package handler

import (
	"fgame/fgame/common/lang"
	activitytemplate "fgame/fgame/game/activity/template"
	activetypes "fgame/fgame/game/activity/types"
	fourgodlogic "fgame/fgame/game/fourgod/logic"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeFourGod, command.CommandHandlerFunc(handleFourGod))
}

func handleFourGod(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理四神遗迹")

	err = enterFourGod(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理跳转四神遗迹,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理跳转四神遗迹,完成")
	return
}

func enterFourGod(p scene.Player) (err error) {
	pl := p.(player.Player)
	activeTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activetypes.ActivityTypeFourGod)
	flag, err := fourgodlogic.PlayerEnterFourGodScene(pl, activeTemplate)
	if err != nil {
		return
	}
	if !flag {
		playerlogic.SendSystemMessage(pl, lang.FourGodFixedTime)
		return
	}
	return
}
