package battle

import (
	"fgame/fgame/common/lang"
	coretypes "fgame/fgame/core/types"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterCheckEnterHandler(scenetypes.SceneTypeChuangShiZhiZhanFuShu, scene.CheckEnterSceneHandlerFunc(checkEnterScene))
}

// 创世之战
func checkEnterScene(spl scene.Player, s scene.Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	enterPos = pos
	pl, ok := spl.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("chuangshi: 不是玩家")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeChuangShiZhiZhan) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("chuangshi: 进入场景失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	if pl.GetCamp() == chuangshitypes.ChuangShiCampTypeNone {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 进入场景失败，没有加入阵营")
		playerlogic.SendSystemMessage(pl, lang.ChuangShiNotCamp)
		return
	}

	flag = true
	return
}
