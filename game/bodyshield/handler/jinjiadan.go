package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/bodyshield/bodyshield"
	bodyShieldlogic "fgame/fgame/game/bodyshield/logic"
	"fgame/fgame/game/bodyshield/pbutil"
	playerbshield "fgame/fgame/game/bodyshield/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BODYSHIELD_JJDAN_TYPE), dispatch.HandlerFunc(handleBodyShieldJJDan))
}

//处理护体盾金甲丹信息
func handleBodyShieldJJDan(s session.Session, msg interface{}) (err error) {
	log.Debug("bodyShield:处理获取护体盾金甲丹消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csBodyShieldJJDan := msg.(*uipb.CSBodyShieldJJDan)
	num := csBodyShieldJJDan.GetNum()

	err = bodyShieldJJDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("bodyShield:处理获取护体盾金甲丹消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("bodyShield:处理获取护体盾金甲丹消息完成")
	return nil

}

// //食用护体盾金甲丹
// func bodyShieldJJDan(pl player.Player) (err error) {
// 	bshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbshield.PlayerBodyShieldDataManager)
// 	bodyShieldInfo := bshieldManager.GetBodyShiedInfo()
// 	advancedId := bodyShieldInfo.AdvanceId
// 	jinJiaLevel := bodyShieldInfo.JinjiadanLevel
// 	bodyShieldTemplate := bodyshield.GetBodyShieldService().GetBodyShieldNumber(int32(advancedId))
// 	if bodyShieldTemplate == nil {
// 		return
// 	}
// 	jinJiaDanTemplate := bodyshield.GetBodyShieldService().GetBodyShieldJinJia(jinJiaLevel + 1)
// 	if jinJiaDanTemplate == nil {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Warn("bodyShiled:护体金甲丹食丹等级满级")
// 		playerlogic.SendSystemMessage(pl, lang.BodyShieldEatJJDanReachedFull)
// 		return
// 	}

// 	if jinJiaLevel >= bodyShieldTemplate.ShidanLimit {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Warn("bodyShiled:护体金甲丹食丹等级已达最大,请进阶后再试")
// 		playerlogic.SendSystemMessage(pl, lang.BodyShieldEatJJDanReachedLimit)
// 		return
// 	}

// 	useItemMap := jinJiaDanTemplate.GetUseItemTemplate()
// 	if len(useItemMap) != 0 {
// 		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		flag := inventoryManager.HasEnoughItems(useItemMap)
// 		if !flag {
// 			log.WithFields(log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("bodyShiled:当前金甲丹药数量不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}

// 		//消耗物品
// 		reasonText := commonlog.InventoryLogReasonShieldJJDan.String()
// 		flag = inventoryManager.BatchRemove(useItemMap, commonlog.InventoryLogReasonShieldJJDan, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("bodyShiled:BatchRemove should be ok"))
// 		}
// 		//同步物品
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}

// 	//培养判断
// 	pro, _, sucess := bodyshieldlogic.BodyShieldJinJiaDan(bodyShieldInfo.JinjiadanNum, bodyShieldInfo.JinjiadanPro, jinJiaDanTemplate)
// 	bshieldManager.EatJinJiaDan(pro, sucess)
// 	if sucess {
// 		//同步属性
// 		bodyShieldlogic.BodyShieldPropertyChanged(pl)
// 	}
// 	scBodyShieldJJDan := pbutil.BuildSCBodyShieldJJDan(bodyShieldInfo.JinjiadanLevel, bodyShieldInfo.JinjiadanPro)
// 	pl.SendMsg(scBodyShieldJJDan)
// 	return
// }

//食用护体盾金甲丹
func bodyShieldJJDan(pl player.Player, num int32) (err error) {
	bshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbshield.PlayerBodyShieldDataManager)
	bodyShieldInfo := bshieldManager.GetBodyShiedInfo()
	advancedId := bodyShieldInfo.AdvanceId
	jinJiaLevel := bodyShieldInfo.JinjiadanLevel
	bodyShieldTemplate := bodyshield.GetBodyShieldService().GetBodyShieldNumber(int32(advancedId))
	if bodyShieldTemplate == nil {
		return
	}
	jinJiaDanTemplate := bodyshield.GetBodyShieldService().GetBodyShieldJinJia(jinJiaLevel + 1)
	if jinJiaDanTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bodyShiled:护体金甲丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.BodyShieldEatJJDanReachedFull)
		return
	}

	if jinJiaLevel >= bodyShieldTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bodyShiled:护体金甲丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.BodyShieldEatJJDanReachedLimit)
		return
	}

	reachJinJiaDanTemplate, flag := bodyshield.GetBodyShieldService().GetBodyShieldEatJinJiaTemplate(jinJiaLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("mount:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachJinJiaDanTemplate.Level > bodyShieldTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("mount:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := jinJiaDanTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("bodyShiled:当前金甲丹数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonShieldJJDan.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonShieldJJDan, reasonText)
		if !flag {
			panic(fmt.Errorf("bodyShiled:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	bshieldManager.EatJinJiaDan(reachJinJiaDanTemplate.Level)
	//同步属性
	bodyShieldlogic.BodyShieldPropertyChanged(pl)

	scBodyShieldJJDan := pbutil.BuildSCBodyShieldJJDan(bodyShieldInfo.JinjiadanLevel, bodyShieldInfo.JinjiadanPro)
	pl.SendMsg(scBodyShieldJJDan)
	return
}
