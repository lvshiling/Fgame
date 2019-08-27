package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	collecttypes "fgame/fgame/game/collect/types"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/types"
	scenetypes "fgame/fgame/game/scene/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"strconv"
)

/*场景部怪配置*/
type BiologyTemplate struct {
	*BiologyTemplateVO
	//生物
	biologySetType types.BiologySetType
	//生物类型
	biologyType types.BiologyType
	//ai类型
	biologyScriptType types.BiologyScriptType
	//追击范围平方
	randradiusSquare float64
	//阵营
	factionType types.FactionType
	//主动怪 被动怪
	threatType types.ThreatType
	//战斗属性
	battlePropertyMap map[int32]int64
	//基本攻击技能
	attackSkill *SkillTemplate
	//技能1
	skill1 *SkillTemplate
	//技能2
	skill2 *SkillTemplate
	//技能列表
	skills []int32
	//所有技能
	allSkills []int32
	//主动技能
	positiveSkills []int
	//主动触发技能
	attackProbabilitySkills []int
	//被动触发技能
	attackedProbabilitySkills []int
	//技能触发概率
	skillRateMap    map[int32]int32
	minAttackRadius float64
	//最小攻击距离
	minAttackRadiusSquare float64
	//警戒范围
	alertRadiusSquare float64
	//脱离战斗是否回满血
	biologyAutoRecoverType types.BiologyAutoRecoverType
	//重生类型
	biologyRebornType types.BiologyRebornType
	//重生时间
	rebornTime int64
	//掉落组合
	dropIdList []int32
	//采集物被采集选的掉落id组合
	caijiChooseDropIdList map[collecttypes.CollectChooseFinishType]int32

	//掉落类型
	dropType scenetypes.DropType
	//掉落判断
	dropJudgeType scenetypes.DropJudgeType
	//buff
	buffIdList []int32
	//活动id
	groupIdList []int32

	portalTemplate *PortalTemplate
	//采集点采集完成是否消失
	collectPointSult collecttypes.CollectPointFinishType
	miZangTemplate   *BossMiZangTemplate
}

func (bt *BiologyTemplate) TemplateId() int {
	return bt.Id
}

func (bt *BiologyTemplate) GetFactionType() types.FactionType {
	return bt.factionType
}

func (bt *BiologyTemplate) GetRebornTime(now int64) int64 {
	switch bt.biologyRebornType {
	case scenetypes.BiologyRebornTypeSecond:
		{
			return bt.rebornTime
		}
	case scenetypes.BiologyRebornTypeTime:
		{
			begin, _ := timeutils.BeginOfNow(now)
			return bt.rebornTime + begin
		}
	}
	return bt.rebornTime
}

func (bt *BiologyTemplate) GetRebornType() types.BiologyRebornType {
	return bt.biologyRebornType
}

func (bt *BiologyTemplate) GetCollectPointFinishType() collecttypes.CollectPointFinishType {
	return bt.collectPointSult
}

func (bt *BiologyTemplate) GetPositiveSkills() []int {
	return bt.positiveSkills
}

func (bt *BiologyTemplate) GetSkills() []int32 {
	return bt.skills
}

func (bt *BiologyTemplate) GetAllSkills() []int32 {
	return bt.allSkills
}

func (bt *BiologyTemplate) GetDropIdList() []int32 {
	return bt.dropIdList
}

func (bt *BiologyTemplate) GetCaiJiChooseDropId(typ collecttypes.CollectChooseFinishType) (dropId int32, ok bool) {
	dropId, ok = bt.caijiChooseDropIdList[typ]
	return
}

func (bt *BiologyTemplate) GetAttackSkill() int32 {
	if bt.attackSkill == nil {
		return int32(0)
	}
	return int32(bt.attackSkill.Id)
}

func (bt *BiologyTemplate) GetSkillRate(skillId int32) int32 {
	w, ok := bt.skillRateMap[skillId]
	if !ok {
		return 0
	}
	return w
}

func (bt *BiologyTemplate) GetAttackedProbabilitySkills() []int {
	return bt.attackedProbabilitySkills
}

func (bt *BiologyTemplate) GetAttackProbalitySkills() []int {
	return bt.attackProbabilitySkills
}

func (bt *BiologyTemplate) GetBattlePropertyMap() map[int32]int64 {
	return bt.battlePropertyMap
}

func (bt *BiologyTemplate) IsPositive() bool {
	return bt.threatType == types.ThreateTypePositive
}

//获取追击距离
func (bt *BiologyTemplate) GetRandradiusSquare() float64 {
	return bt.randradiusSquare
}

func (bt *BiologyTemplate) GetAlertRadiusSquare() float64 {
	return bt.alertRadiusSquare
}

//获取最小攻击距离
func (bt *BiologyTemplate) GetMinAttackRadiusSquare() float64 {
	return bt.minAttackRadiusSquare
}

//获取最小攻击距离
func (bt *BiologyTemplate) GetMinAttackRadius() float64 {
	return bt.minAttackRadius
}

//获取脱离战斗是否回血
func (bt *BiologyTemplate) IsRecover() bool {
	return bt.biologyAutoRecoverType == types.BiologyAutoRecoverTypeYes
}

func (bt *BiologyTemplate) GetBiologySetType() types.BiologySetType {
	return bt.biologySetType
}

func (bt *BiologyTemplate) GetBiologyType() types.BiologyType {
	return bt.biologyType
}

func (bt *BiologyTemplate) GetBiologyScriptType() types.BiologyScriptType {
	return bt.biologyScriptType
}

func (bt *BiologyTemplate) CanReborn() bool {
	if bt.biologyRebornType != types.BiologyRebornTypeSecond {
		return true
	}
	if bt.rebornTime != 0 {
		return true
	}
	return false
}

func (bt *BiologyTemplate) GetDropType() scenetypes.DropType {
	return bt.dropType
}

func (bt *BiologyTemplate) GetDropJudgeType() scenetypes.DropJudgeType {
	return bt.dropJudgeType
}

func (bt *BiologyTemplate) GetBuffIdList() []int32 {
	return bt.buffIdList
}

func (bt *BiologyTemplate) GetGroupIdList() []int32 {
	return bt.groupIdList
}

func (bt *BiologyTemplate) GetPortalTemplate() *PortalTemplate {
	return bt.portalTemplate
}

func (bt *BiologyTemplate) GetMizangTemplate() *BossMiZangTemplate {
	return bt.miZangTemplate
}

func (bt *BiologyTemplate) PatchAfterCheck() {

}

func (bt *BiologyTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(bt.FileName(), bt.TemplateId(), err)
			return
		}
	}()
	bt.biologySetType = types.BiologySetType(bt.SetType)
	if !bt.biologySetType.Valid() {
		err = template.NewTemplateFieldError("SetType", fmt.Errorf("%d is invalid", bt.SetType))
		return
	}

	bt.biologyType = types.BiologyType(bt.Type)
	if !bt.biologyType.Valid() {
		err = template.NewTemplateFieldError("type", fmt.Errorf("%d is invalid", bt.Type))
		return
	}

	bt.biologyScriptType = types.BiologyScriptType(bt.ScriptType)
	if !bt.biologyType.Valid() {
		err = template.NewTemplateFieldError("ScriptType", fmt.Errorf("%d is invalid", bt.ScriptType))
		return
	}

	bt.factionType = types.FactionType(bt.Faction)

	if !bt.factionType.Valid() {
		err = template.NewTemplateFieldError("faction", fmt.Errorf("faction is invalid"))
		return
	}

	bt.threatType = types.ThreatType(bt.ThreatType)
	if !bt.threatType.Valid() {
		err = template.NewTemplateFieldError("threatType", fmt.Errorf("threatType is invalid"))
		return
	}

	bt.biologyAutoRecoverType = types.BiologyAutoRecoverType(bt.AutoRecoverMaxhealth)
	if !bt.biologyAutoRecoverType.Valid() {
		err = fmt.Errorf("%d invalid", bt.AutoRecoverMaxhealth)
		err = template.NewTemplateFieldError("autoRecoverMaxhealth", err)
		return
	}
	if bt.AttackId != 0 {

		to := template.GetTemplateService().Get(bt.AttackId, (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("%d no exist", bt.AttackId)
			err = template.NewTemplateFieldError("attack", err)
			return
		}

		bt.attackSkill, _ = to.(*SkillTemplate)

	}

	if bt.SkillId1 != 0 {
		to := template.GetTemplateService().Get(bt.SkillId1, (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("%d no exist", bt.SkillId1)
			err = template.NewTemplateFieldError("skillId1", err)
			return
		}
		bt.skill1, _ = to.(*SkillTemplate)

	}

	if bt.SkillId2 != 0 {
		to := template.GetTemplateService().Get(bt.SkillId2, (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("%d no exist", bt.SkillId2)
			err = template.NewTemplateFieldError("skillId2", err)
			return
		}
		bt.skill2, _ = to.(*SkillTemplate)

	}

	bt.randradiusSquare = float64(bt.Randradius * bt.Randradius)
	bt.alertRadiusSquare = float64(bt.Alertradius * bt.Alertradius)

	bt.battlePropertyMap = make(map[int32]int64)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeMaxHP)] = int64(bt.Hp)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeAttack)] = int64(bt.Attack)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeDefend)] = int64(bt.Defence)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeCrit)] = int64(bt.Critical)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeTough)] = int64(bt.Tough)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeAbnormality)] = int64(bt.Abnormality)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeBlock)] = int64(bt.Block)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeDodge)] = int64(bt.Dodge)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeHit)] = int64(bt.Hit)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeHuanYunAttack)] = int64(bt.HunyuanAtt)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeHuanYunDef)] = int64(bt.HunyuanDef)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeMoveSpeed)] = int64(bt.SpeedMove)
	bt.battlePropertyMap[int32(propertytypes.BattlePropertyTypeForce)] = int64(bt.Force)

	//掉落组合
	dropIdList, err := coreutils.SplitAsIntArray(bt.DropCombine)
	if err != nil {
		err = template.NewTemplateFieldError("DropCombine", err)
		return
	}
	bt.dropIdList = dropIdList

	//采集物被采集选的掉落id
	bt.caijiChooseDropIdList = make(map[collecttypes.CollectChooseFinishType]int32)
	chooseDropIdList, err := coreutils.SplitAsIntArray(bt.CaiJiChooseDrop)
	if err != nil {
		err = template.NewTemplateFieldError("CaiJiChooseDrop", err)
		return
	}
	for typeInt, val := range chooseDropIdList {
		finishType := collecttypes.CollectChooseFinishType(typeInt)
		if !finishType.Valid() {
			err = template.NewTemplateFieldError("CaiJiChooseDrop Type limit", err)
			return
		}
		bt.caijiChooseDropIdList[finishType] = val
	}

	bt.collectPointSult = collecttypes.CollectPointFinishType(bt.CaiJiIsXiaoShi)
	if !bt.collectPointSult.Valid() {
		err = template.NewTemplateFieldError("CaiJiIsXiaoShi", fmt.Errorf("%d is invalid", bt.CaiJiIsXiaoShi))
		return
	}

	bt.biologyRebornType = types.BiologyRebornType(bt.RebornType)
	if !bt.biologyRebornType.Valid() {
		err = fmt.Errorf("%d invalid", bt.RebornType)
		err = template.NewTemplateFieldError("rebornType", err)
		return
	}
	switch bt.biologyRebornType {
	case scenetypes.BiologyRebornTypeTime:
		{
			rebornTime, err := timeutils.ParseDayOfHHMM(bt.RebornTime)
			if err != nil {
				err = template.NewTemplateFieldError("rebornTime", err)
				return err
			}
			bt.rebornTime = rebornTime
		}
	case scenetypes.BiologyRebornTypeSecond:
		{
			rebornTime, err := strconv.ParseInt(bt.RebornTime, 10, 64)
			if err != nil {
				err = template.NewTemplateFieldError("rebornTime", err)
				return err
			}
			bt.rebornTime = rebornTime
		}
	}

	bt.dropType = scenetypes.DropType(bt.DropType)
	bt.dropJudgeType = scenetypes.DropJudgeType(bt.DropOwnerType)
	bt.buffIdList, err = coreutils.SplitAsIntArray(bt.BuffIds)
	if err != nil {
		err = template.NewTemplateFieldError("BuffIds", err)
		return
	}

	bt.groupIdList, err = coreutils.SplitAsIntArray(bt.ActivityGroupId)
	if err != nil {
		err = template.NewTemplateFieldError("ActivityGroupId", err)
		return
	}
	if bt.PortalId != 0 {
		tempPortalTemplate := template.GetTemplateService().Get(int(bt.PortalId), (*PortalTemplate)(nil))
		if tempPortalTemplate == nil {
			err = fmt.Errorf("[%d] invalid", bt.PortalId)
			err = template.NewTemplateFieldError("PortalId", err)
			return
		}
		bt.portalTemplate = tempPortalTemplate.(*PortalTemplate)
	}

	if bt.MiZangId != 0 {
		tempMiZangTemplate := template.GetTemplateService().Get(int(bt.MiZangId), (*BossMiZangTemplate)(nil))
		if tempMiZangTemplate == nil {
			err = fmt.Errorf("[%d] invalid", bt.MiZangId)
			err = template.NewTemplateFieldError("MiZangId", err)
			return
		}
		bt.miZangTemplate = tempMiZangTemplate.(*BossMiZangTemplate)
	}
	return nil
}

const (
	attackRadiusDeviation = 0.1
)

//获取生物站的距离
func (bt *BiologyTemplate) GetMinStandRadius() float64 {
	return bt.minAttackRadius - attackRadiusDeviation
}

//是否免疫
func (bt *BiologyTemplate) IsImmune(effect skilltypes.SkillSpecialEffectType) bool {
	switch effect {
	case skilltypes.SkillSpecialEffectTypeClose:
		if (bt.IsJitui & int32(skilltypes.SkillImmuneTypeClose)) != 0 {
			return true
		}
		break
	case skilltypes.SkillSpecialEffectTypeRepel:
		if (bt.IsJitui & int32(skilltypes.SkillImmuneTypeRepel)) != 0 {
			return true
		}
	default:
		return true
	}
	return false
}

//检查
func (bt *BiologyTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(bt.FileName(), bt.TemplateId(), err)
			return
		}
	}()
	if bt.attackSkill != nil {
		if bt.attackSkill.IsDynamic() {
			err = fmt.Errorf("%d should be static", bt.AttackId)
			err = template.NewTemplateFieldError("attack", err)
			return
		}
		if !bt.attackSkill.IsPositive() {

			err = fmt.Errorf("%d should be positive ", bt.AttackId)
			err = template.NewTemplateFieldError("attack", err)
			return
		}
	}
	if bt.skill1 != nil {
		if bt.skill1.IsDynamic() {
			err = fmt.Errorf("%d should be static", bt.SkillId1)
			err = template.NewTemplateFieldError("skillId1", err)
			return
		}
		if bt.skill1.IsPassive() {
			err = fmt.Errorf("%d should not be passive", bt.SkillId1)
			err = template.NewTemplateFieldError("skillId1", err)
			return
		}
	}
	if bt.skill2 != nil {
		if bt.skill2.IsDynamic() {
			err = fmt.Errorf("%d should be static", bt.SkillId2)
			err = template.NewTemplateFieldError("skillId2", err)
			return
		}
		if bt.skill2.IsPassive() {
			err = fmt.Errorf("%d should not be passive", bt.SkillId2)
			err = template.NewTemplateFieldError("skillId2", err)
			return
		}
	}
	minAttackRadius := float64(0)
	if bt.attackSkill != nil {
		minAttackRadius = bt.attackSkill.GetAttackDistance()
		bt.allSkills = append(bt.allSkills, int32(bt.attackSkill.TemplateId()))
	}
	bt.skillRateMap = make(map[int32]int32)
	if bt.skill1 != nil {
		bt.skills = append(bt.skills, int32(bt.skill1.Id))
		bt.allSkills = append(bt.allSkills, int32(bt.skill1.Id))

		switch bt.skill1.GetSkillSecondType() {
		case skilltypes.SkillSecondTypePositive:
			bt.positiveSkills = append(bt.positiveSkills, bt.skill1.Id)
			bt.skillRateMap[int32(bt.skill1.Id)] = int32(bt.SkillRate1)

			if bt.skill1.GetAttackDistance() < minAttackRadius {
				minAttackRadius = bt.skill1.GetAttackDistance()
			}
			break
		case skilltypes.SkillSecondTypeAttackProbability:
			bt.attackProbabilitySkills = append(bt.attackProbabilitySkills, bt.skill1.Id)
			break
		case skilltypes.SkillSecondTypeAttackedProbability:
			bt.attackedProbabilitySkills = append(bt.attackedProbabilitySkills, bt.skill1.Id)
			break
		}
	}

	if bt.skill2 != nil {
		bt.skills = append(bt.skills, int32(bt.skill2.Id))
		bt.allSkills = append(bt.allSkills, int32(bt.skill2.Id))

		switch bt.skill2.GetSkillSecondType() {
		case skilltypes.SkillSecondTypePositive:
			bt.positiveSkills = append(bt.positiveSkills, bt.skill2.Id)
			bt.skillRateMap[int32(bt.skill2.Id)] = int32(bt.SkillRate2)

			if bt.skill2.GetAttackDistance() < minAttackRadius {
				minAttackRadius = float64(bt.skill2.GetAttackDistance())
			}
			break
		case skilltypes.SkillSecondTypeAttackProbability:
			bt.attackProbabilitySkills = append(bt.attackProbabilitySkills, bt.skill2.Id)
			break
		case skilltypes.SkillSecondTypeAttackedProbability:
			bt.attackedProbabilitySkills = append(bt.attackedProbabilitySkills, bt.skill2.Id)
			break
		}
	}

	bt.minAttackRadius = minAttackRadius
	bt.minAttackRadiusSquare = bt.minAttackRadius * bt.minAttackRadius

	//等级
	if err = validator.MinValidate(float64(bt.Level), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("level", err)
		return
	}

	//验证属性
	for typInt, val := range bt.battlePropertyMap {
		typ := propertytypes.BattlePropertyType(typInt)
		if typ == propertytypes.BattlePropertyTypeMaxHP || typ == propertytypes.BattlePropertyTypeMaxTP {
			if err = validator.MinValidate(float64(val), float64(0), false); err != nil {
				err = template.NewTemplateFieldError(typ.String(), err)
				return
			}
		} else {
			if err = validator.MinValidate(float64(val), float64(0), true); err != nil {
				err = template.NewTemplateFieldError(typ.String(), err)
				return
			}
		}
	}

	if err = validator.MinValidate(float64(bt.ForceHit), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("forceHit", err)
		return
	}
	if err = validator.MinValidate(float64(bt.ForceCritical), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("forceCritical", err)
		return
	}

	if err = validator.MinValidate(float64(bt.BaseDamge), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("base_damge", err)
		return
	}

	if err = validator.MinValidate(float64(bt.DropOwnerType), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("drop_owner_type", err)
		return
	}

	if err = validator.MinValidate(float64(bt.ExpBase), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("exp_base", err)
		return
	}

	if err = validator.MinValidate(float64(bt.ExpPoint), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("exp_point", err)
		return
	}

	if err = validator.MinValidate(float64(bt.Patrolradius), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("patrolradius", err)
		return
	}

	if err = validator.MinValidate(float64(bt.Alertradius), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("alertradius", err)
		return
	}

	if err = validator.MinValidate(float64(bt.Randradius), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("randradius", err)
		return
	}

	if err = validator.RangeValidate(float64(bt.SkillRate1), float64(0), true, float64(common.MAX_RATE), true); err != nil {
		err = template.NewTemplateFieldError("skillRate1", err)
		return
	}

	if err = validator.RangeValidate(float64(bt.SkillRate2), float64(0), true, float64(common.MAX_RATE), true); err != nil {
		err = template.NewTemplateFieldError("skillRate2", err)
		return
	}

	if !bt.dropType.Valid() {
		err = fmt.Errorf("[%d] invalid", bt.DropType)
		err = template.NewTemplateFieldError("dropType", err)
		return
	}

	if bt.biologyAutoRecoverType == types.BiologyAutoRecoverTypeYes {
		if err = validator.MinValidate(float64(bt.AutoRecoverTime), float64(0), false); err != nil {
			err = template.NewTemplateFieldError("AutoRecoverTime", err)
			return
		}
	}

	if bt.dropType == scenetypes.DropTypePercent {
		if err = validator.RangeValidate(float64(bt.DropFlag), float64(0), false, float64(common.MAX_RATE), true); err != nil {
			err = template.NewTemplateFieldError("dropFlag", err)
			return
		}
		//不能配置回血
		if bt.biologyAutoRecoverType == types.BiologyAutoRecoverTypeYes {
			err = fmt.Errorf("[%d] invalid", bt.AutoRecoverMaxhealth)
			err = template.NewTemplateFieldError("autoRecover", err)
			return
		}
	}

	if !bt.dropJudgeType.Valid() {
		err = fmt.Errorf("[%d] invalid", bt.DropOwnerType)
		err = template.NewTemplateFieldError("DropOwnerType", err)
		return
	}

	for _, buffId := range bt.buffIdList {
		buffTemplate := template.GetTemplateService().Get(int(buffId), (*BuffTemplate)(nil))
		if buffTemplate == nil {
			err = template.NewTemplateFieldError("buffIds", err)
			return
		}
	}

	if bt.CaiJiTime != 0 {
		if err = validator.MinValidate(float64(bt.CaiJiTime), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("CaiJiTime", err)
			return
		}
	}

	if bt.CaiJiRecoverTime != 0 {
		if err = validator.MinValidate(float64(bt.CaiJiRecoverTime), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("CaiJiRecoverTime", err)
			return
		}
		if err = validator.MinValidate(float64(bt.CaiJiLimitCount), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("CaiJiLimitCount", err)
			return
		}
	}

	return nil
}

func (bt *BiologyTemplate) FileName() string {
	return "tb_biology.json"
}

func init() {
	template.Register((*BiologyTemplate)(nil))
}
