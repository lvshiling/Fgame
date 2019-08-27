package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTuMoClearNum, command.CommandHandlerFunc(handleTuMoClearNum))
}

//屠魔次数置为0
func handleTuMoClearNum(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空屠魔次数")

	err = clearTuMoNum(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空屠魔次数,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空屠魔次数完成")
	return
}

//屠魔次数置为0
func clearTuMoNum(p scene.Player) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	manager.GMClearTuMoNum()

	num, buyNum, leftNum := manager.GetTuMoNum()
	scQuestTuMoNumGet := pbutil.BuildSCQuestTuMoNumGet(num, buyNum, leftNum)
	pl.SendMsg(scQuestTuMoNumGet)
	return
}
