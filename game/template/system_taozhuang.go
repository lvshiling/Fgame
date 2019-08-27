package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	additionsystypes "fgame/fgame/game/additionsys/types"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

//套装配置
type SystemTaozhuangTemplate struct {
	*SystemTaozhuangTemplateVO
	//套装系统类型
	taozhuangType additionsystypes.AdditionSysType
	//套装装备品质
	taozhuangQuality itemtypes.ItemQualityType
}

func (t *SystemTaozhuangTemplate) TemplateId() int {
	return t.Id
}

func (t *SystemTaozhuangTemplate) GetTaozhuangType() additionsystypes.AdditionSysType {
	return t.taozhuangType
}

func (t *SystemTaozhuangTemplate) GetTaozhuangQuality() itemtypes.ItemQualityType {
	return t.taozhuangQuality
}

func (t *SystemTaozhuangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//套装系统类型
	t.taozhuangType = additionsystypes.AdditionSysType(t.Type)
	if !t.taozhuangType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//套装物品品质属性
	t.taozhuangQuality = itemtypes.ItemQualityType(t.Pos1Quality)
	if !t.taozhuangQuality.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Pos1Quality)
		return template.NewTemplateFieldError("Pos1Quality", err)
	}

	return nil
}

func (t *SystemTaozhuangTemplate) PatchAfterCheck() {

}

func (t *SystemTaozhuangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//套装数量
	if err = validator.MinValidate(float64(t.Number), float64(0), false); err != nil {
		err = template.NewTemplateFieldError("number", err)
		return
	}

	//属性加成
	if err = validator.MinValidate(float64(t.AttrPercent), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("AttrPercent", err)
		return
	}

	return nil
}

func (t *SystemTaozhuangTemplate) FileName() string {
	return "tb_system_taozhuang.json"
}

func init() {
	template.Register((*SystemTaozhuangTemplate)(nil))
}
