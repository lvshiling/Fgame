package handler

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	playeralliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeYaoPai, command.CommandHandlerFunc(handleYaoPai))
}

func handleYaoPai(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理腰牌")
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	yaoPaiStr := c.Args[0]
	yaoPai, err := strconv.ParseInt(yaoPaiStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"yaoPai": yaoPai,
				"error":  err,
			}).Warn("gm:处理腰牌,腰牌id不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = setYaoPai(pl, int32(yaoPai))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"yaoPai": yaoPai,
				"error":  err,
			}).Warn("gm:处理腰牌,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":     pl.GetId(),
			"yaoPai": yaoPai,
		}).Debug("gm:处理腰牌,完成")
	return
}

func setYaoPai(p scene.Player, yaoPai int32) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)

	currentYaoPai := manager.GetYaoPai()
	needAdd := yaoPai - currentYaoPai
	if needAdd == 0 {
		return
	}
	if needAdd < 0 {
		reasonText := commonlog.YaoPaiLogReasonGM.String()
		flag := manager.CostYaoPai(-needAdd, commonlog.YaoPaiLogReasonGM, reasonText)
		if !flag {
			panic(fmt.Errorf("gm:cost 腰牌 should be ok"))
		}

	} else {
		reasonText := commonlog.YaoPaiLogReasonGM.String()
		manager.AddYaoPai(needAdd, commonlog.YaoPaiLogReasonGM, reasonText)
	}

	return
}
