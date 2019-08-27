package action

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
)

func init() {
	scene.RegisterDefaultGuaJiActionFactory(battle.PlayerStateAttack, scene.GuaJiActionFactoryFunc(newAttackAction))
}

type attackAction struct {
	*scene.DummyGuaJiAction
}

func (a *attackAction) GuaJi(p scene.Player) {
	now := global.GetGame().GetTimeService().Now()
	elapse := now - p.GetSkillTime()
	//没到时间
	if elapse < p.GetSkillActionTime() {
		return
	}
	p.GuaJiTrace()
}

func newAttackAction() scene.GuaJiAction {
	a := &attackAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
