package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildISPlayerMountSync(hidden bool) *crosspb.ISPlayerMountSync {
	isPlayerMountSync := &crosspb.ISPlayerMountSync{}
	isPlayerMountSync.Hidden = &hidden

	return isPlayerMountSync
}
