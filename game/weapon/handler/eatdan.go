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
	weaponlogic "fgame/fgame/game/weapon/logic"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"
	"fgame/fgame/game/weapon/weapon"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WEAPON_EATDAN_TYPE), dispatch.HandlerFunc(handleWeaponEatDan))
}

//处理兵魂食培养丹信息
func handleWeaponEatDan(s session.Session, msg interface{}) (err error) {
	log.Debug("weapon:处理兵魂食培养丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWeaponEatDan := msg.(*uipb.CSWeaponEatDan)
	weaponId := csWeaponEatDan.GetWeaponId()
	num := csWeaponEatDan.GetNum()

	err = weaponEatDan(tpl, weaponId, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weaponId": weaponId,
				"num":      num,
				"error":    err,
			}).Error("weapon:处理兵魂食培养丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("weapon:处理兵魂食培养丹信息完成")
	return nil

}

// //兵魂食培养丹逻辑
// func weaponEatDan(pl player.Player, weaponId int32) (err error) {
// 	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
// 	if !weaponManager.IsValid(weaponId) {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 			"weaponId": weaponId,
// 		}).Warn("weapon:参数无效")
// 		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	flag := weaponManager.IfWeaponExist(weaponId)
// 	if !flag {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 			"weaponId": weaponId,
// 		}).Warn("weapon:未激活的兵魂,无法食培养丹")
// 		playerlogic.SendSystemMessage(pl, lang.WeaponNotActiveNotEat)
// 		return
// 	}

// 	flag = weaponManager.IfCanPeiYang(weaponId)
// 	if !flag {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 			"weaponId": weaponId,
// 		}).Warn("weapon:食丹等级已达最大")
// 		playerlogic.SendSystemMessage(pl, lang.WeaponEatDanReachedLimit)
// 		return
// 	}
// 	weaponInfo := weaponManager.GetWeapon(weaponId)
// 	culLevel := weaponInfo.CulLevel
// 	weaponTemplate := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
// 	if culLevel >= weaponTemplate.EatDan {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Warn("weapon:食丹等级已达最大")
// 		playerlogic.SendSystemMessage(pl, lang.WeaponEatDanReachedLimit)
// 		return
// 	}

// 	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
// 	culTemplate := to.GetWeaponPeiYangByLevel(culLevel + 1)
// 	if culTemplate == nil {
// 		return
// 	}

// 	useItemMap := culTemplate.GetUseItemTemplate()
// 	if len(useItemMap) != 0 {
// 		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		flag := inventoryManager.HasEnoughItems(useItemMap)
// 		if !flag {
// 			log.WithFields(log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("weapon:当前培养丹数量不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}

// 		//消耗物品
// 		reasonText := commonlog.InventoryLogReasonWeaponEatDan.String()
// 		flag = inventoryManager.BatchRemove(useItemMap, commonlog.InventoryLogReasonWeaponEatDan, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("weapon:BatchRemove should be ok"))
// 		}
// 		//同步物品
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}
// 	//兵魂培养判断
// 	pro, _, sucess := weaponlogic.WeaponPeiYang(weaponInfo.CulNum, weaponInfo.CulPro, culTemplate)
// 	weaponManager.EatCulDan(weaponId, pro, sucess)
// 	if sucess {
// 		//同步属性
// 		weaponlogic.WeaponPropertyChanged(pl)
// 	}

// 	scWeaponEatDan := pbutil.BuildSCWeaponEatDan(weaponId, weaponInfo.CulLevel, weaponInfo.CulPro)
// 	pl.SendMsg(scWeaponEatDan)
// 	return
// }

//兵魂食培养丹逻辑
func weaponEatDan(pl player.Player, weaponId int32, num int32) (err error) {
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	if !weaponManager.IsValid(weaponId) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag := weaponManager.IfWeaponExist(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:未激活的兵魂,无法食培养丹")
		playerlogic.SendSystemMessage(pl, lang.WeaponNotActiveNotEat)
		return
	}

	flag = weaponManager.IfCanPeiYang(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:食丹等级已达最大")
		playerlogic.SendSystemMessage(pl, lang.WeaponEatDanReachedLimit)
		return
	}
	weaponInfo := weaponManager.GetWeapon(weaponId)
	culLevel := weaponInfo.CulLevel
	weaponTemplate := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if culLevel >= weaponTemplate.EatDan {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
			"num":      num,
		}).Warn("weapon:食丹等级已达最大")
		playerlogic.SendSystemMessage(pl, lang.WeaponEatDanReachedLimit)
		return
	}

	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if to == nil {
		return
	}
	culTemplate := to.GetWeaponPeiYangByLevel(culLevel + 1)
	if culTemplate == nil {
		return
	}

	reachPeiYangTemplate, flag := to.GetWeaponEatPeiYangTemplate(culLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
			"num":      num,
		}).Warn("weapon:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachPeiYangTemplate.Level > weaponTemplate.EatDan {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
			"num":      num,
		}).Warn("weapon:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := culTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("weapon:当前幻化丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonWeaponEatDan.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonWeaponEatDan, reasonText)
		if !flag {
			panic(fmt.Errorf("weapon:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	weaponManager.EatCulDan(weaponId, reachPeiYangTemplate.Level)
	//同步属性
	weaponlogic.WeaponPropertyChanged(pl)

	scWeaponEatDan := pbutil.BuildSCWeaponEatDan(weaponId, weaponInfo.CulLevel, weaponInfo.CulPro)
	pl.SendMsg(scWeaponEatDan)
	return
}
