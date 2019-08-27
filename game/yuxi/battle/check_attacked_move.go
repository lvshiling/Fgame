package battle

import (
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	yuxiscene "fgame/fgame/game/yuxi/scene"
	yuxitypes "fgame/fgame/game/yuxi/types"
)

func init() {
	scene.RegisterCheckAttackedMoveHandler(scenetypes.SceneTypeYuXi, scene.CheckAttackedMoveHandlerFunc(yuXiCheckAttackedMove))
}

func yuXiCheckAttackedMove(attackObject scene.BattleObject, defenceObject scene.BattleObject, pos coretypes.Position) bool {
	spl, ok := defenceObject.(scene.Player)
	if !ok {
		return false
	}

	s := attackObject.GetScene()
	if s == nil {
		return false
	}

	sd, ok := s.SceneDelegate().(yuxiscene.YuXiSceneData)
	if !ok {
		return false
	}

	if !sd.GetScene().MapTemplate().IsSafe(pos) {
		return true
	}

	owner, _ := sd.GetOwerYuXinInfo()
	if owner == spl {
		sd.YuXiRest(spl, yuxitypes.YuXiReborTypePlayerEnterSafe)
	}

	return true
}
