package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type WushuangWeaponUnlockTemplate struct {
	*WushuangWeaponUnlockTemplateVO
	IsUnlock bool
}

func (t *WushuangWeaponUnlockTemplate) TemplateId() int {
	return t.Id
}

//检查有效性
func (t *WushuangWeaponUnlockTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//身体部位
	err = validator.RangeValidate(float64(t.Position), float64(0), true, float64(5), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Position)
		return template.NewTemplateFieldError("Position", err)
	}

	//是否解锁
	err = validator.RangeValidate(float64(t.IsJiesuo), float64(0), true, float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.IsJiesuo)
		return template.NewTemplateFieldError("IsJiesuo", err)
	}

	return
}

//组合成需要的数据
func (t *WushuangWeaponUnlockTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	return
}

//检验后组合
func (t *WushuangWeaponUnlockTemplate) PatchAfterCheck() {
	switch t.IsJiesuo {
	case int32(1):
		t.IsUnlock = true
	case int32(0):
		t.IsUnlock = false
	}
}

func (t *WushuangWeaponUnlockTemplate) FileName() string {
	return "tb_wushuang_jiesuo.json"
}

func init() {
	template.Register((*WushuangWeaponUnlockTemplate)(nil))
}
