package relive_handler

import (
	crossloginflow "fgame/fgame/game/cross/loginflow"
	crosstypes "fgame/fgame/game/cross/types"
	lianyulogic "fgame/fgame/game/lianyu/logic"
	"fgame/fgame/game/player"
)

func init() {
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeLianYu, crossloginflow.HandlerFunc(lianYuSendAttend))
}

//无间炼狱发送参加
func lianYuSendAttend(pl player.Player, crossType crosstypes.CrossType, crossArgs ...string) (err error) {
	lianyulogic.LianYuAttendSend(pl)
	return
}
