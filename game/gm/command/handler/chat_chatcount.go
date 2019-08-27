package handler

import (
	"fgame/fgame/common/lang"
	playerchat "fgame/fgame/game/chat/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypePlayerChatCount, command.CommandHandlerFunc(handlePlayerChatCount))
}

//发言奖励次数
func handlePlayerChatCount(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理发言奖励次数")
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	numStr := c.Args[0]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   num,
				"error": err,
			}).Warn("gm:处理发言奖励次数,类型num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	err = playerChatCount(pl, int32(num))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理发言奖励次数,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理发言奖励次数完成")
	return
}

func playerChatCount(pl player.Player, count int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerChatDataManagerType).(*playerchat.PlayerChatDataManager)
	manager.GMSetChatCount(count)
	return
}
