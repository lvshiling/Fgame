package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//装备配置
type SystemEquipShengHenTemplate struct {
	*SystemEquipShengHenTemplateVO
	// //一级强化模板
	// minSystemStrengthenTemplate *SystemStrengthenTemplate
	// //强化模板Map
	// strengthenTemplateMap map[int32]*SystemStrengthenTemplate
	//套装
	tempTaozhuangTemplate *SystemTaozhuangTemplate
	//属性
	battlePropertyMap map[propertytypes.BattlePropertyType]int64
	//需要的进阶物品
	needItemMap map[int32]int32
	//分解返还物品
	returnItemMap    map[int32]int32
	nextItemTemplate *ItemTemplate
}

func (t *SystemEquipShengHenTemplate) TemplateId() int {
	return t.Id
}

func (t *SystemEquipShengHenTemplate) GetTaozhuangTemplate() *SystemTaozhuangTemplate {
	return t.tempTaozhuangTemplate
}

func (t *SystemEquipShengHenTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *SystemEquipShengHenTemplate) GetReturnItemMap() map[int32]int32 {
	return t.returnItemMap
}

func (t *SystemEquipShengHenTemplate) HasCondition() bool {
	return len(t.needItemMap) != 0
}

func (t *SystemEquipShengHenTemplate) GetNextItemTemplate() *ItemTemplate {
	return t.nextItemTemplate
}

// func (t *SystemEquipShengHenTemplate) GetStrengthenTemplate(level int32) *SystemStrengthenTemplate {
// 	return t.strengthenTemplateMap[level]
// }

func (t *SystemEquipShengHenTemplate) GetTushiExp() int32 {
	return t.TushiExp
}

func (t *SystemEquipShengHenTemplate) GetHp() int32 {
	return t.Hp
}

func (t *SystemEquipShengHenTemplate) GetAttack() int32 {
	return t.Attack
}

func (t *SystemEquipShengHenTemplate) GetDefence() int32 {
	return t.Defence
}

func (t *SystemEquipShengHenTemplate) GetSuccessRate() int32 {
	return t.SuccessRate
}

func (t *SystemEquipShengHenTemplate) GetBattleProperty() map[propertytypes.BattlePropertyType]int64 {
	return t.battlePropertyMap
}

func (t *SystemEquipShengHenTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//套装
	if t.SuitGroup != 0 {
		tempTaozhuangTemplate := template.GetTemplateService().Get(int(t.SuitGroup), (*SystemTaozhuangTemplate)(nil))
		if tempTaozhuangTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.SuitGroup)
			err = template.NewTemplateFieldError("SuitGroup", err)
			return
		}
		t.tempTaozhuangTemplate = tempTaozhuangTemplate.(*SystemTaozhuangTemplate)
	}

	//属性
	t.battlePropertyMap = make(map[propertytypes.BattlePropertyType]int64)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.Hp)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = int64(t.Attack)
	t.battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = int64(t.Defence)

	itemArr, err := coreutils.SplitAsIntArray(t.NeedItem)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedItem)
		return template.NewTemplateFieldError("NeedItem", err)
	}
	numArr, err := coreutils.SplitAsIntArray(t.NeedItemNum)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedItemNum)
		return template.NewTemplateFieldError("NeedItemNum", err)
	}
	if len(itemArr) != len(numArr) {
		err = fmt.Errorf("NeedItem[%s]NeedItemNum[%s]长度不相等", t.NeedItem, t.NeedItemNum)
		return template.NewTemplateFieldError("NeedItemAmount", err)
	}

	t.needItemMap = make(map[int32]int32)
	for i := 0; i < len(itemArr); i++ {
		t.needItemMap[itemArr[i]] = numArr[i]
	}

	//分解返还物品
	returnItemArr, err := coreutils.SplitAsIntArray(t.MeltingReturnId)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MeltingReturnId)
		return template.NewTemplateFieldError("MeltingReturnId", err)
	}
	returnNumArr, err := coreutils.SplitAsIntArray(t.MeltingReturnCount)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MeltingReturnCount)
		return template.NewTemplateFieldError("MeltingReturnCount", err)
	}
	if len(itemArr) != len(numArr) {
		err = fmt.Errorf("MeltingReturnId[%s]MeltingReturnCount[%s]长度不相等", t.MeltingReturnId, t.MeltingReturnCount)
		return template.NewTemplateFieldError("MeltingReturn", err)
	}

	t.returnItemMap = make(map[int32]int32)
	for i := 0; i < len(returnItemArr); i++ {
		t.returnItemMap[returnItemArr[i]] = returnNumArr[i]
	}

	if t.Next != 0 {
		//下一阶装备
		tempNextItemTemplate := template.GetTemplateService().Get(int(t.Next), (*ItemTemplate)(nil))
		if tempNextItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.Next)
			return template.NewTemplateFieldError("next", err)
		}
		t.nextItemTemplate = tempNextItemTemplate.(*ItemTemplate)
	}

	return nil
}

func (t *SystemEquipShengHenTemplate) PatchAfterCheck() {
	// //动态强化模板
	// t.strengthenTemplateMap = make(map[int32]*GoldEquipStrengthenTemplate)
	// //赋值 strengthenTemplateMap
	// for strengthenTemplate := t.SystemStrengthenTemplate; strengthenTemplate != nil; strengthenTemplate = strengthenTemplate.nextSystemStrengthenTemplate {
	// 	level := strengthenTemplate.Level
	// 	t.strengthenTemplateMap[level] = strengthenTemplate
	// }
}

func (t *SystemEquipShengHenTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// // 动态强化模板
	// if t.SystemStrenId == 0 {
	// 	return fmt.Errorf("UpgradeBeginId [%d] invalid", t.SystemStrenId)
	// }
	// tempSystemStrengthenTemplate := template.GetTemplateService().Get(int(t.SystemStrenId), (*SystemStrengthenTemplate)(nil))
	// if tempSystemStrengthenTemplate == nil {
	// 	return fmt.Errorf("UpgradeBeginId [%d] invalid", t.SystemStrenId)
	// }
	// systemStrengthenTemplate, ok := tempSystemStrengthenTemplate.(*SystemStrengthenTemplate)
	// if !ok {
	// 	return fmt.Errorf("UpgradeBeginId [%d] invalid", t.SystemStrenId)
	// }
	// if systemStrengthenTemplate.Level != 1 {
	// 	return fmt.Errorf("UpgradeBeginId [%d] invalid", t.SystemStrenId)
	// }
	// t.minSystemStrengthenTemplate = systemStrengthenTemplate

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

	//物品校验
	for itemId, itemNum := range t.needItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.NeedItem)
			return template.NewTemplateFieldError("NeedItem", err)
		}
		if itemNum <= 0 {
			err = fmt.Errorf("[%s] invalid", t.NeedItemNum)
			return template.NewTemplateFieldError("NeedItemNum", err)
		}
	}
	for itemId, itemNum := range t.returnItemMap {
		itemTemplate := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.MeltingReturnId)
			return template.NewTemplateFieldError("MeltingReturnId", err)
		}
		if itemNum <= 0 {
			err = fmt.Errorf("[%s] invalid", t.MeltingReturnCount)
			return template.NewTemplateFieldError("MeltingReturnCount", err)
		}
	}

	//
	if t.nextItemTemplate != nil {
		equipmentTemplate := t.nextItemTemplate.GetSystemEquipTemplate()
		if equipmentTemplate == nil {
			err = fmt.Errorf("[%d] 不是装备", t.Next)
			return template.NewTemplateFieldError("next", err)
		}
		if !equipmentTemplate.HasCondition() {
			err = fmt.Errorf("[%d] 没有条件", t.Next)
			return template.NewTemplateFieldError("next", err)
		}
		//成功概率1-10000
		err = validator.RangeValidate(float64(t.SuccessRate), float64(1), true, float64(common.MAX_RATE), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.SuccessRate)
			err = template.NewTemplateFieldError("successRate", err)
			return
		}
	}
	return nil
}

func (edt *SystemEquipShengHenTemplate) FileName() string {
	return "tb_system_equip_shenghen.json"
}

func init() {
	template.Register((*SystemEquipShengHenTemplate)(nil))
}
