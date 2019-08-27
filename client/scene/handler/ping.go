package handler

import (
	"fgame/fgame/client/player/player"
	"fgame/fgame/client/processor"
	clientsession "fgame/fgame/client/session"
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_SC_PING_TYPE), dispatch.HandlerFunc(handlePing))
}
func handlePing(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:ping")
	cs := clientsession.SessionInContext(s.Context())
	pl := cs.Player().(*player.Player)

	err = playerPing(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.Id(),
			}).Debug("scene:ping")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.Id(),
		}).Debug("scene:ping")
	return nil
}

func playerPing(pl *player.Player) (err error) {
	pl.Ping()
	return
}
