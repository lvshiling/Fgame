package player

import (
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fmt"
)

type ItemUseHandler interface {
	Use(pl player.Player, it *PlayerItemObject, useNum int32, chooseIndexList []int32, args string) (flag bool, err error)
}

type ItemUseHandleFunc func(pl player.Player, it *PlayerItemObject, useNum int32, chooseIndexList []int32, args string) (flag bool, err error)

func (iuhf ItemUseHandleFunc) Use(pl player.Player, it *PlayerItemObject, useNum int32, chooseIndexList []int32, args string) (flag bool, err error) {
	return iuhf(pl, it, useNum, chooseIndexList, args)
}

var (
	itemUseMap = make(map[itemtypes.ItemType]map[itemtypes.ItemSubType]ItemUseHandler)
)

func RegisterUseHandler(itemType itemtypes.ItemType, itemSubType itemtypes.ItemSubType, h ItemUseHandler) {
	itemUseSubMap, exist := itemUseMap[itemType]
	if !exist {
		itemUseSubMap = make(map[itemtypes.ItemSubType]ItemUseHandler)
		itemUseMap[itemType] = itemUseSubMap
	}

	_, exist = itemUseSubMap[itemSubType]
	if exist {
		panic(fmt.Errorf("repeate register %s type use handler", itemType.String()))
	}

	itemUseSubMap[itemSubType] = h
}

func GetUseHandler(itemType itemtypes.ItemType, itemSubType itemtypes.ItemSubType) ItemUseHandler {
	itemUseSubType, exist := itemUseMap[itemType]
	if !exist {
		return nil
	}

	h, exist := itemUseSubType[itemSubType]
	if !exist {
		return nil
	}
	return h
}
