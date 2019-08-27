package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

//护体盾配置
type BodyShieldTemplate struct {
	*BodyShieldTemplateVO
	useItemTemplate    *ItemTemplate //进阶物品
	battleAttrTemplate *AttrTemplate //阶别属性
}

func (bst *BodyShieldTemplate) TemplateId() int {
	return bst.Id
}

func (bst *BodyShieldTemplate) GetUseItemTemplate() *ItemTemplate {
	return bst.useItemTemplate
}

func (bst *BodyShieldTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return bst.battleAttrTemplate
}

func (bst *BodyShieldTemplate) GetIsClear() bool {
	return bst.IsClear != 0
}

func (bst *BodyShieldTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(bst.FileName(), bst.TemplateId(), err)
			return
		}
	}()

	//验证 Attr
	if bst.Attr != 0 {
		tempAttrTemplate := template.GetTemplateService().Get(int(bst.Attr), (*AttrTemplate)(nil))
		if tempAttrTemplate == nil {
			err = fmt.Errorf("[%d] invalid", bst.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := tempAttrTemplate.(*AttrTemplate)
		bst.battleAttrTemplate = attrTemplate
	}

	//验证 UseItem
	if bst.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(bst.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", bst.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		bst.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(bst.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", bst.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	return nil
}

func (bst *BodyShieldTemplate) PatchAfterCheck() {

}
func (bst *BodyShieldTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(bst.FileName(), bst.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if bst.NextId != 0 {
		to := template.GetTemplateService().Get(int(bst.NextId), (*BodyShieldTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", bst.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp := to.(*BodyShieldTemplate)

		diff := nextTemp.Number - int32(bst.Number)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", bst.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(bst.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_money
	err = validator.MinValidate(float64(bst.UseMoney), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.UseMoney)
		err = template.NewTemplateFieldError("UseMoney", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(bst.TimesMin), float64(0), true, float64(bst.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(bst.TimesMax), float64(bst.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(bst.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(bst.AddMin), float64(0), true, float64(bst.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(bst.AddMax), float64(bst.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(bst.ZhufuMax), float64(bst.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 UseYinliang
	err = validator.MinValidate(float64(bst.UseYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.UseYinliang)
		err = template.NewTemplateFieldError("UseYinliang", err)
		return
	}

	//验证 ShidanLimit
	err = validator.MinValidate(float64(bst.ShidanLimit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.ShidanLimit)
		err = template.NewTemplateFieldError("ShidanLimit", err)
		return
	}

	err = validator.MinValidate(float64(bst.IsClear), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bst.IsClear)
		return template.NewTemplateFieldError("IsClear", err)
	}

	return nil
}

func (bst *BodyShieldTemplate) FileName() string {
	return "tb_body_shield.json"
}

func init() {
	template.Register((*BodyShieldTemplate)(nil))
}
