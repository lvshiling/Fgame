package drop

import (
	commonlogtypes "fgame/fgame/common/log/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fmt"
)

type DropResHandler interface {
	AddRes(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool
}

type DropResHandlerFunc func(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool

func (f DropResHandlerFunc) AddRes(pl player.Player, resNum int64, reason commonlogtypes.LogReason, reasonText string) bool {
	return f(pl, resNum, reason, reasonText)
}

var (
	resHandlerMap = make(map[itemtypes.ItemAutoUseResSubType]DropResHandler)
)

func RegistDropResHandler(resType itemtypes.ItemAutoUseResSubType, h DropResHandler) {
	_, ok := resHandlerMap[resType]
	if ok {
		panic(fmt.Errorf("掉落资源处理器类型已经注册，resType:%d", resType))
	}

	resHandlerMap[resType] = h
}

func GetDropResHandler(resType itemtypes.ItemAutoUseResSubType) DropResHandler {
	h, ok := resHandlerMap[resType]
	if !ok {
		return nil
	}

	return h
}
