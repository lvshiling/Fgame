package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//每消费奖励
type GroupTemplateRewardsCost struct {
	*welfaretemplate.GroupTemplateBase //
}

func (gt *GroupTemplateRewardsCost) Init() (err error) {
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

// 消费领取奖励(每消费多少领取)-兑换率
func (gt *GroupTemplateRewardsCost) GetCostRewardsConvertRate() int32 {
	return gt.GetFirstValue1()
}

func CreateGroupTemplateRewardsCost(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateRewardsCost{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeCost, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateRewardsCost))
}
