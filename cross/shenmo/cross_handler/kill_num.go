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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_SHENMO_KILLNUM_CHANGED_TYPE), dispatch.HandlerFunc(handleShenMoKillNumChanged))
}

//处理神魔战场击杀的人数
func handleShenMoKillNumChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理神魔战场击杀的人数")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = shenMoKillNumChanged(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("shenmo:处理神魔战场击杀的人数,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenmo:处理神魔战场击杀的人数,完成")
	return nil
}

//神魔战场击杀的人数
func shenMoKillNumChanged(pl *player.Player) (err error) {
	return
}
