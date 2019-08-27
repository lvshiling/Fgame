package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	qixueeventtypes "fgame/fgame/game/qixue/event/types"
	qixueventtypes "fgame/fgame/game/qixue/event/types"
	qixuetemplate "fgame/fgame/game/qixue/template"
	weaponlogic "fgame/fgame/game/weapon/logic"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"
)

//玩家泣血枪降阶
func playerQiXueDegrade(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	eventData := data.(*qixueeventtypes.PlayerQiXueChangedWeaponEventData)
	oldLev := eventData.GetOldLevel()
	oldStar := eventData.GetOldStar()
	newLev := eventData.GetNewLevel()
	newStar := eventData.GetNewStar()

	oldQixueTemplate := qixuetemplate.GetQiXueTemplateService().GetQiXueTemplateByLevel(oldLev, oldStar)
	if oldQixueTemplate == nil {
		return
	}
	newQixueTemplate := qixuetemplate.GetQiXueTemplateService().GetQiXueTemplateByLevel(newLev, newStar)
	if newQixueTemplate == nil {
		return
	}

	//兵魂变化
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	oldWeaponId := oldQixueTemplate.WeaponId
	newWeaponId := newQixueTemplate.WeaponId
	if oldWeaponId != 0 && newWeaponId == 0 {
		flag := weaponManager.WeaponRemoveTemp(oldWeaponId)
		if !flag {
			return
		}

		//同步属性
		weaponlogic.WeaponPropertyChanged(pl)
		weaponWear := weaponManager.GetWeaponWear()
		weaponMap := weaponManager.GetAllWeapon()
		scWeaponGet := pbutil.BuildSCWeaponGet(weaponWear, weaponMap)
		pl.SendMsg(scWeaponGet)
	}

	return
}

func init() {
	gameevent.AddEventListener(qixueventtypes.EventTypeQiXueDegrade, event.EventListenerFunc(playerQiXueDegrade))
}
