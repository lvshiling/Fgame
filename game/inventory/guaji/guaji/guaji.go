package guaji

import (
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fmt"
)

//挂机自动使用
type GuaJiItemAutoUseHandler interface {
	AutoUseItem(pl player.Player, index int32, num int32)
}

type GuaJiItemAutoUseHandlerFunc func(pl player.Player, index int32, num int32)

func (f GuaJiItemAutoUseHandlerFunc) AutoUseItem(pl player.Player, index int32, num int32) {
	f(pl, index, num)
}

var (
	guaJiItemAutoUseMap = make(map[itemtypes.ItemType]map[itemtypes.ItemSubType]GuaJiItemAutoUseHandler)
)

func GetGuaJiItemAutoUseHandler(itemType itemtypes.ItemType, itemSubType itemtypes.ItemSubType) GuaJiItemAutoUseHandler {
	tempMap, ok := guaJiItemAutoUseMap[itemType]
	if !ok {
		return nil
	}
	h, ok := tempMap[itemSubType]
	if !ok {
		return nil
	}
	return h
}

func RegisterGuaJiItemAutoUseHandler(itemType itemtypes.ItemType, itemSubType itemtypes.ItemSubType, h GuaJiItemAutoUseHandler) {
	tempMap, exist := guaJiItemAutoUseMap[itemType]
	if !exist {
		tempMap = make(map[itemtypes.ItemSubType]GuaJiItemAutoUseHandler)
		guaJiItemAutoUseMap[itemType] = tempMap
	}

	_, exist = tempMap[itemSubType]
	if exist {
		panic(fmt.Errorf("重复注册%s挂机物品自动使用处理器", itemType.String()))
	}

	tempMap[itemSubType] = h
}

//类型使用
type GuaJiItemUseHandler interface {
	UseItem(pl player.Player, index int32, num int32)
}

type GuaJiItemUseHandlerFunc func(pl player.Player, index int32, num int32)

func (f GuaJiItemUseHandlerFunc) UseItem(pl player.Player, index int32, num int32) {
	f(pl, index, num)
}

var (
	guaJiItemUseMap = make(map[itemtypes.ItemType]GuaJiItemUseHandler)
)

func GetGuaJiItemUseHandler(itemType itemtypes.ItemType) GuaJiItemUseHandler {
	h, ok := guaJiItemUseMap[itemType]
	if !ok {
		return nil
	}

	return h
}

func RegisterGuaJiItemUseHandler(itemType itemtypes.ItemType, h GuaJiItemUseHandler) {
	_, exist := guaJiItemUseMap[itemType]
	if exist {
		panic(fmt.Errorf("重复注册%s挂机物品使用处理器", itemType.String()))
	}
	guaJiItemUseMap[itemType] = h
}
