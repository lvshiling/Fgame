package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shenmologic "fgame/fgame/game/shenmo/logic"
	"fgame/fgame/game/shenmo/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_SHENMO_FINISH_LINEUP_CANCLE_TYPE), dispatch.HandlerFunc(handleShenMoFinishLineUpCancle))
}

//处理神魔战场结束通知排队人员
func handleShenMoFinishLineUpCancle(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理神魔战场结束通知排队人员")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = shenMoFinishLineUpCancle(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("shenmo:处理神魔战场结束通知排队人员,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenmo:处理神魔战场结束通知排队人员,完成")
	return nil

}

//神魔战场结束通知排队人员
func shenMoFinishLineUpCancle(pl player.Player) (err error) {
	pl.ShenMoLineUp(false)
	shenmologic.ShenMoFinishLineUpCancle(pl)
	scShenMoFinishToLineUp := pbutil.BuildSCShenMoFinishToLineUp()
	pl.SendMsg(scShenMoFinishToLineUp)
	return
}
