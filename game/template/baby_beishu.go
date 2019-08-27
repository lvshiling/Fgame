package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	commontypes "fgame/fgame/game/common/types"
	"fmt"
)

//宝宝属性倍数配置
type BabyBeiShuTemplate struct {
	*BabyBeiShuTemplateVO
	attrType commontypes.DanBeiPropertyType
}

func (t *BabyBeiShuTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyBeiShuTemplate) PatchAfterCheck() {}

func (t *BabyBeiShuTemplate) GetDanBeiType() commontypes.DanBeiPropertyType {
	return t.attrType
}

func (t *BabyBeiShuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *BabyBeiShuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.attrType = commontypes.DanBeiPropertyType(t.Type)
	if !t.attrType.Valid() {
		err = fmt.Errorf("[%s] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//属性点
	err = validator.MinValidate(float64(t.Attr), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attr)
		return template.NewTemplateFieldError("Attr", err)
	}

	return nil
}

func (t *BabyBeiShuTemplate) FileName() string {
	return "tb_baobao_danbei.json"
}

func init() {
	template.Register((*BabyBeiShuTemplate)(nil))
}
