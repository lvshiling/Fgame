package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildISMassacreDrop(attackId int64, attackName string) *crosspb.ISMassacreDrop {
	massacreDrop := &crosspb.ISMassacreDrop{}
	massacreDrop.AttackerId = &attackId
	massacreDrop.AttackerName = &attackName
	return massacreDrop
}
