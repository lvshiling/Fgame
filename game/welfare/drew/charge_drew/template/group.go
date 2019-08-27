package template

import (
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

//充值抽奖
type GroupTemplateChargeDrew struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateChargeDrew) Init() (err error) {
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
func (gt *GroupTemplateChargeDrew) GetChargeDrewNeedGold() int64 {
	return int64(gt.GetFirstValue1())
}

//兑换次数的元宝基数
func (gt *GroupTemplateChargeDrew) GetChargeDrewConvertRate() int32 {
	return gt.GetFirstValue2()
}

//每日兑换次数限制
func (gt *GroupTemplateChargeDrew) GetChargeDrewConvertLimit() int32 {
	return gt.GetFirstValue3()
}

//最小兑换次数
func (gt *GroupTemplateChargeDrew) GetChargeDrewMinCycleTimes() int32 {
	return gt.GetFirstValue4()
}

func CreateGroupTemplateChargeDrew(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateChargeDrew{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeChargeDrew, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateChargeDrew))
}
