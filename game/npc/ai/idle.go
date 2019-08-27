package ai

import (
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

func init() {
	scene.RegisterDefaultAction(scene.NPCStateInit, scene.NPCActionHandler(idleAction))
}

func idleAction(n scene.NPC) {
	bo := n.GetForeverAttackTarget()
	if bo != nil && !bo.IsDead() {
		n.SetAttackTarget(bo)
		n.Trace()
		return
	}

	//查找敌人
	e := scenelogic.FindHatestEnemy(n)

	if e == nil {
		//查找默认目标
		bo = n.GetDefaultAttackTarget()
		if bo == nil {
			return
		}
		n.SetAttackTarget(bo)
		n.Trace()
		return
	}
	n.SetAttackTarget(e.BattleObject)
	n.Trace()
}
