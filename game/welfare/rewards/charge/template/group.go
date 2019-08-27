package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//每充值奖励
type GroupTemplateRewardsCharge struct {
	*welfaretemplate.GroupTemplateBase //
}

func (gt *GroupTemplateRewardsCharge) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	if len(gt.GetOpenTempMap()) != 1 {
		return fmt.Errorf("应该只有一条配置")
	}

	return nil
}

// 充值领取奖励(每消费多少领取)-兑换率
func (gt *GroupTemplateRewardsCharge) GetChargeRewardsConvertRate() int32 {
	return gt.GetFirstValue1()
}

func CreateGroupTemplateRewardsCharge(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateRewardsCharge{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeCharge, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRewardsCharge))
}
