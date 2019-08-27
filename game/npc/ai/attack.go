package ai

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
)

func init() {
	scene.RegisterDefaultAction(scene.NPCStateAttack, scene.NPCActionHandler(attackAction))
}

//攻击动作
func attackAction(n scene.NPC) {
	now := global.GetGame().GetTimeService().Now()
	elapse := now - n.GetSkillTime()
	//没到时间
	if elapse < n.GetSkillActionTime() {
		return
	}
	n.Trace()
}
