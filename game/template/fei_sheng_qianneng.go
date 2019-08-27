package template

import (
	"fgame/fgame/core/template"
	feishengtypes "fgame/fgame/game/feisheng/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//飞升潜能配置
type FeiShengQianNengTemplate struct {
	*FeiShengQianNengTemplateVO
	qnType        feishengtypes.QianNengType
	battleAttrMap map[propertytypes.BattlePropertyType]int64 //飞升等级属性
}

func (t *FeiShengQianNengTemplate) TemplateId() int {
	return t.Id
}

func (t *FeiShengQianNengTemplate) GetQianNengType() feishengtypes.QianNengType {
	return t.qnType
}

func (t *FeiShengQianNengTemplate) GetBattleAttrMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrMap
}

func (t *FeiShengQianNengTemplate) PatchAfterCheck() {
}

func (t *FeiShengQianNengTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//属性
	t.battleAttrMap = make(map[propertytypes.BattlePropertyType]int64)
	if t.Hp > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	}
	if t.Attack > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	}
	if t.Defence > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	}

	return nil
}

func (t *FeiShengQianNengTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.qnType = feishengtypes.QianNengType(t.Type)
	if !t.qnType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	return nil
}

func (t *FeiShengQianNengTemplate) FileName() string {
	return "tb_feisheng_qianneng.json"
}

func init() {
	template.Register((*FeiShengQianNengTemplate)(nil))
}
