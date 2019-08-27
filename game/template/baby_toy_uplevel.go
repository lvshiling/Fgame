package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	babytypes "fgame/fgame/game/baby/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//玩具强化配置
type BabyToyUplevelTemplate struct {
	*BabyToyUplevelTemplateVO
	suitType                   babytypes.ToySuitType
	posType                    inventorytypes.BodyPositionType
	nextBabyToyUplevelTemplate *BabyToyUplevelTemplate //下一级强化
	failReturnTemplate         *BabyToyUplevelTemplate
	battlePropertyMap          map[propertytypes.BattlePropertyType]int64 //属性
}

func (t *BabyToyUplevelTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyToyUplevelTemplate) GetSuitType() babytypes.ToySuitType {
	return t.suitType
}

func (t *BabyToyUplevelTemplate) GetPosType() inventorytypes.BodyPositionType {
	return t.posType
}

func (t *BabyToyUplevelTemplate) GetNextTemplate() *BabyToyUplevelTemplate {
	return t.nextBabyToyUplevelTemplate
}

func (t *BabyToyUplevelTemplate) GetFaildReturnTemplate() *BabyToyUplevelTemplate {
	return t.failReturnTemplate
}

func (t *BabyToyUplevelTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *BabyToyUplevelTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//下一阶强化
	if t.NextId != 0 {
		if t.NextId-int32(t.Id) != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}

		tempNextBabyToyUplevelTemplate := template.GetTemplateService().Get(int(t.NextId), (*BabyToyUplevelTemplate)(nil))
		if tempNextBabyToyUplevelTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("next", err)
		}
		t.nextBabyToyUplevelTemplate = tempNextBabyToyUplevelTemplate.(*BabyToyUplevelTemplate)
	}

	// 
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)

	return nil
}
func (t *BabyToyUplevelTemplate) PatchAfterCheck() {

}
func (t *BabyToyUplevelTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	t.suitType = babytypes.ToySuitType(t.SuitType)
	if !t.suitType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.SuitType)
		return template.NewTemplateFieldError("SuitType", err)
	}

	t.posType = inventorytypes.BodyPositionType(t.Position)
	if !t.posType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Position)
		return template.NewTemplateFieldError("Position", err)
	}

	//验证 等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 概率
	err = validator.MinValidate(float64(t.Rate), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Rate)
		return template.NewTemplateFieldError("Rate", err)
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

	//
	if t.NeedItem > 0 || t.ItemCount > 0 {
		to := template.GetTemplateService().Get(int(t.NeedItem), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NeedItem)
			err = template.NewTemplateFieldError("NeedItem", err)
			return
		}
		err = validator.MinValidate(float64(t.ItemCount), float64(0), false)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.ItemCount)
			return template.NewTemplateFieldError("ItemCount", err)
		}
	}

	//失败回退等级
	if t.FailReturnStrengthenId != 0 {
		//失败
		to := template.GetTemplateService().Get(int(t.FailReturnStrengthenId), (*BabyToyUplevelTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.FailReturnStrengthenId)
			return template.NewTemplateFieldError("FailReturnStrengthenId", err)
		}
		t.failReturnTemplate = to.(*BabyToyUplevelTemplate)
	}

	return nil
}

func (t *BabyToyUplevelTemplate) FileName() string {
	return "tb_baobao_wanju_level.json"
}

func init() {
	template.Register((*BabyToyUplevelTemplate)(nil))
}
