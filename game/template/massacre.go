package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//戮仙刃配置
type MassacreTemplate struct {
	*MassacreTemplateVO
	useItemTemplate   *ItemTemplate   //进阶物品
	atvWeaponTemplate *WeaponTemplate //激活兵魂
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
}

func (t *MassacreTemplate) TemplateId() int {
	return t.Id
}

func (t *MassacreTemplate) GetUseItemTemplate() *ItemTemplate {
	return t.useItemTemplate
}

func (t *MassacreTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *MassacreTemplate) GetAtvWeaponTemplate() *WeaponTemplate {
	return t.atvWeaponTemplate
}

func (t *MassacreTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//阶别兵魂
	if t.WeaponId != 0 {
		to := template.GetTemplateService().Get(int(t.WeaponId), (*WeaponTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.WeaponId)
			return template.NewTemplateFieldError("WeaponId", err)
		}
		weaponTemplate, _ := to.(*WeaponTemplate)
		t.atvWeaponTemplate = weaponTemplate
	}

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

func (t *MassacreTemplate) PatchAfterCheck() {

}
func (t *MassacreTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*MassacreTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
	}

	//验证type戮仙刃阶别
	err = validator.MinValidate(float64(t.Type), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证star戮仙刃星数
	err = validator.MinValidate(float64(t.Star), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Star)
		err = template.NewTemplateFieldError("Star", err)
		return
	}

	//验证UpdatePercent
	err = validator.RangeValidate(float64(t.UpdatePercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdatePercent)
		err = template.NewTemplateFieldError("UpdatePercent", err)
		return
	}

	//验证use_money
	err = validator.MinValidate(float64(t.UseMoney), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseMoney)
		err = template.NewTemplateFieldError("UseMoney", err)
		return
	}

	//验证UseGas
	err = validator.MinValidate(float64(t.UseGas), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseGas)
		err = template.NewTemplateFieldError("UseGas", err)
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

	//验证 GasMin
	err = validator.RangeValidate(float64(t.GasMin), float64(0), true, float64(t.GasMax), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GasMin)
		err = template.NewTemplateFieldError("GasMin", err)
		return
	}

	//验证 GasMax
	err = validator.MinValidate(float64(t.GasMax), float64(t.GasMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GasMax)
		err = template.NewTemplateFieldError("GasMax", err)
		return
	}

	//验证 GasPercent
	err = validator.RangeValidate(float64(t.GasPercent), float64(0), true, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.GasPercent)
		err = template.NewTemplateFieldError("GasPercent", err)
		return
	}

	//验证StarCount
	err = validator.MinValidate(float64(t.StarCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.StarCount)
		err = template.NewTemplateFieldError("StarCount", err)
		return
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

	return nil
}

func (t *MassacreTemplate) FileName() string {
	return "tb_massacre.json"
}

func init() {
	template.Register((*MassacreTemplate)(nil))
}
