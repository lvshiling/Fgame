package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	activitytypes "fgame/fgame/game/activity/types"
	scenetypes "fgame/fgame/game/scene/types"

	"fmt"
)

//无间炼狱常量配置
type LianYuConstantTemplate struct {
	*LianYuConstantTemplateVO
	biologyTemplate *BiologyTemplate
	activityType    activitytypes.ActivityType
}

func (t *LianYuConstantTemplate) GetActivityType() activitytypes.ActivityType {
	return t.activityType
}

func (t *LianYuConstantTemplate) TemplateId() int {
	return t.Id
}

func (t *LianYuConstantTemplate) GetBiologyTemplate() *BiologyTemplate {
	return t.biologyTemplate
}

func (t *LianYuConstantTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	to := template.GetTemplateService().Get(int(t.BossId), (*BiologyTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.BossId)
		return template.NewTemplateFieldError("BossId", err)
	}
	t.biologyTemplate = to.(*BiologyTemplate)

	return nil
}

func (t *LianYuConstantTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//活动类型
	activityType := activitytypes.ActivityType(t.Type)
	if !activityType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return err
	}
	t.activityType = activityType

	err = validator.MinValidate(float64(t.BossTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BossTime)
		err = template.NewTemplateFieldError("BossTime", err)
		return
	}

	err = validator.MinValidate(float64(t.PlayerLimitCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PlayerLimitCount)
		err = template.NewTemplateFieldError("PlayerLimitCount", err)
		return
	}

	if t.biologyTemplate.GetBiologyScriptType() != scenetypes.BiologyScriptTypeCrossLianYuBoss {
		err = fmt.Errorf("[%d] invalid", t.BossId)
		err = template.NewTemplateFieldError("BossId", err)
		return
	}

	return nil
}
func (t *LianYuConstantTemplate) PatchAfterCheck() {

}
func (t *LianYuConstantTemplate) FileName() string {
	return "tb_lianyu_constant.json"
}

func init() {
	template.Register((*LianYuConstantTemplate)(nil))
}
