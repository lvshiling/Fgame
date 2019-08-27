package logic

import (
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/lang"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	playerlogic "fgame/fgame/game/player/logic"
	propertytypes "fgame/fgame/game/property/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
	"fgame/fgame/game/soul/soul"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"math"
	"math/rand"

	log "github.com/Sirupsen/logrus"
)

func HandleObjectAttack(pl scene.Player, uid int64, pos coretypes.Position, angle float64, skillId int32) {
	if uid == pl.GetId() {
		HandlePlayerAttack(pl, pos, angle, skillId)
		return
	}
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"uId":      uid,
				"pos":      pos,
				"angle":    angle,
				"skillId":  skillId,
			}).Warn("scene:处理对象移动消息,场景为空")
		return
	}
	lingTong := s.GetLingTong(uid)
	if lingTong == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"uId":      uid,
				"pos":      pos,
				"angle":    angle,
			}).Warn("scene:处理对象移动消息,灵童不存在")
		return
	}
	if lingTong.GetOwner() != pl {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"uId":      uid,
				"pos":      pos,
				"angle":    angle,
			}).Warn("scene:处理对象移动消息,灵童不是他的")
		return
	}
	HandleLingTongAttack(lingTong, pos, angle, skillId)
}

//攻击
func HandlePlayerAttack(pl scene.Player, pos coretypes.Position, angle float64, skillId int32) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
				"angle":    angle,
				"skillId":  skillId,
			}).Warn("scene:处理对象攻击消息,场景为空")
		return
	}

	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	if skillTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"pos":      pos,
			"angle":    angle,
			"skillId":  skillId,
		}).Warn("scene:处理对象攻击消息,技能不存在")
		return
	}
	//被动技能
	if skillTemplate.IsPassive() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"pos":      pos,
			"angle":    angle,
			"skillId":  skillId,
		}).Warnln("scene:处理对象攻击消息,被动技能")
		return
	}

	flag := Attack(pl, pos, angle, skillTemplate, true)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"pos":      pos,
			"angle":    angle,
			"skillId":  skillId,
		}).Warnln("scene:处理对象攻击消息,攻击失败")
		return
	}
}

//攻击
func HandleLingTongAttack(lingTong scene.LingTong, pos coretypes.Position, angle float64, skillId int32) {

	pl := lingTong.GetOwner()
	if pl == nil {
		log.WithFields(
			log.Fields{
				"playerId": lingTong.GetOwner().GetId(),
				"pos":      pos,
				"angle":    angle,
				"skillId":  skillId,
			}).Warn("scene:处理灵童攻击消息,主角为空")
		return
	}
	s := lingTong.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": lingTong.GetOwner().GetId(),
				"pos":      pos,
				"angle":    angle,
				"skillId":  skillId,
			}).Warn("scene:处理灵童攻击消息,场景为空")
		return
	}
	ownerScene := pl.GetScene()
	if ownerScene != s {
		log.WithFields(
			log.Fields{
				"playerId": lingTong.GetOwner().GetId(),
				"pos":      pos,
				"angle":    angle,
				"skillId":  skillId,
			}).Warn("scene:处理灵童攻击消息,场景不一致")
		return
	}
	attackId := lingTong.GetLingTongTemplate().AttackId
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(attackId)
	if skillTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"pos":      pos,
			"angle":    angle,
			"attackId": attackId,
		}).Warn("scene:处理灵童攻击消息,技能不存在")
		return
	}
	//被动技能
	if skillTemplate.IsPassive() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"pos":      pos,
			"angle":    angle,
			"attackId": attackId,
		}).Warnln("scene:处理灵童攻击消息,被动技能")
		return
	}

	flag := LingTongAttack(lingTong, pos, angle, skillTemplate, true)
	if !flag {
		// log.WithFields(log.Fields{
		// 	"playerId": pl.GetId(),
		// 	"pos":      pos,
		// 	"angle":    angle,
		// 	"attackId": attackId,
		// }).Warnln("scene:处理灵童攻击消息,攻击失败")
		return
	}
	return
}

//获取命中
func getHit(bo scene.BattleObject) int64 {
	hit := bo.GetBattleProperty(propertytypes.BattlePropertyTypeHit)
	hitPercent := bo.GetBattleProperty(propertytypes.BattlePropertyTypeHitRatePercent)
	hit += int64(math.Ceil(float64(hit) * (float64(hitPercent) / float64(common.MAX_RATE))))
	return hit
}

//获取闪避
func getDodge(bo scene.BattleObject) int64 {
	dodge := bo.GetBattleProperty(propertytypes.BattlePropertyTypeDodge)
	dodgePercent := bo.GetBattleProperty(propertytypes.BattlePropertyTypeDodgeRatePercent)
	dodge += int64(math.Ceil(float64(dodge) * (float64(dodgePercent) / float64(common.MAX_RATE))))
	return dodge
}

//获取暴击
func getCrit(bo scene.BattleObject) int64 {
	crit := bo.GetBattleProperty(propertytypes.BattlePropertyTypeCrit)

	return crit
}

//获取坚韧
func getTough(bo scene.BattleObject) int64 {
	tough := bo.GetBattleProperty(propertytypes.BattlePropertyTypeTough)

	return tough
}

func Attack(attackObject scene.BattleObject, pos coretypes.Position, angle float64, skillTemplate *gametemplate.SkillTemplate, active bool) (flag bool) {

	skillId := int32(skillTemplate.TypeId)

	//判断是否有这个技能
	if skillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeSubSkill && skillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeCastingSoul {
		ski := attackObject.GetSkill(skillId)
		if ski == nil {
			log.WithFields(log.Fields{
				"attackObject": attackObject,
				"skillId":      skillId,
			}).Warnln("scene: 技能不存在")
			return false
		}
	}

	//判断是否在cd中
	if attackObject.IsSkillInCd(skillId) {
		log.WithFields(log.Fields{
			"attackObject": attackObject,
			"skillId":      skillId,
		}).Warnln("scene: skill in cd")
		return false
	}

	if skillTemplate.GetSkillFirstType() == skilltypes.SkillFirstTypeJump {
		//不能使用轻功
		if attackObject.GetBattleLimit()&scenetypes.BattleLimitTypeQingGong.Mask() != 0 {
			log.WithFields(log.Fields{
				"attackObject": attackObject,
				"skillId":      skillId,
			}).Warnln("scene: 当前限制使用轻功")
			return false
		}
	}

	//不能使用技能
	if attackObject.GetBattleLimit()&scenetypes.BattleLimitTypeSkill.Mask() != 0 {
		log.WithFields(log.Fields{
			"attackObject": attackObject,
			"skillId":      skillId,
		}).Warnln("scene: 当前限制使用技能")
		return false
	}

	//使用技能
	if !attackObject.UseSkill(skillId) {
		panic(fmt.Errorf("scene:使用技能应该成功"))
	}

	//坐骑隐藏
	switch attackObj := attackObject.(type) {
	case scene.Player:
		if !attackObj.IsMountHidden() {
			attackObj.MountHidden(true)
		}
		break
	}

	costHpValue := int64(skillTemplate.CostHpValue)
	costHpPercent := int64(skillTemplate.CostHpPersent)
	if costHpValue > 0 || costHpPercent > 0 {
		percentHpValue := int64(math.Ceil(float64(costHpPercent) / float64(common.MAX_RATE) * float64(attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP))))
		totalCostHp := costHpValue + percentHpValue
		if totalCostHp > attackObject.GetHP() {
			totalCostHp = attackObject.GetHP() - 1
		}
		if totalCostHp > 0 {

			dead := attackObject.CostHP(totalCostHp, 0)
			if dead {
				panic(fmt.Errorf("scene:不能自杀"))
			}
		}
	}

	//消耗体力
	costTpValue := int64(skillTemplate.CostTpValue)
	if costTpValue > 0 {
		flag := attackObject.CostTP(costTpValue)
		if !flag {
			log.WithFields(log.Fields{
				"attackObject": attackObject,
				"skillId":      skillId,
			}).Warnln("scene: 体力不足")
			return false
		}
	}

	switch attackObj := attackObject.(type) {
	case scene.Player:
		scScenePlayerSkillUse := pbutil.BuildSCScenePlayerSkillUse(attackObj, int32(skillTemplate.Id))
		attackObj.SendMsg(scScenePlayerSkillUse)
		break
	}
	//校正
	Move(attackObject, pos, angle, 0, scenetypes.MoveTypeNormal, true, false)

	//遍历所有对象
	if active {
		//发送攻击
		scObjectAttack := pbutil.BuildSCObjectAttack(attackObject, pos, angle, skillId)
		BroadcastNeighbor(attackObject, scObjectAttack)
		gameevent.Emit(sceneeventtypes.EventTypeBattleObjectAttack, attackObject, nil)
	}

	if skillTemplate.BuffId3 != 0 {
		AddBuff(attackObject, int32(skillTemplate.BuffId3), attackObject.GetId(), int64(skillTemplate.BuffId3Rate))
	}
	//添加buff
	if skillTemplate.BuffId != 0 {
		AddBuff(attackObject, int32(skillTemplate.BuffId), attackObject.GetId(), common.MAX_RATE)
	}

	return true
}

func PetAttack(pl scene.Player, objType int32, pos coretypes.Position, angle float64, skillId int32, active bool) (flag bool) {

	scObjectAttack := pbutil.BuildSCPetAttack(pl, objType, pos, angle, skillId)
	BroadcastNeighbor(pl, scObjectAttack)
	for _, nei := range pl.GetNeighbors() {
		bo, ok := nei.(scene.BattleObject)
		if !ok {
			continue
		}
		//死亡
		if bo.IsDead() {
			continue
		}
		if !pl.IsEnemy(bo) {
			continue
		}
		damageInt := int64(1)
		damageType := scenetypes.DamageTypePet
		CostHP(bo, damageInt, skillId, pl.GetId(), damageType)
		// bo.CostHP(damageInt, pl.GetId())

		// scObjectDamage := pbutil.BuildSCObjectDamage(bo, damageType, -damageInt, skillId, pl.GetId())
		// BroadcastNeighborIncludeSelf(bo, scObjectDamage)

	}
	return true
}

//灵童攻击
func LingTongAttack(attackObject scene.LingTong, pos coretypes.Position, angle float64, skillTemplate *gametemplate.SkillTemplate, active bool) (flag bool) {
	if !CheckIfLingTongAndPlayerSameScene(attackObject) {
		log.WithFields(log.Fields{
			"attackObject": attackObject,
		}).Warnln("scene: 不在同一个场景")
		return
	}
	owner := attackObject.GetOwner()
	skillId := int32(skillTemplate.TypeId)
	//判断是否有这个技能
	if skillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeSubSkill {
		//判断是否有这个技能
		ski := attackObject.GetSkill(skillId)
		if ski == nil {
			log.WithFields(log.Fields{
				"attackObject": attackObject,
				"skillId":      skillId,
			}).Warnln("scene: 技能不存在")
			return false
		}
	}

	//判断是否在cd中
	if attackObject.IsSkillInCd(skillId) {
		// log.WithFields(log.Fields{
		// 	"attackObject": attackObject,
		// 	"skillId":      skillId,
		// }).Warnln("scene: skill in cd")
		return false
	}

	//使用技能
	if !attackObject.UseSkill(skillId) {
		panic(fmt.Errorf("scene:使用技能应该成功"))
	}

	//校正
	Move(attackObject, pos, angle, 0, scenetypes.MoveTypeNormal, true, false)

	//遍历所有对象
	if active {
		//发送攻击
		scObjectAttack := pbutil.BuildSCObjectAttack(attackObject, pos, angle, skillId)
		BroadcastNeighbor(attackObject, scObjectAttack)
		gameevent.Emit(sceneeventtypes.EventTypeBattleObjectAttack, attackObject, nil)
	}

	if skillTemplate.BuffId3 != 0 {
		AddBuff(owner, int32(skillTemplate.BuffId3), owner.GetId(), int64(skillTemplate.BuffId3Rate))
	}
	//添加buff
	if skillTemplate.BuffId != 0 {
		AddBuff(attackObject, int32(skillTemplate.BuffId), attackObject.GetId(), common.MAX_RATE)
	}

	return true
}

func CalculateLingTongAttack(attackObject scene.LingTong, pos coretypes.Position, angle float64, skillTemplate *gametemplate.SkillTemplate, active bool) (flag bool) {
	if !CheckIfLingTongAndPlayerSameScene(attackObject) {
		log.WithFields(log.Fields{
			"attackObject": attackObject,
		}).Warnln("scene: 不在同一个场景")
		return
	}

	owner := attackObject.GetOwner()
	sb := attackObject.GetScene()

	skillId := int32(skillTemplate.TypeId)
	level := int32(1)
	//判断是否有这个技能
	if skillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeSubSkill {
		//判断是否有这个技能
		ski := attackObject.GetSkill(skillId)
		if ski == nil {
			log.WithFields(log.Fields{
				"attackObject": attackObject,
				"skillId":      skillId,
			}).Warnln("scene: 技能不存在")
			return false
		}
		level = ski.GetLevel()
	}

	skillTemplate = skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(skillId, level)
	specialEffectType, specialTarget, specialDistance, specialAnimationTime, specialTime := GetSpecialEffect(attackObject, skillTemplate)
	//技能区域检测
	sos := getAttackBattleObjectList(attackObject, skillTemplate)
	skillFourthType := skillTemplate.GetSkillFourthType()
	s := attackObject.GetScene()
	switch skillFourthType {
	case skilltypes.SkillFourthTypeAttack:

		damagePersent := int32(0)
		damageAttack := float64(0)
		damageValueBase := int32(0)
		if skillTemplate.IsDynamic() {
			skillLevelTemplate := skillTemplate.GetSkillByLevel(level)
			damageValueBase = skillLevelTemplate.DamageValueBase
			damageAttack = skillLevelTemplate.GetDamageAttack()
		} else {
			damagePersent = skillTemplate.DamagePersent
			damageAttack = skillTemplate.GetDamageAttack()
			damageValueBase = skillTemplate.DamageValueBase
		}

		//判断帝魂加成
		if skillTemplate.GetSkillFirstType() == skilltypes.SkillFirstTypeGuHun {
			switch attackObj := attackObject.(type) {
			case scene.Player:
				soulChainTemplate := soul.GetSoulService().GetSoulChainTemplate(attackObj.GetSoulAwakenNum())
				if soulChainTemplate != nil {
					guHunPercent := soulChainTemplate.Value
					damageAttack += float64(guHunPercent) / float64(common.MAX_RATE)
				}
				break
			}
		}

		//遍历作用目标
		for _, defenceObject := range sos {
			//复活保护
			if defenceObject.GetBattleLimit()&scenetypes.BattleLimitTypeNoAttacked.Mask() != 0 {
				playerlogic.SendSystemMessage(owner, lang.BattleProtectNoAttacked)
				continue
			}

			switch defenceObj := defenceObject.(type) {
			case scene.NPC:
				{
					h := scene.GetCheckNPCAttackHandler(defenceObj.GetBiologyTemplate().GetBiologyScriptType())
					if h != nil {
						flag := h.CheckAttack(attackObject, defenceObj)
						if !flag {
							continue
						}
					}
				}
			case scene.Player:

				// 检查攻击
				h := scene.GetCheckAttackHandler(s.MapTemplate().GetMapType())
				if h != nil {
					if !h.CheckAttack(attackObject) {
						continue
					}
				}

				//判断安全区
				if s.MapTemplate().IsSafe(attackObject.GetPosition()) {
					log.WithFields(log.Fields{
						"attackObject": attackObject,
						"skillId":      skillId,
					}).Warnln("scene: 攻方安全区")
					continue
				}
				//判断安全区
				if s.MapTemplate().IsSafe(defenceObj.GetPosition()) {
					log.WithFields(log.Fields{
						"defenceObj": defenceObj,
						"skillId":    skillId,
					}).Warnln("scene: 防御方安全区")
					continue
				}

				//过场保护
				if defenceObj.GetBattleLimit()&scenetypes.BattleLimitTypeAttackedChangeScene.Mask() != 0 {
					playerlogic.SendSystemMessage(owner, lang.BattleProtectChangeScene)
					continue
				}
				//复活保护
				if defenceObj.GetBattleLimit()&scenetypes.BattleLimitTypeAttackedRelive.Mask() != 0 {
					playerlogic.SendSystemMessage(owner, lang.BattleProtectRelive)
					continue
				}

				//复活保护
				if defenceObj.GetBattleLimit()&scenetypes.BattleLimitTypeNoAttacked.Mask() != 0 {
					playerlogic.SendSystemMessage(owner, lang.BattleProtectPK)
					continue
				}

				//pk保护
				protectedLevel := s.MapTemplate().ProtectLevel
				if defenceObj.GetLevel() < protectedLevel {
					playerlogic.SendSystemMessage(owner, lang.PkStateProtectCanNotAttack, fmt.Sprintf("%d", protectedLevel))
					continue
				}
			}

			if skillTemplate.TargetAction&skilltypes.SkillTargetActionTypePlayer.Mask() == 0 {
				//无效
				var scObjectDamage *scenepb.SCObjectDamage
				if active {
					scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeWuXiao, 0, skillId, attackObject.GetId())
				} else {
					scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeWuXiao, 0, 0, attackObject.GetId())
				}
				BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)
				continue
			}

			//算命中
			hit := getHit(attackObject) //attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeHit)

			dodge := getDodge(defenceObject) //defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeDodge)
			hit -= dodge
			if hit < 0 {
				hit = 0
			}
			hitFlag := mathutils.RandomHit(common.MAX_RATE, int(hit))
			//随机
			if !hitFlag {
				log.WithFields(log.Fields{
					"attackObject":  attackObject,
					"defenceObject": defenceObject,
					"skillId":       skillId,
					"pos":           pos,
					"angle":         angle,
					"active":        active,
					"hit":           hit,
					"dodge":         dodge,
				}).Debug("scene: 闪避了")
				//发送闪避
				var scObjectDamage *scenepb.SCObjectDamage
				if active {
					scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeDodge, 0, skillId, attackObject.GetId())
				} else {
					scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeDodge, 0, 0, attackObject.GetId())
				}
				BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)
				continue
			}
			//命中了
			//算出伤害
			damage := float64(0)
			var damageInt int64
			var totalDamage int64
			damageType := scenetypes.DamageTypeLingTong
			var crit int64
			var tough int64
			var critRate float64
			var isCrit bool
			var block int64
			var isBlock bool
			var tbreak int64
			var blockRate float64
			var fanTan int64
			var fanTanPercent int64
			var totalFanTan int64
			//var specialEffectType skilltypes.SkillSpecialEffectType
			var isDead bool
			// var scObjectDamage *scenepb.SCObjectDamage
			var isBaseDamage bool
			defendNPC, ok := defenceObject.(scene.NPC)
			if ok {
				if defendNPC.GetBiologyTemplate().BaseDamge != 0 {
					damageInt = int64(defendNPC.GetBiologyTemplate().BaseDamge)
					isBaseDamage = true
					goto AfterDamage
				}
			}

			if damagePersent == 0 {
				ownerPercent := attackObject.GetLingTongTemplate().PlayerAttackPercent
				ownerAttack := int64(math.Floor(float64(owner.GetBattleProperty(propertytypes.BattlePropertyTypeAttack)) * float64(ownerPercent) / float64(common.MAX_RATE)))
				damage = scene.GetDamage(attackObject, ownerAttack, defenceObject, damageAttack, float64(damageValueBase))
				//至少为1
				if damage <= 0 {
					damage = 1
				}
				//计算浮动效果

				//上下浮动0.1
				wave := rand.Float64()/10 - 0.05
				damage *= (1 + wave)
			} else {
				defenceMaxHP := attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
				damage = float64(damagePersent) / common.MAX_RATE * float64(defenceMaxHP)
			}
			//判断暴击
			crit = getCrit(attackObject) //attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeCrit)
			tough = getTough(defenceObject)
			critRate = 0
			if crit != 0 {
				//TODO 暴击率计算
				critRate = (float64(crit*crit) * 2.75) / float64(crit*crit*2+tough*tough*60)
			}
			critRate += float64(skillTemplate.AddCritical) / float64(common.MAX_RATE)
			//增加攻击暴击几率万分比
			critRate += float64(attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeCritRatePercent)) / float64(common.MAX_RATE)

			isCrit = mathutils.RandomOneHit(critRate)

			if isCrit {
				critRatio := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCrit)) / common.MAX_RATE
				//增加暴击伤害万分比
				critRatio += float64(attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeCritHarmPercent)) / float64(common.MAX_RATE)
				damage *= critRatio
			}

			//判断格挡
			block = defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeBlock)
			tbreak = attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeAbnormality)
			blockRate = 0
			if block != 0 {
				blockRate = (float64(block*block) * 2.75) / float64(block*block*2+tbreak*tbreak*60) //1 - (float64(tbreak*tbreak)+float64(block*block)*0.4)/float64(tbreak*tbreak+block*block)
			}

			blockRate += float64(defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeBlockRatePercent)) / float64(common.MAX_RATE)

			isBlock = mathutils.RandomOneHit(blockRate)

			if isBlock {
				blockRatio := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBlock)) / common.MAX_RATE
				damage *= blockRatio
			}

			//specialEffectType = skillTemplate.GetSpecialEffectType()
			//有特殊效果
			if specialEffectType != skilltypes.SkillSpecialEffectTypeNone {

				hasSpecailEffect := false
				switch defendObj := defenceObject.(type) {
				case scene.NPC:
					{
						//if skillTemplate.SpecialTarget&skilltypes.SkillSpecialTargetMonster.Mask() != 0 {
						if specialTarget&skilltypes.SkillSpecialTargetMonster.Mask() != 0 {
							if !defendObj.GetBiologyTemplate().IsImmune(specialEffectType) {
								hasSpecailEffect = true
							}
						}
						break
					}
				case scene.Player:
					{
						//if skillTemplate.SpecialTarget&skilltypes.SkillSpecialTargetPlayer.Mask() != 0 {
						if specialTarget&skilltypes.SkillSpecialTargetPlayer.Mask() != 0 {
							hasSpecailEffect = true
						}
						break
					}
				}
				if hasSpecailEffect {
					specialHit := mathutils.RandomHit(common.MAX_RATE, int(skillTemplate.SpecialEffectRate))
					if specialHit {
						distance := coreutils.Distance(attackObject.GetPosition(), defenceObject.GetPosition())
						destPos := defenceObject.GetPosition()
						faceAngle := coreutils.GetAngle(defenceObject.GetPosition(), attackObject.GetPosition())
						hitSpeed := float64(0)
						//hitTime := skillTemplate.GetSpecialAnimationTime()
						hitTime := specialAnimationTime
						//stopTime := skillTemplate.GetSpecialTime()
						stopTime := specialTime
						//计算特殊效果
						switch /*skillTemplate.GetSpecialEffectType()*/ specialEffectType {
						//拉近
						case skilltypes.SkillSpecialEffectTypeClose:
							if distance > skillTemplate.GetAttackDistance() {
								t := skillTemplate.GetAttackDistance() / distance
								destPos = coreutils.Lerp(attackObject.GetPosition(), defenceObject.GetPosition(), t)
								destPos.Y = sb.MapTemplate().GetMap().GetHeight(destPos.X, destPos.Z)
								hitSpeed = (distance - skillTemplate.GetAttackDistance()) / hitTime
								//保留原地
								if !sb.MapTemplate().GetMap().IsMask(destPos.X, destPos.Z) {
									hitSpeed = float64(0)
									destPos = defenceObject.GetPosition()
								}
								h := scene.GetCheckAttackedMoveHandler(sb.MapTemplate().GetMapType())
								if h != nil {
									flag := h.CheckAttackedMove(attackObject, defenceObject, destPos)
									if !flag {
										hitSpeed = float64(0)
										destPos = defenceObject.GetPosition()
									}
								}
							}
							//计算目的地
							AttackedMove(defenceObject, destPos, faceAngle, hitSpeed, stopTime)
						//击退
						case skilltypes.SkillSpecialEffectTypeRepel:
							{
								//t := skillTemplate.GetSpecialDistance() / distance
								t := specialDistance / distance
								destPos := coreutils.Lerp(attackObject.GetPosition(), defenceObject.GetPosition(), 1+t)
								destPos.Y = sb.MapTemplate().GetMap().GetHeight(destPos.X, destPos.Z)
								faceAngle := coreutils.GetAngle(defenceObject.GetPosition(), attackObject.GetPosition())
								//hitTime := skillTemplate.GetSpecialTime()
								hitTime := specialTime
								//hitSpeed := skillTemplate.GetSpecialDistance() / hitTime
								hitSpeed := specialDistance / hitTime
								//保留原地
								if !sb.MapTemplate().GetMap().IsMask(destPos.X, destPos.Z) {
									hitSpeed = float64(0)
									destPos = defenceObject.GetPosition()
								}
								h := scene.GetCheckAttackedMoveHandler(sb.MapTemplate().GetMapType())
								if h != nil {
									flag := h.CheckAttackedMove(attackObject, defenceObject, destPos)
									if !flag {
										hitSpeed = float64(0)
										destPos = defenceObject.GetPosition()
									}
								}

								AttackedMove(defenceObject, destPos, faceAngle, hitSpeed, stopTime)
								//计算目的地
							}
						//冲刺
						case skilltypes.SkillSpecialEffectTypeSprint:
						}
					}
				}

			}

			//pvp系数

			switch defenceObject.(type) {
			case scene.Player:
				{
					damage *= (float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypePVP)) / common.MAX_RATE)
				}
			}

			//判断反弹
			fanTan = defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeFanTan)
			fanTanPercent = defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeFanTanPercent)
			totalFanTan = fanTan + int64(math.Ceil(damage*float64(fanTanPercent)/float64(common.MAX_RATE)))
			//计算伤害
			damageInt = int64(math.Ceil(damage))

			damageType = scenetypes.DamageTypeLingTong

			if !owner.IsDead() {
				if totalFanTan > 0 {
					//改变漂字
					attackDamageType := scenetypes.DamageTypeFISH
					//反弹给攻击者
					// owner.CostHP(totalFanTan, defenceObject.GetId())
					// scFanTanDefenceObjectDamage := pbutil.BuildSCObjectDamage(owner, attackDamageType, -totalFanTan, 0, defenceObject.GetId())
					// BroadcastNeighborIncludeSelf(owner, scFanTanDefenceObjectDamage)
					CostHP(owner, totalFanTan, 0, defenceObject.GetId(), attackDamageType)
				}
			}

		AfterDamage:
			beHarmSkillTemplate := getMaxDamageAttackedSkill(defenceObject, true)
			if beHarmSkillTemplate != nil {
				maxDamage := int64(math.Ceil(float64(beHarmSkillTemplate.BeHarmLimitHp) / float64(common.MAX_RATE) * float64(defenceObject.GetMaxHP())))
				//触发最大伤害技能
				if damageInt > maxDamage {
					damageInt = maxDamage
					Attack(defenceObject, defenceObject.GetPosition(), defenceObject.GetAngle(), beHarmSkillTemplate, false)
				}
			}
			totalDamage = damageInt
			hate := defenceObject.GetHate(owner)
			hate = int(math.Ceil(float64(skillTemplate.HatredPersent) / float64(common.MAX_RATE) * float64(hate)))
			hate += int(skillTemplate.HatredValue)
			hate += int(damageInt)
			if hate > 0 {
				defenceObject.AddHate(owner, hate)
			}

			//发送事件
			if active {
				gameevent.Emit(sceneeventtypes.EventTypeBattleObjectAttacked, defenceObject, owner)
			}

			//攻击者进入战斗
			owner.Battle()

			//进入战斗
			switch defendObj := defenceObject.(type) {
			case scene.Player:
				defendObj.Battle()
				owner.PvpBattle()
				defendObj.PvpBattle()
				scPlayerAttacked := pbutil.BuildSCPlayerAttacked(owner.GetId())
				defendObj.SendMsg(scPlayerAttacked)
				break
			}

			//判断死亡
			// isDead = defenceObject.CostHP(damageInt, owner.GetId())

			switch skillTemplate.GetSkillFirstType() {
			case skilltypes.SkillFirstTypeGuHun:
				{
					damageType = scenetypes.DamageTypeSoul
					break
				}
			case skilltypes.SkillFirstTypeShiHunFanAdvancedSkill:
				{
					damageType = scenetypes.DamageTypeXiXue
					break
				}
			}

			// if active {
			// 	scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, damageType, -damageInt, skillId, attackObject.GetId())
			// } else {
			// 	scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, damageType, -damageInt, 0, attackObject.GetId())
			// }
			// BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)
			if active {
				isDead = CostHP(defenceObject, damageInt, skillId, owner.GetId(), damageType)
			} else {
				isDead = CostHP(defenceObject, damageInt, 0, owner.GetId(), damageType)
			}

			if isBaseDamage {
				goto AfterHunYuan
			}
			if !isDead {
				//计算混元伤害
				hunYuanAttack := attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeHuanYunAttack)
				hunYuanDef := defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeHuanYunDef)
				hunYuanDamage := hunYuanAttack - hunYuanDef
				if hunYuanDamage > 0 {
					totalDamage += hunYuanDamage
					hunYuanDamageType := scenetypes.DamageTypeHunYuan
					if active {
						isDead = CostHP(defenceObject, hunYuanDamage, skillId, owner.GetId(), hunYuanDamageType)
					} else {
						isDead = CostHP(defenceObject, hunYuanDamage, 0, owner.GetId(), hunYuanDamageType)
					}
					// //判断死亡
					// isDead = defenceObject.CostHP(hunYuanDamage, attackObject.GetId())
					// //发送伤害数据
					// var scHunYuanObjectDamage *scenepb.SCObjectDamage
					// if active {
					// 	scHunYuanObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeHunYuan, -hunYuanDamage, skillId, attackObject.GetId())
					// } else {
					// 	scHunYuanObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeHunYuan, -hunYuanDamage, 0, attackObject.GetId())
					// }
					// BroadcastNeighborIncludeSelf(defenceObject, scHunYuanObjectDamage)
				}
			}
		AfterHunYuan:

			//受击光效
			attackedBuffTemplate := skillTemplate.GetAttackedBuffTemplate()
			if attackedBuffTemplate != nil {
				AddBuff(defenceObject, int32(attackedBuffTemplate.TemplateId()), owner.GetId(), common.MAX_RATE)
			}

			//死亡
			if isDead {
				log.WithFields(log.Fields{
					"attackObject":  attackObject,
					"defenceObject": defenceObject,
					"skillId":       skillId,
					"pos":           pos,
					"angle":         angle,
					"active":        active,
				}).Debugln("sceAddBuffWithTianFuListne: 死亡")
				continue
			}

			//添加技能buff
			for buffId, buffRate := range skillTemplate.GetBuffMap() {
				//添加buff
				buffHit := mathutils.RandomHit(common.MAX_RATE, int(buffRate))
				if buffHit {
					//TODO 一般buff是不会永久保存的
					AddBuff(defenceObject, buffId, owner.GetId(), int64(buffRate))
				}
			}

			//天赋buff
			addTianFuBuff(defenceObject, attackObject, skillTemplate)

			//主动施法且不限制被动
			if skillTemplate.IsPositive() && skillTemplate.GetLimitTouchType() == skilltypes.SkillLimitTouchTypeNo {
				//获取对象的被动触发技能
				attackedProbabilitySkillMap := defenceObject.GetSkills(skilltypes.SkillSecondTypeAttackedProbability)
				if len(attackedProbabilitySkillMap) != 0 {
					for _, attackedProbabilitySkill := range attackedProbabilitySkillMap {
						attackedProbabilitySkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(attackedProbabilitySkill.GetSkillId(), attackedProbabilitySkill.GetLevel())
						if defenceObject.IsSkillInCd(int32(attackedProbabilitySkillTempalte.TypeId)) {
							continue
						}

						//不能触发
						if attackedProbabilitySkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypePlayer.Mask() == 0 {
							continue
						}

						attackedFlag := mathutils.RandomHit(common.MAX_RATE, int(attackedProbabilitySkillTempalte.Rate))
						if attackedFlag {
							Attack(defenceObject, defenceObject.GetPosition(), defenceObject.GetAngle(), attackedProbabilitySkillTempalte, false)
							break
						}
					}
				}

				//获取对象的被动触发技能
				percentSkillMap := defenceObject.GetSkills(skilltypes.SkillSecondTypeHp)
				if len(percentSkillMap) != 0 {
					for _, percentSkill := range percentSkillMap {
						percentSkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(percentSkill.GetSkillId(), percentSkill.GetLevel())
						if defenceObject.IsSkillInCd(int32(percentSkillTempalte.TypeId)) {
							continue
						}
						hpPercent := float64(defenceObject.GetHP()) / float64(defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP))

						if hpPercent*common.MAX_RATE > float64(percentSkillTempalte.HpTrigger) {
							continue
						}

						//不能触发
						if percentSkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypePlayer.Mask() == 0 {
							continue
						}

						Attack(defenceObject, defenceObject.GetPosition(), defenceObject.GetAngle(), percentSkillTempalte, false)
					}
				}
			}
		}
		//主动施法
		if skillTemplate.IsPositive() {
			//获取自己主动触发技能
			attackProbabilitySkillMap := attackObject.GetSkills(skilltypes.SkillSecondTypeAttackProbability)
			for _, attackProbabilitySkill := range attackProbabilitySkillMap {
				attackProbabilitySkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(attackProbabilitySkill.GetSkillId(), attackProbabilitySkill.GetLevel())
				if attackObject.IsSkillInCd(int32(attackProbabilitySkillTempalte.TypeId)) {
					continue
				}
				attackFlag := mathutils.RandomHit(common.MAX_RATE, int(attackProbabilitySkillTempalte.Rate))
				if attackFlag {
					Attack(attackObject, attackObject.GetPosition(), attackObject.GetAngle(), attackProbabilitySkillTempalte, false)
					break
				}
			}
		}

	case skilltypes.SkillFourthTypeHelp:
		//遍历作用目标
		for _, defenceObject := range sos {

			//获取治疗值
			maxHP := defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
			cure := float64(skillTemplate.CurePersent) / float64(common.MAX_RATE) * float64(maxHP)
			cure += float64(skillTemplate.CurePersent)
			cureInt := int64(math.Ceil(cure))
			if cureInt != 0 {
				//发送事件
				if active {
					gameevent.Emit(sceneeventtypes.EventTypeBattleObjectCure, defenceObject, nil)
				}
				defenceObject.AddHP(cureInt)
				//发送伤害事件

				//发送伤害数据
				scObjectDamage := pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeRecovery, cureInt, skillId, attackObject.GetId())
				BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)

			}
			for buffId, buffRate := range skillTemplate.GetBuffMap() {
				//添加buff
				buffHit := mathutils.RandomHit(common.MAX_RATE, int(buffRate))
				if buffHit {
					//TODO 一般buff是不会永久保存的
					AddBuff(defenceObject, buffId, owner.GetId(), int64(buffRate))
				}
			}

			//天赋buff
			addTianFuBuff(defenceObject, attackObject, skillTemplate)
		}
	}

	return true
}

func CalculateAttack(attackObject scene.BattleObject, pos coretypes.Position, angle float64, skillTemplate *gametemplate.SkillTemplate, active bool) (flag bool) {
	sb := attackObject.GetScene()
	skillId := int32(skillTemplate.TypeId)
	level := int32(1)
	//判断是否有这个技能
	if skillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeSubSkill && skillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeCastingSoul {
		//判断是否有这个技能
		ski := attackObject.GetSkill(skillId)
		if ski == nil {
			log.WithFields(log.Fields{
				"attackObject": attackObject,
				"skillId":      skillId,
			}).Warnln("scene: 技能不存在")
			return false
		}
		level = ski.GetLevel()
	}
	skillTemplate = skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(skillId, level)

	specialEffectType, specialTarget, specialDistance, specialAnimationTime, specialTime := GetSpecialEffect(attackObject, skillTemplate)
	//技能区域检测
	sos := getAttackBattleObjectList(attackObject, skillTemplate)
	skillFourthType := skillTemplate.GetSkillFourthType()
	s := attackObject.GetScene()
	switch skillFourthType {
	case skilltypes.SkillFourthTypeAttack:

		damagePersent := int32(0)
		damageAttack := float64(0)
		damageValueBase := int32(0)
		if skillTemplate.IsDynamic() {
			skillLevelTemplate := skillTemplate.GetSkillByLevel(level)
			damageValueBase = skillLevelTemplate.DamageValueBase
			damageAttack = skillLevelTemplate.GetDamageAttack()
		} else {
			damagePersent = skillTemplate.DamagePersent
			damageAttack = skillTemplate.GetDamageAttack()
			damageValueBase = skillTemplate.DamageValueBase
		}

		//判断帝魂加成
		if skillTemplate.GetSkillFirstType() == skilltypes.SkillFirstTypeGuHun {
			switch attackObj := attackObject.(type) {
			case scene.Player:
				soulChainTemplate := soul.GetSoulService().GetSoulChainTemplate(attackObj.GetSoulAwakenNum())
				if soulChainTemplate != nil {
					guHunPercent := soulChainTemplate.Value
					damageAttack += float64(guHunPercent) / float64(common.MAX_RATE)
				}
				break
			}
		}

		isAttackPlayer := attackObject.GetSceneObjectType() == scenetypes.BiologyTypePlayer
		//遍历作用目标
		for _, defenceObject := range sos {
			//复活保护
			if defenceObject.GetBattleLimit()&scenetypes.BattleLimitTypeNoAttacked.Mask() != 0 {
				attackObj, ok := attackObject.(scene.Player)
				if ok {
					playerlogic.SendSystemMessage(attackObj, lang.BattleProtectNoAttacked)
				}
				continue
			}
			switch defenceObj := defenceObject.(type) {
			case scene.NPC:
				{
					h := scene.GetCheckNPCAttackHandler(defenceObj.GetBiologyTemplate().GetBiologyScriptType())
					if h != nil {
						flag := h.CheckAttack(attackObject, defenceObj)
						if !flag {
							continue
						}
					}
				}
			case scene.Player:
				//过场保护
				if defenceObj.GetBattleLimit()&scenetypes.BattleLimitTypeAttackedChangeScene.Mask() != 0 {
					attackObj, ok := attackObject.(scene.Player)
					if ok {
						playerlogic.SendSystemMessage(attackObj, lang.BattleProtectChangeScene)
					}
					continue
				}

				// 检查攻击
				h := scene.GetCheckAttackHandler(s.MapTemplate().GetMapType())
				if h != nil {
					if !h.CheckAttack(attackObject) {
						continue
					}
				}

				if isAttackPlayer {
					//判断安全区
					if s.MapTemplate().IsSafe(attackObject.GetPosition()) {
						log.WithFields(log.Fields{
							"attackObject": attackObject,
							"skillId":      skillId,
						}).Warnln("scene: 攻方安全区")
						continue
					}
					//判断安全区
					if s.MapTemplate().IsSafe(defenceObj.GetPosition()) {
						log.WithFields(log.Fields{
							"defenceObj": defenceObj,
							"skillId":    skillId,
						}).Warnln("scene: 防御方安全区")
						continue
					}
					attackObj := attackObject.(scene.Player)

					//复活保护
					if defenceObj.GetBattleLimit()&scenetypes.BattleLimitTypeAttackedRelive.Mask() != 0 {
						playerlogic.SendSystemMessage(attackObj, lang.BattleProtectRelive)
						continue
					}
					//复活保护
					if defenceObj.GetBattleLimit()&scenetypes.BattleLimitTypeAttackedPKProtect.Mask() != 0 {
						playerlogic.SendSystemMessage(attackObj, lang.BattleProtectPK)
						continue
					}

					//pk保护
					protectedLevel := s.MapTemplate().ProtectLevel
					if defenceObj.GetLevel() < protectedLevel {
						playerlogic.SendSystemMessage(attackObj, lang.PkStateProtectCanNotAttack, fmt.Sprintf("%d", protectedLevel))
						continue
					}
				}
			}
			isPlayer := defenceObject.GetSceneObjectType() == scenetypes.BiologyTypePlayer

			if isPlayer {
				if skillTemplate.TargetAction&skilltypes.SkillTargetActionTypePlayer.Mask() == 0 {
					//无效
					var scObjectDamage *scenepb.SCObjectDamage
					if active {
						scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeWuXiao, 0, skillId, attackObject.GetId())
					} else {
						scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeWuXiao, 0, 0, attackObject.GetId())
					}
					BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)
					continue
				}
			} else {
				if skillTemplate.TargetAction&skilltypes.SkillTargetActionTypeMonster.Mask() == 0 {
					//无效
					var scObjectDamage *scenepb.SCObjectDamage
					if active {
						scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeWuXiao, 0, skillId, attackObject.GetId())
					} else {
						scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeWuXiao, 0, 0, attackObject.GetId())
					}
					BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)
					continue
				}
			}

			//算命中
			hit := getHit(attackObject) //attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeHit)

			dodge := getDodge(defenceObject) //defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeDodge)
			hit -= dodge
			if hit < 0 {
				hit = 0
			}
			hitFlag := mathutils.RandomHit(common.MAX_RATE, int(hit))
			//随机
			if !hitFlag {
				log.WithFields(log.Fields{
					"attackObject":  attackObject,
					"defenceObject": defenceObject,
					"skillId":       skillId,
					"pos":           pos,
					"angle":         angle,
					"active":        active,
					"hit":           hit,
					"dodge":         dodge,
				}).Debug("scene: 闪避了")
				//发送闪避
				var scObjectDamage *scenepb.SCObjectDamage
				if active {
					scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeDodge, 0, skillId, attackObject.GetId())
				} else {
					scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeDodge, 0, 0, attackObject.GetId())
				}
				BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)
				continue
			}
			//命中了
			//算出伤害
			damage := float64(0)
			var damageInt int64
			var totalDamage int64
			damageType := scenetypes.DamageTypeAttack
			var crit int64
			var tough int64
			var critRate float64
			var isCrit bool
			var block int64
			var isBlock bool
			var tbreak int64
			var blockRate float64
			var fanTan int64
			var fanTanPercent int64
			var totalFanTan int64
			//var specialEffectType skilltypes.SkillSpecialEffectType
			var isDead bool
			// var scObjectDamage *scenepb.SCObjectDamage
			var isBaseDamage bool
			defendNPC, ok := defenceObject.(scene.NPC)
			if ok {
				if defendNPC.GetBiologyTemplate().BaseDamge != 0 {
					damageInt = int64(defendNPC.GetBiologyTemplate().BaseDamge)
					isBaseDamage = true
					goto AfterDamage
				}
			}

			if damagePersent == 0 {
				damage = scene.GetBasicDamage(attackObject, defenceObject, damageAttack, float64(damageValueBase))
				//至少为1
				if damage <= 0 {
					damage = 1
				}
				//计算浮动效果

				//上下浮动0.1
				wave := rand.Float64()/10 - 0.05
				damage *= (1 + wave)
			} else {
				defenceMaxHP := attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
				damage = float64(damagePersent) / common.MAX_RATE * float64(defenceMaxHP)
			}
			//判断暴击
			crit = getCrit(attackObject) //attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeCrit)
			tough = getTough(defenceObject)
			critRate = 0
			if crit != 0 {
				//TODO 暴击率计算
				critRate = (float64(crit*crit) * 2.75) / float64(crit*crit*2+tough*tough*60)
			}
			critRate += float64(skillTemplate.AddCritical) / float64(common.MAX_RATE)
			//增加攻击暴击几率万分比
			critRate += float64(attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeCritRatePercent)) / float64(common.MAX_RATE)

			isCrit = mathutils.RandomOneHit(critRate)

			if isCrit {
				critRatio := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCrit)) / common.MAX_RATE
				//增加暴击伤害万分比
				critRatio += float64(attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeCritHarmPercent)) / float64(common.MAX_RATE)
				damage *= critRatio
			}

			//判断格挡
			block = defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeBlock)
			tbreak = attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeAbnormality)
			blockRate = 0
			if block != 0 {
				blockRate = (float64(block*block) * 2.75) / float64(block*block*2+tbreak*tbreak*60) //1 - (float64(tbreak*tbreak)+float64(block*block)*0.4)/float64(tbreak*tbreak+block*block)
			}

			blockRate += float64(defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeBlockRatePercent)) / float64(common.MAX_RATE)

			isBlock = mathutils.RandomOneHit(blockRate)

			if isBlock {
				blockRatio := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBlock)) / common.MAX_RATE
				damage *= blockRatio
			}

			//specialEffectType = skillTemplate.GetSpecialEffectType()
			//有特殊效果
			if specialEffectType != skilltypes.SkillSpecialEffectTypeNone {

				hasSpecailEffect := false
				switch defendObj := defenceObject.(type) {
				case scene.NPC:
					{
						//if skillTemplate.SpecialTarget&skilltypes.SkillSpecialTargetMonster.Mask() != 0 {
						if specialTarget&skilltypes.SkillSpecialTargetMonster.Mask() != 0 {
							if !defendObj.GetBiologyTemplate().IsImmune(specialEffectType) {
								hasSpecailEffect = true
							}
						}
						break
					}
				case scene.Player:
					{
						//if skillTemplate.SpecialTarget&skilltypes.SkillSpecialTargetPlayer.Mask() != 0 {
						if specialTarget&skilltypes.SkillSpecialTargetPlayer.Mask() != 0 {
							hasSpecailEffect = true
						}
						break
					}
				}
				if hasSpecailEffect {
					specialHit := mathutils.RandomHit(common.MAX_RATE, int(skillTemplate.SpecialEffectRate))
					if specialHit {
						distance := coreutils.Distance(attackObject.GetPosition(), defenceObject.GetPosition())
						destPos := defenceObject.GetPosition()
						faceAngle := coreutils.GetAngle(defenceObject.GetPosition(), attackObject.GetPosition())
						hitSpeed := float64(0)
						//hitTime := skillTemplate.GetSpecialAnimationTime()
						hitTime := specialAnimationTime
						//stopTime := skillTemplate.GetSpecialTime()
						stopTime := specialTime
						//计算特殊效果
						switch /*skillTemplate.GetSpecialEffectType()*/ specialEffectType {
						//拉近
						case skilltypes.SkillSpecialEffectTypeClose:
							if distance > skillTemplate.GetAttackDistance() {
								t := skillTemplate.GetAttackDistance() / distance
								destPos = coreutils.Lerp(attackObject.GetPosition(), defenceObject.GetPosition(), t)
								destPos.Y = sb.MapTemplate().GetMap().GetHeight(destPos.X, destPos.Z)
								hitSpeed = (distance - skillTemplate.GetAttackDistance()) / hitTime
								//保留原地
								if !sb.MapTemplate().GetMap().IsMask(destPos.X, destPos.Z) {
									hitSpeed = float64(0)
									destPos = defenceObject.GetPosition()
								}
								h := scene.GetCheckAttackedMoveHandler(sb.MapTemplate().GetMapType())
								if h != nil {
									flag := h.CheckAttackedMove(attackObject, defenceObject, destPos)
									if !flag {
										hitSpeed = float64(0)
										destPos = defenceObject.GetPosition()
									}
								}
							}
							//计算目的地
							AttackedMove(defenceObject, destPos, faceAngle, hitSpeed, stopTime)
						//击退
						case skilltypes.SkillSpecialEffectTypeRepel:
							{
								//t := skillTemplate.GetSpecialDistance() / distance
								t := specialDistance / distance
								destPos := coreutils.Lerp(attackObject.GetPosition(), defenceObject.GetPosition(), 1+t)
								destPos.Y = sb.MapTemplate().GetMap().GetHeight(destPos.X, destPos.Z)
								faceAngle := coreutils.GetAngle(defenceObject.GetPosition(), attackObject.GetPosition())
								//hitTime := skillTemplate.GetSpecialTime()
								hitTime := specialTime
								//hitSpeed := skillTemplate.GetSpecialDistance() / hitTime
								hitSpeed := specialDistance / hitTime
								//保留原地
								if !sb.MapTemplate().GetMap().IsMask(destPos.X, destPos.Z) {
									hitSpeed = float64(0)
									destPos = defenceObject.GetPosition()
								}
								h := scene.GetCheckAttackedMoveHandler(sb.MapTemplate().GetMapType())
								if h != nil {
									flag := h.CheckAttackedMove(attackObject, defenceObject, destPos)
									if !flag {
										hitSpeed = float64(0)
										destPos = defenceObject.GetPosition()
									}
								}

								AttackedMove(defenceObject, destPos, faceAngle, hitSpeed, stopTime)
								//计算目的地
							}
						//冲刺
						case skilltypes.SkillSpecialEffectTypeSprint:
						}
					}
				}

			}

			//pvp系数
			switch attackObject.(type) {
			case scene.Player:
				switch defenceObject.(type) {
				case scene.Player:
					{
						damage *= (float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypePVP)) / common.MAX_RATE)
					}
				}
			}

			//判断反弹
			fanTan = defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeFanTan)
			fanTanPercent = defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeFanTanPercent)
			totalFanTan = fanTan + int64(math.Ceil(damage*float64(fanTanPercent)/float64(common.MAX_RATE)))
			//计算伤害
			damageInt = int64(math.Ceil(damage))

			damageType = scenetypes.DamageTypeAttack
			if isCrit && isBlock {
				damageType = scenetypes.DamageTypeCritGeDang
			} else if isCrit {
				damageType = scenetypes.DamageTypeCrit
			} else if isBlock {
				damageType = scenetypes.DamageTypeAttackGeDang
			}

			if totalFanTan > 0 {
				//改变漂字
				attackDamageType := scenetypes.DamageTypeFISH

				//反弹给攻击者
				// attackDead := attackObject.CostHP(totalFanTan, defenceObject.GetId())
				// scFanTanDefenceObjectDamage := pbutil.BuildSCObjectDamage(attackObject, attackDamageType, -totalFanTan, 0, defenceObject.GetId())
				// BroadcastNeighborIncludeSelf(attackObject, scFanTanDefenceObjectDamage)
				attackDead := CostHP(attackObject, totalFanTan, 0, defenceObject.GetId(), attackDamageType)
				//攻击者死亡了
				if attackDead {
					break
				}

			}

		AfterDamage:
			//计算最大伤害
			beHarmSkillTemplate := getMaxDamageAttackedSkill(defenceObject, isAttackPlayer)
			if beHarmSkillTemplate != nil {
				//TODO 判断cd
				maxDamage := int64(math.Ceil(float64(beHarmSkillTemplate.BeHarmLimitHp) / float64(common.MAX_RATE) * float64(defenceObject.GetMaxHP())))
				//触发最大伤害技能
				if damageInt > maxDamage {
					damageInt = maxDamage
					Attack(defenceObject, defenceObject.GetPosition(), defenceObject.GetAngle(), beHarmSkillTemplate, false)
				}
			}
			totalDamage = damageInt
			hate := defenceObject.GetHate(attackObject)
			hate = int(math.Ceil(float64(skillTemplate.HatredPersent) / float64(common.MAX_RATE) * float64(hate)))
			hate += int(skillTemplate.HatredValue)
			hate += int(damageInt)
			if hate > 0 {
				defenceObject.AddHate(attackObject, hate)
			}
			// //假如是npc的话
			// switch defendObj := defenceObject.(type) {
			// //npc 加仇恨
			// case scene.NPC:
			// 	hate := defendObj.GetHate(attackObject)
			// 	hate = int(math.Ceil(float64(skillTemplate.HatredPersent) / float64(common.MAX_RATE) * float64(hate)))
			// 	hate += int(skillTemplate.HatredValue)
			// 	hate += int(damageInt)
			// 	if hate > 0 {
			// 		defendObj.AddHate(attackObject, hate)
			// 	}
			// 	break
			// case scene.Player:
			// 	hate := defendObj.GetHate(attackObject)
			// 	hate = int(math.Ceil(float64(skillTemplate.HatredPersent) / float64(common.MAX_RATE) * float64(hate)))
			// 	hate += int(skillTemplate.HatredValue)
			// 	hate += int(damageInt)
			// 	if hate > 0 {
			// 		defendObj.AddHate(attackObject, hate)
			// 	}
			// 	break
			// }
			//发送事件
			if active {
				gameevent.Emit(sceneeventtypes.EventTypeBattleObjectAttacked, defenceObject, attackObject)
			}

			//攻击者进入战斗
			switch attackObj := attackObject.(type) {
			case scene.Player:
				attackObj.Battle()
				break
			}

			//进入战斗
			switch defendObj := defenceObject.(type) {
			case scene.Player:
				defendObj.Battle()
				switch attackObj := attackObject.(type) {
				case scene.Player:
					attackObj.PvpBattle()
					defendObj.PvpBattle()
					scPlayerAttacked := pbutil.BuildSCPlayerAttacked(attackObj.GetId())
					defendObj.SendMsg(scPlayerAttacked)
				}
				break
			}

			//判断死亡
			// isDead = defenceObject.CostHP(damageInt, attackObject.GetId())
			switch skillTemplate.GetSkillFirstType() {
			case skilltypes.SkillFirstTypeGuHun:
				{
					damageType = scenetypes.DamageTypeSoul
					break
				}
			case skilltypes.SkillFirstTypeShiHunFanAdvancedSkill:
				{
					damageType = scenetypes.DamageTypeXiXue
					break
				}
			}

			// if active {
			// 	scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, damageType, -damageInt, skillId, attackObject.GetId())
			// } else {
			// 	scObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, damageType, -damageInt, 0, attackObject.GetId())
			// }
			// BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)

			if active {
				isDead = CostHP(defenceObject, damageInt, skillId, attackObject.GetId(), damageType)
			} else {
				isDead = CostHP(defenceObject, damageInt, 0, attackObject.GetId(), damageType)
			}

			if isBaseDamage {
				goto AfterHunYuan
			}
			if !isDead {
				//计算混元伤害
				hunYuanAttack := attackObject.GetBattleProperty(propertytypes.BattlePropertyTypeHuanYunAttack)
				hunYuanDef := defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeHuanYunDef)
				hunYuanDamage := hunYuanAttack - hunYuanDef
				if hunYuanDamage > 0 {
					totalDamage += hunYuanDamage
					// //判断死亡
					// isDead = defenceObject.CostHP(hunYuanDamage, attackObject.GetId())
					// //发送伤害数据
					// var scHunYuanObjectDamage *scenepb.SCObjectDamage
					// if active {
					// 	scHunYuanObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeHunYuan, -hunYuanDamage, skillId, attackObject.GetId())
					// } else {
					// 	scHunYuanObjectDamage = pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeHunYuan, -hunYuanDamage, 0, attackObject.GetId())
					// }
					// BroadcastNeighborIncludeSelf(defenceObject, scHunYuanObjectDamage)
					hunYuanDamageType := scenetypes.DamageTypeHunYuan
					if active {
						isDead = CostHP(defenceObject, hunYuanDamage, skillId, attackObject.GetId(), hunYuanDamageType)
					} else {
						isDead = CostHP(defenceObject, hunYuanDamage, 0, attackObject.GetId(), hunYuanDamageType)
					}
				}
			}
		AfterHunYuan:

			//受击光效
			attackedBuffTemplate := skillTemplate.GetAttackedBuffTemplate()
			if attackedBuffTemplate != nil {
				AddBuff(defenceObject, int32(attackedBuffTemplate.TemplateId()), attackObject.GetId(), common.MAX_RATE)
				// scObjectBuff := pbutil.BuildSCObjectAttackedBuff(defenceObject, attackedBuffTemplate.TemplateId())
				// broadcastNeighborIncludeSelf(defenceObject, scObjectBuff)
			}

			// 添加buff的buff
			for _, buffObj := range attackObject.GetBuffs() {
				attackBuffId := buffObj.GetBuffId()
				attackBuffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(attackBuffId)
				if attackBuffTemplate.GetTouchType() != scenetypes.BuffTouchTypeObjectDamage {
					continue
				}
				//判断是否触发状态
				for subBuffId, subBuffRage := range attackBuffTemplate.GetSubBuffMap() {
					subRateHit := mathutils.RandomHit(common.MAX_RATE, int(subBuffRage))
					if subRateHit {
						attackId := int64(0)
						if attackObject.GetSceneObjectType() == scenetypes.BiologyTypePlayer {
							attackId = attackObject.GetId()
						}
						//添加buff
						AddBuff(defenceObject, subBuffId, attackId, int64(subBuffRage))
					}
				}

				// if attackBuffTemplate.GetSubBuffTemplate() != nil {
				// 	subRateHit := mathutils.RandomHit(common.MAX_RATE, int(attackBuffTemplate.SubRate))
				// 	if subRateHit {
				// 		attackId := int64(0)
				// 		if attackObject.GetSceneObjectType() == scenetypes.BiologyTypePlayer {
				// 			attackId = attackObject.GetId()
				// 		}
				// 		//添加buff
				// 		AddBuff(defenceObject, int32(attackBuffTemplate.GetSubBuffTemplate().TemplateId()), attackId, int64(attackBuffTemplate.SubRate))
				// 	}
				// }
			}
			//死亡
			if isDead {
				log.WithFields(log.Fields{
					"attackObject":  attackObject,
					"defenceObject": defenceObject,
					"skillId":       skillId,
					"pos":           pos,
					"angle":         angle,
					"active":        active,
				}).Debugln("scene: 死亡")
				continue
			}

			if active {
				//判断特殊技能
				for _, teShuSkill := range attackObject.GetTeShuSkills() {
					if teShuSkill.GetChuFaRate() <= 0 {
						continue
					}
					//判断技能cd
					if attackObject.IsSkillInCd(teShuSkill.GetSkillId()) {
						continue
					}
					teShuSkillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(teShuSkill.GetSkillId())
					if teShuSkillTemplate == nil {
						continue
					}
					// chuFaRate := teShuSkill.GetChuFaRate()
					// diKangSkill := defenceObject.GetTeShuSkill(teShuSkill.GetSkillId())
					// diKangRate := int32(0)
					// if diKangSkill != nil {
					// 	diKangRate = diKangSkill.GetDiKangRate()
					// }
					// chuFaRate -= diKangRate
					chuFaRate := common.MAX_RATE
					chuFa := mathutils.RandomHit(common.MAX_RATE, int(chuFaRate))
					if !chuFa {
						continue
					}
					Attack(attackObject, attackObject.GetPosition(), attackObject.GetAngle(), teShuSkillTemplate, false)
				}
			}

			//添加技能buff
			for buffId, buffRate := range skillTemplate.GetBuffMap() {
				//添加buff
				buffHit := mathutils.RandomHit(common.MAX_RATE, int(buffRate))
				if buffHit {
					//TODO 一般buff是不会永久保存的
					AddBuff(defenceObject, buffId, attackObject.GetId(), int64(buffRate))
				}
			}

			//天赋buff
			addTianFuBuff(defenceObject, attackObject, skillTemplate)

			//主动施法且不限制被动
			if skillTemplate.IsPositive() && skillTemplate.GetLimitTouchType() == skilltypes.SkillLimitTouchTypeNo {
				//获取对象的被动触发技能
				attackedProbabilitySkillMap := defenceObject.GetSkills(skilltypes.SkillSecondTypeAttackedProbability)
				if len(attackedProbabilitySkillMap) != 0 {
					for _, attackedProbabilitySkill := range attackedProbabilitySkillMap {
						attackedProbabilitySkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(attackedProbabilitySkill.GetSkillId(), attackedProbabilitySkill.GetLevel())
						if defenceObject.IsSkillInCd(int32(attackedProbabilitySkillTempalte.TypeId)) {
							continue
						}
						if isAttackPlayer {
							//不能触发
							if attackedProbabilitySkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypePlayer.Mask() == 0 {
								continue
							}
						} else {
							//不能触发
							if attackedProbabilitySkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypeMonster.Mask() == 0 {
								continue
							}
						}

						attackedFlag := mathutils.RandomHit(common.MAX_RATE, int(attackedProbabilitySkillTempalte.Rate))
						if attackedFlag {
							Attack(defenceObject, defenceObject.GetPosition(), defenceObject.GetAngle(), attackedProbabilitySkillTempalte, false)
							break
						}
					}
				}

				//获取对象的被动触发技能
				percentSkillMap := defenceObject.GetSkills(skilltypes.SkillSecondTypeHp)
				if len(percentSkillMap) != 0 {
					for _, percentSkill := range percentSkillMap {
						percentSkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(percentSkill.GetSkillId(), percentSkill.GetLevel())
						if defenceObject.IsSkillInCd(int32(percentSkillTempalte.TypeId)) {
							continue
						}
						hpPercent := float64(defenceObject.GetHP()) / float64(defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP))

						if hpPercent*common.MAX_RATE > float64(percentSkillTempalte.HpTrigger) {
							continue
						}
						if isAttackPlayer {
							//不能触发
							if percentSkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypePlayer.Mask() == 0 {
								continue
							}
						} else {
							//不能触发
							if percentSkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypeMonster.Mask() == 0 {
								continue
							}
						}

						Attack(defenceObject, defenceObject.GetPosition(), defenceObject.GetAngle(), percentSkillTempalte, false)
					}
				}

				switch defenceObj := defenceObject.(type) {
				case scene.Player:
					{
						defenceLingTong := defenceObj.GetLingTong()
						if defenceLingTong != nil && CheckIfLingTongAndPlayerSameScene(defenceLingTong) {
							attackedProbabilitySkillMap := defenceLingTong.GetSkills(skilltypes.SkillSecondTypeAttackedProbability)
							if len(attackedProbabilitySkillMap) != 0 {
								for _, attackedProbabilitySkill := range attackedProbabilitySkillMap {
									attackedProbabilitySkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(attackedProbabilitySkill.GetSkillId(), attackedProbabilitySkill.GetLevel())
									if defenceLingTong.IsSkillInCd(int32(attackedProbabilitySkillTempalte.TypeId)) {
										continue
									}
									if isAttackPlayer {
										//不能触发
										if attackedProbabilitySkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypePlayer.Mask() == 0 {
											continue
										}
									} else {
										//不能触发
										if attackedProbabilitySkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypeMonster.Mask() == 0 {
											continue
										}
									}

									attackedFlag := mathutils.RandomHit(common.MAX_RATE, int(attackedProbabilitySkillTempalte.Rate))
									if attackedFlag {
										Attack(defenceLingTong, defenceLingTong.GetPosition(), defenceLingTong.GetAngle(), attackedProbabilitySkillTempalte, false)
										break
									}
								}
							}

							percentSkillMap := defenceLingTong.GetSkills(skilltypes.SkillSecondTypeHp)
							if len(percentSkillMap) != 0 {
								for _, percentSkill := range percentSkillMap {
									percentSkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(percentSkill.GetSkillId(), percentSkill.GetLevel())
									if defenceLingTong.IsSkillInCd(int32(percentSkillTempalte.TypeId)) {
										continue
									}
									hpPercent := float64(defenceObject.GetHP()) / float64(defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP))

									if hpPercent*common.MAX_RATE > float64(percentSkillTempalte.HpTrigger) {
										continue
									}
									if isAttackPlayer {
										//不能触发
										if percentSkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypePlayer.Mask() == 0 {
											continue
										}
									} else {
										//不能触发
										if percentSkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypeMonster.Mask() == 0 {
											continue
										}
									}

									Attack(defenceLingTong, defenceLingTong.GetPosition(), defenceLingTong.GetAngle(), percentSkillTempalte, false)
								}
							}
						}
					}
				}
			}
		}
		//主动施法
		if skillTemplate.IsPositive() {
			//获取自己主动触发技能
			attackProbabilitySkillMap := attackObject.GetSkills(skilltypes.SkillSecondTypeAttackProbability)
			for _, attackProbabilitySkill := range attackProbabilitySkillMap {
				attackProbabilitySkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(attackProbabilitySkill.GetSkillId(), attackProbabilitySkill.GetLevel())
				if attackObject.IsSkillInCd(int32(attackProbabilitySkillTempalte.TypeId)) {
					continue
				}
				attackFlag := mathutils.RandomHit(common.MAX_RATE, int(attackProbabilitySkillTempalte.Rate))
				if attackFlag {
					Attack(attackObject, attackObject.GetPosition(), attackObject.GetAngle(), attackProbabilitySkillTempalte, false)
					break
				}
			}
		}

	case skilltypes.SkillFourthTypeHelp:
		//遍历作用目标
		for _, defenceObject := range sos {

			//获取治疗值
			maxHP := defenceObject.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
			cure := float64(skillTemplate.CurePersent) / float64(common.MAX_RATE) * float64(maxHP)
			cure += float64(skillTemplate.CurePersent)
			cureInt := int64(math.Ceil(cure))
			if cureInt != 0 {
				//发送事件
				if active {
					gameevent.Emit(sceneeventtypes.EventTypeBattleObjectCure, defenceObject, nil)
				}
				defenceObject.AddHP(cureInt)
				//发送伤害事件

				//发送伤害数据
				scObjectDamage := pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeRecovery, cureInt, skillId, attackObject.GetId())
				BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)

			}
			for buffId, buffRate := range skillTemplate.GetBuffMap() {
				//添加buff
				buffHit := mathutils.RandomHit(common.MAX_RATE, int(buffRate))
				if buffHit {
					//TODO 一般buff是不会永久保存的
					AddBuff(defenceObject, buffId, attackObject.GetId(), int64(buffRate))
				}
			}
			//天赋buff
			addTianFuBuff(defenceObject, attackObject, skillTemplate)
		}
	}

	return true
}

const (
	defaultPlayerLength = 0.1
	defaultPlayerWidth  = 0.1
)

//获取技能作用对象
func getAttackBattleObjectList(so scene.BattleObject, skillTemplate *gametemplate.SkillTemplate) (sos []scene.BattleObject) {
	//area := skillTemplate.GetSkillArea()
	area := GetSkillArea(so, skillTemplate)
	if skillTemplate.TargetSelect&skilltypes.SkillTargetSelectTypeSelf.Mask() != 0 {
		switch tbo := so.(type) {
		//灵童需要添加作用目标自身主人
		case scene.LingTong:
			if !tbo.GetOwner().IsDead() {
				sos = append(sos, tbo.GetOwner())
			}
			break
		default:
			sos = append(sos, so)
			break
		}
	}
	if area != nil {
		//TODO: 优化复用
		ps := make([]scene.BattleObject, 0, 8)
		ms := make([]scene.BattleObject, 0, 8)
		// s := so.GetScene()
		for _, neighbor := range so.GetNeighbors() {
			target, ok := neighbor.(scene.BattleObject)
			if !ok {
				continue
			}
			_, ok = neighbor.(scene.LingTong)
			if ok {
				continue
			}
			//无敌排除
			if target.GetBattleLimit()&scenetypes.BattleLimitTypeAttacked.Mask() != 0 {
				continue
			}
			//死亡排除
			if target.IsDead() {
				continue
			}

			switch targetObj := target.(type) {
			case scene.NPC:
				//特殊处理城门守卫
				if targetObj.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeXianMengNPC {
					//敌人
					if so.IsEnemy(target) {
						if skillTemplate.TargetSelect&skilltypes.SkillTargetSelectTypeEnemy.Mask() == 0 {
							continue
						}
					} else {
						//好友
						if skillTemplate.TargetSelect&skilltypes.SkillTargetSelectTypeAlliance.Mask() == 0 {
							continue
						}
					}
				}
				//筛选个数
				pengZhuang := targetObj.GetBiologyTemplate().Pengzhuang
				pengZhuangKuan := targetObj.GetBiologyTemplate().PengzhuangKuan
				if area.PositionInArea(so.GetPosition(), float64(so.GetAngle()), targetObj.GetPosition(), targetObj.GetAngle(), pengZhuang, pengZhuangKuan) {
					ms = append(ms, target)
				}
				break
			case scene.Player:
				//敌人
				if so.IsEnemy(target) {
					if skillTemplate.TargetSelect&skilltypes.SkillTargetSelectTypeEnemy.Mask() == 0 {
						continue
					}
				} else {
					//好友
					if skillTemplate.TargetSelect&skilltypes.SkillTargetSelectTypeAlliance.Mask() == 0 {
						continue
					}
				}
				if area.PositionInArea(so.GetPosition(), float64(so.GetAngle()), targetObj.GetPosition(), targetObj.GetAngle(), defaultPlayerLength, defaultPlayerWidth) {
					ps = append(ps, target)
				}
				break
			}
		}
		if len(ps) >= int(skillTemplate.TargetCount) {
			sos = append(sos, ps[:int(skillTemplate.TargetCount)]...)
		} else {
			sos = append(sos, ps...)
			remain := int(skillTemplate.TargetCount) - len(ps)
			if remain > len(ms) {
				sos = append(sos, ms...)
			} else {
				sos = append(sos, ms[:remain]...)
			}
		}
	}
	return sos
}

//获取最大伤害上限
func getMaxDamageAttackedSkill(defenceObject scene.BattleObject, isAttackPlayer bool) (attackedSkillTemplate *gametemplate.SkillTemplate) {
	attackedProbabilitySkillMap := defenceObject.GetSkills(skilltypes.SkillSecondTypeAttackedProbability)
	if len(attackedProbabilitySkillMap) != 0 {
		for _, attackedProbabilitySkill := range attackedProbabilitySkillMap {
			attackedProbabilitySkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(attackedProbabilitySkill.GetSkillId(), attackedProbabilitySkill.GetLevel())
			if defenceObject.IsSkillInCd(int32(attackedProbabilitySkillTempalte.TypeId)) {
				continue
			}
			if isAttackPlayer {
				//不能触发
				if attackedProbabilitySkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypePlayer.Mask() == 0 {
					continue
				}
			} else {
				//不能触发
				if attackedProbabilitySkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypeMonster.Mask() == 0 {
					continue
				}
			}
			if attackedProbabilitySkillTempalte.BeHarmLimitHp <= 0 {
				continue
			}
			if attackedSkillTemplate == nil {
				attackedSkillTemplate = attackedProbabilitySkillTempalte
			} else {
				if attackedProbabilitySkillTempalte.BeHarmLimitHp < attackedSkillTemplate.BeHarmLimitHp {
					attackedSkillTemplate = attackedProbabilitySkillTempalte
				}
			}
		}
	}
	return
}

//获取复活技能
func getRebornAttackedSkill(defenceObject scene.BattleObject, isAttackPlayer bool) (attackedSkillTemplate *gametemplate.SkillTemplate) {
	attackedProbabilitySkillMap := defenceObject.GetSkills(skilltypes.SkillSecondTypeAttackedProbability)
	if len(attackedProbabilitySkillMap) != 0 {
		for _, attackedProbabilitySkill := range attackedProbabilitySkillMap {
			attackedProbabilitySkillTempalte := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(attackedProbabilitySkill.GetSkillId(), attackedProbabilitySkill.GetLevel())
			if isAttackPlayer {
				//不能触发
				if attackedProbabilitySkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypePlayer.Mask() == 0 {
					continue
				}
			} else {
				//不能触发
				if attackedProbabilitySkillTempalte.BeTriggered&skilltypes.SkillBeTriggerTypeMonster.Mask() == 0 {
					continue
				}
			}
			if attackedProbabilitySkillTempalte.GetRebornSkillTemplate() == nil {
				continue
			}
			if defenceObject.IsSkillInCd(int32(attackedProbabilitySkillTempalte.GetRebornSkillTemplate().TypeId)) {
				continue
			}
			if attackedProbabilitySkillTempalte.RebornSkillRate > 0 {
				return attackedProbabilitySkillTempalte
			}
		}
	}
	return
}

//获取天赋buff
func getTianFuBuff(bo scene.BattleObject, skillTemplate *gametemplate.SkillTemplate) (buffTianFuMap map[int32][]int32, buffRateMap map[int32]int32) {
	skillObj := bo.GetSkill(int32(skillTemplate.TemplateId()))
	if skillObj == nil {
		return
	}
	tianFuList := skillObj.GetTianFuList()
	if len(tianFuList) == 0 {
		return
	}
	buffTianFuMap = make(map[int32][]int32)
	buffRateMap = make(map[int32]int32)
	for _, tianFuInfo := range tianFuList {
		tianFuId := tianFuInfo.TianFuId
		level := tianFuInfo.Level

		tianFuTemplate := skillTemplate.GetTianFuTemplate(tianFuId)
		if tianFuTemplate == nil {
			continue
		}
		tianFuLevelTemplate := tianFuTemplate.GetTianFuLevelByLevel(level)
		if tianFuLevelTemplate == nil {
			continue
		}
		tempBuffRateMap := tianFuLevelTemplate.GetBuffRateMap()
		for buffId, buffRate := range tempBuffRateMap {
			if buffRateMap[buffId] < buffRate {
				buffRateMap[buffId] = buffRate
			}
		}
		tempBuffDongTaiMap := tianFuLevelTemplate.GetBuffDongTaiMap()
		for buffId, buffDongTaiId := range tempBuffDongTaiMap {
			buffTianFuMap[buffId] = append(buffTianFuMap[buffId], buffDongTaiId)
		}
	}
	return
}

func addTianFuBuff(bo scene.BattleObject, so scene.BattleObject, skillTemplate *gametemplate.SkillTemplate) {
	skillObj := so.GetSkill(int32(skillTemplate.TemplateId()))
	if skillObj == nil {
		return
	}
	buffTianFuMap, buffRateMap := getTianFuBuff(so, skillTemplate)
	for buffId, buffRate := range buffRateMap {
		buffTianFuList := buffTianFuMap[buffId]
		buffHit := mathutils.RandomHit(common.MAX_RATE, int(buffRate))
		if buffHit {
			//TODO 一般buff是不会永久保存的
			AddBuffWithTianFuList(bo, buffId, so.GetId(), int64(buffRate), buffTianFuList)
		}
	}

}

func GetSkillArea(so scene.BattleObject, skillTemplate *gametemplate.SkillTemplate) (area skilltypes.SkillArea) {
	area = skillTemplate.GetSkillArea()
	skillObj := so.GetSkill(int32(skillTemplate.TemplateId()))
	if skillObj == nil {
		return
	}
	tianFuList := skillObj.GetTianFuList()
	if len(tianFuList) == 0 {
		return
	}
	areaType := skilltypes.SkillAreaTypeDefault
	for _, tianFuInfo := range tianFuList {
		tianFuId := tianFuInfo.TianFuId
		level := tianFuInfo.Level

		tianFuTemplate := skillTemplate.GetTianFuTemplate(tianFuId)
		if tianFuTemplate == nil {
			continue
		}
		tianFuLevelTemplate := tianFuTemplate.GetTianFuLevelByLevel(level)
		if tianFuLevelTemplate == nil {
			continue
		}

		areaType = tianFuLevelTemplate.GetAreaType()
		if areaType == skilltypes.SkillAreaTypeDefault {
			continue
		}
		area = tianFuLevelTemplate.GetSkillArea()
	}
	return
}

func GetSpecialEffect(so scene.BattleObject,
	skillTemplate *gametemplate.SkillTemplate) (specialEffectType skilltypes.SkillSpecialEffectType,
	specialTarget int32,
	specialDistance float64,
	specialAnimationTime float64,
	specialTime float64,
) {
	specialEffectType = skillTemplate.GetSpecialEffectType()
	specialTarget = skillTemplate.SpecialTarget
	specialDistance = skillTemplate.GetSpecialDistance()
	specialAnimationTime = skillTemplate.GetSpecialAnimationTime()
	skillObj := so.GetSkill(int32(skillTemplate.TemplateId()))
	if skillObj == nil {
		return
	}
	tianFuList := skillObj.GetTianFuList()
	if len(tianFuList) == 0 {
		return
	}

	for _, tianFuInfo := range tianFuList {
		tianFuId := tianFuInfo.TianFuId
		level := tianFuInfo.Level

		tianFuTemplate := skillTemplate.GetTianFuTemplate(tianFuId)
		if tianFuTemplate == nil {
			continue
		}
		tianFuLevelTemplate := tianFuTemplate.GetTianFuLevelByLevel(level)
		if tianFuLevelTemplate == nil {
			continue
		}

		if tianFuLevelTemplate.GetSpecialEffectType() != 0 {
			specialEffectType = tianFuLevelTemplate.GetSpecialEffectType()
		}

		if tianFuLevelTemplate.GetSpecialAnimationTime() != 0 {
			specialDistance = tianFuLevelTemplate.GetSpecialDistance()
		}

		if tianFuLevelTemplate.GetSpecialTime() != 0 {
			specialTime = tianFuLevelTemplate.GetSpecialTime()
		}

		if tianFuLevelTemplate.SpecialTarget != 0 {
			specialTarget = tianFuLevelTemplate.SpecialTarget
		}
	}
	return
}

//伤害
func CostHP(defenceObject scene.BattleObject, damageInt int64, skillId int32, attackId int64, damageType scenetypes.DamageType) bool {
	dead := false

	if defenceObject.GetEffectNum(scenetypes.BuffEffectTypeHuDun)+defenceObject.GetHP() < damageInt {
		isAttackPlayer := false
		rebornSkillTemplate := getRebornAttackedSkill(defenceObject, isAttackPlayer)
		if rebornSkillTemplate != nil {
			hitFlag := mathutils.RandomHit(common.MAX_RATE, int(rebornSkillTemplate.RebornSkillRate))
			//随机
			if hitFlag {
				//死亡
				damageInt = defenceObject.GetEffectNum(scenetypes.BuffEffectTypeHuDun) + defenceObject.GetHP() - 1
				dead = defenceObject.CostHP(damageInt, attackId)
				scObjectDamage := pbutil.BuildSCObjectDamage(defenceObject, damageType, -damageInt, skillId, attackId)
				BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)

				defenceObject.AddHP(defenceObject.GetMaxHP())
				scObjectRebornDamage := pbutil.BuildSCObjectDamage(defenceObject, scenetypes.DamageTypeRecovery, defenceObject.GetMaxHP(), 0, 0)
				BroadcastNeighborIncludeSelf(defenceObject, scObjectRebornDamage)
				//加血
				Attack(defenceObject, defenceObject.GetPosition(), defenceObject.GetAngle(), rebornSkillTemplate.GetRebornSkillTemplate(), false)
				return false
			}
		}
	}
	dead = defenceObject.CostHP(damageInt, attackId)
	scObjectDamage := pbutil.BuildSCObjectDamage(defenceObject, damageType, -damageInt, skillId, attackId)
	BroadcastNeighborIncludeSelf(defenceObject, scObjectDamage)

	if dead {
		return defenceObject.Dead(attackId)
	}
	return dead
}
