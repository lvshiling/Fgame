package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_USE_EQUIP_TYPE), dispatch.HandlerFunc(handleInventoryUseEquip))
}

//使用装备
func handleInventoryUseEquip(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理使用装备")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryUseEquip := msg.(*uipb.CSInventoryUseEquip)
	index := csInventoryUseEquip.GetIndex()
	if index < 0 {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = useEquip(tpl, index)
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

//使用装备
func useEquip(pl player.Player, index int32) (err error) {
	return logic.UseEquip(pl, index)
}
