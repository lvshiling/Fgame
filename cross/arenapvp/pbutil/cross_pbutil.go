package pbutil

import crosspb "fgame/fgame/common/codec/pb/cross"

func BuildISArenapvpAttend(isLineup bool, sceneId int64) *crosspb.ISArenapvpAttend {
	isMsg := &crosspb.ISArenapvpAttend{}
	isMsg.IsLineUp = &isLineup
	isMsg.SceneId = &sceneId
	return isMsg
}

func BuildISArenapvpRelive() *crosspb.ISArenapvpRelive {
	isMsg := &crosspb.ISArenapvpRelive{}
	return isMsg
}

func BuildISArenapvpResetReliveTimes() *crosspb.ISArenapvpResetReliveTimes {
	isMsg := &crosspb.ISArenapvpResetReliveTimes{}
	return isMsg
}

func BuildISArenapvpAttendSuccess() *crosspb.ISArenapvpAttendSuccess {
	isMsg := &crosspb.ISArenapvpAttendSuccess{}
	return isMsg
}

func BuildISArenapvpResultBattle(win bool, pvpType int32) *crosspb.ISArenapvpResultBattle {
	isMsg := &crosspb.ISArenapvpResultBattle{}
	isMsg.Win = &win
	isMsg.PvpType = &pvpType
	return isMsg
}

func BuildISArenapvpResultElection(win bool, ranking, pvpType int32) *crosspb.ISArenapvpResultElection {
	isMsg := &crosspb.ISArenapvpResultElection{}
	isMsg.Ranking = &ranking
	isMsg.PvpType = &pvpType
	isMsg.Win = &win
	return isMsg
}

func BuildISAreanapvpElectionLuckyRew() *crosspb.ISAreanapvpElectionLuckyRew {
	isMsg := &crosspb.ISAreanapvpElectionLuckyRew{}
	return isMsg
}
