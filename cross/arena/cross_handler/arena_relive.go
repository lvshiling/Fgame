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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_ARENA_RELIVE_TYPE), dispatch.HandlerFunc(handleArenaRelive))
}

//竞技场复活
func handleArenaRelive(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:竞技场复活")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	siArenaRelive := msg.(*crosspb.SIArenaRelive)
	err = arenaRelive(tpl, siArenaRelive.GetSuccess())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:竞技场复活,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:竞技场复活,完成")
	return nil

}

//竞技场获胜
func arenaRelive(pl *player.Player, suceess bool) (err error) {
	if suceess {
		pl.Reborn(pl.GetPos())
	} else {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Info("arena:竞技场复活,失败")
		//TODO 是否需要退出场景
	}
	return
}
