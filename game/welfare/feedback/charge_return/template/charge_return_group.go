package template

import (
	"fgame/fgame/core/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//充值返还
type GroupTemplateChargeReturn struct {
	*welfaretemplate.GroupTemplateBase
	returnType welfaretypes.ChargeReturnType
}

func (gt *GroupTemplateChargeReturn) Init() (err error) {
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
		//培养奖励类型
		gt.returnType = welfaretypes.ChargeReturnType(t.Value1)
		if !gt.returnType.Valid() {
			err = fmt.Errorf("[%d] invalid", t.Value1)
			err = template.NewTemplateFieldError("Value1", err)
			return
		}
	}

	return nil
}

//返还类型
func (gt *GroupTemplateChargeReturn) GetReturnType() welfaretypes.ChargeReturnType {
	return gt.returnType
}

//返还比例
func (gt *GroupTemplateChargeReturn) GetReturnRatio() int32 {
	return gt.GetFirstValue2()
}

func CreateGroupTemplateChargeReturn(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateChargeReturn{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturn, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateChargeReturn))
}
