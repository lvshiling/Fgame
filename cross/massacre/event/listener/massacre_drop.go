package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/massacre/pbutil"
	"fgame/fgame/cross/player/player"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//戮仙刃掉落检测
func massacreDropCheck(target event.EventTarget, data event.EventData) (err error) {
	bo, ok := target.(scene.BattleObject)
	if !ok {
		return
	}
	pl, ok := bo.(*player.Player)
	if !ok {
		return
	}

	attackId := data.(int64)
	spl := player.GetOnlinePlayerManager().GetPlayerById(attackId)
	if spl == nil {
		return
	}

	if pl.GetScene() != spl.GetScene() {
		return
	}

	//该地图能否掉落杀气
	s := spl.GetScene()
	if s.MapTemplate().CanShaqiDrop() == false {
		return
	}

	isMassacreDrop := pbutil.BuildISMassacreDrop(attackId, spl.GetName())
	pl.SendMsg(isMassacreDrop)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(massacreDropCheck))
}
