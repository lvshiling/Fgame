package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/game/scene/scene"
)

func BuildISGodSiegeAttend(godType int32, isLineUp bool, lineUpPos int32) *crosspb.ISGodSiegeAttend {
	isGodSiegeAttend := &crosspb.ISGodSiegeAttend{}
	isGodSiegeAttend.GodType = &godType
	isGodSiegeAttend.IsLineUp = &isLineUp
	isGodSiegeAttend.BeforeNum = &lineUpPos
	return isGodSiegeAttend
}

func BuildISGodSiegeCancleUp(godType int32) *crosspb.ISGodSiegeCancleLineUp {
	isGodSiegeCancleLineUp := &crosspb.ISGodSiegeCancleLineUp{}
	isGodSiegeCancleLineUp.GodType = &godType
	return isGodSiegeCancleLineUp
}

func BuildISGodSiegeLineUpSuccess(godType int32) *crosspb.ISGodSiegeLineUpSuccess {
	isGodSiegeLineUpSuccess := &crosspb.ISGodSiegeLineUpSuccess{}
	isGodSiegeLineUpSuccess.GodType = &godType
	return isGodSiegeLineUpSuccess
}

func BuildISGodSiegeFinshToLineUp(godType int32) *crosspb.ISGodSiegeFinishLineUpCancle {
	isGodSiegeFinishLineUpCancle := &crosspb.ISGodSiegeFinishLineUpCancle{}
	isGodSiegeFinishLineUpCancle.GodType = &godType
	return isGodSiegeFinishLineUpCancle
}

func BuildISDenseWatSync(p scene.Player) *crosspb.ISDenseWatSync {
	isDenseWatSync := &crosspb.ISDenseWatSync{}
	num := p.GetDenseWatNum()
	endTime := p.GetDenseWatEndTime()
	isDenseWatSync.DensWatData = buildDenseWat(num, endTime)
	return isDenseWatSync
}

func buildDenseWat(num int32, endTime int64) *crosspb.DenseWatData {
	denseWatData := &crosspb.DenseWatData{}
	denseWatData.Num = &num
	denseWatData.EndTime = &endTime
	return denseWatData
}
