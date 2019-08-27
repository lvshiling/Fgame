package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendlogic "fgame/fgame/game/friend/logic"
	friendtemplate "fgame/fgame/game/friend/template"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
)

//玩家转生变化
func playerZhuanShengChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	zhuanShu, ok := data.(int32)
	if !ok {
		return
	}
	noticeType := friendtypes.FriendNoticeTypeZhuanSheng
	noticeTempList := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeTemplate(noticeType)
	for _, noticeTemp := range noticeTempList {
		if zhuanShu != noticeTemp.TiaoJian {
			continue
		}

		// 推送消息
		friendlogic.BroadcastFriendNotice(pl, noticeType, noticeTemp.TiaoJian, "")
	}
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerZhuanShengChanged, event.EventListenerFunc(playerZhuanShengChanged))
}
