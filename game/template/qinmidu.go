package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//亲密度配置
type QinMiDuTemplate struct {
	*QinMiDuTemplateVO
}

func (qt *QinMiDuTemplate) TemplateId() int {
	return qt.Id
}

func (qt *QinMiDuTemplate) PatchAfterCheck() {

}

func (qt *QinMiDuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(qt.FileName(), qt.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (qt *QinMiDuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(qt.FileName(), qt.TemplateId(), err)
			return
		}
	}()

	err = validator.MinValidate(float64(qt.QinMiDuMin), float64(1), true)
	if err != nil {
		return template.NewTemplateFieldError("QinMiDuMin", err)
	}

	if qt.NextId != 0 {
		err = validator.MinValidate(float64(qt.QinMiDuMax), float64(qt.QinMiDuMin), true)
		if err != nil {
			return template.NewTemplateFieldError("QinMiDuMax", err)
		}
	}

	//验证 next_id
	if qt.NextId != 0 {
		diff := qt.NextId - int32(qt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", qt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return err
		}
		to := template.GetTemplateService().Get(int(qt.NextId), (*QinMiDuTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", qt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		qto := to.(*QinMiDuTemplate)
		//验证 level
		diffLeve := qto.QinMiDuMin - qt.QinMiDuMax
		if diffLeve != 1 {
			err = fmt.Errorf("[%d] invalid", qt.QinMiDuMax)
			return template.NewTemplateFieldError("QinMiDuMax", err)
		}
	}

	return nil
}

func (qt *QinMiDuTemplate) FileName() string {
	return "tb_qinmidu.json"
}

func init() {
	template.Register((*QinMiDuTemplate)(nil))
}
