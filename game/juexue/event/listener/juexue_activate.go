package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	juexueeventtypes "fgame/fgame/game/juexue/event/types"
	juexuetypes "fgame/fgame/game/juexue/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fmt"
)

//玩家绝学激活
func playerJueXueActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	typ := data.(juexuetypes.JueXueType)

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	juexueName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(typ.String()))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.JueXueActivateNotice), playerName, juexueName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(juexueeventtypes.EventTypeJueXueAcitivate, event.EventListenerFunc(playerJueXueActivate))
}
