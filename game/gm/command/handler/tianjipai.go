package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTianJiPai, command.CommandHandlerFunc(handleTianJiPai))
}

func handleTianJiPai(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理完成天机牌任务")
	pl := p.(player.Player)
	if len(c.Args) <= 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	subTypeStr := c.Args[0]
	numStr := c.Args[1]
	subType, err := strconv.ParseInt(subTypeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"subType": subType,
				"error":   err,
			}).Warn("gm:处理天机牌任务,子类型subType不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"num":   num,
				"error": err,
			}).Warn("gm:处理天机牌任务,num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	questSubType := questtypes.QuestSubType(subType)
	if !questSubType.Valid() {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"subType": subType,
				"error":   err,
			}).Warn("gm:处理天机牌任务,子类型subType不是有效的")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
	}
	questlogic.GMIncreaseQuestData(pl, questSubType, 0, int32(num))

	log.WithFields(
		log.Fields{
			"id":      pl.GetId(),
			"subType": subType,
			"num":     num,
		}).Debug("gm:处理天机牌任务,完成")
	return
}
