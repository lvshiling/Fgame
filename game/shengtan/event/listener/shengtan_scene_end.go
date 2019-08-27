package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	shengtaneventtypes "fgame/fgame/game/shengtan/event/types"
	"fgame/fgame/game/shengtan/pbutil"
	shengtanscene "fgame/fgame/game/shengtan/scene"
	"fgame/fgame/game/shengtan/shengtan"
)

//圣坛结束
func shengTanSceneEnd(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(shengtanscene.ShengTanSceneData)
	if !ok {
		return
	}
	shengtan.GetShengTanService().ShengTanSceneClose(sd.GetAllianceId())
	scShengTanSceneEnd := pbutil.BuildSCShengTanSceneEnd()
	sd.GetScene().BroadcastMsg(scShengTanSceneEnd)
	return nil
}

func init() {
	gameevent.AddEventListener(shengtaneventtypes.EventTypeShengTanSceneEnd, event.EventListenerFunc(shengTanSceneEnd))
}
