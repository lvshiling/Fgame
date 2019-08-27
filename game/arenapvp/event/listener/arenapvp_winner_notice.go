package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	arenapvpeventtypes "fgame/fgame/game/arenapvp/event/types"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	chatlogic "fgame/fgame/game/chat/logic"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fmt"
)

//晋级推送
func arenapvpWinnerNotice(target event.EventTarget, data event.EventData) (err error) {
	battlePlList, ok := target.([]*arenapvpdata.PvpPlayerInfo)
	if !ok {
		return
	}

	pvpType, ok := data.(arenapvptypes.ArenapvpType)
	if !ok {
		return
	}

	var noticeList []*arenapvpdata.PvpPlayerInfo
	for _, battlePl := range battlePlList {
		battleData := battlePl.GetBattleData(pvpType)
		if battleData == nil {
			continue
		}

		noticeList = append(noticeList, battlePl)
	}

	nameText := ""
	for index, battlePl := range noticeList {
		if index == 0 {
			nameText += battlePl.PlayerName
		} else {
			nameText += "," + battlePl.PlayerName
		}
	}

	nameText = chatlogic.FormatMailKeyWordNoticeStr(nameText)
	content := ""
	switch pvpType {
	case arenapvptypes.ArenapvpTypeTop32:
		{
			content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpWinnerTOP32Content), nameText)
		}
	case arenapvptypes.ArenapvpTypeTop16:
		{
			content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpWinnerTOP16Content), nameText)
		}
	case arenapvptypes.ArenapvpTypeTop8:
		{
			content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpWinnerTOP8Content), nameText)
		}
	case arenapvptypes.ArenapvpTypeTop4:
		{
			content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpWinnerTOP4Content), nameText)
		}
	case arenapvptypes.ArenapvpTypeFinals:
		{
			content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpWinnerFinalsContent), nameText)
		}
	case arenapvptypes.ArenapvpTypeChampion:
		{
			winnerName := ""
			for _, battlePl := range noticeList {
				battleData := battlePl.GetBattleData(pvpType)
				if battlePl.PlayerId != battleData.WinnerId {
					continue
				}
				winnerName = battlePl.PlayerName
			}

			winnerName = chatlogic.FormatMailKeyWordNoticeStr(winnerName)
			content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpWinnerChampionContent), winnerName)
		}
	}

	//公告推送
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpWinnerNotice, event.EventListenerFunc(arenapvpWinnerNotice))
}
