package ai

import (
	coreutils "fgame/fgame/core/utils"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
	"fgame/fgame/pkg/randomutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterDefaultAction(scene.NPCStateTrace, scene.NPCActionHandler(traceAction))
}

//追踪动作
func traceAction(n scene.NPC) {

	//超过追击范围
	if n.GetBiologyTemplate().GetRandradiusSquare() < coreutils.DistanceSquare(n.GetPosition(), n.GetBornPosition()) {
		n.Back()
		return
	}
	bo := n.GetAttackTarget()
	if bo == nil {
		n.Back()
		return
	}
	// TRICK 需要修改
	if bo.GetScene() == nil {
		n.Back()
		return
	}

	if bo.IsDead() {
		n.Back()
		return
	}

	traceTarget(n, bo)
}

// func trace(n scene.NPC, bo scene.BattleObject) (flag bool) {
// 	log.WithFields(
// 		log.Fields{
// 			"npc":   n,
// 			"enemy": bo,
// 		}).Debugln("scene.ai:追击")

// 	//计算位置
// 	destPos := bo.GetPosition() //getDestPosition(n, bo)
// 	//追击
// 	return scenelogic.MoveTo(n, destPos, 300)
// }

//随机技能
func traceTarget(n scene.NPC, bo scene.BattleObject) {
	log.WithFields(
		log.Fields{
			"npc":   n,
			"enemy": bo,
		}).Debugln("scene.ai:进行追击")

	//获取距离
	distanceSquare := coreutils.DistanceSquare(n.GetPosition(), bo.GetPosition())
	if distanceSquare > n.GetBiologyTemplate().GetMinAttackRadiusSquare() {
		log.WithFields(
			log.Fields{
				"npc":   n,
				"enemy": bo,
			}).Debugln("scene.ai:超出最小攻击范围")
		//追击
		// trace(n, bo)
		flag := n.SetDestPosition(bo.GetPosition())
		if !flag {
			log.WithFields(
				log.Fields{
					"npc":   n.GetName(),
					"pos":   bo.GetPosition(),
					"mapId": n.GetScene().MapId(),
				}).Warn("npc:追踪找不到路")
			n.SetAttackTarget(nil)
			return
		}
		return
	}

	//被限制攻击
	if n.GetBattleLimit()&scenetypes.BattleLimitTypeSkill.Mask() != 0 {
		return
	}

	log.WithFields(
		log.Fields{
			"npc":   n,
			"enemy": bo,
		}).Debugln("scene.ai:进入攻击范围")

	//获取空闲的主动技能
	positiveSkills := n.GetSkills(skilltypes.SkillSecondTypePositive)
	//过滤可以释放的技能
	var remainSkills []int32
	var weights []int
	for _, positiveSkill := range positiveSkills {
		positiveSkillId := positiveSkill.GetSkillId()
		if !n.IsSkillInCd(positiveSkillId) {
			remainSkills = append(remainSkills, positiveSkillId)
			weights = append(weights, int(n.GetBiologyTemplate().GetSkillRate(positiveSkillId)))
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
		attack(n, bo, randomSkill)
		return
	}

	log.WithFields(
		log.Fields{
			"npc":   n,
			"enemy": bo,
		}).Debugln("scene.ai:没有攻击技能可以使用")
	//判断普通攻击
	randomSkill = n.GetBiologyTemplate().GetAttackSkill()
	if randomSkill == 0 {
		log.WithFields(
			log.Fields{
				"npc":   n,
				"enemy": bo,
			}).Debugln("scene.ai:普通攻击技能不存在")
		//追击
		// trace(n, bo)
		return
	}
	randomSkillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(randomSkill)
	randomSkill = randomSkillTemplate.TypeId
	if n.IsSkillInCd(randomSkill) {
		log.WithFields(
			log.Fields{
				"npc":   n,
				"enemy": bo,
			}).Debugln("scene.ai:普通攻击技能在cd")
		return
	}
	log.WithFields(
		log.Fields{
			"npc":   n,
			"enemy": bo,
		}).Debugln("scene.ai:普通攻击技能攻击")
	//攻击
	attack(n, bo, randomSkill)
}

//攻击
func attack(n scene.NPC, bo scene.BattleObject, skillId int32) {

	//设置角度
	angle := coreutils.GetAngle(n.GetPosition(), bo.GetPosition())
	//获取场景
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	scenelogic.Attack(n, n.GetPosition(), angle, skillTemplate, true)
}
