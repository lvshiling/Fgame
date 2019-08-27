package handler

import (
	"fgame/fgame/common/lang"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeEmailAdd, command.CommandHandlerFunc(handleEmailAdd))
}

func handleEmailAdd(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理发送邮件")
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	title := c.Args[0]
	content := c.Args[1]
	attachmentInfo := c.Args[2:]
	if len(attachmentInfo)%2 != 0 {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"title":          title,
				"content":        content,
				"attachmentInfo": attachmentInfo,
			}).Warn("gm:处理发送邮件,附件参数错误")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	attachmentCount := len(attachmentInfo) / 4
	itemIdArr := make([]int32, 0, attachmentCount)
	itemCount := make([]int32, 0, attachmentCount)
	for index, number := range attachmentInfo {
		num, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			log.WithFields(
				log.Fields{
					"id":             pl.GetId(),
					"title":          title,
					"content":        content,
					"attachmentInfo": attachmentInfo,
					"error":          err,
				}).Warn("gm:处理发送邮件，附件信息不是数字")
			playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
			return err
		}

		if index%2 == 0 || index%2 == 2 {
			if !isValid(num) {
				log.WithFields(
					log.Fields{
						"id":             pl.GetId(),
						"title":          title,
						"content":        content,
						"attachmentInfo": attachmentInfo,
					}).Warn("gm:处理发送邮件，附件物品ID不存在")
				playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
				return err
			}
			itemIdArr = append(itemIdArr, int32(num))
		} else {
			itemCount = append(itemCount, int32(num))
		}
	}

	attachmentInfoMap := make(map[int32]int32)
	for index, ietmId := range itemIdArr {
		if _, ok := attachmentInfoMap[ietmId]; ok {
			attachmentInfoMap[ietmId] += itemCount[index]
		} else {
			attachmentInfoMap[ietmId] = itemCount[index]
		}
	}

	err = addEmail(pl, title, content, attachmentInfoMap)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":             pl.GetId(),
				"title":          title,
				"content":        content,
				"attachmentInfo": attachmentInfo,
				"error":          err,
			}).Warn("gm:处理发送邮件,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id":             pl.GetId(),
			"title":          title,
			"content":        content,
			"attachmentInfo": attachmentInfo,
		}).Debug("gm:处理发送邮件,完成")
	return
}

func addEmail(pl player.Player, title string, content string, attachments map[int32]int32) (err error) {
	now := global.GetGame().GetTimeService().Now()
	var newItemList []*droptemplate.DropItemData
	for itemId, num := range attachments {
		newData := droptemplate.CreateItemData(itemId, num, 0, itemtypes.ItemBindTypeBind)
		newItemList = append(newItemList, newData)
	}
	emaillogic.AddEmailItemLevel(pl, title, content, now, newItemList)
	return
}

func isValid(id int64) bool {
	itemTemplate := item.GetItemService().GetItem(int(id))
	if itemTemplate == nil {
		return false
	}
	return true
}
