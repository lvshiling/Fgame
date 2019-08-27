package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	souleventtypes "fgame/fgame/game/soul/event/types"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"

	"fmt"
)

//帝魂激活
func soulActive(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	soulTag, ok := data.(soultypes.SoulType)
	if !ok {
		return
	}

	temp := soul.GetSoulService().GetSoulActiveTemplate(soulTag)
	if temp == nil {
		return
	}

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	soulName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(temp.Name))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.SoulActivateNotice), playerName, soulName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulActive, event.EventListenerFunc(soulActive))
}
