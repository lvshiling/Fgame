package types

import (
	huiyuantypes "fgame/fgame/game/huiyuan/types"
)

type HuiYuanEventType string

const (
	EventTypeHuiYuanRewards HuiYuanEventType = "HuiYuanRewads" //会员奖励
	EventTypeHuiYuanBuy                      = "HuiYuanBuy"    //会员购买
)

type HuiYuanRewardsEventData struct {
	huiyuanType huiyuantypes.HuiYuanType //会员类型
	dayNumber   int32                    //奖励天数
	buyTime     int64                    //购买时间
}

func CreateHuiYuanRewardsEventData(typ huiyuantypes.HuiYuanType, num int32, buyTime int64) *HuiYuanRewardsEventData {
	data := &HuiYuanRewardsEventData{
		huiyuanType: typ,
		dayNumber:   num,
		buyTime:     buyTime,
	}

	return data
}

func (d *HuiYuanRewardsEventData) GetHuiYuanType() huiyuantypes.HuiYuanType {
	return d.huiyuanType
}

func (d *HuiYuanRewardsEventData) GetRewardsDayNum() int32 {
	return d.dayNumber
}

func (d *HuiYuanRewardsEventData) GetBuyTime() int64 {
	return d.buyTime
}
