package fuben

import (
	"fgame/fgame/game/battle/battle"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeFuBen, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

//副本挂机
func (a *idleAction) GuaJi(p scene.Player) {
	s := p.GetScene()
	//结束了
	if s.IsFinish() {
		p.BackLastScene()
		return
	}

	e := scenelogic.FindHatestEnemy(p)

	if e == nil {
		//查找默认目标
		bo := p.GetDefaultAttackTarget()
		if bo != nil {
			p.SetAttackTarget(bo)
			p.GuaJiTrace()
			return
		}
		return
	}
	p.SetAttackTarget(e.BattleObject)
	p.GuaJiTrace()
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
