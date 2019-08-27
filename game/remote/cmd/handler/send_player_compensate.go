package handler

import (
	compensatelogic "fgame/fgame/game/compensate/logic"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fmt"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_SEND_PLAYER_COMPENSATE_TYPE), cmd.CmdHandlerFunc(handlePlayerCompensate))
}

// 部分补偿
func handlePlayerCompensate(msg proto.Message) (err error) {
	log.Info("cmd:请求发玩家补偿")
	cmdMsg := msg.(*cmdpb.CmdSendPlayerCompensate)
	playerIdStrList := cmdMsg.GetPlayerIdList()
	title := cmdMsg.GetTitle()

	content := cmdMsg.GetContent()
	attachmentStr := cmdMsg.GetAttachmentStr()
	bind := cmdMsg.GetBind()
	if len(title) <= 0 {
		err = cmd.ErrorCodeMailFormatWrong
		log.WithFields(
			log.Fields{
				"playerIdStr":   playerIdStrList,
				"title":         title,
				"content":       content,
				"attachmentStr": attachmentStr,
				"err":           err,
			}).Warn("cmd:请求发玩家补偿,邮件格式错误")
		return
	}

	if len(content) <= 0 {
		err = cmd.ErrorCodeMailFormatWrong
		log.WithFields(
			log.Fields{
				"playerIdStr":   playerIdStrList,
				"title":         title,
				"content":       content,
				"attachmentStr": attachmentStr,
				"err":           err,
			}).Warn("cmd:请求发玩家补偿,邮件格式错误")
		return
	}

	playerIdList, err := parsePlayerList(playerIdStrList)
	if err != nil {
		err = cmd.ErrorCodeMailPlayerFormatWrong
		log.WithFields(
			log.Fields{
				"playerIdStr":   playerIdStrList,
				"title":         title,
				"content":       content,
				"attachmentStr": attachmentStr,
				"err":           err,
			}).Warn("cmd:请求发玩家补偿,玩家格式错误")
		return
	}
	attachmentList, err := parseAttachmentList(attachmentStr, bind)
	if err != nil {
		err = cmd.ErrorCodeMailAttachmentFormatWrong
		log.WithFields(
			log.Fields{
				"playerIdStr":   playerIdStrList,
				"title":         title,
				"content":       content,
				"attachmentStr": attachmentStr,
				"err":           err,
			}).Warn("cmd:请求发玩家补偿,附件格式错误")
		return
	}
	err = playerListCompensate(playerIdList, title, content, attachmentList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerIdStr":   playerIdStrList,
				"title":         title,
				"content":       content,
				"attachmentStr": attachmentStr,
				"err":           err,
			}).Error("cmd:请求发玩家补偿,发送补偿错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerIdStr":   playerIdStrList,
			"title":         title,
			"content":       content,
			"attachmentStr": attachmentStr,
		}).Info("cmd:请求发玩家补偿,成功")

	return
}

func playerListCompensate(playerIdList []int64, title string, content string, attachmentList []*droptemplate.DropItemData) (err error) {
	now := global.GetGame().GetTimeService().Now()
	for _, playerId := range playerIdList {
		compensatelogic.SendPlayerCompensateEmail(playerId, title, content, now, attachmentList)
	}
	return nil
}

func parsePlayerList(playerListStr string) (playerIdList []int64, err error) {
	playerIdStrList := strings.Split(playerListStr, ",")
	for _, playerIdStr := range playerIdStrList {
		playerId, err := strconv.ParseInt(playerIdStr, 10, 64)
		if err != nil {
			return nil, err
		}
		playerIdList = append(playerIdList, playerId)
	}
	return
}

func parseAttachmentList(itemStr string, bind bool) (itemDataList []*droptemplate.DropItemData, err error) {
	if len(itemStr) == 0 {
		return nil, nil
	}
	itemArr := strings.Split(itemStr, ",")
	for _, tempItem := range itemArr {
		itemNumArr := strings.Split(tempItem, ":")
		if len(itemNumArr) != 2 {
			return nil, fmt.Errorf("格式不对[%s]", tempItem)
		}
		itemId, err := strconv.ParseInt(itemNumArr[0], 10, 64)
		if err != nil {
			return nil, err
		}
		itemIdInt := int32(itemId)

		itemTemplate := item.GetItemService().GetItem(int(itemIdInt))
		if itemTemplate == nil {
			return nil, fmt.Errorf("物品不存在[%d]", itemIdInt)
		}

		itemNum, err := strconv.ParseInt(itemNumArr[1], 10, 64)
		if err != nil {
			return nil, err
		}
		bindType := itemtypes.ItemBindTypeUnBind
		if bind {
			bindType = itemtypes.ItemBindTypeBind
		}
		itemNumInt := int32(itemNum)
		itemData := droptemplate.CreateItemData(int32(itemId), itemNumInt, 0, bindType)
		itemDataList = append(itemDataList, itemData)
	}
	return
}
