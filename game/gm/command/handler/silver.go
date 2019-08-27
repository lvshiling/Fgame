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

	command.Register(gmcommandtypes.CommandTypeSilver, command.CommandHandlerFunc(handleSilver))
}

func handleSilver(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理设置银两数量")
	if len(c.Args[0]) <= 0 {
		log.Warn("gm:处理设置银两数量,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	silverStr := c.Args[0]
	silver, err := strconv.ParseInt(silverStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"error":  err,
				"silver": silverStr,
			}).Warn("gm:处理设置银两数量,silver不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	if silver < 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	//TODO 修改物品数量
	err = setSilver(pl, silver)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"error":  err,
				"silver": silver,
			}).Warn("gm:设置银两数量,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":     pl.GetId(),
			"silver": silver,
		}).Debug("gm:设置银两数量,完成")
	return
}

func setSilver(p scene.Player, silver int64) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	currentSilver := manager.GetSilver()
	needAdd := silver - currentSilver
	if needAdd == 0 {
		return
	}
	if needAdd < 0 {
		reasonText := commonlog.SilverLogReasonGM.String()
		flag := manager.CostSilver(-needAdd, commonlog.SilverLogReasonGM, reasonText)
		if !flag {
			panic(fmt.Errorf("gm:cost silver should be ok"))
		}

	} else {
		reasonText := commonlog.SilverLogReasonGM.String()
		manager.AddSilver(needAdd, commonlog.SilverLogReasonGM, reasonText)
	}
	propertylogic.SnapChangedProperty(pl)
	return
}
