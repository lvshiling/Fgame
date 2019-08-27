package relive_handler

import (
	arenapvplogic "fgame/fgame/cross/arenapvp/logic"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterAutoReliveHandler(scenetypes.SceneTypeArenapvp, scene.AutoReliveHandlerFunc(autoRelive))
	scene.RegisterAutoReliveHandler(scenetypes.SceneTypeArenapvpHaiXuan, scene.AutoReliveHandlerFunc(autoRelive))
}

//pvp复活处理
func autoRelive(pl scene.Player) bool {
	s := pl.GetScene()
	if s == nil {
		return false
	}

	reliveTimes := pl.GetArenapvpReliveTimes()
	switch tsd := s.SceneDelegate().(type) {
	case arenapvpscene.ArenapvpSceneData:
		{
			remainTimes := tsd.GetPvpTemp().GetRemainReliveTimes(reliveTimes)
			if remainTimes > 0 {
				arenapvplogic.Reborn(pl)
				return true
			}
		}
	case arenapvpscene.ArenapvpBattleSceneData:
		{
			remainTimes := tsd.GetPvpTemp().GetRemainReliveTimes(reliveTimes)
			if remainTimes > 0 {
				arenapvplogic.Reborn(pl)
				return true
			}
		}
	}

	return false
}
