package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	tulongequiplogic "fgame/fgame/game/tulongequip/logic"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TULONG_TAKE_OFF_EQUIP_TYPE), dispatch.HandlerFunc(handleTakeOffTuLongEquip))
}

//处理脱下屠龙装
func handleTakeOffTuLongEquip(s session.Session, msg interface{}) (err error) {
	log.Debug("tulongequip:处理脱下屠龙装")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTuLongTakeOffEquip)
	slotId := csMsg.GetSlotId()
	suitInt := csMsg.GetSuitType()

	//参数不对
	slotPosition := inventorytypes.BodyPositionType(slotId)
	if !slotPosition.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	suitType := tulongequiptypes.TuLongSuitType(suitInt)
	if !suitType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"slotId":    slotId,
			}).Warn("inventory:使用元神屠龙装备,装备类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = takeOff(tpl, suitType, slotPosition)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("tulongequip:处理脱下屠龙装,错误")

		return err
	}
	log.Debug("tulongequip:处理脱下屠龙装,完成")
	return nil
}

//脱下
func takeOff(pl player.Player, suitType tulongequiptypes.TuLongSuitType, pos inventorytypes.BodyPositionType) (err error) {
	return tulongequiplogic.HandleTakeOff(pl, suitType, pos)
}
