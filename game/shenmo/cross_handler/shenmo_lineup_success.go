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
	shenmologic "fgame/fgame/game/shenmo/logic"
	"fgame/fgame/game/shenmo/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_SHENMO_LINEUP_SUCCESS_TYPE), dispatch.HandlerFunc(handleShenMoLineUpSuccess))
}

//处理神魔战场排队成功
func handleShenMoLineUpSuccess(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理神魔战场排队成功")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = shenMoLineUpSuccess(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
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
func shenMoLineUpSuccess(pl player.Player) (err error) {
	pl.ShenMoLineUp(false)
	shenmologic.ShenMoLineUpSuccess(pl)
	scShenMoLineUpSuccess := pbutil.BuildSCShenMoLineUpSuccess()
	pl.SendMsg(scShenMoLineUpSuccess)
	//进入跨服神魔战场
	crosslogic.CrossPlayerDataLogin(pl)
	return
}
