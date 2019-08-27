package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

//战翼皮肤升星配置
type WingUpstarTemplate struct {
	*WingUpstarTemplateVO
	needItemMap            map[int32]int32 //升星需要物品
	nextWingUpstarTemplate *WingUpstarTemplate
	useItemTemplate        *ItemTemplate
}

func (wut *WingUpstarTemplate) TemplateId() int {
	return wut.Id
}

func (wut *WingUpstarTemplate) GetNeedItemMap() map[int32]int32 {
	return wut.needItemMap
}

func (wut *WingUpstarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wut.FileName(), wut.TemplateId(), err)
			return
		}
	}()

	wut.needItemMap = make(map[int32]int32)
	//验证 upstar_item_id
	if wut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(wut.UpstarItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", wut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}

		err = validator.MinValidate(float64(wut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", wut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}
		wut.needItemMap[wut.UpstarItemId] = wut.UpstarItemCount

		wut.useItemTemplate = to.(*ItemTemplate)
	}

	//验证 next_id
	if wut.NextId != 0 {
		to := template.GetTemplateService().Get(int(wut.NextId), (*WingUpstarTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*WingUpstarTemplate)
			diffLevel := nextTemplate.Level - wut.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			wut.nextWingUpstarTemplate = nextTemplate
		}
	}

	return nil
}

func (wut *WingUpstarTemplate) PatchAfterCheck() {

}

func (wut *WingUpstarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(wut.FileName(), wut.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(wut.UpstarRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.UpstarRate)
		err = template.NewTemplateFieldError("UpstarRate", err)
		return
	}

	//验证 equip_percent
	err = validator.RangeValidate(float64(wut.WingPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.WingPercent)
		err = template.NewTemplateFieldError("WingPercent", err)
		return
	}

	//验证 level
	err = validator.MinValidate(float64(wut.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 hp
	err = validator.MinValidate(float64(wut.Hp), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(wut.Attack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(wut.Defence), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	if wut.useItemTemplate != nil {
		if wut.useItemTemplate.GetItemType() != itemtypes.ItemTypeWing {
			err = fmt.Errorf("UpstarItemId [%d]  invalid", wut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(wut.TimesMin), float64(0), true, float64(wut.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(wut.TimesMax), float64(wut.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(wut.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(wut.AddMin), float64(0), true, float64(wut.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(wut.AddMax), float64(wut.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(wut.ZhufuMax), float64(wut.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", wut.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	return nil
}

func (wut *WingUpstarTemplate) FileName() string {
	return "tb_wing_upstar.json"
}

func init() {
	template.Register((*WingUpstarTemplate)(nil))
}
