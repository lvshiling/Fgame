package listener

import (
	"fgame/fgame/core/event"
	alliancebossscene "fgame/fgame/game/alliance/boss_scene"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//玩家进入仙盟boss
func playerEnterAllianceBossScene(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(alliancebossscene.AllianceBossSceneData)
	p := data.(player.Player)

	scAllianceBossEnter := pbutil.BuildSCAllianceBossEnter(sd)
	p.SendMsg(scAllianceBossEnter)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypePlayerEnterAllianceBossScene, event.EventListenerFunc(playerEnterAllianceBossScene))
}
