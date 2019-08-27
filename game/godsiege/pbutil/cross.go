package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSIGodSiegeAttend(godType int32) *crosspb.SIGodSiegeAttend {
	siGodSiegeAttend := &crosspb.SIGodSiegeAttend{}
	siGodSiegeAttend.GodType = &godType
	return siGodSiegeAttend
}

func BuildSIGodSiegeCancleLineUp(godType int32) *crosspb.SIGodSiegeCancleLineUp {
	siGodSiegeCancleLineUp := &crosspb.SIGodSiegeCancleLineUp{}
	siGodSiegeCancleLineUp.GodType = &godType
	return siGodSiegeCancleLineUp
}

func BuildSIGodSiegeLineUpSuccess(godType int32) *crosspb.SIGodSiegeLineUpSuccess {
	siGodSiegeLineUpSuccess := &crosspb.SIGodSiegeLineUpSuccess{}
	siGodSiegeLineUpSuccess.GodType = &godType
	return siGodSiegeLineUpSuccess
}

func BuildSIGodSiegeFinishLineUpCancle(godType int32) *crosspb.SIGodSiegeFinishLineUpCancle {
	siGodSiegeFinishLineUpCancle := &crosspb.SIGodSiegeFinishLineUpCancle{}
	siGodSiegeFinishLineUpCancle.GodType = &godType
	return siGodSiegeFinishLineUpCancle
}
