package handler

import (
	"fgame/fgame/common/lang"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/gm/command"
	"fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fmt"

	commonlog "fgame/fgame/common/log"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeItem, command.CommandHandlerFunc(handleItem))
}
func handleItem(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理设置物品数量")
	if len(c.Args) < 2 {
		log.Warn("gm:设置物品数量,参数少于2")

		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	itemIdStr := c.Args[0]
	itemId, err := strconv.ParseInt(itemIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"error":  err,
				"itemId": itemIdStr,
			}).Warn("gm:设置物品数量,itemId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	numStr := c.Args[1]
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
				"num":   numStr,
			}).Warn("gm:设置物品数量,数量不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	if num < 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	//TODO 修改物品数量
	err = setItem(pl, int32(itemId), int32(num))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"error":  err,
				"itemId": itemIdStr,
				"num":    numStr,
			}).Warn("gm:设置物品数量,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":     pl.GetId(),
			"itemId": itemIdStr,
			"num":    numStr,
		}).Debug("gm:处理设置物品数量完成")
	return
}

func setItem(p scene.Player, itemId int32, num int32) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	currentNum := manager.NumOfItems(itemId)
	needAdd := num - currentNum
	if needAdd == 0 {
		return
	}
	if needAdd < 0 {
		flag := manager.UseItem(itemId, -needAdd, commonlog.InventoryLogReasonGM, commonlog.InventoryLogReasonGM.String())
		if !flag {
			panic(fmt.Errorf("gm:use item should be ok"))
		}

	} else {
		itemData := droptemplate.CreateItemData(itemId, needAdd, 0, itemtypes.ItemBindTypeUnBind)
		flag := manager.AddItemLevel(itemData, commonlog.InventoryLogReasonGM, commonlog.InventoryLogReasonGM.String())
		if !flag {
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
	}
	logic.SnapInventoryChanged(pl)
	return
}
