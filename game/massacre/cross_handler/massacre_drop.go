package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	massacrelogic "fgame/fgame/game/massacre/logic"
	"fgame/fgame/game/massacre/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_MASSACRE_DROP_TYPE), dispatch.HandlerFunc(handleMassacreDrop))
}

//处理跨服戮仙刃掉落
func handleMassacreDrop(s session.Session, msg interface{}) (err error) {
	log.Debug("massacre:处理跨服戮仙刃掉落")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isMassacreDrop := msg.(*crosspb.ISMassacreDrop)
	attackId := isMassacreDrop.GetAttackerId()
	attackName := isMassacreDrop.GetAttackerName()
	err = massacreDrop(tpl, attackId, attackName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("massacre:处理跨服戮仙刃掉落,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("massacre:处理跨服服戮仙刃掉落,完成")
	return nil

}

//跨服戮仙刃掉落
func massacreDrop(pl player.Player, attackId int64, attackName string) (err error) {
	//直接开始计算掉落吧
	itemId, dropNum := massacrelogic.MassacreProcessDrop(pl, attackId, attackName)

	if dropNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("massacre:处理获取戮仙刃掉落,掉落数量错误2")
		return
	}

	siMassacreDrop := pbutil.BuildSIMassacreDrop(itemId, dropNum, attackId)
	pl.SendCrossMsg(siMassacreDrop)
	return
}
