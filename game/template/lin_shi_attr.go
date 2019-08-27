package template

import (
	"fgame/fgame/core/template"
	propertytypes "fgame/fgame/game/property/types"
)

//临时属性
type LinShiAttrTemplate struct {
	*LinShiAttrTemplateVO

	battlePropertyMap        map[propertytypes.BattlePropertyType]int64
	battlePropertyPercentMap map[propertytypes.BattlePropertyType]int64
}

func (t *LinShiAttrTemplate) TemplateId() int {
	return t.Id
}

func (t *LinShiAttrTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *LinShiAttrTemplate) GetBattlePropertyPercentMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyPercentMap
}

func (t *LinShiAttrTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)

	t.battlePropertyPercentMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.HpPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeAttack] = int64(t.AttackPercent)
	t.battlePropertyPercentMap[propertytypes.BattlePropertyTypeDefend] = int64(t.DefPercent)

	return nil
}

func (t *LinShiAttrTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}
func (t *LinShiAttrTemplate) PatchAfterCheck() {

}
func (t *LinShiAttrTemplate) FileName() string {
	return "tb_linshi_attr.json"
}

func init() {
	template.Register((*LinShiAttrTemplate)(nil))
}
