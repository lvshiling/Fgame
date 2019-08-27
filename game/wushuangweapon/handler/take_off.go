package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	wushuangweaponlogic "fgame/fgame/game/wushuangweapon/logic"
	"fgame/fgame/game/wushuangweapon/pbutil"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_TAKE_OFF_TYPE), dispatch.HandlerFunc(handlerTakeOff))
}

func handlerTakeOff(s session.Session, msg interface{}) (err error) {
	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSWushuangWeaponTakeOff)
	csBodyPos := csMsg.GetBodyPos()
	bodyPos := wushuangweapontypes.WushuangWeaponPart(csBodyPos)

	if !bodyPos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("WushuangWeapon:部位脱掉请求，类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = takeOff(tpl, bodyPos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("WushuangWeapon:处理穿上无双神器，错误")
		return
	}
	return
}

func takeOff(pl player.Player, bodyPos wushuangweapontypes.WushuangWeaponPart) (err error) {
	wushuangDataManager := pl.GetPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType).(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	slotObj := wushuangDataManager.GetSlotObjectFromBodyPos(bodyPos)

	//判断是否已经装备
	if !slotObj.IsEquip() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("wushuangWeapon:部位没有装备")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag := wushuangweaponlogic.TakeOffLogic(pl, bodyPos)
	if !flag {
		return
	}

	//属性更新
	wushuangweaponlogic.WushuangWeaponPropertyChanged(pl)

	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	//发消息
	scMsg := pbutil.BuildSCWushuangWeaponTakeOff()
	pl.SendMsg(scMsg)
	return
}
