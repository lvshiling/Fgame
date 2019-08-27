package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	houseventtypes "fgame/fgame/game/house/event/types"
	"fgame/fgame/game/house/house"
	playerhouse "fgame/fgame/game/house/player"
	housetypes "fgame/fgame/game/house/types"
	"fgame/fgame/game/player"
)

//玩家房子激活
func playerHouseActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	houseObj, ok := data.(*playerhouse.PlayerHouseObject)
	if !ok {
		return
	}

	operateType := housetypes.HouseOperateTypeActivate
	house.GetHouseService().AddLog(pl.GetName(), houseObj.GetHouseIndex(), int32(houseObj.GetHouseType()), houseObj.GetHouseLevel(), operateType)
	return
}

func init() {
	gameevent.AddEventListener(houseventtypes.EventTypeHouseActivate, event.EventListenerFunc(playerHouseActivate))
}
