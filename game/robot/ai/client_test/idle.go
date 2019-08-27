package client_test

import (
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/robot/robot"
	robottypes "fgame/fgame/game/robot/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	skillcommon "fgame/fgame/game/skill/common"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
	"fgame/fgame/pkg/randomutils"
	"math"
	"math/rand"
)

func init() {
	robot.RegisterActionFactory(robottypes.RobotTypeTest, robot.RobotPlayerStateIdle, robot.RobotActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*robot.DummyAction
}

func (a *idleAction) Action(p scene.RobotPlayer) {
	//随机范围走动和放技能
	s := p.GetScene()
	if s == nil {
		return
	}
	isAttack := rand.Float64() > 0.5
	if isAttack {
		allSkills := p.GetSkills(skilltypes.SkillSecondTypePositive)
		tempSkills := make([]skillcommon.SkillObject, 0, len(allSkills))
		weights := make([]int, 0, len(tempSkills))
		for _, ski := range allSkills {
			if p.IsSkillInCd(ski.GetSkillId()) {
				continue
			}
			tempSkills = append(tempSkills, ski)
			weights = append(weights, 1)
		}
		//判断是否有技能
		var randomSkill skillcommon.SkillObject
		//随机技能
		if len(weights) > 0 {
			skillIndex := randomutils.RandomWeights(weights)
			if skillIndex >= 0 {
				randomSkill = tempSkills[skillIndex]
			}
		}

		if randomSkill != nil {
			skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(randomSkill.GetSkillId())
			scenelogic.Attack(p, p.GetPos(), p.GetAngle(), skillTemplate, true)
			return
		}
	}
	t := 100
	//随机角度
	angle := rand.Float64() * math.Pi * 2
	//算出位置
	moveSpeed := float64(p.GetBattleProperty(propertytypes.BattlePropertyTypeMoveSpeed)) / float64(common.MILL_METER)
	distance := moveSpeed * float64(t) / float64(int64(common.SECOND))
	destPos := coretypes.Position{
		X: p.GetPosition().X + math.Cos(angle)*distance,
		Y: p.GetPosition().Y,
		Z: p.GetPosition().Z + math.Sin(angle)*distance,
	}
	if !s.MapTemplate().GetMap().IsMask(destPos.X, destPos.Z) {
		return
	}
	destPos.Y = s.MapTemplate().GetMap().GetHeight(destPos.X, destPos.Z)
	scenelogic.Move(p, destPos, angle, moveSpeed, scenetypes.MoveTypeNormal, false, false)
	return
}

func newIdleAction() scene.RobotAction {
	a := &idleAction{}
	a.DummyAction = robot.NewDummyAction()
	return a
}
