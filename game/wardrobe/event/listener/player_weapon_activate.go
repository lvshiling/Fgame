package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	weaponeventtypes "fgame/fgame/game/weapon/event/types"
	"fgame/fgame/game/weapon/weapon"
)

//玩家兵魂激活
func playerWeaponActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	weaponId := data.(int32)
	weaponTemplate := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if weaponTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.ActiveSeqId(wardrobetypes.WardrobeSysTypeWeapon, weaponId)
	return
}

func init() {
	gameevent.AddEventListener(weaponeventtypes.EventTypeWeaponActivate, event.EventListenerFunc(playerWeaponActivate))
}
