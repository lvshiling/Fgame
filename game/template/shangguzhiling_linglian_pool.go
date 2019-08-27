package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type ShangguzhilingLinglianPoolTemplate struct {
	*ShangguzhilingLinglianPoolTemplateVO
	nextPoolTemp *ShangguzhilingLinglianPoolTemplate
}

func (t *ShangguzhilingLinglianPoolTemplate) GetNextPoolTemplate() *ShangguzhilingLinglianPoolTemplate {
	return t.nextPoolTemp
}

func (t *ShangguzhilingLinglianPoolTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingLinglianPoolTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingLinglianPoolTemplate) PatchAfterCheck() {

}

func (t *ShangguzhilingLinglianPoolTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//星级
	err = validator.MinValidate(float64(t.Star), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Star)
		return template.NewTemplateFieldError("Star", err)
	}

	//权重
	err = validator.MinValidate(float64(t.Rate), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Rate)
		return template.NewTemplateFieldError("Rate", err)
	}

	//下一星级模板
	if t.NextId != 0 {
		nextTempInterface := template.GetTemplateService().Get(int(t.NextId), (*ShangguzhilingLinglianPoolTemplate)(nil))
		if nextTempInterface == nil {
			err = fmt.Errorf("ShangguzhilingLinglianPoolTemplate [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp, ok := nextTempInterface.(*ShangguzhilingLinglianPoolTemplate)
		if !ok {
			err = fmt.Errorf("ShangguzhilingLinglianPoolTemplate assert [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		if nextTemp.Biaoshi != t.Biaoshi+1 {
			err = fmt.Errorf("ShangguzhilingLinglianPoolTemplate [%d] invalid, curBiaoshi [%d], nextTempBiaoshi [%d]", t.NextId, t.Biaoshi, nextTemp.Biaoshi)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextPoolTemp = nextTemp
	}

	//hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//攻击
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//防御
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//品质
	err = validator.MinValidate(float64(t.Quality), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Quality)
		return template.NewTemplateFieldError("Quality", err)
	}

	return nil
}

func (t *ShangguzhilingLinglianPoolTemplate) FileName() string {
	return "tb_sgzl_linglian_attr_pool.json"
}

func init() {
	template.Register((*ShangguzhilingLinglianPoolTemplate)(nil))
}
