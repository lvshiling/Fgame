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
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_TAKE_OFF_EQUIP_TYPE), dispatch.HandlerFunc(handleInventoryTakeOffEquip))
}

//处理脱下
func handleInventoryTakeOffEquip(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理脱下装备")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryTakeOffEquip := msg.(*uipb.CSInventoryTakeOffEquip)
	slotId := csInventoryTakeOffEquip.GetSlotId()
	slotPosition := inventorytypes.BodyPositionType(slotId)
	//参数不对
	if !slotPosition.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = takeOff(tpl, slotPosition)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理脱下装备,错误")

		return err
	}
	log.Debug("inventory:处理脱下装备,完成")
	return nil
}

//脱下
func takeOff(pl player.Player, pos inventorytypes.BodyPositionType) (err error) {
	return logic.TakeOff(pl, pos)

}
