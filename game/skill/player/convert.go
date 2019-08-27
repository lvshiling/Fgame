package player

import (
	"fgame/fgame/game/skill/common"
)

func convertTianFu(tianFuMap map[int32]int32) (tianFuList []*common.TianFuInfo) {
	tianFuList = make([]*common.TianFuInfo, 0, 3)
	for tianFuId, level := range tianFuMap {
		tianFuInfo := &common.TianFuInfo{
			TianFuId: tianFuId,
			Level:    level,
		}
		tianFuList = append(tianFuList, tianFuInfo)
	}
	return
}
