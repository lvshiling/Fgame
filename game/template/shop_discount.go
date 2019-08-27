package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/common/common"
	shoptypes "fgame/fgame/game/shop/types"
	shopdiscounttypes "fgame/fgame/game/shopdiscount/types"
	"fmt"
)

//商城促销配置
type ShopDiscountTemplate struct {
	*ShopDiscountTemplateVO
	discountType shopdiscounttypes.ShopDiscountType //打折类型类型
	shopType     shoptypes.ShopConsumeType          //商城类型
}

func (t *ShopDiscountTemplate) TemplateId() int {
	return t.Id
}

func (t *ShopDiscountTemplate) GetDiscountType() shopdiscounttypes.ShopDiscountType {
	return t.discountType
}

func (t *ShopDiscountTemplate) GetShopType() shoptypes.ShopConsumeType {
	return t.shopType
}

func (t *ShopDiscountTemplate) PatchAfterCheck() {
}

func (t *ShopDiscountTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 type
	t.discountType = shopdiscounttypes.ShopDiscountType(t.Type)
	if !t.discountType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//验证 shopType
	t.shopType = shoptypes.ShopConsumeType(t.ShopType)
	if !t.shopType.Valid() {
		err = fmt.Errorf("[%d] invalid", t.ShopType)
		err = template.NewTemplateFieldError("ShopType", err)
		return
	}

	return nil
}

func (t *ShopDiscountTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证discount
	err = validator.RangeValidate(float64(t.Discount), float64(0), false, float64(common.MAX_RATE), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.Discount)
		err = template.NewTemplateFieldError("Discount", err)
		return
	}

	return nil
}

func (t *ShopDiscountTemplate) FileName() string {
	return "tb_shop_discount.json"
}

func init() {
	template.Register((*ShopDiscountTemplate)(nil))
}
