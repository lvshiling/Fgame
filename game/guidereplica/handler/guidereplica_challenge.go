package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	guidereplicalogic "fgame/fgame/game/guidereplica/logic"
	"fgame/fgame/game/guidereplica/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GUIDE_REPLICA_CHALLENGE_TYPE), dispatch.HandlerFunc(handlerGuideReplicaChallenge))
}

//引导副本挑战
func handlerGuideReplicaChallenge(s session.Session, msg interface{}) (err error) {
	log.Debug("guidereplica:处理引导副本挑战请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGuideReplicaChallenge)
	questId := csMsg.GetQuestId()

	err = guidereplicaChallenge(tpl, questId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"questId":  questId,
				"err":      err,
			}).Error("guidereplica:处理引导副本挑战请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"questId":  questId,
		}).Debug("guidereplica:处理引导副本挑战请求完成")

	return
}

//引导副本挑战逻辑
func guidereplicaChallenge(pl player.Player, questId int32) (err error) {

	//进入场景
	flag := guidereplicalogic.PlayerEnterQuestGuideReplica(pl, questId)
	if !flag {
		return
	}

	scMsg := pbutil.BuildSCGuideReplicaChallenge()
	pl.SendMsg(scMsg)
	return
}
