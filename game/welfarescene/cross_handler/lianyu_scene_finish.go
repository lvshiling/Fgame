package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	lianyulogic "fgame/fgame/game/lianyu/logic"
	"fgame/fgame/game/lianyu/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_LIANYU_FINISH_LINEUP_CANCLE_TYPE), dispatch.HandlerFunc(handleLianYuFinishLineUpCancle))
}

//处理无间炼狱结束通知排队人员
func handleLianYuFinishLineUpCancle(s session.Session, msg interface{}) (err error) {
	log.Debug("lianyu:处理无间炼狱结束通知排队人员")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = lianYuFinishLineUpCancle(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("lianyu:处理无间炼狱结束通知排队人员,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lianyu:处理无间炼狱结束通知排队人员,完成")
	return nil

}

//无间炼狱结束通知排队人员
func lianYuFinishLineUpCancle(pl player.Player) (err error) {
	pl.LianYuLineUp(false)
	lianyulogic.LianYuFinishLineUpCancle(pl)
	scLianYuFinishToLineUp := pbutil.BuildSCLianYuFinishToLineUp()
	pl.SendMsg(scLianYuFinishToLineUp)
	return
}
