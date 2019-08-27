package battle

import (
	coretypes "fgame/fgame/core/types"
	arenascene "fgame/fgame/cross/arena/scene"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterCheckMoveHandler(scenetypes.SceneTypeArena, scene.CheckMoveHandlerFunc(chenArenaMove))
}

// pos：移动的位置
func chenArenaMove(p scene.Player, pos coretypes.Position) (flag bool, fixPos coretypes.Position) {
	//倒计时中不能移动
	s := p.GetScene()
	if s == nil {
		return
	}
	sd := s.SceneDelegate().(arenascene.ArenaSceneData)
	if sd.GetState() == arenascene.ArenaSceneStateInit {
		return false, p.GetPos()
	}
	flag = true
	return
}
