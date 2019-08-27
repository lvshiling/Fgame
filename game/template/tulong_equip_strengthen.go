package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	inventorytypes "fgame/fgame/game/inventory/types"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fmt"
)

//屠龙装备强化配置
type TuLongEquipStrengthenTemplate struct {
	*TuLongEquipStrengthenTemplateVO
	suitType                          tulongequiptypes.TuLongSuitType
	posType                           inventorytypes.BodyPositionType
	nextTuLongEquipStrengthenTemplate *TuLongEquipStrengthenTemplate //下一级强化
}

func (t *TuLongEquipStrengthenTemplate) TemplateId() int {
	return t.Id
}

func (t *TuLongEquipStrengthenTemplate) GetSuitType() tulongequiptypes.TuLongSuitType {
	return t.suitType
}

func (t *TuLongEquipStrengthenTemplate) GetPosType() inventorytypes.BodyPositionType {
	return t.posType
}

func (t *TuLongEquipStrengthenTemplate) GetNextTemplate() *TuLongEquipStrengthenTemplate {
	return t.nextTuLongEquipStrengthenTemplate
}

func (t *TuLongEquipStrengthenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//下一阶强化
	if t.NextId != 0 {
		if t.NextId-int32(t.Id) != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}

		tempNextTuLongEquipStrengthenTemplate := template.GetTemplateService().Get(int(t.NextId), (*TuLongEquipStrengthenTemplate)(nil))
		if tempNextTuLongEquipStrengthenTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}
		t.nextTuLongEquipStrengthenTemplate = tempNextTuLongEquipStrengthenTemplate.(*TuLongEquipStrengthenTemplate)
	}

	return nil
}
func (t *TuLongEquipStrengthenTemplate) PatchAfterCheck() {

}
func (t *TuLongEquipStrengthenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	t.suitType = tulongequiptypes.TuLongSuitType(t.Type)
	if !t.suitType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	t.posType = inventorytypes.BodyPositionType(t.SubType)
	if !t.posType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SubType)
		return template.NewTemplateFieldError("SubType", err)
	}

	//验证 等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 概率
	err = validator.MinValidate(float64(t.Rate), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Rate)
		return template.NewTemplateFieldError("Rate", err)
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

	//
	to := template.GetTemplateService().Get(int(t.NeedItem), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%s] invalid", t.NeedItem)
		err = template.NewTemplateFieldError("NeedItem", err)
		return
	}
	err = validator.MinValidate(float64(t.ItemCount), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ItemCount)
		return template.NewTemplateFieldError("ItemCount", err)
	}

	return nil
}

func (edt *TuLongEquipStrengthenTemplate) FileName() string {
	return "tb_tulongequip_strengthen.json"
}

func init() {
	template.Register((*TuLongEquipStrengthenTemplate)(nil))
}
