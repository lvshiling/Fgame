package battle

import (
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	yuxiscene "fgame/fgame/game/yuxi/scene"
	yuxitypes "fgame/fgame/game/yuxi/types"
)

func init() {
	scene.RegisterCheckMoveHandler(scenetypes.SceneTypeYuXi, scene.CheckMoveHandlerFunc(yuXiCheckMove))
}

func yuXiCheckMove(p scene.Player, pos coretypes.Position) (flag bool, fixPos coretypes.Position) {
	s := p.GetScene()
	if s == nil {
		flag = false
		return
	}

	sd, ok := s.SceneDelegate().(yuxiscene.YuXiSceneData)
	if !ok {
		flag = false
		return
	}

	if !sd.GetScene().MapTemplate().IsSafe(pos) {
		flag = true
		return
	}

	sd.YuXiRest(p, yuxitypes.YuXiReborTypePlayerEnterSafe)

	flag = true
	return
}
