package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

type GemstoneSlotTemplate struct {
	*GemstoneSlotTemplateVO
	bodyPosition inventorytypes.BodyPositionType
	gemSubType   itemtypes.ItemSubType
	needItemMap  map[int32]int32
}

func (t *GemstoneSlotTemplate) GetBodyPosition() inventorytypes.BodyPositionType {
	return t.bodyPosition
}

func (t *GemstoneSlotTemplate) GetGemSubType() itemtypes.ItemSubType {
	return t.gemSubType
}

func (t *GemstoneSlotTemplate) GetNeedItemMap() map[int32]int32 {
	return t.needItemMap
}

func (t *GemstoneSlotTemplate) TemplateId() int {
	return t.Id
}

func (t *GemstoneSlotTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	t.bodyPosition = inventorytypes.BodyPositionType(t.Position)
	t.gemSubType = itemtypes.CreateItemGemSubType(t.GemstoneType)

	//物品
	t.needItemMap = make(map[int32]int32)
	needItemIdList, err := utils.SplitAsIntArray(t.NeedItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.NeedItemId)
		return template.NewTemplateFieldError("NeedItemId", err)
	}
	needItemCountList, err := utils.SplitAsIntArray(t.NeedItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.NeedItemCount)
		return template.NewTemplateFieldError("NeedItemCount", err)
	}
	if len(needItemIdList) != len(needItemCountList) {
		err = fmt.Errorf("[%s] invalid", t.NeedItemCount)
		return template.NewTemplateFieldError("NeedItemCount Or NeedItemId", err)
	}
	if len(needItemIdList) > 0 {
		for index, itemId := range needItemIdList {
			t.needItemMap[itemId] += needItemCountList[index]
		}
	}

	return nil
}

func (t *GemstoneSlotTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	if !t.bodyPosition.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Position)
		return template.NewTemplateFieldError("Position", err)
	}

	if !t.gemSubType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.GemstoneType)
		return template.NewTemplateFieldError("GemstoneType", err)
	}

	//验证order
	err = validator.MinValidate(float64(t.Order), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Order)
		return template.NewTemplateFieldError("Order", err)
	}

	err = validator.MinValidate(float64(t.NeedLevel), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedLevel)
		return template.NewTemplateFieldError("NeedLevel", err)
	}

	err = validator.MinValidate(float64(t.NeedLayer), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.NeedLayer)
		return template.NewTemplateFieldError("NeedLayer", err)
	}

	//验证  物品
	for itemId, num := range t.needItemMap {
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("NeedItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			err = template.NewTemplateFieldError("NeedItemCount", err)
			return
		}
	}

	return nil
}

func (t *GemstoneSlotTemplate) PatchAfterCheck() {

}

func (t *GemstoneSlotTemplate) FileName() string {
	return "tb_gemstone_slot.json"
}

func init() {
	template.Register((*GemstoneSlotTemplate)(nil))
}
