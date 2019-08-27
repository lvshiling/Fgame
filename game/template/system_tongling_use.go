package template

import (
	"fgame/fgame/core/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fmt"
)

//通灵系统物品配置
type SystemTongLingUseTemplate struct {
	*SystemTongLingUseTemplateVO
	useItemTemplate *ItemTemplate //通灵物品
}

func (mclt *SystemTongLingUseTemplate) TemplateId() int {
	return mclt.Id
}

func (mclt *SystemTongLingUseTemplate) GetUseItemTemplate() *ItemTemplate {
	return mclt.useItemTemplate
}

func (mclt *SystemTongLingUseTemplate) Patch() (err error) {
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

func (mclt *SystemTongLingUseTemplate) Check() (err error) {
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
func (mclt *SystemTongLingUseTemplate) PatchAfterCheck() {

}
func (mclt *SystemTongLingUseTemplate) FileName() string {
	return "tb_system_tongling_use.json"
}

func init() {
	template.Register((*SystemTongLingUseTemplate)(nil))
}
