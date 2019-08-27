package listener

import (
	"fgame/fgame/core/event"
	tulongeventtypes "fgame/fgame/cross/tulong/event/types"
	"fgame/fgame/cross/tulong/pbutil"
	tulongscene "fgame/fgame/cross/tulong/scene"
	"fgame/fgame/cross/tulong/tulong"
	gameevent "fgame/fgame/game/event"
)

//屠龙场景结束
func tuLongSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(tulongscene.TuLongSceneData)
	if !ok {
		return
	}

	//屠龙结果
	allianceMap := sd.GetAllianceMap()
	allPlayers := sd.GetScene().GetAllPlayers()
	for playerId, pl := range allPlayers {
		allianceId := pl.GetAllianceId()
		allianceInfo, exist := allianceMap[allianceId]
		if !exist {
			continue
		}
		killNum := allianceInfo.GetKillNum()
		itemMap := allianceInfo.GetItemMap(playerId)

		scTuLongResult := pbutil.BuildSCTuLongResult(killNum, itemMap)
		pl.SendMsg(scTuLongResult)
	}

	tulong.GetTuLongService().TuLongSceneFinish()
	return
}

func init() {
	gameevent.AddEventListener(tulongeventtypes.EventTypeTuLongSceneFinish, event.EventListenerFunc(tuLongSceneFinish))
}
