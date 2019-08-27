package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_USE_GOLD_EQUIP_TYPE), dispatch.HandlerFunc(handleUseGoldEquip))
}

//使用元神金装
func handleUseGoldEquip(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理使用元神金装")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csUseGoldEquip := msg.(*uipb.CSUseGoldEquip)
	index := csUseGoldEquip.GetIndex()
	if index < 0 {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = useGoldEquip(tpl, index)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理脱下元神金装,错误")

		return err
	}
	log.Debug("inventory:处理脱下元神金装,完成")
	return nil
}

//使用元神金装
func useGoldEquip(pl player.Player, index int32) (err error) {
	return goldequiplogic.HandleUseGoldEquip(pl, index)
}
