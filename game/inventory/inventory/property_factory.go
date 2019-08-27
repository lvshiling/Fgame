package inventory

import (
	babytypes "fgame/fgame/game/baby/types"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	ringtypes "fgame/fgame/game/ring/types"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fmt"
)

type PropertyFactory interface {
	CreatePropertyData(base *inventorytypes.ItemPropertyDataBase) inventorytypes.ItemPropertyData
}

type PropertyFactoryFunc func(base *inventorytypes.ItemPropertyDataBase) inventorytypes.ItemPropertyData

func (pf PropertyFactoryFunc) CreatePropertyData(base *inventorytypes.ItemPropertyDataBase) inventorytypes.ItemPropertyData {
	return pf(base)
}

var (
	propertyFactoryMap = make(map[itemtypes.ItemType]PropertyFactory)
)

func RegisterPropertyFactory(itemType itemtypes.ItemType, f PropertyFactory) {
	_, ok := propertyFactoryMap[itemType]
	if ok {
		panic(fmt.Errorf("item:repeat register property; ItemType:%d", itemType))
	}

	propertyFactoryMap[itemType] = f
}

func CreatePropertyDataInterface(typ itemtypes.ItemType, base *inventorytypes.ItemPropertyDataBase) inventorytypes.ItemPropertyData {
	f, ok := propertyFactoryMap[typ]
	if !ok {
		return base
	}

	return f.CreatePropertyData(base)
}

func init() {
	RegisterPropertyFactory(itemtypes.ItemTypeGoldEquip, PropertyFactoryFunc(goldequiptypes.CreateGoldEquipPropertyData))
	RegisterPropertyFactory(itemtypes.ItemTypeTuLongEquip, PropertyFactoryFunc(tulongequiptypes.CreateTuLongEquipPropertyData))
	RegisterPropertyFactory(itemtypes.ItemTypeBaoBaoCard, PropertyFactoryFunc(babytypes.CreateBabyPropertyData))
	RegisterPropertyFactory(itemtypes.ItemTypeTeRing, PropertyFactoryFunc(ringtypes.CreateRingPropertyData))
}
