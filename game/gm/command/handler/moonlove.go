package handler

import (
	"fgame/fgame/common/lang"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	moonlovelogic "fgame/fgame/game/moonlove/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeMoonloveJoin, command.CommandHandlerFunc(handleMoonloveJoin))
}

func handleMoonloveJoin(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:进入月下情缘")
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	activeIdStr := c.Args[0]
	activeId, err := strconv.ParseInt(activeIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":       pl.GetId(),
				"activeId": activeIdStr,
				"error":    err,
			}).Warn("gm:处理进入月下情缘，活动id不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	activeType := activitytypes.ActivityType(activeId)
	err = moonloveJoin(pl, activeType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:进入月下情缘,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:进入月下情缘,完成")
	return
}

func moonloveJoin(pl player.Player, activeType activitytypes.ActivityType) (err error) {

	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activeType)
	_, err = moonlovelogic.PlayerEnterMoonlove(pl, activityTemplate)
	if err != nil {
		return
	}
	return
}
