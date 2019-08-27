package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	fushitypes "fgame/fgame/game/fushi/types"
	"fmt"
)

type FuShiTemplate struct {
	*FuShiTemplateVO
	fushiType fushitypes.FuShiType
}

func (t *FuShiTemplate) GetFuShiType() fushitypes.FuShiType {
	return t.fushiType
}

func (t *FuShiTemplate) TemplateId() int {
	return t.Id
}

func (t *FuShiTemplate) Patch() (err error) {

	return
}

func (t *FuShiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 符石类型
	t.fushiType = fushitypes.FuShiType(t.Type)
	if !t.fushiType.Vaild() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	// 八卦秘境层数
	err = validator.MinValidate(float64(t.NeedBaGuaMiJing), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedBaGuaMiJing)
		return template.NewTemplateFieldError("NeedBaGuaMiJing", err)
	}

	return
}

func (t *FuShiTemplate) PatchAfterCheck() {
}

func (t *FuShiTemplate) FileName() string {
	return "tb_baguafushi.json"
}

func init() {
	template.Register((*FuShiTemplate)(nil))
}
