package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	qixueventtypes "fgame/fgame/game/qixue/event/types"
	qixuetemplate "fgame/fgame/game/qixue/template"
	weaponlogic "fgame/fgame/game/weapon/logic"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"
)

//玩家泣血枪进阶
func playerQiXueAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*qixueventtypes.PlayerQiXueChangedWeaponEventData)
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

	//激活兵魂
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	oldWeaponId := oldQixueTemplate.WeaponId
	newWeaponId := newQixueTemplate.WeaponId
	if oldWeaponId == 0 && newWeaponId != 0 {
		flag := weaponManager.WeaponActiveTemp(newWeaponId)
		if !flag {
			return
		}
		//同步属性
		weaponlogic.WeaponPropertyChanged(pl)
		scWeaponActive := pbutil.BuildSCWeaponActive(newWeaponId)
		pl.SendMsg(scWeaponActive)
	}
	return
}

func init() {
	gameevent.AddEventListener(qixueventtypes.EventTypeQiXueAdvanced, event.EventListenerFunc(playerQiXueAdavanced))
}
