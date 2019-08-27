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
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	tianmologic "fgame/fgame/game/tianmo/logic"
	"fgame/fgame/game/tianmo/pbutil"
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotemplate "fgame/fgame/game/tianmo/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TIANMOTI_EAT_DAN_TYPE), dispatch.HandlerFunc(handleTianMoEatDan))
}

//处理使用天魔丹
func handleTianMoEatDan(s session.Session, msg interface{}) (err error) {
	log.Debug("tianMo:处理使用天魔丹")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSTianMoEatDan)
	num := csMsg.GetNum()

	err = tianMoEatDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("tianMo:处理使用天魔丹,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tianMo:处理使用天魔丹完成")
	return nil

}

// //食用天魔丹
// func tianMoEatDan(pl player.Player) (err error) {
// 	tianMoManager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianMo.PlayerTianMoDataManager)
// 	tianMoInfo := tianMoManager.GetTianMoInfo()
// 	advancedId := tianMoInfo.AdvanceId
// 	tianMoLevel := tianMoInfo.TianMoDanLevel
// 	tianMoTemplate := tianMo.GetTianMoTemplateService().GetTianMoNumber(int32(advancedId))
// 	if tianMoTemplate == nil {
// 		return
// 	}
// 	tianMoDanTemplate := tianMo.GetTianMoTemplateService().GetTianMoDan(tianMoLevel + 1)
// 	if tianMoDanTemplate == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("tianMo:天魔丹食丹等级满级")
// 		playerlogic.SendSystemMessage(pl, lang.TianMoEatDanReachedFull)
// 		return
// 	}

// 	if tianMoLevel >= tianMoTemplate.CulturingDanLimit {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("tianMo:天魔丹食丹等级已达最大,请进阶后再试")
// 		playerlogic.SendSystemMessage(pl, lang.TianMoEatDanReachedLimit)
// 		return
// 	}

// 	useItemMap := tianMoDanTemplate.GetUseItemTemplate()
// 	if len(useItemMap) != 0 {
// 		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		flag := inventoryManager.HasEnoughItems(useItemMap)
// 		if !flag {
// 			log.WithFields(
// 				log.Fields{
// 					"playerId": pl.GetId(),
// 				}).Warn("tianMo:当前天魔丹药数量不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}

// 		//消耗物品
// 		reasonText := commonlog.InventoryLogReasonTianMoEatDan.String()
// 		flag = inventoryManager.BatchRemove(useItemMap, commonlog.InventoryLogReasonTianMoEatDan, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("tianMo:BatchRemove should be ok"))
// 		}
// 		//同步物品
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}

// 	//培养判断
// 	pro, _, sucess := tianMologic.TianMoDan(tianMoInfo.TianMoDanNum, tianMoInfo.TianMoDanPro, tianMoDanTemplate)
// 	tianMoManager.EatTianMoDan(pro, sucess)
// 	if sucess {
// 		//同步属性
// 		tianMologic.TianMoPropertyChanged(pl)
// 	}

// 	scTianMoEatDan := pbutil.BuildSCTianMoEatDan(tianMoInfo.TianMoDanLevel, tianMoInfo.TianMoDanPro)
// 	pl.SendMsg(scTianMoEatDan)
// 	return
// }

//食用天魔丹
func tianMoEatDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("tianMo:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	tianMoManager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianMoInfo := tianMoManager.GetTianMoInfo()
	advancedId := tianMoInfo.AdvanceId
	tianMoLevel := tianMoInfo.TianMoDanLevel
	tianMoTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(advancedId)
	if tianMoTemplate == nil {
		return
	}
	tianMoDanTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoDan(tianMoLevel + 1)
	if tianMoDanTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tianMo:天魔丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.TianMoEatDanReachedFull)
		return
	}

	if tianMoLevel >= tianMoTemplate.ShidanLimit {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tianMo:天魔丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.TianMoEatDanReachedLimit)
		return
	}

	reachDanTemplate, flag := tianmotemplate.GetTianMoTemplateService().GetTianMoEatDanTemplate(tianMoLevel, num)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tianMo:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if reachDanTemplate.Level > tianMoTemplate.ShidanLimit {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("tianMo:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := tianMoDanTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("tianMo:当前丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		useItemReason := commonlog.InventoryLogReasonTianMoEatDan
		flag = inventoryManager.UseItem(useItem, num, useItemReason, useItemReason.String())
		if !flag {
			panic(fmt.Errorf("tianMo:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	tianMoManager.EatTianMoDan(reachDanTemplate.Level)
	//同步属性
	tianmologic.TianMoPropertyChanged(pl)

	scTianMoEatDan := pbutil.BuildSCTianMoEatDan(tianMoInfo.TianMoDanLevel, tianMoInfo.TianMoDanPro)
	pl.SendMsg(scTianMoEatDan)
	return
}
