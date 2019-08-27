package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	anqilogic "fgame/fgame/game/anqi/logic"
	"fgame/fgame/game/anqi/pbutil"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ANQI_EAT_DAN_TYPE), dispatch.HandlerFunc(handleAnqiEatDan))
}

//处理使用暗器丹
func handleAnqiEatDan(s session.Session, msg interface{}) (err error) {
	log.Debug("anqi:处理使用暗器丹")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAnqiEatDan := msg.(*uipb.CSAnqiEatDan)
	num := csAnqiEatDan.GetNum()

	err = anqiEatDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("anqi:处理使用暗器丹,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("anqi:处理使用暗器丹完成")
	return nil

}

// //食用暗器丹
// func anqiEatDan(pl player.Player) (err error) {
// 	anqiManager := pl.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
// 	anqiInfo := anqiManager.GetAnqiInfo()
// 	advancedId := anqiInfo.AdvanceId
// 	anqiLevel := anqiInfo.AnqiDanLevel
// 	anqiTemplate := anqi.GetAnqiTemplateService().GetAnqiNumber(int32(advancedId))
// 	if anqiTemplate == nil {
// 		return
// 	}
// 	anqiDanTemplate := anqi.GetAnqiTemplateService().GetAnqiDan(anqiLevel + 1)
// 	if anqiDanTemplate == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("anqi:暗器丹食丹等级满级")
// 		playerlogic.SendSystemMessage(pl, lang.AnqiEatDanReachedFull)
// 		return
// 	}

// 	if anqiLevel >= anqiTemplate.CulturingDanLimit {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("anqi:暗器丹食丹等级已达最大,请进阶后再试")
// 		playerlogic.SendSystemMessage(pl, lang.AnqiEatDanReachedLimit)
// 		return
// 	}

// 	useItemMap := anqiDanTemplate.GetUseItemTemplate()
// 	if len(useItemMap) != 0 {
// 		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		flag := inventoryManager.HasEnoughItems(useItemMap)
// 		if !flag {
// 			log.WithFields(
// 				log.Fields{
// 					"playerId": pl.GetId(),
// 				}).Warn("anqi:当前暗器丹药数量不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}

// 		//消耗物品
// 		reasonText := commonlog.InventoryLogReasonAnqiEatDan.String()
// 		flag = inventoryManager.BatchRemove(useItemMap, commonlog.InventoryLogReasonAnqiEatDan, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("anqi:BatchRemove should be ok"))
// 		}
// 		//同步物品
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}

// 	//培养判断
// 	pro, _, sucess := anqilogic.AnqiDan(anqiInfo.AnqiDanNum, anqiInfo.AnqiDanPro, anqiDanTemplate)
// 	anqiManager.EatAnqiDan(pro, sucess)
// 	if sucess {
// 		//同步属性
// 		anqilogic.AnqiPropertyChanged(pl)
// 	}

// 	scAnqiEatDan := pbutil.BuildSCAnqiEatDan(anqiInfo.AnqiDanLevel, anqiInfo.AnqiDanPro)
// 	pl.SendMsg(scAnqiEatDan)
// 	return
// }

//食用暗器丹
func anqiEatDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("anqi:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	anqiManager := pl.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiInfo := anqiManager.GetAnqiInfo()
	advancedId := anqiInfo.AdvanceId
	anqiLevel := anqiInfo.AnqiDanLevel
	anqiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(advancedId))
	if anqiTemplate == nil {
		return
	}
	anqiDanTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiDan(anqiLevel + 1)
	if anqiDanTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("anqi:暗器丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.AnqiEatDanReachedFull)
		return
	}

	if anqiLevel >= anqiTemplate.CulturingDanLimit {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("anqi:暗器丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.AnqiEatDanReachedLimit)
		return
	}

	reachDanTemplate, flag := anqitemplate.GetAnqiTemplateService().GetAnQiEatDanTemplate(anqiLevel, num)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("anqi:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachDanTemplate.Level > anqiTemplate.CulturingDanLimit {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("anqi:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := anqiDanTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("anqi:当前丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonAnqiEatDan.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonAnqiEatDan, reasonText)
		if !flag {
			panic(fmt.Errorf("anqi:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	anqiManager.EatAnqiDan(reachDanTemplate.Level)
	//同步属性
	anqilogic.AnqiPropertyChanged(pl)

	scAnqiEatDan := pbutil.BuildSCAnqiEatDan(anqiInfo.AnqiDanLevel, anqiInfo.AnqiDanPro)
	pl.SendMsg(scAnqiEatDan)
	return
}
