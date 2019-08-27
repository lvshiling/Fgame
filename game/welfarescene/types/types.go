package types

import (
	welfaretypes "fgame/fgame/game/welfare/types"
)

type WelfareSceneType int32

const (
	WelfareSceneTypeQiYuDao WelfareSceneType = iota + 1 //奇遇岛
)

func (t WelfareSceneType) Valid() bool {
	switch t {
	case WelfareSceneTypeQiYuDao:
		return true
	}

	return false
}

var (
	welfareSceneTypeStringMap = map[WelfareSceneType]string{
		WelfareSceneTypeQiYuDao: "奇遇岛",
	}
)

func (t WelfareSceneType) String() string {
	return welfareSceneTypeStringMap[t]
}

var (
	wsTypeMpa = map[welfaretypes.OpenActivityType]map[welfaretypes.OpenActivitySubType]WelfareSceneType{
		welfaretypes.OpenActivityTypeHuHu: {
			welfaretypes.OpenActivitySpecialSubTypeQiYu: WelfareSceneTypeQiYuDao,
		},
	}
)

func ConvertToWelfareSceneType(typ welfaretypes.OpenActivityType, subType welfaretypes.OpenActivitySubType) (wsType WelfareSceneType, isExit bool) {
	subMap, ok := wsTypeMpa[typ]
	if !ok {
		return
	}

	wsType, isExit = subMap[subType]
	return
}
