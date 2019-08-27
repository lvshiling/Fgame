package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type ShangguzhilingLingwenLevelTemplate struct {
	*ShangguzhilingLingwenLevelTemplateVO
	nextLevelTemp *ShangguzhilingLingwenLevelTemplate
}

func (t *ShangguzhilingLingwenLevelTemplate) GetNextLevelTemp() *ShangguzhilingLingwenLevelTemplate {
	return t.nextLevelTemp
}

func (t *ShangguzhilingLingwenLevelTemplate) GetExperience() int64 {
	return t.Experience
}

func (t *ShangguzhilingLingwenLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingLingwenLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingLingwenLevelTemplate) PatchAfterCheck() {

}

func (t *ShangguzhilingLingwenLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//下一等级模板
	if t.NextId != 0 {
		nextTempInterface := template.GetTemplateService().Get(int(t.NextId), (*ShangguzhilingLingwenLevelTemplate)(nil))
		if nextTempInterface == nil {
			err = fmt.Errorf("ShangguzhilingLingwenLevelTemplate [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp, ok := nextTempInterface.(*ShangguzhilingLingwenLevelTemplate)
		if !ok {
			err = fmt.Errorf("ShangguzhilingLingwenLevelTemplate assert [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		if nextTemp.Level != t.Level+1 {
			err = fmt.Errorf("ShangguzhilingLingwenLevelTemplate [%d] invalid, curLevel [%d], nextTempLevel [%d]", t.NextId, t.Level, nextTemp.Level)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextLevelTemp = nextTemp
	}

	//升级所需经验
	err = validator.MinValidate(float64(t.Experience), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Experience)
		return template.NewTemplateFieldError("Experience", err)
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

	return nil
}

func (t *ShangguzhilingLingwenLevelTemplate) FileName() string {
	return "tb_sgzl_lingwen_level.json"
}

func init() {
	template.Register((*ShangguzhilingLingwenLevelTemplate)(nil))
}
