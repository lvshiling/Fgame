package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	jxtypes "fgame/fgame/game/juexue/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

//绝学配置
type JueXueTemplate struct {
	*JueXueTemplateVO
	juexueType       jxtypes.JueXueType
	jxsType          jxtypes.JueXueStageType
	useItempTemplate *ItemTemplate
}

func (jxt *JueXueTemplate) TemplateId() int {
	return jxt.Id
}

func (jxt *JueXueTemplate) GetType() jxtypes.JueXueType {
	return jxt.juexueType
}

func (jxt *JueXueTemplate) GetInsight() jxtypes.JueXueStageType {
	return jxt.jxsType
}

func (jxt *JueXueTemplate) GetUseItempTemplate() *ItemTemplate {
	return jxt.useItempTemplate
}

func (jxt *JueXueTemplate) PatchAfterCheck() {
}

func (jxt *JueXueTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(jxt.FileName(), jxt.TemplateId(), err)
			return
		}
	}()

	jxt.juexueType = jxtypes.JueXueType(jxt.Type)
	if !jxt.juexueType.Valid() {
		err = fmt.Errorf("[%d] invalid", jxt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 is_insight
	jxt.jxsType = jxtypes.JueXueStageType(jxt.IsInsight)
	if !jxt.jxsType.Valid() {
		err = fmt.Errorf("[%d] invalid", jxt.IsInsight)
		err = template.NewTemplateFieldError("IsInsight", err)
		return
	}

	//验证 need_item
	if jxt.NeedItemId != 0 {
		to := template.GetTemplateService().Get(int(jxt.NeedItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", jxt.NeedItemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return err
		}

		err = validator.MinValidate(float64(jxt.NeedItemNum), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", jxt.NeedItemNum)
			return template.NewTemplateFieldError("NeedItemNum", err)
		}

		jxt.useItempTemplate = to.(*ItemTemplate)
	}

	//验证 next_id
	if jxt.NextId != 0 {
		to := template.GetTemplateService().Get(int(jxt.NextId), (*JueXueTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", jxt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return err
		}

		nextTemplate := to.(*JueXueTemplate)
		if nextTemplate.Type != jxt.Type {
			err = fmt.Errorf("[%d] invalid", nextTemplate.Type, jxt.Type)
			err = template.NewTemplateFieldError("Type", err)
			return err
		}

		if nextTemplate.jxsType == jxtypes.JueXueStageTypeAorU {
			diffLevel := nextTemplate.Level - jxt.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				err = template.NewTemplateFieldError("Level", err)
				return err
			}
		}

		if nextTemplate.jxsType == jxtypes.JueXueStageTypeInsight {
			if jxt.jxsType == jxtypes.JueXueStageTypeInsight {
				diffLevel := nextTemplate.Level - jxt.Level
				if diffLevel != 1 {
					err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
					err = template.NewTemplateFieldError("Level", err)
					return err
				}
			}
		}
	}

	return nil
}

func (jxt *JueXueTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(jxt.FileName(), jxt.TemplateId(), err)
			return
		}
	}()

	//验证 power
	err = validator.MinValidate(float64(jxt.Power), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", jxt.Power)
		return template.NewTemplateFieldError("Power", err)
	}

	//验证 skill
	to := template.GetTemplateService().Get(int(jxt.Skill), (*SkillTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", jxt.Skill)
		err = template.NewTemplateFieldError("Skill", err)
		return err
	}
	typ := to.(*SkillTemplate).GetSkillFirstType()
	if typ != skilltypes.SkillFirstTypeJueXue {
		err = fmt.Errorf("[%d] invalid", jxt.Skill)
		err = template.NewTemplateFieldError("Skill", err)
		return err
	}

	return nil
}

func (jxt *JueXueTemplate) FileName() string {
	return "tb_juexue.json"
}

func init() {
	template.Register((*JueXueTemplate)(nil))
}
