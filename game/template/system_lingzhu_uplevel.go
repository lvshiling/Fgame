package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

type SystemLingZhuUpLevelTemplate struct {
	*SystemLingZhuUpLevelTemplateVO
	nextSystemLingZhuUpLevelTemplate *SystemLingZhuUpLevelTemplate
}

func (t *SystemLingZhuUpLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *SystemLingZhuUpLevelTemplate) GetNextStrengthenTemplate() *SystemLingZhuUpLevelTemplate {
	return t.nextSystemLingZhuUpLevelTemplate
}

//检查有效性
func (t *SystemLingZhuUpLevelTemplate) Check() (err error) {
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

	//Hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//攻击力
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//防御力
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//消耗物品数量
	err = validator.MinValidate(float64(t.UseItemCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItemCount)
		return template.NewTemplateFieldError("UseItemCount", err)
	}

	//升级成功率万分比
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

	//祝福值随机最小值
	err = validator.MinValidate(float64(t.AddMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMin)
		return template.NewTemplateFieldError("AddMin", err)
	}

	//祝福值随机最大值
	err = validator.MinValidate(float64(t.AddMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMax)
		return template.NewTemplateFieldError("AddMax", err)
	}

	//最大祝福值
	err = validator.MinValidate(float64(t.ZhufuMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		return template.NewTemplateFieldError("ZhufuMax", err)
	}

	//检查nextId可不可靠
	if t.nextSystemLingZhuUpLevelTemplate != nil {
		diff := t.nextSystemLingZhuUpLevelTemplate.Level - t.Level
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	return
}

//组合成需要的数据
func (t *SystemLingZhuUpLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if t.NextId == 0 {
		t.nextSystemLingZhuUpLevelTemplate = nil
	} else {
		temp := template.GetTemplateService().Get(int(t.NextId), (*SystemLingZhuUpLevelTemplate)(nil))
		t.nextSystemLingZhuUpLevelTemplate, _ = temp.(*SystemLingZhuUpLevelTemplate)
		if t.nextSystemLingZhuUpLevelTemplate == nil {
			err = fmt.Errorf("SystemLingZhuUpLevelTemplate[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("MagicConditionParameter", err)
		}
	}

	return
}

//检验后组合
func (t *SystemLingZhuUpLevelTemplate) PatchAfterCheck() {
}

func (t *SystemLingZhuUpLevelTemplate) FileName() string {
	return "tb_lingtong_lingzhu_level.json"
}

func init() {
	template.Register((*SystemLingZhuUpLevelTemplate)(nil))
}
