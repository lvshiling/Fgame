package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"

	"fmt"
)

//宝宝读书配置
type BabyLearnTemplate struct {
	*BabyLearnTemplateVO
	returnItemMap map[int32]int32
	nextTemp      *BabyLearnTemplate
}

func (t *BabyLearnTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyLearnTemplate) GetNextTemplate() *BabyLearnTemplate {
	return t.nextTemp
}

func (t *BabyLearnTemplate) GetReturnItemMap() map[int32]int32 {
	return t.returnItemMap
}

func (t *BabyLearnTemplate) PatchAfterCheck() {}

func (t *BabyLearnTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//
	t.returnItemMap = make(map[int32]int32)
	returnItemIdArr, err := utils.SplitAsIntArray(t.ZsReturnItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ZsReturnItem)
		return template.NewTemplateFieldError("ZsReturnItem", err)
	}
	returnItemCountArr, err := utils.SplitAsIntArray(t.ZsReturnCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ZsReturnCount)
		return template.NewTemplateFieldError("ZsReturnCount", err)
	}

	if len(returnItemIdArr) != len(returnItemCountArr) {
		err = fmt.Errorf("ZsReturnItem[%s],ZsReturnCount[%s]无效", t.ZsReturnItem, t.ZsReturnCount)
		return template.NewTemplateFieldError("ZsReturnItem or ZsReturnCount", err)
	}

	for index, itemId := range returnItemIdArr {
		itemTmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itemTmpObj == nil {
			err = fmt.Errorf("[%s] invalid", t.ZsReturnItem)
			return template.NewTemplateFieldError("ZsReturnItem", err)
		}

		err = validator.MinValidate(float64(returnItemCountArr[index]), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("ZsReturnCount", err)
		}
		t.returnItemMap[itemId] += returnItemCountArr[index]
	}

	return nil
}

func (t *BabyLearnTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//下一阶强化
	if t.NextId != 0 {
		if t.NextId-t.Id != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		to := template.GetTemplateService().Get(t.NextId, (*BabyLearnTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*BabyLearnTemplate)

		if t.nextTemp.Level-t.Level != 1 {
			err = fmt.Errorf("[%d] invalid", t.Level)
			return template.NewTemplateFieldError("Level", err)
		}
	}

	//等级
	err = validator.MinValidate(float64(t.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//经验
	err = validator.MinValidate(float64(t.Experience), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Experience)
		return template.NewTemplateFieldError("Experience", err)
	}

	//属性倍数
	err = validator.MinValidate(float64(t.BeiShu), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BeiShu)
		return template.NewTemplateFieldError("BeiShu", err)
	}

	return nil
}

func (t *BabyLearnTemplate) FileName() string {
	return "tb_baobao_level.json"
}

func init() {
	template.Register((*BabyLearnTemplate)(nil))
}
