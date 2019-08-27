package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//转生大礼包
type GroupTemplateZhaunSheng struct {
	*welfaretemplate.GroupTemplateBase
	gitfType welfaretypes.ZhuanShengGiftValue1Type
}

func (gt *GroupTemplateZhaunSheng) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	if len(gt.GetOpenTempMap()) != 1 {
		return fmt.Errorf("应该只有一条配置")
	}

	// 礼包商城类型
	firstTemp := gt.GetFirstOpenTemp()
	gt.gitfType = welfaretypes.ZhuanShengGiftValue1Type(firstTemp.Value1)

	if gt.GetType() == welfaretypes.OpenActivityTypeDiscount && gt.GetSubType() == welfaretypes.OpenActivityDiscountSubTypeZhuanSheng {
		switch gt.gitfType {
		case welfaretypes.ZhuanShengGiftValue1TypePoint:
			{
				//积分商城
				//验证 value1
				err = validator.MinValidate(float64(gt.GetFirstValue2()), float64(1), true)
				if err != nil {
					err = fmt.Errorf("[%d] invalid", gt.GetFirstValue2())
					err = template.NewTemplateFieldError("Value2", err)
					return
				}

				//验证 value2
				err = validator.MinValidate(float64(gt.GetFirstValue3()), float64(1), true)
				if err != nil {
					err = fmt.Errorf("[%d] invalid", gt.GetFirstValue3())
					err = template.NewTemplateFieldError("Value3", err)
					return
				}
			}
		}
	}

	return nil
}

//获取礼包商城类型
func (gt *GroupTemplateZhaunSheng) GetGiftValue1Type() welfaretypes.ZhuanShengGiftValue1Type {
	return gt.gitfType
}

//获取总充值积分
func (gt *GroupTemplateZhaunSheng) GetTotalAndRemainPoint(chargeNum int64, usePoint int32) (totalPoint int32, leftPoint int32) {
	if gt.GetFirstValue2() <= 0 {
		return
	}
	totalVal := chargeNum / int64(gt.GetFirstValue2())
	totalPoint = gt.GetFirstValue3() * int32(totalVal)
	leftPoint = totalPoint - usePoint
	return
}

func CreateGroupTemplateZhaunSheng(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateZhaunSheng{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeZhuanSheng, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateZhaunSheng))
}
