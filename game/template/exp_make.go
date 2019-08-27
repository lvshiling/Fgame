package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//炼制经验配置
type MadeTemplate struct {
	*MadeTemplateVO
	costList []int32
}

func (t *MadeTemplate) IsFullTimes(times int32) bool {
	if len(t.costList) <= int(times) {
		return true
	}
	return false
}

func (t *MadeTemplate) GetNeedCost(times int32) int32 {
	for curTimes, cost := range t.costList {
		if int32(curTimes) == times {
			return cost
		}
	}
	return -1
}

func (t *MadeTemplate) TemplateId() int {
	return t.Id
}

func (t *MadeTemplate) PatchAfterCheck() {
}

func (t *MadeTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//炼制消耗
	t.costList, err = utils.SplitAsIntArray(t.CostBase)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.CostBase)
		return template.NewTemplateFieldError("CostBase", err)
	}

	return nil
}

func (t *MadeTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 等级
	err = validator.MinValidate(float64(t.LevelMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelMin)
		err = template.NewTemplateFieldError("LevelMin", err)
		return
	}
	err = validator.MinValidate(float64(t.LevelMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelMax)
		err = template.NewTemplateFieldError("LevelMax", err)
		return
	}

	//验证 经验
	err = validator.MinValidate(float64(t.Exp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Exp)
		err = template.NewTemplateFieldError("Exp", err)
		return
	}

	//验证 经验点
	err = validator.MinValidate(float64(t.ExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ExpPoint)
		err = template.NewTemplateFieldError("ExpPoint", err)
		return
	}

	//消耗
	for _, cost := range t.costList {
		err = validator.MinValidate(float64(cost), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", cost)
			err = template.NewTemplateFieldError("CostBase", err)
			return
		}
	}

	return nil
}

func (t *MadeTemplate) FileName() string {
	return "tb_exp_make.json"
}

func init() {
	template.Register((*MadeTemplate)(nil))
}
