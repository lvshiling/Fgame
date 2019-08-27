package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildISQiXueDrop(attackId int64, attackName string) *crosspb.ISQiXueDrop {
	qixueDrop := &crosspb.ISQiXueDrop{}
	qixueDrop.AttackerId = &attackId
	qixueDrop.AttackerName = &attackName
	return qixueDrop
}
