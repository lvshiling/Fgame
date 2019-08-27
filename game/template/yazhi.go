package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//收益率配置
type YaZhiTemplate struct {
	*YaZhiTemplateVO
}

func (t *YaZhiTemplate) TemplateId() int {
	return t.Id
}

func (t *YaZhiTemplate) PatchAfterCheck() {
}

func (t *YaZhiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *YaZhiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 最低等级
	err = validator.MinValidate(float64(t.LevelMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelMin)
		err = template.NewTemplateFieldError("LevelMin", err)
		return
	}
	//验证 最高等级
	err = validator.MinValidate(float64(t.LevelMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevelMax)
		err = template.NewTemplateFieldError("LevelMax", err)
		return
	}

	// 经验率
	err = validator.MinValidate(float64(t.ExpPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ExpPercent)
		err = template.NewTemplateFieldError("ExpPercent", err)
		return
	}

	// 掉宝率
	err = validator.MinValidate(float64(t.DropPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropPercent)
		err = template.NewTemplateFieldError("DropPercent", err)
		return
	}
	return nil
}

func (t *YaZhiTemplate) FileName() string {
	return "tb_yazhi.json"
}

func init() {
	template.Register((*YaZhiTemplate)(nil))
}
