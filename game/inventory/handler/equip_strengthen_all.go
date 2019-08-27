package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/inventory/pbutil"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_EQUIP_STRENGTHEN_ALL_TYPE), dispatch.HandlerFunc(handleInventoryEquipStrengthenAll))
}

//处理一键强化
func handleInventoryEquipStrengthenAll(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理槽位强化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = equipStrengthenAll(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理槽位强化装备,错误")

		return err
	}
	log.Debug("inventory:处理槽位强化装备,完成")
	return nil
}

//强化所有
func equipStrengthenAll(pl player.Player) (err error) {
	resultList := equipmentSlotStrengthenUpgradeAll(pl)
	if len(resultList) == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("inventory:处理槽位强化装备,没有可以强化的槽位")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotStrengthNoSlot)
		return
	}
	//同步改变
	logic.SnapInventoryEquipChanged(pl)
	logic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//强化成功
	scInventoryEquipmentSlotStrengthAll := pbutil.BuildSCInventoryEquipmentSlotStrengthAll(resultList)
	pl.SendMsg(scInventoryEquipmentSlotStrengthAll)

	return
}

//强化所有
func equipmentSlotStrengthenUpgradeAll(pl player.Player) (resultList []*inventorytypes.StrengthenResult) {
	for pos := inventorytypes.BodyPositionTypeWeapon; pos <= inventorytypes.BodyPositionTypeRing; pos++ {
		tempResult, flag := equipmentSlotStrengthenUpgrade(pl, pos, true)
		if !flag {
			continue
		}
		result := &inventorytypes.StrengthenResult{
			Pos:    pos,
			Result: tempResult,
		}
		resultList = append(resultList, result)
	}
	return resultList
}
