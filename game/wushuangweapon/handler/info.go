package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/wushuangweapon/pbutil"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_INFO_TYPE), dispatch.HandlerFunc(handlerInfo))
}

func handlerInfo(s session.Session, msg interface{}) (err error) {
	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = info(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("WushuangWeapon:处理部位信息请求，错误")
		return
	}
	return
}

func info(pl player.Player) (err error) {
	wushuangDataManager := pl.GetPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType).(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	var bodyposinfoList []*playerwushuangweapon.PlayerWushuangWeaponSlotObject
	for _, slotObj := range wushuangDataManager.GetSlotObjectMap() {
		if slotObj.IsEquip() {
			bodyposinfoList = append(bodyposinfoList, slotObj)
		}
	}
	scWushuangWeaponInfo := pbutil.BuildSCWushuangWeaponInfo(bodyposinfoList)
	pl.SendMsg(scWushuangWeaponInfo)
	return
}
