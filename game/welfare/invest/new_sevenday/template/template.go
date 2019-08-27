package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

// 新七日投资
type GroupTemplateNewInvestDay struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateNewInvestDay) Init() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(gt.GetActivityName(), -1, err)
			return
		}
	}()

	// 验证
	for _, t := range gt.GetOpenTempMap() {

		err = validator.MinValidate(float64(t.Value1), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}

		err = validator.MinValidate(float64(t.Value2), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Value2)
			err = template.NewTemplateFieldError("Value2", err)
			return
		}

		err = validator.MinValidate(float64(t.Value3), float64(0), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", t.Value3)
			err = template.NewTemplateFieldError("Value3", err)
			return
		}

	}

	return
}

func (gt *GroupTemplateNewInvestDay) CheckInvestType(typ int32) bool {
	for _, temp := range gt.GetOpenTempMap() {
		if temp.Value1 == typ {
			return true
		}
	}
	return false
}

func (gt *GroupTemplateNewInvestDay) GetMaxSigleChargeNum(typ int32) int32 {
	for _, temp := range gt.GetOpenTempMap() {
		if temp.Value1 == typ {
			return temp.Value3
		}
	}
	return 0
}

//七日投资最大天数
func (gt *GroupTemplateNewInvestDay) GetInvestDayMaxRewardsLevel(typ int32) int32 {
	maxDay := int32(0)
	for _, temp := range gt.GetOpenTempMap() {
		if temp.Value1 != typ {
			continue
		}
		if maxDay < temp.Value2 {
			maxDay = temp.Value2
		}
	}
	return maxDay
}

//七日投资最大天数
func (gt *GroupTemplateNewInvestDay) GetInvestDayMaxDay() int32 {
	return gt.GetMaxValue2()
}

//七日投资-可领取奖励
func (gt *GroupTemplateNewInvestDay) GetInvestDayRewTempList(typ, maxExclude, maxInclude int32) (newTempList []*gametemplate.OpenserverActivityTemplate) {
	for _, temp := range gt.GetOpenTempMap() {
		if temp.Value1 != typ {
			continue
		}

		//领取条件
		if temp.Value2 <= maxExclude || temp.Value2 > maxInclude {
			continue
		}

		newTempList = append(newTempList, temp)
	}

	return
}

func CreateGroupTemplateNewInvestDay(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateNewInvestDay{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewServenDay, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateNewInvestDay))
}
