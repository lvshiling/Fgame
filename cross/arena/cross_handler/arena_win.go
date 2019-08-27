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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_ARENA_WIN_TYPE), dispatch.HandlerFunc(handleArenaWin))
}

//竞技场获胜
func handleArenaWin(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:竞技场获胜")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = arenaWin(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:竞技场获胜,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:竞技场获胜,完成")
	return nil

}

//竞技场获胜
func arenaWin(pl *player.Player) (err error) {
	//判断是否有组队

	return
}
