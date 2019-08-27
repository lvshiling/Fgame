package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

//法宝通灵配置
type FaBaoTongLingTemplate struct {
	*FaBaoTongLingTemplateVO
	needItemMap               map[int32]int32 //通灵需要物品
	nextFaBaoTongLingTemplate *FaBaoTongLingTemplate
	useItemTemplate           *ItemTemplate
}

func (fut *FaBaoTongLingTemplate) TemplateId() int {
	return fut.Id
}

func (fut *FaBaoTongLingTemplate) GetNeedItemMap() map[int32]int32 {

	return fut.needItemMap
}

func (fut *FaBaoTongLingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fut.FileName(), fut.TemplateId(), err)
			return
		}
	}()

	fut.needItemMap = make(map[int32]int32)
	//验证 use_item
	if fut.UseItem != 0 {
		to := template.GetTemplateService().Get(int(fut.UseItem), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}

		err = validator.MinValidate(float64(fut.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", fut.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}

		fut.needItemMap[fut.UseItem] = fut.ItemCount
	}

	//验证 next_id
	if fut.NextId != 0 {
		to := template.GetTemplateService().Get(int(fut.NextId), (*FaBaoTongLingTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*FaBaoTongLingTemplate)
			diffLevel := nextTemplate.Level - fut.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			fut.nextFaBaoTongLingTemplate = nextTemplate
		}
	}

	return nil
}

func (fut *FaBaoTongLingTemplate) PatchAfterCheck() {

}

func (fut *FaBaoTongLingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(fut.FileName(), fut.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(fut.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 fabao_percent
	err = validator.RangeValidate(float64(fut.FaBaoPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.FaBaoPercent)
		err = template.NewTemplateFieldError("FaBaoPercent", err)
		return
	}

	//验证 level
	err = validator.MinValidate(float64(fut.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	for itemId, _ := range fut.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
		itemTemplate := to.(*ItemTemplate)

		if itemTemplate.GetItemType() != itemtypes.ItemTypeFaBao {
			err = fmt.Errorf("UpstarItemId [%d]  invalid", fut.UseItem)
			return template.NewTemplateFieldError("UpstarItemId", err)
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
	err = validator.MinValidate(float64(fut.ZhufuMax), float64(fut.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	return nil
}

func (fut *FaBaoTongLingTemplate) FileName() string {
	return "tb_fabao_tongling.json"
}

func init() {
	template.Register((*FaBaoTongLingTemplate)(nil))
}
