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
	mountlogic "fgame/fgame/game/mount/logic"
	"fgame/fgame/game/mount/mount"
	"fgame/fgame/game/mount/pbutil"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOUNT_UNREALDAN_TYPE), dispatch.HandlerFunc(handleMountUnrealDan))
}

//处理坐骑食幻化丹信息
func handleMountUnrealDan(s session.Session, msg interface{}) (err error) {
	log.Debug("mount:处理坐骑食幻化丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMountUnrealDan := msg.(*uipb.CSMountUnrealDan)
	num := csMountUnrealDan.GetNum()

	err = mountUnrealDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("mount:处理坐骑食幻化丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mount:处理坐骑食幻化丹信息完成")
	return nil

}

// //坐骑食幻化丹的逻辑
// func mountUnrealDan(pl player.Player) (err error) {

// 	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
// 	mountInfo := mountManager.GetMountInfo()
// 	advancedId := mountInfo.AdvanceId
// 	unrealLevel := mountInfo.UnrealLevel
// 	mountTemplate := mount.GetMountService().GetMountNumber(int32(advancedId))
// 	if mountTemplate == nil {
// 		return
// 	}
// 	hunaHuaTemplate := mount.GetMountService().GetMountHuanHuaTemplate(unrealLevel + 1)
// 	if hunaHuaTemplate == nil {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Warn("mount:幻化丹食丹等级满级")
// 		playerlogic.SendSystemMessage(pl, lang.MountEatUnDanReachedFull)
// 		return
// 	}

// 	if unrealLevel >= mountTemplate.ShidanLimit {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Warn("mount:幻化丹食丹等级已达最大,请进阶后再试")
// 		playerlogic.SendSystemMessage(pl, lang.MountEatUnDanReachedLimit)
// 		return
// 	}

// 	useItemMap := hunaHuaTemplate.GetUseItemTemplate()
// 	if len(useItemMap) != 0 {
// 		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		flag := inventoryManager.HasEnoughItems(useItemMap)
// 		if !flag {
// 			log.WithFields(log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("mount:当前幻化丹药数量不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}
// 		//消耗物品
// 		reasonText := commonlog.InventoryLogReasonMountEatUn.String()
// 		flag = inventoryManager.BatchRemove(useItemMap, commonlog.InventoryLogReasonMountEatUn, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("mount:BatchRemove should be ok"))
// 		}
// 		//同步物品
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}

// 	//坐骑幻化丹培养判断
// 	pro, _, sucess := mountlogic.MountHunaHuaFeed(mountInfo.UnrealNum, mountInfo.UnrealPro, hunaHuaTemplate)
// 	mountManager.EatUnrealDan(pro, sucess)
// 	if sucess {
// 		//同步属性
// 		mountlogic.MountPropertyChanged(pl)
// 	}

// 	scMountUnrealDan := pbutil.BuildSCMountUnrealDan(mountInfo.UnrealLevel, mountInfo.UnrealPro)
// 	pl.SendMsg(scMountUnrealDan)
// 	return
// }

//坐骑食幻化丹的逻辑
func mountUnrealDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("mount:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountInfo := mountManager.GetMountInfo()
	advancedId := mountInfo.AdvanceId
	unrealLevel := mountInfo.UnrealLevel
	mountTemplate := mount.GetMountService().GetMountNumber(int32(advancedId))
	if mountTemplate == nil {
		return
	}
	hunaHuaTemplate := mount.GetMountService().GetMountHuanHuaTemplate(unrealLevel + 1)
	if hunaHuaTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("mount:幻化丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.MountEatUnDanReachedFull)
		return
	}

	if unrealLevel >= mountTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("mount:幻化丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.MountEatUnDanReachedLimit)
		return
	}

	reachHuanHuaTemplate, flag := mount.GetMountService().GetMountEatHuanHuanTemplate(unrealLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("mount:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachHuanHuaTemplate.Level > mountTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("mount:参数错误")
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
			}).Warn("mount:当前幻化丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonMountEatUn.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonMountEatUn, reasonText)
		if !flag {
			panic(fmt.Errorf("mount:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	mountManager.EatUnrealDan(reachHuanHuaTemplate.Level)
	//同步属性
	mountlogic.MountPropertyChanged(pl)

	scMountUnrealDan := pbutil.BuildSCMountUnrealDan(mountInfo.UnrealLevel, mountInfo.UnrealPro)
	pl.SendMsg(scMountUnrealDan)
	return
}
