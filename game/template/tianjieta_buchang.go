package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//天劫塔补偿配置
type TianJieTaBuChangTemplate struct {
	*TianJieTaBuChangTemplateVO
	returnItemMap map[int32]int32 //补偿的物品
}

func (t *TianJieTaBuChangTemplate) TemplateId() int {
	return t.Id
}

func (t *TianJieTaBuChangTemplate) GetReturnItemMap() map[int32]int32 {
	return t.returnItemMap
}

func (t *TianJieTaBuChangTemplate) PatchAfterCheck() {}

func (t *TianJieTaBuChangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//物品
	t.returnItemMap = make(map[int32]int32)
	rewItemIdList, err := utils.SplitAsIntArray(t.ReturnItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ReturnItemId)
		return template.NewTemplateFieldError("ReturnItemId", err)
	}
	rewItemCountList, err := utils.SplitAsIntArray(t.ReturnItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ReturnItemCount)
		return template.NewTemplateFieldError("ReturnItemCount", err)
	}
	if len(rewItemIdList) != len(rewItemCountList) {
		err = fmt.Errorf("[%s] invalid", t.ReturnItemCount)
		return template.NewTemplateFieldError("ReturnItemCount Or ReturnItemId", err)
	}
	if len(rewItemIdList) > 0 {
		for index, itemId := range rewItemIdList {
			t.returnItemMap[itemId] += rewItemCountList[index]
		}
	}

	return nil
}

func (t *TianJieTaBuChangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 层数
	err = validator.MinValidate(float64(t.Number), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Number)
		err = template.NewTemplateFieldError("Number", err)
		return
	}

	//验证  物品
	for itemId, num := range t.returnItemMap {
		itemTemp := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTemp == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("ReturnItemId", err)
			return
		}

		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			err = template.NewTemplateFieldError("ReturnItemCount", err)
			return
		}
	}

	return nil
}

func (t *TianJieTaBuChangTemplate) FileName() string {
	return "tb_tianjieta_buchang.json"
}

func init() {
	template.Register((*TianJieTaBuChangTemplate)(nil))
}
