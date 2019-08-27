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
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	winglogic "fgame/fgame/game/wing/logic"
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"
	"fgame/fgame/game/wing/wing"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_WING_UNREALDAN_TYPE), dispatch.HandlerFunc(handleWingUnrealDan))

}

//处理战翼食幻化丹信息
func handleWingUnrealDan(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理战翼食幻化丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWingUnrealDan := msg.(*uipb.CSWingUnrealDan)
	num := csWingUnrealDan.GetNum()

	err = wingUnrealDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("wing:处理战翼食幻化丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("wing:处理战翼食幻化丹信息完成")
	return nil

}

// //战翼食幻化丹的逻辑
// func wingUnrealDan(pl player.Player) (err error) {
// 	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
// 	wingInfo := wingManager.GetWingInfo()
// 	advancedId := wingInfo.AdvanceId
// 	unrealLevel := wingInfo.UnrealLevel
// 	wingTemplate := wing.GetWingService().GetWingNumber(int32(advancedId))
// 	if wingTemplate == nil {
// 		return
// 	}

// 	huanHuaTemplate := wing.GetWingService().GetWingHuanHuaTemplate(unrealLevel + 1)
// 	if huanHuaTemplate == nil {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Warn("wing:幻化丹食丹等级满级")
// 		playerlogic.SendSystemMessage(pl, lang.WingUnrealDanReachedFull)
// 		return
// 	}

// 	if unrealLevel >= wingTemplate.ShidanLimit {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Warn("wing:幻化丹食丹等级已达最大,请进阶后再试")
// 		playerlogic.SendSystemMessage(pl, lang.WingUnrealDanReachedLimit)
// 		return
// 	}

// 	useItemMap := huanHuaTemplate.GetUseItemTemplate()
// 	if len(useItemMap) != 0 {
// 		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		flag := inventoryManager.HasEnoughItems(useItemMap)
// 		if !flag {
// 			log.WithFields(log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("wing:当前幻化丹药数量不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}
// 		//消耗物品
// 		reasonText := commonlog.InventoryLogReasonWingEatUn.String()
// 		flag = inventoryManager.BatchRemove(useItemMap, commonlog.InventoryLogReasonWingEatUn, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("wing:BatchRemove should be ok"))
// 		}
// 		//同步物品
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}

// 	//战翼幻化丹判断
// 	pro, _, sucess := winglogic.WingHuanHua(wingInfo.UnrealNum, wingInfo.UnrealPro, huanHuaTemplate)
// 	wingManager.EatUnrealDan(pro, sucess)
// 	if sucess {
// 		//同步属性
// 		winglogic.WingPropertyChanged(pl)
// 	}
// 	scWingShiDan := pbutil.BuildSCWingUnrealDan(wingInfo.UnrealLevel, wingInfo.UnrealPro)
// 	pl.SendMsg(scWingShiDan)
// 	return
// }

//战翼食幻化丹的逻辑
func wingUnrealDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("wing:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	advancedId := wingInfo.AdvanceId
	unrealLevel := wingInfo.UnrealLevel
	wingTemplate := wing.GetWingService().GetWingNumber(int32(advancedId))
	if wingTemplate == nil {
		return
	}

	huanHuaTemplate := wing.GetWingService().GetWingHuanHuaTemplate(unrealLevel + 1)
	if huanHuaTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("wing:幻化丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.WingUnrealDanReachedFull)
		return
	}

	if unrealLevel >= wingTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("wing:幻化丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.WingUnrealDanReachedLimit)
		return
	}

	reachHuanHuaTemplate, flag := wing.GetWingService().GetWingEatHuanHuanTemplate(unrealLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("wing:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachHuanHuaTemplate.Level > wingTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("wing:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := huanHuaTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("wing:当前幻化丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonWingEatUn.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonWingEatUn, reasonText)
		if !flag {
			panic(fmt.Errorf("wing:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	wingManager.EatUnrealDan(reachHuanHuaTemplate.Level)
	//同步属性
	winglogic.WingPropertyChanged(pl)

	scWingShiDan := pbutil.BuildSCWingUnrealDan(wingInfo.UnrealLevel, wingInfo.UnrealPro)
	pl.SendMsg(scWingShiDan)
	return
}
