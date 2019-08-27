package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fgame/fgame/game/welfare/welfare"
	"fmt"
)

//添加抽奖日志
func addDrewLog(target event.EventTarget, data event.EventData) (err error) {
	groupId := target.(int32)
	eventData, ok := data.(*welfareeventtypes.DrewAddLogEventData)
	if !ok {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		return
	}

	itemId := eventData.GetItemId()
	plName := eventData.GetPlayerName()
	itemNum := eventData.GetItemNum()
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return
	}
	quality := itemTemplate.GetQualityType()
	if quality < itemtypes.ItemQualityTypeOrange {
		return
	}
	// 添加日志
	welfare.GetWelfareService().AddDrewLog(groupId, plName, itemId, itemNum)

	// 系统公告
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	itemName := coreutils.FormatColor(quality.GetColor(), coreutils.FormatNoticeStrUnderline(itemTemplate.FormateItemNameOfNum(itemNum)))
	linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
	itemNameLink := coreutils.FormatLink(itemName, linkArgs)
	plyerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(plName))
	acName := chatlogic.FormatModuleNameNoticeStr(timeTemp.Name)
	args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
	link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityDrewRewNotice), plyerName, acName, itemNameLink, link)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeDrewAddLog, event.EventListenerFunc(addDrewLog))
}
