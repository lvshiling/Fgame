package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//暗器配置
type AnqiTemplate struct {
	*AnqiTemplateVO
	useItemTemplate   *ItemTemplate  //进阶物品
	skillTemplate     *SkillTemplate //暗器技能
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
}

func (t *AnqiTemplate) TemplateId() int {
	return t.Id
}

func (t *AnqiTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *AnqiTemplate) GetUseItemTemplate() *ItemTemplate {
	return t.useItemTemplate
}

func (t *AnqiTemplate) GetIsClear() bool {
	return t.IsClear != 0
}

func (t *AnqiTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 UseItem
	if t.UseItem != 0 {
		useItemTemplateVO := template.GetTemplateService().Get(int(t.UseItem), (*ItemTemplate)(nil))
		if useItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}
		t.useItemTemplate = useItemTemplateVO.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
	}

	//属性
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)

	return nil
}

func (t *AnqiTemplate) PatchAfterCheck() {

}
func (t *AnqiTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*AnqiTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp := to.(*AnqiTemplate)

		diff := nextTemp.Number - int32(t.Number)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
		return
	}

	//验证use_money
	err = validator.MinValidate(float64(t.UseMoney), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseMoney)
		err = template.NewTemplateFieldError("UseMoney", err)
		return
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

	//验证 ZhufuMax
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	//验证 UseYinliang
	err = validator.MinValidate(float64(t.UseYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseYinliang)
		err = template.NewTemplateFieldError("UseYinliang", err)
		return
	}

	//验证 ShidanLimit
	err = validator.MinValidate(float64(t.CulturingDanLimit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.CulturingDanLimit)
		err = template.NewTemplateFieldError("CulturingDanLimit", err)
		return
	}

	//暗器技能
	if t.Skill != 0 {
		to := template.GetTemplateService().Get(int(t.Skill), (*SkillTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.Skill)
			return template.NewTemplateFieldError("Skill", err)
		}
		t.skillTemplate = to.(*SkillTemplate)
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

	err = validator.MinValidate(float64(t.IsClear), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.IsClear)
		return template.NewTemplateFieldError("IsClear", err)
	}

	return nil
}

func (t *AnqiTemplate) FileName() string {
	return "tb_anqi.json"
}

func init() {
	template.Register((*AnqiTemplate)(nil))
}
