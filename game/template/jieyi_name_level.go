package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

//结义信物等级配置
type JieYiNameLevelTemplate struct {
	*JieYiNameLevelTemplateVO
	nextTemp *JieYiNameLevelTemplate
}

func (t *JieYiNameLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *JieYiNameLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return
}

func (t *JieYiNameLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*JieYiNameLevelTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*JieYiNameLevelTemplate)
		diff := t.nextTemp.Level - int32(t.Level)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	// 验证等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	// 验证升级成功率
	err = validator.RangeValidate(float64(t.UpLevPercent), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpLevPercent)
		return template.NewTemplateFieldError("Rate", err)
	}

	// 验证声威值
	err = validator.MinValidate(float64(t.UseShengWei), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 hp
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	//验证 attack
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	//验证 defence
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	//验证 TimesMin
	err = validator.RangeValidate(float64(t.TimesMin), float64(0), true, float64(t.TimesMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		err = template.NewTemplateFieldError("TimesMin", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(t.TimesMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	//验证 TimesMax
	err = validator.MinValidate(float64(t.TimesMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		err = template.NewTemplateFieldError("TimesMax", err)
		return
	}

	// 验证声威等级最小掉落
	err = validator.RangeValidate(float64(t.DeathMinLevel), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DeathMinLevel)
		return template.NewTemplateFieldError("DeathMinLevel", err)
	}

	// 验证声威等级最大掉落
	err = validator.RangeValidate(float64(t.DeathMaxLevel), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DeathMaxLevel)
		return template.NewTemplateFieldError("DeathMaxLevel", err)
	}
	if t.DeathMinLevel > t.DeathMaxLevel {
		err = fmt.Errorf("[%d] and [%d]invalid", t.DeathMinLevel, t.DeathMaxLevel)
		return template.NewTemplateFieldError("DeathMinLevel and DeathMaxLevel", err)
	}

	// 验证死亡掉落概率
	err = validator.RangeValidate(float64(t.DropPercent), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropPercent)
		return template.NewTemplateFieldError("DropPercent", err)
	}

	// 验证掉落在地的声威值
	err = validator.MinValidate(float64(t.StarCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.StarCount)
		return template.NewTemplateFieldError("StarCount", err)
	}

	return nil
}

func (t *JieYiNameLevelTemplate) PatchAfterCheck() {
}

func (t *JieYiNameLevelTemplate) FileName() string {
	return "tb_jieyi_weiming_level.json"
}

func init() {
	template.Register((*JieYiNameLevelTemplate)(nil))
}
