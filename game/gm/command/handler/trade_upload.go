package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"

	inventorytypes "fgame/fgame/game/inventory/types"
	tradelogic "fgame/fgame/game/trade/logic"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTradeUpload, command.CommandHandlerFunc(handleTradeUpload))
}

func handleTradeUpload(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理交易上传")
	tpl, ok := pl.(player.Player)
	if !ok {
		log.Warn("gm:处理交易上传,不是玩家")
		return
	}

	if len(c.Args) < 4 {
		log.Warn("gm:处理交易上传,参数少于1")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	bagTypeStr := c.Args[0]

	bagTypeInt, err := strconv.ParseInt(bagTypeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"bagType": bagTypeStr,
			}).Warn("gm:处理交易上传,bagType不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	bagType := inventorytypes.BagType(bagTypeInt)
	if !bagType.Valid() {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"error":   err,
				"bagType": bagTypeStr,
			}).Warn("gm:处理交易上传,bagType无效")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	indexStr := c.Args[1]

	index, err := strconv.ParseInt(indexStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"index": indexStr,
			}).Warn("gm:处理交易上传,index不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	numStr := c.Args[2]

	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"num":   numStr,
			}).Warn("gm:处理交易上传,num不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	goldStr := c.Args[3]

	gold, err := strconv.ParseInt(goldStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"num":   numStr,
			}).Warn("gm:处理交易上传,gold不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	tempGold := int32(gold)
	err = tradeUpload(tpl, bagType, int32(index), int32(num), tempGold)

	log.WithFields(
		log.Fields{
			"id":      pl.GetId(),
			"bagType": bagType,
			"index":   indexStr,
			"num":     numStr,
		}).Debug("gm:处理交易上传完成")
	return
}

func tradeUpload(p player.Player, bagType inventorytypes.BagType, index int32, num int32, gold int32) error {
	err := tradelogic.TradeUploadItem(p, bagType, index, num, gold)
	if err != nil {
		return err
	}
	return nil
}
