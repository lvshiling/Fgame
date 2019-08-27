package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendentity "fgame/fgame/game/friend/entity"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	playerfriend "fgame/fgame/game/friend/player"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/idutil"
)

func playerFriendBlack(target event.EventTarget, data event.EventData) (err error) {
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
		friendBlackLog(friendId, pl.GetId(), friendtypes.FriendLogTypeBlack)
	} else {
		frCtx := scene.WithPlayer(context.Background(), fr)
		playerFriendBlackMsg := message.NewScheduleMessage(onPlayerFriendBlack, frCtx, pl.GetId(), nil)
		fr.Post(playerFriendBlackMsg)
	}
	return
}

func friendBlackLog(playerId int64, frinedId int64, logType friendtypes.FriendLogType) {
	//写离线日志
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	playerFriendLogEntity := &friendentity.PlayerFriendLogEntity{
		Id:         id,
		PlayerId:   playerId,
		FriendId:   frinedId,
		Type:       int32(logType),
		CreateTime: now,
		DeleteTime: 0,
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(playerFriendLogEntity)
}

func onPlayerFriendBlack(ctx context.Context, result interface{}, err error) error {

	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	friendId := result.(int64)

	manager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	manager.ReverseBlackFriend(friendId)
	return nil
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendBlack, event.EventListenerFunc(playerFriendBlack))
}
