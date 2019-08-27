package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	droptypes "fgame/fgame/game/drop/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//宝库掉落配置
type DropBaoKuTemplate struct {
	*DropBaoKuTemplateVO
	bindType                itemtypes.ItemBindType
	goldequipFuJiaTemp      *GoldEquipFuJiaTemplate
	startUpstarPoolTemplate *GoldEquipStrengthenPoolTemplate
	upstarPoolList          []*GoldEquipStrengthenPoolTemplate
}

func (t *DropBaoKuTemplate) TemplateId() int {
	return t.Id
}

func (t *DropBaoKuTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	return nil
}

func (t *DropBaoKuTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 drop_id
	err = validator.MinValidate(float64(t.DropId), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.DropId)
		err = template.NewTemplateFieldError("DropId", err)
		return
	}

	//验证 item_id
	to := template.GetTemplateService().Get(int(t.ItemId), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.ItemId)
		err = template.NewTemplateFieldError("ItemId", err)
		return
	}
	itemTemp := to.(*ItemTemplate)
	if itemTemp.itemType != itemtypes.ItemTypeGoldEquip {
		if t.GoldEquipMin > 0 || t.GoldEquipMax > 0 {
			err = fmt.Errorf("[%d] invalid, can`t have level", t.Id)
			return template.NewTemplateFieldError("id", err)
		}
	}
	if itemTemp.IsGoldEquip() {
		//验证 max_count
		err = validator.MaxValidate(float64(t.MaxCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.MaxCount)
			err = template.NewTemplateFieldError("MaxCount", err)
			return
		}
	}

	//验证 rate
	err = validator.RangeValidate(float64(t.Rate), float64(0), true, float64(droptypes.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Rate)
		err = template.NewTemplateFieldError("Rate", err)
		return
	}

	//验证 min_count
	err = validator.MinValidate(float64(t.MinCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MinCount)
		err = template.NewTemplateFieldError("MinCount", err)
		return
	}

	//验证 max_count
	err = validator.MinValidate(float64(t.MaxCount), float64(1), true)
	if err != nil || t.MaxCount < t.MinCount {
		err = fmt.Errorf("[%d] invalid", t.MaxCount)
		err = template.NewTemplateFieldError("MaxCount", err)
		return
	}

	//验证 min_stack
	err = validator.MinValidate(float64(t.MinStack), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.MinStack)
		err = template.NewTemplateFieldError("MinStack", err)
		return
	}

	//验证 max_stack
	err = validator.MinValidate(float64(t.MaxStack), float64(1), true)
	if err != nil || t.MaxStack < t.MinStack {
		err = fmt.Errorf("[%d] invalid", t.MaxStack)
		err = template.NewTemplateFieldError("MaxStack", err)
		return
	}

	// //验证 gold_equip_min
	// err = validator.MinValidate(float64(t.GoldEquipMin), float64(1), true)
	// if err != nil {
	// 	err = fmt.Errorf("[%d] invalid", t.GoldEquipMin)
	// 	err = template.NewTemplateFieldError("GoldEquipMin", err)
	// 	return
	// }

	// //验证 gold_equip_max
	// err = validator.MinValidate(float64(t.GoldEquipMax), float64(1), true)
	// if err != nil || t.GoldEquipMax < t.GoldEquipMin {
	// 	err = fmt.Errorf("[%d] invalid", t.GoldEquipMax)
	// 	err = template.NewTemplateFieldError("GoldEquipMax", err)
	// 	return
	// }

	//验证 bind_type
	t.bindType = itemtypes.ItemBindType(t.BindType)
	if !t.bindType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.BindType)
		err = template.NewTemplateFieldError("BindType", err)
		return
	}

	//验证 exist_time
	err = validator.MinValidate(float64(t.ExistTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ExistTime)
		err = template.NewTemplateFieldError("ExistTime", err)
		return
	}

	//验证 protected_time
	err = validator.MinValidate(float64(t.ProtectedTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ProtectedTime)
		err = template.NewTemplateFieldError("ProtectedTime", err)
		return
	}

	//验证 fail_time
	err = validator.MinValidate(float64(t.FailTime), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.FailTime)
		err = template.NewTemplateFieldError("FailTime", err)
		return
	}

	//验证 is_drop_agin
	err = validator.MinValidate(float64(t.IsDropAgin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.IsDropAgin)
		err = template.NewTemplateFieldError("IsDropAgin", err)
		return
	}

	//验证 is_drop_suppress
	err = validator.MinValidate(float64(t.IsDropSuppress), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.IsDropSuppress)
		err = template.NewTemplateFieldError("IsDropSuppress", err)
		return
	}

	//金装附件属性
	if t.FristFujiaId > 0 {
		tempObj := template.GetTemplateService().Get(int(t.FristFujiaId), (*GoldEquipFuJiaTemplate)(nil))
		if tempObj == nil {
			err = fmt.Errorf("[%d] invalid", t.FristFujiaId)
			err = template.NewTemplateFieldError("FristFujiaId", err)
			return err
		}
		t.goldequipFuJiaTemp = tempObj.(*GoldEquipFuJiaTemplate)
	}

	//金装升星池起始id
	if t.StrengthenPoolId > 0 {
		tempObj := template.GetTemplateService().Get(int(t.StrengthenPoolId), (*GoldEquipStrengthenPoolTemplate)(nil))
		if tempObj == nil {
			err = fmt.Errorf("[%d] invalid", t.StrengthenPoolId)
			err = template.NewTemplateFieldError("StrengthenPoolId", err)
			return err
		}
		t.startUpstarPoolTemplate = tempObj.(*GoldEquipStrengthenPoolTemplate)
	}

	return nil
}

func (t *DropBaoKuTemplate) PatchAfterCheck() {
	// 金装随机升星等级
	if t.startUpstarPoolTemplate != nil {
		for startTemp := t.startUpstarPoolTemplate; startTemp != nil; startTemp = startTemp.GetNextTemplate() {
			t.upstarPoolList = append(t.upstarPoolList, startTemp)
		}
	}
}

func (t *DropBaoKuTemplate) RandomNum() int32 {
	minCount := t.MinCount
	maxCount := t.MaxCount
	count := mathutils.RandomRange(int(minCount), int(maxCount))
	return int32(count)
}

func (t *DropBaoKuTemplate) RandomGoldEquipLevel() int32 {
	minCount := t.GoldEquipMin
	maxCount := t.GoldEquipMax
	count := mathutils.RandomRange(int(minCount), int(maxCount))
	return int32(count)
}

func (t *DropBaoKuTemplate) RandomGoldEquipUpstarLevel() int32 {
	if len(t.upstarPoolList) < 1 {
		return 0
	}

	var weights []int64
	for _, poolTemp := range t.upstarPoolList {
		weights = append(weights, int64(poolTemp.Rate))
	}
	index := mathutils.RandomWeights(weights)
	if index < 0 {
		return 0
	}

	return t.upstarPoolList[index].Level
}

func (t *DropBaoKuTemplate) RandomGoldEquipAttr() (attrList []int32, isRandom bool) {
	if t.goldequipFuJiaTemp == nil {
		return
	}
	return t.goldequipFuJiaTemp.RandomAttr(), true
}

func (t *DropBaoKuTemplate) RandomStack() int32 {
	minStack := t.MinStack
	maxStack := t.MaxStack
	stack := mathutils.RandomRange(int(minStack), int(maxStack))
	return int32(stack)
}

func (t *DropBaoKuTemplate) GetBindType() itemtypes.ItemBindType {
	return t.bindType
}

func (t *DropBaoKuTemplate) FileName() string {
	return "tb_drop_baoku.json"
}

func init() {
	template.Register((*DropBaoKuTemplate)(nil))
}
