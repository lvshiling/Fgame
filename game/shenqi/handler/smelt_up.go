package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shenqieventtypes "fgame/fgame/game/shenqi/event/types"
	shenqilogic "fgame/fgame/game/shenqi/logic"
	"fgame/fgame/game/shenqi/pbutil"
	playershenqi "fgame/fgame/game/shenqi/player"
	shenqitypes "fgame/fgame/game/shenqi/types"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENQI_SMELT_UP_TYPE), dispatch.HandlerFunc(handleShenQiSmeltUp))
}

//处理神器淬炼升级
func handleShenQiSmeltUp(s session.Session, msg interface{}) (err error) {
	log.Debug("smeltup:处理神器淬炼升级")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csShenqiZhuling := msg.(*uipb.CSShenqiSmeltUp)
	typInt := csShenqiZhuling.GetShenQiType()
	slotIdInt := csShenqiZhuling.GetSlotId()
	auto := csShenqiZhuling.GetAuto()
	typ := shenqitypes.ShenQiType(typInt)
	slotId := shenqitypes.SmeltType(slotIdInt)

	//参数不对
	if !typ.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typInt":   typInt,
			}).Warn("smeltup:淬炼类型,错误")
		return
	}
	if !slotId.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"slotIdInt": slotIdInt,
			}).Warn("smeltup:淬炼类型,错误")
		return
	}

	err = shenQiSmeltUp(tpl, typ, slotId, auto)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("smeltup:处理神器淬炼升级,错误")

		return err
	}
	log.Debug("smeltup:处理神器淬炼升级,完成")
	return nil
}

//神器淬炼升级
func shenQiSmeltUp(pl player.Player, typ shenqitypes.ShenQiType, slotId shenqitypes.SmeltType, buyFlag bool) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenQiSmelt) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("smeltup:淬炼升级失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	shenQiManager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	nextTemplate := shenQiManager.GetNextSmeltUpTemplate(typ, slotId)
	if nextTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"slotId":   slotId.String(),
			}).Warn("smeltup:处理神器淬炼升级,已经最高级")
		playerlogic.SendSystemMessage(pl, lang.ShenQiSlotLevelMax)
		return
	}

	if nextTemplate.NeedShenQiLevel > shenQiManager.GetShenQiDebrisMinLevelByShenQi(typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"slotId":   slotId.String(),
			}).Warn("smeltup:处理神器淬炼升级,神器等级不足")
		playerlogic.SendSystemMessage(pl, lang.ShenQiLevelNotEnough)
		return
	}

	slotObj := shenQiManager.GetShenQiSmeltMapByArg(typ, slotId)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//升级需要消耗的银两
	costSilver := int64(0)
	//升级需要的消耗的绑元
	costBindGold := int32(0)
	//升级需要消耗的元宝
	costGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := nextTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := nextTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = nextTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if buyFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"buyFlag":  buyFlag,
			}).Warn("smeltup:神器淬炼进物品不足,无法升级")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动升级
		needBuyNum := itemCount - totalNum
		itemCount = totalNum

		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItem) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"buyFlag":  buyFlag,
				}).Warn("smeltup:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"buyFlag":  buyFlag,
				}).Warn("smeltup:购买物品失败,自动升级已停止")
				playerlogic.SendSystemMessage(pl, lang.ShopAdvancedAutoBuyItemFail)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			costGold += int32(shopNeedGold)
			costBindGold += int32(shopNeedBindGold)
			costSilver += shopNeedSilver
		}
	}

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"buyFlag":  buyFlag,
			}).Warn("smeltup:银两不足,无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"buyFlag":  buyFlag,
			}).Warn("smeltup:元宝不足,无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够绑元
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"buyFlag":  buyFlag,
			}).Warn("smeltup:元宝不足,无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGold := commonlog.GoldLogReasonShenQiSmeltUpCost
	reasonSliver := commonlog.SilverLogReasonShenQiSmeltUpCost
	reasonGoldText := fmt.Sprintf(reasonGold.String(), typ.String(), slotId.String())
	reasonSliverText := fmt.Sprintf(reasonSliver.String(), typ.String(), slotId.String())
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), reasonGold, reasonGoldText, costSilver, reasonSliver, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("smeltup: uplevel Cost should be ok"))
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonShenQiSmeltUpCost
		reasonText := fmt.Sprintf(inventoryReason.String(), typ.String(), slotId.String())
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("smeltup: uplevel use item should be ok"))
		}
	}

	if itemCount > 0 {
		inventorylogic.SnapInventoryChanged(pl)
	}

	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//升级判断
	sucess, pro, _, addTimes, _ := shenqilogic.SmeltUpJudge(pl, slotObj.UpNum, slotObj.UpPro, nextTemplate)
	befLev := slotObj.Level
	shenQiManager.SmeltUpAdvanced(typ, slotId, pro, addTimes, sucess)
	if sucess {
		shenQiManager.RefreshShenQiSmeltMinLevel(typ)
		//同步属性
		shenqilogic.ShenQiPropertyChanged(pl)
		//日志
		logReason := commonlog.ShenQiLogReasonRelatedUpLevel
		reasonText := fmt.Sprintf(logReason.String(), typ.String(), slotId.String(), commontypes.AdvancedTypeElement.String())
		logData := shenqieventtypes.CreatePlayerShenQiRelatedUpLevelLogEventData(befLev, slotObj.Level, logReason, reasonText)
		gameevent.Emit(shenqieventtypes.EventTypeShenQiRelatedUpLevelLog, pl, logData)
	}

	scMsg := pbutil.BuildSCShenQiSmeltUp(slotObj, buyFlag)
	pl.SendMsg(scMsg)
	return
}
