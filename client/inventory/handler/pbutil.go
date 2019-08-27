package handler

import "fgame/fgame/client/inventory"
import uipb "fgame/fgame/common/codec/pb/ui"

func convertFromSlotItems(items []*uipb.SlotItem) (itemObjects []*inventory.ItemObject) {
	for _, item := range items {
		itemObject := convertFromSlotItem(item)
		itemObjects = append(itemObjects, itemObject)
	}
	return itemObjects
}

func convertFromSlotItem(item *uipb.SlotItem) *inventory.ItemObject {
	itemObject := &inventory.ItemObject{}
	itemObject.ItemId = item.GetItemId()
	itemObject.Index = item.GetIndex()
	itemObject.Num = item.GetNum()
	return itemObject
}
