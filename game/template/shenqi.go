package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"fmt"
)

//神器配置
type ShenQiTemplate struct {
	*ShenQiTemplateVO
	nextShenQiTemplate *ShenQiTemplate //下一级
	shenQiType         shenqitypes.ShenQiType
	skillTemplate      *SkillTemplate //技能
}

func (est *ShenQiTemplate) TemplateId() int {
	return est.Id
}

func (est *ShenQiTemplate) GetNextTemplate() *ShenQiTemplate {
	return est.nextShenQiTemplate
}

func (est *ShenQiTemplate) GetShenQiType() shenqitypes.ShenQiType {
	return est.shenQiType
}

func (est *ShenQiTemplate) GetSkillTemplate() *SkillTemplate {
	return est.skillTemplate
}

func (et *ShenQiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//下一级
	if et.NextId != 0 {
		diff := et.NextId - int32(et.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", et.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		tempNextShenQiTemplate := template.GetTemplateService().Get(int(et.NextId), (*ShenQiTemplate)(nil))
		if tempNextShenQiTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		et.nextShenQiTemplate = tempNextShenQiTemplate.(*ShenQiTemplate)
		diffLev := et.nextShenQiTemplate.Level - et.Level
		if diffLev != 1 {
			err = fmt.Errorf("[%d] invalid", et.nextShenQiTemplate.Level)
			err = template.NewTemplateFieldError("Next Level", err)
			return
		}
	}

	return nil
}
func (et *ShenQiTemplate) PatchAfterCheck() {

}
func (et *ShenQiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//验证 ShenQiType
	et.shenQiType = shenqitypes.ShenQiType(et.ShenQiType)
	if !et.shenQiType.Valid() {
		err = fmt.Errorf("[%d] invalid", et.ShenQiType)
		return template.NewTemplateFieldError("ShenQiType", err)
	}

	//等级
	err = validator.MinValidate(float64(et.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", et.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//技能
	if et.SkillId != 0 {
		to := template.GetTemplateService().Get(int(et.SkillId), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", et.SkillId)
			return template.NewTemplateFieldError("SkillId", err)
		}
		et.skillTemplate = to.(*SkillTemplate)
	}

	return nil
}

func (edt *ShenQiTemplate) FileName() string {
	return "tb_shenqi.json"
}

func init() {
	template.Register((*ShenQiTemplate)(nil))
}
