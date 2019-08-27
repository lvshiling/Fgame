package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/qixue/pbutil"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//戮仙刃掉落检测
func qiXueDropCheck(target event.EventTarget, data event.EventData) (err error) {
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

	//该地图能否掉落杀戮心
	s := spl.GetScene()
	if !s.MapTemplate().IfCanShaLuDrop() {
		return
	}

	isMsg := pbutil.BuildISQiXueDrop(attackId, spl.GetName())
	pl.SendMsg(isMsg)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(qiXueDropCheck))
}
