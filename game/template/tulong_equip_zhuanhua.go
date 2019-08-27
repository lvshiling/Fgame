package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fmt"
)

//屠龙转化配置
type TuLongEquipZhuanHuaTemplate struct {
	*TuLongEquipZhuanHuaTemplateVO
	posType inventorytypes.BodyPositionType
}

func (t *TuLongEquipZhuanHuaTemplate) TemplateId() int {
	return t.Id
}

func (t *TuLongEquipZhuanHuaTemplate) GetPosType() inventorytypes.BodyPositionType {
	return t.posType
}

func (t *TuLongEquipZhuanHuaTemplate) Patch() (err error) {

	return nil
}

func (t *TuLongEquipZhuanHuaTemplate) PatchAfterCheck() {

}

func (t *TuLongEquipZhuanHuaTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.posType = inventorytypes.BodyPositionType(t.SubType)
	if !t.posType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SubType)
		return template.NewTemplateFieldError("Type", err)
	}

	//转生数
	if err = validator.MinValidate(float64(t.Level), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	return nil
}

func (t *TuLongEquipZhuanHuaTemplate) FileName() string {
	return "tb_tulongequip_zhuanhua.json"
}

func init() {
	template.Register((*TuLongEquipZhuanHuaTemplate)(nil))
}
