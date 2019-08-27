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
	scene.RegisterCheckEnterHandler(scenetypes.SceneTypeLongGong, scene.CheckEnterSceneHandlerFunc(checkEnterScene))
}

//龙宫探宝 进入处理
func checkEnterScene(spl scene.Player, s scene.Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	enterPos = pos

	pl, ok := spl.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("longgong: 不是玩家")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeLongGongTanBao) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("longgong: 进入场景失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	flag = true
	return
}
