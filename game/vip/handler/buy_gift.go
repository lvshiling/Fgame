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
	"fgame/fgame/game/vip/pbutil"
	playervip "fgame/fgame/game/vip/player"
	viptemplate "fgame/fgame/game/vip/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_VIP_GIFT_BUY_TYPE), dispatch.HandlerFunc(handleVipGiftBuy))
}

//处理购买vip礼包
func handleVipGiftBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理购买vip礼包消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSVipGiftBuy)
	buyLevel := csMsg.GetGiftLevel()
	star := csMsg.GetGiftStar()

	err = buyVipGift(tpl, buyLevel, star)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("vip:处理购买vip礼包消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("vip:处理购买vip礼包消息完成")
	return nil

}

//购买vip礼包界面逻辑
func buyVipGift(pl player.Player, buyLevel, star int32) (err error) {
	vipTemplate := viptemplate.GetVipTemplateService().GetVipTemplate(buyLevel, star)
	if vipTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"buyLevel": buyLevel,
			}).Warn("vip:处理购买vip礼包消息,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	curLevel, _ := vipManager.GetVipLevel()
	if curLevel < buyLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"curLevel": curLevel,
				"buyLevel": buyLevel,
			}).Warn("vip:处理购买vip礼包消息,vip等级不足")
		playerlogic.SendSystemMessage(pl, lang.VipLevelToLow)
		return
	}

	if !vipManager.IsCanBuyGift(buyLevel) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"buyLevel": buyLevel,
			}).Warn("vip:处理购买vip礼包消息, 不能重复购买")
		playerlogic.SendSystemMessage(pl, lang.VipHadBuyGift)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	needGold := vipTemplate.Price
	if !propertyManager.HasEnoughGold(int64(needGold), false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
				"buyLevel": buyLevel,
			}).Warn("vip:处理购买vip礼包消息, 元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemMap := vipTemplate.GetDisCountItemMap()
	if !inventoryManager.HasEnoughSlots(itemMap) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemMap":  itemMap,
			}).Warn("vip:处理购买vip礼包消息, 背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	goldCostReason := commonlog.GoldLogReasonBuyVipGift
	flag := propertyManager.CostGold(int64(needGold), false, goldCostReason, goldCostReason.String())
	if !flag {
		panic(fmt.Errorf("vip: 消耗元宝应该成功"))
	}

	silverRew := vipTemplate.GiftSilver
	if silverRew > 0 {
		getSilverReason := commonlog.SilverLogReasonVipGiftRew
		propertyManager.AddSilver(silverRew, getSilverReason, getSilverReason.String())
	}

	itemGetReason := commonlog.InventoryLogReasonVipGiftGet
	flag = inventoryManager.BatchAdd(itemMap, itemGetReason, itemGetReason.String())
	if !flag {
		panic(fmt.Errorf("vip: 批量添加物品应该成功"))
	}

	//添加购买记录
	vipManager.BuyVipGift(buyLevel)

	// 同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	showItemMap := vipTemplate.GetEmailDisCountItemMap()
	scMsg := pbutil.BuildSCVipGiftBuy(buyLevel, star, showItemMap)
	pl.SendMsg(scMsg)
	return
}
