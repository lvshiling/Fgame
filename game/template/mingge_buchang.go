package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//命格补偿配置
type MingGeBuchangTemplate struct {
	*MingGeBuchangTemplateVO
	returnItemMap map[int32]int32
}

func (t *MingGeBuchangTemplate) TemplateId() int {
	return t.Id
}

func (t *MingGeBuchangTemplate) GetReturnItemMap() map[int32]int32 {
	return t.returnItemMap
}

func (t *MingGeBuchangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	t.returnItemMap = make(map[int32]int32)
	returnItemIdArr, err := utils.SplitAsIntArray(t.ReturnItemId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ReturnItemId)
		return template.NewTemplateFieldError("ReturnItemId", err)
	}
	returnItemCountArr, err := utils.SplitAsIntArray(t.ReturnItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ReturnItemCount)
		return template.NewTemplateFieldError("ReturnItemCount", err)
	}

	if len(returnItemIdArr) != len(returnItemCountArr) {
		err = fmt.Errorf("ReturnItemId[%s],ReturnItemCount[%s]无效", t.ReturnItemId, t.ReturnItemCount)
		return template.NewTemplateFieldError("ReturnItemId or ReturnItemCount", err)
	}

	for index, itemId := range returnItemIdArr {
		t.returnItemMap[itemId] += returnItemCountArr[index]
	}

	return nil
}

func (t *MingGeBuchangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	for itemId, itemNum := range t.returnItemMap {
		itetmpObj := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if itetmpObj == nil {
			err = fmt.Errorf("[%s] invalid", t.ReturnItemId)
			return template.NewTemplateFieldError("ReturnItemId", err)
		}

		err = validator.MinValidate(float64(itemNum), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.ReturnItemCount)
			return template.NewTemplateFieldError("ReturnItemCount", err)
		}
	}

	return nil
}

func (t *MingGeBuchangTemplate) PatchAfterCheck() {

}

func (t *MingGeBuchangTemplate) FileName() string {
	return "tb_minggong_buchang.json"
}

func init() {
	template.Register((*MingGeBuchangTemplate)(nil))
}
