package action

import (
	"fgame/fgame/game/global"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

func init() {
	scene.RegisterLingTongActionFactory(scene.LingTongStateAttack, scene.LingTongActionFactoryFunc(newAttackAction))
}

type attackAction struct {
	*scene.DummyLingTongAction
}

func (a *attackAction) Action(lingTong scene.LingTong) {
	//跟随
	if !scenelogic.CheckIfLingTongAndPlayerSameScene(lingTong) {
		lingTong.Idle()
		return
	}
	now := global.GetGame().GetTimeService().Now()
	elapse := now - lingTong.GetSkillTime()
	//没到时间
	if elapse < lingTong.GetSkillActionTime() {
		return
	}
	lingTong.Idle()
}

func newAttackAction() scene.LingTongAction {
	a := &attackAction{}
	a.DummyLingTongAction = scene.NewDummyLingTongAction()
	return a
}
