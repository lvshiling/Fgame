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

//玩家等级变化
func playerLevelChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	// 获取等级条件
	noticeType := friendtypes.FriendNoticeTypeUplevel
	noticeTempList := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeTemplate(noticeType)
	for _, noticeTemp := range noticeTempList {
		if pl.GetLevel() != noticeTemp.TiaoJian {
			continue
		}

		// 推送消息
		friendlogic.BroadcastFriendNotice(pl, noticeType, noticeTemp.TiaoJian, "")
	}
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerLevelChanged, event.EventListenerFunc(playerLevelChanged))
}
