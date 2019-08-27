package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"
)

//灵池争夺合服
func oneArenaMergeServer(target event.EventTarget, data event.EventData) (err error) {
	playerId, ok := target.(int64)
	if !ok {
		return
	}

	emailTitle := lang.GetLangService().ReadLang(lang.OneArenaMergeServerTitle)
	emailContent := lang.GetLangService().ReadLang(lang.OneArenaMergeServerContent)
	emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, nil)
	return
}

func init() {
	gameevent.AddEventListener(onearenaeventtypes.EventTypeOneArenaMergeServer, event.EventListenerFunc(oneArenaMergeServer))
}
