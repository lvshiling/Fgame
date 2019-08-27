package ai

import (
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/global"
	"fgame/fgame/game/robot/robot"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
	"fgame/fgame/pkg/randomutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	robot.RegisterDefaultActionFactory(robot.RobotPlayerStateTrace, robot.RobotActionFactoryFunc(newTraceAction))
}

type traceAction struct {
	*robot.DummyAction
	lastTime int64
}

func (a *traceAction) OnEnter() {
	now := global.GetGame().GetTimeService().Now()
	a.lastTime = now
	return
}

func (a *traceAction) OnExit() {
	return
}
func (a *traceAction) Action(p scene.RobotPlayer) {
	bo := p.GetAttackTarget()
	if bo == nil {
		p.Idle()
		return
	}

	traceTarget(p, bo)

	return
}

func newTraceAction() scene.RobotAction {
	a := &traceAction{}
	a.DummyAction = robot.NewDummyAction()
	return a
}

// func trace(p scene.RobotPlayer, bo scene.BattleObject, elapse int64) (flag bool) {
// 	//计算位置
// 	destPos := bo.GetPosition() //getDestPosition(n, bo)
// 	//追击
// 	return scenelogic.MoveTo(p, destPos, elapse)
// }

//随机技能
func traceTarget(p scene.RobotPlayer, bo scene.BattleObject) {

	//获取距离
	distanceSquare := coreutils.DistanceSquare(p.GetPosition(), bo.GetPosition())
	minAttackDistanceSquare := float64(1.0)
	if distanceSquare > minAttackDistanceSquare {
		flag := p.SetDestPosition(bo.GetPosition())
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
				}).Warn("robot:机器人找不到路")
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

//攻击
func attack(n scene.RobotPlayer, bo scene.BattleObject, skillId int32) {
	//设置角度
	angle := coreutils.GetAngle(n.GetPosition(), bo.GetPosition())
	//获取场景
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	scenelogic.Attack(n, n.GetPosition(), angle, skillTemplate, true)
}
