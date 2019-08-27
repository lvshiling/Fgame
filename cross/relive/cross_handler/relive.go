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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_RELIVE_TYPE), dispatch.HandlerFunc(handlerPlayerRelive))
}

//复活
func handlerPlayerRelive(s session.Session, msg interface{}) (err error) {
	log.Debug("relive:玩家复活")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	siPlayerRelive := msg.(*crosspb.SIPlayerRelive)
	succ := siPlayerRelive.GetSuccess()
	err = playerRelive(tpl, succ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("relive:玩家复活,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("relive:玩家复活,完成")
	return nil

}

//玩家复活
func playerRelive(pl *player.Player, succ bool) (err error) {
	if !succ {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Debug("relive:玩家复活,失败")
		return
	}
	if !pl.IsDead() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Debug("relive:玩家复活,已经复活")
		return
	}
	//复活
	pl.Reborn(pl.GetPos())
	pl.Relive()
	
	return
}
