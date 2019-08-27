package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/marry/types"

	"fmt"
)

//婚戒培养配置
type MarryRingTemplate struct {
	*MarryRingTemplateVO
	ringType           types.MarryRingType //婚戒类型
	useItemTemplate    *ItemTemplate       //进阶物品
	battleAttrTemplate *AttrTemplate       //阶别属性
}

func (mrt *MarryRingTemplate) TemplateId() int {
	return mrt.Id
}

func (mrt *MarryRingTemplate) GetRingType() types.MarryRingType {
	return mrt.ringType
}

func (mrt *MarryRingTemplate) GetUseItemTemplate() *ItemTemplate {
	return mrt.useItemTemplate
}

func (mrt *MarryRingTemplate) GetBattleAttrTemplate() *AttrTemplate {
	return mrt.battleAttrTemplate
}

func (mrt *MarryRingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mrt.FileName(), mrt.TemplateId(), err)
			return
		}
	}()

	mrt.ringType = types.MarryRingType(mrt.Type)
	if !mrt.ringType.Valid() {
		err = fmt.Errorf("[%d] invalid", mrt.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//阶别attr属性
	if mrt.Attr != 0 {
		to := template.GetTemplateService().Get(int(mrt.Attr), (*AttrTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mrt.Attr)
			return template.NewTemplateFieldError("Attr", err)
		}
		attrTemplate, _ := to.(*AttrTemplate)
		mrt.battleAttrTemplate = attrTemplate
	}

	//验证 UseItem
	if mrt.UseItem != 0 {
		useItemTemplate := template.GetTemplateService().Get(int(mrt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", mrt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		mrt.useItemTemplate = useItemTemplate.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(mrt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mrt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	return nil
}

func (mrt *MarryRingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mrt.FileName(), mrt.TemplateId(), err)
			return
		}
	}()

	err = validator.MinValidate(float64(mrt.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 next_id
	if mrt.NextId != 0 {
		diff := mrt.NextId - int32(mrt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", mrt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(mrt.NextId), (*MarryRingTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mrt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		ringTemplate := to.(*MarryRingTemplate)
		diffLevel := ringTemplate.Level - mrt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", mrt.Level)
			err = template.NewTemplateFieldError("Level", err)
			return
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mrt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_silver
	err = validator.MinValidate(float64(mrt.UseSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.UseSilver)
		err = template.NewTemplateFieldError("UseSilver", err)
		return
	}

	//验证use_gold
	err = validator.MinValidate(float64(mrt.UseGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.UseGold)
		err = template.NewTemplateFieldError("UseGold", err)
		return
	}

	//验证use_bindgold
	err = validator.MinValidate(float64(mrt.UseBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.UseBindGold)
		err = template.NewTemplateFieldError("UseBindGold", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mrt.TimesMin), float64(0), true, float64(mrt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mrt.TimesMax), float64(mrt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mrt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mrt.AddMin), float64(0), true, float64(mrt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mrt.AddMax), float64(mrt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mrt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	return nil
}
func (mrt *MarryRingTemplate) PatchAfterCheck() {

}
func (mrt *MarryRingTemplate) FileName() string {
	return "tb_marry_ring.json"
}

func init() {
	template.Register((*MarryRingTemplate)(nil))
}
