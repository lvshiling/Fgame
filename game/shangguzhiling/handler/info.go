package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shangguzhilinglogic "fgame/fgame/game/shangguzhiling/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_INFO_TYPE), dispatch.HandlerFunc(handleShangguzhilingInfo))
}

func handleShangguzhilingInfo(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = shangguzhilingInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shangguzhiling:上古之灵信息请求,错误")

		return err
	}
	return nil
}

func shangguzhilingInfo(pl player.Player) (err error) {
	err = shangguzhilinglogic.SendLingShouInfo(pl)
	return
}
