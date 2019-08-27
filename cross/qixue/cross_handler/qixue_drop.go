package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	qixuetemplate "fgame/fgame/game/qixue/template"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_QIXUE_DROP_TYPE), dispatch.HandlerFunc(handleQiXueKilledDrop))
}

func handleQiXueKilledDrop(s session.Session, msg interface{}) (err error) {
	log.Debug("qixue:泣血枪 被杀掉落")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siMsg := msg.(*crosspb.SIQiXueDrop)
	itemId := siMsg.GetItemId()
	itemNum := siMsg.GetItemNum()
	attackId := siMsg.GetAttackerId()
	err = qiXueKillerDrop(tpl, itemId, itemNum, attackId)
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
func qiXueKillerDrop(pl *player.Player, itemId int32, itemNum int64, attackId int64) (err error) {
	if itemId == 0 || itemNum == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"itemNum":  itemNum,
				"attackId": attackId,
			}).Warn("qixue:处理跨服泣血枪掉落,失败")
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
			}).Warn("qixue:处理跨服泣血枪掉落,失败")
		return
	}

	minStack := int(qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropMinStack)
	maxStack := int(qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropMaxStack) + 1
	protectedTime := qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropProtectedTime
	existTime := qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropFailTime
	stack := int32(mathutils.RandomRange(minStack, maxStack))
	scenelogic.CustomItemDrop(s, pl.GetPosition(), ownId, itemId, int32(itemNum), stack, protectedTime, existTime)
	return
}
