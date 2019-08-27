package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/game/shop/types"
	"fmt"
)

//商店配置
type ShopTemplate struct {
	*ShopTemplateVO
	shopConsumeType types.ShopConsumeType //类型
}

func (st *ShopTemplate) TemplateId() int {
	return st.Id
}

func (st *ShopTemplate) GetShopConsumeType() types.ShopConsumeType {
	return st.shopConsumeType
}

//同种商店购买道具优先级
func (st *ShopTemplate) Priority(old *ShopTemplate) bool {
	oldType := types.ShopConsumeType(old.ConsumeType)
	curType := types.ShopConsumeType(st.ConsumeType)
	flag := curType.Priority(oldType)
	if flag {
		return true
	}
	if st.ConsumeData1 < old.ConsumeData1 {
		return true
	}
	return false
}

func (st *ShopTemplate) GetConsumeData(num int32) (needGold, needBindGold int32, needSilver int64) {
	if num <= 0 {
		return
	}
	needGold = 0
	needBindGold = 0
	needSilver = 0
	consume := st.ConsumeData1 * num
	switch types.ShopConsumeType(st.ConsumeType) {
	case types.ShopConsumeTypeBindGold:
		{
			needBindGold = consume
		}
	case types.ShopConsumeTypeGold:
		{
			needGold = consume
		}
	case types.ShopConsumeTypeSliver:
		{
			needSilver = int64(consume)
		}
	}
	return
}

func (st *ShopTemplate) GetConsumeMoney(num int32) (costMoney int32) {
	if num <= 0 {
		return
	}
	switch types.ShopConsumeType(st.ConsumeType) {
	case types.ShopConsumeTypeBindGold:
		{
			costMoney = st.ConsumeData1 * num
		}
	case types.ShopConsumeTypeGold:
		{
			costMoney = st.ConsumeData1 * num
		}
	case types.ShopConsumeTypeSliver:
		{
			costMoney = st.ConsumeData1 * num
		}
	}
	return
}

func (st *ShopTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()

	//验证 consume_type
	st.shopConsumeType = types.ShopConsumeType(st.ConsumeType)
	if !st.shopConsumeType.Valid() {
		err = fmt.Errorf("[%d] invalid", st.ConsumeType)
		return template.NewTemplateFieldError("ConsumeType", err)
	}

	return nil
}

func (st *ShopTemplate) PatchAfterCheck() {

}

func (st *ShopTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(st.FileName(), st.TemplateId(), err)
			return
		}
	}()

	//验证 buyCount
	err = validator.MinValidate(float64(st.BuyCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.BuyCount)
		err = template.NewTemplateFieldError("BuyCount", err)
		return
	}

	//验证 maxCount
	err = validator.MinValidate(float64(st.MaxCount), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.MaxCount)
		err = template.NewTemplateFieldError("MaxCount", err)
		return
	}

	//验证 consumeData1
	err = validator.MinValidate(float64(st.ConsumeData1), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.ConsumeData1)
		err = template.NewTemplateFieldError("ConsumeData1", err)
		return
	}

	//验证
	if st.LimitCount != 0 {
		maxCount := st.BuyCount * st.MaxCount
		if maxCount > st.LimitCount {
			err = fmt.Errorf("[%d] invalid", st.MaxCount)
			err = template.NewTemplateFieldError("MaxCount", err)
			return
		}
	}

	//验证 order
	err = validator.MinValidate(float64(st.Order), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", st.Order)
		err = template.NewTemplateFieldError("Order", err)
		return
	}

	// 消耗物品
	if st.ConsumeItemId > 0 {
		if st.shopConsumeType != types.ShopConsumeTypeItem {
			err = fmt.Errorf("[%d] invalid", st.ConsumeItemId)
			err = template.NewTemplateFieldError("ConsumeItemId", err)
			return
		}

		to := template.GetTemplateService().Get(int(st.ConsumeItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", st.ConsumeItemId)
			err = template.NewTemplateFieldError("ConsumeItemId", err)
			return
		}
	}

	return nil
}

func (st *ShopTemplate) FileName() string {
	return "tb_shop.json"
}

func init() {
	template.Register((*ShopTemplate)(nil))
}
