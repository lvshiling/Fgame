package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/inventory/logic"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_EQUIP_UPGRADE_TYPE), dispatch.HandlerFunc(handleInventoryEquipUpgrade))
}

//处理升阶
func handleInventoryEquipUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理装备升阶")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryEquipUpgrade := msg.(*uipb.CSInventoryEquipUpgrade)
	slotId := csInventoryEquipUpgrade.GetSlotId()
	slotPosition := inventorytypes.BodyPositionType(slotId)
	//参数不对
	if !slotPosition.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      slotPosition.String(),
			}).Warn("inventory:强化升阶,参数不对")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = equipSlotUpgrade(tpl, slotPosition)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理装备升阶,错误")

		return err
	}
	log.Debug("inventory:处理装备升阶,完成")
	return nil
}

//升阶
func equipSlotUpgrade(pl player.Player, pos inventorytypes.BodyPositionType) (err error) {
	logic.EquipSlotUpgrade(pl, pos)
	return
}
