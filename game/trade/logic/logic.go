package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playertrade "fgame/fgame/game/trade/player"
	tradetemplate "fgame/fgame/game/trade/template"
	"fgame/fgame/game/trade/trade"
	playervip "fgame/fgame/game/vip/player"
	viptemplate "fgame/fgame/game/vip/template"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

//交易上架
func TradeUploadItem(pl player.Player, typ inventorytypes.BagType, index int32, num int32, gold int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//数量小于1
	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"num":      num,
			}).Warn("trade:处理交易上架，数量不大于0")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	itemObject := inventoryManager.FindItemByIndex(typ, index)
	if itemObject == nil || itemObject.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"num":      num,
			}).Warn("trade:处理交易上架，物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//绑定
	if itemObject.BindType == itemtypes.ItemBindTypeBind {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"num":      num,
			}).Warn("trade:处理交易上架，物品是绑定物品")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}
	//判断过期
	if itemObject.PropertyData.GetExpireType() != inventorytypes.NewItemLimitTimeTypeNone {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"num":      num,
			}).Warn("trade:处理交易上架，有时效物品不能交易")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}
	itemId := itemObject.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"num":      num,
			}).Warn("trade:处理交易上架，物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	//是否可以交易
	if !itemTemplate.CanTrade() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"num":      num,
			}).Warn("trade:处理交易上架，不能交易")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotTrade)
		return
	}
	//判断定价
	if gold < itemTemplate.MarketMinPrice {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"num":      num,
				"gold":     gold,
				"minPrice": itemTemplate.MarketMinPrice,
			}).Warn("trade:处理交易上架，定价太低")
		playerlogic.SendSystemMessage(pl, lang.TradeItemPriceTooLow)
		return
	}

	numOfGlobalTradeItems := trade.GetTradeService().GetNumOfGlobalTradeItems()
	if numOfGlobalTradeItems >= tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate().AllCountMax {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"num":      num,
				"gold":     gold,
				"minPrice": itemTemplate.MarketMinPrice,
			}).Warn("trade:处理交易上架，总上架物品个数太多")
		playerlogic.SendSystemMessage(pl, lang.TradeItemTotalNumLimit)
		return
	}

	propertyData := itemObject.PropertyData
	//复制属性
	copyPropertyData := propertyData.Copy()
	level := itemObject.Level
	//数量
	if itemObject.Num < num {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"num":      num,
			}).Warn("trade:处理交易上架，物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//判断上架上限
	err = trade.GetTradeService().UploadItem(pl, itemId, num, copyPropertyData, level, gold)
	if err != nil {
		return
	}
	reasonText := fmt.Sprintf(commonlog.InventoryLogReasonTradeUpload.String(), gold)
	flag, _ := inventoryManager.RemoveIndex(typ, index, num, commonlog.InventoryLogReasonTradeUpload, reasonText)
	if !flag {
		panic("inventory:移除物品应该是可以的")
	}
	inventorylogic.SnapInventoryChanged(pl)
	return
}

//交易下架
func TradeWithdraw(pl player.Player, tradeId int64) (err error) {
	err = trade.GetTradeService().WithdrawItem(pl, tradeId)
	if err != nil {
		return
	}
	return
}

//交易
func TradeItem(pl player.Player, tradeId int64) (err error) {
	tradeItem := trade.GetTradeService().GetGlobalTradeItem(tradeId)
	if tradeItem == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"tradeId":  tradeId,
		}).Warn("trade:物品不存在,无法交易")
		playerlogic.SendSystemMessage(pl, lang.TradeItemNoExist)
		return
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyData := tradeItem.GetPropertyData()
	if !inventoryManager.HasEnoughSlotItemLevelWithProperty(tradeItem.GetItemId(), tradeItem.GetItemNum(), tradeItem.GetLevel(), itemtypes.ItemBindTypeUnBind, propertyData.GetExpireType(), propertyData.GetItemGetTime(), propertyData.GetExpireTimestamp()) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"tradeId":  tradeId,
		}).Warn("trade:交易物品,背包格子不够")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	costGold := int64(tradeItem.GetGold())
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughGold(costGold, false)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"tradeId":  tradeId,
		}).Warn("trade:元宝不足,无法交易")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}
	_, err = trade.GetTradeService().TradeItem(pl, tradeId)
	if err != nil {
		return
	}
	goldLog := commonlog.GoldLogReasonTrade
	reasonText := fmt.Sprintf(goldLog.String(), tradeId)
	//扣钱
	flag = propertyManager.CostGold(costGold, false, goldLog, reasonText)
	if !flag {
		panic(fmt.Errorf("trade:花钱应该成功的"))
	}
	propertylogic.SnapChangedProperty(pl)
	return
}

//交易成功
func OnPlayerTradeItem(pl player.Player, tradeOrderObject *trade.TradeOrderObject) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemId := tradeOrderObject.GetItemId()
	itemNum := tradeOrderObject.GetNum()
	level := tradeOrderObject.GetLevel()
	bindType := itemtypes.ItemBindTypeUnBind
	propertyData := tradeOrderObject.GetPropertyData()
	reason := commonlog.InventoryLogReasonTradeBuy
	reasonText := fmt.Sprintf(reason.String(), tradeOrderObject.GetId(), tradeOrderObject.GetTradeId(), tradeOrderObject.GetGold())
	//格子足够
	if inventoryManager.HasEnoughSlotItemLevelWithProperty(itemId, itemNum, level, bindType, propertyData.GetExpireType(), propertyData.GetItemGetTime(), propertyData.GetExpireTimestamp()) {
		inventoryManager.AddItemLevelWithPropertyData(itemId, itemNum, level, bindType, propertyData, reason, reasonText)
		inventorylogic.SnapInventoryChanged(pl)
	} else {
		refundTitle := lang.GetLangService().ReadLang(lang.TradeItemTitle)
		refundContent := lang.GetLangService().ReadLang(lang.TradeItemContent)
		now := global.GetGame().GetTimeService().Now()
		dropItemData := inventorylogic.ConverToItemData(tradeOrderObject.GetItemId(), tradeOrderObject.GetNum(), level, itemtypes.ItemBindTypeUnBind, tradeOrderObject.GetPropertyData())
		emaillogic.AddEmailItemLevel(pl, refundTitle, refundContent, now, []*droptemplate.DropItemData{dropItemData})
	}
	tradeManager := pl.GetPlayerDataManager(playertypes.PlayerTradeDataManagerType).(*playertrade.PlayerTradeManager)
	tradeManager.AddBuyLog(
		tradeOrderObject.GetTradeId(),
		tradeOrderObject.GetSellServerId(),
		tradeOrderObject.GetSellPlayerId(),
		tradeOrderObject.GetSellPlayerName(),
		tradeOrderObject.GetServerId(),
		tradeOrderObject.GetPlayerId(),
		pl.GetName(),
		tradeOrderObject.GetGold(),
		tradeOrderObject.GetGold(),
		0,
		tradeOrderObject.GetItemId(),
		tradeOrderObject.GetNum(),
		tradeOrderObject.GetPropertyData().Copy(),
		tradeOrderObject.GetLevel(),
		tradeOrderObject.GetTradeTime(),
	)
	trade.GetTradeService().EndTradeItem(pl, tradeOrderObject)

}

func OnPlayerSellItem(pl player.Player, tradeItemObject *trade.TradeItemObject) {
	sellTitle := lang.GetLangService().ReadLang(lang.TradeSellTitle)
	buyPlayerName := tradeItemObject.GetBuyPlayerName()
	itemTemplate := item.GetItemService().GetItem(int(tradeItemObject.GetItemId()))
	if itemTemplate == nil {
		return
	}
	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)

	vipLevel, star := vipManager.GetVipLevel()
	itemNum := tradeItemObject.GetNum()
	gold := tradeItemObject.GetGold()
	fee := int32(0)
	vipTemplate := viptemplate.GetVipTemplateService().GetVipTemplate(vipLevel, star)
	if vipTemplate != nil {
		fee = vipTemplate.Shouxu
	}
	feeGold := int32(math.Ceil(float64(gold) * float64(fee) / float64(common.MAX_RATE)))
	getGold := gold - feeGold
	buyPlayerNameStr := chatlogic.FormatMailKeyWordNoticeStr(coreutils.FormatColor(chattypes.ColorTypePlayerName, buyPlayerName))
	itemNameStr := chatlogic.FormatMailKeyWordNoticeStr(coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), itemTemplate.Name))
	itemNumStr := chatlogic.FormatMailKeyWordNoticeStr(coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, fmt.Sprintf("%d", itemNum)))
	goldStr := coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, fmt.Sprintf("%d", gold))
	vipLevelStr := coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, fmt.Sprintf("%d", vipLevel))
	feeStr := coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, fmt.Sprintf("%d", fee/100))
	feeGoldStr := coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, fmt.Sprintf("%d", feeGold))
	sellContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.TradeSellContent), buyPlayerNameStr, itemNameStr, itemNumStr, goldStr, vipLevelStr, feeStr, feeGoldStr)
	sellAttach := make(map[int32]int32)
	if getGold > 0 {
		sellAttach[constanttypes.GoldItem] = int32(getGold)
	}
	emaillogic.AddEmail(pl, sellTitle, sellContent, sellAttach)
	tradeManager := pl.GetPlayerDataManager(playertypes.PlayerTradeDataManagerType).(*playertrade.PlayerTradeManager)
	tradeManager.AddSellLog(
		tradeItemObject.GetGlobalTradeId(),
		tradeItemObject.GetServerId(),
		tradeItemObject.GetPlayerId(),
		tradeItemObject.GetPlayerName(),
		tradeItemObject.GetBuyServerId(),
		tradeItemObject.GetBuyPlayerId(),
		tradeItemObject.GetBuyPlayerName(),
		tradeItemObject.GetGold(),
		getGold,
		feeGold,
		tradeItemObject.GetItemId(),
		tradeItemObject.GetNum(),
		tradeItemObject.GetPropertyData().Copy(),
		tradeItemObject.GetLevel(),
		tradeItemObject.GetTradeTime(),
		fee,
	)
	trade.GetTradeService().EndSellItem(pl, tradeItemObject)

}

func OnPlayerTradeRecycle(pl player.Player) {
	tradeConstantTemplate := tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate()
	//小于需要的充值数
	if pl.GetChargeGoldNum() > int64(tradeConstantTemplate.NeedChongzhi) {
		return
	}
	tradeManager := pl.GetPlayerDataManager(playertypes.PlayerTradeDataManagerType).(*playertrade.PlayerTradeManager)
	canRecycleTradeList := trade.GetTradeService().GetCanRecycleTradeList(pl)
	for _, recycleTradeObj := range canRecycleTradeList {
		flag := tradeManager.IfCanRecycle(int64(recycleTradeObj.GetGold()))
		if !flag {
			continue
		}
		//判断个人是否已经超过
		flag = trade.GetTradeService().SystemRecycleTrade(recycleTradeObj.GetGlobalTradeId())
		if !flag {
			continue
		}
		flag = tradeManager.AddRecycleGold(int64(recycleTradeObj.GetGold()))
		if !flag {
			panic(fmt.Errorf("trade:交易回收应该成功"))
		}
		break
	}
}
