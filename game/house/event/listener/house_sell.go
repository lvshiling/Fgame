package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	houseventtypes "fgame/fgame/game/house/event/types"
	"fgame/fgame/game/house/house"
	housetypes "fgame/fgame/game/house/types"
	"fgame/fgame/game/player"
)

//玩家房子出售
func playerHouseSell(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*houseventtypes.PlayerHouseSellEventData)
	if !ok {
		return
	}

	operateType := housetypes.HouseOperateTypeSell
	house.GetHouseService().AddLog(pl.GetName(), eventData.GetHouseIndex(), int32(eventData.GetHouseType()), eventData.GetHouseLevel(), operateType)
	return
}

func init() {
	gameevent.AddEventListener(houseventtypes.EventTypeHouseSell, event.EventListenerFunc(playerHouseSell))
}
