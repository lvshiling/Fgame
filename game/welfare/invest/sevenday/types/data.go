package types

import (
	"fgame/fgame/game/global"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

// 投资-七日
type InvestDayInfo struct {
	ReceiveDay int32 `json:"receiveDay"` //奖励领取记录
	BuyTime    int64 `json:"buyTime"`    //购买时间
}

func (info *InvestDayInfo) GetCurDay() (curDay int32) {
	now := global.GetGame().GetTimeService().Now()
	diffDay, _ := timeutils.DiffDay(now, info.BuyTime)
	curDay = diffDay + 1
	return
}

func (info *InvestDayInfo) IsCanReceive(maxRewardsDay int32) bool {
	if info.BuyTime <= 0 {
		return false
	}

	if info.GetCurDay() <= info.ReceiveDay {
		return false
	}

	if info.ReceiveDay >= maxRewardsDay {
		return false
	}

	return true
}

func (info *InvestDayInfo) UpdataBuyTime(now int64) {
	info.BuyTime = now
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeServenDay, (*InvestDayInfo)(nil))
}
