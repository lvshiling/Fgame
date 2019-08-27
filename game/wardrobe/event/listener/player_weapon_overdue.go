package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	weaponeventtypes "fgame/fgame/game/weapon/event/types"
)

//玩家兵魂失效
func playerWeaponOverdue(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	weaponId := data.(int32)
	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.RemoveSeqId(wardrobetypes.WardrobeSysTypeWeapon, weaponId)
	return
}

func init() {
	gameevent.AddEventListener(weaponeventtypes.EventTypeWeaponRemove, event.EventListenerFunc(playerWeaponOverdue))
}
