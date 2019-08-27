package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	shopdiscounttypes "fgame/fgame/game/shopdiscount/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//商店特权
type GroupTemplateShopDiscount struct {
	*welfaretemplate.GroupTemplateBase
	sdType shopdiscounttypes.ShopDiscountType
}

func (gt *GroupTemplateShopDiscount) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	if len(gt.GetOpenTempMap()) != 1 {
		return fmt.Errorf("应该只有一条配置")
	}

	for _, t := range gt.GetOpenTempMap() {
		gt.sdType = shopdiscounttypes.ShopDiscountType(t.Value1)
		if !gt.sdType.Valid() {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}
		err = validator.MinValidate(float64(t.Value2), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Value2)
			err = template.NewTemplateFieldError("Value2", err)
			return
		}
	}

	return
}

//特权类型
func (gt *GroupTemplateShopDiscount) GetShopDiscountType() shopdiscounttypes.ShopDiscountType {
	return gt.sdType
}

//需要充值数
func (gt *GroupTemplateShopDiscount) GetShopDiscountNeedChargeNum() int32 {
	return gt.GetFirstValue2()
}

func CreateGroupTemplateShopDiscount(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateShopDiscount{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeShopDiscount, welfaretypes.OpenActivityDefaultSubTypeDefault, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateShopDiscount))
}
