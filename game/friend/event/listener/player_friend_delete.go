package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/friend"
	"fgame/fgame/game/friend/pbutil"
	playerfriend "fgame/fgame/game/friend/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

func playerFriendDelete(target event.EventTarget, data event.EventData) (err error) {
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

	manager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	manager.DeleteFriend(friendId)
	scFriendDelete := pbutil.BuildSCFriendDelete(friendId)
	pl.SendMsg(scFriendDelete)

	fr := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fr == nil {
		return
	}
	frCtx := scene.WithPlayer(context.Background(), fr)
	playerFriendRemoveMsg := message.NewScheduleMessage(onPlayerFriendDelete, frCtx, fo, nil)
	fr.Post(playerFriendRemoveMsg)
	return
}

func onPlayerFriendDelete(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	fo, ok := result.(*friend.FriendObject)
	if !ok {
		return nil
	}

	friendId := fo.FriendId
	if friendId == pl.GetId() {
		friendId = fo.PlayerId
	}

	manager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	manager.DeleteFriend(friendId)
	scFriendDelete := pbutil.BuildSCFriendDelete(friendId)
	pl.SendMsg(scFriendDelete)
	return nil
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendDelete, event.EventListenerFunc(playerFriendDelete))
}
