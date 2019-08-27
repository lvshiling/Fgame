package listener

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/friend/pbutil"
	playerfriend "fgame/fgame/game/friend/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}

	//TODO:YLZ 需要把日志放在管理器加载 不然会阻塞主进程
	friendManager := p.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)

	//好友信息
	var friendList []*uipb.FriendInfo
	var blackList []*uipb.FriendInfo
	for _, fri := range friendManager.GetFriends() {
		friendId := fri.FriendId
		if friendId == p.GetId() {
			friendId = fri.PlayerId
		}
		flag := friendManager.IsBlack(friendId)
		if flag {
			continue
		}

		point := fri.Point
		friendInfo, err := player.GetPlayerService().GetPlayerInfo(friendId)
		if err != nil {
			return err
		}
		if friendInfo == nil {
			continue
		}
		isBlacked := friendManager.IsBlacked(friendId)
		tempFri := pbutil.BuildFriend(friendId, point, friendInfo, isBlacked)
		friendList = append(friendList, tempFri)
	}
	for _, fri := range friendManager.GetBlacks() {
		friendId := fri.FriendId
		friendInfo, err := player.GetPlayerService().GetPlayerInfo(friendId)
		if err != nil {
			return err
		}
		if friendInfo == nil {
			continue
		}
		tempFri := pbutil.BuildFriend(friendId, 0, friendInfo, false)
		blackList = append(blackList, tempFri)
	}

	feedbackList := friendManager.GetFriendFeedbackList()
	record := friendManager.GetReceiveRewRecord()
	dummyNum := friendManager.GetDummyFriendNum()
	scFriendsGet := pbutil.BuildSCFriendsGet(friendList, blackList, feedbackList, record, dummyNum)
	p.SendMsg(scFriendsGet)

	inviteMap := friendManager.GetFriendInviteMap()
	scFriendInviteList := pbutil.BuildSCFriendInviteList(inviteMap)
	p.SendMsg(scFriendInviteList)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
