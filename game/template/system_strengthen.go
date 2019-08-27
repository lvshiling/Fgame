package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//装备强化配置
type SystemStrengthenTemplate struct {
	*SystemStrengthenTemplateVO
	nextSystemStrengthenTemplate *SystemStrengthenTemplate //下一级强化
	failSystemStrengthenTemplate *SystemStrengthenTemplate //回退等级
	needItemMap                  map[int32]int32
	//属性
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
}

func (est *SystemStrengthenTemplate) TemplateId() int {
	return est.Id
}

func (est *SystemStrengthenTemplate) GetNextTemplate() *SystemStrengthenTemplate {
	return est.nextSystemStrengthenTemplate
}

func (est *SystemStrengthenTemplate) GetFailTemplate() *SystemStrengthenTemplate {
	return est.failSystemStrengthenTemplate
}

func (est *SystemStrengthenTemplate) GetNeedItemMap() map[int32]int32 {
	tempMap := make(map[int32]int32)
	for itemId, num := range est.needItemMap {
		_, ok := tempMap[itemId]
		if ok {
			tempMap[itemId] += num
		} else {
			tempMap[itemId] = num
		}
	}
	return tempMap
}

func (est *SystemStrengthenTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return est.battlePropertyMap
}

func (et *SystemStrengthenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	//下一阶强化
	if et.NextId != 0 {
		tempNextSystemStrengthenTemplate := template.GetTemplateService().Get(int(et.NextId), (*SystemStrengthenTemplate)(nil))
		if tempNextSystemStrengthenTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		et.nextSystemStrengthenTemplate = tempNextSystemStrengthenTemplate.(*SystemStrengthenTemplate)
	}

	//回退强化
	if et.FailBacklevel != 0 {
		tempFailSystemStrengthenTemplate := template.GetTemplateService().Get(int(et.FailBacklevel), (*SystemStrengthenTemplate)(nil))
		if tempFailSystemStrengthenTemplate == nil {
			err = fmt.Errorf("[%d] invalid", et.FailBacklevel)
			return template.NewTemplateFieldError("FailBacklevel", err)
		}
		et.failSystemStrengthenTemplate = tempFailSystemStrengthenTemplate.(*SystemStrengthenTemplate)
	}

	//需要物品
	et.needItemMap = make(map[int32]int32)
	needItemIdList, err := utils.SplitAsIntArray(et.CostItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", et.CostItemId)
		return template.NewTemplateFieldError("CostItemId", err)
	}
	needItemCountList, err := utils.SplitAsIntArray(et.CostItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", et.CostItemCount)
		return template.NewTemplateFieldError("CostItemCount", err)
	}
	if len(needItemIdList) != len(needItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", et.CostItemId, et.CostItemCount)
		return template.NewTemplateFieldError("CostItemId or CostItemCount", err)
	}
	if len(needItemIdList) > 0 {
		//组合数据
		for index, itemId := range needItemIdList {
			_, ok := et.needItemMap[itemId]
			if ok {
				et.needItemMap[itemId] += needItemCountList[index]
			} else {
				et.needItemMap[itemId] = needItemCountList[index]
			}
		}
	}

	//属性
	et.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	et.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(et.Hp)
	et.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(et.Attack)
	et.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(et.Defence)

	return nil
}
func (et *SystemStrengthenTemplate) PatchAfterCheck() {

}
func (et *SystemStrengthenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(et.FileName(), et.TemplateId(), err)
			return
		}
	}()

	// 消耗物品id
	for itemId, num := range et.needItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("CostItemId", fmt.Errorf("[%d] invalid", itemId))
		}
		if err = validator.MinValidate(float64(num), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("CostItemCount", err)
			return
		}
	}

	//验证 ProtectItemId
	if et.ProtectItemId != 0 {
		protectItemTemplateVO := template.GetTemplateService().Get(int(et.ProtectItemId), (*ItemTemplate)(nil))
		if protectItemTemplateVO == nil {
			err = fmt.Errorf("[%d] invalid", et.ProtectItemId)
			err = template.NewTemplateFieldError("ProtectItemId", err)
			return
		}

		//验证 ProtectItemCount
		err = validator.MinValidate(float64(et.ProtectItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", et.ProtectItemCount)
			err = template.NewTemplateFieldError("ProtectItemCount", err)
			return
		}
	}

	return nil
}

func (edt *SystemStrengthenTemplate) FileName() string {
	return "tb_system_strengthen.json"
}

func init() {
	template.Register((*SystemStrengthenTemplate)(nil))
}
