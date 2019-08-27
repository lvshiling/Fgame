package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/scene/types"
	"fmt"
)

//屠龙常量配置
type TuLongConstantTemplate struct {
	*TuLongConstantTemplateVO
	mapTemplate             *MapTemplate
	bigEggBiologyTemplate   *BiologyTemplate
	smallEggBiologyTemplate *BiologyTemplate
}

func (tl *TuLongConstantTemplate) TemplateId() int {
	return tl.Id
}

func (tl *TuLongConstantTemplate) GetBigEggBiologyTemplate() *BiologyTemplate {
	return tl.bigEggBiologyTemplate
}

func (tl *TuLongConstantTemplate) GetSmallEggBiologyTemplate() *BiologyTemplate {
	return tl.smallEggBiologyTemplate
}

func (tl *TuLongConstantTemplate) GetMapTemplate() *MapTemplate {
	return tl.mapTemplate
}

func (tl *TuLongConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tl.FileName(), tl.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (tl *TuLongConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tl.FileName(), tl.TemplateId(), err)
			return
		}
	}()
	// err = validator.MinValidate(float64(tl.UnionRankQian), float64(1), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", tl.UnionRankQian)
	// 	return template.NewTemplateFieldError("UnionRankQian", err)
	// }

	to := template.GetTemplateService().Get(int(tl.BigEgg), (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", tl.BigEgg)
		return template.NewTemplateFieldError("BigEgg", err)
	}
	tl.bigEggBiologyTemplate = to.(*BiologyTemplate)
	if tl.bigEggBiologyTemplate.GetBiologyScriptType() != types.BiologyScriptTypeCrossBigEgg {
		err = fmt.Errorf("[%d] invalid", tl.BigEgg)
		return template.NewTemplateFieldError("BigEgg", err)
	}

	to = template.GetTemplateService().Get(int(tl.SmallEgg), (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", tl.SmallEgg)
		return template.NewTemplateFieldError("SmallEgg", err)
	}
	tl.smallEggBiologyTemplate = to.(*BiologyTemplate)
	if tl.smallEggBiologyTemplate.GetBiologyScriptType() != types.BiologyScriptTypeCrossSmallEgg {
		err = fmt.Errorf("[%d] invalid", tl.SmallEgg)
		return template.NewTemplateFieldError("SmallEgg", err)
	}

	to = template.GetTemplateService().Get(int(tl.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", tl.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}
	tl.mapTemplate = to.(*MapTemplate)
	if tl.mapTemplate.GetMapType() != types.SceneTypeCrossTuLong {
		err = fmt.Errorf("[%d] invalid", tl.MapId)
		return template.NewTemplateFieldError("MapId", err)
	}

	err = validator.MinValidate(float64(tl.BossTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tl.BossTime)
		return template.NewTemplateFieldError("BossTime", err)
	}

	err = validator.MinValidate(float64(tl.CaiJiTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", tl.CaiJiTime)
		return template.NewTemplateFieldError("CaiJiTime", err)
	}
	return nil
}
func (tl *TuLongConstantTemplate) PatchAfterCheck() {

}
func (tl *TuLongConstantTemplate) FileName() string {
	return "tb_tulong_constant.json"
}

func init() {
	template.Register((*TuLongConstantTemplate)(nil))
}
