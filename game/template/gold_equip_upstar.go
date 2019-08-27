package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	propertytypes "fgame/fgame/game/property/types"
	"fmt"
)

//元神金装强化配置
type GoldEquipUpstarTemplate struct {
	*GoldEquipUpstarTemplateVO
	nextTemp           *GoldEquipUpstarTemplate
	needItemMap        map[int32]int32
	tunshiReturnMap    map[int32]int32
	battleAttrMap      map[propertytypes.BattlePropertyType]int64
	failReturnTemplate *GoldEquipUpstarTemplate
}

func (t *GoldEquipUpstarTemplate) TemplateId() int {
	return t.Id
}

func (t *GoldEquipUpstarTemplate) GetBattleAttrMap() map[propertytypes.BattlePropertyType]int64 {
	return t.battleAttrMap
}

func (t *GoldEquipUpstarTemplate) GetNextTemplate() *GoldEquipUpstarTemplate {
	return t.nextTemp
}

func (t *GoldEquipUpstarTemplate) GetFaildReturnTemplate() *GoldEquipUpstarTemplate {
	return t.failReturnTemplate
}

func (t *GoldEquipUpstarTemplate) GetNeedItemMap() map[int32]int32 {
	needItemMap := make(map[int32]int32)
	for itemId, num := range t.needItemMap {
		_, ok := needItemMap[itemId]
		if ok {
			needItemMap[itemId] += num
		} else {
			needItemMap[itemId] = num
		}
	}
	return needItemMap
}

func (t *GoldEquipUpstarTemplate) GetReturnItemMap() map[int32]int32 {
	return t.tunshiReturnMap
}

func (t *GoldEquipUpstarTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//属性
	t.battleAttrMap = make(map[propertytypes.BattlePropertyType]int64)
	if t.AddHp > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeMaxHP] = int64(t.AddHp)
	}
	if t.AddAttack > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeAttack] = int64(t.AddAttack)
	}
	if t.AddDef > 0 {
		t.battleAttrMap[propertytypes.BattlePropertyTypeDefend] = int64(t.AddDef)
	}

	//所需物品
	t.needItemMap = make(map[int32]int32)
	needItemIdList, err := utils.SplitAsIntArray(t.NeedItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.NeedItem)
		return template.NewTemplateFieldError("NeedItem", err)
	}
	needItemCountList, err := utils.SplitAsIntArray(t.NeedCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.NeedCount)
		return template.NewTemplateFieldError("NeedCount", err)
	}
	if len(needItemIdList) != len(needItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.NeedItem, t.NeedCount)
		return template.NewTemplateFieldError("NeedItem or NeedCount", err)
	}
	if len(needItemIdList) > 0 {
		//组合数据
		for index, itemId := range needItemIdList {
			_, ok := t.needItemMap[itemId]
			if ok {
				t.needItemMap[itemId] += needItemCountList[index]
			} else {
				t.needItemMap[itemId] = needItemCountList[index]
			}
		}
	}

	// 吞噬返还
	t.tunshiReturnMap = make(map[int32]int32)
	returnItemIdList, err := utils.SplitAsIntArray(t.MeltingReturnId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MeltingReturnId)
		return template.NewTemplateFieldError("MeltingReturnId", err)
	}
	returnItemCountList, err := utils.SplitAsIntArray(t.MeltingReturnCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MeltingReturnCount)
		return template.NewTemplateFieldError("MeltingReturnCount", err)
	}
	if len(returnItemIdList) != len(returnItemCountList) {
		err = fmt.Errorf("[%s][%s] invalid", t.MeltingReturnId, t.MeltingReturnCount)
		return template.NewTemplateFieldError("MeltingReturnId or MeltingReturnCount", err)
	}
	if len(returnItemIdList) > 0 {
		//组合数据
		for index, itemId := range returnItemIdList {
			_, ok := t.tunshiReturnMap[itemId]
			if ok {
				t.tunshiReturnMap[itemId] += returnItemCountList[index]
			} else {
				t.tunshiReturnMap[itemId] = returnItemCountList[index]
			}
		}
	}

	return nil
}

func (t *GoldEquipUpstarTemplate) PatchAfterCheck() {
}

func (t *GoldEquipUpstarTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 next_id
	if t.NextId != 0 {
		diff := t.NextId - int32(t.Id)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
		to := template.GetTemplateService().Get(int(t.NextId), (*GoldEquipUpstarTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		t.nextTemp = to.(*GoldEquipUpstarTemplate)
	}

	//等级
	if err = validator.MinValidate(float64(t.Level), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//成功率
	if err = validator.MinValidate(float64(t.SuccessRate), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("SuccessRate", err)
		return
	}

	//失败率
	if err = validator.MinValidate(float64(t.FailReturnRate), float64(0), true); err != nil {
		err = template.NewTemplateFieldError("FailReturnRate", err)
		return
	}

	//失败回退等级
	if t.FailReturnLevel != 0 {
		//失败
		tempFailReturnGoldEquipUpstarTemplate := template.GetTemplateService().Get(int(t.FailReturnLevel), (*GoldEquipUpstarTemplate)(nil))
		if tempFailReturnGoldEquipUpstarTemplate == nil {
			err = fmt.Errorf("[%d] invalid", t.FailReturnLevel)
			return template.NewTemplateFieldError("FailReturnLevel", err)
		}
		t.failReturnTemplate = tempFailReturnGoldEquipUpstarTemplate.(*GoldEquipUpstarTemplate)
	}

	// 消耗物品id
	for itemId, num := range t.needItemMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("UseItem", fmt.Errorf("[%d] invalid", itemId))
		}
		if err = validator.MinValidate(float64(num), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("UseCount", err)
			return
		}
	}

	// 返回数量
	for itemId, num := range t.tunshiReturnMap {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("MeltingReturnId", fmt.Errorf("[%d] invalid", itemId))
		}
		if err = validator.MinValidate(float64(num), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("MeltingReturnCount", err)
			return
		}
	}

	// 防爆物品
	if t.ProtectItemId > 0 {
		itemTmpObj := template.GetTemplateService().Get(int(t.ProtectItemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			return template.NewTemplateFieldError("ProtectItemId", fmt.Errorf("[%d] invalid", t.ProtectItemId))
		}
		if err = validator.MinValidate(float64(t.ProtectItemCount), float64(1), true); err != nil {
			err = template.NewTemplateFieldError("ProtectItemCount", err)
			return
		}
	}

	return nil
}

func (t *GoldEquipUpstarTemplate) FileName() string {
	return "tb_goldequip_upstar.json"
}

func init() {
	template.Register((*GoldEquipUpstarTemplate)(nil))
}
