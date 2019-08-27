package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type ShangguzhilingLinglianLockTemplate struct {
	*ShangguzhilingLinglianLockTemplateVO
	nextTemplate *ShangguzhilingLinglianLockTemplate
}

func (t *ShangguzhilingLinglianLockTemplate) GetNextLockTemp() *ShangguzhilingLinglianLockTemplate {
	return t.nextTemplate
}

func (t *ShangguzhilingLinglianLockTemplate) IsMaxTemp() bool {
	if t.NextId == 0 {
		return true
	}
	return false
}

func (t *ShangguzhilingLinglianLockTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingLinglianLockTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingLinglianLockTemplate) PatchAfterCheck() {

}

func (t *ShangguzhilingLinglianLockTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//锁定次数
	err = validator.MinValidate(float64(t.Times), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Times)
		return template.NewTemplateFieldError("Times", err)
	}
	//消耗锁定物品数量
	err = validator.MinValidate(float64(t.SuodingUseItemCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.SuodingUseItemCount)
		return template.NewTemplateFieldError("SuodingUseItemCount", err)
	}
	//验证下一Id
	if t.NextId != 0 {
		nextTempInterface := template.GetTemplateService().Get(int(t.NextId), (*ShangguzhilingLinglianLockTemplate)(nil))
		if nextTempInterface == nil {
			err = fmt.Errorf("ShangguzhilingLinglianLockTemplate [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp, ok := nextTempInterface.(*ShangguzhilingLinglianLockTemplate)
		if !ok {
			err = fmt.Errorf("ShangguzhilingLinglianLockTemplate assert [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		if nextTemp.Times != t.Times+1 {
			err = fmt.Errorf("ShangguzhilingLinglianLockTemplate [%d]", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemplate = nextTemp
	}

	return nil
}

func (t *ShangguzhilingLinglianLockTemplate) FileName() string {
	return "tb_sgzl_linglian_suoding.json"
}

func init() {
	template.Register((*ShangguzhilingLinglianLockTemplate)(nil))
}
