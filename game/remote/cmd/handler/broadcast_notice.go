package handler

import (
	"fgame/fgame/game/notice/notice"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_BROADCAST_NOTICE), cmd.CmdHandlerFunc(handleBroadcastNotice))
}

// 系统公告
func handleBroadcastNotice(msg proto.Message) (err error) {
	log.Info("cmd:请求发送系统公告")
	cmdMsg := msg.(*cmdpb.CmdBroadcastNotice)
	content := cmdMsg.GetContent()
	beginTime := cmdMsg.GetBeginTime()
	endTime := cmdMsg.GetEndTime()
	interval := cmdMsg.GetIntervalTime()

	if len(content) <= 0 {
		err = cmd.ErrorCodeCommonArgumentInvalid
		log.WithFields(
			log.Fields{
				"content":   content,
				"beginTime": beginTime,
				"endTime":   endTime,
				"interval":  interval,
				"err":       err,
			}).Warn("cmd:请求发送系统公告,内容错误")
		return
	}
	if beginTime < 0 || beginTime > endTime {
		err = cmd.ErrorCodeCommonArgumentInvalid
		log.WithFields(
			log.Fields{
				"content":   content,
				"beginTime": beginTime,
				"endTime":   endTime,
				"interval":  interval,
				"err":       err,
			}).Warn("cmd:请求发送系统公告,公告发送时间错误")
		return
	}

	notice.GetNoticeService().AddGmNotice(content, beginTime, endTime, interval)

	log.WithFields(
		log.Fields{
			"content":   content,
			"beginTime": beginTime,
			"endTime":   endTime,
			"interval":  interval,
		}).Info("cmd:请求发送系统公告,成功")

	return
}
