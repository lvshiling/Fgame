package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/jieyi/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_SHENGWEI_DROP_TYPE), dispatch.HandlerFunc(handleShengWeiDrop))
}

//处理跨服声威掉落
func handleShengWeiDrop(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi:处理跨服声威掉落")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isMsg := msg.(*crosspb.ISShengWeiDrop)
	attackId := isMsg.GetAttackerId()
	attackName := isMsg.GetAttackerName()
	err = shengWeiDrop(tpl, attackId, attackName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("jieyi:处理跨服声威掉落,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("jieyi:处理跨服服声威掉落,完成")
	return nil

}

//跨服声威掉落
func shengWeiDrop(pl player.Player, attackId int64, attackName string) (err error) {
	//直接开始计算掉落
	itemId, dropNum := jieyilogic.DropShengWeiZhi(pl, attackId, attackName)

	if dropNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"dropNum":  dropNum,
			}).Warn("jieyi:处理获取声威掉落,掉落数量错误")
		return
	}

	siMsg := pbutil.BuildSIShengWeiDrop(itemId, int64(dropNum), attackId)
	pl.SendCrossMsg(siMsg)
	return
}
