package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"

	"fmt"
)

//爱情树培养配置
type MarryTreeTemplate struct {
	*MarryTreeTemplateVO
	useItemTemplate    *ItemTemplate //进阶物品
	battleAttrTemplate *AttrTemplate //阶别属性
}

func (mtt *MarryTreeTemplate) TemplateId() int {
	return mtt.Id
}

func (mtt *MarryTreeTemplate) GetUseItemTemplate() *ItemTemplate {
	return mtt.useItemTemplate
}

func (mtt *MarryTreeTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return mtt.battleAttrTemplate
}

func (mtt *MarryTreeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mtt.FileName(), mtt.TemplateId(), err)
			return
		}
	}()

	//阶别attr属性
	if mtt.Attr != 0 {
		to := template.GetTemplateService().Get(int(mtt.Attr), (*AttrTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mtt.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := to.(*AttrTemplate)
		mtt.battleAttrTemplate = attrTemplate
	}

	//验证 UseItem
	if mtt.UseItem != 0 {
		useItemTemplate := template.GetTemplateService().Get(int(mtt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", mtt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		mtt.useItemTemplate = useItemTemplate.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(mtt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mtt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	return nil
}

func (mtt *MarryTreeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mtt.FileName(), mtt.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if mtt.NextId != 0 {
		diff := mtt.NextId - int32(mtt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", mtt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(mtt.NextId), (*MarryTreeTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mtt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mtt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mtt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_silver
	err = validator.MinValidate(float64(mtt.UseSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mtt.UseSilver)
		err = template.NewTemplateFieldError("UseSilver", err)
		return
	}

	//验证use_gold
	err = validator.MinValidate(float64(mtt.UseGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mtt.UseGold)
		err = template.NewTemplateFieldError("UseGold", err)
		return
	}

	//验证use_bindgold
	err = validator.MinValidate(float64(mtt.UseBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mtt.UseBindGold)
		err = template.NewTemplateFieldError("UseBindGold", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mtt.TimesMin), float64(0), true, float64(mtt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mtt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mtt.TimesMax), float64(mtt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mtt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mtt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mtt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mtt.AddMin), float64(0), true, float64(mtt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mtt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mtt.AddMax), float64(mtt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mtt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	return nil
}
func (mtt *MarryTreeTemplate) PatchAfterCheck() {

}
func (mtt *MarryTreeTemplate) FileName() string {
	return "tb_marry_tree.json"
}

func init() {
	template.Register((*MarryTreeTemplate)(nil))
}
