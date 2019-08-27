package pbutil

import crosspb "fgame/fgame/common/codec/pb/cross"

func BuildISLianYuAttend(isLineUp bool, lineUpPos int32) *crosspb.ISLianYuAttend {
	isLianYuAttend := &crosspb.ISLianYuAttend{}
	isLianYuAttend.IsLineUp = &isLineUp
	isLianYuAttend.BeforeNum = &lineUpPos
	return isLianYuAttend
}

func BuildISLianYuCancleUp() *crosspb.ISLianYuCancleLineUp {
	isLianYuCancleLineUp := &crosspb.ISLianYuCancleLineUp{}
	return isLianYuCancleLineUp
}

func BuildISLianYuLineUpSuccess() *crosspb.ISLianYuLineUpSuccess {
	isLianYuLineUpSuccess := &crosspb.ISLianYuLineUpSuccess{}
	return isLianYuLineUpSuccess
}

func BuildISLianYuFinshToLineUp() *crosspb.ISLianYuFinishLineUpCancle {
	isLianYuFinishLineUpCancle := &crosspb.ISLianYuFinishLineUpCancle{}
	return isLianYuFinishLineUpCancle
}
