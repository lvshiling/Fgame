package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/property/types"
	"fmt"
)

//食丹等级配置
type EatDanTemplate struct {
	*EatDanTemplateVO
	eatDanMap map[int]int32
	//整级的属性
	battlePropertyMap map[types.BattlePropertyType]int64
}

func (edt *EatDanTemplate) TemplateId() int {
	return edt.Id
}

func (edt *EatDanTemplate) GetAllEatDan() map[int]int32 {
	return edt.eatDanMap
}

func (edt *EatDanTemplate) GetBattlePropertyMap() map[types.BattlePropertyType]int64 {
	return edt.battlePropertyMap
}

func (edt *EatDanTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(edt.FileName(), edt.TemplateId(), err)
			return
		}
	}()

	edt.eatDanMap = make(map[int]int32)
	edt.eatDanMap[int(edt.DanId1)] = int32(edt.DanIdNum1)
	edt.eatDanMap[int(edt.DanId2)] = int32(edt.DanIdNum2)
	edt.eatDanMap[int(edt.DanId3)] = int32(edt.DanIdNum3)
	edt.eatDanMap[int(edt.DanId4)] = int32(edt.DanIdNum4)
	edt.eatDanMap[int(edt.DanId5)] = int32(edt.DanIdNum5)
	edt.eatDanMap[int(edt.DanId6)] = int32(edt.DanIdNum6)
	edt.eatDanMap[int(edt.DanId7)] = int32(edt.DanIdNum7)

	return nil
}

func (edt *EatDanTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(edt.FileName(), edt.TemplateId(), err)
			return
		}
	}()

	// 验证DanId 和 DanIdNum
	for danId, danIdNum := range edt.eatDanMap {
		to := template.GetTemplateService().Get(int(danId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", danId)
			return template.NewTemplateFieldError("DanId", err)
		}

		err = validator.MinValidate(float64(danIdNum), float64(1), true)
		if err != nil {
			err = template.NewTemplateFieldError("DanIdNum", err)
			return
		}
	}

	//验证 next_id
	if edt.NextId != 0 {
		diff := edt.NextId - int32(edt.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", edt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(edt.NextId), (*EatDanTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", edt.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
	}

	return nil
}

func (edt *EatDanTemplate) PatchAfterCheck() {
	edt.battlePropertyMap = make(map[types.BattlePropertyType]int64)
	for itemId, num := range edt.eatDanMap {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		itemTemplate, _ := to.(*ItemTemplate)
		attrTemplate := itemTemplate.GetDanAttrTemplate()
		for typ, val := range attrTemplate.GetAllBattleProperty() {
			total, ok := edt.battlePropertyMap[typ]
			if !ok {
				total = 0
			}
			total += val * int64(num)
			edt.battlePropertyMap[typ] = total
		}
	}
}

func (edt *EatDanTemplate) FileName() string {
	return "tb_eat_dan.json"
}

func init() {
	template.Register((*EatDanTemplate)(nil))
}
