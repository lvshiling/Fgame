package relive_handler

import (
	crossloginflow "fgame/fgame/game/cross/loginflow"
	crosstypes "fgame/fgame/game/cross/types"
	godsiegelogic "fgame/fgame/game/godsiege/logic"
	"fgame/fgame/game/player"
)

func init() {
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeGodSiegeQiLin, crossloginflow.HandlerFunc(godSiegeSendAttend))
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeGodSiegeHuoFeng, crossloginflow.HandlerFunc(godSiegeSendAttend))
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeGodSiegeDuLong, crossloginflow.HandlerFunc(godSiegeSendAttend))
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeDenseWat, crossloginflow.HandlerFunc(godSiegeSendAttend))
}

//神兽攻城发送参加
func godSiegeSendAttend(pl player.Player, crossType crosstypes.CrossType, crossArgs ...string) (err error) {
	godsiegelogic.GodSiegeAttendSend(pl, crossType)
	return
}
