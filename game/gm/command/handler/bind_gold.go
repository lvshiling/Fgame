package handler

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeBindGold, command.CommandHandlerFunc(handleBindGold))
}

func handleBindGold(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理设置绑定元宝数量")
	if len(c.Args[0]) <= 0 {
		log.Warn("gm:处理设置绑定元宝数量,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	goldStr := c.Args[0]
	gold, err := strconv.ParseInt(goldStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"gold":  goldStr,
			}).Warn("gm:处理设置绑定元宝数量,gold不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	if gold < 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	//TODO 修改物品数量
	err = setBindGold(pl, gold)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"gold":  gold,
			}).Warn("gm:设置元宝数量,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":   pl.GetId(),
			"gold": gold,
		}).Debug("gm:设置元宝数量,完成")
	return
}

func setBindGold(p scene.Player, gold int64) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	currentGold := manager.GetBindGlod()
	needAdd := gold - currentGold
	if needAdd == 0 {
		return
	}
	if needAdd < 0 {
		reasonText := commonlog.GoldLogReasonGM.String()
		flag := manager.CostGold(-needAdd, true, commonlog.GoldLogReasonGM, reasonText)
		if !flag {
			panic(fmt.Errorf("gm:cost gold should be ok"))
		}

	} else {
		reasonText := commonlog.GoldLogReasonGM.String()
		manager.AddGold(needAdd, true, commonlog.GoldLogReasonGM, reasonText)
	}
	propertylogic.SnapChangedProperty(pl)

	return
}
