package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	constanttypes "fgame/fgame/game/constant/types"
	housetypes "fgame/fgame/game/house/types"
	"fmt"
)

type HouseTemplate struct {
	*HouseTemplateVO
	useItemMap     map[int32]int32
	repairItemMap  map[int32]int32
	rewardsItemMap map[int32]int32
	emailRentMap   map[int32]int32
	houseType      housetypes.HouseType
	nextTemp       *HouseTemplate
}

func (t *HouseTemplate) TemplateId() int {
	return t.Id
}

func (t *HouseTemplate) FileName() string {
	return "tb_fangzi.json"
}

//组合成需要的数据
func (t *HouseTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//升级消耗
	t.useItemMap = make(map[int32]int32)
	useItemIdArr, err := utils.SplitAsIntArray(t.UseItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}
	useItemCountArr, err := utils.SplitAsIntArray(t.UseItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UseItemCount)
		return template.NewTemplateFieldError("UseItemCount", err)
	}
	if len(useItemIdArr) != len(useItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.UseItemId, t.UseItemCount)
		return template.NewTemplateFieldError("UseItemId or UseItemCount", err)
	}
	if len(useItemIdArr) > 0 {
		for index, itemId := range useItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.UseItemId)
				return template.NewTemplateFieldError("UseItemId", err)
			}

			err = validator.MinValidate(float64(useItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("UseItemCount", err)
			}

			t.useItemMap[itemId] = useItemCountArr[index]
		}
	}

	//验证：维修物品
	t.repairItemMap = make(map[int32]int32)
	repairItemIdArr, err := utils.SplitAsIntArray(t.FixItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FixItemId)
		return template.NewTemplateFieldError("FixItemId", err)
	}
	repairItemCountArr, err := utils.SplitAsIntArray(t.FixItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FixItemCount)
		return template.NewTemplateFieldError("FixItemCount", err)
	}
	if len(repairItemIdArr) != len(repairItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.FixItemId, t.FixItemCount)
		return template.NewTemplateFieldError("FixItemId or FixItemCount", err)
	}
	if len(repairItemIdArr) > 0 {
		for index, itemId := range repairItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.FixItemId)
				return template.NewTemplateFieldError("FixItemId", err)
			}

			err = validator.MinValidate(float64(repairItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("FixItemCount", err)
			}

			t.repairItemMap[itemId] = repairItemCountArr[index]
		}
	}

	//验证：挑战通关奖励物品
	t.rewardsItemMap = make(map[int32]int32)
	rewItemIdArr, err := utils.SplitAsIntArray(t.UplevGetItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UplevGetItem)
		return template.NewTemplateFieldError("UplevGetItem", err)
	}
	rewItemCountArr, err := utils.SplitAsIntArray(t.UplevGetItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.UplevGetItemCount)
		return template.NewTemplateFieldError("UplevGetItemCount", err)
	}
	if len(rewItemIdArr) != len(rewItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.UplevGetItem, t.UplevGetItemCount)
		return template.NewTemplateFieldError("UplevGetItem or UplevGetItemCount", err)
	}
	if len(rewItemIdArr) > 0 {
		for index, itemId := range rewItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.UplevGetItem)
				return template.NewTemplateFieldError("UplevGetItem", err)
			}

			err = validator.MinValidate(float64(rewItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("UplevGetItemCount", err)
			}

			t.rewardsItemMap[itemId] = rewItemCountArr[index]
		}
	}

	return nil
}

//检查有效性
func (t *HouseTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 next_id
	if t.NextId != 0 {
		to := template.GetTemplateService().Get(int(t.NextId), (*HouseTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*HouseTemplate)

		if t.nextTemp.HouseIndex != t.HouseIndex {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		if t.nextTemp.Type != t.Type {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		diff := t.nextTemp.Level - int32(t.Level)
		if diff != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}
	}

	t.houseType = housetypes.HouseType(t.Type)
	if !t.houseType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		return template.NewTemplateFieldError("Type", err)
	}

	//验证： 等级
	err = validator.MinValidate(float64(t.Level), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Level", err)
	}

	//验证：损坏率
	err = validator.MinValidate(float64(t.BrokenPercent), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("BrokenPercent", err)
	}

	//验证：租金
	err = validator.MinValidate(float64(t.Rent), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("Rent", err)
	}

	//验证：房价
	err = validator.MinValidate(float64(t.HousePrice), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("HousePrice", err)
	}
	//验证：提前出手房价
	err = validator.MinValidate(float64(t.AdvanceSalePercent), float64(0), true)
	if err != nil {
		return template.NewTemplateFieldError("AdvanceSalePercent", err)
	}

	return nil
}

//检验后组合
func (t *HouseTemplate) PatchAfterCheck() {
	t.emailRentMap = make(map[int32]int32)
	switch t.houseType {
	case housetypes.HouseTypeSilver:
		{
			t.emailRentMap[constanttypes.SilverItem] += t.Rent
		}
	case housetypes.HouseTypeBindGold:
		{
			t.emailRentMap[constanttypes.BindGoldItem] += t.Rent
		}
	}
}

//获取房子类型
func (t *HouseTemplate) GetHouseType() housetypes.HouseType {
	return t.houseType
}

//获取升级消耗
func (t *HouseTemplate) GetUseItemMap() map[int32]int32 {
	return t.useItemMap
}

//获取维修消耗物品
func (t *HouseTemplate) GetRepairItemMap() map[int32]int32 {
	return t.repairItemMap
}

//获取装修奖励物品
func (t *HouseTemplate) GetRewardsItemMap() map[int32]int32 {
	return t.rewardsItemMap
}

//获取租金邮件
func (t *HouseTemplate) GetRentItemMap() map[int32]int32 {
	return t.emailRentMap
}

//获取装修奖励物品
func (t *HouseTemplate) IsMaxLevel() bool {
	return t.NextId == 0
}

//下一级模板
func (t *HouseTemplate) GetNextTemp() *HouseTemplate {
	return t.nextTemp
}

func init() {
	template.Register((*HouseTemplate)(nil))
}
