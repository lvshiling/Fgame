package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	"fmt"
)

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeDaoJuTypeChange, event.EventListenerFunc(jieYiDaoJuTypeHelpChange))
}

// 监听到道具改变
func jieYiDaoJuTypeHelpChange(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	daoJu, _ := data.(jieyitypes.JieYiDaoJuType)
	if daoJu != jieyitypes.JieYiDaoJuTypeHigh {
		return
	}

	// 是高级，发送时装
	daojuTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiDaoJuTemplate(daoJu)

	fashionMap := daojuTemp.GetFashionMap()
	if len(fashionMap) != 0 {
		title := lang.GetLangService().ReadLang(lang.EmailJieYiDaoJuFashionTitle)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailJieYiDaoJuFashionContent))
		emaillogic.AddEmail(pl, title, content, fashionMap)
	}
	return
}
