package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

type YingLingPuSuitTemplate struct {
	*YingLingPuSuitTemplateVO
	battleAttrMap        map[propertytypes.BattlePropertyType]int64
	battleAttrPercentMap map[propertytypes.BattlePropertyType]int64
	suitConditionIdList  []int32
}

func (t *YingLingPuSuitTemplate) TemplateId() int {
	return t.Id
}

func (t *YingLingPuSuitTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrMap
}

func (t *YingLingPuSuitTemplate) GetBattlePropertyPercentMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrPercentMap
}

func (t *YingLingPuSuitTemplate) GetSuitCondition() []int32 {
	return t.suitConditionIdList
}

func (t *YingLingPuSuitTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//套装条件
	t.suitConditionIdList, err = utils.SplitAsIntArray(t.YinLingPuId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.YinLingPuId)
		err = template.NewTemplateFieldError("YinLingPuId", err)
		return
	}

	// 套装属性
	t.battleAttrMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battleAttrMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battleAttrMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battleAttrMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)

	// 套装属性百分比
	t.battleAttrPercentMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battleAttrPercentMap[propertytypes.BattlePropertyTypeAttack] = int64(t.AttackPercent)
	t.battleAttrPercentMap[propertytypes.BattlePropertyTypeDefend] = int64(t.DefPercent)
	t.battleAttrPercentMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.HpPercent)

	return nil
}

func (t *YingLingPuSuitTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//属性
	for typ, val := range t.battleAttrMap {
		err = validator.MinValidate(float64(val), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", val)
			err = template.NewTemplateFieldError(fmt.Sprintf("%s", typ.String()), err)
			return
		}
	}

	err = validator.MinValidate(float64(t.HpPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.HpPercent)
		err = template.NewTemplateFieldError("HpPercent", err)
		return
	}
	err = validator.MinValidate(float64(t.AttackPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.AttackPercent)
		err = template.NewTemplateFieldError("AttackPercent", err)
		return
	}
	err = validator.MinValidate(float64(t.DefPercent), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DefPercent)
		err = template.NewTemplateFieldError("DefPercent", err)
		return
	}

	return nil
}

func (t *YingLingPuSuitTemplate) PatchAfterCheck() {
	return
}

func (t *YingLingPuSuitTemplate) FileName() string {
	return "tb_yinglingpu_taozhuang.json"
}

func init() {
	template.Register((*YingLingPuSuitTemplate)(nil))
}
