package template

import (
	"fgame/fgame/core/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fmt"
)

//神铸系统物品配置
type SystemShenZhuUseTemplate struct {
	*SystemShenZhuUseTemplateVO
	useItemMap map[additionsystypes.SlotPositionType]*ItemTemplate //神铸消耗物品
}

func (mclt *SystemShenZhuUseTemplate) TemplateId() int {
	return mclt.Id
}

func (mclt *SystemShenZhuUseTemplate) GetUseItemByPos(pos additionsystypes.SlotPositionType) *ItemTemplate {
	temp, ok := mclt.useItemMap[pos]
	if !ok {
		return nil
	}
	return temp
}

func (mclt *SystemShenZhuUseTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mclt.FileName(), mclt.TemplateId(), err)
			return
		}
	}()

	mclt.useItemMap = make(map[additionsystypes.SlotPositionType]*ItemTemplate)
	//验证 Pos1UseItem
	pos1UseItemTemplateVO := template.GetTemplateService().Get(int(mclt.Pos1UseItemId), (*ItemTemplate)(nil))
	if pos1UseItemTemplateVO == nil {
		err = fmt.Errorf("[%d] invalid", mclt.Pos1UseItemId)
		err = template.NewTemplateFieldError("Pos1UseItemId", err)
		return
	}
	mclt.useItemMap[additionsystypes.SlotPositionTypeOne] = pos1UseItemTemplateVO.(*ItemTemplate)

	//验证 Pos2UseItem
	pos2UseItemTemplateVO := template.GetTemplateService().Get(int(mclt.Pos2UseItemId), (*ItemTemplate)(nil))
	if pos2UseItemTemplateVO == nil {
		err = fmt.Errorf("[%d] invalid", mclt.Pos2UseItemId)
		err = template.NewTemplateFieldError("Pos2UseItemId", err)
		return
	}
	mclt.useItemMap[additionsystypes.SlotPositionTypeTwo] = pos2UseItemTemplateVO.(*ItemTemplate)

	//验证 Pos3UseItem
	pos3UseItemTemplateVO := template.GetTemplateService().Get(int(mclt.Pos3UseItemId), (*ItemTemplate)(nil))
	if pos3UseItemTemplateVO == nil {
		err = fmt.Errorf("[%d] invalid", mclt.Pos3UseItemId)
		err = template.NewTemplateFieldError("Pos3UseItemId", err)
		return
	}
	mclt.useItemMap[additionsystypes.SlotPositionTypeThree] = pos3UseItemTemplateVO.(*ItemTemplate)

	//验证 Pos4UseItem
	pos4UseItemTemplateVO := template.GetTemplateService().Get(int(mclt.Pos4UseItemId), (*ItemTemplate)(nil))
	if pos4UseItemTemplateVO == nil {
		err = fmt.Errorf("[%d] invalid", mclt.Pos4UseItemId)
		err = template.NewTemplateFieldError("Pos4UseItemId", err)
		return
	}
	mclt.useItemMap[additionsystypes.SlotPositionTypeFour] = pos4UseItemTemplateVO.(*ItemTemplate)

	return nil
}

func (mclt *SystemShenZhuUseTemplate) Check() (err error) {
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
func (mclt *SystemShenZhuUseTemplate) PatchAfterCheck() {

}
func (mclt *SystemShenZhuUseTemplate) FileName() string {
	return "tb_system_shenzhu_use.json"
}

func init() {
	template.Register((*SystemShenZhuUseTemplate)(nil))
}
