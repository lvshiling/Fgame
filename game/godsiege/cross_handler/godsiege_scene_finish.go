package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	godsiegelogic "fgame/fgame/game/godsiege/logic"
	"fgame/fgame/game/godsiege/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_FINISH_LINEUP_CANCLE_TYPE), dispatch.HandlerFunc(handleGodSiegeFinishLineUpCancle))
}

//处理神兽攻城结束通知排队人员
func handleGodSiegeFinishLineUpCancle(s session.Session, msg interface{}) (err error) {
	log.Debug("godsiege:处理神兽攻城结束通知排队人员")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isGodSiegeFinishLineUpCancle := msg.(*crosspb.ISGodSiegeFinishLineUpCancle)
	godType := isGodSiegeFinishLineUpCancle.GetGodType()
	err = godSiegeFinishLineUpCancle(tpl, godType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"godType":  godType,
				"err":      err,
			}).Error("godsiege:处理神兽攻城结束通知排队人员,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("godsiege:处理神兽攻城结束通知排队人员,完成")
	return nil

}

//神兽攻城结束通知排队人员
func godSiegeFinishLineUpCancle(pl player.Player, godType int32) (err error) {
	pl.GodSiegeCancleLineUp()
	godsiegelogic.GodSiegeFinishLineUpCancle(pl, godType)
	scGodSiegeFinishToLineUp := pbutil.BuildSCGodSiegeFinishToLineUp(godType)
	pl.SendMsg(scGodSiegeFinishToLineUp)
	return
}
