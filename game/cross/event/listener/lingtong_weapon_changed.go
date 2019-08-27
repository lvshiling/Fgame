package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//灵童领域变化
func lingTongWeaponChanged(target event.EventTarget, data event.EventData) (err error) {
	lingTong, ok := target.(scene.LingTong)
	if !ok {
		return
	}
	owner := lingTong.GetOwner()
	pl, ok := owner.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}
	weaponId := lingTong.GetLingTongWeaponId()
	weaponState := lingTong.GetLingTongWeaponState()
	lingTongWeaponChanged := pbutil.BuildLingTongWeaponChanged(weaponId, weaponState)
	pl.SendCrossMsg(lingTongWeaponChanged)

	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeBattleLingTongShowWeaponChanged, event.EventListenerFunc(lingTongWeaponChanged))
}
