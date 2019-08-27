package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/skill/types"
	"fmt"
)

type SkillTemplate struct {
	*SkillTemplateVO
	//技能类型
	skillFirstType types.SkillFirstType
	// 被动或主动
	skillSecondType types.SkillSecondType
	//静态或动态
	skillThirdType types.SkillThirdType
	//攻击或辅助
	skillFourthType types.SkillFourthType
	//cd组
	cdGroup *CdGroupTemplate

	ruleBreakType types.SkillRuleBreakType

	faceNeedType types.SkillFaceNeedType
	//技能范围
	areaType  types.SkillAreaType
	skillArea types.SkillArea

	//角色需求
	roleType playertypes.RoleType
	//攻击buff
	attackBuffTemplate *BuffTemplate
	//受攻击buff
	attackedBuffTemplate *BuffTemplate
	//自己buff
	selfBuffTemplate *BuffTemplate
	//加buff
	buffTemplate *BuffTemplate
	//添加buff
	buffMap map[int32]int32

	//加被动属性
	attrTemplate      *AttrTemplate
	specialEffectType types.SkillSpecialEffectType
	//属性光环(增益对象:仙盟成员)
	attrAuraTemplate *AttrTemplate

	//距离
	specialDistance float64
	//时间
	specialTime float64
	//表现时间
	specialAnimationTime float64
	//动态技能
	skillLevelTemplate *SkillLevelTemplate
	//限制被动技能触发
	limitTouchType types.SkillLimitTouchType
	//下一个等级技能
	nextSkillTemplate *SkillTemplate
	//damageAttack
	damageAttack float64

	//攻击距离
	attackDistance float64
	//动态技能等级map
	skillLevelTemplateMap map[int32]*SkillLevelTemplate
	//动态技能使用
	maxLevel int32
	//延迟时间
	delayTime int64
	//天赋
	tianFuTemplateMap map[int32]*TianFuTemplate
	tianFuTemplate    *TianFuTemplate
	//复活技能
	rebornSkillTemplate *SkillTemplate
}

func (st *SkillTemplate) GetSkillByLevel(level int32) *SkillLevelTemplate {
	if skillLevelTemplate, ok := st.skillLevelTemplateMap[level]; ok {
		return skillLevelTemplate
	}
	return nil
}

func (st *SkillTemplate) GetRebornSkillTemplate() *SkillTemplate {

	return st.rebornSkillTemplate
}

func (st *SkillTemplate) IsStatic() bool {
	return st.skillThirdType == types.SkillThirdTypeStatic
}

func (st *SkillTemplate) IsDynamic() bool {
	return st.skillThirdType == types.SkillThirdTypeDynamic
}

func (st *SkillTemplate) IsPassive() bool {
	return st.skillSecondType == types.SkillSecondTypePassive
}

func (st *SkillTemplate) IsPositive() bool {
	return st.skillSecondType == types.SkillSecondTypePositive
}

func (st *SkillTemplate) GetDamageAttack() float64 {
	return st.damageAttack
}

func (st *SkillTemplate) GetSkillFirstType() types.SkillFirstType {
	return st.skillFirstType
}

func (st *SkillTemplate) GetSkillSecondType() types.SkillSecondType {
	return st.skillSecondType
}

func (st *SkillTemplate) GetSkillThirdType() types.SkillThirdType {
	return st.skillThirdType
}

func (st *SkillTemplate) GetSkillFourthType() types.SkillFourthType {
	return st.skillFourthType
}

func (st *SkillTemplate) GetCdGroup() *CdGroupTemplate {
	return st.cdGroup
}

func (st *SkillTemplate) GetAttackDistance() float64 {
	return st.attackDistance
}

func (st *SkillTemplate) GetRuleBreakType() types.SkillRuleBreakType {
	return st.ruleBreakType
}

func (st *SkillTemplate) GetFaceNeedType() types.SkillFaceNeedType {
	return st.faceNeedType
}

func (st *SkillTemplate) GetAreaType() types.SkillAreaType {
	return st.areaType
}

func (st *SkillTemplate) GetSkillArea() types.SkillArea {
	return st.skillArea
}

func (st *SkillTemplate) IsLimitRole() bool {
	return st.ProNeed != 0
}

func (st *SkillTemplate) GetRoleType() playertypes.RoleType {
	return st.roleType
}

func (st *SkillTemplate) GetBuffMap() map[int32]int32 {
	return st.buffMap
}

func (st *SkillTemplate) GetAttackBuffTemplate() *BuffTemplate {
	return st.attackBuffTemplate
}

func (st *SkillTemplate) GetAttackedBuffTemplate() *BuffTemplate {
	return st.attackedBuffTemplate
}

func (st *SkillTemplate) GetSelfBuffTemplate() *BuffTemplate {
	return st.selfBuffTemplate
}

func (st *SkillTemplate) GetAttrTemplate() *AttrTemplate {
	return st.attrTemplate
}

func (st *SkillTemplate) GetAttrAuraTemplate() *AttrTemplate {
	return st.attrAuraTemplate
}

func (st *SkillTemplate) GetSpecialEffectType() types.SkillSpecialEffectType {
	return st.specialEffectType
}

func (st *SkillTemplate) GetSpecialTime() float64 {
	return st.specialTime
}
func (st *SkillTemplate) GetSpecialAnimationTime() float64 {
	return st.specialAnimationTime
}

func (st *SkillTemplate) GetSpecialDistance() float64 {
	return st.specialDistance
}

func (st *SkillTemplate) GetSkillLevelTemplate() *SkillLevelTemplate {
	return st.skillLevelTemplate
}

func (st *SkillTemplate) GetLimitTouchType() types.SkillLimitTouchType {
	return st.limitTouchType
}

func (st *SkillTemplate) GetNextSkillTemplate() *SkillTemplate {
	return st.nextSkillTemplate
}

func (st *SkillTemplate) GetMaxLevel() int32 {
	return st.maxLevel
}

func (st *SkillTemplate) GetDelayTime() int64 {
	return st.delayTime
}

func (st *SkillTemplate) TemplateId() int {
	return st.Id
}

func (st *SkillTemplate) GetTianFuTemplate(tianFuId int32) *TianFuTemplate {
	if len(st.tianFuTemplateMap) == 0 {
		return nil
	}
	tianFuTemplate, ok := st.tianFuTemplateMap[tianFuId]
	if !ok {
		return nil
	}
	return tianFuTemplate
}

func (st *SkillTemplate) PatchAfterCheck() {
	//动态技能
	if st.skillThirdType == types.SkillThirdTypeDynamic {
		st.skillLevelTemplateMap = make(map[int32]*SkillLevelTemplate)
		//赋值skillLevelTemplateMap
		for levelTemplate := st.skillLevelTemplate; levelTemplate != nil; levelTemplate = levelTemplate.nextSkillLevelTemplate {
			level := levelTemplate.Level
			st.skillLevelTemplateMap[level] = levelTemplate
			st.maxLevel = level
		}
	}

	//天赋
	if st.tianFuTemplate != nil {
		st.tianFuTemplateMap = make(map[int32]*TianFuTemplate)
		//赋值tianFuTemplateMap
		for tempTemplate := st.tianFuTemplate; tempTemplate != nil; tempTemplate = tempTemplate.nextTianFuTemplate {
			tianFuId := int32(tempTemplate.TemplateId())
			st.tianFuTemplateMap[tianFuId] = tempTemplate
		}
	}
}
func (st *SkillTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()

	st.skillFirstType = types.SkillFirstType(st.FirstType)
	if !st.skillFirstType.Valid() {
		return fmt.Errorf("firstType [%d] invalid", st.FirstType)
	}
	st.skillSecondType = types.SkillSecondType(st.SecondType)
	if !st.skillSecondType.Valid() {
		return fmt.Errorf("secondType [%d] invalid", st.SecondType)
	}

	if st.BuffId != 0 {

		// to := template.GetTemplateService().Get(st.BuffId, (*BuffTemplate)(nil))
		// if to == nil {
		// 	err = fmt.Errorf("buffid [%d] should valid", st.BuffId)
		// 	return template.NewTemplateFieldError("buffId", err)
		// }
		// st.attackBuffTemplate = to.(*BuffTemplate)
	}

	if st.BuffId2 != 0 {
		to := template.GetTemplateService().Get(st.BuffId2, (*BuffTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("buffid2 [%d] should valid", st.BuffId2)
			return template.NewTemplateFieldError("buffId2", err)
		}
		st.attackedBuffTemplate = to.(*BuffTemplate)
	}

	if st.BuffId3 != 0 {
		to := template.GetTemplateService().Get(int(st.BuffId3), (*BuffTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("buffid3 [%d] should valid", st.BuffId3)
			return template.NewTemplateFieldError("buffId3", err)
		}
		st.selfBuffTemplate = to.(*BuffTemplate)
	}

	st.skillThirdType = types.SkillThirdType(st.ThirdType)
	if !st.skillThirdType.Valid() {
		return fmt.Errorf("thirdType [%d] invalid", st.ThirdType)
	}
	st.skillFourthType = types.SkillFourthType(st.FourthType)
	if !st.skillFourthType.Valid() {
		return fmt.Errorf("fourthType [%d] invalid", st.FourthType)
	}

	if st.skillSecondType != types.SkillSecondTypePassive {
		//验证cd组
		tempCdGroupTemplate := template.GetTemplateService().Get(int(st.CdGroup), (*CdGroupTemplate)(nil))
		if tempCdGroupTemplate == nil {
			return fmt.Errorf("cdGroup [%d] invalid", st.CdGroup)
		}
		cdGroupTemplate, ok := tempCdGroupTemplate.(*CdGroupTemplate)
		if !ok {
			return fmt.Errorf("cdGroup [%d] invalid", st.CdGroup)
		}
		st.cdGroup = cdGroupTemplate

		//无视安全区域
		st.ruleBreakType = types.SkillRuleBreakType(st.RuleBreak)
		if !st.ruleBreakType.Valid() {
			return fmt.Errorf("ruleBreak [%d] invalid", st.RuleBreak)
		}
		//朝向
		st.faceNeedType = types.SkillFaceNeedType(st.FaceNeed)
		if !st.faceNeedType.Valid() {
			return fmt.Errorf("faceNeed [%d] invalid", st.FaceNeed)
		}

		//区域类型
		st.areaType = types.SkillAreaType(st.AreaType)
		if !st.areaType.Valid() {
			return fmt.Errorf("areaType [%d] invalid", st.AreaType)
		}
		//组成技能范围
		switch st.areaType {
		case types.SkillAreaTypeLine:
			st.skillArea = types.NewSkillAreaRectangle(float64(st.AreaRange)/common.MILL_METER, float64(st.AreaRadius)/common.MILL_METER)
		case types.SkillAreaTypeFan:
			st.skillArea = types.NewSkillAreaFan(float64(st.AreaRange), float64(st.AreaRadius)/common.MILL_METER)
		case types.SkillAreaTypeRound:
			st.skillArea = types.NewSkillAreaRound(float64(st.AreaRadius) / common.MILL_METER)
		}
	}
	st.attackDistance = float64(st.Distance) / common.MILL_METER
	st.damageAttack = float64(st.SpellDamage) / common.MAX_RATE * float64(st.SpellPower) / common.MAX_RATE

	if st.ProNeed != 0 {
		st.roleType = playertypes.RoleType(st.ProNeed)
		if !st.roleType.Valid() {
			return fmt.Errorf("proNeed [%d] invalid", st.ProNeed)
		}
	}
	addIdArr, err := utils.SplitAsIntArray(st.AddStatus)
	if err != nil {
		return template.NewTemplateFieldError("AddStatus", err)
	}
	addRateArr, err := utils.SplitAsIntArray(st.AddStatusRate)
	if err != nil {
		return template.NewTemplateFieldError("AddStatusRate", err)
	}
	if len(addIdArr) != len(addRateArr) {
		err = fmt.Errorf("AddStatus[%s]AddStatusRate[%s]数量不一致", st.AddStatus, st.AddStatusRate)
		return template.NewTemplateFieldError("subId,subRate", err)
	}
	st.buffMap = make(map[int32]int32)
	for i, subId := range addIdArr {
		st.buffMap[subId] = addRateArr[i]
	}

	//被动属性
	if st.AddAttrId != 0 {
		tempAttrTemplate := template.GetTemplateService().Get(int(st.AddAttrId), (*AttrTemplate)(nil))
		if tempAttrTemplate == nil {
			return fmt.Errorf("addAttrId [%d] invalid", st.AddAttrId)
		}
		attrTemplate, ok := tempAttrTemplate.(*AttrTemplate)
		if !ok {
			return fmt.Errorf("addAttrId [%d] invalid", st.AddAttrId)
		}
		st.attrTemplate = attrTemplate
	}
	//属性光环
	if st.UnionAttrId != 0 {
		tempAttrTemplate := template.GetTemplateService().Get(int(st.UnionAttrId), (*AttrTemplate)(nil))
		if tempAttrTemplate == nil {
			return fmt.Errorf("UnionAttrId [%d] invalid", st.UnionAttrId)
		}
		attrAuraTemplate, ok := tempAttrTemplate.(*AttrTemplate)
		if !ok {
			return fmt.Errorf("UnionAttrId [%d] invalid", st.UnionAttrId)
		}
		st.attrAuraTemplate = attrAuraTemplate
	}

	//特殊效果
	st.specialEffectType = types.SkillSpecialEffectType(st.SpecialEffect)
	if !st.specialEffectType.Valid() {
		return fmt.Errorf("specialEffect [%d] invalid", st.SpecialEffect)
	}

	st.specialDistance = float64(st.SpecialEffectValue) / common.MILL_METER
	st.specialTime = float64(st.SpecialEffectValue2) / float64(common.SECOND)
	st.specialAnimationTime = float64(st.SpecialEffectValue3) / float64(common.SECOND)
	//动态技能
	if st.skillThirdType == types.SkillThirdTypeDynamic {
		if st.SpellUpgradeBeginId == 0 {
			return fmt.Errorf("spellUpgradeBeginId [%d] invalid", st.SpellUpgradeBeginId)
		}

		tempSkillLevelTemplate := template.GetTemplateService().Get(int(st.SpellUpgradeBeginId), (*SkillLevelTemplate)(nil))
		if tempSkillLevelTemplate == nil {
			return fmt.Errorf("spellUpgradeBeginId [%d] invalid", st.SpellUpgradeBeginId)
		}
		skillLevelTemplate, ok := tempSkillLevelTemplate.(*SkillLevelTemplate)
		if !ok {
			return fmt.Errorf("spellUpgradeBeginId [%d] invalid", st.SpellUpgradeBeginId)
		}
		if skillLevelTemplate.Level != 1 {
			return fmt.Errorf("spellUpgradeBeginId [%d] invalid", st.SpellUpgradeBeginId)
		}
		st.skillLevelTemplate = skillLevelTemplate

	}

	//限制被动技能
	st.limitTouchType = types.SkillLimitTouchType(st.LimitTouch)
	if !st.limitTouchType.Valid() {
		return fmt.Errorf("limitTouch [%d] invalid", st.LimitTouch)
	}

	//下一个等级
	if st.SpellNext != 0 {
		tempNextSkillTemplate := template.GetTemplateService().Get(int(st.SpellNext), (*SkillTemplate)(nil))
		if tempNextSkillTemplate == nil {
			return fmt.Errorf("spellNext [%d] invalid", st.SpellNext)
		}
		nextSkillTemplate, ok := tempNextSkillTemplate.(*SkillTemplate)
		if !ok {
			return fmt.Errorf("spellNext [%d] invalid", st.SpellNext)
		}
		st.nextSkillTemplate = nextSkillTemplate
	}
	delayTimeArr, err := utils.SplitAsFloatArray(st.DelayTimesServer)
	if err != nil {
		return template.NewTemplateFieldError("DelayTimesServer", err)
	}
	if len(delayTimeArr) > 0 {
		st.delayTime = int64(delayTimeArr[0] * float64(common.SECOND))
	}

	if st.TianFuBeginId != 0 {
		tempTianFuTemplate := template.GetTemplateService().Get(int(st.TianFuBeginId), (*TianFuTemplate)(nil))
		if tempTianFuTemplate == nil {
			return fmt.Errorf("tianFuBeginId [%d] invalid", st.TianFuBeginId)
		}
		st.tianFuTemplate = tempTianFuTemplate.(*TianFuTemplate)
	}

	if st.RebornSkill != 0 {
		tempRebornSkillTemplate := template.GetTemplateService().Get(int(st.RebornSkill), (*SkillTemplate)(nil))
		if tempRebornSkillTemplate == nil {
			return fmt.Errorf("RebornSkill [%d] invalid", st.RebornSkill)
		}
		st.rebornSkillTemplate = tempRebornSkillTemplate.(*SkillTemplate)
	}
	return nil
}

func (st *SkillTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()

	st.skillFourthType = types.SkillFourthType(st.FourthType)
	if st.skillFourthType == types.SkillFourthTypeAttack {
		if st.TargetSelect&types.SkillTargetSelectTypeSelf.Mask() != 0 {
			return fmt.Errorf("TargetSelect [%d] invalid", st.TargetSelect)
		}
	}

	//职业技能必须是动态技能
	if st.skillFirstType == types.SkillFirstTypeRole {
		if st.skillThirdType != types.SkillThirdTypeDynamic {
			return fmt.Errorf("type_3rd [%d] invalid", st.ThirdType)
		}
	}
	if st.skillSecondType != types.SkillSecondTypePassive {
		//触发几率
		if err = validator.MinValidate(float64(st.Rate), float64(0), true); err != nil {
			return template.NewTemplateFieldError("Rate", err)
		}
		//cd时间
		if err = validator.MinValidate(float64(st.CdTime), float64(0), true); err != nil {
			return template.NewTemplateFieldError("CdTime", err)
		}

		//施法距离
		if err = validator.MinValidate(float64(st.Distance), float64(0), true); err != nil {
			return template.NewTemplateFieldError("Distance", err)

		}

		//作用半径
		if err = validator.MinValidate(float64(st.AreaRadius), float64(0), true); err != nil {
			return template.NewTemplateFieldError("AreaRadius", err)
		}

		//作用范围0-360
		if st.areaType == types.SkillAreaTypeFan || st.areaType == types.SkillAreaTypeRound {
			if err = validator.RangeValidate(float64(st.AreaRange), float64(common.MIN_ANGLE), true, float64(common.MAX_ANGLE), true); err != nil {
				return template.NewTemplateFieldError("AreaRange", err)
			}
		} else if st.areaType == types.SkillAreaTypeLine {
			if err = validator.MinValidate(float64(st.AreaRange), float64(0), false); err != nil {
				return template.NewTemplateFieldError("AreaRange", err)
			}
		}
		//目标数量
		if err = validator.MinValidate(float64(st.TargetCount), float64(0), false); err != nil {
			return template.NewTemplateFieldError("TargetCount", err)

		}

		//等级需要
		if err = validator.MinValidate(float64(st.LevNeed), float64(0), false); err != nil {
			return template.NewTemplateFieldError("LevNeed", err)

		}
		// //消耗生命值
		// if err = validator.MinValidate(float64(st.CostHpValue), float64(0), true); err != nil {
		// 	return
		// }

		// //消耗生命万分比
		// if err = validator.RangeValidate(float64(st.CostHpPersent), float64(0), true, float64(common.MAX_RATE), false); err != nil {
		// 	return
		// }
		// //消耗体力值
		// if err = validator.MinValidate(float64(st.CostTpValue), float64(0), true); err != nil {
		// 	return
		// }

		//伤害固定值
		if err = validator.MinValidate(float64(st.DamageValueBase), float64(0), true); err != nil {
			return template.NewTemplateFieldError("DamageValueBase", err)

		}
		//伤害固定值成长值
		if err = validator.MinValidate(float64(st.DamageValue), float64(0), true); err != nil {
			return template.NewTemplateFieldError("DamageValue", err)

		}
		//技能威力值
		if err = validator.MinValidate(float64(st.SpellDamage), float64(0), true); err != nil {
			return template.NewTemplateFieldError("SpellDamage", err)
		}
		//技能伤害
		if err = validator.MinValidate(float64(st.SpellPower), float64(0), true); err != nil {
			return template.NewTemplateFieldError("SpellPower", err)
		}
		//技能伤害万分比
		if err = validator.MinValidate(float64(st.DamagePersent), float64(0), true); err != nil {
			return template.NewTemplateFieldError("DamagePersent", err)
		}
		//治疗固定值
		if err = validator.MinValidate(float64(st.CureValue), float64(0), true); err != nil {
			return template.NewTemplateFieldError("CureValue", err)
		}
		//治疗万分比
		if err = validator.MinValidate(float64(st.CurePersent), float64(0), true); err != nil {
			return template.NewTemplateFieldError("CurePersent", err)
		}

		//增加仇恨值
		if err = validator.MinValidate(float64(st.HatredValue), float64(0), true); err != nil {
			return template.NewTemplateFieldError("HatredValue", err)
		}
		//仇恨万分比
		if err = validator.MinValidate(float64(st.HatredPersent), float64(0), true); err != nil {
			return template.NewTemplateFieldError("HatredPersent", err)
		}

		//增加命中
		if err = validator.MinValidate(float64(st.AddHit), float64(0), true); err != nil {
			return template.NewTemplateFieldError("AddHit", err)
		}

		//增加暴击
		if err = validator.MinValidate(float64(st.AddCritical), float64(0), true); err != nil {
			return template.NewTemplateFieldError("AddCritical", err)
		}

		//特殊效果概率
		if err = validator.MinValidate(float64(st.SpecialEffectRate), float64(0), true); err != nil {
			return template.NewTemplateFieldError("SpecialEffectRate", err)
		}

		//特殊效果距离
		if err = validator.MinValidate(float64(st.SpecialEffectValue), float64(0), true); err != nil {
			return template.NewTemplateFieldError("SpecialEffectValue", err)
		}
		//特殊效果时间
		if err = validator.MinValidate(float64(st.SpecialEffectValue2), float64(0), true); err != nil {
			return template.NewTemplateFieldError("SpecialEffectValue2", err)
		}
		//战力增加
		if err = validator.MinValidate(float64(st.AddPower), float64(0), true); err != nil {
			return template.NewTemplateFieldError("AddPower", err)
		}

		//消耗物品数量
		if err = validator.MinValidate(float64(st.ConsumeGoodsCount), float64(0), true); err != nil {
			return template.NewTemplateFieldError("ConsumeGoodsCount", err)
		}
		//消耗经验
		if err = validator.MinValidate(float64(st.ConsumeExperience), float64(0), true); err != nil {
			return template.NewTemplateFieldError("ConsumeExperience", err)
		}
		//消耗银两
		if err = validator.MinValidate(float64(st.ConsumeMoney), float64(0), true); err != nil {
			return template.NewTemplateFieldError("ConsumeMoney", err)
		}

		for buffId, buffRate := range st.buffMap {
			to := template.GetTemplateService().Get(int(buffId), (*BuffTemplate)(nil))
			if to == nil {
				err = fmt.Errorf("AddStatus [%s] should valid", st.AddStatus)
				return template.NewTemplateFieldError("AddStatus", err)
			}
			//状态概率
			if err = validator.MinValidate(float64(buffRate), float64(0), true); err != nil {
				err = fmt.Errorf("AddStatusRate [%s] should valid", st.AddStatusRate)
				return template.NewTemplateFieldError("AddStatusRate", err)
			}
		}
	}

	if st.TianFuBeginId != 0 {
		tempTianFuTemplate := template.GetTemplateService().Get(int(st.TianFuBeginId), (*TianFuTemplate)(nil))
		if tempTianFuTemplate == nil {
			return fmt.Errorf("TianFuBeginId [%d] invalid", st.TianFuBeginId)
		}
	}

	if st.rebornSkillTemplate != nil {
		if st.rebornSkillTemplate.GetSkillFirstType() != types.SkillFirstTypeSubSkill {
			return fmt.Errorf("RebornSkill [%d] invalid", st.RebornSkill)
		}
	}

	return nil
}

func (mt *SkillTemplate) FileName() string {
	return "tb_skill.json"
}

func init() {
	template.Register((*SkillTemplate)(nil))
}
