package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	mountlogic "fgame/fgame/game/mount/logic"
	"fgame/fgame/game/mount/pbutil"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeMount, playerinventory.ItemUseHandleFunc(handleMountAdvancedTicket))
}

//坐骑进阶丹
func handleMountAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountInfo := mountManager.GetMountInfo()
	curAdvancedId := int32(mountInfo.AdvanceId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("mount:进阶丹使用,坐骑阶数条件不符")
		playerlogic.SendSystemMessage(pl, lang.MountAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	mountManager.MountAdvancedTicket(addAdvancedNum)

	mountReason := commonlog.MountLogReasonAdvanced
	reasonText := fmt.Sprintf(mountReason.String(), commontypes.AdvancedTypeTicket.String())
	data := mounteventtypes.CreatePlayerMountAdvancedLogEventData(curAdvancedId, addAdvancedNum, mountReason, reasonText)
	gameevent.Emit(mounteventtypes.EventTypeMountAdvancedLog, pl, data)

	//同步属性
	mountlogic.MountPropertyChanged(pl)
	scMountAdvanced := pbutil.BuildSCMountAdavancedFinshed(int32(mountInfo.AdvanceId), mountInfo.MountId, commontypes.AdvancedTypeTicket)
	pl.SendMsg(scMountAdvanced)

	flag = true
	return
}
