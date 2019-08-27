package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//七日投资
type GroupTemplateInvestDay struct {
	*welfaretemplate.GroupTemplateBase
	buyNeedGold int32
}

func (gt *GroupTemplateInvestDay) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	//购买所需元宝
	gt.buyNeedGold = gt.GetFirstValue2()

	for _, t := range gt.GetOpenTempMap() {
		//七日
		err = validator.MinValidate(float64(t.Value1), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}

		//校验 购买所需元宝
		if t.Value2 != gt.buyNeedGold {
			err = fmt.Errorf("[%d] invalid", t.Value2)
			err = template.NewTemplateFieldError("Value2", err)
			return
		}

	}

	return
}

// 七日投资所需元宝
func (gt *GroupTemplateInvestDay) GetInvestDayNeedGold() int32 {
	return gt.buyNeedGold
}

//七日投资最大天数
func (gt *GroupTemplateInvestDay) GetInvestDayMaxRewardsLevel() int32 {
	return gt.GetMaxValue1()
}

//七日投资-可领取奖励
func (gt *GroupTemplateInvestDay) GetInvestDayRewTempList(maxExclude, maxInclude int32) (newTempList []*gametemplate.OpenserverActivityTemplate) {
	for _, temp := range gt.GetOpenTempMap() {
		//领取条件
		if temp.Value1 <= maxExclude || temp.Value1 > maxInclude {
			continue
		}

		newTempList = append(newTempList, temp)
	}

	return
}

func CreateGroupTemplateInvestDay(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateInvestDay{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeServenDay, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateInvestDay))
}
