package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//消费抽奖
type GroupTemplateCostDrew struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateCostDrew) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	if len(gt.GetOpenTempMap()) != 1 {
		return fmt.Errorf("应该只有一条配置")
	}
	return
}

//每次抽奖元宝消耗
func (gt *GroupTemplateCostDrew) GetCostDrewNeedGold() int64 {
	return int64(gt.GetFirstValue1())
}

//兑换次数的元宝基数
func (gt *GroupTemplateCostDrew) GetCostDrewConvertRate() int32 {
	return gt.GetFirstValue2()
}

//每日兑换次数限制
func (gt *GroupTemplateCostDrew) GetCostDrewConvertLimit() int32 {
	return gt.GetFirstValue3()
}

//最小兑换次数
func (gt *GroupTemplateCostDrew) GetCostDrewMinCycleTimes() int32 {
	return gt.GetFirstValue4()
}

func CreateGroupTemplateCostDrew(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateCostDrew{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeCostDrew, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateCostDrew))
}
