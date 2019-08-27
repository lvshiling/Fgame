package cross_loginflow

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/chuangshi/pbutil"
	crossloginflow "fgame/fgame/game/cross/loginflow"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	crossloginflow.RegisterCrossLoginFlow(crosstypes.CrossTypeChuangShi, crossloginflow.HandlerFunc(enterChuangShi))
}

//进入创世城池
func enterChuangShi(pl player.Player, crossType crosstypes.CrossType, crossArgs ...string) (err error) {
	if len(crossArgs) <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"len":      len(crossArgs),
			}).Warn("login:玩家加载数据,参数不对")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	cityId, err := strconv.ParseInt(crossArgs[0], 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"arg":      crossArgs[0],
			}).Warn("login:玩家加载数据,参数不对")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	siMsg := pbutil.BuildSIChuangShiEnterCity(cityId)
	pl.SendCrossMsg(siMsg)
	return
}
