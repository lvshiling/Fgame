package types

import (
	groupcollectenum "fgame/fgame/game/welfare/group/collect/enum"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/mathutils"
)

// 次数奖励
type CollectRewInfo struct {
	HadPokerList  []int32                      `json:"hadPokerList"`  //已收集的卡牌
	NonePokerList []int32                      `json:"nonePokerList"` //未收集的卡牌
	RewRecord     []groupcollectenum.PokerType `json:"rewRecord"`     //奖励记录
}

func (info *CollectRewInfo) AddPoker() int32 {
	min := int(0)
	max := len(info.NonePokerList)
	if max == 0 {
		info.NonePokerList = groupcollectenum.GetInitPokerList()
		info.HadPokerList = []int32{}
		info.RewRecord = []groupcollectenum.PokerType{}
	}
	randomNum := mathutils.RandomRange(min, max)
	poker := info.NonePokerList[randomNum]
	info.HadPokerList = append(info.HadPokerList, poker)
	info.NonePokerList = append(info.NonePokerList[:randomNum], info.NonePokerList[randomNum+1:]...)
	return poker
}

func (info *CollectRewInfo) CheckCollect() (bool, groupcollectenum.PokerType) {

	noneCollectMap := make(map[groupcollectenum.PokerType]int32)
	for _, poker := range info.HadPokerList {
		pokerType, _ := groupcollectenum.MaskDecode(poker)
		if info.isFinishCollect(pokerType) {
			continue
		}

		// 花色计数
		_, ok := noneCollectMap[pokerType]
		if !ok {
			noneCollectMap[pokerType] = 1
		} else {
			noneCollectMap[pokerType] += 1
		}
	}

	for pokerType, collectNum := range noneCollectMap {
		if !groupcollectenum.IsFinishCollect(pokerType, collectNum) {
			continue
		}

		return true, pokerType
	}

	return false, -1
}

func (info *CollectRewInfo) isFinishCollect(pokerType groupcollectenum.PokerType) bool {
	for _, rewPokerType := range info.RewRecord {
		if rewPokerType != pokerType {
			continue
		}
		return true
	}

	return false
}

func (info *CollectRewInfo) AddRewRecord(pokerType groupcollectenum.PokerType) {
	info.RewRecord = append(info.RewRecord, pokerType)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeGroup, welfaretypes.OpenActivityGroupSubTypeCollectPoker, (*CollectRewInfo)(nil))
}
