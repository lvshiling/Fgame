package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
	"fgame/fgame/game/lianyu/pbutil"
	lianyuscene "fgame/fgame/game/lianyu/scene"
)

//无间炼狱排行榜变化
func lianYuRankChanged(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(lianyuscene.LianYuSceneData)

	rankList := sd.GetRank()
	scLianYuRankChanged := pbutil.BuildSCLianYuRankChanged(rankList)
	sd.GetScene().BroadcastMsg(scLianYuRankChanged)
	return
}

func init() {
	gameevent.AddEventListener(lianyueventtypes.EventTypeLianYuShaQiRankChanged, event.EventListenerFunc(lianYuRankChanged))
}
