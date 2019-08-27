package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	activityeventtypes "fgame/fgame/game/activity/event/types"
	activitytemplate "fgame/fgame/game/activity/template"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

//活动开始提醒
func activityNoticeStart(target event.EventTarget, data event.EventData) (err error) {
	activityTimeTemp, ok := data.(*gametemplate.ActivityTimeTemplate)
	if !ok {
		return
	}

	activityTemp := activitytemplate.GetActivityTemplateService().GetActiveTemplate(int32(activityTimeTemp.ActivityId))
	if activityTemp == nil {
		return
	}
	minuteInt := activityTimeTemp.EndNoticeTime / int64(common.MINUTE)
	activityName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(activityTemp.Name))
	noticeContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ActivityStartNotice), minuteInt, activityName)

	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(noticeContent))
	noticelogic.NoticeNumBroadcast([]byte(noticeContent), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(activityeventtypes.EventTypeActivityNoticeStart, event.EventListenerFunc(activityNoticeStart))
}
