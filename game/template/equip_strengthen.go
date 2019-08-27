package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	inventorytypes "fgame/fgame/game/inventory/types"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//装备强化配置
type EquipStrengthenTemplate struct {
	*EquipStrengthenTemplateVO
	//强化类型
	equipStrengthType inventorytypes.EquipmentStrengthenType
	//位置
	position inventorytypes.BodyPositionType
	//属性id
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	//下一阶升装备
	nextEquipStrengthenTemplate *EquipStrengthenTemplate
	//回退等级
	failEquipStrengthenTemplate *EquipStrengthenTemplate
	//进阶所需物品
	needItemMap map[int32]int32
}

func (est *EquipStrengthenTemplate) GetEquipStrengthType() inventorytypes.EquipmentStrengthenType {
	return est.equipStrengthType
}

func (est *EquipStrengthenTemplate) GetPosition() inventorytypes.BodyPositionType {
	return est.position
}

func (est *EquipStrengthenTemplate) GetBattlePropertyMap() map[propertytypes.BattlePropertyType]int64 {
	return est.battlePropertyMap
}

func (est *EquipStrengthenTemplate) GetNextEquipStrengthenTemplate() *EquipStrengthenTemplate {
	return est.nextEquipStrengthenTemplate
}

func (est *EquipStrengthenTemplate) GetFailEquipStrengthenTemplate() *EquipStrengthenTemplate {
	return est.failEquipStrengthenTemplate
}

func (est *EquipStrengthenTemplate) GetNeedItemMap() map[int32]int32 {
	return est.needItemMap
}

//判断是否有条件
func (est *EquipStrengthenTemplate) HasCondition() bool {
	if len(est.needItemMap) != 0 {
		return true
	}
	if est.SilverNum > 0 {
		return true
	}
	return false
}

func (est *EquipStrengthenTemplate) TemplateId() int {
	return est.Id
}

func (et *EquipStrengthenTemplate) Patch() (err error) {
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
	if et.NeedItem != 0 {
		//下一阶装备
		tempNeedItemTemplate := template.GetTemplateService().Get(int(et.NeedItem), (*ItemTemplate)(nil))
		if tempNeedItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NeedItem)
			return template.NewTemplateFieldError("needItem", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(et.NeedItemNum), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", et.NeedItemNum)
			err = template.NewTemplateFieldError("needItemNum", err)
			return
		}
		needItemTemplate := tempNeedItemTemplate.(*ItemTemplate)
		et.needItemMap[int32(needItemTemplate.TemplateId())] = et.NeedItemNum
	}

	if et.NextId != 0 {
		//下一阶装备
		tempNextEquipStrengthenTemplate := template.GetTemplateService().Get(int(et.NextId), (*EquipStrengthenTemplate)(nil))
		if tempNextEquipStrengthenTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NextId)
			return template.NewTemplateFieldError("next", err)
		}
		et.nextEquipStrengthenTemplate = tempNextEquipStrengthenTemplate.(*EquipStrengthenTemplate)
	}

	if et.FailReturnStrengthenId != 0 {
		//失败
		tempFailReturnEquipStrengthenTemplate := template.GetTemplateService().Get(int(et.FailReturnStrengthenId), (*EquipStrengthenTemplate)(nil))
		if tempFailReturnEquipStrengthenTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.FailReturnStrengthenId)
			return template.NewTemplateFieldError("fail_return_strengthen_id", err)
		}
		et.failEquipStrengthenTemplate = tempFailReturnEquipStrengthenTemplate.(*EquipStrengthenTemplate)
	}
	//位置
	et.position = inventorytypes.BodyPositionType(et.Position)
	if !et.position.Valid() {
		err = fmt.Errorf("[%d] invalid", et.Position)
		err = template.NewTemplateFieldError("position", err)
		return
	}
	//强化类型
	et.equipStrengthType = inventorytypes.EquipmentStrengthenType(et.Type)
	if !et.equipStrengthType.Valid() {
		err = fmt.Errorf("[%d] invalid", et.Type)
		err = template.NewTemplateFieldError("type", err)
		return
	}
	return nil
}
func (et *EquipStrengthenTemplate) PatchAfterCheck() {

}
func (et *EquipStrengthenTemplate) Check() (err error) {
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

	//可以升阶
	if et.nextEquipStrengthenTemplate != nil {
		//成功概率1-10000
		err = validator.RangeValidate(float64(et.SuccessRate), float64(1), true, float64(common.MAX_RATE), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", et.SuccessRate)
			err = template.NewTemplateFieldError("successRate", err)
			return
		}
		if !et.nextEquipStrengthenTemplate.HasCondition() {
			err = fmt.Errorf("可以升阶但是没有条件")
			return
		}
	}

	if et.failEquipStrengthenTemplate != nil {
		//失败概率1-10000
		err = validator.RangeValidate(float64(et.ReturnRate), float64(0), true, float64(common.MAX_RATE), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", et.ReturnRate)
			err = template.NewTemplateFieldError("returnRate", err)
			return
		}
	}

	return nil
}

func (edt *EquipStrengthenTemplate) FileName() string {
	return "tb_equip_strengthen.json"
}

func init() {
	template.Register((*EquipStrengthenTemplate)(nil))
}
