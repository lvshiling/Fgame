package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arenapvp/arenapvp"
	arenapvpeventtypes "fgame/fgame/cross/arenapvp/event/types"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	arenapvpdata "fgame/fgame/game/arenapvp/data"
	gameevent "fgame/fgame/game/event"
)

//竞技场场景结束
func arenapvpElectionSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(arenapvpscene.ArenapvpSceneData)
	if !ok {
		return
	}

	winnerList, ok := data.([]*arenapvpdata.PvpPlayerInfo)
	if !ok {
		return
	}

	s := sd.GetScene()
	nextPvpTemp := sd.GetPvpTemp().GetNextTemp()
	if nextPvpTemp == nil {
		return
	}

	arenapvp.GetArenapvpService().ArenapvpElectionFinish(s, winnerList)
	return
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpElectionSceneFinish, event.EventListenerFunc(arenapvpElectionSceneFinish))
}
