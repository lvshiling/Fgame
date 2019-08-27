package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSIQiXueDrop(itemId int32, itemNum int64, attackId int64) *crosspb.SIQiXueDrop {
	qixueDrop := &crosspb.SIQiXueDrop{}
	qixueDrop.ItemId = &itemId
	qixueDrop.ItemNum = &itemNum
	qixueDrop.AttackerId = &attackId

	return qixueDrop
}
