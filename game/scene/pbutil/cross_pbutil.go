package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSIPlayerKillBiology(biologyId int32) *crosspb.SIPlayerKillBiology {
	siPlayerKillBiology := &crosspb.SIPlayerKillBiology{}
	siPlayerKillBiology.BiologyId = &biologyId
	return siPlayerKillBiology
}
