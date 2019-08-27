package found

import (
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	"fmt"
)

//资源找回信息
type FoundObjDataHandler interface {
	GetFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32)
}

type FoundObjDataHandlerFunc func(pl player.Player) (resLevel int32, maxTimes int32, group int32)

func (f FoundObjDataHandlerFunc) GetFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return f(pl)
}

func RegistFoundDataHandler(resType foundtypes.FoundResourceType, data FoundObjDataHandler) {
	_, ok := foundMap[resType]
	if ok {
		panic("found:重复注册资源找回处理器")
	}

	foundMap[resType] = data
}

func GetFoundDataHandler(resType foundtypes.FoundResourceType) FoundObjDataHandler {
	h, ok := foundMap[resType]
	if !ok {
		return nil
	}
	return h
}

var (
	foundMap = make(map[foundtypes.FoundResourceType]FoundObjDataHandler)
)

// 资源找回条件判断
type FoundCheckHandler interface {
	IsCanFoundBack(pl player.Player) (flag bool)
}

type FoundCheckHandlerFunc func(pl player.Player) (flag bool)

func (f FoundCheckHandlerFunc) IsCanFoundBack(pl player.Player) (flag bool) {
	return f(pl)
}

var (
	foundCheckHandlerMap = make(map[foundtypes.FoundResourceType]FoundCheckHandler)
	defaultH             = FoundCheckHandlerFunc(defaultHandler)
)

func RegistFoundCheckHandler(resType foundtypes.FoundResourceType, h FoundCheckHandler) {
	_, ok := foundCheckHandlerMap[resType]
	if ok {
		panic(fmt.Errorf("found:重复注册资源找回判断处理器,资源类型：%d", resType))
	}

	foundCheckHandlerMap[resType] = h
}

func defaultHandler(pl player.Player) bool {
	return true
}

func GetFoundCheckHandler(resType foundtypes.FoundResourceType) FoundCheckHandler {
	h, ok := foundCheckHandlerMap[resType]
	if !ok {
		return defaultH
	}
	return h
}
