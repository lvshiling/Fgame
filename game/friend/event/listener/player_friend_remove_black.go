package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	playerfriend "fgame/fgame/game/friend/player"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

func playerFriendRemoveBlack(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	friendId, ok := data.(int64)
	if !ok {
		return
	}

	fr := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fr == nil {
		//写日志
		friendBlackLog(friendId, pl.GetId(), friendtypes.FriendLogTypeRemoveBlack)
	} else {
		frCtx := scene.WithPlayer(context.Background(), fr)
		playerFriendBlackMsg := message.NewScheduleMessage(onPlayerFriendRemoveBlack, frCtx, pl.GetId(), nil)
		fr.Post(playerFriendBlackMsg)
	}
	return
}

func onPlayerFriendRemoveBlack(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	friendId := result.(int64)

	manager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	manager.ReverseRemoveBlackFriend(friendId)
	return nil
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendRemoveBlack, event.EventListenerFunc(playerFriendRemoveBlack))
}
