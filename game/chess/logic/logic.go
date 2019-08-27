package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/chess/chess"
	"fgame/fgame/game/chess/pbutil"
	playerchess "fgame/fgame/game/chess/player"
	chesstemplate "fgame/fgame/game/chess/template"
	chesstypes "fgame/fgame/game/chess/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//破解苍龙棋局逻辑
func ChessAttend(pl player.Player, typ chesstypes.ChessType, logTime int64, autoFlag bool) (err error) {

	chessManager := pl.GetPlayerDataManager(playertypes.PlayerChessDataManagerType).(*playerchess.PlayerChessDataManager)
	curChessId := chessManager.GetChessId(typ)
	chessTemplate := chesstemplate.GetChessTemplateService().GetChessByTypAndChessId(typ, curChessId)
	if chessTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"curChessId": curChessId,
				"typ":        typ.String(),
			}).Warn("chess:破解棋局错误,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//次数限制
	flag := chessManager.IsEnoughTimes(typ, 1)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("chess:破解棋局错误，破解次数不足")
		playerlogic.SendSystemMessage(pl, lang.ChessNotEnougTimes)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needGold := int64(chessTemplate.GoldUse)
	needBindGold := int64(chessTemplate.BindGoldUse)
	needSilver := int64(chessTemplate.SilverUse)
	needItemId := chessTemplate.UseItemId
	needItemCount := chessTemplate.UseItemCount
	giftItemId := chessTemplate.GiftItem
	giftItemNum := chessTemplate.GiftItemCount

	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	//物品是否足够
	totalNum := inventoryManager.NumOfItems(int32(needItemId))
	if totalNum < needItemCount {
		if !autoFlag {
			log.WithFields(
				log.Fields{
					"playerId":      pl.GetId(),
					"typ":           typ,
					"needItemId":    needItemId,
					"needItemCount": needItemCount,
				}).Warn("chess:破解棋局错误，道具不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动购买
		needBuyNum := needItemCount - totalNum
		needItemCount = totalNum
		//获取价格
		// shopTemplate := shop.GetShopService().GetShopTemplateByItem(needItemId)
		// if shopTemplate == nil {
		// 	log.WithFields(log.Fields{
		// 		"playerId": pl.GetId(),
		// 		"autoFlag": autoFlag,
		// 	}).Warn("chess:商铺没有该道具,无法自动购买")
		// 	playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
		// 	return
		// }
		// shopNeedGold, shopNeedBindGold, shopNeedSilver := shopTemplate.GetConsumeData(needBuyNum)
		// needSilver += shopNeedSilver
		// needGold += int64(shopNeedGold)
		// needBindGold += int64(shopNeedBindGold)
		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(needItemId) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("chess:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, needItemId, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("chess:购买物品失败,棋局破解失败")
				playerlogic.SendSystemMessage(pl, lang.ShopChessAutoBuyItemFail)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			needGold += shopNeedGold
			needBindGold += shopNeedBindGold
			needSilver += shopNeedSilver
		}
	}

	//是否足够银两
	flag = propertyManager.HasEnoughSilver(needSilver)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("chess:破解棋局错误，银两不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//是否足够元宝
	flag = propertyManager.HasEnoughGold(needGold, false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("chess:破解棋局错误，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}
	//是否足够绑元
	needCostBindGold := needBindGold + needGold
	flag = propertyManager.HasEnoughGold(needCostBindGold, true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("chess:破解棋局错误，绑元不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//破解棋局
	rewList := chessManager.GetChessDrop(typ, 1)
	if len(rewList) == 0 {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"chessType":    typ,
				"attendTimes ": 1,
			}).Warn("chess:破解棋局错误，掉落为空")
		playerlogic.SendSystemMessage(pl, lang.ChessNotGetRewards)
		return
	}

	logRewList := make([]*droptemplate.DropItemData, len(rewList))
	copy(logRewList, rewList)
	if giftItemId > 0 {
		giftData := droptemplate.CreateItemData(giftItemId, giftItemNum, 0, itemtypes.ItemBindTypeUnBind)
		rewList = append(rewList, giftData)
	}

	//背包空间
	var rewItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(rewList) > 0 {
		rewItemList, resMap = droplogic.SeperateItemDatas(rewList)
	}

	if !inventoryManager.HasEnoughSlotsOfItemLevel(rewItemList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("chess:破解棋局错误,背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗银两
	if needSilver > 0 {
		silverUseReason := commonlog.SilverLogReasonChessUse
		silverUseReasonText := fmt.Sprintf(silverUseReason.String(), typ.String())
		flag := propertyManager.CostSilver(needSilver, silverUseReason, silverUseReasonText)
		if !flag {
			panic("chess：消耗银两应该成功")
		}
	}

	//消耗元宝
	if needGold > 0 {
		goldUseReason := commonlog.GoldLogReasonChessUse
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), typ.String())
		flag := propertyManager.CostGold(needGold, false, goldUseReason, goldUseReasonText)
		if !flag {
			panic("chess:消耗元宝应该成功")
		}
	}
	//消耗绑元
	if needBindGold > 0 {
		goldUseReason := commonlog.GoldLogReasonChessUse
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), typ.String())
		flag := propertyManager.CostGold(needBindGold, true, goldUseReason, goldUseReasonText)
		if !flag {
			panic("chess:消耗元宝应该成功")
		}
	}

	//消耗物品
	if needItemCount > 0 {
		itemUseReason := commonlog.InventoryLogReasonChessAttend
		itemUseReasonText := fmt.Sprintf(itemUseReason.String(), typ)
		if flag := inventoryManager.UseItem(needItemId, needItemCount, itemUseReason, itemUseReasonText); !flag {
			panic("chess: attend chess use item should be ok")
		}
	}

	//增加掉落
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonChessGet
		silverReason := commonlog.SilverLogReasonChessGet
		levelReason := commonlog.LevelLogReasonChessGet
		goldReasonText := fmt.Sprintf(goldReason.String(), typ.String())
		silverReasonText := fmt.Sprintf(silverReason.String(), typ.String())

		levelReasonText := fmt.Sprintf(levelReason.String(), typ.String())
		err = droplogic.AddRes(pl, resMap, goldReason, goldReasonText, silverReason, silverReasonText, levelReason, levelReasonText)
		if err != nil {
			return
		}
	}

	if len(rewItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonChessGet
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ.String())
		flag = inventoryManager.BatchAddOfItemLevel(rewItemList, itemGetReason, itemGetReasonText)
		if !flag {
			panic("chess:增加物品应该成功")
		}
	}

	for _, itemData := range logRewList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()

		//更新棋局
		chessManager.AttendChess(typ)
		//生成日志
		chess.GetChessService().AddLog(pl.GetName(), itemId, num)

		//稀有道具公告
		inventorylogic.PrecioustemBroadcast(pl, itemId, num, lang.InventoryChessItemNotice)
	}

	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	logList := chess.GetChessService().GetLogByTime(logTime)
	scChessAttend := pbutil.BuildSCChessAttend(rewList, typ, logList, autoFlag)
	pl.SendMsg(scChessAttend)
	return
}
