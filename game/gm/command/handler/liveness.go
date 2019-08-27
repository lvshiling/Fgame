package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/liveness/pbutil"
	playerliveness "fgame/fgame/game/liveness/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLiveness, command.CommandHandlerFunc(handleLivenessSet))
}

func handleLivenessSet(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置活跃度,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if level < 0 {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"level": levelStr,
				"error": err,
			}).Warn("gm:处理设置活跃度,level小于等于0")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLivenessDataManagerType).(*playerliveness.PlayerLivenessDataManager)
	manager.GmSetLiveness(level)

	livenessObj := manager.GetLiveness()
	livenessMap := manager.GetLivenessQuestMap()
	scLivenessGet := pbutil.BuildSCLivenessGet(livenessObj, livenessMap)
	p.SendMsg(scLivenessGet)
	return
}
