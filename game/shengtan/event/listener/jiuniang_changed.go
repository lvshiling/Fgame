package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	shengtaneventtypes "fgame/fgame/game/shengtan/event/types"
	"fgame/fgame/game/shengtan/pbutil"
	shengtanscene "fgame/fgame/game/shengtan/scene"
)

//圣坛
func jiuNiangChanged(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(shengtanscene.ShengTanSceneData)
	if !ok {
		return
	}
	jiuNiangNum, jiuNiangPercent := sd.GetJiuNiangNum()
	//广播酒酿变化
	scShengTanSceneJiuNiangChanged := pbutil.BuildSCShengTanSceneJiuNiangChanged(jiuNiangNum, jiuNiangPercent)
	sd.GetScene().BroadcastMsg(scShengTanSceneJiuNiangChanged)

	return nil
}

func init() {
	gameevent.AddEventListener(shengtaneventtypes.EventTypeShengTanSceneJiuNiangChanged, event.EventListenerFunc(jiuNiangChanged))
}
