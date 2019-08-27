package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fmt"
)

//宝宝洞房配置
type BabyDongFangTemplate struct {
	*BabyDongFangTemplateVO
	failReturnItemMap map[int32]int32
	nextTemp          *BabyDongFangTemplate
}

func (t *BabyDongFangTemplate) TemplateId() int {
	return t.Id
}

func (t *BabyDongFangTemplate) GetFailReturnItemMap() map[int32]int32 {
	return t.failReturnItemMap
}

func (t *BabyDongFangTemplate) PatchAfterCheck() {}

func (t *BabyDongFangTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//洞房失败返还
	t.failReturnItemMap = make(map[int32]int32)
	returnItemIdArr, err := utils.SplitAsIntArray(t.FailReturnItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FailReturnItem)
		return template.NewTemplateFieldError("FailReturnItem", err)
	}
	returnItemCountArr, err := utils.SplitAsIntArray(t.FailReturnItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FailReturnItemCount)
		return template.NewTemplateFieldError("FailReturnItemCount", err)
	}
	if len(returnItemIdArr) != len(returnItemCountArr) {
		err = fmt.Errorf("[%s] invalid", t.FailReturnItem)
		return template.NewTemplateFieldError("FailReturnItem and FailReturnItemCount", err)
	}

	for index, itemId := range returnItemIdArr {
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%s] invalid", t.FailReturnItem)
			return template.NewTemplateFieldError("FailReturnItem", err)
		}

		err = validator.MinValidate(float64(returnItemCountArr[index]), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", t.FailReturnItemCount)
			return template.NewTemplateFieldError("FailReturnItemCount", err)
		}
		t.failReturnItemMap[itemId] += returnItemCountArr[index]
	}

	//
	if t.NextId != 0 {
		if t.NextId-t.Id != 1 {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}

		to := template.GetTemplateService().Get(t.NextId, (*BabyDongFangTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.NextId)
			return template.NewTemplateFieldError("NextId", err)
		}
		t.nextTemp = to.(*BabyDongFangTemplate)

		if t.nextTemp.BabyCount-t.BabyCount != 1 {
			err = fmt.Errorf("[%d] invalid", t.BabyCount)
			return template.NewTemplateFieldError("BabyCount", err)
		}
	}

	return nil
}

func (t *BabyDongFangTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//数量
	err = validator.MinValidate(float64(t.BabyCount), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.BabyCount)
		return template.NewTemplateFieldError("BabyCount", err)
	}

	//道具
	to := template.GetTemplateService().Get(int(t.PregnantItem), (*ItemTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", t.PregnantItem)
		err = template.NewTemplateFieldError("PregnantItem", err)
		return
	}
	err = validator.MinValidate(float64(t.PregnantCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PregnantCount)
		return template.NewTemplateFieldError("PregnantCount", err)
	}

	// 概率
	err = validator.MinValidate(float64(t.PregnantRate), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.PregnantRate)
		return template.NewTemplateFieldError("PregnantRate", err)
	}

	return nil
}

func (t *BabyDongFangTemplate) FileName() string {
	return "tb_baobao_dongfang.json"
}

func init() {
	template.Register((*BabyDongFangTemplate)(nil))
}
