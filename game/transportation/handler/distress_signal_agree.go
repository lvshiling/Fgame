package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/transportation/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_AGREE_DISTRESS_SIGNAL_TYPE), dispatch.HandlerFunc(handlerAgreeDistressSignal))
}

//处理穿云箭推送
func handlerAgreeDistressSignal(s session.Session, msg interface{}) (err error) {
	log.Debug("transport:处理穿云箭推送")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csAgreeDistressSignal := msg.(*uipb.CSAgreeDistressSignal)
	targetId := csAgreeDistressSignal.GetTargetId()

	err = agreeDistressSignal(tpl, targetId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"targetId": targetId,
				"err":      err,
			}).Error("transport:处理穿云箭推送，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"targetId": targetId,
		}).Debug("transport:处理穿云箭推送完成")

	return
}

func agreeDistressSignal(pl player.Player, targetId int64) (err error) {
	targetPlayer := player.GetOnlinePlayerManager().GetPlayerById(targetId)
	if targetPlayer == nil {
		return
	}

	targetScene := targetPlayer.GetScene()
	if targetScene == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"targetId": targetId,
			}).Warn("transport:进入押镖场景失败,求援玩家不在场景中")
		return
	}

	targetPosition := targetPlayer.GetPosition()
	if !scenelogic.PlayerEnterScene(pl, targetScene, targetPosition) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"targetId": targetId,
			}).Warn("transport:进入押镖场景失败")
		return
	}

	//TODO:xzk:PK模式改为仙盟PK模式

	scAgreeDistressSignal := pbutil.BuildSCAgreeDistressSignal(targetPosition)
	pl.SendMsg(scAgreeDistressSignal)

	return
}
