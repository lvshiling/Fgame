package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"sort"
)

type XiuxianBookInfo struct {
	FirstTimeRewRecord int32      `json:"firstTimeRewRecord"` //活动初始时可领取的奖励
	ChargeNum          int32      `json:"chargeNum"`          //充值金额
	HasReceiveRecord   recordType `json:"hasReceiveRecord"`   //已经领取的记录 (等级)
	MaxLevel           int32      `json:"maxLevel"`           //历史最高等级
}

//分组模板排序类型
type recordType []int32

func (s recordType) Len() int           { return len(s) }
func (s recordType) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s recordType) Less(i, j int) bool { return s[i] < s[j] }

func (info *XiuxianBookInfo) GetMinCanReceiveLevel(canReceiveList []int32) int32 {
	var isCanAdd bool
	for _, addNum := range canReceiveList {
		if addNum < info.FirstTimeRewRecord {
			continue
		}
		isCanAdd = true
		for _, delNum := range info.HasReceiveRecord {
			if addNum == delNum {
				isCanAdd = false
				break
			}
		}
		if isCanAdd {
			return addNum
		}
	}
	return 0
}

// 获取当前最高等级，MaxLevel只脱下装备时候的记录最高状态
func (info *XiuxianBookInfo) GetHighestLevelInHistory(countLevel int32) int32 {
	if info.MaxLevel > countLevel {
		return info.MaxLevel
	} else {
		return countLevel
	}
}

func (info *XiuxianBookInfo) IsCanReceiveReward(needLevel int32, curLevel int32, needChargeNum int32) bool {
	if needLevel < info.FirstTimeRewRecord {
		return false
	}
	if needLevel > curLevel {
		return false
	}
	if needChargeNum > info.ChargeNum {
		return false
	}
	for _, value := range info.HasReceiveRecord {
		if needLevel == value {
			return false
		}
	}
	return true
}

func (info *XiuxianBookInfo) AddReceiveRecord(needLevel int32) {
	info.HasReceiveRecord = append(info.HasReceiveRecord, needLevel)
	sort.Sort(info.HasReceiveRecord)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipStrength, (*XiuxianBookInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipOpenLight, (*XiuxianBookInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeEquipUpStar, (*XiuxianBookInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeLingTong, (*XiuxianBookInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeDianXing, (*XiuxianBookInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeShenQi, (*XiuxianBookInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillXinFa, (*XiuxianBookInfo)(nil))
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeXiuxianBook, welfaretypes.OpenActivityXiuxianBookSubTypeSkillDiHun, (*XiuxianBookInfo)(nil))
}
