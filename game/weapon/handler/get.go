package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WEAPON_GET_TYPE), dispatch.HandlerFunc(handleWeaponGet))
}

//处理获取兵魂信息
func handleWeaponGet(s session.Session, msg interface{}) (err error) {
	log.Debug("weapon:处理获取兵魂信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = weaponGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("weapon:处理获取兵魂信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("weapon:处理获取兵魂信息完成")
	return nil
}

//处理兵魂界面信息逻辑
func weaponGet(pl player.Player) (err error) {
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	weaponWear := weaponManager.GetWeaponWear()
	weaponMap := weaponManager.GetAllWeapon()
	scWeaponGet := pbutil.BuildSCWeaponGet(weaponWear, weaponMap)
	pl.SendMsg(scWeaponGet)
	return
}
