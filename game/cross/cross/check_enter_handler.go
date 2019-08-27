package cross

import (
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	"fmt"
)

// 跨服进入检查
type CheckEnterCrossHandler interface {
	CheckEnter(pl player.Player, crossType crosstypes.CrossType) bool
}

type CheckEnterCrossHandlerFunc func(pl player.Player, crossType crosstypes.CrossType) bool

func (f CheckEnterCrossHandlerFunc) CheckEnter(pl player.Player, crossType crosstypes.CrossType) bool {
	return f(pl, crossType)
}

var (
	checkEnterMap = make(map[crosstypes.CrossType]CheckEnterCrossHandler)
)

func RegisterCrossCheckEnterHandler(crossType crosstypes.CrossType, h CheckEnterCrossHandler) {
	_, ok := checkEnterMap[crossType]
	if ok {
		panic(fmt.Errorf("跨服进入检查器重复注册，跨服类型：%s", crossType.String()))
	}

	checkEnterMap[crossType] = h
}

func CheckEnterCross(pl player.Player, crossType crosstypes.CrossType) bool {
	h, ok := checkEnterMap[crossType]
	if !ok {
		return false
	}

	return h.CheckEnter(pl, crossType)
}
