package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_XUE_CHI_SYNC_TYPE), dispatch.HandlerFunc(handleXueChiSync))
}

//血池同步
func handleXueChiSync(s session.Session, msg interface{}) (err error) {
	log.Debug("xuechi:血池同步")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)

	isXueChiSync := msg.(*crosspb.ISXueChiSync)
	blood := isXueChiSync.GetXueChiData().GetBlood()
	bloodLine := isXueChiSync.GetXueChiData().GetBloodLine()

	err = playerSyncBlood(tpl, blood, bloodLine)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"blood":     blood,
				"bloodLine": bloodLine,
				"error":     err,
			}).Error("xuechi:血池同步,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"blood":     blood,
			"bloodLine": bloodLine,
		}).Debug("xuechi:血池同步")
	return nil
}

//处理设置血池线逻辑
func playerSyncBlood(pl scene.Player, blood int64, bloodLine int32) (err error) {
	pl.SyncBlood(blood, bloodLine)
	return
}
