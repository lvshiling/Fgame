package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

//领域升星配置
type FieldUpstarTemplate struct {
	*FieldUpstarTemplateVO
	needItemMap             map[int32]int32 //升星需要物品
	nextFieldUpstarTemplate *FieldUpstarTemplate
	useItemTemplate         *ItemTemplate
}

func (fut *FieldUpstarTemplate) TemplateId() int {
	return fut.Id
}

func (fut *FieldUpstarTemplate) GetNeedItemMap() map[int32]int32 {
	return fut.needItemMap
}

func (fut *FieldUpstarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fut.FileName(), fut.TemplateId(), err)
			return
		}
	}()

	fut.needItemMap = make(map[int32]int32)
	//验证 upstar_item_id
	if fut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(fut.UpstarItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}

		err = validator.MinValidate(float64(fut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}
		fut.needItemMap[fut.UpstarItemId] = fut.UpstarItemCount

		fut.useItemTemplate = to.(*ItemTemplate)
	}

	//验证 next_id
	if fut.NextId != 0 {
		to := template.GetTemplateService().Get(int(fut.NextId), (*FieldUpstarTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*FieldUpstarTemplate)
			diffLevel := nextTemplate.Level - fut.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			fut.nextFieldUpstarTemplate = nextTemplate
		}
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(fut.TimesMin), float64(0), true, float64(fut.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(fut.TimesMax), float64(fut.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(fut.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(fut.AddMin), float64(0), true, float64(fut.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(fut.AddMax), float64(fut.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 ZhufuMax
	err = validator.MinValidate(float64(fut.ZhufuMax), float64(fut.ZhufuMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	return nil
}

func (fut *FieldUpstarTemplate) PatchAfterCheck() {

}

func (fut *FieldUpstarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fut.FileName(), fut.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(fut.UpstarRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.UpstarRate)
		err = template.NewTemplateFieldError("UpstarRate", err)
		return
	}

	//验证 field_percent
	err = validator.RangeValidate(float64(fut.FieldPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.FieldPercent)
		err = template.NewTemplateFieldError("FieldPercent", err)
		return
	}

	//验证 level
	err = validator.MinValidate(float64(fut.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 hp
	err = validator.MinValidate(float64(fut.Hp), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(fut.Attack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(fut.Defence), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	if fut.useItemTemplate != nil {
		if fut.useItemTemplate.GetItemType() != itemtypes.ItemTypeLingyu {
			err = fmt.Errorf("UpstarItemId [%d]  invalid", fut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}
	}

	return nil
}

func (fut *FieldUpstarTemplate) FileName() string {
	return "tb_field_upstar.json"
}

func init() {
	template.Register((*FieldUpstarTemplate)(nil))
}
