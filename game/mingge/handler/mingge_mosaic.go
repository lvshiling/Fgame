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
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	minggelogic "fgame/fgame/game/mingge/logic"
	"fgame/fgame/game/mingge/pbutil"
	playermingge "fgame/fgame/game/mingge/player"
	minggetemplate "fgame/fgame/game/mingge/template"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_MINGGE_PAN_MOSAIC_TYPE), dispatch.HandlerFunc(handleMingGeMosaic))
}

//处理命盘洗练信息
func handleMingGeMosaic(s session.Session, msg interface{}) (err error) {
	log.Debug("mingge:处理命盘洗练信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMingGePanMosaic := msg.(*uipb.CSMingGePanMosaic)
	panType := csMingGePanMosaic.GetPanType()
	mingGeType := csMingGePanMosaic.GetMingGeType()
	slot := csMingGePanMosaic.GetSlot()
	itemId := csMingGePanMosaic.GetItemId()

	err = mingGeMosaic(tpl, minggetypes.MingGeType(panType), minggetypes.MingGeAllSubType(mingGeType), minggetypes.MingGeSlotType(slot), itemId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("mingge:处理命盘洗练信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mingge:处理命盘洗练信息完成")
	return nil
}

//处理命盘洗练信息逻辑
func mingGeMosaic(pl player.Player, mingGeType minggetypes.MingGeType, mingGeSubType minggetypes.MingGeAllSubType, slot minggetypes.MingGeSlotType, itemId int32) (err error) {
	if !mingGeType.Valid() || !mingGeSubType.Valid() {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"mingGeType":    mingGeType,
			"mingGeSubType": mingGeSubType,
			"slot":          slot,
			"itemId":        itemId,
		}).Warn("mingge:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil || itemTemplate.GetItemType() != itemtypes.ItemTypeMingGe {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"mingGeType":    mingGeType,
			"mingGeSubType": mingGeSubType,
			"slot":          slot,
			"itemId":        itemId,
		}).Warn("mingge:物品不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	typeFlag1 := itemTemplate.TypeFlag1
	mingGeTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeTempalte(typeFlag1)
	if mingGeTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"mingGeType":    mingGeType,
			"mingGeSubType": mingGeSubType,
			"slot":          slot,
			"itemId":        itemId,
		}).Warn("mingge:命格物品不合适")
		playerlogic.SendSystemMessage(pl, lang.MingGeMosaicItemNoCorrect)
		return
	}
	if !slot.Vaild() {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"mingGeType":    mingGeType,
			"mingGeSubType": mingGeSubType,
			"slot":          slot,
			"itemId":        itemId,
		}).Warn("mingge:槽位不存在")
		playerlogic.SendSystemMessage(pl, lang.MingGeMosaicItemSlotNoCorrect)
		return
	}

	//普通可以镶嵌超级的
	if mingGeType == minggetypes.MingGeTypeNormal {
		if mingGeTemplate.GetMingGeSubType() != mingGeSubType {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"mingGeType":    mingGeType,
				"mingGeSubType": mingGeSubType,
				"slot":          slot,
				"itemId":        itemId,
			}).Warn("mingge:命格物品不合适")
			playerlogic.SendSystemMessage(pl, lang.MingGeMosaicItemNoCorrect)
			return
		}
	} else {
		if mingGeTemplate.GetMingGeType() != mingGeType || mingGeTemplate.GetMingGeSubType() != mingGeSubType {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"mingGeType":    mingGeType,
				"mingGeSubType": mingGeSubType,
				"slot":          slot,
				"itemId":        itemId,
			}).Warn("mingge:命格物品不合适")
			playerlogic.SendSystemMessage(pl, lang.MingGeMosaicItemNoCorrect)
			return
		}
	}

	//判断背包物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItem(itemId, 1) {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"mingGeType":    mingGeType,
			"mingGeSubType": mingGeSubType,
			"slot":          slot,
			"itemId":        itemId,
		}).Warn("mingge:命格物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	curItemId := manager.GetSlotItem(mingGeType, mingGeSubType, slot)
	if curItemId != 0 {
		curItemTemplate := item.GetItemService().GetItem(int(curItemId))
		if curItemTemplate.GetQualityType() > itemTemplate.GetQualityType() ||
			(curItemTemplate.GetQualityType() == itemTemplate.GetQualityType() && curItemTemplate.TypeFlag2 >= itemTemplate.TypeFlag2) {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"mingGeType":    mingGeType,
				"mingGeSubType": mingGeSubType,
				"slot":          slot,
				"itemId":        itemId,
			}).Warn("mingge:命格品质不高于当前镶嵌的命格")
			playerlogic.SendSystemMessage(pl, lang.MingGeMosaicItemQualityNoHigher)
			return
		}
	}

	//使用物品
	inventoryReason := commonlog.InventoryLogReasonMingGeMosaicUse
	reasonText := inventoryReason.String()
	flag := inventoryManager.UseItem(itemId, 1, inventoryReason, reasonText)
	if !flag {
		panic(fmt.Errorf("mingge: mingGeMosaic UseItem should be ok"))
	}
	flag = manager.SlotMosaic(mingGeType, mingGeSubType, slot, itemId)
	if !flag {
		panic("mingge: mingGeMosaic use item should be ok")
	}

	//卸下当前命格物品
	if curItemId != 0 {
		inventoryReason := commonlog.InventoryLogReasonMingGeUnloadAdd
		reasonText := inventoryReason.String()
		flag := inventoryManager.AddItem(curItemId, 1, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("mingge: mingGeMosaic UseItem should be ok"))
		}
	}

	//推送背包变化
	inventorylogic.SnapInventoryChanged(pl)

	// 属性变化
	minggelogic.MingGePropertyChanged(pl)
	scMingGeMosaic := pbutil.BuildSCMingGeMosaic(int32(mingGeType), int32(mingGeSubType), int32(slot), itemId)
	pl.SendMsg(scMingGeMosaic)
	return
}
