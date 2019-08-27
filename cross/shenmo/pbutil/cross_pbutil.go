package pbutil

import crosspb "fgame/fgame/common/codec/pb/cross"

func BuildISShenMoAttend(isLineUp bool, lineUpPos int32) *crosspb.ISShenMoAttend {
	isShenMoAttend := &crosspb.ISShenMoAttend{}
	isShenMoAttend.IsLineUp = &isLineUp
	isShenMoAttend.BeforeNum = &lineUpPos
	return isShenMoAttend
}

func BuildISShenMoCancleUp() *crosspb.ISShenMoCancleLineUp {
	isShenMoCancleLineUp := &crosspb.ISShenMoCancleLineUp{}
	return isShenMoCancleLineUp
}

func BuildISShenMoLineUpSuccess() *crosspb.ISShenMoLineUpSuccess {
	isShenMoLineUpSuccess := &crosspb.ISShenMoLineUpSuccess{}
	return isShenMoLineUpSuccess
}

func BuildISShenMoFinishLineUpCancle() *crosspb.ISShenMoFinishLineUpCancle {
	isShenMoFinishLineUpCancle := &crosspb.ISShenMoFinishLineUpCancle{}
	return isShenMoFinishLineUpCancle
}

func BuildISShenMoKillNumChanged(killNum int32) *crosspb.ISShenMoKillNumChanged {
	isShenMoKillNumChanged := &crosspb.ISShenMoKillNumChanged{}
	isShenMoKillNumChanged.KillNum = &killNum
	return isShenMoKillNumChanged
}

func BuildISPlayerGongXunAdd(addNum int32) *crosspb.ISPlayerGongXunAdd {
	isPlayerGongXunAdd := &crosspb.ISPlayerGongXunAdd{}
	isPlayerGongXunAdd.Num = &addNum
	return isPlayerGongXunAdd
}

func BuildISPlayerGongXunSub(subNum int32) *crosspb.ISPlayerGongXunSub {
	isPlayerGongXunSub := &crosspb.ISPlayerGongXunSub{}
	isPlayerGongXunSub.Num = &subNum
	return isPlayerGongXunSub
}

func BuildISPlayerGongXunChanged() *crosspb.ISPlayerGongXunChanged {
	isPlayerGongXunChanged := &crosspb.ISPlayerGongXunChanged{}
	return isPlayerGongXunChanged
}
