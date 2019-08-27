package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

type GroupTemplateDiscountYunYinShop struct {
	*welfaretemplate.GroupTemplateBase
}

func (g *GroupTemplateDiscountYunYinShop) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(g.GetActivityName(), g.GetGroupId(), err)
			return
		}
	}()

	for _, temp := range g.GetOpenTempMap() {
		// 验证value_1
		err = validator.MinValidate(float64(temp.Value1), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", temp.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}

	}
	return
}

func (t *GroupTemplateDiscountYunYinShop) GetCanReceiveRewardTemplateList(goldNum int32) (costTempList []*gametemplate.OpenserverActivityTemplate) {
	openTempMap := t.GetOpenTempMap()
	for _, openTemp := range openTempMap {
		if openTemp.Value1 <= goldNum {
			costTempList = append(costTempList, openTemp)
		}
	}
	return
}

func (t *GroupTemplateDiscountYunYinShop) GetCanReceiveRewardList(goldNum int32) (costList []int32) {
	openTempMap := t.GetOpenTempMap()
	for _, openTemp := range openTempMap {
		if openTemp.Value1 <= goldNum {
			costList = append(costList, openTemp.Value1)
		}
	}
	return
}

func CreateGroupTemplateDiscountYunYinShop(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	t := &GroupTemplateDiscountYunYinShop{}
	t.GroupTemplateBase = base
	return t
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeYunYin, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateDiscountYunYinShop))
}
