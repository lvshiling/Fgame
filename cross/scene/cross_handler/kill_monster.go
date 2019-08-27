package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLYAER_KILL_BIOLOGY_TYPE), dispatch.HandlerFunc(handlePlayerKillBiology))
}

func handlePlayerKillBiology(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:怪物被杀")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("scene:怪物被杀,完成")
	return nil
}
