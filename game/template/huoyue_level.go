package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	questtypes "fgame/fgame/game/quest/types"
	"fmt"
)

//活跃度配置
type HuoYueLevelTemplate struct {
	*HuoYueLevelTemplateVO
}

func (tt *HuoYueLevelTemplate) TemplateId() int {
	return tt.Id
}

func (tt *HuoYueLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (tt *HuoYueLevelTemplate) PatchAfterCheck() {

}

func (tt *HuoYueLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()

	//quest_id
	if tt.QuestId != 0 {
		to := template.GetTemplateService().Get(int(tt.QuestId), (*QuestTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", tt.QuestId)
			return template.NewTemplateFieldError("QuestId", err)
		}

		questTemplate := to.(*QuestTemplate)
		if questTemplate.GetQuestType() != questtypes.QuestTypeLiveness {
			err = fmt.Errorf("[%d] invalid", tt.QuestId)
			return template.NewTemplateFieldError("QuestId", err)
		}
	}

	err = validator.MinValidate(float64(tt.LevelMin), float64(1), true)
	if err != nil {
		err = template.NewTemplateFieldError("LevelMin", err)
		return
	}

	err = validator.MinValidate(float64(tt.LevelMax), float64(tt.LevelMin), true)
	if err != nil {
		err = template.NewTemplateFieldError("LevelMax", err)
		return
	}

	err = validator.MinValidate(float64(tt.HuoYue), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("HuoYue", err)
		return
	}

	err = validator.MinValidate(float64(tt.XueLian), float64(0), true)
	if err != nil {
		err = template.NewTemplateFieldError("XueLian", err)
		return
	}

	return nil
}

func (tt *HuoYueLevelTemplate) FileName() string {
	return "tb_huoyue_level.json"
}

func init() {
	template.Register((*HuoYueLevelTemplate)(nil))
}
