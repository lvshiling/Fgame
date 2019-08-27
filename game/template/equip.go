package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//装备配置
type EquipTemplate struct {
	*EquipTemplateVO
	//属性id
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	//下一阶装备
	nextItemTemplate *ItemTemplate
	//进阶所需物品
	needItemMap map[int32]int32
	//套装
	taozhuangTemplate *TaozhuangTemplate
}

func (et *EquipTemplate) TemplateId() int {
	return et.Id
}

func (et *EquipTemplate) GetTaozhuangTemplate() *TaozhuangTemplate {
	return et.taozhuangTemplate
}

func (et *EquipTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return et.battlePropertyMap
}

func (et *EquipTemplate) GetNextItemTemplate() *ItemTemplate {
	return et.nextItemTemplate
}

func (et *EquipTemplate) GetNeedItemMap() map[int32]int32 {
	return et.needItemMap
}

func (et *EquipTemplate) HasCondition() bool {
	return len(et.needItemMap) != 0
}

func (et *EquipTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()
	et.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	et.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = et.Hp
	et.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = et.Attack
	et.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = et.Defence

	et.needItemMap = make(map[int32]int32)
	//验证升级材料
	if et.NeedItem1 != 0 {
		//下一阶装备
		tempNeedItemTemplate := template.GetTemplateService().Get(int(et.NeedItem1), (*ItemTemplate)(nil))
		if tempNeedItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NeedItem1)
			return template.NewTemplateFieldError("needItem1", err)
		}

		needItemTemplate := tempNeedItemTemplate.(*ItemTemplate)
		et.needItemMap[int32(needItemTemplate.TemplateId())] = et.NeedItemNum1
	}
	if et.NeedItem2 != 0 {
		//下一阶装备
		tempNeedItemTemplate := template.GetTemplateService().Get(int(et.NeedItem2), (*ItemTemplate)(nil))
		if tempNeedItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NeedItem2)
			return template.NewTemplateFieldError("needItem2", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(et.NeedItemNum2), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", et.NeedItemNum2)
			err = template.NewTemplateFieldError("needItemNum2", err)
			return
		}
		needItemTemplate := tempNeedItemTemplate.(*ItemTemplate)
		et.needItemMap[int32(needItemTemplate.TemplateId())] = et.NeedItemNum2
	}
	if et.NeedItem3 != 0 {
		//下一阶装备
		tempNeedItemTemplate := template.GetTemplateService().Get(int(et.NeedItem3), (*ItemTemplate)(nil))
		if tempNeedItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NeedItem3)
			return template.NewTemplateFieldError("needItem3", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(et.NeedItemNum3), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", et.NeedItemNum3)
			err = template.NewTemplateFieldError("needItemNum3", err)
			return
		}
		needItemTemplate := tempNeedItemTemplate.(*ItemTemplate)
		et.needItemMap[int32(needItemTemplate.TemplateId())] = et.NeedItemNum3
	}

	if et.TaozhuangId != 0 {
		tempTaozhuangTemplate := template.GetTemplateService().Get(int(et.TaozhuangId), (*TaozhuangTemplate)(nil))
		if tempTaozhuangTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.TaozhuangId)
			err = template.NewTemplateFieldError("taozhuangId", err)
			return
		}
		et.taozhuangTemplate = tempTaozhuangTemplate.(*TaozhuangTemplate)
	}
	return nil
}

func (et *EquipTemplate) PatchAfterCheck() {

}
func (et *EquipTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()
	for typ, val := range et.battlePropertyMap {
		//验证数量至少0
		err = validator.MinValidate(float64(val), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", val)
			err = template.NewTemplateFieldError(typ.String(), err)
			return
		}
	}
	if et.Next != 0 {
		//下一阶装备
		tempNextItemTemplate := template.GetTemplateService().Get(int(et.Next), (*ItemTemplate)(nil))
		if tempNextItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.Next)
			return template.NewTemplateFieldError("next", err)
		}
		et.nextItemTemplate = tempNextItemTemplate.(*ItemTemplate)
		equipmentTemplate := et.nextItemTemplate.GetEquipmentTemplate()
		if equipmentTemplate == nil {
			err = fmt.Errorf("[%d] 不是装备", et.Next)
			return template.NewTemplateFieldError("next", err)
		}
		if !equipmentTemplate.HasCondition() {
			err = fmt.Errorf("[%d] 没有条件", et.Next)
			return template.NewTemplateFieldError("next", err)
		}
		//成功概率1-10000
		err = validator.RangeValidate(float64(et.SuccessRate), float64(1), true, float64(common.MAX_RATE), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", et.SuccessRate)
			err = template.NewTemplateFieldError("successRate", err)
			return
		}
	}

	return nil
}

func (edt *EquipTemplate) FileName() string {
	return "tb_equip.json"
}

func init() {
	template.Register((*EquipTemplate)(nil))
}
