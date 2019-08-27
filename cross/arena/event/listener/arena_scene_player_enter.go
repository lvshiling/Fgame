package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/arena"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	"fgame/fgame/cross/arena/pbutil"
	arenascene "fgame/fgame/cross/arena/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

// 竞技场进入
func arenaScenePlayerEnter(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(arenascene.ArenaSceneData)
	pl := data.(scene.Player)
	scArenaSceneInfo := pbutil.BuildSCArenaSceneInfo(sd)
	pl.SendMsg(scArenaSceneInfo)

	t := arena.GetArenaService().GetArenaTeamByPlayerId(pl.GetId())
	if t == nil {
		return
	}

	//发送竞技场信息
	scPlayerArenaData := pbutil.BuildSCPlayerArenaData(pl)
	pl.SendMsg(scPlayerArenaData)
	scArenaPlayerDataEnterSceneChanged := pbutil.BuildSCArenaPlayerDataEnterSceneChanged(pl)
	sd.GetScene().BroadcastMsg(scArenaPlayerDataEnterSceneChanged)

	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaScenePlayerEnter, event.EventListenerFunc(arenaScenePlayerEnter))
}
