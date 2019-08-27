package logic

import (
	"fgame/fgame/cross/relive/pbutil"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/scene/scene"
)

//一般的正常复活
func Relive(pl scene.Player) {
	maxLevel := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeReliveNoNeedItemsBeforeLevel)
	if pl.GetLevel() < maxLevel {
		pl.Reborn(pl.GetPosition())
		return
	}

	pl.RefreshReliveTime()

	culTime := pl.GetCulReliveTime()
	culTime += 1
	isPlayerRelive := pbutil.BuildISPlayerRelive(culTime)
	pl.SendMsg(isPlayerRelive)
}
