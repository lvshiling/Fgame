package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

//身法升星配置
type ShenFaUpstarTemplate struct {
	*ShenFaUpstarTemplateVO
	needItemMap              map[int32]int32 //升星需要物品
	nextShenFaUpstarTemplate *ShenFaUpstarTemplate
	useItemTemplate          *ItemTemplate
}

func (sut *ShenFaUpstarTemplate) TemplateId() int {
	return sut.Id
}

func (sut *ShenFaUpstarTemplate) GetNeedItemMap() map[int32]int32 {
	return sut.needItemMap
}

func (sut *ShenFaUpstarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(sut.FileName(), sut.TemplateId(), err)
			return
		}
	}()

	sut.needItemMap = make(map[int32]int32)
	//验证 upstar_item_id
	if sut.UpstarItemId != 0 {
		to := template.GetTemplateService().Get(int(sut.UpstarItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", sut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}

		err = validator.MinValidate(float64(sut.UpstarItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", sut.UpstarItemCount)
			return template.NewTemplateFieldError("UpstarItemCount", err)
		}
		sut.needItemMap[sut.UpstarItemId] = sut.UpstarItemCount

		sut.useItemTemplate = to.(*ItemTemplate)
	}

	//验证 next_id
	if sut.NextId != 0 {
		to := template.GetTemplateService().Get(int(sut.NextId), (*ShenFaUpstarTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*ShenFaUpstarTemplate)
			diffLevel := nextTemplate.Level - sut.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			sut.nextShenFaUpstarTemplate = nextTemplate
		}
	}

	return nil
}

func (sut *ShenFaUpstarTemplate) PatchAfterCheck() {

}

func (sut *ShenFaUpstarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(sut.FileName(), sut.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(sut.UpstarRate), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.UpstarRate)
		err = template.NewTemplateFieldError("UpstarRate", err)
		return
	}

	//验证 field_percent
	err = validator.RangeValidate(float64(sut.ShenFaPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.ShenFaPercent)
		err = template.NewTemplateFieldError("ShenFaPercent", err)
		return
	}

	//验证 level
	err = validator.MinValidate(float64(sut.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 hp
	err = validator.MinValidate(float64(sut.Hp), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(sut.Attack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(sut.Defence), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	if sut.useItemTemplate != nil {
		if sut.useItemTemplate.GetItemType() != itemtypes.ItemTypeShenFa {
			err = fmt.Errorf("UpstarItemId [%d]  invalid", sut.UpstarItemId)
			return template.NewTemplateFieldError("UpstarItemId", err)
		}
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(sut.TimesMin), float64(0), true, float64(sut.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(sut.TimesMax), float64(sut.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(sut.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(sut.AddMin), float64(0), true, float64(sut.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(sut.AddMax), float64(sut.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(sut.ZhufuMax), float64(sut.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", sut.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	return nil
}

func (sut *ShenFaUpstarTemplate) FileName() string {
	return "tb_shenfa_upstar.json"
}

func init() {
	template.Register((*ShenFaUpstarTemplate)(nil))
}
