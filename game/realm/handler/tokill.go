package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	realmlogic "fgame/fgame/game/realm/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_REALM_TOKILL_TYPE), dispatch.HandlerFunc(handleRealmToKill))
}

//处理前往击杀信息
func handleRealmToKill(s session.Session, msg interface{}) (err error) {
	log.Debug("realm:处理获取前往击杀消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = realmToKill(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("realm:处理获取前往击杀消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("realm:处理获取前往击杀消息完成")
	return nil

}

//获取前往击杀界面信息的逻辑
func realmToKill(pl player.Player) (err error) {
	return realmlogic.HandleTianJieTa(pl)
}
