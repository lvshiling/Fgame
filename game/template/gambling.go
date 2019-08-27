package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	gemtypes "fgame/fgame/game/gem/types"
	"fmt"
)

//赌石配置
type GamblingTemplate struct {
	*GamblingTemplateVO
	typ             gemtypes.GemGambleType //赌石类型
	intervalDropMap map[int32]int32        //间隔掉落包
	useItemTemplate *ItemTemplate
}

func (gt *GamblingTemplate) TemplateId() int {
	return gt.Id
}

func (gt *GamblingTemplate) GetUseItemTemplate() *ItemTemplate {
	return gt.useItemTemplate
}

func (gt *GamblingTemplate) GetIntervalDropMap() map[int32]int32 {
	return gt.intervalDropMap
}

func (gt *GamblingTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(gt.FileName(), gt.TemplateId(), err)
			return
		}
	}()

	//验证 typ
	gt.typ = gemtypes.GemGambleType(gt.Type)
	if !gt.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", gt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 need_item
	if gt.NeedItem != 0 {
		to := template.GetTemplateService().Get(int(gt.NeedItem), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", gt.NeedItem)
			err = template.NewTemplateFieldError("NeedItem", err)
			return
		}

		//验证 need_item_num
		err = validator.MinValidate(float64(gt.NeedItemNum), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", gt.NeedItemNum)
			err = template.NewTemplateFieldError("NeedItemNum", err)
			return
		}
		gt.useItemTemplate = to.(*ItemTemplate)
	}

	gt.intervalDropMap = make(map[int32]int32)
	//验证 interval_num_1
	err = validator.MinValidate(float64(gt.IntervalNum1), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", gt.IntervalNum1)
		err = template.NewTemplateFieldError("IntervalNum1", err)
		return
	}
	gt.intervalDropMap[gt.IntervalNum1] = gt.DropId1

	//验证 interval_num_2
	err = validator.MinValidate(float64(gt.IntervalNum2), float64(1), true)
	if err != nil || gt.IntervalNum1 == gt.IntervalNum2 {
		err = fmt.Errorf("[%d] invalid", gt.IntervalNum2)
		err = template.NewTemplateFieldError("IntervalNum2", err)
		return
	}

	gt.intervalDropMap[gt.IntervalNum2] = gt.DropId2

	//验证 interval_num_3
	err = validator.MinValidate(float64(gt.IntervalNum3), float64(1), true)
	if err != nil || gt.IntervalNum1 == gt.IntervalNum3 {
		err = fmt.Errorf("[%d] invalid", gt.IntervalNum3)
		err = template.NewTemplateFieldError("IntervalNum3", err)
		return
	}

	gt.intervalDropMap[gt.IntervalNum3] = gt.DropId3

	return nil
}

func (gt *GamblingTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(gt.FileName(), gt.TemplateId(), err)
			return
		}
	}()

	//验证 need_yinliang
	err = validator.MinValidate(float64(gt.NeedYinLiang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", gt.NeedYinLiang)
		err = template.NewTemplateFieldError("NeedYinLiang", err)
		return
	}

	//验证 need_gold
	err = validator.MinValidate(float64(gt.NeedGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", gt.NeedGold)
		err = template.NewTemplateFieldError("NeedGold", err)
		return
	}

	//验证 need_yuanshi
	err = validator.MinValidate(float64(gt.NeedYuanShi), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", gt.NeedYuanShi)
		err = template.NewTemplateFieldError("NeedYuanShi", err)
		return
	}

	return nil
}

func (gt *GamblingTemplate) PatchAfterCheck() {

}

func (gt *GamblingTemplate) FileName() string {
	return "tb_gambling.json"
}

func init() {
	template.Register((*GamblingTemplate)(nil))
}
