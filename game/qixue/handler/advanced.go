package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	qixuelogic "fgame/fgame/game/qixue/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QIXUE_ADVANCED_TYPE), dispatch.HandlerFunc(handleQiXueAdvanced))
}

//处理泣血枪进阶信息
func handleQiXueAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("qixue:处理获取泣血枪进阶消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = qixuelogic.HandleQiXueAdvanced(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("qixue:处理获取泣血枪进阶消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("qixue:处理获取泣血枪进阶消息完成")
	return nil

}
