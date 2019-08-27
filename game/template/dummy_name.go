package template

import (
	"fgame/fgame/core/template"
	dummytypes "fgame/fgame/game/dummy/types"
	"fmt"
)

//假人名配置
type DummyNameTemplate struct {
	*DummyNameTemplateVO
	dummyType dummytypes.DummyType
}

func (dt *DummyNameTemplate) TemplateId() int {
	return dt.Id
}

func (dt *DummyNameTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(dt.FileName(), dt.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (dt *DummyNameTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(dt.FileName(), dt.TemplateId(), err)
			return
		}
	}()

	//验证 type
	typ := dummytypes.DummyType(dt.Type)
	if !typ.Valid() {
		err = fmt.Errorf("[%d] Type", dt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}
	dt.dummyType = typ

	return nil
}

func (dt *DummyNameTemplate) PatchAfterCheck() {

}

func (dt *DummyNameTemplate) FileName() string {
	return "tb_dummy_name.json"
}

func (dt *DummyNameTemplate) GetDummyType() dummytypes.DummyType {
	return dt.dummyType
}

func init() {
	template.Register((*DummyNameTemplate)(nil))
}
