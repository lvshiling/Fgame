package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildISDropItemGet(itemId int32, num int32, level int32, attrList []int32, upstar int32) *crosspb.ISPlayerGetDropItem {
	isGetDropItem := &crosspb.ISPlayerGetDropItem{}
	isGetDropItem.ItemId = &itemId
	isGetDropItem.Num = &num
	isGetDropItem.Level = &level
	isGetDropItem.AttrList = attrList
	isGetDropItem.Upstar = &upstar
	return isGetDropItem
}
