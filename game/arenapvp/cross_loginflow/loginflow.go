package cross_loginflow

import (
	"fgame/fgame/game/arenapvp/pbutil"
	crossloginflow "fgame/fgame/game/cross/loginflow"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
)

func init() {
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeArenapvp, crossloginflow.HandlerFunc(arenapvpAttend))
}

//参加pvp
func arenapvpAttend(pl player.Player, crossType crosstypes.CrossType, crossArgs ...string) (err error) {
	siMsg := pbutil.BuildSIArenapvpAttend()
	pl.SendCrossMsg(siMsg)
	return
}
