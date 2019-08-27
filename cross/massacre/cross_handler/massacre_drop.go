package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_MASSACRE_DROP_TYPE), dispatch.HandlerFunc(handleMassacreKilledDrop))
}

func handleMassacreKilledDrop(s session.Session, msg interface{}) (err error) {
	log.Debug("massacre:戮仙刃 被杀掉落")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siMassacreDrop := msg.(*crosspb.SIMassacreDrop)
	itemId := siMassacreDrop.GetItemId()
	itemNum := siMassacreDrop.GetItemNum()
	attackId := siMassacreDrop.GetAttackerId()
	err = massacreKillerDrop(tpl, itemId, itemNum, attackId)
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
func massacreKillerDrop(pl *player.Player, itemId int32, itemNum int64, attackId int64) (err error) {
	if itemId == 0 || itemNum == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"itemNum":  itemNum,
				"attackId": attackId,
			}).Warn("massacre:处理跨服戮仙刃掉落,失败")
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
			}).Warn("massacre:处理跨服戮仙刃掉落,失败")
		return
	}

	minStack := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqMinStack))
	maxStack := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqMaxStack)) + 1
	protectedTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqProtectedTime)
	existTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqExistTime)
	stack := int32(mathutils.RandomRange(minStack, maxStack))
	scenelogic.CustomItemDrop(s, pl.GetPosition(), ownId, itemId, int32(itemNum), stack, protectedTime, existTime)

	return
}
