package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSILianYuAttend() *crosspb.SILianYuAttend {
	siLianYuAttend := &crosspb.SILianYuAttend{}
	return siLianYuAttend
}

func BuildSILianYuCancleLineUp() *crosspb.SILianYuCancleLineUp {
	siLianYuCancleLineUp := &crosspb.SILianYuCancleLineUp{}
	return siLianYuCancleLineUp
}

func BuildSILianYuLineUpSuccess() *crosspb.SILianYuLineUpSuccess {
	siLianYuLineUpSuccess := &crosspb.SILianYuLineUpSuccess{}
	return siLianYuLineUpSuccess
}

func BuildSILianYuFinishLineUpCancle() *crosspb.SILianYuFinishLineUpCancle {
	siLianYuFinishLineUpCancle := &crosspb.SILianYuFinishLineUpCancle{}
	return siLianYuFinishLineUpCancle
}
