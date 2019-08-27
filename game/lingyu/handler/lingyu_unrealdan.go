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
	lingyulogic "fgame/fgame/game/lingyu/logic"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGYU_UNREALDAN_TYPE), dispatch.HandlerFunc(handleLingyuUnrealDan))

}

//处理领域食幻化丹信息
func handleLingyuUnrealDan(s session.Session, msg interface{}) (err error) {
	log.Debug("lingyu:处理领域食幻化丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingyuUnrealDan := msg.(*uipb.CSLingyuUnrealDan)
	num := csLingyuUnrealDan.GetNum()

	err = lingyuUnrealDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("lingyu:处理领域食幻化丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingyu:处理领域食幻化丹信息完成")
	return nil

}

// //领域食幻化丹的逻辑
// func lingyuUnrealDan(pl player.Player) (err error) {
// 	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
// 	lingyuInfo := lingyuManager.GetLingyuInfo()
// 	advancedId := lingyuInfo.AdvanceId
// 	unrealLevel := lingyuInfo.UnrealLevel
// 	lingyuTemplate := lingyu.GetLingyuTemplateService().GetLingyuByNumber(int32(advancedId))

// 	hunaHuaTemplate := lingyu.GetLingyuTemplateService().GetLingyuHuanHuaTemplate(unrealLevel + 1)
// 	if hunaHuaTemplate == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerid": pl.GetId(),
// 			}).Warn("lingyu:幻化丹食丹等级满级")
// 		playerlogic.SendSystemMessage(pl, lang.LingyuUnrealDanReachedFull)
// 		return
// 	}

// 	if unrealLevel >= lingyuTemplate.ShidanLimit {
// 		log.WithFields(
// 			log.Fields{
// 				"playerid": pl.GetId(),
// 			}).Warn("lingyu:幻化丹食丹等级已达最大,请进阶后再试")
// 		playerlogic.SendSystemMessage(pl, lang.LingyuUnrealDanReachedLimit)
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
// 				}).Warn("lingyu:当前幻化丹药数量不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}

// 		//消耗物品
// 		reasonText := commonlog.InventoryLogReasonLingyuEatUn.String()
// 		flag = inventoryManager.BatchRemove(useItemMap, commonlog.InventoryLogReasonLingyuEatUn, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("lingyu:BatchRemove should be ok"))
// 		}
// 		//同步物品
// 		inventorylogic.SnapInventoryChanged(pl)

// 	}

// 	//领域幻化丹培养判断
// 	pro, _, sucess := lingyulogic.LingYuHuanHua(lingyuInfo.UnrealNum, lingyuInfo.UnrealPro, hunaHuaTemplate)
// 	lingyuManager.EatUnrealDan(pro, sucess)
// 	if sucess {
// 		//同步属性
// 		lingyulogic.LingyuPropertyChanged(pl)
// 	}

// 	scLingyuShiDan := pbutil.BuildSCLingyuUnrealDan(lingyuInfo.UnrealLevel, lingyuInfo.UnrealPro)
// 	pl.SendMsg(scLingyuShiDan)
// 	return
// }

//领域食幻化丹的逻辑
func lingyuUnrealDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("lingyu:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuInfo := lingyuManager.GetLingyuInfo()
	advancedId := lingyuInfo.AdvanceId
	if advancedId <= 0 {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("lingyu:幻化丹食丹错误，领域阶数为0")
		playerlogic.SendSystemMessage(pl, lang.LingyuAdvanceToLow)
		return
	}

	unrealLevel := lingyuInfo.UnrealLevel
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(int32(advancedId))

	hunaHuaTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuHuanHuaTemplate(unrealLevel + 1)
	if hunaHuaTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("lingyu:幻化丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.LingyuUnrealDanReachedFull)
		return
	}

	if unrealLevel >= lingyuTemplate.ShidanLimit {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("lingyu:幻化丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.LingyuUnrealDanReachedLimit)
		return
	}

	reachHuanHuaTemplate, flag := lingyutemplate.GetLingyuTemplateService().GetLingYuEatHuanHuaTemplate(unrealLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("lingyu:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachHuanHuaTemplate.Level > lingyuTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("lingyu:参数错误")
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
			}).Warn("lingyu:当前幻化丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonLingyuEatUn.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonLingyuEatUn, reasonText)
		if !flag {
			panic(fmt.Errorf("lingyu:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	lingyuManager.EatUnrealDan(reachHuanHuaTemplate.Level)
	//同步属性
	lingyulogic.LingyuPropertyChanged(pl)

	scLingyuShiDan := pbutil.BuildSCLingyuUnrealDan(lingyuInfo.UnrealLevel, lingyuInfo.UnrealPro)
	pl.SendMsg(scLingyuShiDan)
	return
}
