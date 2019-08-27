package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"

	"fmt"
)

//护体仙羽配置
type FeatherTemplate struct {
	*FeatherTemplateVO
	useItemTemplate    *ItemTemplate //进阶物品
	battleAttrTemplate *AttrTemplate //阶别属性
}

func (ft *FeatherTemplate) TemplateId() int {
	return ft.Id
}

func (ft *FeatherTemplate) GetUseItemTemplate() *ItemTemplate {
	return ft.useItemTemplate
}

func (ft *FeatherTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return ft.battleAttrTemplate
}

func (ft *FeatherTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(ft.FileName(), ft.TemplateId(), err)
			return
		}
	}()

	//阶别attr属性
	if ft.Attr != 0 {
		to := template.GetTemplateService().Get(int(ft.Attr), (*AttrTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", ft.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := to.(*AttrTemplate)
		ft.battleAttrTemplate = attrTemplate
	}

	//验证 UseItem
	if ft.UseItem != 0 {
		useItemTemplate := template.GetTemplateService().Get(int(ft.UseItem), (*ItemTemplate)(nil))
		if useItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", ft.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		ft.useItemTemplate = useItemTemplate.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(ft.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", ft.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	return nil
}

func (ft *FeatherTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(ft.FileName(), ft.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if ft.NextId != 0 {
		diff := ft.NextId - int32(ft.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", ft.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(ft.NextId), (*FeatherTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", ft.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(ft.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_silver
	err = validator.MinValidate(float64(ft.UseSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.UseSilver)
		err = template.NewTemplateFieldError("UseSilver", err)
		return
	}

	//验证use_gold
	err = validator.MinValidate(float64(ft.UseGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.UseGold)
		err = template.NewTemplateFieldError("UseGold", err)
		return
	}

	//验证use_bindgold
	err = validator.MinValidate(float64(ft.UseBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.UseBindGold)
		err = template.NewTemplateFieldError("UseBindGold", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(ft.TimesMin), float64(0), true, float64(ft.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(ft.TimesMax), float64(ft.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(ft.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(ft.AddMin), float64(0), true, float64(ft.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(ft.AddMax), float64(ft.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 NeedRate
	err = validator.MinValidate(float64(ft.NeedRate), float64(ft.NeedRate), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", ft.NeedRate)
		err = template.NewTemplateFieldError("NeedRate", err)
		return
	}

	return nil
}
func (ft *FeatherTemplate) PatchAfterCheck() {

}
func (ft *FeatherTemplate) FileName() string {
	return "tb_feather.json"
}

func init() {
	template.Register((*FeatherTemplate)(nil))
}
