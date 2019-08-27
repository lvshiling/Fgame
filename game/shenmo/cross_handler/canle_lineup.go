package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosslogic "fgame/fgame/game/cross/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_SHENMO_CANCLE_LINEUP_TYPE), dispatch.HandlerFunc(handleShenMoCancleLineUp))
}

//处理神魔战场取消排队
func handleShenMoCancleLineUp(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理神魔战场取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = shenMoCancleLineUp(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
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
func shenMoCancleLineUp(pl player.Player) (err error) {
	pl.ShenMoLineUp(false)
	//退出跨服
	crosslogic.PlayerExitCross(pl)
	return
}
