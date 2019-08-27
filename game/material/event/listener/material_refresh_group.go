package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	materialeventtypes "fgame/fgame/game/material/event/types"
	playermaterial "fgame/fgame/game/material/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//更新怪物波数
func playerMaterialRefreshGroup(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*materialeventtypes.RefreshGroupEventData)
	group := eventData.GetGroup()
	typ := eventData.GetMaterialType()

	materialManager := pl.GetPlayerDataManager(playertypes.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
	materialManager.RefreshGroup(group, typ)

	return
}

func init() {
	gameevent.AddEventListener(materialeventtypes.EventTypeMaterialRefreshGroup, event.EventListenerFunc(playerMaterialRefreshGroup))
}
