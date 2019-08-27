package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/fireworks/pbutil"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/shop/shop"
	"fmt"

	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	noticelogic "fgame/fgame/game/notice/logic"
	playerlogic "fgame/fgame/game/player/logic"
	playerproperty "fgame/fgame/game/property/player"

	inventorylogic "fgame/fgame/game/inventory/logic"
	propertylogic "fgame/fgame/game/property/logic"
	shoplogic "fgame/fgame/game/shop/logic"

	log "github.com/Sirupsen/logrus"
)

//广播烟花消息
func BroadcastMsg(pl player.Player, itemId int32, num int32) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	scFireworksBroadcast := pbutil.BuildSCFireworksBroadcast(itemId, num)
	s.BroadcastMsg(scFireworksBroadcast)
}

func ShootFireworks(pl player.Player, itemId int32, num int32, isNotice bool) (isReturn bool) {
	isReturn = true
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
			"num":      num,
		}).Warn("fireworks:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return
	}
	itemType := itemTemplate.GetItemType()
	itemSubType := itemTemplate.GetItemSubType()

	if itemType != itemtypes.ItemTypeAutoUseRes {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
			"num":      num,
		}).Warn("fireworks:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if itemSubType != itemtypes.ItemAutoUseResSubTypeNormalFireworks &&
		itemSubType != itemtypes.ItemAutoUseResSubTypeSeniorFireworks {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
			"num":      num,
		}).Warn("fireworks:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	costGold := int32(0)
	costBindGold := int32(0)
	costSilver := int64(0)
	totalNum := int32(0)
	itemCount := num

	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	totalNum = inventoryManager.NumOfItems(itemId)
	if totalNum < itemCount {
		//自动购买
		needBuyNum := itemCount - totalNum
		itemCount = totalNum
		//获取价格
		// shopTemplate := shop.GetShopService().GetShopTemplateByItem(itemId)
		// if shopTemplate == nil {
		// 	log.WithFields(log.Fields{
		// 		"playerId": pl.GetId(),
		// 		"itemId":   itemId,
		// 		"num":      num,
		// 	}).Warn("fireworks:商铺没有该道具,无法自动购买")
		// 	playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
		// 	return
		// }

		// shopNeedGold, shopNeedBindGold, shopNeedSilver := shopTemplate.GetConsumeData(needBuyNum)
		// costGold += shopNeedGold
		// costBindGold += shopNeedBindGold
		// costSilver += shopNeedSilver
		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(itemId) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"num":      num,
				}).Warn("fireworks:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, itemId, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"num":      num,
				}).Warn("fireworks:购买烟花物品失败")
				playerlogic.SendSystemMessage(pl, lang.ShopFireworksAutoBuyItemFail)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			costGold += int32(shopNeedGold)
			costBindGold += int32(shopNeedBindGold)
			costSilver += shopNeedSilver
		}
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("fireworks:银两不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//元宝是否足够
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("fireworks:元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//绑元是否足够
	needCostBindGold := costGold + costBindGold
	if needCostBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needCostBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("fireworks:元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	goldReason := commonlog.GoldLogReasonFireworks
	silverReason := commonlog.SilverLogReasonFireworks
	goldReasonText := fmt.Sprintf(goldReason.String(), itemId, num)
	silverReasonText := fmt.Sprintf(silverReason.String(), itemId, num)
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), goldReason, goldReasonText, int64(costSilver), silverReason, silverReasonText)
	if !flag {
		panic(fmt.Errorf("fireworks: ShootFireworks CostGold should be ok"))
	}
	//同步银两
	propertylogic.SnapChangedProperty(pl)

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonShootFireworks
		reasonText := fmt.Sprintf(inventoryReason.String(), num)
		flag := inventoryManager.UseItem(itemId, itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("fireworks: ShootFireworks use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	isReturn = false
	if isNotice {
		name := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		itemName := coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemplate.FormateItemNameOfNum(num)))
		linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
		itemNameLink := coreutils.FormatLink(itemName, linkArgs)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.FireworksNotice), name, itemNameLink)
		//跑马灯
		noticelogic.NoticeNumBroadcastScene(pl.GetScene(), []byte(content), 0, int32(1))
		//系统频道
		chatlogic.BroadcastScene(pl.GetScene(), chattypes.MsgTypeText, []byte(content))
		BroadcastMsg(pl, itemId, num)
	}
	return
}
