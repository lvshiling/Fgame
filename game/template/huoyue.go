package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	funcopentypes "fgame/fgame/game/funcopen/types"
	questtypes "fgame/fgame/game/quest/types"
	"fmt"
)

//活跃度配置
type HuoYueTemplate struct {
	*HuoYueTemplateVO
	funcOpenTyp funcopentypes.FuncOpenType //模块功能开启类型
}

func (tt *HuoYueTemplate) TemplateId() int {
	return tt.Id
}

func (tt *HuoYueTemplate) GetFuncOpenTyp() funcopentypes.FuncOpenType {
	return tt.funcOpenTyp
}

func (tt *HuoYueTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (tt *HuoYueTemplate) PatchAfterCheck() {

}

func (tt *HuoYueTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	//module_opened_id
	if tt.ModuleOpenedId != 0 {
		to := template.GetTemplateService().Get(int(tt.ModuleOpenedId), (*ModuleOpenedTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", tt.ModuleOpenedId)
			return template.NewTemplateFieldError("ModuleOpenedId", err)
		}

		tt.funcOpenTyp = to.(*ModuleOpenedTemplate).GetFuncOpenType()
	}

	//quest_id
	if tt.Id != 0 {
		to := template.GetTemplateService().Get(int(tt.Id), (*QuestTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", tt.Id)
			return template.NewTemplateFieldError("Id", err)
		}

		questTemplate := to.(*QuestTemplate)
		if questTemplate.GetQuestType() != questtypes.QuestTypeLiveness {
			err = fmt.Errorf("[%d] invalid", tt.Id)
			return template.NewTemplateFieldError("Id", err)
		}
	}

	err = validator.MinValidate(float64(tt.RewardCountLimit), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("RewardCountLimit", err)
		return
	}

	return nil
}

func (tt *HuoYueTemplate) FileName() string {
	return "tb_huoyue.json"
}

func init() {
	template.Register((*HuoYueTemplate)(nil))
}
