package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

var (
	isPlayerExitCross = &crosspb.ISPlayerExitCross{}
)

func BuildISPlayerExitCross() *crosspb.ISPlayerExitCross {
	return isPlayerExitCross
}

func BuildISPlayerKillBiology(monsterId int32) *crosspb.ISPlayerKillBiology {
	iSPlayerKillBiology := &crosspb.ISPlayerKillBiology{}
	iSPlayerKillBiology.BiologyId = &monsterId
	return iSPlayerKillBiology
}
