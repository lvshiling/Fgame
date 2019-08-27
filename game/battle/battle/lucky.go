package battle

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/scene/scene"
)

type LuckyData struct {
	itemType    itemtypes.ItemType
	itemSubType itemtypes.ItemSubType
	rate        int32
	expireTime  int64
}

// 幸运系统
type PlayerLuckyManager struct {
	p         scene.Player
	luckyList []*LuckyData
}

func (m *PlayerLuckyManager) GetLuckyRate(typ itemtypes.ItemType, subType itemtypes.ItemSubType) int32 {
	index, data := m.getLucky(typ, subType)
	if data == nil {
		return 0
	}

	now := global.GetGame().GetTimeService().Now()
	if now > data.expireTime {
		m.delLucky(index)
		return 0
	}

	return data.rate
}

func (m *PlayerLuckyManager) AddLucky(typ itemtypes.ItemType, subType itemtypes.ItemSubType, rate int32, expireTime int64) {
	_, data := m.getLucky(typ, subType)
	if data != nil {
		data.expireTime = expireTime
		data.rate = rate
	} else {
		d := &LuckyData{
			itemType:    typ,
			itemSubType: subType,
			rate:        rate,
			expireTime:  expireTime,
		}

		m.luckyList = append(m.luckyList, d)
	}
}

func (m *PlayerLuckyManager) getLucky(typ itemtypes.ItemType, subType itemtypes.ItemSubType) (int, *LuckyData) {
	for index, data := range m.luckyList {
		if data.itemType == typ && data.itemSubType == subType {
			return index, data
		}
	}
	return -1, nil
}

func (m *PlayerLuckyManager) delLucky(index int) {
	newList := make([]*LuckyData, 0, 1)
	newList = append(newList, m.luckyList[:index]...)
	newList = append(newList, m.luckyList[index+1:]...)
	m.luckyList = newList
}

func CreatePlayerLuckyManagerWithData(p scene.Player, luckyMap map[int32]int64) *PlayerLuckyManager {
	m := &PlayerLuckyManager{
		p: p,
	}

	for itemId, expireTime := range luckyMap {
		itemTemplate := item.GetItemService().GetItem(int(itemId))
		typ := itemTemplate.GetItemType()
		subType := itemTemplate.GetItemSubType()
		rate := itemTemplate.TypeFlag1

		m.AddLucky(typ, subType, rate, expireTime)
	}

	return m
}

func CreatePlayerLuckyManager(p scene.Player) *PlayerLuckyManager {
	m := &PlayerLuckyManager{
		p: p,
	}

	return m
}
