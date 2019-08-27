package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	transportationeventtypes "fgame/fgame/game/transportation/event/types"
	transportationtypes "fgame/fgame/game/transportation/types"
)

//玩家参加仙盟镖车
func playerAcceptTransportation(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	typ, ok := data.(transportationtypes.TransportationType)
	if !ok {
		return
	}

	err = transportationDart(pl, typ)
	if err != nil {
		return
	}

	err = transportationCar(pl, typ)
	if err != nil {
		return
	}

	switch typ {
	case transportationtypes.TransportationTypeAlliance:
		{
			return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAlliance, 0, 1)
		}
	}
	return
}

//进行X次押镖
func transportationDart(pl player.Player, typ transportationtypes.TransportationType) (err error) {
	switch typ {
	case transportationtypes.TransportationTypeSilver,
		transportationtypes.TransportationTypeGold:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeDart, 0, 1)
	}
	return
}

//所在仙盟开启X次仙盟镖车
func transportationCar(pl player.Player, typ transportationtypes.TransportationType) (err error) {
	switch typ {
	case transportationtypes.TransportationTypeSilver,
		transportationtypes.TransportationTypeGold,
		transportationtypes.TransportationTypeAlliance:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeDartCar, 0, 1)
	}
	return
}

func init() {
	gameevent.AddEventListener(transportationeventtypes.EventTypeTransportationAccept, event.EventListenerFunc(playerAcceptTransportation))
}
