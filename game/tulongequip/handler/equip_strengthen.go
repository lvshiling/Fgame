package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	tulongequipeventtypes "fgame/fgame/game/tulongequip/event/types"
	tulongequiplogic "fgame/fgame/game/tulongequip/logic"
	"fgame/fgame/game/tulongequip/pbutil"
	playertulongequip "fgame/fgame/game/tulongequip/player"
	tulongequiptemplate "fgame/fgame/game/tulongequip/template"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TULONG_EQUIP_STRENGTHEN_TYPE), dispatch.HandlerFunc(handleTuLongEquipStrengthen))
}

//处理强化
func handleTuLongEquipStrengthen(s session.Session, msg interface{}) (err error) {
	log.Debug("tulongequip:处理装备槽屠龙强化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTuLongEquipStrengthen)
	slotId := csMsg.GetSlotId()
	suitInt := csMsg.GetSuitType()

	slotPosition := inventorytypes.BodyPositionType(slotId)
	if !slotPosition.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	suitType := tulongequiptypes.TuLongSuitType(suitInt)
	if !suitType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = tuLongEquipStrengthen(tpl, suitType, slotPosition)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"slotPosition": slotPosition,
				"suitType":     suitType,
				"error":        err,
			}).Error("tulongequip:处理装备槽屠龙装备强化,错误")

		return err
	}
	log.Debug("tulongequip:处理装备槽屠龙装备强化,完成")
	return nil
}

//强化
func tuLongEquipStrengthen(pl player.Player, suitType tulongequiptypes.TuLongSuitType, posType inventorytypes.BodyPositionType) (err error) {

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	tulongequipManager := pl.GetPlayerDataManager(playertypes.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	targetIt := tulongequipManager.GetTuLongEquipByPos(suitType, posType)

	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitType": suitType,
				"posType":  posType,
			}).Warn("tulongequip:强化升级失败,强化目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	curLevel := targetIt.GetLevel()

	//能否被强化
	strengthenTemplate := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipStrengthenTemplate(suitType, posType, curLevel)
	if strengthenTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitType": suitType,
				"posType":  posType,
				"curLevel": curLevel,
			}).Warn("tulongequip:强化升级失败,该屠龙装备无法被强化")
		playerlogic.SendSystemMessage(pl, lang.TuLongEquipStrengthenNotAllow)
		return
	}

	//强化等级判断
	if strengthenTemplate.NextId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitType": suitType,
				"posType":  posType,
			}).Warn("tulongequip:强化升级失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.TuLongEquipReachStrengthenMax)
		return
	}

	nextStrengthenTemplate := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipStrengthenTemplate(suitType, posType, curLevel+1)
	needItemId := nextStrengthenTemplate.NeedItem
	needItemNum := nextStrengthenTemplate.ItemCount
	if !inventoryManager.HasEnoughItem(needItemId, needItemNum) {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"needItemId":  needItemId,
				"needItemNum": needItemNum,
			}).Warn("tulongequip:强化升级失败,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//消耗材料
	if needItemNum > 0 {
		reason := commonlog.InventoryLogReasonTuLongEquipStrengthenUse
		flag := inventoryManager.UseItem(needItemId, needItemNum, reason, reason.String())
		if !flag {
			panic(fmt.Errorf("tulongequip:背包强化升级移除材料应该成功"))
		}
	}
	eventdata := tulongequipeventtypes.CreatePlayerTuLongEquipUseItemEventData(needItemId, nextStrengthenTemplate.ItemCount)
	gameevent.Emit(tulongequipeventtypes.EventTypeTuLongEquipUseItem, pl, eventdata)

	//计算成功
	success := mathutils.RandomHit(common.MAX_RATE, int(strengthenTemplate.Rate))
	if success {
		flag := tulongequipManager.UpdateTuLongEquipLevel(suitType, posType)
		if !flag {
			panic(fmt.Errorf("tulongequip: 强化升级应该成功"))
		}

		tulongequiplogic.TuLongEquipPropertyChanged(pl)
		tulongequiplogic.SnapInventoryTuLongEquipChanged(pl)
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCTuLongEquipStrengthen(int32(suitType), int32(posType), targetIt.GetLevel(), success)
	pl.SendMsg(scMsg)

	return
}
