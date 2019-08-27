package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/lingtongdev/types"
	propertytypes "fgame/fgame/game/property/types"

	"fmt"
)

//灵翼培养配置
type LingTongWingPeiYangTemplate struct {
	*LingTongWingPeiYangTemplateVO
	useItemMap                map[int32]int32 //培养物品
	culItemTemplate           *ItemTemplate
	battlePropertyMap         map[propertytypes.BattlePropertyType]int64
	lingTongBattlePropertyMap map[propertytypes.BattlePropertyType]int64
	next                      LingTongDevPeiYangTemplate
}

func (t *LingTongWingPeiYangTemplate) TemplateId() int {
	return t.Id
}

func (t *LingTongWingPeiYangTemplate) GetNextId() int32 {
	return t.NextId
}

func (t *LingTongWingPeiYangTemplate) GetUpdateWfb() int32 {
	return t.UpdateWfb
}

func (t *LingTongWingPeiYangTemplate) GetAddMin() int32 {
	return t.AddMin
}

func (t *LingTongWingPeiYangTemplate) GetAddMax() int32 {
	return t.AddMax
}

func (t *LingTongWingPeiYangTemplate) GetTimesMin() int32 {
	return t.TimesMin
}

func (t *LingTongWingPeiYangTemplate) GetTimesMax() int32 {
	return t.TimesMax
}

func (t *LingTongWingPeiYangTemplate) GetZhuFuMax() int32 {
	return t.ZhufuMax
}

func (t *LingTongWingPeiYangTemplate) GetItemMap() map[int32]int32 {
	return t.useItemMap
}

func (t *LingTongWingPeiYangTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *LingTongWingPeiYangTemplate) GetLingTongBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.lingTongBattlePropertyMap
}

func (t *LingTongWingPeiYangTemplate) GetLevel() int32 {
	return t.Level
}

func (t *LingTongWingPeiYangTemplate) GetNext() LingTongDevPeiYangTemplate {
	return t.next
}

func (t *LingTongWingPeiYangTemplate) GetItemCount() int32 {
	return t.ItemCount
}

func (t *LingTongWingPeiYangTemplate) GetItemId() int32 {
	return t.UseItem
}

func (t *LingTongWingPeiYangTemplate) GetClassType() types.LingTongDevSysType {
	return types.LingTongDevSysTypeLingQi
}

func (t *LingTongWingPeiYangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.useItemMap = make(map[int32]int32)
	//验证 UseItem
	if t.UseItem != 0 {
		useItemTemplate := template.GetTemplateService().Get(int(t.UseItem), (*ItemTemplate)(nil))
		if useItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			err = template.NewTemplateFieldError("UseItem", err)
			return
		}

		t.culItemTemplate = useItemTemplate.(*ItemTemplate)

		//验证 ItemCount
		err = validator.MinValidate(float64(t.ItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			err = template.NewTemplateFieldError("ItemCount", err)
			return
		}
		t.useItemMap[t.UseItem] = t.ItemCount
	}

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

	//属性
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)

	err = validator.MinValidate(float64(t.LingTongAttack), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongAttack)
		return template.NewTemplateFieldError("LingTongAttack", err)
	}

	err = validator.MinValidate(float64(t.LingTongCritical), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongCritical)
		return template.NewTemplateFieldError("LingTongCritical", err)
	}

	err = validator.MinValidate(float64(t.LingTongHit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongHit)
		return template.NewTemplateFieldError("LingTongHit", err)
	}

	err = validator.MinValidate(float64(t.LingTongAbnormality), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LingTongAbnormality)
		return template.NewTemplateFieldError("LingTongAbnormality", err)
	}

	t.lingTongBattlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.lingTongBattlePropertyMap[propertytypes.BattlePropertyTypeAttack] = t.LingTongAttack
	t.lingTongBattlePropertyMap[propertytypes.BattlePropertyTypeCrit] = t.LingTongCritical
	t.lingTongBattlePropertyMap[propertytypes.BattlePropertyTypeHit] = t.LingTongHit
	t.lingTongBattlePropertyMap[propertytypes.BattlePropertyTypeAbnormality] = t.LingTongAbnormality

	return nil
}

func (t *LingTongWingPeiYangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*LingTongWingPeiYangTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.next = to.(*LingTongWingPeiYangTemplate)

		diffLevel := t.next.GetLevel() - t.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", t.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//验证update_wfb
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		err = template.NewTemplateFieldError("UpdateWfb", err)
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

	//验证 zhufu_max
	err = validator.MinValidate(float64(t.ZhufuMax), float64(t.AddMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ZhufuMax)
		err = template.NewTemplateFieldError("ZhufuMax", err)
		return
	}

	if t.culItemTemplate != nil {
		if t.culItemTemplate.GetItemSubType() != itemtypes.ItemLingTongWingSubTypePeiYangDan {
			err = fmt.Errorf("[%d] invalid", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
	}

	return nil
}
func (t *LingTongWingPeiYangTemplate) PatchAfterCheck() {

}
func (t *LingTongWingPeiYangTemplate) FileName() string {
	return "tb_lingtong_wing_peiyang.json"
}

func init() {
	template.Register((*LingTongWingPeiYangTemplate)(nil))
}
