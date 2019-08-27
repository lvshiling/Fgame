package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

type ChuangShiChengFangJianSheTemplate struct {
	*ChuangShiChengFangJianSheTemplateVO
	nextTemp *ChuangShiChengFangJianSheTemplate
	//天气模板
	tianQiTemp *ChuangShiChengFangJianSheTianQiTemplate
	//属性
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
}

func (t *ChuangShiChengFangJianSheTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *ChuangShiChengFangJianSheTemplate) TemplateId() int {
	return t.Id
}

func (t *ChuangShiChengFangJianSheTemplate) GetNextTemp() *ChuangShiChengFangJianSheTemplate {
	return t.nextTemp
}

func (t *ChuangShiChengFangJianSheTemplate) GetTianQiTemp() *ChuangShiChengFangJianSheTianQiTemplate {
	return t.tianQiTemp
}

func (t *ChuangShiChengFangJianSheTemplate) FileName() string {
	return "tb_chuangshi_cfjianshe.json"
}

func (t *ChuangShiChengFangJianSheTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// next_id
	if t.NextId > 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*ChuangShiChengFangJianSheTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*ChuangShiChengFangJianSheTemplate)
	}

	//
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)

	//天气关联id
	if t.TianqiId > 0 {
		to := template.GetTemplateService().Get(int(t.TianqiId), (*ChuangShiChengFangJianSheTianQiTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.TianqiId)
			return template.NewTemplateFieldError("TianqiId", err)
		}
		t.tianQiTemp = to.(*ChuangShiChengFangJianSheTianQiTemplate)
	}

	return
}

func (t *ChuangShiChengFangJianSheTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	if t.nextTemp != nil {
		diff := t.nextTemp.Level - t.Level
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
	}

	//
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}
	//
	err = validator.MinValidate(float64(t.NeedExp), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedExp)
		return template.NewTemplateFieldError("NeedExp", err)
	}
	//
	err = validator.MinValidate(float64(t.FallMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FallMin)
		return template.NewTemplateFieldError("FallMin", err)
	}
	//
	err = validator.MinValidate(float64(t.FallMax), float64(1), true)
	if err != nil || t.FallMax < t.FallMin {
		err = fmt.Errorf("[%d] invalid", t.FallMax)
		return template.NewTemplateFieldError("FallMax", err)
	}
	//
	err = validator.RangeValidate(float64(t.FallPercent), float64(1), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FallPercent)
		return template.NewTemplateFieldError("FallPercent", err)
	}

	//生命
	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}
	//攻击
	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}
	//防御
	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	return
}

func (t *ChuangShiChengFangJianSheTemplate) PatchAfterCheck() {

}

func init() {
	template.Register((*ChuangShiChengFangJianSheTemplate)(nil))
}
