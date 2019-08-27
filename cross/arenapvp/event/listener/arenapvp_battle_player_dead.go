package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arenapvp/arenapvp"
	"fgame/fgame/cross/arenapvp/pbutil"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

// 玩家死亡
func pvpPlayerDead(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArenapvp {
		return
	}

	sd := s.SceneDelegate().(arenapvpscene.ArenapvpBattleSceneData)
	scMsg := pbutil.BuildSCArenapvpPlayerShowDataDeadChanged(pl)
	s.BroadcastMsg(scMsg)

	if sd.GetState() == arenapvpscene.ArenaSceneStateCompete {
		remain := sd.GetPvpTemp().GetRemainReliveTimes(pl.GetArenapvpReliveTimes())
		if remain <= 0 {
			pvpPlayer, flag := arenapvp.GetArenapvpService().PvpPlayerFailed(pl.GetId())
			if flag {
				scArenapvpPlayerStateChanged := pbutil.BuildSCArenapvpPlayerStateChanged(pvpPlayer)
				s.BroadcastMsg(scArenapvpPlayerStateChanged)
			}

			sd.Judge(false)
			return
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(pvpPlayerDead))
}
