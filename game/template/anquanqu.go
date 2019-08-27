package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fmt"
)

//安全区
type AnquanquTemplate struct {
	*AnquanquTemplateVO
	pos                  coretypes.Position
	nextAnquanquTemplate *AnquanquTemplate
}

func (t *AnquanquTemplate) TemplateId() int {
	return t.Id
}

func (t *AnquanquTemplate) GetPos() coretypes.Position {
	return t.pos
}

func (t *AnquanquTemplate) GetNext() *AnquanquTemplate {
	return t.nextAnquanquTemplate
}

func (t *AnquanquTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.pos = coretypes.Position{
		X: t.PosX,
		Z: t.PosZ,
	}
	//验证 next_id
	if t.NextId != 0 {
		tempNextAnquanquTemplate := template.GetTemplateService().Get(int(t.NextId), (*AnquanquTemplate)(nil))
		if tempNextAnquanquTemplate != nil {
			t.nextAnquanquTemplate = tempNextAnquanquTemplate.(*AnquanquTemplate)
		}
	}
	return nil
}

func (t *AnquanquTemplate) PatchAfterCheck() {

}
func (t *AnquanquTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.NextId != 0 && t.nextAnquanquTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.NextId)
		return template.NewTemplateFieldError("NextId", err)
	}

	return nil
}

func (t *AnquanquTemplate) FileName() string {
	return "tb_anquanqu.json"
}

func init() {
	template.Register((*AnquanquTemplate)(nil))
}
