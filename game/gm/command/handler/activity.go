package handler

import (
	"fgame/fgame/common/lang"
	activitylogic "fgame/fgame/game/activity/logic"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeActivity, command.CommandHandlerFunc(handleActivity))
}

func handleActivity(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) > 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	typStr := c.Args[0]
	typ, err := strconv.ParseInt(typStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"typ":   typ,
				"error": err,
			}).Warn("gm:处理活动,类型typ不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	//参数验证
	activeType := activitytypes.ActivityType(typ)

	if !activeType.Valid() {
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"activeType": activeType,
			}).Warn("activityGm:活动类型,错误")
		return
	}

	return activitylogic.HandleActiveAttend(pl, activeType, "")
}
