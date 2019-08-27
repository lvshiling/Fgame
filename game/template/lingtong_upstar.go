package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//灵体升星配置
type LingTongUpstarTemplate struct {
	*LingTongUpstarTemplateVO
	needItemMap       map[int32]int32 //升星需要物品
	useItemTemplate   *ItemTemplate
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	nextTemplate      *LingTongUpstarTemplate
}

func (t *LingTongUpstarTemplate) TemplateId() int {
	return t.Id
}

func (t *LingTongUpstarTemplate) GetNextId() int32 {
	return t.NextId
}

func (t *LingTongUpstarTemplate) GetUpdateWfb() int32 {
	return t.UpdateWfb
}

func (t *LingTongUpstarTemplate) GetAddMin() int32 {
	return t.AddMin
}

func (t *LingTongUpstarTemplate) GetAddMax() int32 {
	return t.AddMax
}

func (t *LingTongUpstarTemplate) GetTimesMin() int32 {
	return t.TimesMin
}

func (t *LingTongUpstarTemplate) GetTimesMax() int32 {
	return t.TimesMax
}

func (t *LingTongUpstarTemplate) GetZhuFuMax() int32 {
	return t.ZhufuMax
}

func (t *LingTongUpstarTemplate) GetItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *LingTongUpstarTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *LingTongUpstarTemplate) GetLevel() int32 {
	return t.Level
}

func (t *LingTongUpstarTemplate) GetNext() *LingTongUpstarTemplate {
	return t.nextTemplate
}

func (t *LingTongUpstarTemplate) Patch() (err error) {
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
	//验证 upstar_item_id
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
		to := template.GetTemplateService().Get(int(t.NextId), (*LingTongUpstarTemplate)(nil))
		if to != nil {
			nextTemp := to.(*LingTongUpstarTemplate)
			diffLevel := nextTemp.Level - t.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemp.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			t.nextTemplate = nextTemp
		}
	}

	return nil
}

func (t *LingTongUpstarTemplate) PatchAfterCheck() {

}

func (t *LingTongUpstarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 upstar_rate
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		err = template.NewTemplateFieldError("UpstarRate", err)
		return
	}

	//验证 level
	err = validator.MinValidate(float64(t.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	for itemId, _ := range t.needItemMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
		// itemTemplate := to.(*ItemTemplate)

		// if itemTemplate.GetItemType() != itemtypes.ItemTypeLingTong {
		// 	err = fmt.Errorf("UpstarItemId [%d]  invalid", t.UseItem)
		// 	return template.NewTemplateFieldError("UseItem", err)
		// }
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

	return nil
}

func (t *LingTongUpstarTemplate) FileName() string {
	return "tb_lingtong_upstar.json"
}

func init() {
	template.Register((*LingTongUpstarTemplate)(nil))
}
