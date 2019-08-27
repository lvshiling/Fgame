package types

import (
	"fgame/fgame/game/global"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

// 新七日投资
type NewInvestDayInfo struct {
	MaxSingleChargeNum int32           `json:"maxSingleChargeNum"` // 最大单笔充值
	ReceiveMap         map[int32]int32 `json:"receiveMap"`         // 奖励领取记录
	BuyTimeMap         map[int32]int64 `json:"buyTimeMap"`         // 购买时间
	IsEmail            bool            `json:"isEmail"`            // 是否发放邮件
}

// 是否已经购买
func (info *NewInvestDayInfo) IsEnoughMaxSigleChargeNum(goldNum int32) bool {
	if info.MaxSingleChargeNum < goldNum {
		return false
	}
	return true
}

// 是否已经购买
func (info *NewInvestDayInfo) IsAlreadyBuy(typ int32) bool {
	_, ok := info.BuyTimeMap[typ]
	if !ok {
		return false
	}
	return true
}

func (info *NewInvestDayInfo) GetCurDay(typ int32) (curDay int32, flag bool) {
	buyTime, ok := info.BuyTimeMap[typ]
	if !ok {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	diffDay, _ := timeutils.DiffDay(now, buyTime)
	curDay = diffDay + 1
	flag = true
	return
}

// 是否有可领取奖励
func (info *NewInvestDayInfo) IsCanReceiveRettot(maxRewardsDay int32) bool {
	for typ, _ := range info.ReceiveMap {
		if info.IsCanReceive(typ, maxRewardsDay) {
			return true
		}
	}
	return false
}

func (info *NewInvestDayInfo) GetReceiveByType(typ int32) (day int32, flag bool) {
	day, ok := info.ReceiveMap[typ]
	if !ok {
		return
	}
	flag = true
	return
}

// 是否可领取
func (info *NewInvestDayInfo) IsCanReceive(typ int32, maxRewardsDay int32) bool {
	buyTime, ok := info.BuyTimeMap[typ]
	if !ok {
		return false
	}
	if buyTime <= 0 {
		return false
	}

	curDay, flag := info.GetCurDay(typ)
	if !flag {
		return false
	}

	receiveDay, ok := info.ReceiveMap[typ]
	if !ok {
		receiveDay = 0
		info.ReceiveMap[typ] = receiveDay
	}

	if curDay <= receiveDay {
		return false
	}

	if receiveDay >= maxRewardsDay {
		return false
	}

	return true
}

// 是否可领取
func (info *NewInvestDayInfo) IsCanReceiveAboutEmail(typ int32, day int32) bool {
	buyTime, ok := info.BuyTimeMap[typ]
	if !ok {
		return false
	}
	if buyTime <= 0 {
		return false
	}

	receiveDay, ok := info.ReceiveMap[typ]
	if !ok {
		return false
	}

	if day <= receiveDay {
		return false
	}

	return true
}

// 购买成功，初始化
func (info *NewInvestDayInfo) InitNewSevenDayInvest(typ int32, now int64) {
	info.ReceiveMap[typ] = 0
	info.BuyTimeMap[typ] = now
}

// 领取成功，更新
func (info *NewInvestDayInfo) UpdateNewSevenDayInvest(typ int32, curDay int32) {
	info.ReceiveMap[typ] = curDay
}

// 领取成功，更新
func (info *NewInvestDayInfo) UpdateNewSevenDayInvestAboutEmail() {
	for typ, _ := range info.ReceiveMap {
		curDay, _ := info.GetCurDay(typ)
		info.ReceiveMap[typ] = curDay
	}
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewServenDay, (*NewInvestDayInfo)(nil))
}
