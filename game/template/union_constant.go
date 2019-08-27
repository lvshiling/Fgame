package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

//仙盟常量配置
type UnionConstantTemplate struct {
	*UnionConstantTemplateVO
	limitRenameTime int64
}

func (t *UnionConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *UnionConstantTemplate) GetRenameLimitTime() int64 {
	return t.limitRenameTime
}

func (t *UnionConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *UnionConstantTemplate) PatchAfterCheck() {
	t.limitRenameTime = t.GaimingTime * int64(common.DAY)
}

func (t *UnionConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	err = validator.MinValidate(float64(t.GaimingTime), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GaimingTime)
		return template.NewTemplateFieldError("GaimingTime", err)
	}
	//最小常量
	err = validator.MinValidate(float64(t.GaimingItemCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GaimingItemCount)
		return template.NewTemplateFieldError("GaimingItemCount", err)
	}

	return nil
}

func (tt *UnionConstantTemplate) FileName() string {
	return "tb_union_constant.json"
}

func init() {
	template.Register((*UnionConstantTemplate)(nil))
}
