package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fgame/fgame/game/welfare/welfare"
	"fmt"
)

//添加拉霸日志
func addLaBaLog(target event.EventTarget, data event.EventData) (err error) {
	groupId := target.(int32)
	eventData, ok := data.(*welfareeventtypes.LaBaAddLogEventData)
	if !ok {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		return
	}

	plName := eventData.GetPlayerName()
	rewGold := eventData.GetRewGold()
	costGold := eventData.GetCostGold()
	welfare.GetWelfareService().AddLaBaLog(groupId, plName, costGold, rewGold)

	// 系统公告
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	plyerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(plName))
	useGlodText := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), costGold)))
	rewGoldText := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), rewGold)))
	acName := chatlogic.FormatModuleNameNoticeStr(timeTemp.Name)
	args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
	link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityLaBaRewGoldNotice), plyerName, acName, useGlodText, rewGoldText, link)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeLaBaAddLog, event.EventListenerFunc(addLaBaLog))
}
