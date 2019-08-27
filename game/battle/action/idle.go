package action

import (
	"fgame/fgame/game/battle/battle"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

func init() {
	scene.RegisterDefaultGuaJiActionFactory(battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

func (a *idleAction) GuaJi(p scene.Player) {

	//世界地图定制
	s := p.GetScene()
	if s == nil {
		return
	}

	e := scenelogic.FindHatestEnemy(p)

	if e != nil {
		p.SetAttackTarget(e.BattleObject)
		p.GuaJiTrace()
		return
	}

	defaultTarget := p.GetDefaultAttackTarget()
	if defaultTarget != nil {
		p.SetAttackTarget(defaultTarget)
		p.GuaJiTrace()
		return
	}
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
