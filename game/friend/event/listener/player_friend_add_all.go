package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/pbutil"
	playerfriend "fgame/fgame/game/friend/player"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

func playerFriendInviteAll(target event.EventTarget, data event.EventData) (err error) {
	plList, ok := target.([]player.Player)
	if !ok {
		return
	}
	pl, ok := data.(player.Player)
	if !ok {
		return
	}

	for _, friPl := range plList {
		firPlCtx := scene.WithPlayer(context.Background(), friPl)
		playerFriendInviteMsg := message.NewScheduleMessage(onPlayerFriendInviteAll, firPlCtx, pl.GetId(), nil)
		friPl.Post(playerFriendInviteMsg)
	}
	return
}

func onPlayerFriendInviteAll(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	friendId, ok := result.(int64)
	if !ok {
		return nil
	}

	manager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	//判断好友数量
	numFriends := manager.NumOfFriend()
	maxFriends := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFriendLimit)
	if numFriends >= int(maxFriends) {
		return nil
	}

	flag := manager.IsBlack(friendId)
	if flag {
		return nil
	}
	playerInfo, err := player.GetPlayerService().GetPlayerInfo(friendId)
	if err != nil {
		return err
	}
	if playerInfo == nil {
		return nil
	}

	manager.FriendInvite(friendId)
	now := global.GetGame().GetTimeService().Now()
	scFriendInvitePushPeer := pbutil.BuildSCFriendInvitePushPeer(playerInfo, now)
	pl.SendMsg(scFriendInvitePushPeer)
	return nil
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendAddAll, event.EventListenerFunc(playerFriendInviteAll))
}
