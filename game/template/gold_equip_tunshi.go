package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//元神等级配置
type GoldYuanTemplate struct {
	*GoldYuanTemplateVO
	//属性id
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	//下级模板
	nextGoldYuanTemplate *GoldYuanTemplate
}

func (est *GoldYuanTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return est.battlePropertyMap
}

func (est *GoldYuanTemplate) GetNextGoldYuanTemplate() *GoldYuanTemplate {
	return est.nextGoldYuanTemplate
}

func (est *GoldYuanTemplate) TemplateId() int {
	return est.Id
}

func (t *GoldYuanTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	//属性
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.AddHp)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.AddAttack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.AddDefect)

	if t.NextId != 0 {
		//下一阶装备
		tempNextGoldYuanTemplate := template.GetTemplateService().Get(int(t.NextId), (*GoldYuanTemplate)(nil))
		if tempNextGoldYuanTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}
		t.nextGoldYuanTemplate = tempNextGoldYuanTemplate.(*GoldYuanTemplate)
	}

	return nil
}
func (t *GoldYuanTemplate) PatchAfterCheck() {

}
func (t *GoldYuanTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	for typ, val := range t.battlePropertyMap {
		//验证数量至少0
		err = validator.MinValidate(float64(val), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", val)
			err = template.NewTemplateFieldError(typ.String(), err)
			return
		}
	}

	// 经验
	err = validator.MinValidate(float64(t.Exp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Exp)
		err = template.NewTemplateFieldError("Exp", err)
		return
	}

	return nil
}

func (edt *GoldYuanTemplate) FileName() string {
	return "tb_goldequip_tunshi.json"
}

func init() {
	template.Register((*GoldYuanTemplate)(nil))
}
