package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fmt"
)

type SystemAwakeUseTemplate struct {
	*SystemAwakeUseTemplateVO
	sysType additionsystypes.AdditionSysType
}

func (t *SystemAwakeUseTemplate) TemplateId() int {
	return t.Id
}

func (t *SystemAwakeUseTemplate) GetMinAwakeAdvanced() int32 {
	return t.NeedNumber
}

func (t *SystemAwakeUseTemplate) GetSysType() additionsystypes.AdditionSysType {
	return t.sysType
}

func (t *SystemAwakeUseTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证SysType
	sysType := additionsystypes.AdditionSysType(t.SysType)
	if !sysType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SysType)
		return template.NewTemplateFieldError("SysType", err)
	}
	t.sysType = sysType

	return
}

func (t *SystemAwakeUseTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证UseItem
	useItemTemplateVO := template.GetTemplateService().Get(int(t.UseItem), (*ItemTemplate)(nil))
	if useItemTemplateVO == nil {
		err = fmt.Errorf("[%d] invalid", t.UseItem)
		err = template.NewTemplateFieldError("UseItem", err)
		return
	}

	//验证need_number
	err = validator.MinValidate(float64(t.NeedNumber), 1, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedNumber)
		err = template.NewTemplateFieldError("NeedNumber", err)
		return
	}

	return
}

func (t *SystemAwakeUseTemplate) PatchAfterCheck() {

}

func (t *SystemAwakeUseTemplate) FileName() string {
	return "tb_system_juexing_use.json"
}

func init() {
	template.Register((*SystemAwakeUseTemplate)(nil))
}
