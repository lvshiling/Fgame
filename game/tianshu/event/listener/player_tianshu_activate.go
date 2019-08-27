package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	tianshueventtypes "fgame/fgame/game/tianshu/event/types"
	tianshulogic "fgame/fgame/game/tianshu/logic"
	playertianshu "fgame/fgame/game/tianshu/player"
	tianshutemplate "fgame/fgame/game/tianshu/template"
)

//玩家天书激活
func tianshuActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	obj, ok := data.(*playertianshu.PlayerTianShuObject)
	if !ok {
		return
	}

	tianshulogic.TianShuPropertyChanged(pl)

	// 同步
	typ := obj.GetType()
	level := obj.GetLevel()
	tianshuTemp := tianshutemplate.GetTianShuTemplateService().GetTianShuTemplate(typ, level)
	rate := tianshuTemp.Tequan
	pl.AddTianShu(typ, rate)

	return
}

func init() {
	gameevent.AddEventListener(tianshueventtypes.EventTypePlayerTianShuActivate, event.EventListenerFunc(tianshuActivate))
}
