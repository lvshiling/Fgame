package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"

	"fmt"
)

//神盾尖刺配置
type ShieldTemplate struct {
	*ShieldTemplateVO
	useItemTemplate    *ItemTemplate //进阶物品
	battleAttrTemplate *AttrTemplate //阶别属性
}

func (st *ShieldTemplate) TemplateId() int {
	return st.Id
}

func (st *ShieldTemplate) GetUseItemTemplate() *ItemTemplate {
	return st.useItemTemplate
}

func (st *ShieldTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return st.battleAttrTemplate
}

func (st *ShieldTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()

	//阶别attr属性
	if st.Attr != 0 {
		to := template.GetTemplateService().Get(int(st.Attr), (*AttrTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", st.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := to.(*AttrTemplate)
		st.battleAttrTemplate = attrTemplate
	}

	//验证 UseItem
	if st.UseItem != 0 {
		useItemTemplate := template.GetTemplateService().Get(int(st.UseItem), (*ItemTemplate)(nil))
		if useItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", st.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		st.useItemTemplate = useItemTemplate.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(st.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", st.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	return nil
}

func (st *ShieldTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()

	//验证Number
	err = validator.MinValidate(float64(st.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
	}

	//验证 next_id
	if st.NextId != 0 {
		diff := st.NextId - int32(st.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", st.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(st.NextId), (*FeatherTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", st.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(st.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_silver
	err = validator.MinValidate(float64(st.UseSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.UseSilver)
		err = template.NewTemplateFieldError("UseSilver", err)
		return
	}

	//验证use_gold
	err = validator.MinValidate(float64(st.UseGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.UseGold)
		err = template.NewTemplateFieldError("UseGold", err)
		return
	}

	//验证use_bindgold
	err = validator.MinValidate(float64(st.UseBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.UseBindGold)
		err = template.NewTemplateFieldError("UseBindGold", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(st.TimesMin), float64(0), true, float64(st.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(st.TimesMax), float64(st.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(st.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(st.AddMin), float64(0), true, float64(st.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(st.AddMax), float64(st.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	return nil
}
func (st *ShieldTemplate) PatchAfterCheck() {

}
func (st *ShieldTemplate) FileName() string {
	return "tb_shield.json"
}

func init() {
	template.Register((*ShieldTemplate)(nil))
}
