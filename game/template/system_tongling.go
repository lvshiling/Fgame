package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

//通灵配置
type SystemTongLingTemplate struct {
	*SystemTongLingTemplateVO
	nextSystemTongLingTemplate *SystemTongLingTemplate //下一级
}

func (mclt *SystemTongLingTemplate) TemplateId() int {
	return mclt.Id
}

func (mclt *SystemTongLingTemplate) GetNextTemplate() *SystemTongLingTemplate {
	return mclt.nextSystemTongLingTemplate
}

func (mclt *SystemTongLingTemplate) GetLevel() int32 {
	return mclt.Level
}

func (mclt *SystemTongLingTemplate) GetHp() int32 {
	return 0
}

func (mclt *SystemTongLingTemplate) GetAttack() int32 {
	return 0
}

func (mclt *SystemTongLingTemplate) GetDefence() int32 {
	return 0
}

func (mclt *SystemTongLingTemplate) GetPercent() int32 {
	return mclt.Percent
}

func (mclt *SystemTongLingTemplate) GetUseMoney() int32 {
	return 0
}

func (mclt *SystemTongLingTemplate) GetItemCount() int32 {
	return mclt.ItemCount
}

func (mclt *SystemTongLingTemplate) GetUpdateWfb() int32 {
	return mclt.UpdateWfb
}

func (mclt *SystemTongLingTemplate) GetZhufuMax() int32 {
	return mclt.ZhufuMax
}

func (mclt *SystemTongLingTemplate) GetAddMin() int32 {
	return mclt.AddMin
}

func (mclt *SystemTongLingTemplate) GetAddMax() int32 {
	return mclt.AddMax
}

func (mclt *SystemTongLingTemplate) GetTimesMin() int32 {
	return mclt.TimesMin
}

func (mclt *SystemTongLingTemplate) GetTimesMax() int32 {
	return mclt.TimesMax
}

func (mclt *SystemTongLingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//下一级
	if mclt.NextId != 0 {
		tempNextSystemTongLingTemplate := template.GetTemplateService().Get(int(mclt.NextId), (*SystemTongLingTemplate)(nil))
		if tempNextSystemTongLingTemplate == nil {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		mclt.nextSystemTongLingTemplate = tempNextSystemTongLingTemplate.(*SystemTongLingTemplate)
	}

	return nil
}

func (mclt *SystemTongLingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(mclt.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 next_id
	if mclt.NextId != 0 {
		diff := mclt.NextId - int32(mclt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(mclt.NextId), (*SystemTongLingTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mclt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		tongLingTemplate := to.(*SystemTongLingTemplate)

		diffLevel := tongLingTemplate.Level - mclt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", mclt.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(mclt.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证 ItemCount
	err = validator.MinValidate(float64(mclt.ItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.ItemCount)
		err = template.NewTemplateFieldError("ItemCount", err)
		return
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(mclt.TimesMin), float64(0), true, float64(mclt.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mclt.TimesMax), float64(mclt.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(mclt.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 AddMin
	err = validator.RangeValidate(float64(mclt.AddMin), float64(0), true, float64(mclt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(mclt.AddMax), float64(mclt.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(mclt.ZhufuMax), float64(mclt.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证percent
	err = validator.RangeValidate(float64(mclt.Percent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mclt.Percent)
		err = template.NewTemplateFieldError("Percent", err)
		return
	}

	return nil
}
func (mclt *SystemTongLingTemplate) PatchAfterCheck() {

}
func (mclt *SystemTongLingTemplate) FileName() string {
	return "tb_system_tongling.json"
}

func init() {
	template.Register((*SystemTongLingTemplate)(nil))
}
