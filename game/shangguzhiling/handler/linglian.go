package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shangguzhilinglogic "fgame/fgame/game/shangguzhiling/logic"
	"fgame/fgame/game/shangguzhiling/pbutil"
	playershangguzhiling "fgame/fgame/game/shangguzhiling/player"
	shangguzhilingtemplate "fgame/fgame/game/shangguzhiling/template"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_LINGLIAN_TYPE), dispatch.HandlerFunc(handleShangguzhilingLingLian))
}

func handleShangguzhilingLingLian(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSShangguzhilingLingLian)
	lingshouType := shangguzhilingtypes.LingshouType(csMsg.GetType())
	autoFlag := csMsg.GetAutoFlag()
	changeStatusList := []shangguzhilingtypes.LinglianPosType{}
	for _, each := range csMsg.GetChangeStatusSubTypeList() {
		subType := shangguzhilingtypes.LinglianPosType(each)
		if !subType.Valid() {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"error":        err,
					"lingshouType": lingshouType,
					"subType":      subType,
				}).Warn("shangguzhiling:灵炼请求,灵炼部位类型错误")
			playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		}
		changeStatusList = append(changeStatusList, subType)
	}

	if !lingshouType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
			}).Warn("shangguzhiling:灵炼请求,灵兽类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = lingLian(tpl, lingshouType, autoFlag, changeStatusList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shangguzhiling:灵炼请求,错误")

		return err
	}
	return nil
}

//changeStatusList: 要锁定的部位
func lingLian(pl player.Player, lingshouType shangguzhilingtypes.LingshouType, autoFlag bool, changeStatusList []shangguzhilingtypes.LinglianPosType) (err error) {
	lingShouManager := pl.GetPlayerDataManager(playertypes.PlayerShangguzhilingDataManagerType).(*playershangguzhiling.PlayerShangguzhilingDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//是否解锁
	if !lingShouManager.IsLingShouUnlock(lingshouType) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
			}).Warn("shangguzhiling:灵炼请求,灵兽未解锁")
		playerlogic.SendSystemMessage(pl, lang.ShangguzhilingLingShouUnLock)
		return
	}
	obj := lingShouManager.GetLingShouObj(lingshouType)
	//校验输入部位是否解锁
	for _, pos := range changeStatusList {
		if !lingShouManager.IsLingLianPosJiesuo(lingshouType, pos) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"error":        err,
					"lingshouType": lingshouType,
					"pos":          pos,
				}).Warn("shangguzhiling:灵炼请求,灵兽未解锁")
			playerlogic.SendSystemMessage(pl, lang.ShangguzhilingLingLianUnLock)
			return
		}
	}

	//常量表
	constantTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetConstant()
	if constantTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
			}).Warn("shangguzhiling:灵炼请求,常量模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	//锁定消耗的物品
	itemMap := make(map[int32]int32)
	lockItemNum := lingShouManager.GetLingLianLockNeedNum(lingshouType, changeStatusList)
	lockItemId := constantTemp.LinglianSuodingUseItemId
	lockItemTemp := item.GetItemService().GetItem(int(lockItemId))
	if lockItemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
				"lockItemId":   lockItemId,
			}).Warn("shangguzhiling:灵炼请求,锁定使用的物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	if lockItemNum > 0 {
		itemMap[lockItemId] = lockItemNum
	}

	//使用的物品ID
	linglianItemId := constantTemp.LinglianItemId
	linglianItemTemp := item.GetItemService().GetItem(int(linglianItemId))
	if linglianItemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"error":          err,
				"lingshouType":   lingshouType,
				"linglianItemId": linglianItemId,
			}).Warn("shangguzhiling:灵炼请求,灵炼使用的物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//计算洗炼消耗数量
	linglianTimes := obj.GetLingLianTimes()
	maxTimes := constantTemp.LinglianNum
	if linglianTimes > maxTimes {
		linglianTimes = maxTimes
	}
	linglianItemCnt := constantTemp.LinglianItemCount*coreutils.Pow(linglianTimes+1, uint32(constantTemp.LinglianCoefficient)) + constantTemp.LinglianCoefficient2
	if linglianItemCnt > 0 {
		itemMap[linglianItemId] += linglianItemCnt
	}

	//物品数量
	shopIdMap := make(map[int32]int32)
	finalUseItemMap := make(map[int32]int32)
	needGold := int64(0)
	needBindGold := int64(0)
	needSilver := int64(0)
	for useItemId, useItemCount := range itemMap {
		shopIdTempMap := make(map[int32]int32)
		isEnoughBuyTimes := true
		finalUseItemMap[useItemId] = useItemCount
		itemCount := inventoryManager.NumOfItems(useItemId)
		if itemCount < useItemCount {
			if !autoFlag {
				log.WithFields(
					log.Fields{
						"playerId":     pl.GetId(),
						"error":        err,
						"lingshouType": lingshouType,
						"itemCount":    itemCount,
						"useItemCount": useItemCount,
					}).Warn("shangguzhiling:灵炼升级请求,选择使用的物品数量不足")
				playerlogic.SendSystemMessage(pl, lang.ShangguzhilingUseItemCountNotEnough)
				return
			}
			//自动进阶
			needBuyNum := useItemCount - itemCount
			finalUseItemMap[useItemId] = itemCount
			if needBuyNum > 0 {
				if !shop.GetShopService().ShopIsSellItem(useItemId) {
					log.WithFields(log.Fields{
						"playerId":  pl.GetId(),
						"useItemId": useItemId,
						"autoFlag":  autoFlag,
					}).Warn("shangguzhiling:商铺没有该道具,无法自动购买")
					playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
					return
				}

				isEnoughBuyTimes, shopIdTempMap = shoplogic.MaxBuyTimesForPlayer(pl, useItemId, needBuyNum)
				if !isEnoughBuyTimes {
					log.WithFields(log.Fields{
						"playerId":  pl.GetId(),
						"useItemId": useItemId,
						"autoFlag":  autoFlag,
					}).Warn("shangguzhiling:购买物品失败,上古之灵升级失败")
					playerlogic.SendSystemMessage(pl, lang.ShopBuyNumInvalid)
					return
				}

				shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdTempMap)
				needGold += shopNeedGold
				needBindGold += shopNeedBindGold
				needSilver += shopNeedSilver
				shopIdMap = coreutils.MergeMap(shopIdMap, shopIdTempMap)
			}
		}
	}
	//是否足够银两
	flag := propertyManager.HasEnoughSilver(needSilver)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shangguzhiling:上古之灵，银两不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//是否足够元宝
	flag = propertyManager.HasEnoughGold(needGold, false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shangguzhiling:上古之灵，元宝不足")
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
			}).Warn("shangguzhiling:上古之灵，绑元不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//-----------------------分割线-----------------------

	//灵炼
	lingShouManager.LingLian(lingshouType, changeStatusList)

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//自动购买消耗金钱
	//消耗银两
	if needSilver > 0 {
		silverUseReason := commonlog.SilverLogReasonLingLinglian
		silverUseReasonText := fmt.Sprintf(silverUseReason.String(), lingshouType.String())
		flag := propertyManager.CostSilver(needSilver, silverUseReason, silverUseReasonText)
		if !flag {
			panic("shangguzhiling:消耗银两应该成功")
		}
	}

	//消耗元宝
	if needGold > 0 {
		goldUseReason := commonlog.GoldLogReasonLingshouLinglian
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), lingshouType.String())
		flag := propertyManager.CostGold(needGold, false, goldUseReason, goldUseReasonText)
		if !flag {
			panic("shangguzhiling:消耗元宝应该成功")
		}
	}

	//消耗绑元
	if needBindGold > 0 {
		goldUseReason := commonlog.GoldLogReasonLingshouLinglian
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), lingshouType.String())
		flag := propertyManager.CostGold(needBindGold, true, goldUseReason, goldUseReasonText)
		if !flag {
			panic("shangguzhiling:消耗元宝应该成功")
		}
	}

	//同步物品（删掉吃掉的物品）
	useReason := commonlog.InventoryLogReasonLingShouLinglian
	useReasonText := fmt.Sprintf(useReason.String(), lingshouType.String())
	if len(finalUseItemMap) > 0 {
		flag = inventoryManager.BatchRemove(finalUseItemMap, useReason, useReasonText)
		if !flag {
			panic("inventory:移除物品应该是可以的")
		}
	}

	shangguzhilinglogic.LingShouPropertyChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//发送消息过滤锁定和未激活的
	linglianMap := make(map[shangguzhilingtypes.LinglianPosType]*shangguzhilingtypes.LinglianInfo)
	isLockMap := make(map[shangguzhilingtypes.LinglianPosType]bool)
	for _, subType := range changeStatusList {
		isLockMap[subType] = true
	}
	jiesuoLinglian := lingShouManager.GetLingLianPosJiesuoList(obj.GetLingShouType())
	for _, subType := range jiesuoLinglian {
		if isLockMap[subType] {
			continue
		}
		linglianMap[subType] = obj.GetLingLianInfo(subType)
	}
	scMsg := pbutil.BuildSCShangguzhilingLingLian(lingshouType, linglianMap, obj.GetLingLianTimes())
	pl.SendMsg(scMsg)
	return
}
