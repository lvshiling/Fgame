package template

import (
	"fgame/fgame/core/template/validator"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//每充值奖励(次数限制)
type GroupTemplateRewardsLimit struct {
	*welfaretemplate.GroupTemplateBase //
}

func (gt *GroupTemplateRewardsLimit) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	for _, t := range gt.GetOpenTempMap() {
		// 兑换率
		err = validator.MinValidate(float64(t.Value1), float64(0), false)
		if err != nil {
			err = fmt.Errorf("Value1 [%d] invalid", t.Value1)
			err = welfaretypes.NewWelfareRecordError(t.Id, err)
			return
		}

		// 全服次数
		err = validator.MinValidate(float64(t.Value2), float64(0), false)
		if err != nil {
			err = fmt.Errorf("Value2 [%d] invalid", t.Value2)
			err = welfaretypes.NewWelfareRecordError(t.Id, err)
			return
		}

		// 个人次数
		err = validator.MinValidate(float64(t.Value3), float64(0), false)
		if err != nil {
			err = fmt.Errorf("Value3 [%d] invalid", t.Value3)
			err = welfaretypes.NewWelfareRecordError(t.Id, err)
			return
		}
	}

	return nil
}

func CreateGroupTemplateRewardsLimit(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateRewardsLimit{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeChargeLimit, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRewardsLimit))
}
