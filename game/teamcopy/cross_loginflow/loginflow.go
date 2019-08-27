package relive_handler

import (
	crosslogic "fgame/fgame/game/cross/logic"
	crossloginflow "fgame/fgame/game/cross/loginflow"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/team/team"
	teamcopylogic "fgame/fgame/game/teamcopy/logic"
)

func init() {
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeTeamCopy, crossloginflow.HandlerFunc(teamCopySendStartBattle))
}

//组队副本发送开始战斗
func teamCopySendStartBattle(pl player.Player, crossType crosstypes.CrossType, crossArgs ...string) (err error) {
	teamObj := team.GetTeamService().GetTeamByPlayerId(pl.GetId())
	//而且还要是队长
	if teamObj != nil && teamObj.IsCopyBattle() {
		//发送开始战斗
		teamcopylogic.TeamCopyStartBattleSend(pl)
	} else {
		crosslogic.PlayerExitCross(pl)
	}
	return
}
