package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	commontypes "fgame/fgame/game/common/types"
	propertytypes "fgame/fgame/game/property/types"

	"fmt"
)

//天魔配置
type TianMoTemplate struct {
	*TianMoTemplateVO
	activateType      commontypes.SpecialAdvancedType //天魔激活类型
	activateUniteType commontypes.AdvancedUnitePiFuType
	useItemTemplate   *ItemTemplate //进阶物品
	//属性
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
}

func (t *TianMoTemplate) TemplateId() int {
	return t.Id
}

func (t *TianMoTemplate) GetActivateType() commontypes.SpecialAdvancedType {
	return t.activateType
}

func (t *TianMoTemplate) GetActivateUniteType() commontypes.AdvancedUnitePiFuType {
	return t.activateUniteType
}

func (t *TianMoTemplate) GetUseItemTemplate() *ItemTemplate {
	return t.useItemTemplate
}

func (t *TianMoTemplate) GetBattleAttrMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *TianMoTemplate) GetIsClear() bool {
	return t.IsClear != 0
}

func (t *TianMoTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//天魔类型
	t.activateType = commontypes.SpecialAdvancedType(t.ShengjieType)
	if !t.activateType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.ShengjieType)
		return template.NewTemplateFieldError("ShengjieType", err)
	}

	// 关联皮肤
	t.activateUniteType = commontypes.AdvancedUnitePiFuType(t.WaiguanType)
	if !t.activateUniteType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.WaiguanType)
		err = template.NewTemplateFieldError("WaiguanType", err)
		return
	}

	// 属性
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)

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

	return nil
}

func (t *TianMoTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if t.NextId != 0 {

		to := template.GetTemplateService().Get(int(t.NextId), (*TianMoTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		nextTemp := to.(*TianMoTemplate)

		diff := nextTemp.Number - int32(t.Number)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	//验证Number
	err = validator.MinValidate(float64(t.Number), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
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
	err = validator.MinValidate(float64(t.ShidanLimit), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ShidanLimit)
		err = template.NewTemplateFieldError("ShidanLimit", err)
		return
	}

	err = validator.MinValidate(float64(t.IsClear), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.IsClear)
		return template.NewTemplateFieldError("IsClear", err)
	}

	return nil
}

func (t *TianMoTemplate) PatchAfterCheck() {
}

func (t *TianMoTemplate) FileName() string {
	return "tb_tianmoti.json"
}

func init() {
	template.Register((*TianMoTemplate)(nil))
}
