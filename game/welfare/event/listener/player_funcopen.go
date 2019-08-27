package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretemplate "fgame/fgame/game/welfare/template"
)

//玩家功能开启
func playerFuncOpen(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	funcType := data.(funcopentypes.FuncOpenType)

	// 功能开启邮件提醒
	funcOpenMailNotice(pl, funcType)
	return
}

func funcOpenMailNotice(pl player.Player, funcType funcopentypes.FuncOpenType) {
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByFuncType(funcType)
	for _, timeTemp := range timeTempList {
		welfarelogic.SendOpenNoticeMail(pl, timeTemp.Group)
	}
}

func init() {
	gameevent.AddEventListener(funcopeneventtypes.EventTypeFuncOpen, event.EventListenerFunc(playerFuncOpen))
}
