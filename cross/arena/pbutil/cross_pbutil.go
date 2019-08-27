package pbutil

import crosspb "fgame/fgame/common/codec/pb/cross"

func BuildISArenaMatch() *crosspb.ISArenaMatch {
	isArenaMatch := &crosspb.ISArenaMatch{}
	return isArenaMatch
}

func BuildISArenaStopMatch(success bool) *crosspb.ISArenaStopMatch {
	isArenaStopMatch := &crosspb.ISArenaStopMatch{}
	isArenaStopMatch.Success = &success
	return isArenaStopMatch
}

func BuildISArenaMatchResult(result bool) *crosspb.ISArenaMatchResult {
	isArenaMatchResult := &crosspb.ISArenaMatchResult{}
	isArenaMatchResult.Result = &result
	return isArenaMatchResult
}

func BuildISArenaWin(level int32, extra bool) *crosspb.ISArenaWin {
	isArenaWin := &crosspb.ISArenaWin{}
	isArenaWin.Level = &level
	isArenaWin.Extra = &extra
	return isArenaWin
}

var (
	isArenaCollectExpTree = &crosspb.ISArenaCollectExpTree{}
)

func BuildISArenaCollectExpTree() *crosspb.ISArenaCollectExpTree {

	return isArenaCollectExpTree
}

func BuildISArenaCollectBox(boxId int32) *crosspb.ISArenaCollectBox {
	isArenaCollectBox := &crosspb.ISArenaCollectBox{}
	isArenaCollectBox.BoxId = &boxId
	return isArenaCollectBox
}

func BuildISArenaGiveUp() *crosspb.ISArenaGiveUp {
	isMsg := &crosspb.ISArenaGiveUp{}
	return isMsg
}

func BuildISArenaResetReliveTimes() *crosspb.ISArenaResetReliveTimes {
	isMsg := &crosspb.ISArenaResetReliveTimes{}
	return isMsg
}
