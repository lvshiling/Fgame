package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/game/scene/scene"
)

func BuildISXueChiSync(pl scene.Player) *crosspb.ISXueChiSync {
	isXueChiSync := &crosspb.ISXueChiSync{}
	xueChiData := &crosspb.XueChiData{}
	blood := pl.GetBlood()
	bloodLine := pl.GetBloodLine()
	xueChiData.Blood = &blood
	xueChiData.BloodLine = &bloodLine
	isXueChiSync.XueChiData = xueChiData
	return isXueChiSync
}
