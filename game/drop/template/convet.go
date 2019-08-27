package template

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
)

func ConvertToItemDataList(itemMap map[int32]int32, bindType itemtypes.ItemBindType) (dataList []*DropItemData) {
	level := int32(0)
	itemGetTime := global.GetGame().GetTimeService().Now()
	for itemId, num := range itemMap {
		itemTemp := item.GetItemService().GetItem(int(itemId))
		if itemTemp == nil {
			continue
		}
		expireTime := itemTemp.GetExpireTime()
		expireType := itemTemp.GetLimitTimeType()
 
		dataList = append(dataList, CreateItemDataWithExpire(itemId, num, level, bindType, expireType, expireTime, itemGetTime))
	}
	return
}

func ConvertToItemDataListDefault(itemMap map[int32]int32) (dataList []*DropItemData) {
	return ConvertToItemDataList(itemMap, itemtypes.ItemBindTypeUnBind)
}
