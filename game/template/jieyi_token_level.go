package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	"fmt"
)

//结义信物等级配置
type JieYiTokenLevelTemplate struct {
	*JieYiTokenLevelTemplateVO
	needItemMap map[int32]int32
	nextTemp    *JieYiTokenLevelTemplate
}

func (t *JieYiTokenLevelTemplate) TemplateId() int {
	return t.Id
}

func (t *JieYiTokenLevelTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *JieYiTokenLevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.needItemMap = make(map[int32]int32)
	//验证物品id
	if t.UseItemId != 0 {
		to := template.GetTemplateService().Get(int(t.UseItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemId)
			return template.NewTemplateFieldError("UseItemId", err)
		}

		err = validator.MinValidate(float64(t.UseItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.UseItemCount)
			return template.NewTemplateFieldError("UseItemCount", err)
		}
		t.needItemMap[t.UseItemId] = t.UseItemCount

	}

	return
}

func (t *JieYiTokenLevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*JieYiTokenLevelTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*JieYiTokenLevelTemplate)
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
	err = validator.RangeValidate(float64(t.Rate), float64(0), true, common.MAX_RATE, true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Rate)
		return template.NewTemplateFieldError("Rate", err)
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

	//验证 AddMin
	err = validator.RangeValidate(float64(t.AddMin), float64(0), true, float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMin)
		err = template.NewTemplateFieldError("AddMin", err)
		return
	}

	//验证 AddMax
	err = validator.MinValidate(float64(t.AddMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AddMax)
		err = template.NewTemplateFieldError("AddMax", err)
		return
	}

	//验证 zhufu_max
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	return nil
}

func (t *JieYiTokenLevelTemplate) PatchAfterCheck() {
}

func (t *JieYiTokenLevelTemplate) FileName() string {
	return "tb_jieyi_xinwu_strengthen.json"
}

func init() {
	template.Register((*JieYiTokenLevelTemplate)(nil))
}
