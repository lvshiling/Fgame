package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	alliancetypes "fgame/fgame/game/alliance/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

//活动模板配置
type AllianceSkillTemplate struct {
	*AllianceSkillTemplateVO
	skillType alliancetypes.AllianceSkillType
	nextTemp  *AllianceSkillTemplate
}

func (at *AllianceSkillTemplate) Patch() (err error) {
	return
}
func (at *AllianceSkillTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(at.FileName(), at.TemplateId(), err)
			return
		}
	}()

	//仙术类型
	at.skillType = alliancetypes.AllianceSkillType(at.Type)
	if !at.skillType.Valid() {
		err = fmt.Errorf("[%d] invalid", at.Type)
		return template.NewTemplateFieldError("type", err)
	}

	//验证 skill_id
	if at.SkillId !=  0 {
		tempSkillTemplate := template.GetTemplateService().Get(int(at.SkillId), (*SkillTemplate)(nil))
		if tempSkillTemplate == nil {
			err = fmt.Errorf("[%d] invalid", at.SkillId)
			return template.NewTemplateFieldError("SkillId", err)
		}
		skilltyp := tempSkillTemplate.(*SkillTemplate).GetSkillFirstType()
		if skilltyp != skilltypes.SkillFirstTypeAlliance {
			err = fmt.Errorf("[%d] invalid", at.SkillId)
			return template.NewTemplateFieldError("SkillId", err)
		}
	}

	//仙术等级
	err = validator.MinValidate(float64(at.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", at.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//升级需的贡献
	err = validator.MinValidate(float64(at.NeedContribution), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", at.NeedContribution)
		return template.NewTemplateFieldError("NeedContribution", err)
	}

	//升级所需仙盟等级
	err = validator.MinValidate(float64(at.NeedUnionLevel), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", at.NeedUnionLevel)
		return template.NewTemplateFieldError("NeedUnionLevel", err)
	}

	//仙术开启所需仙盟等级
	err = validator.MinValidate(float64(at.OpenNeedLevel), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", at.OpenNeedLevel)
		return template.NewTemplateFieldError("OpenNeedLevel", err)
	}

	//验证 next_id
	if at.NextId != 0 {
		diff := at.NextId - int32(at.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", at.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(at.NextId), (*AllianceSkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", at.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		at.nextTemp = to.(*AllianceSkillTemplate)
	}

	return
}

func (at *AllianceSkillTemplate) GetNextTemplate() *AllianceSkillTemplate {
	return at.nextTemp
}

func (at *AllianceSkillTemplate) PatchAfterCheck() {

}
func (at *AllianceSkillTemplate) TemplateId() int {
	return at.Id
}

func (at *AllianceSkillTemplate) FileName() string {
	return "tb_union_skill.json"
}

func (at *AllianceSkillTemplate) GetSkillType() alliancetypes.AllianceSkillType {
	return at.skillType
}

func init() {
	template.Register((*AllianceSkillTemplate)(nil))
}
