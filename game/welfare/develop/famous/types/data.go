package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//培养名人
type DevelopFameInfo struct {
	FavorableNum    int32           `json:"favorableNum"`    //好感度
	FeedTimesMap    map[int32]int32 `json:"feedTimesMap"`    //物品使用次数map,物品:次数
	RewRecord       []int32         `json:"rewRecord"`       //奖励领取记录
	IsEmail         bool            `json:"isEmail"`         //是否邮件
	DayFavorableNum int32           `json:"dayFavorableNum"` //每日好感度
}

func (info *DevelopFameInfo) IsCanReceiveRewards(needFavorableNum int32) bool {
	if info.FavorableNum < needFavorableNum {
		return false
	}
	//领取记录
	for _, value := range info.RewRecord {
		if value == needFavorableNum {
			return false
		}
	}

	return true
}

func (info *DevelopFameInfo) IfCanFeed(feedItemId, addTimes, timesLimit int32) bool {
	if timesLimit == 0 {
		return true
	}
	feedTimes := info.FeedTimesMap[feedItemId]
	if feedTimes+addTimes <= timesLimit {
		return true
	}

	return false
}

func (info *DevelopFameInfo) AddRecord(needFavorableNum int32) {
	info.RewRecord = append(info.RewRecord, needFavorableNum)
}

func (info *DevelopFameInfo) AddFeedTimes(feedItemId, addTimes int32) {
	_, ok := info.FeedTimesMap[feedItemId]
	if !ok {
		info.FeedTimesMap[feedItemId] = addTimes
	} else {
		info.FeedTimesMap[feedItemId] += addTimes
	}
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeDevelop, welfaretypes.OpenActivityDefaultSubTypeDefault, (*DevelopFameInfo)(nil))
}
