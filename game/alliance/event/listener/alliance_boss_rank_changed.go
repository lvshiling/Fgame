package listener

import (
	"fgame/fgame/core/event"
	alliancebossscene "fgame/fgame/game/alliance/boss_scene"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	gameevent "fgame/fgame/game/event"
)

//仙盟boss伤害排名变更
func allianceBossRankChanged(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancebossscene.AllianceBossSceneData)
	s := sd.GetScene()
	if s == nil {
		return
	}
	scAllianceBossRank := pbutil.BuildSCAllianceBossRank(sd)
	s.BroadcastMsg(scAllianceBossRank)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceBossRankChanged, event.EventListenerFunc(allianceBossRankChanged))
}
