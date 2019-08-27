package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/game/common/common"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	xianzuncardeventtypes "fgame/fgame/game/xianzuncard/event/types"
	xianzuncardtemplate "fgame/fgame/game/xianzuncard/template"
	"fmt"
)

func init() {
	gameevent.AddEventListener(xianzuncardeventtypes.EventTypeXianZunCardCrossDay, event.EventListenerFunc(playerCrossDay))
}

func playerCrossDay(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*xianzuncardeventtypes.PlayerXianZunCardCrossDayEventData)
	if !ok {
		return
	}
	typ := eventData.GetType()
	diffDay := eventData.GetDiffDay()

	xianZunTemp := xianzuncardtemplate.GetXianZunCardTemplateService().GetXianZunCardTemplate(typ)
	if xianZunTemp == nil {
		return
	}
	num := diffDay
	maxDay := int32(xianZunTemp.Duration / int64(common.DAY))
	if diffDay > maxDay {
		num = maxDay
	}

	now := global.GetGame().GetTimeService().Now()
	// 多少天没领取全部补送,分发邮件
	itemMap := xianZunTemp.GetEmailItemMap()
	if len(itemMap) != 0 {
		title := lang.GetLangService().ReadLang(lang.EmailXianZunCardCrossTitle)
		titleStr := fmt.Sprintf(title, xianZunTemp.Name)
		content := lang.GetLangService().ReadLang(lang.EmailXianZunCardCrossContent)
		contentStr := fmt.Sprintf(content, xianZunTemp.Name)
		for i := 0; i < int(num); i++ {
			creatTime := now - (int64(common.DAY) * int64(diffDay-1+int32(i)))
			emaillogic.AddEmailDefinTime(pl, titleStr, contentStr, creatTime, itemMap)
		}
	}

	return
}
