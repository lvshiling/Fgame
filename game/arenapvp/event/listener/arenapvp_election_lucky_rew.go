package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	arenapvpeventtypes "fgame/fgame/game/arenapvp/event/types"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	"fmt"
)

//竞猜推送
func arenapvpElectionLuckyRew(target event.EventTarget, data event.EventData) (err error) {
	electionData, ok := target.(*arenapvpdata.ElectionData)
	if !ok {
		return
	}

	electionIndexStr := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpLuckyIndexText), electionData.ElectionIndex+1))
	luckyNameStr := chatlogic.FormatMailKeyWordNoticeStr(electionData.LuckyNameText)
	rewNameStr := chatlogic.FormatMailKeyWordNoticeStr(lang.GetLangService().ReadLang(lang.ArenapvpLuckyRewNameText))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpLuckySystemContent), electionIndexStr, luckyNameStr, rewNameStr)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	return
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpElectionLuckyRew, event.EventListenerFunc(arenapvpElectionLuckyRew))
}
