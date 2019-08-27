package inventory

import uipb "fgame/fgame/common/codec/pb/ui"

func buildInventoryGet(page int32) *uipb.CSInventory {
	csInventoryGet := &uipb.CSInventoryGet{}
	csInventoryGet.Page = &page
	return csInventoryGet
}

func buildInventoryUse(index int32, num int32) *uipb.CSInventoryItemUse {
	csInventoryUse := &uipb.CSInventoryItemUse{}
	csInventoryUse.Index = &index
	csInventoryUse.Num = &num
	return csInventoryUse
}

func buildInventoryMerge() *uipb.CSInventoryMerge {
	csInventoryMerge := &uipb.CSInventoryMerge{}
	return csInventoryMerge
}
