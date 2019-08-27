package handler

import (
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fgame/fgame/game/trade/trade"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_CUSTOM_TRADE_RECYCLE_GOLD_TYPE), cmd.CmdHandlerFunc(handleCustomRecycleGoldSet))
}

func handleCustomRecycleGoldSet(msg proto.Message) (err error) {
	log.Info("cmd:请求自定义回购池设置")
	cmdChatSet := msg.(*cmdpb.CmdCustomRecycleGold)
	recycleGold := cmdChatSet.GetGold()

	err = customRecycleGoldSet(recycleGold)
	if err != nil {
		log.WithFields(
			log.Fields{
				"recycleGold": recycleGold,
				"err":         err,
			}).Error("cmd:请求自定义回购池设置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"recycleGold": recycleGold,
		}).Info("cmd:请求自定义回购池设置，成功")
	return
}

func customRecycleGoldSet(recycleGold int64) (err error) {
	trade.GetTradeService().GMSetCustomRecycle(recycleGold)
	return
}
