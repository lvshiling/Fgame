package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

//血盾配置
type BloodShieldTemplate struct {
	*BloodShieldTemplateVO
	useItemMap         map[int32]int32 //进阶物品
	unrealItemTemplate *ItemTemplate
}

func (t *BloodShieldTemplate) TemplateId() int {
	return t.Id
}

func (t *BloodShieldTemplate) GetUseItemTemplate() map[int32]int32 {
	return t.useItemMap
}

func (t *BloodShieldTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.useItemMap = make(map[int32]int32)
	//验证 UseItem
	if t.UseItem != "" {
		itemArr, err := utils.SplitAsIntArray(t.UseItem)
		if err != nil {
			return err
		}

		numArr, err := utils.SplitAsIntArray(t.ItemCount)
		if err != nil {
			return err
		}

		if len(itemArr) != len(numArr) {
			err = fmt.Errorf("[%s][%s] invalid", t.UseItem, t.ItemCount)
			err = template.NewTemplateFieldError("UseItem or ItemCount", err)
			return err
		}

		for index, _ := range itemArr {
			t.useItemMap[itemArr[index]] = numArr[index]
		}
	}

	return nil
}

func (t *BloodShieldTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(t.Star), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Star)
		err = template.NewTemplateFieldError("Star", err)
		return
	}

	//验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*BloodShieldTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		bloodShieldTemplate := to.(*BloodShieldTemplate)

		diffLevel := bloodShieldTemplate.Star - t.Star
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", t.Star)
			return template.NewTemplateFieldError("Star", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(t.UpdatePercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdatePercent)
		err = template.NewTemplateFieldError("UpdatePercent", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(t.TimesMin), float64(0), true, float64(t.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(t.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(t.AddMin), float64(0), true, float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(t.AddMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 medicinal_limit
	err = validator.MinValidate(float64(t.MedicinalLimit), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MedicinalLimit)
		err = template.NewTemplateFieldError("MedicinalLimit", err)
		return
	}

	//验证 spell_id
	to := template.GetTemplateService().Get(int(t.SpellId), (*SkillTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.SpellId)
		err = template.NewTemplateFieldError("SpellId", err)
		return
	}
	skillTemplate := to.(*SkillTemplate)
	skillFirstType := skillTemplate.GetSkillFirstType()
	if skillFirstType != skilltypes.SkillFirstTypeXueDun {
		err = fmt.Errorf("[%d] invalid", t.SpellId)
		err = template.NewTemplateFieldError("SpellId", err)
		return
	}

	//验证物品
	for itemId, num := range t.useItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}

		itemTemplate := to.(*ItemTemplate)
		if itemTemplate.GetItemSubType() != itemtypes.ItemXueDunSubTypeUpstar {
			err = fmt.Errorf("[%s] invalid", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}

		//验证 ItemCount
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	return nil
}
func (t *BloodShieldTemplate) PatchAfterCheck() {

}
func (t *BloodShieldTemplate) FileName() string {
	return "tb_blood_shield.json"
}

func init() {
	template.Register((*BloodShieldTemplate)(nil))
}
