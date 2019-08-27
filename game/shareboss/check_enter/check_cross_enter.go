package check_enter

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/cross/cross"
	crosstypes "fgame/fgame/game/cross/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	cross.RegisterCrossCheckEnterHandler(crosstypes.CrossTypeWorldboss, cross.CheckEnterCrossHandlerFunc(checkEnterCross))
}

//跨服世界boss 进入处理
func checkEnterCross(pl player.Player, crossType crosstypes.CrossType) (flag bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeCrossWorldBoss) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("cangjingge: 进入场景失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}
	huiYuanManager := pl.GetPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	if !huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 进入场景失败，不是至尊会员")
		playerlogic.SendSystemMessage(pl, lang.HuiYuanNotHuiYuan)
		return
	}
	return true
}
