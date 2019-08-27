package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//每累充
type GroupTemplateChargeReturnMultiple struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateChargeReturnMultiple) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	if len(gt.GetOpenTempMap()) != 1 {
		return fmt.Errorf("应该只有一条配置")
	}

	//验证 value1
	err = validator.MinValidate(float64(gt.GetFirstValue1()), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", gt.GetFirstValue1())
		err = template.NewTemplateFieldError("Value1", err)
		return
	}

	//验证 value2
	err = validator.MinValidate(float64(gt.GetFirstValue2()), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", gt.GetFirstValue2())
		err = template.NewTemplateFieldError("Value2", err)
		return
	}

	return nil
}

//每累计充值数
func (gt *GroupTemplateChargeReturnMultiple) GetPerChargeNum() int32 {
	return gt.GetFirstValue1()
}

//领取上限
func (gt *GroupTemplateChargeReturnMultiple) GetRewardLimitCnt() int32 {
	return gt.GetFirstValue2()
}

func CreateGroupTemplateChargeReturnMultiple(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateChargeReturnMultiple{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateChargeReturnMultiple))
}
