package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/mingge/types"
	"fmt"
)

//命格单倍配置
type MingGeDanBeiTemplate struct {
	*MingGeDanBeiTemplateVO
	mingGePropertyType types.MingGePropertyType
}

func (mt *MingGeDanBeiTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MingGeDanBeiTemplate) GetMingGePropertyType() types.MingGePropertyType {
	return mt.mingGePropertyType
}

func (mt *MingGeDanBeiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	mt.mingGePropertyType = types.MingGePropertyType(mt.Type)
	if !mt.mingGePropertyType.Valid() {
		err = fmt.Errorf("[%d] invalid", mt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	return nil
}

func (mt *MingGeDanBeiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	err = validator.MinValidate(float64(mt.Attr), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Attr)
		err = template.NewTemplateFieldError("Attr", err)
		return
	}

	err = validator.MinValidate(float64(mt.Rate), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Rate)
		err = template.NewTemplateFieldError("Rate", err)
		return
	}
	return nil
}

func (mt *MingGeDanBeiTemplate) PatchAfterCheck() {

}

func (mt *MingGeDanBeiTemplate) FileName() string {
	return "tb_mingge_danbei.json"
}

func init() {
	template.Register((*MingGeDanBeiTemplate)(nil))
}
