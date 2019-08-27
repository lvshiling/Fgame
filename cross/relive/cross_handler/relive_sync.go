package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_RELIVE_SYNC_TYPE), dispatch.HandlerFunc(handlerPlayerReliveSync))
}

//复活
func handlerPlayerReliveSync(s session.Session, msg interface{}) (err error) {
	log.Debug("relive:玩家复活同步确认")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("relive:玩家复活同步确认,完成")
	return nil

}
