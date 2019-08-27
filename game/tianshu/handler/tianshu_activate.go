package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tianshu/pbutil"
	playertianshu "fgame/fgame/game/tianshu/player"
	tianshutemplate "fgame/fgame/game/tianshu/template"
	tianshutypes "fgame/fgame/game/tianshu/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TIANSHU_ACTIVATE_TYPE), dispatch.HandlerFunc(handleTianShuActivate))
}

//处理天书激活
func handleTianShuActivate(s session.Session, msg interface{}) (err error) {
	log.Debug("tianshu:处理激活天书消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTianShuActivate)
	typ := csMsg.GetType()

	tianshuType := tianshutypes.TianShuType(typ)
	if !tianshuType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Warn("tianshu:处理激活天书消息,错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = tianshuActivate(tpl, tianshuType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("tianshu:处理激活天书消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tianshu:处理激活天书消息完成")
	return nil
}

const (
	initLevel = int32(1)
)

//激活天书信息
func tianshuActivate(pl player.Player, typ tianshutypes.TianShuType) (err error) {
	tianshuManager := pl.GetPlayerDataManager(playertypes.PlayerTianShuDataManagerType).(*playertianshu.PlayerTianShuDataManager)
	if tianshuManager.IsActivateTianShu(typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("tianshu:处理激活天书，已经激活天书")
		playerlogic.SendSystemMessage(pl, lang.TianShuHadActivate)
		return
	}

	tianshuTemp := tianshutemplate.GetTianShuTemplateService().GetTianShuTemplate(typ, initLevel)
	if tianshuTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"level":    initLevel,
			}).Warn("tianshu:处理激活天书，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	activateNeed := int64(tianshuTemp.NeedGold)
	countChargeNum := pl.GetChargeGoldNum()
	if countChargeNum < activateNeed {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"activateNeed":   activateNeed,
				"countChargeNum": countChargeNum,
			}).Warn("tianshu:处理激活天书，不满足激活条件")
		playerlogic.SendSystemMessage(pl, lang.TianShuNotEnoughCondition)
		return
	}

	//消耗物品
	needItemMap := tianshuTemp.GetNeedItemMap()
	if len(needItemMap) > 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		if !inventoryManager.HasEnoughItems(needItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"needItemMap": needItemMap,
				}).Warn("tianshu:处理激活天书，物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		itemUseReason := commonlog.InventoryLogReasonTianShuUplevel
		reasonText := fmt.Sprintf(itemUseReason.String(), typ.String())
		flag := inventoryManager.BatchRemove(needItemMap, itemUseReason, reasonText)
		if !flag {
			panic("tianshu:批量移除物品应该成功")
		}

		inventorylogic.SnapInventoryChanged(pl)
	}

	tianshuManager.ActivateTianShu(typ)

	scMsg := pbutil.BuildSCTianShuActivate(typ)
	pl.SendMsg(scMsg)
	return
}
