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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_TULONG_KILL_BOSS_TYPE), dispatch.HandlerFunc(handleTuLongKillBoss))
}

func handleTuLongKillBoss(s session.Session, msg interface{}) (err error) {
	log.Debug("tulong:屠龙击杀boss")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("tulong:屠龙击杀boss,完成")
	return nil
}
