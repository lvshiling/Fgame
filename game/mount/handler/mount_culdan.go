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

	processor.Register(codec.MessageType(uipb.MessageType_CS_MOUNT_CULDAN_TYPE), dispatch.HandlerFunc(handleMountCulDan))

}

//处理坐骑食培养丹信息
func handleMountCulDan(s session.Session, msg interface{}) (err error) {
	log.Debug("mount:处理坐骑食培养丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMountCulDan := msg.(*uipb.CSMountCulDan)
	num := csMountCulDan.GetNum()

	err = mountCulDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("mount:处理坐骑食培养丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mount:处理坐骑食培养丹信息完成")
	return nil

}

//坐骑食培养丹逻辑
// func mountCulDan(pl player.Player) (err error) {
// 	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
// 	mountInfo := mountManager.GetMountInfo()
// 	advancedId := mountInfo.AdvanceId
// 	culLevel := mountInfo.CulLevel
// 	mountTemplate := mount.GetMountService().GetMountNumber(int32(advancedId))
// 	if mountTemplate == nil {
// 		return
// 	}
// 	caoLiaoTemplate := mount.GetMountService().GetMountCaoLiaoTemplate(culLevel + 1)
// 	if caoLiaoTemplate == nil {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Warn("mount:培养丹食丹等级满级")
// 		playerlogic.SendSystemMessage(pl, lang.MountEatCulDanReachedFull)
// 		return
// 	}

// 	if culLevel >= mountTemplate.CulturingDanLimit {
// 		log.WithFields(log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Warn("mount:培养丹食丹等级已达最大,请进阶后再试")
// 		playerlogic.SendSystemMessage(pl, lang.MountEatCulDanReachedLimit)
// 		return
// 	}

// 	useItemMap := caoLiaoTemplate.GetUseItemTemplate()
// 	if len(useItemMap) != 0 {
// 		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		flag := inventoryManager.HasEnoughItems(useItemMap)
// 		if !flag {
// 			log.WithFields(log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warn("mount:当前培养丹数量不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}
// 		//消耗物品
// 		reasonText := commonlog.InventoryLogReasonMountEatClu.String()
// 		flag = inventoryManager.BatchRemove(useItemMap, commonlog.InventoryLogReasonMountEatClu, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("mount:BatchRemove should be ok"))
// 		}
// 		//同步物品
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}

// 	//坐骑草料喂养判断
// 	pro, _, sucess := mountlogic.MountCaoLiaoFeed(mountInfo.CulNum, mountInfo.CulPro, caoLiaoTemplate)
// 	mountManager.EatCulDan(pro, sucess)
// 	if sucess {
// 		//同步属性
// 		mountlogic.MountPropertyChanged(pl)
// 	}
// 	scMountShiDan := pbutil.BuildSCMountCulDan(mountInfo.CulLevel, mountInfo.CulPro)
// 	pl.SendMsg(scMountShiDan)
// 	return
// }

//坐骑食培养丹逻辑
func mountCulDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("mount:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountInfo := mountManager.GetMountInfo()
	advancedId := mountInfo.AdvanceId
	culLevel := mountInfo.CulLevel
	mountTemplate := mount.GetMountService().GetMountNumber(int32(advancedId))
	if mountTemplate == nil {
		return
	}
	caoLiaoTemplate := mount.GetMountService().GetMountCaoLiaoTemplate(culLevel + 1)
	if caoLiaoTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("mount:培养丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.MountEatCulDanReachedFull)
		return
	}

	if culLevel >= mountTemplate.CulturingDanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("mount:培养丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.MountEatCulDanReachedLimit)
		return
	}

	reachCaoLiaoTemplate, flag := mount.GetMountService().GetMountEatCaoLiaoTemplate(culLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("mount:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachCaoLiaoTemplate.Level > mountTemplate.CulturingDanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("mount:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := caoLiaoTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("mount:当前培养丹数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonMountEatClu.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonMountEatClu, reasonText)
		if !flag {
			panic(fmt.Errorf("mount:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	mountManager.EatCulDan(reachCaoLiaoTemplate.Level)
	mountlogic.MountPropertyChanged(pl)

	scMountShiDan := pbutil.BuildSCMountCulDan(mountInfo.CulLevel, mountInfo.CulPro)
	pl.SendMsg(scMountShiDan)
	return
}
