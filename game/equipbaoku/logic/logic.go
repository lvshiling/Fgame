package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/common/common"
	commontypes "fgame/fgame/game/common/types"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/equipbaoku/equipbaoku"
	"fgame/fgame/game/equipbaoku/pbutil"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equipbaokutemplate "fgame/fgame/game/equipbaoku/template"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

//探索装备宝库逻辑
func EquipBaoKuAttend(pl player.Player, logTime int64, autoFlag bool, typ equipbaokutypes.BaoKuType) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeEquipBaoKu) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:探索宝库错误,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}
	equipBaoKuManager := pl.GetPlayerDataManager(playertypes.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	equipBaoKuTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetEquipBaoKuByLevAndZhuanNum(pl.GetLevel(), pl.GetZhuanSheng(), typ)
	if equipBaoKuTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:探索宝库错误,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needGold := int64(equipBaoKuTemplate.GoldUse)
	needBindGold := int64(equipBaoKuTemplate.BindGoldUse)
	needSilver := int64(equipBaoKuTemplate.SilverUse)
	needItemId := equipBaoKuTemplate.UseItemId
	needItemCount := equipBaoKuTemplate.UseItemCount

	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	//物品是否足够
	totalNum := inventoryManager.NumOfItems(int32(needItemId))
	if totalNum < needItemCount {
		if !autoFlag {
			log.WithFields(
				log.Fields{
					"playerId":      pl.GetId(),
					"needItemId":    needItemId,
					"needItemCount": needItemCount,
				}).Warn("equipbaoku:探索宝库错误，道具不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动购买
		needBuyNum := needItemCount - totalNum
		needItemCount = totalNum
		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(needItemId) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("equipbaoku:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, needItemId, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("equipbaoku:购买物品失败,宝库探索失败")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNumInvalid)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			needGold += shopNeedGold
			needBindGold += shopNeedBindGold
			needSilver += shopNeedSilver
		}
	}

	//是否足够银两
	flag := propertyManager.HasEnoughSilver(needSilver)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:探索宝库错误，银两不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//是否足够元宝
	flag = propertyManager.HasEnoughGold(needGold, false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:探索宝库错误，元宝不足")
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
			}).Warn("equipbaoku:探索宝库错误，绑元不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//探索宝库
	rewList := equipBaoKuManager.GetEquipBaoKuDrop(1, typ)
	if len(rewList) == 0 {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"attendTimes ": 1,
			}).Warn("equipbaoku:探索宝库错误，掉落为空")
		playerlogic.SendSystemMessage(pl, lang.EquipBaoKuNotGetRewards)
		return
	}

	//背包空间
	var rewItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(rewList) > 0 {
		rewItemList, resMap = droplogic.SeperateItemDatas(rewList)
	}

	if !inventoryManager.HasEnoughSlotsOfItemLevelMiBao(rewItemList, typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:探索宝库错误,秘宝仓库空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryMiBaoDepotSlotNoEnough)
		return
	}

	//消耗银两
	if needSilver > 0 {
		silverUseReason := commonlog.SilverLogReasonEquipBaoKuUse
		silverUseReasonText := fmt.Sprintf(silverUseReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		flag := propertyManager.CostSilver(needSilver, silverUseReason, silverUseReasonText)
		if !flag {
			panic("equipbaoku:消耗银两应该成功")
		}
	}

	//消耗元宝
	if needGold > 0 {
		goldUseReason := commonlog.GoldLogReasonEquipBaoKuUse
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		flag := propertyManager.CostGold(needGold, false, goldUseReason, goldUseReasonText)
		if !flag {
			panic("equipbaoku:消耗元宝应该成功")
		}
	}
	//消耗绑元
	if needBindGold > 0 {
		goldUseReason := commonlog.GoldLogReasonEquipBaoKuUse
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		flag := propertyManager.CostGold(needBindGold, true, goldUseReason, goldUseReasonText)
		if !flag {
			panic("equipbaoku:消耗元宝应该成功")
		}
	}

	//消耗物品
	if needItemCount > 0 {
		itemUseReason := commonlog.InventoryLogReasonEquipBaoKuAttend
		itemUseReasonText := fmt.Sprintf(itemUseReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		if flag := inventoryManager.UseItem(needItemId, needItemCount, itemUseReason, itemUseReasonText); !flag {
			panic("equipbaoku: attend equipbaoku use item should be ok")
		}
	}

	//增加掉落
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonEquipBaoKuGet
		silverReason := commonlog.SilverLogReasonEquipBaoKuGet
		levelReason := commonlog.LevelLogReasonEquipBaoKuGet
		goldReasonText := fmt.Sprintf(goldReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		silverReasonText := fmt.Sprintf(silverReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		levelReasonText := fmt.Sprintf(levelReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		err = droplogic.AddRes(pl, resMap, goldReason, goldReasonText, silverReason, silverReasonText, levelReason, levelReasonText)
		if err != nil {
			return
		}
	}

	if len(rewItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonEquipBaoKuGet
		itemGetReasonText := fmt.Sprintf(itemGetReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		flag = inventoryManager.BatchAddOfItemLevelMiBao(rewItemList, itemGetReason, itemGetReasonText, typ)
		if !flag {
			panic("equipbaoku:增加物品应该成功")
		}
	}

	for _, itemData := range rewList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		//生成日志
		equipbaoku.GetEquipBaoKuService().AddLog(pl.GetName(), itemId, num, typ)
		//稀有道具公告
		if typ == equipbaokutypes.BaoKuTypeEquip {
			inventorylogic.PrecioustemBroadcast(pl, itemId, num, lang.InventoryEquipBaoKuItemNotice)
		} else {
			inventorylogic.PrecioustemBroadcast(pl, itemId, num, lang.InventoryMaterialBaoKuItemNotice)
		}
	}
	//宝库积分,幸运值
	addXingYunZhi := equipBaoKuTemplate.GiftXingYunZhi
	addJiFen := equipBaoKuTemplate.GiftJiFen
	isDouble, luckyPointCritNum, attendPointCritNum := welfarelogic.IsCanDrewBaoKuCrit()
	if isDouble {
		addXingYunZhi = addXingYunZhi + int32(math.Ceil(float64(luckyPointCritNum)/float64(common.MAX_RATE)*float64(addXingYunZhi)))
		addJiFen = addJiFen + int32(math.Ceil(float64(attendPointCritNum)/float64(common.MAX_RATE)*float64(addJiFen)))
	}
	equipBaoKuManager.AttendEquipBaoKu(addXingYunZhi, addJiFen, 1, commontypes.ChangeTypeAttendGet, typ)

	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapMiBaoDepotChanged(pl, typ)
	inventorylogic.SnapInventoryChanged(pl)

	luckyPoints := equipBaoKuManager.GetEquipBaoKuObj(typ).GetLuckyPoints()
	attendPoints := equipBaoKuManager.GetEquipBaoKuObj(typ).GetAttendPoints()
	logList := equipbaoku.GetEquipBaoKuService().GetLogByTime(logTime, typ)
	scEquipBaoKuAttend := pbutil.BuildSCEquipBaoKuAttend(rewList, logList, luckyPoints, attendPoints, autoFlag, int32(typ))
	pl.SendMsg(scEquipBaoKuAttend)
	return
}

//计算分解装备
func CountResolveMiBaoDepotEquip(pl player.Player, itemIndexList []int32) (totalExp int64, returnItemMap map[int32]int32, flag bool) {
	returnItemMap = make(map[int32]int32)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	for _, itemIndex := range itemIndexList {
		it := inventoryManager.FindMiBaoDepotItemByIndex(itemIndex, equipbaokutypes.BaoKuTypeEquip)
		if it == nil {
			log.WithFields(
				log.Fields{
					"index": itemIndex,
				}).Warn("equipbaoku:格子不存在")
			return
		}
		if it.ItemId == 0 {
			log.WithFields(
				log.Fields{
					"index": itemIndex,
				}).Warn("equipbaoku:格子没有物品")
			return
		}

		itemTemp := item.GetItemService().GetItem(int(it.ItemId))
		// 分解金装
		if itemTemp.IsGoldEquip() {
			goldEquipTemp := itemTemp.GetGoldEquipTemplate()

			// 计算装备提供的经验
			totalExp += int64(goldEquipTemp.TunshiExp)

			// 返还物品
			data := it.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
			if data.OpenLightLevel >= 0 {
				openLightTemp := goldEquipTemp.GetOpenLightTemplate(data.OpenLightLevel)
				if openLightTemp == nil {
					goto UpstarLevel
				}
				for itemId, num := range openLightTemp.GetReturnItemMap() {
					_, ok := returnItemMap[itemId]
					if !ok {
						returnItemMap[itemId] = num
					} else {
						returnItemMap[itemId] += num
					}
				}
			}
		UpstarLevel:
			if data.UpstarLevel >= 0 {
				upstarTemp := goldEquipTemp.GetUpstarTemplate(data.UpstarLevel)
				if upstarTemp == nil {
					goto Level
				}
				for itemId, num := range upstarTemp.GetReturnItemMap() {
					_, ok := returnItemMap[itemId]
					if !ok {
						returnItemMap[itemId] = num
					} else {
						returnItemMap[itemId] += num
					}
				}
			}
		Level:
			if it.Level >= 0 {
				strengthenTemp := goldEquipTemp.GetStrengthenTemplate(it.Level)
				if strengthenTemp == nil {
					goto Done
				}
				for itemId, num := range strengthenTemp.GetReturnItemMap() {
					_, ok := returnItemMap[itemId]
					if !ok {
						returnItemMap[itemId] = num
					} else {
						returnItemMap[itemId] += num
					}
				}
			}
		Done:
		}

		// 分解系统装备
		sysEquipTemp := itemTemp.GetSystemEquipTemplate()
		if sysEquipTemp != nil {
			totalExp += int64(sysEquipTemp.GetTushiExp())
			for itemId, num := range sysEquipTemp.GetReturnItemMap() {
				_, ok := returnItemMap[itemId]
				if !ok {
					returnItemMap[itemId] = num
				} else {
					returnItemMap[itemId] += num
				}
			}
		}

	}
	flag = true
	return
}
