package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	qixuelogic "fgame/fgame/game/qixue/logic"
	"fgame/fgame/game/qixue/pbutil"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_QIXUE_DROP_TYPE), dispatch.HandlerFunc(handleQiXueDrop))
}

//处理跨服泣血枪掉落
func handleQiXueDrop(s session.Session, msg interface{}) (err error) {
	log.Debug("qixue:处理跨服泣血枪掉落")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isMsg := msg.(*crosspb.ISQiXueDrop)
	attackId := isMsg.GetAttackerId()
	attackName := isMsg.GetAttackerName()
	err = qixueDrop(tpl, attackId, attackName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("qixue:处理跨服泣血枪掉落,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("qixue:处理跨服服泣血枪掉落,完成")
	return nil

}

//跨服泣血枪掉落
func qixueDrop(pl player.Player, attackId int64, attackName string) (err error) {
	//直接开始计算掉落吧
	itemId, dropNum := qixuelogic.QiXueProcessDrop(pl, attackId, attackName)

	if dropNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"dropNum":  dropNum,
			}).Warn("qixue:处理获取泣血枪掉落,掉落数量错误")
		return
	}

	siMsg := pbutil.BuildSIQiXueDrop(itemId, dropNum, attackId)
	pl.SendCrossMsg(siMsg)
	return
}
