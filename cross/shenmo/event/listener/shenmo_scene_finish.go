package listener

import (
	"fgame/fgame/core/event"
	shenmologic "fgame/fgame/cross/shenmo/logic"
	"fgame/fgame/cross/shenmo/pbutil"
	gameevent "fgame/fgame/game/event"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
	shenmoscene "fgame/fgame/game/shenmo/scene"
	"fgame/fgame/game/shenmo/shenmo"

	log "github.com/Sirupsen/logrus"
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
	log.Info("shenmo 4")
	shenmo.GetShenMoService().ShenMoSceneFinish()
	return
}

func init() {
	gameevent.AddEventListener(shenmoeventtypes.EventTypeShenMoSceneFinish, event.EventListenerFunc(shenMoSceneFinish))
}
