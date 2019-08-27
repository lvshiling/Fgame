package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSIXueChiAdd(blood int64) *crosspb.SIXueChiAdd {
	siXueChiAdd := &crosspb.SIXueChiAdd{}
	siXueChiAdd.Blood = &blood
	return siXueChiAdd
}
