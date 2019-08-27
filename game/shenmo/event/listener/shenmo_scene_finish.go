package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
	shenmologic "fgame/fgame/game/shenmo/logic"
	"fgame/fgame/game/shenmo/pbutil"
	shenmoscene "fgame/fgame/game/shenmo/scene"
	"fgame/fgame/game/shenmo/shenmo"
)

//神魔战场场景结束
func shenMoSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(shenmoscene.ShenMoSceneData)
	if !ok {

		return
	}

	//神魔战场结果
	allPlayer := sd.GetScene().GetAllPlayers()
	for _, pl := range allPlayer {
		if pl == nil {
			continue
		}

		scShenMoSceneEnd := pbutil.BuildSCShenMoSceneEnd()
		pl.SendMsg(scShenMoSceneEnd)
	}
	lineList := shenmo.GetShenMoService().GetAllLineUpList()
	shenmologic.BroadShenMoFinishToLineUpCancle(lineList)
	shenmo.GetShenMoService().ShenMoSceneFinish()
	return
}

func init() {
	gameevent.AddEventListener(shenmoeventtypes.EventTypeShenMoSceneFinish, event.EventListenerFunc(shenMoSceneFinish))
}
