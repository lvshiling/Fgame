package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arenapvp/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//复活次数变化
func playerArenapvpReliveTimesChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}

	//海选
	electionScene(pl, s)

	//pvp
	battlepvpScene(pl, s)

	return
}

func electionScene(pl scene.Player, s scene.Scene) {
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArenapvpHaiXuan {
		return
	}

	scMsg := pbutil.BuildSCArenapvpElectionSceneDataChanged(pl.GetArenapvpReliveTimes())
	pl.SendMsg(scMsg)
}

func battlepvpScene(pl scene.Player, s scene.Scene) {
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArenapvp {
		return
	}

	scMsg := pbutil.BuildSCArenapvpPlayerShowDataReliveTimeChanged(pl)
	s.BroadcastMsg(scMsg)
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerArenapvpReliveTimesChanged, event.EventListenerFunc(playerArenapvpReliveTimesChanged))
}
