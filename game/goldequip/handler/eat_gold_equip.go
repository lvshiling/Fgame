package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EAT_GOLD_EQUIP_TYP), dispatch.HandlerFunc(handleGoldEquipEat))
}

//处理吞噬元神金装
func handleGoldEquipEat(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理元神金装吞噬")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csEatGoldEquip := msg.(*uipb.CSEatGoldEquip)
	itemIndexList := csEatGoldEquip.GetIndexList()

	err = goldEquipEat(tpl, itemIndexList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
				"error":     err,
			}).Error("goldequip:处理元神金装吞噬,错误")

		return err
	}
	log.Debug("goldequip:处理元神金装吞噬,完成")
	return nil
}

//吞噬
func goldEquipEat(pl player.Player, itemIndexList []int32) (err error) {
	isAuto := int32(0)
	return goldequiplogic.HandleGoldEquipEat(pl, isAuto, itemIndexList)
}
