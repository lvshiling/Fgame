package template

import (
	"fgame/fgame/game/common/common"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"math"
)

type GroupTemplateChargeArenapvpAssistReturn struct {
	*welfaretemplate.GroupTemplateBase
}

func (gt *GroupTemplateChargeArenapvpAssistReturn) Init() (err error) {
	defer func() {
		if err != nil {
			err = welfaretypes.NewWelfareError(gt.GetActivityName(), gt.GetGroupId(), err)
			return
		}
	}()

	return
}
func (gt *GroupTemplateChargeArenapvpAssistReturn) GetReturnGoldNum(rank int32, costGold int64) (goldnum int32, rateShow int32, maxGoldNum int32) {
	for _, temp := range gt.GetOpenTempMap() {
		if int32(temp.Value1) <= rank && int32(temp.Value2) >= rank {
			rate := int32(temp.Value3)
			rateShow := rate / 100
			maxGoldNum := int32(temp.Value4)
			returnGoldNum := int32(math.Ceil(float64(costGold) * float64(rate) / float64(common.MAX_RATE)))
			if returnGoldNum >= maxGoldNum {
				returnGoldNum = maxGoldNum
				return returnGoldNum, rateShow, maxGoldNum
			} else {
				return returnGoldNum, rateShow, maxGoldNum
			}
		}
	}
	return
}

func CreateGroupTemplate(base *welfaretemplate.GroupTemplateBase) welfaretemplate.GroupTemplateI {
	gt := &GroupTemplateChargeArenapvpAssistReturn{}
	gt.GroupTemplateBase = base
	return gt
}

func init() {
	welfaretemplate.RegisterGroupTemplate(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeWuLian, welfaretemplate.GroupTemplateIFactoryFunc(CreateGroupTemplate))
}