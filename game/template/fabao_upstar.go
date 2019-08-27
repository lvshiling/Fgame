package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

//法宝升星配置
type FaBaoUpstarTemplate struct {
	*FaBaoUpstarTemplateVO
	needItemMap             map[int32]int32 //升星需要物品
	nextFaBaoUpstarTemplate *FaBaoUpstarTemplate
	useItemTemplate         *ItemTemplate
}

func (fut *FaBaoUpstarTemplate) TemplateId() int {
	return fut.Id
}

func (fut *FaBaoUpstarTemplate) GetNeedItemMap() map[int32]int32 {

	return fut.needItemMap
}

func (fut *FaBaoUpstarTemplate) Patch() (err error) {
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
	}

	//验证 next_id
	if fut.NextId != 0 {
		to := template.GetTemplateService().Get(int(fut.NextId), (*FaBaoUpstarTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*FaBaoUpstarTemplate)
			diffLevel := nextTemplate.Level - fut.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			fut.nextFaBaoUpstarTemplate = nextTemplate
		}
	}

	return nil
}

func (fut *FaBaoUpstarTemplate) PatchAfterCheck() {

}

func (fut *FaBaoUpstarTemplate) Check() (err error) {
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

	//验证 fabao_percent
	err = validator.RangeValidate(float64(fut.FaBaoPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", fut.FaBaoPercent)
		err = template.NewTemplateFieldError("FaBaoPercent", err)
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

	for itemId, _ := range fut.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", fut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}
		itemTemplate := to.(*ItemTemplate)

		if itemTemplate.GetItemType() != itemtypes.ItemTypeFaBao {
			err = fmt.Errorf("UpstarItemId [%d]  invalid", fut.UpstarItemId)
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

func (fut *FaBaoUpstarTemplate) FileName() string {
	return "tb_fabao_upstar.json"
}

func init() {
	template.Register((*FaBaoUpstarTemplate)(nil))
}
