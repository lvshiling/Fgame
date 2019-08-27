package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
)

type CharacterLevelTemplate struct {
	*CharacterLevelTemplateVO
	expRatio          float64
	nextLevelTemplate *CharacterLevelTemplate
}

func (mt *CharacterLevelTemplate) GetExpRatio() float64 {
	return mt.expRatio
}

func (mt *CharacterLevelTemplate) TemplateId() int {
	return mt.Id
}

func (mt *CharacterLevelTemplate) GetNextLevelTemplate() *CharacterLevelTemplate {
	return mt.nextLevelTemplate
}

func (mt *CharacterLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()
	if mt.NextId != 0 {
		tempNextLevelTemplate := template.GetTemplateService().Get(int(mt.NextId), (*CharacterLevelTemplate)(nil))
		mt.nextLevelTemplate = tempNextLevelTemplate.(*CharacterLevelTemplate)
	}
	if mt.nextLevelTemplate == nil {
		mt.expRatio = float64(mt.Experience) / float64(mt.HotrexpModul)
	} else {
		mt.expRatio = float64(mt.nextLevelTemplate.Experience) / float64(mt.nextLevelTemplate.HotrexpModul)
	}
	return nil
}

func (mt *CharacterLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()
	err = validator.MinValidate(float64(mt.TpReplyTime), float64(0), false)
	if err != nil {
		return
	}

	return nil
}
func (mt *CharacterLevelTemplate) PatchAfterCheck() {

}

func (mt *CharacterLevelTemplate) FileName() string {
	return "tb_character_level.json"
}
func init() {
	template.Register((*CharacterLevelTemplate)(nil))
}
