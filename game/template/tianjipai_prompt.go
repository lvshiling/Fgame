package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//天机牌提示配置
type TianJiPaiPromptTemplate struct {
	*TianJiPaiPromptTemplateVO
}

func (tt *TianJiPaiPromptTemplate) TemplateId() int {
	return tt.Id
}

func (tt *TianJiPaiPromptTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (tt *TianJiPaiPromptTemplate) PatchAfterCheck() {

}

func (tt *TianJiPaiPromptTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	if tt.ModuleOpenedId != 0 {
		to := template.GetTemplateService().Get(int(tt.ModuleOpenedId), (*ModuleOpenedTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", tt.ModuleOpenedId)
			return template.NewTemplateFieldError("ModuleOpenedId", err)
		}
	}

	err = validator.MinValidate(float64(tt.IsAuto), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("IsAuto", err)
		return
	}

	return nil
}

func (tt *TianJiPaiPromptTemplate) FileName() string {
	return "tb_tianjipai_prompt.json"
}

func init() {
	template.Register((*TianJiPaiPromptTemplate)(nil))
}
