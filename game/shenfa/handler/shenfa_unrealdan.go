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
	shenfalogic "fgame/fgame/game/shenfa/logic"
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENFA_UNREALDAN_TYPE), dispatch.HandlerFunc(handleShenfaUnrealDan))

}

//处理身法食幻化丹信息
func handleShenfaUnrealDan(s session.Session, msg interface{}) (err error) {
	log.Debug("shenfa:处理身法食幻化丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShenfaUnrealDan := msg.(*uipb.CSShenfaUnrealDan)
	num := csShenfaUnrealDan.GetNum()

	err = shenfaUnrealDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("shenfa:处理身法食幻化丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenfa:处理身法食幻化丹信息完成")
	return nil

}

// //身法食幻化丹的逻辑
// func shenfaUnrealDan(pl player.Player) (err error) {
// 	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
// 	shenfaInfo := shenfaManager.GetShenfaInfo()
// 	advancedId := shenfaInfo.AdvanceId
// 	unrealLevel := shenfaInfo.UnrealLevel
// 	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(int32(advancedId))

// 	hunaHuaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaHuanHuaTemplate(unrealLevel + 1)
// 	if hunaHuaTemplate == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerid": pl.GetId(),
// 			}).Warn("shenfa:幻化丹食丹等级满级")
// 		playerlogic.SendSystemMessage(pl, lang.ShenfaUnrealDanReachedFull)
// 		return
// 	}

// 	if unrealLevel >= shenfaTemplate.ShidanLimit {
// 		log.WithFields(
// 			log.Fields{
// 				"playerid": pl.GetId(),
// 			}).Warn("shenfa:幻化丹食丹等级已达最大,请进阶后再试")
// 		playerlogic.SendSystemMessage(pl, lang.ShenfaUnrealDanReachedLimit)
// 		return
// 	}

// 	useItemMap := hunaHuaTemplate.GetUseItemTemplate()
// 	if len(useItemMap) != 0 {
// 		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		flag := inventoryManager.HasEnoughItems(useItemMap)
// 		if !flag {
// 			log.WithFields(
// 				log.Fields{
// 					"playerid": pl.GetId(),
// 				}).Warn("shenfa:当前幻化丹药数量不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}

// 		//消耗物品
// 		reasonText := commonlog.InventoryLogReasonShenfaEatUn.String()
// 		flag = inventoryManager.BatchRemove(useItemMap, commonlog.InventoryLogReasonShenfaEatUn, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("shenfa:BatchRemove should be ok"))
// 		}
// 		//同步物品
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}

// 	//身法幻化丹培养判断
// 	pro, _, sucess := shenfalogic.ShenFaHuanHua(shenfaInfo.UnrealNum, shenfaInfo.UnrealPro, hunaHuaTemplate)
// 	shenfaManager.EatUnrealDan(pro, sucess)
// 	if sucess {
// 		//同步属性
// 		shenfalogic.ShenfaPropertyChanged(pl)
// 	}

// 	scShenfaShiDan := pbutil.BuildSCShenfaUnrealDan(shenfaInfo.UnrealLevel, shenfaInfo.UnrealPro)
// 	pl.SendMsg(scShenfaShiDan)
// 	return
// }

//身法食幻化丹的逻辑
func shenfaUnrealDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"num":      num,
			}).Warn("shenfa:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenfaInfo := shenfaManager.GetShenfaInfo()
	advancedId := shenfaInfo.AdvanceId
	unrealLevel := shenfaInfo.UnrealLevel
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(int32(advancedId))

	hunaHuaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaHuanHuaTemplate(unrealLevel + 1)
	if hunaHuaTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("shenfa:幻化丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.ShenfaUnrealDanReachedFull)
		return
	}

	if unrealLevel >= shenfaTemplate.ShidanLimit {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("shenfa:幻化丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.ShenfaUnrealDanReachedLimit)
		return
	}

	reachHuanHuaTemplate, flag := shenfatemplate.GetShenfaTemplateService().GetShenFaEatHuanHuaTemplate(unrealLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("wing:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachHuanHuaTemplate.Level > shenfaTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("wing:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := hunaHuaTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("shenfa:当前幻化丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonShenfaEatUn.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonShenfaEatUn, reasonText)
		if !flag {
			panic(fmt.Errorf("shenfa:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	shenfaManager.EatUnrealDan(reachHuanHuaTemplate.Level)
	//同步属性
	shenfalogic.ShenfaPropertyChanged(pl)

	scShenfaShiDan := pbutil.BuildSCShenfaUnrealDan(shenfaInfo.UnrealLevel, shenfaInfo.UnrealPro)
	pl.SendMsg(scShenfaShiDan)
	return
}
