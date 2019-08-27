package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	itemtypes "fgame/fgame/game/item/types"
	"fmt"
)

type GodCastingEquipTemplate struct {
	*GodCastingEquipTemplateVO
	useItemIdList    []int32
	useItemCountList []int32
	useItemMap       map[int32]int32 //key:物品ID，value:物品数量
	nextItemTemplate *ItemTemplate
}

func (t *GodCastingEquipTemplate) TemplateId() int {
	return t.Id
}

func (t *GodCastingEquipTemplate) GetUseItemMap() map[int32]int32 {
	return t.useItemMap
}

func (t *GodCastingEquipTemplate) GetNextItemTemplate() *ItemTemplate {
	return t.nextItemTemplate
}

//检查有效性
func (t *GodCastingEquipTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//物品ID
	err = validator.MinValidate(float64(t.ItemId), float64(0), false)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.ItemId)
		return template.NewTemplateFieldError("ItemId", err)
	}

	//用于强化的物品列表
	t.useItemIdList, err = utils.SplitAsIntArray(t.UseItemId)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItemId)
		return template.NewTemplateFieldError("UseItemId", err)
	}

	//用于强化的物品数量列表
	t.useItemCountList, err = utils.SplitAsIntArray(t.UseItemCount)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UseItemCount)
		return template.NewTemplateFieldError("UseItemCount", err)
	}

	//升级成功率万分比
	err = validator.RangeValidate(float64(t.UpdateWfb), float64(0), true, float64(10000), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.UpdateWfb)
		return template.NewTemplateFieldError("UpdateWfb", err)
	}

	//最小次数
	err = validator.MinValidate(float64(t.TimesMin), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMin)
		return template.NewTemplateFieldError("TimesMin", err)
	}

	//最大次数
	err = validator.MinValidate(float64(t.TimesMax), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.TimesMax)
		return template.NewTemplateFieldError("TimesMax", err)
	}

	//检验下一物品ID是不是元神金装
	if t.nextItemTemplate.GetItemType() != itemtypes.ItemTypeGoldEquip {
		err = fmt.Errorf("GodCastingEquipTemplate[%d] invalid,which is not ItemTypeGoldEquip", t.ItemId)
		err = template.NewTemplateFieldError("MagicConditionParameter", err)
		return
	}

	return
}

//组合成需要的数据
func (t *GodCastingEquipTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	temp := template.GetTemplateService().Get(int(t.ItemId), (*ItemTemplate)(nil))
	itemTemp, _ := temp.(*ItemTemplate)
	if itemTemp == nil {
		err = fmt.Errorf("GodCastingEquipTemplate[%d] invalid", t.ItemId)
		err = template.NewTemplateFieldError("MagicConditionParameter", err)
		return
	}

	t.nextItemTemplate = itemTemp

	return
}

//检验后组合
func (t *GodCastingEquipTemplate) PatchAfterCheck() {
	//物品列表
	t.useItemMap = make(map[int32]int32)
	for i, itemId := range t.useItemIdList {
		cnt := t.useItemCountList[i]
		t.useItemMap[itemId] = cnt
	}
}

func (t *GodCastingEquipTemplate) FileName() string {
	return "tb_shenzhuequip.json"
}

func init() {
	template.Register((*GodCastingEquipTemplate)(nil))
}
