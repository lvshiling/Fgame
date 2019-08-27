package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	eventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/moonlove/logic"
	"fgame/fgame/game/moonlove/pbutil"
	playermoonlove "fgame/fgame/game/moonlove/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
)

//月下情缘增送礼物事件
func playerFriendGift(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	giftEventData, ok := data.(*eventtypes.FriendGiftEventData)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}
	//处于同一个场景，并且是月下情缘场景
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeYueXiaQingYuan {
		return
	}

	friendId := giftEventData.GetFriendId()
	giftNum := giftEventData.GetNum()
	charmNum := giftEventData.GetCharmNum()

	spl := s.GetPlayer(friendId)
	if spl == nil {
		return
	}

	fri, ok := spl.(player.Player)
	if !ok {
		return
	}

	plMoonloveManager := pl.GetPlayerDataManager(types.PlayerMoonloveDataManagerType).(*playermoonlove.PlayerMoonloveDataManager)
	friMoonloveManager := fri.GetPlayerDataManager(types.PlayerMoonloveDataManagerType).(*playermoonlove.PlayerMoonloveDataManager)

	//自己增加豪气
	plMoonloveManager.AddGenerousNum(giftNum)
	//好友增加魅力
	friMoonloveManager.AddCharmNum(charmNum)

	sceneData := pl.GetScene().SceneDelegate().(logic.MoonloveSceneData)

	//玩家
	plGenerousNum := plMoonloveManager.GetGenerousNum()
	ranking := sceneData.GenerousRankChange(plGenerousNum, pl.GetName(), pl.GetId())
	scMoonloveGenerousChanged := pbutil.BuildSCMoonloveGenerousChanged(pl.GetId(), plGenerousNum, ranking)
	pl.SendMsg(scMoonloveGenerousChanged)

	//好友
	friCharm := friMoonloveManager.GetCharmNum()
	ranking = sceneData.ChramRankChange(friCharm, fri.GetName(), fri.GetId())
	scMoonloveCharmChanged := pbutil.BuildSCMoonloveCharmChanged(fri.GetId(), friCharm, ranking)
	fri.SendMsg(scMoonloveCharmChanged)

	//弹幕
	allPlayers := pl.GetScene().GetAllPlayers()
	scMoonloveGiftNotice := pbutil.BuildSCMoonloveGiftNotice(pl.GetName(), fri.GetName(), giftNum)
	for _, spl := range allPlayers {
		spl.SendMsg(scMoonloveGiftNotice)
	}

	return
}

func init() {
	gameevent.AddEventListener(eventtypes.EventTypeFriendGift, event.EventListenerFunc(playerFriendGift))
}
