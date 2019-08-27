package template

import (
	"fgame/fgame/core/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fmt"
)

//化灵系统物品配置
type SystemHuaLingUseTemplate struct {
	*SystemHuaLingUseTemplateVO
	useItemTemplate *ItemTemplate //化灵物品
}

func (mclt *SystemHuaLingUseTemplate) TemplateId() int {
	return mclt.Id
}

func (mclt *SystemHuaLingUseTemplate) GetUseItemTemplate() *ItemTemplate {
	return mclt.useItemTemplate
}

func (mclt *SystemHuaLingUseTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//验证 UseItem
	useItemTemplateVO := template.GetTemplateService().Get(int(mclt.UseItem), (*ItemTemplate)(nil))
	if useItemTemplateVO == nil {
		err = fmt.Errorf("[%d] invalid", mclt.UseItem)
		err = template.NewTemplateFieldError("UseItem", err)
		return
	}
	mclt.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

	return nil
}

func (mclt *SystemHuaLingUseTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	//验证SysType
	sysType := additionsystypes.AdditionSysType(mclt.SysType)
	if !sysType.Valid() {
		err = fmt.Errorf("[%d] invalid", mclt.SysType)
		return template.NewTemplateFieldError("SysType", err)
	}
	return nil
}
func (mclt *SystemHuaLingUseTemplate) PatchAfterCheck() {

}
func (mclt *SystemHuaLingUseTemplate) FileName() string {
	return "tb_system_hualing_use.json"
}

func init() {
	template.Register((*SystemHuaLingUseTemplate)(nil))
}
