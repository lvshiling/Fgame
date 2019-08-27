package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/title/types"
	"fmt"
)

//神魔称号配置
type ShenMoTitleTemplate struct {
	*ShenMoTitleTemplateVO
}

func (t *ShenMoTitleTemplate) TemplateId() int {
	return t.Id
}

func (t *ShenMoTitleTemplate) PatchAfterCheck() {

}

func (t *ShenMoTitleTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *ShenMoTitleTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 kill_min
	err = validator.MinValidate(float64(t.KillMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KillMin)
		err = template.NewTemplateFieldError("KillMin", err)
		return
	}

	//验证 kill_max
	err = validator.MinValidate(float64(t.KillMax), float64(t.KillMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.KillMax)
		err = template.NewTemplateFieldError("KillMax", err)
		return
	}

	//验证 give_gongxun
	err = validator.MinValidate(float64(t.GiveGongXun), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GiveGongXun)
		err = template.NewTemplateFieldError("GiveGongXun", err)
		return
	}

	//验证 give_jifen
	err = validator.MinValidate(float64(t.GiveJiFen), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GiveJiFen)
		err = template.NewTemplateFieldError("GiveJiFen", err)
		return
	}

	//验证 title
	if t.Title != 0 {
		titleTempTemplate := template.GetTemplateService().Get(int(t.Title), (*TitleTemplate)(nil))
		if titleTempTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.Title)
			err = template.NewTemplateFieldError("Title", err)
			return
		}
		titleTemplate := titleTempTemplate.(*TitleTemplate)
		if titleTemplate.GetTitleType() != types.TitleTypeShenMo {
			err = fmt.Errorf("[%d] invalid", t.Title)
			err = template.NewTemplateFieldError("Title", err)
			return
		}
	}

	return nil
}

func (t *ShenMoTitleTemplate) FileName() string {
	return "tb_shenmo_title.json"
}

func init() {
	template.Register((*ShenMoTitleTemplate)(nil))
}
