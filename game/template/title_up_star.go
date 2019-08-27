package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

// 称号升星配置
type TitleUpStarTemplate struct {
	*TitleUpStarTemplateVO
	needItemMap             map[int32]int32 //升星需要物品
	nextTitleUpStarTemplate *TitleUpStarTemplate
	useItemTemplate         *ItemTemplate
}

func (mut *TitleUpStarTemplate) TemplateId() int {
	return mut.Id
}

func (mut *TitleUpStarTemplate) GetNeedItemMap() map[int32]int32 {
	return mut.needItemMap
}

func (mut *TitleUpStarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mut.FileName(), mut.TemplateId(), err)
			return
		}
	}()

	mut.needItemMap = make(map[int32]int32)
	//验证 upstar_item_id
	if mut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(mut.UpstarItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}

		err = validator.MinValidate(float64(mut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}
		mut.needItemMap[mut.UpstarItemId] = mut.UpstarItemCount

		mut.useItemTemplate = to.(*ItemTemplate)
	}

	//验证 next_id
	if mut.NextId != 0 {
		to := template.GetTemplateService().Get(int(mut.NextId), (*TitleUpStarTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*TitleUpStarTemplate)
			diffLevel := nextTemplate.Level - mut.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			mut.nextTitleUpStarTemplate = nextTemplate
		}
	}

	return nil
}

func (mut *TitleUpStarTemplate) PatchAfterCheck() {

}

func (mut *TitleUpStarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mut.FileName(), mut.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(mut.UpstarRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.UpstarRate)
		err = template.NewTemplateFieldError("UpstarRate", err)
		return
	}

	/*
		//验证 title_percent
		err = validator.RangeValidate(float64(mut.TitlePercent), float64(0), true, float64(common.MAX_RATE), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mut.TitlePercent)
			err = template.NewTemplateFieldError("TitlePercent", err)
			return
		}
	*/

	//验证 level
	err = validator.MinValidate(float64(mut.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 hp
	err = validator.MinValidate(float64(mut.Hp), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(mut.Attack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//验证 defence
	err = validator.MinValidate(float64(mut.Defence), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	if mut.useItemTemplate != nil {
		if mut.useItemTemplate.GetItemType() != itemtypes.ItemTypeTitle {
			err = fmt.Errorf("UpstarItemId [%d]  invalid", mut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mut.TimesMin), float64(0), true, float64(mut.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mut.TimesMax), float64(mut.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mut.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mut.AddMin), float64(0), true, float64(mut.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mut.AddMax), float64(mut.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(mut.ZhufuMax), float64(mut.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mut.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	return nil
}

func (mut *TitleUpStarTemplate) FileName() string {
	return "tb_title_upstar.json"
}

func init() {
	template.Register((*TitleUpStarTemplate)(nil))
}
