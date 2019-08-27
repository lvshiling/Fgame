package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commonlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	mountlogic "fgame/fgame/game/mount/logic"
	"fgame/fgame/game/mount/mount"
	"fgame/fgame/game/mount/pbutil"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeMount, playerinventory.ItemUseHandleFunc(handleMountBlessDan))
}

//坐骑祝福丹
func handleMountBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountInfo := mountManager.GetMountInfo()
	mountAdvancedId := int32(mountInfo.AdvanceId)
	needAdvanceId := itemTemplate.TypeFlag1
	if mountAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"mountAdvancedId": mountAdvancedId,
				"needAdvanceId":   itemTemplate.TypeFlag1,
			}).Warn("mount:祝福丹使用,坐骑阶数不足")
		playerlogic.SendSystemMessage(pl, lang.MountAdvanceNotEqual)
		return
	}

	nextMountTemp := mount.GetMountService().GetMountNumber(mountAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := mountInfo.TimesNum
	curTimes += addTimes
	// isAdvanced, pro := mountlogic.MountEatZhuFuDan(pl, curTimes, mountInfo.Bless, addBlessNum, nextMountTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(mountInfo.Bless, addBlessNum, nextMountTemp.ZhufuMax)
	mountManager.MountAdvanced(pro, addTimes, isAdvanced)

	if isAdvanced {
		mountReason := commonlog.MountLogReasonAdvanced
		reasonText := fmt.Sprintf(mountReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := mounteventtypes.CreatePlayerMountAdvancedLogEventData(mountAdvancedId, 1, mountReason, reasonText)
		gameevent.Emit(mounteventtypes.EventTypeMountAdvancedLog, pl, data)

		//同步属性
		mountlogic.MountPropertyChanged(pl)
		scMountAdvanced := pbutil.BuildSCMountAdavancedFinshed(int32(mountInfo.AdvanceId), mountInfo.MountId, commontypes.AdvancedTypeBlessDan)
		pl.SendMsg(scMountAdvanced)
	} else {
		//进阶不成功
		scMountAdvanced := pbutil.BuildSCMountAdavanced(int32(mountInfo.AdvanceId), mountInfo.MountId, addBlessNum, mountInfo.Bless, mountInfo.BlessTime, commontypes.AdvancedTypeBlessDan, false)
		pl.SendMsg(scMountAdvanced)
	}

	flag = true
	return
}
