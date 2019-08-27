package relive_handler

import (
	arenalogic "fgame/fgame/game/arena/logic"
	crosslogic "fgame/fgame/game/cross/logic"
	crossloginflow "fgame/fgame/game/cross/loginflow"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/team/team"
)

func init() {
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeArena, crossloginflow.HandlerFunc(arenaSendMatch))
}

//3v3发送匹配
func arenaSendMatch(pl player.Player, crossType crosstypes.CrossType, crossArgs ...string) (err error) {
	teamObj := team.GetTeamService().GetTeamByPlayerId(pl.GetId())
	//而且还要是队长
	if teamObj != nil && teamObj.IsMatch() {
		//发送匹配
		arenalogic.ArenaMatchSend(pl)
	} else {
		crosslogic.PlayerExitCross(pl)
	}
	return nil
}
