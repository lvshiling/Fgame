package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	gemeventtypes "fgame/fgame/game/gem/event/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fmt"
)

//玩家矿工激活
func playerMineActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.GemMineActivateNotice), playerName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(gemeventtypes.EventTypeMineActivate, event.EventListenerFunc(playerMineActivate))
}
