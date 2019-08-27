package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//打折配置
type BargainTemplate struct {
	*BargainTemplateVo
	nextTemp *BargainTemplate
}

func (t *BargainTemplate) TemplateId() int {
	return t.Id
}

func (t *BargainTemplate) RandomDaZhe() int32 {
	return int32(mathutils.RandomRange(int(t.DazheMin), int(t.DazheMax+1)))
}

func (t *BargainTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *BargainTemplate) PatchAfterCheck() {
}

func (t *BargainTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//打折次数
	err = validator.MinValidate(float64(t.BargainTimes), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BargainTimes)
		return template.NewTemplateFieldError("BargainTimes", err)
	}

	// nextId
	if t.NextId != 0 {
		nextTempObj := template.GetTemplateService().Get(int(t.NextId), (*BargainTemplate)(nil))
		if nextTempObj != nil {
			nextTemp := nextTempObj.(*BargainTemplate)
			diff := nextTemp.BargainTimes - t.BargainTimes
			if diff != 1 {
				err = fmt.Errorf("[%d] invalid", t.BargainTimes)
				return template.NewTemplateFieldError("BargainTimes", err)
			}
			t.nextTemp = nextTemp
		}
	}

	//打折下限
	err = validator.MinValidate(float64(t.DazheMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DazheMin)
		return template.NewTemplateFieldError("DazheMin", err)
	}
	//打折上限
	err = validator.MinValidate(float64(t.DazheMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DazheMax)
		return template.NewTemplateFieldError("DazheMax", err)
	}

	return nil
}

func (edt *BargainTemplate) FileName() string {
	return "tb_bargain.json"
}

func init() {
	template.Register((*BargainTemplate)(nil))
}
