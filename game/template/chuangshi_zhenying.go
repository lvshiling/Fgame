package template

import (
	"fgame/fgame/core/template"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fmt"
)

type ChuangShiZhenyingTemplate struct {
	*ChuangShiZhenyingTemplateVO
	campType chuangshitypes.ChuangShiCampType
}

func (t *ChuangShiZhenyingTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiZhenyingTemplate) GetCampType() chuangshitypes.ChuangShiCampType {
	return t.campType
}

func (t *ChuangShiZhenyingTemplate) FileName() string {
	return "tb_chuangshi_zhenying.json"
}

func (t *ChuangShiZhenyingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.campType = chuangshitypes.ChuangShiCampType(t.Camp)

	return
}

func (t *ChuangShiZhenyingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	if !t.campType.Valid() {
		err = fmt.Errorf("[%d] 无效", t.Camp)
		return template.NewTemplateFieldError("Camp", err)
	}
	return
}

func (t *ChuangShiZhenyingTemplate) PatchAfterCheck() {

}

func init() {
	template.Register((*ChuangShiZhenyingTemplate)(nil))
}
