package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"

	"fmt"
)

//护体盾金甲丹配置
type BodyShieldJinJiaTemplate struct {
	*BodyShieldJinJiaTemplateVO
	useItemMap         map[int32]int32 //进阶物品
	jinJiaItemTemplate *ItemTemplate
}

func (bsjjt *BodyShieldJinJiaTemplate) TemplateId() int {
	return bsjjt.Id
}

func (bsjjt *BodyShieldJinJiaTemplate) GetUseItemTemplate() map[int32]int32 {
	return bsjjt.useItemMap
}

func (bsjjt *BodyShieldJinJiaTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(bsjjt.FileName(), bsjjt.TemplateId(), err)
			return
		}
	}()

	bsjjt.useItemMap = make(map[int32]int32)
	//验证 UseItem
	if bsjjt.UseItem != 0 {
		useItemTemplate := template.GetTemplateService().Get(int(bsjjt.UseItem), (*ItemTemplate)(nil))
		if useItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", bsjjt.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		bsjjt.jinJiaItemTemplate = useItemTemplate.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(bsjjt.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", bsjjt.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
		bsjjt.useItemMap[bsjjt.UseItem] = bsjjt.ItemCount
	}

	return nil
}

func (bsjjt *BodyShieldJinJiaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(bsjjt.FileName(), bsjjt.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(bsjjt.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 next_id
	if bsjjt.NextId != 0 {
		diff := bsjjt.NextId - int32(bsjjt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", bsjjt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(bsjjt.NextId), (*BodyShieldJinJiaTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", bsjjt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		huanHuaTemplate := to.(*BodyShieldJinJiaTemplate)

		diffLevel := huanHuaTemplate.Level - bsjjt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", bsjjt.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(bsjjt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(bsjjt.TimesMin), float64(0), true, float64(bsjjt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(bsjjt.TimesMax), float64(bsjjt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(bsjjt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(bsjjt.AddMin), float64(0), true, float64(bsjjt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(bsjjt.AddMax), float64(bsjjt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(bsjjt.ZhufuMax), float64(bsjjt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 hp
	err = validator.MinValidate(float64(bsjjt.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.Hp)
		err = template.NewTemplateFieldError("Hp", err)
		return
	}

	//验证 attack
	err = validator.MinValidate(float64(bsjjt.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.Attack)
		err = template.NewTemplateFieldError("Attack", err)
		return
	}

	//验证 defence
	err = validator.MinValidate(float64(bsjjt.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", bsjjt.Defence)
		err = template.NewTemplateFieldError("Defence", err)
		return
	}

	if bsjjt.jinJiaItemTemplate != nil {
		if bsjjt.jinJiaItemTemplate.GetItemSubType() != itemtypes.ItemBodyShieldSubTypeJJDan {
			err = fmt.Errorf("[%d] invalid", bsjjt.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
	}

	return nil
}
func (bsjjt *BodyShieldJinJiaTemplate) PatchAfterCheck() {

}
func (bsjjt *BodyShieldJinJiaTemplate) FileName() string {
	return "tb_body_shield_jinjia.json"
}

func init() {
	template.Register((*BodyShieldJinJiaTemplate)(nil))
}
