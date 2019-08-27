package handler

import (
	compensatelogic "fgame/fgame/game/compensate/logic"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_SEND_SERVER_COMPENSATE_TYPE), cmd.CmdHandlerFunc(handleServerCompensate))
}

// 全服补偿
func handleServerCompensate(msg proto.Message) (err error) {
	log.Info("cmd:请求添加全服补偿")
	cmdMsg := msg.(*cmdpb.CmdSendServerCompensate)
	title := cmdMsg.GetTitle()
	content := cmdMsg.GetContent()
	attachmentStr := cmdMsg.GetAttachmentStr()
	needLevel := cmdMsg.GetNeedLevel()
	needCreateTime := cmdMsg.GetNeedCreateTime()
	bind := cmdMsg.GetBind()
	attachmentList, err := parseAttachmentList(attachmentStr, bind)
	if err != nil {
		err = cmd.ErrorCodeMailAttachmentFormatWrong
		log.WithFields(
			log.Fields{
				"needLevel":      needLevel,
				"needCreateTime": needCreateTime,
				"title":          title,
				"content":        content,
				"attachmentStr":  attachmentStr,
				"err":            err,
			}).Warn("cmd:请求发玩家补偿,附件格式错误")
		return
	}

	err = serverCompensate(needLevel, needCreateTime, title, content, attachmentList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"needLevel":      needLevel,
				"needCreateTime": needCreateTime,
				"title":          title,
				"content":        content,
				"attachmentStr":  attachmentStr,
				"err":            err,
			}).Error("cmd:请求添加全服补偿,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"needLevel":      needLevel,
			"needCreateTime": needCreateTime,
			"title":          title,
			"content":        content,
			"attachmentStr":  attachmentStr,
		}).Info("cmd:请求添加全服补偿,成功")
	return
}

func serverCompensate(needLevel int32, needCreateTime int64, title string, content string, attachmentList []*droptemplate.DropItemData) (err error) {

	compensatelogic.AddServerCompensate(needLevel, needCreateTime, title, content, attachmentList)

	return nil
}
