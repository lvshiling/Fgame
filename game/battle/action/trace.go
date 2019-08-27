package action

import (
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/battle/battle"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
	"fgame/fgame/pkg/randomutils"
)

func init() {
	scene.RegisterDefaultGuaJiActionFactory(battle.PlayerStateTrace, scene.GuaJiActionFactoryFunc(newTraceAction))
}

type traceAction struct {
	*scene.DummyGuaJiAction
}

func (a *traceAction) GuaJi(p scene.Player) {

	bo := p.GetAttackTarget()
	if bo == nil {
		p.GuaJiIdle()
		return
	}
	if bo.IsDead() {
		p.SetAttackTarget(nil)
		p.GuaJiIdle()
		return
	}
	s := p.GetScene()
	if s == nil {
		p.SetAttackTarget(nil)
		p.GuaJiIdle()
		return
	}

	targetScene := bo.GetScene()
	if targetScene == nil {
		p.SetAttackTarget(nil)
		p.GuaJiIdle()
		return
	}
	if s != targetScene {
		flag := scenelogic.PlayerEnterScene(p, targetScene, targetScene.MapTemplate().GetBornPos())
		if !flag {
			p.SetAttackTarget(nil)
			p.GuaJiIdle()
			return
		}
		return
	}

	isPlayer := bo.GetSceneObjectType() == scenetypes.BiologyTypePlayer
	if isPlayer {
		//检查是否在安全区内
		if s.MapTemplate().IsSafe(bo.GetPosition()) {
			p.SetAttackTarget(nil)
			p.GuaJiIdle()
			return
		}
	}
	traceTarget(p, bo)

	return
}

func newTraceAction() scene.GuaJiAction {
	a := &traceAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}

//随机技能
func traceTarget(p scene.Player, bo scene.BattleObject) {

	//获取距离
	distanceSquare := coreutils.DistanceSquare(p.GetPosition(), bo.GetPosition())
	minAttackDistanceSquare := float64(1.0)
	if distanceSquare > minAttackDistanceSquare {
		flag := p.SetDestPosition(bo.GetPosition())
		if !flag {
			p.SetAttackTarget(nil)
			return
		}
		return
	}

	//被限制攻击
	if p.GetBattleLimit()&scenetypes.BattleLimitTypeSkill.Mask() != 0 {
		return
	}

	positiveSkills := p.GetSkills(skilltypes.SkillSecondTypePositive)

	//过滤可以释放的技能
	var remainSkills []int32
	var weights []int
	for _, positiveSkill := range positiveSkills {
		positiveSkillId := positiveSkill.GetSkillId()
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(positiveSkillId)
		//排除跳跃技能
		if skillTemplate.GetSkillFirstType() == skilltypes.SkillFirstTypeJump {
			continue
		}

		if !p.IsSkillInCd(positiveSkillId) {
			remainSkills = append(remainSkills, positiveSkillId)
			weights = append(weights, 1)
		}
	}

	var randomSkill int32
	//随机技能
	if len(weights) > 0 {
		skillIndex := randomutils.RandomWeights(weights)
		if skillIndex >= 0 {
			randomSkill = remainSkills[skillIndex]
		}
	}

	//随机技能
	if randomSkill > 0 {
		//攻击
		attack(p, bo, randomSkill)
		return
	}

}

// //攻击
func attack(n scene.Player, bo scene.BattleObject, skillId int32) {
	//设置角度
	angle := coreutils.GetAngle(n.GetPosition(), bo.GetPosition())
	//获取场景
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	scenelogic.Attack(n, n.GetPosition(), angle, skillTemplate, true)
}
