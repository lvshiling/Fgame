package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/marry/marry"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//送花
func friendGift(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*friendeventtypes.FriendGiftEventData)
	if !ok {
		return
	}
	spouseId := eventData.GetFriendId()
	num := eventData.GetNum()

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	if marryInfo.SpouseId != spouseId {
		return
	}

	marry.GetMarryService().AddPoint(pl.GetId(), num)

	// name := ""
	// spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	// if spl != nil {
	// 	name = spl.GetName()
	// } else {
	// 	playerInfo, _ := player.GetPlayerService().GetPlayerInfo(spouseId)
	// 	name = playerInfo.Name
	// }

	// playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(name))
	// peerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	// content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryPeerGiveGift), playerName, peerName, num)
	// //跑马灯
	// noticelogic.NoticeNumBroadcast([]byte(content), 0, int32(1))
	// //系统公告
	// chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	return
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendGift, event.EventListenerFunc(friendGift))
}
