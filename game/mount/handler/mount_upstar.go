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
	processor.Register(codec.MessageType(uipb.MessageType_CS_MOUNT_UPSTAR_TYPE), dispatch.HandlerFunc(handleMountUpstar))
}

//处理坐骑皮肤升星信息
func handleMountUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("mount:处理坐骑皮肤升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMountUpstar := msg.(*uipb.CSMountUpstar)
	mountId := csMountUpstar.GetMountId()

	err = mountUpstar(tpl, mountId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mountId":  mountId,
				"error":    err,
			}).Error("mount:处理坐骑皮肤升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mount:处理坐骑皮肤升星完成")
	return nil
}

//坐骑皮肤升星的逻辑
func mountUpstar(pl player.Player, mountId int32) (err error) {
	mountTemplate := mount.GetMountService().GetMount(int(mountId))
	if mountTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"mountId":  mountId,
		}).Warn("mount:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if mountTemplate.MountUpstarBeginId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"mountId":  mountId,
		}).Warn("mount:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountOtherInfo, flag := mountManager.IfMountSkinExist(mountId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"mountId":  mountId,
		}).Warn("mount:未激活的坐骑皮肤,无法升星")
		playerlogic.SendSystemMessage(pl, lang.MountSkinUpstarNoActive)
		return
	}

	_, flag = mountManager.IfCanUpStar(mountId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"mountId":  mountId,
		}).Warn("mount:坐骑皮肤已满星")
		playerlogic.SendSystemMessage(pl, lang.MountSkinReacheFullStar)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	level := mountOtherInfo.Level
	nextLevel := level + 1
	to := mount.GetMountService().GetMount(int(mountId))
	if to == nil {
		return
	}
	mountUpstarTemplate := to.GetMountUpstarByLevel(nextLevel)
	if mountUpstarTemplate == nil {
		return
	}

	needItems := mountUpstarTemplate.GetNeedItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"mountId":  mountId,
			}).Warn("mount:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := commonlog.InventoryLogReasonMountUpstar.String()
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonMountUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("mount: mountUpstar use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//坐骑皮肤升星判断
	pro, _, sucess := mountlogic.MountSkinUpstar(mountOtherInfo.UpNum, mountOtherInfo.UpPro, mountUpstarTemplate)
	flag = mountManager.Upstar(mountId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("mount: mountUpstar should be ok"))
	}
	if sucess {
		mountlogic.MountPropertyChanged(pl)
	}
	scMountUpstar := pbutil.BuildSCMountUpstar(mountId, mountOtherInfo.Level, mountOtherInfo.UpPro)
	pl.SendMsg(scMountUpstar)
	return
}
