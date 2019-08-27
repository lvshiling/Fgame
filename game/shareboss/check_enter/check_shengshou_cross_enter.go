package check_enter

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/cross/cross"
	crosstypes "fgame/fgame/game/cross/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	cross.RegisterCrossCheckEnterHandler(crosstypes.CrossTypeShengShou, cross.CheckEnterCrossHandlerFunc(checkEnterShengShouCross))
}

//跨服世界boss 进入处理
func checkEnterShengShouCross(pl player.Player, crossType crosstypes.CrossType) (flag bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShengShouBoss) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("cangjingge: 进入场景失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	return true
}
