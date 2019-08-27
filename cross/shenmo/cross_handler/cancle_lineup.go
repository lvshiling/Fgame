package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/cross/shenmo/pbutil"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/shenmo/shenmo"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_SHENMO_CANCLE_LINEUP_TYPE), dispatch.HandlerFunc(handleShenMoCancleLineUp))
}

//处理神魔战场取消排队
func handleShenMoCancleLineUp(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理神魔战场取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	err = shenMoCancleLineUp(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("shenmo:处理神魔战场取消排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenmo:处理神魔战场取消排队,完成")
	return nil

}

//神魔战场取消排队
func shenMoCancleLineUp(pl *player.Player) (err error) {
	flag := shenmo.GetShenMoService().CancleLineUp(pl.GetId())
	if !flag {
		return
	}
	isShenMoCancleUp := pbutil.BuildISShenMoCancleUp()
	pl.SendMsg(isShenMoCancleUp)
	return
}
