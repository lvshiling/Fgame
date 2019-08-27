package template

import (
	"fgame/fgame/core/template"
	coretypes "fgame/fgame/core/types"
	"fmt"
)

//限制区
type XianZhiQuYuTemplate struct {
	*XianZhiQuYuTemplateVO
	pos                     coretypes.Position
	nextXianZhiQuYuTemplate *XianZhiQuYuTemplate
}

func (t *XianZhiQuYuTemplate) TemplateId() int {
	return t.Id
}

func (t *XianZhiQuYuTemplate) GetPos() coretypes.Position {
	return t.pos
}

func (t *XianZhiQuYuTemplate) GetNext() *XianZhiQuYuTemplate {
	return t.nextXianZhiQuYuTemplate
}

func (t *XianZhiQuYuTemplate) Patch() (err error) {
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
		tempNextXianZhiQuYuTemplate := template.GetTemplateService().Get(int(t.NextId), (*XianZhiQuYuTemplate)(nil))
		if tempNextXianZhiQuYuTemplate != nil {
			t.nextXianZhiQuYuTemplate = tempNextXianZhiQuYuTemplate.(*XianZhiQuYuTemplate)
		}
	}
	return nil
}

func (t *XianZhiQuYuTemplate) PatchAfterCheck() {

}
func (t *XianZhiQuYuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.NextId != 0 && t.nextXianZhiQuYuTemplate == nil {
		err = fmt.Errorf("[%d] invalid", t.NextId)
		return template.NewTemplateFieldError("NextId", err)
	}

	return nil
}

func (t *XianZhiQuYuTemplate) FileName() string {
	return "tb_xianzhiquyu.json"
}

func init() {
	template.Register((*XianZhiQuYuTemplate)(nil))
}
