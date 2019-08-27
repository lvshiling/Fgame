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

func playerFriendAdd(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fo, ok := data.(*friend.FriendObject)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	manager.AddFrined(fo)

	friendId := fo.FriendId
	if friendId == pl.GetId() {
		friendId = fo.PlayerId
	}
	playerInfo, err := player.GetPlayerService().GetPlayerInfo(friendId)
	if err != nil {
		return err
	}
	scFriendAdd := pbutil.BuildSCFriendAdd(friendId, fo.Point, playerInfo, true)
	pl.SendMsg(scFriendAdd)

	fr := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fr == nil {
		return
	}
	frCtx := scene.WithPlayer(context.Background(), fr)
	playerEnterSceneMsg := message.NewScheduleMessage(onPlayerFriendAdd, frCtx, fo, nil)
	fr.Post(playerEnterSceneMsg)
	return
}

func onPlayerFriendAdd(ctx context.Context, result interface{}, err error) error {
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
	manager.AddFrined(fo)

	playerInfo, err := player.GetPlayerService().GetPlayerInfo(friendId)
	if err != nil {
		return err
	}
	scFriendAdd := pbutil.BuildSCFriendAdd(friendId, fo.Point, playerInfo, true)
	pl.SendMsg(scFriendAdd)
	return nil
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendAdd, event.EventListenerFunc(playerFriendAdd))
}
