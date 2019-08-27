package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

type ShangguzhilingLinglianTaozhuangTemplate struct {
	*ShangguzhilingLinglianTaozhuangTemplateVO
	// nextTemplate *ShangguzhilingLinglianTaozhuangTemplate
}

// func (t *ShangguzhilingLinglianTaozhuangTemplate) GetNextTaoZhuangTemp() *ShangguzhilingLinglianTaozhuangTemplate {
// 	return t.nextTemplate
// }

func (t *ShangguzhilingLinglianTaozhuangTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingLinglianTaozhuangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingLinglianTaozhuangTemplate) PatchAfterCheck() {

}

func (t *ShangguzhilingLinglianTaozhuangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//需要星级
	err = validator.MinValidate(float64(t.NeedLevel), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedLevel)
		return template.NewTemplateFieldError("NeedLevel", err)
	}
	//验证 万分比
	err = validator.RangeValidate(float64(t.Percent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Percent)
		err = template.NewTemplateFieldError("Percent", err)
		return
	}
	// //验证下一Id
	// if t.NextId != 0 {
	// 	nextTempInterface := template.GetTemplateService().Get(int(t.NextId), (*ShangguzhilingLinglianTaozhuangTemplate)(nil))
	// 	if nextTempInterface == nil {
	// 		err = fmt.Errorf("ShangguzhilingLinglianTaozhuangTemplate [%d] no exist", t.NextId)
	// 		return template.NewTemplateFieldError("NextId", err)
	// 	}
	// 	nextTemp, ok := nextTempInterface.(*ShangguzhilingLinglianTaozhuangTemplate)
	// 	if !ok {
	// 		err = fmt.Errorf("ShangguzhilingLinglianTaozhuangTemplate assert [%d] no exist", t.NextId)
	// 		return template.NewTemplateFieldError("NextId", err)
	// 	}
	// 	t.nextTemplate = nextTemp
	// }

	return nil
}

func (t *ShangguzhilingLinglianTaozhuangTemplate) FileName() string {
	return "tb_sgzl_linglian_taozhuang.json"
}

func init() {
	template.Register((*ShangguzhilingLinglianTaozhuangTemplate)(nil))
}
