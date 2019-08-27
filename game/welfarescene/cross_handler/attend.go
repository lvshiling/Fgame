package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosslogic "fgame/fgame/game/cross/logic"
	"fgame/fgame/game/lianyu/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_LIANYU_ATTEND_TYPE), dispatch.HandlerFunc(handleLianYuAttend))
}

//处理跨服参加无间炼狱
func handleLianYuAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("lianyu:处理跨服参加无间炼狱")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isLianYuAttend := msg.(*crosspb.ISLianYuAttend)
	isLineUp := isLianYuAttend.GetIsLineUp()
	beforeNum := isLianYuAttend.GetBeforeNum()
	err = lianYuAttend(tpl, isLineUp, beforeNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"isLineUp":  isLineUp,
				"beforeNum": beforeNum,
				"err":       err,
			}).Error("lianyu:处理跨服参加无间炼狱,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"beforeNum": beforeNum,
		}).Debug("lianyu:处理跨服参加无间炼狱,完成")
	return nil

}

//参加无间炼狱
func lianYuAttend(pl player.Player, isLineUp bool, beforeNum int32) (err error) {
	if !isLineUp {
		//进入跨服无间炼狱
		crosslogic.CrossPlayerDataLogin(pl)
	} else {
		pl.LianYuLineUp(isLineUp)
		scLianYuLineUp := pbutil.BuildSCLianYuLineUp(beforeNum)
		pl.SendMsg(scLianYuLineUp)
	}
	return
}
