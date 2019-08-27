package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fmt"
)

//挖矿配置
type MiningTemplate struct {
	*MiningTemplateVO
	useItemTemplate *ItemTemplate
}

func (mt *MiningTemplate) TemplateId() int {
	return mt.Id
}

func (mt *MiningTemplate) GetUseItemTemplate() *ItemTemplate {
	return mt.useItemTemplate
}

func (mt *MiningTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证 item_id_1
	if mt.ItemCount1 != 0 {
		to := template.GetTemplateService().Get(int(mt.ItemId1), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", mt.ItemId1)
			err = template.NewTemplateFieldError("ItemId1", err)
			return
		}

		//验证 item_count_1
		err = validator.MinValidate(float64(mt.ItemCount1), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", mt.ItemCount1)
			err = template.NewTemplateFieldError("ItemCount1", err)
			return
		}
		mt.useItemTemplate = to.(*ItemTemplate)
	}

	return nil
}

func (mt *MiningTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(mt.FileName(), mt.TemplateId(), err)
			return
		}
	}()

	//验证 level
	err = validator.MinValidate(float64(mt.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Level)
		err = template.NewTemplateFieldError("Level", err)
		return
	}

	//验证 interval_time
	err = validator.MinValidate(float64(mt.IntervalTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.IntervalTime)
		err = template.NewTemplateFieldError("IntervalTime", err)
		return
	}

	//验证 revenue
	err = validator.MinValidate(float64(mt.Revenue), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.Revenue)
		err = template.NewTemplateFieldError("Revenue", err)
		return
	}

	//验证 limit_max
	err = validator.MinValidate(float64(mt.LimitMax), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.LimitMax)
		err = template.NewTemplateFieldError("LimitMax", err)
		return
	}

	//验证 need_yinliang
	err = validator.MinValidate(float64(mt.NeedYinLiang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.NeedYinLiang)
		err = template.NewTemplateFieldError("NeedYinLiang", err)
		return
	}

	//验证 need_gold
	err = validator.MinValidate(float64(mt.NeedGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", mt.NeedGold)
		err = template.NewTemplateFieldError("NeedGold", err)
		return
	}

	return nil
}

func (mt *MiningTemplate) PatchAfterCheck() {

}

func (mt *MiningTemplate) FileName() string {
	return "tb_mining.json"
}

func init() {
	template.Register((*MiningTemplate)(nil))
}
