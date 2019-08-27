package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/liveness/pbutil"
	playerliveness "fgame/fgame/game/liveness/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LIVENESS_GET_TYPE), dispatch.HandlerFunc(handleLivenessGet))
}

//处理活跃度信息
func handleLivenessGet(s session.Session, msg interface{}) (err error) {
	log.Debug("liveness:处理活跃度消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = livenessGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("liveness:处理活跃度消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("liveness:处理活跃度消息完成")
	return nil
}

//活跃度信息的逻辑
func livenessGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerLivenessDataManagerType).(*playerliveness.PlayerLivenessDataManager)

	livenessObj := manager.GetLiveness()
	livenessMap := manager.GetLivenessQuestMap()
	scLivenessGet := pbutil.BuildSCLivenessGet(livenessObj, livenessMap)
	pl.SendMsg(scLivenessGet)
	return
}
