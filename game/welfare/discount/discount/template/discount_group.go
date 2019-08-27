package template

import (
	"fgame/fgame/core/template"
	inventorytypes "fgame/fgame/game/inventory/types"
	discountdiscounttypes "fgame/fgame/game/welfare/discount/discount/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//限时折扣礼包
type GroupTemplateDiscount struct {
	*welfaretemplate.GroupTemplateBase
	limitType discountdiscounttypes.TimesLimitType
}

func (gt *GroupTemplateDiscount) Init() (err error) {
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
		gt.limitType = discountdiscounttypes.TimesLimitType(t.Value1)
		if !gt.limitType.Valid() {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}
	}
	return nil
}

// 折扣商店是否全局次数
func (gt *GroupTemplateDiscount) IsGlobalTimesLimit() bool {
	if gt.limitType == discountdiscounttypes.TimesLimitTypeGlobal {
		return true
	}

	return false
}

//获取过期类型
func (gt *GroupTemplateDiscount) GetExpireType() inventorytypes.NewItemLimitTimeType {
	return gt.GetFirstOpenTemp().GetExpireType()
}

//获取过期时间
func (gt *GroupTemplateDiscount) GetExpireTime() int64 {
	return gt.GetFirstOpenTemp().GetExpireTime()
}

func CreateGroupTemplateDiscount(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateDiscount{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeCommon, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateDiscount))
}
