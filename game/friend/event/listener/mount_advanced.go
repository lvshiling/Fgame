package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendlogic "fgame/fgame/game/friend/logic"
	friendtemplate "fgame/fgame/game/friend/template"
	friendtypes "fgame/fgame/game/friend/types"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	"fgame/fgame/game/player"
)

//玩家坐骑进阶
func playerMountAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int)
	if !ok {
		return
	}
	noticeType := friendtypes.FriendNoticeTypeMountAdvanced
	noticeTempList := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeTemplate(noticeType)
	for _, noticeTemp := range noticeTempList {
		if int32(advanceId) != noticeTemp.TiaoJian {
			continue
		}

		// 推送消息
		friendlogic.BroadcastFriendNotice(pl, noticeType, noticeTemp.TiaoJian, "")
	}
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountAdvanced, event.EventListenerFunc(playerMountAdavanced))
}
