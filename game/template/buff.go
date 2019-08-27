package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

type BuffTemplate struct {
	*BuffTemplateVO
	//buff类型
	buffType scenetypes.BuffType
	//触发类型
	touchType scenetypes.BuffTouchType
	//子状态
	// subBuffTemplate *BuffTemplate
	subBuffMap map[int32]int32
	//属性加成
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	//属性万分比加成
	battlePropertyPercentMap map[propertytypes.BattlePropertyType]int64
	//离线保存
	offlineSaveType scenetypes.BuffOfflineSaveType
	//buff飘字
	buffPiaoZi scenetypes.BuffPiaoZi
	//buff抗性
	buffKangXing scenetypes.BuffKangXing
	//特殊效果
	buffEffectType scenetypes.BuffEffectType
	//子技能
	skillTemplate *SkillTemplate
	//前置buff
	parentBuffList []int32
}

func (t *BuffTemplate) TemplateId() int {
	return t.Id
}

func (t *BuffTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.buffType = scenetypes.BuffType(t.BuffType)
	t.touchType = scenetypes.BuffTouchType(t.TypeTouch)
	subIdArr, err := utils.SplitAsIntArray(t.SubId)
	if err != nil {
		return template.NewTemplateFieldError("subId", err)
	}
	subRateArr, err := utils.SplitAsIntArray(t.SubRate)
	if err != nil {
		return template.NewTemplateFieldError("SubRate", err)
	}
	if len(subIdArr) != len(subRateArr) {
		err = fmt.Errorf("subId[%s]和subRate[%s]数量不一致", t.SubId, t.SubRate)
		return template.NewTemplateFieldError("subId,subRate", err)
	}
	t.subBuffMap = make(map[int32]int32)
	for i, subId := range subIdArr {
		t.subBuffMap[subId] = subRateArr[i]
	}

	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.LifemaxAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxTP] = int64(t.TpmaxAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.AttackAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.DefenseAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeCrit] = int64(t.CriticalAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeTough] = int64(t.ToughAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeBlock] = int64(t.BlockAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAbnormality] = int64(t.AbnormalityAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeHit] = int64(t.HitAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDodge] = int64(t.DodgeAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeHuanYunAttack] = int64(t.HunyuanAttAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeHuanYunDef] = int64(t.HunyuanDefAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDamageAdd] = int64(t.HarmBase)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDamageDefend] = int64(t.CuthurtBase)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDamageAddPercent] = int64(t.HarmPercent)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDamageDefendPercent] = int64(t.CuthurtPercent)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeCritRatePercent] = int64(t.CritRatePercent)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeCritHarmPercent] = int64(t.CritHarmPercent)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeHitRatePercent] = int64(t.HitRatePercent)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDodgeRatePercent] = int64(t.DodgeRatePercent)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeSpellCdPercent] = int64(t.SpellCdPercent)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeFanTan] = int64(t.FantanAdd)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeFanTanPercent] = int64(t.FantanPercent)

	t.battlePropertyMap[propertytypes.BattlePropertyTypeBlockRatePercent] = int64(t.BlockRatePercent)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeSpellCdPercent] = int64(t.SpellCdPercent)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAddExp] = int64(t.AddExp)

	t.battlePropertyPercentMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.LifemaxPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeMaxTP] = int64(t.TpmaxPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeAttack] = int64(t.AttackPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeDefend] = int64(t.DefensePercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeCrit] = int64(t.CriticalPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeTough] = int64(t.ToughPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeBlock] = int64(t.BlockPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeAbnormality] = int64(t.AbnormalityPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeHit] = int64(t.HitPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeDodge] = int64(t.DodgePercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeHuanYunAttack] = int64(t.HunyuanAttPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeHuanYunDef] = int64(t.HunyuanDefPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeDamageAdd] = int64(t.HarmPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeDamageDefend] = int64(t.CuthurtPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeMoveSpeed] = int64(t.SpeedMovePercent)

	t.offlineSaveType = scenetypes.BuffOfflineSaveType(t.OfflineSaveType)
	t.buffPiaoZi = scenetypes.BuffPiaoZi(t.BuffPiaozi)
	t.buffKangXing = scenetypes.BuffKangXing(t.RelativeFastness)
	t.buffEffectType = scenetypes.BuffEffectType(t.EffectType)
	tempSkillTemplate := template.GetTemplateService().Get(int(t.SkillId), (*SkillTemplate)(nil))
	if tempSkillTemplate != nil {
		t.skillTemplate = tempSkillTemplate.(*SkillTemplate)
	}
	parentBuffArr, err := utils.SplitAsIntArray(t.ParentBuffId)
	if err != nil {
		return template.NewTemplateFieldError("ParentBuffId", err)
	}
	t.parentBuffList = parentBuffArr
	return
}

func (t *BuffTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	if !t.buffType.Valid() {
		err := fmt.Errorf("[%d] invalid", t.BuffType)
		return template.NewTemplateFieldError("buffType", err)
	}

	if !t.touchType.Valid() {
		err := fmt.Errorf("[%d] invalid", t.TypeTouch)
		return template.NewTemplateFieldError("typeTouch", err)
	}

	if !t.offlineSaveType.Valid() {
		err := fmt.Errorf("[%d] invalid", t.OfflineSaveType)
		return template.NewTemplateFieldError("OfflineSaveType", err)
	}

	err = validator.MinValidate(float64(t.GetExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GetExp)
		return template.NewTemplateFieldError("GetExp", err)
	}

	err = validator.MinValidate(float64(t.GetExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GetExpPoint)
		return template.NewTemplateFieldError("GetExpPoint", err)
	}

	err = validator.MinValidate(float64(t.EffectTypeBase), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.EffectTypeBase)
		return template.NewTemplateFieldError("EffectTypeBase", err)
	}

	err = validator.MinValidate(float64(t.EffectTypePercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.EffectTypePercent)
		return template.NewTemplateFieldError("EffectTypePercent", err)
	}

	for subBuffId, _ := range t.subBuffMap {
		tempBuffTemp := template.GetTemplateService().Get(int(subBuffId), (*BuffTemplate)(nil))
		if tempBuffTemp == nil {
			err = fmt.Errorf("[%d] invalid", t.SubId)
			return template.NewTemplateFieldError("SubId", err)
		}
	}

	if t.skillTemplate != nil {
		if t.skillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeSubSkill {
			err = fmt.Errorf("[%d] invalid", t.SkillId)
			return template.NewTemplateFieldError("SkillId", err)
		}
	}

	if t.buffType == scenetypes.BuffTypeTitleDingZhi && t.TypeRemove != 0 {
		err = fmt.Errorf("[%d] invalid", t.TypeRemove)
		return template.NewTemplateFieldError("TypeRemove", err)
	}

	return nil
}
func (t *BuffTemplate) PatchAfterCheck() {

}
func (t *BuffTemplate) FileName() string {
	return "tb_buff_service.json"
}

func (t *BuffTemplate) GetBuffType() scenetypes.BuffType {
	return t.buffType
}
func (t *BuffTemplate) GetTouchType() scenetypes.BuffTouchType {
	return t.touchType
}

func (t *BuffTemplate) GetSubBuffMap() map[int32]int32 {
	return t.subBuffMap
}

func (t *BuffTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *BuffTemplate) GetBattlePropertyPercentMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyPercentMap
}

func (t *BuffTemplate) GetOfflineSaveType() scenetypes.BuffOfflineSaveType {
	return t.offlineSaveType
}

func (t *BuffTemplate) IsSave() bool {
	return t.offlineSaveType != scenetypes.BuffOfflineSaveTypeNone
}

func (t *BuffTemplate) IsTimer() bool {
	return t.touchType == scenetypes.BuffTouchTypeTimer
}

func (t *BuffTemplate) GetBuffPiaoZi() scenetypes.BuffPiaoZi {
	return t.buffPiaoZi
}

func (t *BuffTemplate) GetBuffKangXing() scenetypes.BuffKangXing {
	return t.buffKangXing
}

func (t *BuffTemplate) GetBuffEffectType() scenetypes.BuffEffectType {
	return t.buffEffectType
}

func (t *BuffTemplate) GetSkillTemplate() *SkillTemplate {
	return t.skillTemplate
}

func (t *BuffTemplate) GetParentBuffList() []int32 {
	return t.parentBuffList
}
func init() {
	template.Register((*BuffTemplate)(nil))
}
