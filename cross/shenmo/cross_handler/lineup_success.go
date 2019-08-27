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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_SHENMO_LINEUP_SUCCESS_TYPE), dispatch.HandlerFunc(handleShenMoLineUpSuccess))
}

//处理神魔战场排队成功
func handleShenMoLineUpSuccess(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理神魔战场排队成功")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = shenMoLineUpSuccess(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("shenmo:处理神魔战场排队成功,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenmo:处理神魔战场排队成功,完成")
	return nil

}

//神魔战场排队成功
func shenMoLineUpSuccess(pl *player.Player) (err error) {
	return
}
