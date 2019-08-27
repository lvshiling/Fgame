package check_enter

import (
	"fgame/fgame/common/lang"
	coretypes "fgame/fgame/core/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterCheckEnterHandler(scenetypes.SceneTypeCangJingGe, scene.CheckEnterSceneHandlerFunc(checkEnterScene))
}

//藏经阁 进入处理
func checkEnterScene(spl scene.Player, s scene.Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	enterPos = pos
	pl, ok := spl.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("cangjingge: 不是玩家")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeCangJingGe) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("cangjingge: 进入场景失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	// huiYuanManager := pl.GetPlayerDataManager(types.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	// isHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	// tempHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypeInterim)
	// if !isHuiYuan && !tempHuiYuan {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":  pl.GetId(),
	// 			"biologyId": biologyId,
	// 		}).Warn("cangjingge:藏经阁boss挑战请求，不是至尊会员")
	// 	playerlogic.SendSystemMessage(pl, lang.HuiYuanNotHuiYuan)
	// 	return
	// }

	flag = true
	return
}
