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
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOUNT_UNREAL_TYPE), dispatch.HandlerFunc(handleMountUnreal))
}

//处理坐骑幻化信息
func handleMountUnreal(s session.Session, msg interface{}) (err error) {
	log.Debug("mount:处理坐骑幻化信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMountUnreal := msg.(*uipb.CSMountUnreal)
	mountId := csMountUnreal.GetMountId()
	if mountId <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mountId":  mountId,
			}).Debug("mount:处理坐骑幻化信息,无效mountId")
		return nil
	}
	err = mountUnreal(tpl, int(mountId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mountId":  mountId,
				"error":    err,
			}).Error("mount:处理坐骑幻化信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"mountId":  mountId,
		}).Debug("mount:处理坐骑幻化信息完成")
	return nil

}

//坐骑幻化的逻辑
func mountUnreal(pl player.Player, mountId int) (err error) {
	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountTemplate := mount.GetMountService().GetMount(int(mountId))
	//校验参数
	if mountTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"mountId":  mountId,
		}).Warn("mount:幻化mountId无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}


	//是否已幻化
	flag := mountManager.IsUnrealed(mountId)
	if !flag {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//获取物品幻化条件3,消耗物品数
		paramItemsMap := mountTemplate.GetMagicParamIMap()
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.HasEnoughItems(paramItemsMap)
			if !flag {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"mountId":  mountId,
				}).Warn("mount:还有幻化条件未达成，无法解锁幻化")
				playerlogic.SendSystemMessage(pl, lang.MountUnrealCondNotReached)
				return
			}

		}

		flag = mountManager.IsCanUnreal(mountId)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"mountId":  mountId,
			}).Warn("mount:还有条件未达成，无法解锁幻化")
			playerlogic.SendSystemMessage(pl, lang.MountUnrealCondNotReached)
			return
		}

		//使用幻化条件3的物品
		inventoryReason := commonlog.InventoryLogReasonMountUnreal
		reasonText := fmt.Sprintf(inventoryReason.String(), mountId)
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.BatchRemove(paramItemsMap, inventoryReason, reasonText)
			if !flag {
				panic(fmt.Errorf("mount:use item should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
		mountManager.AddUnrealInfo(mountId)
		mountlogic.MountPropertyChanged(pl)
	}
	flag = mountManager.Unreal(mountId)
	if !flag {
		panic(fmt.Errorf("mount:幻化应该成功"))
	}
	scMountUnreal := pbutil.BuildSCMountUnreal(int32(mountId))
	pl.SendMsg(scMountUnreal)
	return
}
