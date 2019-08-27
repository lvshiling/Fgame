package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	additionsyseventtypes "fgame/fgame/game/additionsys/event/types"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystypes "fgame/fgame/game/additionsys/types"
	commontypes "fgame/fgame/game/common/types"
	gamevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_SHENG_JI_TYPE), dispatch.HandlerFunc(handleAdditionSysShengJi))
}

//处理附加系统升级
func handleAdditionSysShengJi(s session.Session, msg interface{}) (err error) {
	log.Debug("additionsys:处理附加系统升级")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAdditionSysShengJi := msg.(*uipb.CSAdditionSysShengJi)
	autoFlag := csAdditionSysShengJi.GetAuto()
	sysTypeId := csAdditionSysShengJi.GetSysType()
	sysType := additionsystypes.AdditionSysType(sysTypeId)

	//参数不对
	if !sysType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
			}).Warn("additionsys:升级系统类型,错误")
		return
	}

	err = additionSysShengJi(tpl, sysType, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
				"error":    err,
			}).Warn("additionsys:升级系统类型,错误")
		return
	}

	log.Debug("additionsys:处理附加系统升级,完成")
	return nil

}

//系统升级逻辑
func additionSysShengJi(pl player.Player, typ additionsystypes.AdditionSysType, autoFlag bool) (err error) {
	if !additionsyslogic.GetAdditionSysShengJiFuncOpenByType(pl, typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("inventory:升级失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	levelInfo := additionsysManager.GetAdditionSysLevelInfoByType(typ)

	//判断槽位是否可以升
	nextShengJiTemplate := levelInfo.GetNextShengJiTemplate()
	if nextShengJiTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"autoFlag": autoFlag,
			}).Warn("additionsyse:升级失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysLevelHighest)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的银两
	costSilver := int64(nextShengJiTemplate.GetUseMoney())
	//进阶需要消耗的元宝
	costGold := int32(0)
	//进阶需要消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := nextShengJiTemplate.GetUseItem()
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := nextShengJiTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = nextShengJiTemplate.GetItemCount()
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"autoFlag": autoFlag,
			}).Warn("additionsyse:物品不足,无法升级")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动进阶
		needBuyNum := itemCount - totalNum
		itemCount = totalNum

		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItem) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ.String(),
					"autoFlag": autoFlag,
				}).Warn("additionsyse:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ.String(),
					"autoFlag": autoFlag,
				}).Warn("additionsyse:购买物品失败,自动升级已停止")
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
				"sysType":  typ.String(),
				"autoFlag": autoFlag,
			}).Warn("additionsys:银两不足,无法升级")
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
				"sysType":  typ.String(),
				"autoFlag": autoFlag,
			}).Warn("additionsys:元宝不足,无法升级")
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
				"sysType":  typ.String(),
				"autoFlag": autoFlag,
			}).Warn("additionsys:元宝不足,无法升级")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	goldUseReason := commonlog.GoldLogReasonAdditionSysLevel
	silverUseReason := commonlog.SilverLogReasonAdditionSysLevel
	goldUseReasonText := fmt.Sprintf(commonlog.GoldLogReasonAdditionSysLevel.String(), typ.String(), levelInfo.Level, levelInfo.UpPro, levelInfo.UpNum)
	silverUseReasonText := fmt.Sprintf(commonlog.SilverLogReasonAdditionSysLevel.String(), typ.String(), levelInfo.Level, levelInfo.UpPro, levelInfo.UpNum)
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), goldUseReason, goldUseReasonText, costSilver, silverUseReason, silverUseReasonText)
	if !flag {
		panic(fmt.Errorf("additionsys: additionsys shengji Cost should be ok"))
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonAdditionSysShengJi
		reasonText := fmt.Sprintf(inventoryReason.String(), typ.String(), levelInfo.Level)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("additionsys: shengji use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}
	// if useItemTemplate != nil {
	// 	//消耗事件
	// 	gameevent.Emit(xiantieventtypes.EventTypeXianTiAdvancedCost, pl, int32(xianTiInfo.AdvanceId))
	// }
	if nextShengJiTemplate != nil {
		itemMap := map[int32]int32{
			nextShengJiTemplate.GetUseItem(): nextShengJiTemplate.GetItemCount(),
		}
		useItemEventData := additionsyseventtypes.CreatePlayerAdditionSysUseItemEventData(itemMap)
		gamevent.Emit(additionsyseventtypes.EventTypeAdditionSysUseItem, pl, useItemEventData)
	}

	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	beforeLev := levelInfo.Level
	beforeUpNum := levelInfo.UpNum
	beforeUpPro := levelInfo.UpPro
	//进阶判断
	sucess, pro, _, addTimes := additionsyslogic.AdditionSysLevelJudge(pl, levelInfo.UpNum, levelInfo.UpPro, nextShengJiTemplate)
	additionsysManager.SystemUplevel(typ, pro, addTimes, sucess)

	res := additionsystypes.ShengJiResTypeDefeated //用于通知前段结果
	if sucess {
		//更新属性
		additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)
		res = additionsystypes.ShengJiResTypeSucceed
	}
	//日志
	additionsysReason := commonlog.AdditionSysLogReasonShengJi
	reasonText := fmt.Sprintf(additionsysReason.String(), typ.String(), commontypes.AdvancedTypeSilver.String())
	data := additionsyseventtypes.CreatePlayerAdditionSysShengJiLogEventData(typ, beforeLev, beforeUpNum, beforeUpPro, additionsysReason, reasonText)
	gamevent.Emit(additionsyseventtypes.EventTypeAdditionSysShengJiLog, pl, data)

	scAdditionSysShengJi := pbutil.BuildSCAdditionSysShengJi(int32(levelInfo.SysType), levelInfo.Level, levelInfo.UpPro, int32(res))
	pl.SendMsg(scAdditionSysShengJi)
	return
}
