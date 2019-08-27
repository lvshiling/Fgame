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
	xinfaeventtypes "fgame/fgame/game/xinfa/event/types"
	xinfatypes "fgame/fgame/game/xinfa/types"
	"fgame/fgame/game/xinfa/xinfa"
	"fmt"
)

//玩家心法激活
func playerXinFaActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	xinfaType := data.(xinfatypes.XinFaType)
	xinFaTemplate := xinfa.GetXinFaService().GetXinFaByTypeAndLevel(xinfaType, 1)
	if xinFaTemplate == nil {
		return
	}

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	xinfaName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(xinFaTemplate.Name))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.XinFaActivateNotice), playerName, xinfaName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(xinfaeventtypes.EventTypeXinFaActive, event.EventListenerFunc(playerXinFaActivate))
}
