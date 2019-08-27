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
	"fgame/fgame/game/shenmo/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_SHENMO_ATTEND_TYPE), dispatch.HandlerFunc(handleShenMoAttend))
}

//处理跨服参加神魔战场
func handleShenMoAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理跨服参加神魔战场")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isShenMoAttend := msg.(*crosspb.ISShenMoAttend)
	isLineUp := isShenMoAttend.GetIsLineUp()
	beforeNum := isShenMoAttend.GetBeforeNum()
	err = shenMoAttend(tpl, isLineUp, beforeNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"isLineUp":  isLineUp,
				"beforeNum": beforeNum,
				"err":       err,
			}).Error("shenmo:处理跨服参加神魔战场,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"beforeNum": beforeNum,
		}).Debug("shenmo:处理跨服参加神魔战场,完成")
	return nil

}

//参加神魔战场
func shenMoAttend(pl player.Player, isLineUp bool, beforeNum int32) (err error) {
	if !isLineUp {
		//进入跨服神魔战场
		crosslogic.CrossPlayerDataLogin(pl)
	} else {
		pl.ShenMoLineUp(isLineUp)
		scShenMoLineUp := pbutil.BuildSCShenMoLineUp(beforeNum)
		pl.SendMsg(scShenMoLineUp)
	}
	return
}
