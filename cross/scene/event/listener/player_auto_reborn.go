package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/arena"
	"fgame/fgame/cross/teamcopy/teamcopy"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//自动复活
func playerAutoReborn(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	if s == nil {
		return
	}
	//TODO 修改为跨服处理
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeArena || s.MapTemplate().GetMapType() == scenetypes.SceneTypeArenaShengShou {
		//放弃
		arena.GetArenaService().ArenaMemeberGiveUp(pl)
		//退出跨服
		pl.BackLastScene()
		return
	}

	// if s.MapTemplate().GetMapType() == scenetypes.SceneTypeCrossTeamCopy {
	// 	//放弃
	// 	teamcopy.GetTeamCopyService().TeamCopyMemeberGiveUp(pl)
	// 	//退出跨服
	// 	pl.BackLastScene()
	// 	return
	// }

	flag := scenelogic.AutoReborn(pl)
	if !flag {
		if s.MapTemplate().GetMapType() == scenetypes.SceneTypeCrossTeamCopy {
			//放弃
			teamcopy.GetTeamCopyService().TeamCopyMemeberGiveUp(pl)
			// //退出跨服
			// pl.BackLastScene()
			// return
		}
		pl.BackLastScene()
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerAutoReborn, event.EventListenerFunc(playerAutoReborn))
}
