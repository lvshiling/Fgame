package relive_handler

import (
	crossloginflow "fgame/fgame/game/cross/loginflow"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/tulong/pbutil"
)

func init() {
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeTuLong, crossloginflow.HandlerFunc(tuLongSendAttend))
}

//跨服屠龙发送参加
func tuLongSendAttend(pl player.Player, crossType crosstypes.CrossType, crossArgs ...string) (err error) {
	siTuLongAttend := pbutil.BuildSITuLongAttend()
	pl.SendCrossMsg(siTuLongAttend)
	return
}
