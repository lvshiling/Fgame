package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSIShenMoAttend() *crosspb.SIShenMoAttend {
	siShenMoAttend := &crosspb.SIShenMoAttend{}
	return siShenMoAttend
}

func BuildSIShenMoCancleLineUp() *crosspb.SIShenMoCancleLineUp {
	siShenMoCancleLineUp := &crosspb.SIShenMoCancleLineUp{}
	return siShenMoCancleLineUp
}

func BuildSIShenMoLineUpSuccess() *crosspb.SIShenMoLineUpSuccess {
	siShenMoLineUpSuccess := &crosspb.SIShenMoLineUpSuccess{}
	return siShenMoLineUpSuccess
}

func BuildSIShenMoFinishLineUpCancle() *crosspb.SIShenMoFinishLineUpCancle {
	siShenMoFinishLineUpCancle := &crosspb.SIShenMoFinishLineUpCancle{}
	return siShenMoFinishLineUpCancle
}

func BuildSIPlayerGongXunNumChanged(gongXunNum int32) *crosspb.SIPlayerGongXunChanged {
	siPlayerGongXunChanged := &crosspb.SIPlayerGongXunChanged{}
	siPlayerGongXunChanged.Num = &gongXunNum
	return siPlayerGongXunChanged
}

func BuildSIShenMoKillNumChanged() *crosspb.SIShenMoKillNumChanged {
	siShenMoKillNumChanged := &crosspb.SIShenMoKillNumChanged{}
	return siShenMoKillNumChanged
}
