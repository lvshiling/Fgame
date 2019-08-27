package handler

import (
	"fgame/fgame/common/codec"

	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	"fgame/fgame/game/processor"

	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_PING_TYPE), dispatch.HandlerFunc(handlePing))
}

//处理ping
func handlePing(s session.Session, msg interface{}) error {
	log.Debug("scene:处理ping消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	ping := pbutil.BuildPing()
	tpl.SendMsg(ping)
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("scene:处理ping消息,完成")
	return nil
}
