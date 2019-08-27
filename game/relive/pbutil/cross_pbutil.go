package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSIPlayerRelive(success bool) *crosspb.SIPlayerRelive {
	siPlayerRelive := &crosspb.SIPlayerRelive{}
	siPlayerRelive.Success = &success
	return siPlayerRelive
}
