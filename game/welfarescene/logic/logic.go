package logic

import (
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
)

//进入跨服
func EnterCrossWelareScene(pl player.Player) (flag bool, err error) {
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeWelfare)
	flag = true
	return
}
