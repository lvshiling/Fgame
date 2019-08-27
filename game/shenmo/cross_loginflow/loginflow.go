package relive_handler

import (
	crossloginflow "fgame/fgame/game/cross/loginflow"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	shenmologic "fgame/fgame/game/shenmo/logic"
)

func init() {
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeShenMoWar, crossloginflow.HandlerFunc(shenMoSendAttend))
}

//神魔战场发送参加
func shenMoSendAttend(pl player.Player, crossType crosstypes.CrossType, crossArgs ...string) (err error) {
	shenmologic.ShenMoAttendSend(pl)
	return
}
