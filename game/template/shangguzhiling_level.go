package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type ShangguzhilingLevelTemplate struct {
	*ShangguzhilingLevelTemplateVO
	nextLevelTemp *ShangguzhilingLevelTemplate
}

func (t *ShangguzhilingLevelTemplate) GetNextLevelTemp() *ShangguzhilingLevelTemplate {
	return t.nextLevelTemp
}

func (t *ShangguzhilingLevelTemplate) GeExperience() int64 {
	return t.Experience
}

func (t *ShangguzhilingLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *ShangguzhilingLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShangguzhilingLevelTemplate) PatchAfterCheck() {

}

func (t *ShangguzhilingLevelTemplate) Check() (err error) {
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
		nextTempInterface := template.GetTemplateService().Get(int(t.NextId), (*ShangguzhilingLevelTemplate)(nil))
		if nextTempInterface == nil {
			err = fmt.Errorf("ShangguzhilingLevelTemplate [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp, ok := nextTempInterface.(*ShangguzhilingLevelTemplate)
		if !ok {
			err = fmt.Errorf("ShangguzhilingLevelTemplate assert [%d] no exist", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		if nextTemp.Level != t.Level+1 {
			err = fmt.Errorf("ShangguzhilingLevelTemplate [%d] invalid, curLevel [%d], nextTempLevel [%d]", t.NextId, t.Level, nextTemp.Level)
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

	//宝箱CD
	err = validator.MinValidate(float64(t.BaoxiangCd), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BaoxiangCd)
		return template.NewTemplateFieldError("BaoxiangCd", err)
	}

	//宝箱关联掉落Id
	err = validator.MinValidate(float64(t.BaoxiangDrop), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BaoxiangDrop)
		return template.NewTemplateFieldError("BaoxiangDrop", err)
	}

	return nil
}

func (t *ShangguzhilingLevelTemplate) FileName() string {
	return "tb_sgzl_level.json"
}

func init() {
	template.Register((*ShangguzhilingLevelTemplate)(nil))
}
