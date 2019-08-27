package template

import (
	"fgame/fgame/core/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//充值返还
type GroupTemplateChargeReturnLevel struct {
	*welfaretemplate.GroupTemplateBase
	returnType welfaretypes.ChargeReturnType
}

func (gt *GroupTemplateChargeReturnLevel) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	// 返还类型
	firstTemp := gt.GetFirstOpenTemp()
	if firstTemp != nil {
		gt.returnType = welfaretypes.ChargeReturnType(firstTemp.Value1)
	}

	for _, t := range gt.GetOpenTempMap() {
		//校验 返还类型
		returnType := welfaretypes.ChargeReturnType(t.Value1)
		if !returnType.Valid() || gt.returnType != returnType {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}
	}

	return nil
}

//返还类型
func (gt *GroupTemplateChargeReturnLevel) GetReturnType() welfaretypes.ChargeReturnType {
	return gt.returnType
}

//返还比例
func (gt *GroupTemplateChargeReturnLevel) GetReturnRatio(chargeNum int32) int32 {
	for _, temp := range gt.GetOpenTempMap() {
		rangeMin := temp.Value2
		rangeMax := temp.Value3
		if chargeNum < rangeMin || chargeNum > rangeMax {
			continue
		}

		return temp.Value4
	}
	return 0
}

func CreateGroupTemplateChargeReturnLevel(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateChargeReturnLevel{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnLevel, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateChargeReturnLevel))
}
