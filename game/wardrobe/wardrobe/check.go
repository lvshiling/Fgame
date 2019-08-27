package wardrobe

import (
	"fgame/fgame/game/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fmt"
)

type CheckHandler interface {
	CheckHandle(pl player.Player, seqId int32) bool
}

type CheckHandlerFunc func(pl player.Player, seqId int32) bool

func (hf CheckHandlerFunc) CheckHandle(pl player.Player, seqId int32) bool {
	return hf(pl, seqId)
}

func RegisterCheck(sysType wardrobetypes.WardrobeSysType, h CheckHandler) {
	_, exist := checkHandlerMap[sysType]
	if exist {
		panic(fmt.Sprintf("repeat register sysType %d", sysType))
	}
	checkHandlerMap[sysType] = h
}

func CheckHandle(pl player.Player, sysType wardrobetypes.WardrobeSysType, seqId int32) (flag bool) {
	h, exist := checkHandlerMap[sysType]
	if !exist {
		return
	}
	return h.CheckHandle(pl, seqId)
}

var (
	checkHandlerMap = make(map[wardrobetypes.WardrobeSysType]CheckHandler)
)
