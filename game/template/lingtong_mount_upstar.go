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

//灵骑升星配置
type LingTongMountUpstarTemplate struct {
	*LingTongMountUpstarTemplateVO
	needItemMap               map[int32]int32 //升星需要物品
	useItemTemplate           *ItemTemplate
	battlePropertyMap         map[propertytypes.BattlePropertyType]int64
	lingTongBattlePropertyMap map[propertytypes.BattlePropertyType]int64
	next                      LingTongDevUpstarTemplate
}

func (t *LingTongMountUpstarTemplate) TemplateId() int {
	return t.Id
}

func (t *LingTongMountUpstarTemplate) GetNextId() int32 {
	return t.NextId
}

func (t *LingTongMountUpstarTemplate) GetUpdateWfb() int32 {
	return t.UpdateWfb
}

func (t *LingTongMountUpstarTemplate) GetAddMin() int32 {
	return t.AddMin
}

func (t *LingTongMountUpstarTemplate) GetAddMax() int32 {
	return t.AddMax
}

func (t *LingTongMountUpstarTemplate) GetTimesMin() int32 {
	return t.TimesMin
}

func (t *LingTongMountUpstarTemplate) GetTimesMax() int32 {
	return t.TimesMax
}

func (t *LingTongMountUpstarTemplate) GetZhuFuMax() int32 {
	return t.ZhufuMax
}

func (t *LingTongMountUpstarTemplate) GetItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *LingTongMountUpstarTemplate) GetUpstarPercent() int32 {
	return t.Percent
}

func (t *LingTongMountUpstarTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *LingTongMountUpstarTemplate) GetLingTongBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.lingTongBattlePropertyMap
}

func (t *LingTongMountUpstarTemplate) GetLevel() int32 {
	return t.Level
}

func (t *LingTongMountUpstarTemplate) GetNext() LingTongDevUpstarTemplate {
	return t.next
}

func (t *LingTongMountUpstarTemplate) GetClassType() types.LingTongDevSysType {
	return types.LingTongDevSysTypeLingQi
}

func (t *LingTongMountUpstarTemplate) Patch() (err error) {
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
		to := template.GetTemplateService().Get(int(t.NextId), (*LingTongMountUpstarTemplate)(nil))
		if to != nil {
			nextTemplate := to.(*LingTongMountUpstarTemplate)
			diffLevel := nextTemplate.Level - t.Level
			if diffLevel != 1 {
				err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
				return template.NewTemplateFieldError("Level", err)
			}
			t.next = nextTemplate
		}
	}

	return nil
}

func (t *LingTongMountUpstarTemplate) PatchAfterCheck() {

}

func (t *LingTongMountUpstarTemplate) Check() (err error) {
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

	//验证 fabao_percent
	err = validator.RangeValidate(float64(t.Percent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Percent)
		err = template.NewTemplateFieldError("Percent", err)
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
		itemTemplate := to.(*ItemTemplate)

		if itemTemplate.GetItemType() != itemtypes.ItemTypeLingTongMount {
			err = fmt.Errorf("UpstarItemId [%d]  invalid", t.UseItem)
			return template.NewTemplateFieldError("UseItem", err)
		}
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

func (t *LingTongMountUpstarTemplate) FileName() string {
	return "tb_lingtong_mount_upstar.json"
}

func init() {
	template.Register((*LingTongMountUpstarTemplate)(nil))
}
