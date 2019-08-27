package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fmt"
)

//屠龙装备技能配置
type TuLongEquipSkillTemplate struct {
	*TuLongEquipSkillTemplateVO
	nextTemplate *TuLongEquipSkillTemplate
	suitType     tulongequiptypes.TuLongSuitType //屠龙装备类型
	useItemMap   map[int32]int32
}

func (t *TuLongEquipSkillTemplate) TemplateId() int {
	return t.Id
}

func (t *TuLongEquipSkillTemplate) GetNextTemplate() *TuLongEquipSkillTemplate {
	return t.nextTemplate
}

func (t *TuLongEquipSkillTemplate) GetUseItemMap() map[int32]int32 {
	return t.useItemMap
}

func (t *TuLongEquipSkillTemplate) GetSuitType() tulongequiptypes.TuLongSuitType {
	return t.suitType
}

func (t *TuLongEquipSkillTemplate) PatchAfterCheck() {
}

func (t *TuLongEquipSkillTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	t.useItemMap = make(map[int32]int32)
	useItemIdArr, err := utils.SplitAsIntArray(t.UplevelItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UplevelItem)
		return template.NewTemplateFieldError("UplevelItem", err)
	}
	useItemCountArr, err := utils.SplitAsIntArray(t.UplevelItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UplevelItemCount)
		return template.NewTemplateFieldError("UplevelItemCount", err)
	}
	if len(useItemIdArr) != 0 && len(useItemIdArr) != len(useItemCountArr) {
		err = fmt.Errorf("[%s] invalid", t.UplevelItem)
		return template.NewTemplateFieldError("UplevelItem and UplevelItemCount", err)
	}
	for index, itemId := range useItemIdArr {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", t.UplevelItem)
			return template.NewTemplateFieldError("UplevelItem", err)
		}

		err = validator.MinValidate(float64(useItemCountArr[index]), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.UplevelItemCount)
			return template.NewTemplateFieldError("UplevelItemCount", err)
		}
		t.useItemMap[itemId] += useItemCountArr[index]
	}

	return nil
}

func (t *TuLongEquipSkillTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 type
	t.suitType = tulongequiptypes.TuLongSuitType(t.Type)
	if !t.suitType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 next_id
	if t.NextId != 0 {
		if t.NextId-int32(t.Id) != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		to := template.GetTemplateService().Get(int(t.NextId), (*TuLongEquipSkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		t.nextTemplate = to.(*TuLongEquipSkillTemplate)
	}

	// 条件1 穿戴阶数
	err = validator.MinValidate(float64(t.NeedJieShu), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedJieShu)
		return template.NewTemplateFieldError("Value1", err)
	}
	// 条件2 穿戴数量
	err = validator.MinValidate(float64(t.NeedEquipNum), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedEquipNum)
		return template.NewTemplateFieldError("Value2", err)
	}
	// 条件3 总强化等级
	err = validator.MinValidate(float64(t.NeedStrengthenLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedStrengthenLevel)
		return template.NewTemplateFieldError("Value3", err)
	}

	//验证 UplevelRate
	err = validator.MinValidate(float64(t.UplevelRate), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UplevelRate)
		return template.NewTemplateFieldError("UplevelRate", err)
	}

	//验证 skill_id
	to := template.GetTemplateService().Get(int(t.SkillId), (*SkillTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}

	return nil
}

func (t *TuLongEquipSkillTemplate) FileName() string {
	return "tb_tulongequip_skill.json"
}

func init() {
	template.Register((*TuLongEquipSkillTemplate)(nil))
}
