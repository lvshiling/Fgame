package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
)

//衣橱培养配置
type YiChuPeiYangTemplate struct {
	*YiChuPeiYangTemplateVO
	needItemMap              map[int32]int32 //培养需要物品
	nextYiChuPeiYangTemplate *YiChuPeiYangTemplate
	battlePropertyMap        map[propertytypes.BattlePropertyType]int64
}

func (t *YiChuPeiYangTemplate) TemplateId() int {
	return t.Id
}

func (t *YiChuPeiYangTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *YiChuPeiYangTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *YiChuPeiYangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	err = validator.MinValidate(float64(t.Hp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Hp)
		return template.NewTemplateFieldError("Hp", err)
	}

	err = validator.MinValidate(float64(t.Attack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Attack)
		return template.NewTemplateFieldError("Attack", err)
	}

	err = validator.MinValidate(float64(t.Defence), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Defence)
		return template.NewTemplateFieldError("Defence", err)
	}

	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = t.Hp
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = t.Attack
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = t.Defence

	t.needItemMap = make(map[int32]int32)
	if t.UseItem != 0 {
		to := template.GetTemplateService().Get(int(t.UseItem), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}

		err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}
		t.needItemMap[t.UseItem] = t.ItemCount
	}

	//验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*YiChuPeiYangTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*YiChuPeiYangTemplate)

			diffLevel := nextTemplate.Level - t.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			t.nextYiChuPeiYangTemplate = nextTemplate
		}
	}

	return nil
}

func (t *YiChuPeiYangTemplate) PatchAfterCheck() {

}

func (t *YiChuPeiYangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(t.UpstarWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpstarWfb)
		err = template.NewTemplateFieldError("UpstarWfb", err)
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

	//验证 level
	err = validator.MinValidate(float64(t.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 percent
	err = validator.MinValidate(float64(t.Percent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Percent)
		return template.NewTemplateFieldError("Percent", err)
	}

	//验证技能
	to := template.GetTemplateService().Get(int(t.SkillId), (*SkillTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}
	typ := to.(*SkillTemplate).GetSkillFirstType()
	if typ != skilltypes.SkillFirstTypeWardrobe {
		err = fmt.Errorf("[%d] invalid", t.SkillId)
		err = template.NewTemplateFieldError("SkillId", err)
		return
	}

	return nil
}

func (t *YiChuPeiYangTemplate) FileName() string {
	return "tb_yichu_peiyang.json"
}

func init() {
	template.Register((*YiChuPeiYangTemplate)(nil))
}
