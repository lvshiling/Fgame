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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_ARENA_GIVE_UP_TYPE), dispatch.HandlerFunc(handleArenaGiveUp))
}

//竞技场放弃/退出
func handleArenaGiveUp(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:竞技场放弃/退出")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = arenaGiveUp(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:竞技场放弃/退出,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:竞技场放弃/退出,完成")
	return nil

}

//竞技场放弃/退出
func arenaGiveUp(pl *player.Player) (err error) {
	return
}
