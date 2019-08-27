package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_AUTO_AGREE_JOIN_TYPE), dispatch.HandlerFunc(handleAllianceAutoAgreeApply))
}

//处理仙盟自动处理申请自动审核申请
func handleAllianceAutoAgreeApply(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟自动处理申请加入")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSAllianceAutoAgreeJoin)
	isAuto := csMsg.GetIsAuto()

	err = allianceAutoAgreeJoin(tpl, isAuto)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isAuto":   isAuto,
				"error":    err,
			}).Error("alliance:处理仙盟自动处理申请加入,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"isAuto":   isAuto,
		}).Debug("alliance:处理仙盟自动处理申请加入,完成")
	return nil
}

//仙盟自动处理申请加入
func allianceAutoAgreeJoin(pl player.Player, isAuto int32) (err error) {
	err = alliance.GetAllianceService().ChangeAutoAgreeJoinApply(pl.GetId(), isAuto)
	if err != nil {
		return
	} 

	scMsg := pbutil.BuildSCAllianceAutoAgreeJoin(isAuto)
	pl.SendMsg(scMsg)
	return
}
