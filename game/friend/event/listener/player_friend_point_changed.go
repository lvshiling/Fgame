package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/friend"
	friendlogic "fgame/fgame/game/friend/logic"
	"fgame/fgame/game/player"
)

func playerFriendPointChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fo, ok := data.(*friend.FriendObject)
	if !ok {
		return
	}

	friendId := fo.FriendId
	if friendId == pl.GetId() {
		friendId = fo.PlayerId
	}

	point := fo.Point
	friendlogic.FrinedPointChanged(pl, friendId, point)

	fr := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fr == nil {
		return
	}
	friendlogic.FrinedPointChanged(fr, pl.GetId(), point)

	return
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendPointChanged, event.EventListenerFunc(playerFriendPointChanged))
}
