package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

type BaGuaMiJingBuChangTemplate struct {
	*BaGuaMiJingBuChangTemplateVO
	itemMap map[int32]int32
}

// 返回补偿的物品
func (t *BaGuaMiJingBuChangTemplate) GetItemMap() map[int32]int32 {
	return t.itemMap
}

func (t *BaGuaMiJingBuChangTemplate) TemplateId() int {
	return t.Id
}

func (t *BaGuaMiJingBuChangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 补偿的物品
	t.itemMap = make(map[int32]int32)
	useItemIdArr, err := utils.SplitAsIntArray(t.ItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemId)
		return template.NewTemplateFieldError("ItemId", err)
	}
	useItemCountArr, err := utils.SplitAsIntArray(t.ItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ItemCount)
		return template.NewTemplateFieldError("ItemCount", err)
	}
	if len(useItemIdArr) != len(useItemCountArr) {
		err = fmt.Errorf("[%s][%s] invalid", t.ItemId, t.ItemCount)
		return template.NewTemplateFieldError("ItemId or ItemCount", err)
	}
	if len(useItemIdArr) > 0 {
		for index, itemId := range useItemIdArr {
			itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
			if itemTmpObj == nil {
				err = fmt.Errorf("[%s] invalid", t.ItemId)
				return template.NewTemplateFieldError("ItemId", err)
			}

			err = validator.MinValidate(float64(useItemCountArr[index]), float64(1), true)
			if err != nil {
				return template.NewTemplateFieldError("ItemCount", err)
			}

			t.itemMap[itemId] = useItemCountArr[index]
		}
	}

	return
}

func (t *BaGuaMiJingBuChangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	// 验证 层数
	err = validator.MinValidate(float64(t.Level), float64(0), false)
	if err != nil {
		return template.NewTemplateFieldError("Number", err)
	}

	return
}

func (t *BaGuaMiJingBuChangTemplate) PatchAfterCheck() {
}

func (t *BaGuaMiJingBuChangTemplate) FileName() string {
	return "tb_baguamijing_buchang.json"
}

func init() {
	template.Register((*BaGuaMiJingBuChangTemplate)(nil))
}
