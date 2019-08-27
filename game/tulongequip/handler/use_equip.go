package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	tulongequiplogic "fgame/fgame/game/tulongequip/logic"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TULONG_USE_EQUIP_TYPE), dispatch.HandlerFunc(handleUseTuLongEquip))
}

//使用屠龙装
func handleUseTuLongEquip(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理使用屠龙装")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTuLongUseEquip)

	index := csMsg.GetIndex()
	if index < 0 {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	suitType := tulongequiptypes.TuLongSuitType(csMsg.GetSuitType())
	if !suitType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用元神屠龙装备,装备类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = useTuLongEquip(tpl, index, suitType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理使用屠龙装,错误")

		return err
	}
	log.Debug("inventory:处理使用屠龙装,完成")
	return nil
}

//使用屠龙装
func useTuLongEquip(pl player.Player, index int32, suitType tulongequiptypes.TuLongSuitType) (err error) {
	return tulongequiplogic.HandleUseTuLongEquip(pl, index, suitType)
}
