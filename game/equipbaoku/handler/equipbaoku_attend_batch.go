package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
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
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_ATTEND_BATCH_TYPE), dispatch.HandlerFunc(handleEquipBaoKuAttendBatch))

}

//批量探索装备宝库
func handleEquipBaoKuAttendBatch(s session.Session, msg interface{}) (err error) {
	log.Debug("equipbaoku:批量探索装备宝库")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csEquipBaoKuAttendBatch := msg.(*uipb.CSEquipbaokuAttendBatch)
	attendNum := csEquipBaoKuAttendBatch.GetAttendNum()
	logTime := csEquipBaoKuAttendBatch.GetLogTime()
	autoFlag := csEquipBaoKuAttendBatch.GetAutoFlag()
	typ := equipbaokutypes.BaoKuType(csEquipBaoKuAttendBatch.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"typ":      typ,
				"error":    err,
			}).Warn("equipbaoku:处理探索宝库,宝库类型不合法")
		return
	}

	err = equipBaoKuAttendBatch(tpl, attendNum, logTime, autoFlag, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("equipbaoku:处理批量探索装备宝库,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("equipbaoku:处理批量探索装备宝库完成")
	return nil

}

//批量探索装备宝库逻辑
func equipBaoKuAttendBatch(pl player.Player, attendNum int32, logTime int64, autoFlag bool, typ equipbaokutypes.BaoKuType) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeEquipBaoKu) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:批量探索宝库错误,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}
	equipBaoKuManager := pl.GetPlayerDataManager(playertypes.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	equipBaoKuTemplate := equipbaokutemplate.GetEquipBaoKuTemplateService().GetEquipBaoKuByLevAndZhuanNum(pl.GetLevel(), pl.GetZhuanSheng(), typ)
	if equipBaoKuTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:批量探索宝库错误,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needBindGold := int32(equipBaoKuTemplate.BindGoldUse * attendNum)
	needGold := int32(equipBaoKuTemplate.GoldUse * attendNum)
	needSilver := int64(equipBaoKuTemplate.SilverUse * attendNum)
	needItemId := equipBaoKuTemplate.UseItemId
	needItemCount := equipBaoKuTemplate.UseItemCount * attendNum

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
		//自动进阶
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
			needGold += int32(shopNeedGold)
			needBindGold += int32(shopNeedBindGold)
			needSilver += shopNeedSilver
		}

	}

	//是否足够银两
	flag := propertyManager.HasEnoughSilver(needSilver)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:批量探索宝库错误，银两不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//是否足够元宝
	flag = propertyManager.HasEnoughGold(int64(needGold), false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:批量探索宝库错误，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//是否足够绑元
	needCostBindGold := needBindGold + needGold
	flag = propertyManager.HasEnoughGold(int64(needCostBindGold), true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("equipbaoku:批量探索宝库错误，绑元不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//批量处理
	rewList := equipBaoKuManager.GetEquipBaoKuDrop(attendNum, typ)
	if len(rewList) == 0 {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"attendTimes ": attendNum,
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
		flag := propertyManager.CostGold(int64(needGold), false, goldUseReason, goldUseReasonText)
		if !flag {
			panic("equipbaoku:消耗元宝应该成功")
		}
	}

	//消耗绑元
	if needBindGold > 0 {
		goldUseReason := commonlog.GoldLogReasonEquipBaoKuUse
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		flag := propertyManager.CostGold(int64(needBindGold), true, goldUseReason, goldUseReasonText)
		if !flag {
			panic("equipbaoku:消耗元宝应该成功")
		}
	}

	//消耗物品
	if needItemCount > 0 {
		itemUseReason := commonlog.InventoryLogReasonEquipBaoKuAttend
		itemUseReasonText := fmt.Sprintf(itemUseReason.String(), pl.GetLevel(), pl.GetZhuanSheng())
		if flag := inventoryManager.UseItem(needItemId, needItemCount, itemUseReason, itemUseReasonText); !flag {
			panic(fmt.Errorf("equipbaoku: attend equipbaoku use item should be ok"))
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
	addXingYunZhi := equipBaoKuTemplate.GiftXingYunZhi * attendNum
	addJiFen := equipBaoKuTemplate.GiftJiFen * attendNum
	isDouble, luckyPointCritNum, attendPointCritNum := welfarelogic.IsCanDrewBaoKuCrit()
	if isDouble {
		addXingYunZhi = addXingYunZhi + int32(math.Ceil(float64(luckyPointCritNum)/float64(common.MAX_RATE)*float64(addXingYunZhi)))
		addJiFen = addJiFen + int32(math.Ceil(float64(attendPointCritNum)/float64(common.MAX_RATE)*float64(addJiFen)))
	}
	equipBaoKuManager.AttendEquipBaoKu(addXingYunZhi, addJiFen, attendNum, commontypes.ChangeTypeAttendGet, typ)

	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapMiBaoDepotChanged(pl, typ)
	inventorylogic.SnapInventoryChanged(pl)

	luckyPoints := equipBaoKuManager.GetEquipBaoKuObj(typ).GetLuckyPoints()
	attendPoints := equipBaoKuManager.GetEquipBaoKuObj(typ).GetAttendPoints()
	logList := equipbaoku.GetEquipBaoKuService().GetLogByTime(logTime, typ)
	scEquipBaoKuAttendBatch := pbutil.BuildSCEquipBaoKuAttendBatch(rewList, logList, luckyPoints, attendPoints, autoFlag, int32(typ))
	pl.SendMsg(scEquipBaoKuAttendBatch)
	return
}
