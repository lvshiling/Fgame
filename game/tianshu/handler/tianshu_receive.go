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
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tianshu/pbutil"
	playertianshu "fgame/fgame/game/tianshu/player"
	tianshutemplate "fgame/fgame/game/tianshu/template"
	tianshutypes "fgame/fgame/game/tianshu/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TIANSHU_GIFT_RECEIVE_TYPE), dispatch.HandlerFunc(handleTianShuReceive))
}

//处理天书领取奖励
func handleTianShuReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("tianshu:处理领取天书奖励消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTianShuGiftReceive)
	typ := csMsg.GetType()

	tianshuType := tianshutypes.TianShuType(typ)
	if !tianshuType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Warn("tianshu:处理领取天书奖励消息,错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = tianshuReceive(tpl, tianshuType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("tianshu:处理领取天书奖励消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tianshu:处理领取天书奖励消息完成")
	return nil
}

//领取天书奖励信息
func tianshuReceive(pl player.Player, typ tianshutypes.TianShuType) (err error) {
	tianshuManager := pl.GetPlayerDataManager(playertypes.PlayerTianShuDataManagerType).(*playertianshu.PlayerTianShuDataManager)
	tianshuManager.RefreshReceive()

	if !tianshuManager.IsActivateTianShu(typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("tianshu:处理领取天书奖励，天书未激活")
		playerlogic.SendSystemMessage(pl, lang.TianShuNotActivate)
		return
	}

	if tianshuManager.IsReceive(typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("tianshu:处理领取天书奖励，天书礼包已领取")
		playerlogic.SendSystemMessage(pl, lang.TianShuHadReceive)
		return
	}

	tianshuLevel := tianshuManager.GetTianShuLevel(typ)
	tianshuTemp := tianshutemplate.GetTianShuTemplateService().GetTianShuTemplate(typ, tianshuLevel)
	if tianshuTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"level":    initLevel,
			}).Warn("tianshu:处理领取天书奖励，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	giftItemMap := tianshuTemp.GetGiftItemMap()
	rewSilver := int64(tianshuTemp.FreeGiftSilver)
	rewBindGold := int64(tianshuTemp.FreeGiftBindgold)
	rewGold := int64(tianshuTemp.FreeGiftGold)

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	if len(giftItemMap) > 0 {
		if !inventoryManager.HasEnoughSlots(giftItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"giftItemMap": giftItemMap,
				}).Warn("tianshu:处理领取天书奖励，背包不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
	}

	tianshuManager.ReceiveTianShuGift(typ)

	if len(giftItemMap) > 0 {
		itemAddReason := commonlog.InventoryLogReasonTianShuRew
		reasonText := fmt.Sprintf(itemAddReason.String(), typ)
		flag := inventoryManager.BatchAdd(giftItemMap, itemAddReason, reasonText)
		if !flag {
			panic("tianshu:批量添加物品应该成功")
		}
	}

	silverAddReason := commonlog.SilverLogReasonTianShuRew
	silverReasonText := silverAddReason.String()
	goldAddReason := commonlog.GoldLogReasonTianShuActivate
	goldReasonText := fmt.Sprintf(goldAddReason.String(), typ.String())
	flag := propertyManager.AddMoney(rewBindGold, rewGold, goldAddReason, goldReasonText, rewSilver, silverAddReason, silverReasonText)
	if !flag {
		panic("tianshu:添加天书奖励应该成功")
	}

	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCTianShuGiftReceive(typ)
	pl.SendMsg(scMsg)
	return
}
