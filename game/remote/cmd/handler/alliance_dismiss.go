package handler

import (
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/remote/cmd"
	cmdpb "fgame/fgame/game/remote/cmd/pb"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func init() {
	cmd.RegisterCmdHandler(cmd.CmdType(cmdpb.CmdType_CMD_ALLIANCE_DISMISS_TYPE), cmd.CmdHandlerFunc(handleAllianceDismiss))
}

func handleAllianceDismiss(msg proto.Message) (err error) {
	log.Info("cmd:请求仙盟解散")
	cmdAllianceDismiss := msg.(*cmdpb.CmdAllianceDismiss)
	allianceId := cmdAllianceDismiss.GetAllianceId()

	err = allianceDismiss(allianceId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"allianceId": allianceId,

				"err": err,
			}).Error("cmd:请求仙盟解散,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"allianceId": allianceId,
		}).Info("cmd:请求仙盟解散，成功")
	return
}

func allianceDismiss(allianceId int64) (err error) {

	al, memList, err := alliance.GetAllianceService().GMDismissAlliance(allianceId)
	if err != nil {
		err = cmd.ErrorCodeCommonAllianceActivity
		log.WithFields(
			log.Fields{
				"allianceId": allianceId,
				"err":        err,
			}).Warn("cmd:请求仙盟解散，活动中")
		return
	}

	//解散广播
	for _, member := range memList {
		memberPlayer := player.GetOnlinePlayerManager().GetPlayerById(member.GetMemberId())
		if memberPlayer == nil {
			continue
		}

		dismissBroadcast := pbutil.BuildSCAllianceDismissBroadcast(al.GetAllianceId())
		memberPlayer.SendMsg(dismissBroadcast)
	}
	return
}
