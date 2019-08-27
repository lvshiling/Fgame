package listener

import (
	"fgame/fgame/core/event"
	playerdailiwan "fgame/fgame/game/daliwan/player"
	daliwantemplate "fgame/fgame/game/daliwan/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerDaLiWanDataManagerType).(*playerdailiwan.PlayerDaLiWanManager)
	now := global.GetGame().GetTimeService().Now()
	//添加buff
	for _, obj := range manager.GetDaliWanMap() {
		if obj.IsExpire(now) {
			continue
		}
		linshiTemplate := daliwantemplate.GetDaLiWanTemplateService().GetLinShiTemplate(obj.GetTyp())
		if linshiTemplate == nil {
			continue
		}

		// scenelogic.AddBuff(pl, linshiTemplate.BuffId, pl.GetId(), common.MAX_RATE)
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
