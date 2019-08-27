package action

import (
	coreutils "fgame/fgame/core/utils"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
	"fgame/fgame/pkg/randomutils"
)

func init() {
	scene.RegisterLingTongActionFactory(scene.LingTongStateTrace, scene.LingTongActionFactoryFunc(newTraceAction))
}

type traceAction struct {
	*scene.DummyLingTongAction
}

func (a *traceAction) Action(lingTong scene.LingTong) {
	//跟随
	if !scenelogic.CheckIfLingTongAndPlayerSameScene(lingTong) {
		return
	}
	owner := lingTong.GetOwner()
	if owner.GetAttackTarget() != nil {
		traceObject(lingTong, owner.GetAttackTarget())
		return
	}

	//判断是不是超过近身距离
	if coreutils.DistanceSquare(owner.GetPosition(), lingTong.GetPosition()) <= maxDistanceSquare {
		lingTong.Idle()
		return
	}

	flag := lingTong.SetDestPosition(getLingTongPosition(owner))
	if !flag {
		scenelogic.FixPosition(lingTong, owner.GetPosition())
		return
	}
	return
}

const (
	attackMinDistance = 1
)

func traceObject(lingTong scene.LingTong, bo scene.BattleObject) {
	//获取距离
	distanceSquare := coreutils.DistanceSquare(lingTong.GetPosition(), bo.GetPosition())
	if distanceSquare > attackMinDistance {
		flag := lingTong.SetDestPosition(bo.GetPosition())
		if !flag {
			return
		}
		return
	}

	//获取空闲的主动技能
	positiveSkills := lingTong.GetSkills(skilltypes.SkillSecondTypePositive)
	//过滤可以释放的技能
	var remainSkills []int32
	var weights []int
	for _, positiveSkill := range positiveSkills {
		positiveSkillId := positiveSkill.GetSkillId()
		if !lingTong.IsSkillInCd(positiveSkillId) {
			remainSkills = append(remainSkills, positiveSkillId)
			weights = append(weights, int(1))
		}
	}
	if len(weights) <= 0 {
		return
	}

	var randomSkill int32
	//随机技能

	skillIndex := randomutils.RandomWeights(weights)
	randomSkill = remainSkills[skillIndex]

	angle := coreutils.GetAngle(lingTong.GetPosition(), bo.GetPosition())
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(randomSkill)
	//攻击
	scenelogic.LingTongAttack(lingTong, bo.GetPosition(), angle, skillTemplate, true)
}

func newTraceAction() scene.LingTongAction {
	a := &traceAction{}
	a.DummyLingTongAction = scene.NewDummyLingTongAction()
	return a
}
