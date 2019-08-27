package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	inventorytypes "fgame/fgame/game/inventory/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//屠龙装备配置
type TuLongEquipTemplate struct {
	*TuLongEquipTemplateVO
	//套装
	tempTaozhuangTemplate *TuLongEquipSuitGroupTemplate
	//属性
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	//部位
	posType inventorytypes.BodyPositionType
}

func (t *TuLongEquipTemplate) TemplateId() int {
	return t.Id
}

func (t *TuLongEquipTemplate) GetPosType() inventorytypes.BodyPositionType {
	return t.posType
}

func (t *TuLongEquipTemplate) GetTaozhuangTemplate() *TuLongEquipSuitGroupTemplate {
	return t.tempTaozhuangTemplate
}

func (t *TuLongEquipTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

//获取装备技能
func (m *TuLongEquipTemplate) GetTuLongEquipGroupSuitSkill(equipNum int32) (skillList []int32) {
	groupSuitTemplate := m.tempTaozhuangTemplate
	if groupSuitTemplate == nil {
		return
	}

	return groupSuitTemplate.GetSuitEffectSkillId(equipNum)
}

func (t *TuLongEquipTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//套装
	if t.SuitGroup != 0 {
		tempTaozhuangTemplate := template.GetTemplateService().Get(int(t.SuitGroup), (*TuLongEquipSuitGroupTemplate)(nil))
		if tempTaozhuangTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.SuitGroup)
			err = template.NewTemplateFieldError("SuitGroup", err)
			return
		}
		t.tempTaozhuangTemplate = tempTaozhuangTemplate.(*TuLongEquipSuitGroupTemplate)
	}

	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	return nil
}

func (t *TuLongEquipTemplate) PatchAfterCheck() {}

func (t *TuLongEquipTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//装备阶数
	err = validator.MinValidate(float64(t.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Number)
		return template.NewTemplateFieldError("Number", err)
	}
	//生命
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}
	//攻击
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}
	//防御
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//部位
	t.posType = inventorytypes.BodyPositionType(t.PosType)
	if !t.posType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.PosType)
		return template.NewTemplateFieldError("PosType", err)
	}

	return nil
}

func (edt *TuLongEquipTemplate) FileName() string {
	return "tb_tulongequip.json"
}

func init() {
	template.Register((*TuLongEquipTemplate)(nil))
}
