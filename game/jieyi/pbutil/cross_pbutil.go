package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSIShengWeiDrop(itemId int32, itemNum int64, attackId int64) *crosspb.SIShengWeiDrop {
	siShengWeiDrop := &crosspb.SIShengWeiDrop{}
	siShengWeiDrop.ItemId = &itemId
	siShengWeiDrop.ItemNum = &itemNum
	siShengWeiDrop.AttackerId = &attackId

	return siShengWeiDrop
}
