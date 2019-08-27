package check_attack

import (
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterCheckAttackHandler(scenetypes.SceneTypeArenapvpHaiXuan, scene.CheckAttackHandlerFunc(checkAttack))
}

//比武大会海选 攻击无效
func checkAttack(attackObj scene.BattleObject) (isAttack bool) {

	s := attackObj.GetScene()
	if s == nil {
		return
	}

	// 场景完成状态攻击无效
	if s.IsFinish() {
		return
	}

	isAttack = true
	return
}
