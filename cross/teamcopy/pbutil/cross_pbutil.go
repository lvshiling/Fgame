package pbutil

import crosspb "fgame/fgame/common/codec/pb/cross"

func BuildISTeamCopyStartBattle(sucess bool) *crosspb.ISTeamCopyStartBattle {
	isTeamCopyStartBattle := &crosspb.ISTeamCopyStartBattle{}
	isTeamCopyStartBattle.Sucess = &sucess
	return isTeamCopyStartBattle
}

func BuildISTeamCopyBattleResult(purpose int32, sucess bool) *crosspb.ISTeamCopyBattleResult {
	isTeamCopyBattleResult := &crosspb.ISTeamCopyBattleResult{}
	isTeamCopyBattleResult.Purpose = &purpose
	isTeamCopyBattleResult.Sucess = &sucess
	return isTeamCopyBattleResult
}
