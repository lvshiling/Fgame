package template

import (
	"fgame/fgame/game/common/common"
	gametemplate "fgame/fgame/game/template"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"math"
)

//房产投资
type GroupTemplateHouseInvest struct {
	*welfaretemplate.GroupTemplateBase
}

func CreateGroupTemplateHouseInvest(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	g := &GroupTemplateHouseInvest{}
	g.GroupTemplateBase = base
	return g
}

//房产投资配置
func (gt *GroupTemplateHouseInvest) GetOpenActivityHouseInvest(days int32) *gametemplate.OpenserverActivityTemplate {
	for _, temp := range gt.GetOpenTempMap() {
		if temp.Value1 == days {
			return temp
		}
	}
	return nil
}

//房产投资需要的充值数
func (gt *GroupTemplateHouseInvest) GetOpenActivityHouseInvestChargeNum(days int32) int32 {
	for _, temp := range gt.GetOpenTempMap() {
		if temp.Value1 == days {
			return temp.Value2
		}
	}
	return 0
}

//房产投资需要的充值数
func (gt *GroupTemplateHouseInvest) GetOpenActivityHouseInvestSellNum(days int32) int64 {
	for _, temp := range gt.GetOpenTempMap() {
		if temp.Value1 == days {
			sellNum := int64(math.Ceil(float64(temp.Value3) * float64(temp.Value4) / float64(common.MAX_RATE)))
			return sellNum
		}
	}
	return 0
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseInvest, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplateHouseInvest))
}
