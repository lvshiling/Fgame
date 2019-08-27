package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_XUE_CHI_ADD_TYPE), dispatch.HandlerFunc(handleXueChiAdd))
}

//血池加血
func handleXueChiAdd(s session.Session, msg interface{}) (err error) {
	log.Debug("xuechi:血池加血")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)

	siXueChiAdd := msg.(*crosspb.SIXueChiAdd)
	addBlood := siXueChiAdd.GetBlood()

	err = playerAddBlood(tpl, addBlood)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"addBlood": addBlood,
				"error":    err,
			}).Error("xuechi:血池加血,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"addBlood": addBlood,
		}).Debug("xuechi:血池加血完成")
	return nil
}

//处理设置血池线逻辑
func playerAddBlood(pl scene.Player, addblood int64) (err error) {
	pl.AddBlood(addblood)
	return
}
