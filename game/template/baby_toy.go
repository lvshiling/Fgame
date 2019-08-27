package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	inventorytypes "fgame/fgame/game/inventory/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//玩具配置
type BabyToyTemplate struct {
	*BabyToyTemplateVO
	//套装
	tempTaozhuangTemplate *BabyToySuitGroupTemplate
	//属性
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	//部位
	posType inventorytypes.BodyPositionType
}

func (t *BabyToyTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyToyTemplate) GetPosType() inventorytypes.BodyPositionType {
	return t.posType
}

func (t *BabyToyTemplate) GetTaozhuangTemplate() *BabyToySuitGroupTemplate {
	return t.tempTaozhuangTemplate
}

func (t *BabyToyTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

//获取装备技能
func (m *BabyToyTemplate) GetBabyToyGroupSuitSkill(equipNum int32) (skillList []int32) {
	groupSuitTemplate := m.tempTaozhuangTemplate
	if groupSuitTemplate == nil {
		return
	}

	return groupSuitTemplate.GetSuitEffectSkillId(equipNum)
}

func (t *BabyToyTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//套装
	if t.SuitGroup != 0 {
		tempTaozhuangTemplate := template.GetTemplateService().Get(int(t.SuitGroup), (*BabyToySuitGroupTemplate)(nil))
		if tempTaozhuangTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.SuitGroup)
			err = template.NewTemplateFieldError("SuitGroup", err)
			return
		}
		t.tempTaozhuangTemplate = tempTaozhuangTemplate.(*BabyToySuitGroupTemplate)
	}

	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	return nil
}

func (t *BabyToyTemplate) PatchAfterCheck() {}

func (t *BabyToyTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

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

	//部位
	t.posType = inventorytypes.BodyPositionType(t.PosType)
	if !t.posType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.PosType)
		return template.NewTemplateFieldError("PosType", err)
	}

	return nil
}

func (edt *BabyToyTemplate) FileName() string {
	return "tb_baobao_wanju.json"
}

func init() {
	template.Register((*BabyToyTemplate)(nil))
}
