package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	"fgame/fgame/game/mount/mount"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
)

//玩家幻化激活
func playerMountUnrealActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	mountId := data.(int)
	mountTemplate := mount.GetMountService().GetMount(mountId)
	if mountTemplate == nil {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	manager.ActiveSeqId(wardrobetypes.WardrobeSysTypeMount, int32(mountId))
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountUnrealActivate, event.EventListenerFunc(playerMountUnrealActivate))
}
