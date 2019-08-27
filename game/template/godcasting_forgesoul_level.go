package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type GodCastingForgeSoulLevelTemplate struct {
	*GodCastingForgeSoulLevelTemplateVO
	nextForgeSoulLevelTemplate *GodCastingForgeSoulLevelTemplate
}

func (t *GodCastingForgeSoulLevelTemplate) TemplateId() int {
	return t.Id
}

//升级模板的判断当前等级是不是最大级（如果达到最大等级上限是不会走到这里的）
func (t *GodCastingForgeSoulLevelTemplate) IsMaxLevel(maxLevel int32) bool {
	if t.Level > maxLevel {
		return true
	} else {
		return false
	}
}

func (t *GodCastingForgeSoulLevelTemplate) GetNextStrengthenTemplate() *GodCastingForgeSoulLevelTemplate {
	return t.nextForgeSoulLevelTemplate
}

//检查有效性
func (t *GodCastingForgeSoulLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//下一等级
	err = validator.MinValidate(float64(t.NextId), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NextId)
		return template.NewTemplateFieldError("NextId", err)
	}

	//消耗物品数量
	err = validator.MinValidate(float64(t.UseItemCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItemCount)
		return template.NewTemplateFieldError("UseItemCount", err)
	}

	//触发率万分比
	err = validator.RangeValidate(float64(t.ChufaRate), float64(0), true, float64(10000), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ChufaRate)
		return template.NewTemplateFieldError("ChufaRate", err)
	}

	//抵抗率万分比
	err = validator.RangeValidate(float64(t.DikangRate), float64(0), true, float64(10000), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DikangRate)
		return template.NewTemplateFieldError("DikangRate", err)
	}

	//升级率万分比
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(10000), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		return template.NewTemplateFieldError("UpdateWfb", err)
	}

	//最小次数
	err = validator.MinValidate(float64(t.TimesMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		return template.NewTemplateFieldError("TimesMin", err)
	}

	//最大次数
	err = validator.MinValidate(float64(t.TimesMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		return template.NewTemplateFieldError("TimesMax", err)
	}

	//神铸战力
	err = validator.MinValidate(float64(t.AddPower), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddPower)
		return template.NewTemplateFieldError("AddPower", err)
	}

	//检查nextId可不可靠
	if t.nextForgeSoulLevelTemplate != nil {
		diff := t.nextForgeSoulLevelTemplate.Level - t.Level
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	return
}

//组合成需要的数据
func (t *GodCastingForgeSoulLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.NextId == 0 {
		t.nextForgeSoulLevelTemplate = nil
	} else {
		temp := template.GetTemplateService().Get(int(t.NextId), (*GodCastingForgeSoulLevelTemplate)(nil))
		t.nextForgeSoulLevelTemplate, _ = temp.(*GodCastingForgeSoulLevelTemplate)
		if t.nextForgeSoulLevelTemplate == nil {
			err = fmt.Errorf("GodCastingForgeSoulLevelTemplate[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("MagicConditionParameter", err)
		}
	}

	return
}

//检验后组合
func (t *GodCastingForgeSoulLevelTemplate) PatchAfterCheck() {
}

func (t *GodCastingForgeSoulLevelTemplate) FileName() string {
	return "tb_shenzhuequip_duanhun_level.json"
}

func init() {
	template.Register((*GodCastingForgeSoulLevelTemplate)(nil))
}
