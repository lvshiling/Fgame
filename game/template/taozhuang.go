package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//套装配置
type TaozhuangTemplate struct {
	*TaozhuangTemplateVO
	attrTemplate *AttrTemplate //属性

	battlePropertyMap        map[propertytypes.BattlePropertyType]int64
	battlePropertyPercentMap map[propertytypes.BattlePropertyType]int64
}

func (tt *TaozhuangTemplate) TemplateId() int {
	return tt.Id
}

func (tt *TaozhuangTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return tt.battlePropertyMap
}

func (tt *TaozhuangTemplate) GetBattlePropertyPercentMap() map[propertytypes.BattlePropertyType]int64 {
	return tt.battlePropertyPercentMap
}

func (tt *TaozhuangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()
	tt.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	tt.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = tt.Hp
	tt.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = tt.Attack
	tt.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = tt.Defence
	tt.battlePropertyPercentMap = make(map[propertytypes.BattlePropertyType]int64)
	tt.battlePropertyPercentMap[propertytypes.BattlePropertyTypeMaxHP] = tt.HpPercent
	tt.battlePropertyPercentMap[propertytypes.BattlePropertyTypeAttack] = tt.AttPercent
	tt.battlePropertyPercentMap[propertytypes.BattlePropertyTypeDefend] = tt.DefPercent

	return nil
}

func (tt *TaozhuangTemplate) PatchAfterCheck() {

}

func (tt *TaozhuangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(tt.FileName(), tt.TemplateId(), err)
			return
		}
	}()
	for typ, val := range tt.battlePropertyMap {
		//验证数量至少0
		err = validator.MinValidate(float64(val), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", val)
			err = template.NewTemplateFieldError(typ.String(), err)
			return
		}
	}
	for typ, val := range tt.battlePropertyPercentMap {
		//验证数量至少0
		err = validator.MinValidate(float64(val), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", val)
			err = template.NewTemplateFieldError(typ.String(), err)
			return
		}
	}
	//套装数量
	if err = validator.MinValidate(float64(tt.Number), float64(0), false); err != nil {
		err = template.NewTemplateFieldError("number", err)
		return
	}
	return nil
}

func (tt *TaozhuangTemplate) FileName() string {
	return "tb_taozhuang.json"
}

func init() {
	template.Register((*TaozhuangTemplate)(nil))
}
