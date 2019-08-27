package inventory

import (
	"encoding/json"
	babytypes "fgame/fgame/game/baby/types"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	ringtypes "fgame/fgame/game/ring/types"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fmt"
	"reflect"
)

var (
	propertyDataMap     = make(map[itemtypes.ItemType]reflect.Type)
	defaultPropertyType = reflect.TypeOf((*inventorytypes.ItemPropertyDataBase)(nil)).Elem()
)

func RegisterPropertyData(itemType itemtypes.ItemType, od inventorytypes.ItemPropertyData) {
	_, ok := propertyDataMap[itemType]
	if ok {
		panic(fmt.Errorf("item:repeat register property; ItemType:%d", itemType))
	}

	propertyDataMap[itemType] = reflect.TypeOf(od).Elem()
}

func CreatePropertyData(typ itemtypes.ItemType, content string) (data inventorytypes.ItemPropertyData, err error) {
	dataType, ok := propertyDataMap[typ]
	if !ok {
		dataType = defaultPropertyType
	}

	x := reflect.New(dataType)
	err = json.Unmarshal([]byte(content), x.Interface())
	if err != nil {
		return
	}
	data = x.Interface().(inventorytypes.ItemPropertyData)
	data.InitBase()
	return
}

func init() {
	RegisterPropertyData(itemtypes.ItemTypeGoldEquip, (*goldequiptypes.GoldEquipPropertyData)(nil))
	RegisterPropertyData(itemtypes.ItemTypeTuLongEquip, (*tulongequiptypes.TuLongEquipPropertyData)(nil))
	RegisterPropertyData(itemtypes.ItemTypeBaoBaoCard, (*babytypes.BabyPropertyData)(nil))
	RegisterPropertyData(itemtypes.ItemTypeTeRing, (*ringtypes.RingPropertyData)(nil))
}
