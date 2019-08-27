package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/marry/pbutil"
	marryscene "fgame/fgame/game/marry/scene"
)

//结婚场景玩家名字变化
func marryScenePlayerNameChanged(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(marryscene.MarrySceneData)
	if !ok {
		return
	}
	period := sd.GetPeriod()
	playerId, name, spouseId, spouseName := sd.GetBothName()
	heroismList := sd.GetHeroismList()
	scMarryBanquet := pbuitl.BuildSCMarryBanquet(period, playerId, name, spouseId, spouseName, heroismList)
	sd.GetScene().BroadcastMsg(scMarryBanquet)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryScenePlayerNameChanged, event.EventListenerFunc(marryScenePlayerNameChanged))
}
