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
	minggelogic "fgame/fgame/game/mingge/logic"
	"fgame/fgame/game/mingge/pbutil"
	playermingge "fgame/fgame/game/mingge/player"
	minggetypes "fgame/fgame/game/mingge/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MINGGE_PAN_UNLOAD_TYPE), dispatch.HandlerFunc(handleMingGePanUnload))
}

//处理命盘命格卸下信息
func handleMingGePanUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("mingge:处理命盘命格卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMingGePanUnload := msg.(*uipb.CSMingGePanUnload)
	panType := csMingGePanUnload.GetPanType()
	mingGeType := csMingGePanUnload.GetMingGeType()
	slot := csMingGePanUnload.GetSlot()

	err = mingGePanUnload(tpl, minggetypes.MingGeType(panType), minggetypes.MingGeAllSubType(mingGeType), minggetypes.MingGeSlotType(slot))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("mingge:处理命盘命格卸下信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mingge:处理命盘命格卸下信息完成")
	return nil
}

//处理命盘命格卸下信息逻辑
func mingGePanUnload(pl player.Player, mingGeType minggetypes.MingGeType, mingGeSubType minggetypes.MingGeAllSubType, slot minggetypes.MingGeSlotType) (err error) {
	if !mingGeType.Valid() || !mingGeSubType.Valid() {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"mingGeType":    mingGeType,
			"mingGeSubType": mingGeSubType,
			"slot":          slot,
		}).Warn("mingge:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !slot.Vaild() {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"mingGeType":    mingGeType,
			"mingGeSubType": mingGeSubType,
			"slot":          slot,
		}).Warn("mingge:槽位不存在")
		playerlogic.SendSystemMessage(pl, lang.MingGeMosaicItemSlotNoCorrect)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	curItemId := manager.GetSlotItem(mingGeType, mingGeSubType, slot)
	if curItemId == 0 {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"mingGeType":    mingGeType,
			"mingGeSubType": mingGeSubType,
			"slot":          slot,
		}).Warn("mingge:槽位没有镶嵌命格")
		playerlogic.SendSystemMessage(pl, lang.MingGeSlotUnloadNoItemId)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//卸下当前命格物品
	if curItemId != 0 {
		inventoryReason := commonlog.InventoryLogReasonMingGeUnloadAdd
		reasonText := inventoryReason.String()
		flag := inventoryManager.AddItem(curItemId, 1, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("mingge: mingGeMosaic UseItem should be ok"))
		}
		//推送背包变化
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag := manager.SlotUnload(mingGeType, mingGeSubType, slot)
	if !flag {
		panic("mingge: SlotUnload  should be ok")
	}

	// 属性变化
	minggelogic.MingGePropertyChanged(pl)
	scMingGeUnload := pbutil.BuildSCMingGeUnload(int32(mingGeType), int32(mingGeSubType), int32(slot))
	pl.SendMsg(scMingGeUnload)
	return
}
