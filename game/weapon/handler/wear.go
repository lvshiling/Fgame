package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_WEAPON_WEAR_TYPE), dispatch.HandlerFunc(handleWeaponWear))

}

//处理兵魂穿戴信息
func handleWeaponWear(s session.Session, msg interface{}) (err error) {
	log.Debug("weapon:处理兵魂穿戴信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWeaponWear := msg.(*uipb.CSWeaponWear)
	weaponId := csWeaponWear.GetWeaponWear()

	err = weaponWear(tpl, weaponId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weaponId": weaponId,
				"error":    err,
			}).Error("weapon:处理兵魂穿戴信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("weapon:处理兵魂穿戴信息完成")
	return nil

}

//兵魂穿戴逻辑
func weaponWear(pl player.Player, weaponId int32) (err error) {
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	flag := weaponManager.IsValid(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = weaponManager.IfWeaponExist(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:未激活的兵魂,无法穿戴")
		playerlogic.SendSystemMessage(pl, lang.WeaponNotActiveNotWear)
		return
	}

	flag = weaponManager.Wear(weaponId)
	if !flag {
		panic(fmt.Errorf("weapon: weaponWear should be ok"))
	}
	scWeaponWear := pbutil.BuildSCWeaponWear(weaponId)
	pl.SendMsg(scWeaponWear)
	return
}
