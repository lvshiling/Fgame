package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_SHENGWEI_DROP_TYPE), dispatch.HandlerFunc(handleShengWeiKilledDrop))
}

func handleShengWeiKilledDrop(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi:被杀声威掉落")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siMsg := msg.(*crosspb.SIShengWeiDrop)
	itemId := siMsg.GetItemId()
	itemNum := siMsg.GetItemNum()
	attackId := siMsg.GetAttackerId()
	err = shengWeiKillerDrop(tpl, itemId, itemNum, attackId)
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
		}).Debug("jieyi:处理跨服声威掉落,完成")

	return nil
}

//跨服声威掉落
func shengWeiKillerDrop(pl *player.Player, itemId int32, itemNum int64, attackId int64) (err error) {
	if itemId == 0 || itemNum == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"itemNum":  itemNum,
				"attackId": attackId,
			}).Warn("jieyi:处理跨服声威掉落,失败")
		return
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(attackId)
	ownId := int64(0)
	if spl != nil {
		ownId = spl.GetId()
	}

	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"itemNum":  itemNum,
				"attackId": attackId,
			}).Warn("jieyi:处理跨服声威掉落,失败")
		return
	}

	minStack := int(jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate().DropMinStack)
	maxStack := int(jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate().DropMaxStack) + 1
	protectedTime := int32(jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate().DropProtectedTime)
	existTime := int32(jieyitemplate.GetJieYiTemplateService().GetJieYiConstantTemplate().DropFailTime)
	stack := int32(mathutils.RandomRange(minStack, maxStack))
	scenelogic.CustomItemDrop(s, pl.GetPosition(), ownId, itemId, int32(itemNum), stack, protectedTime, existTime)
	return
}
