package check_enter

import (
	"fgame/fgame/common/lang"
	coretypes "fgame/fgame/core/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	shengtanscene "fgame/fgame/game/shengtan/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterCheckEnterHandler(scenetypes.SceneTypeAllianceShengTan, scene.CheckEnterSceneHandlerFunc(checkEnterScene))
}

//仙盟圣坛 进入处理
func checkEnterScene(spl scene.Player, s scene.Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	enterPos = pos
	pl, ok := spl.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 不是玩家")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeAllianceAltar) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 进入场景失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}
	sd, ok := s.SceneDelegate().(shengtanscene.ShengTanSceneData)
	if !ok {
		return
	}
	if pl.GetAllianceId() != sd.GetAllianceId() {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"allianceId":    pl.GetAllianceId(),
				"sceneAlliance": sd.GetAllianceId(),
			}).Warn("shenyu: 进入场景失败，没有加入仙盟")
		playerlogic.SendSystemMessage(pl, lang.ShengTanAllianceUserNotSameAlliance)
		return
	}

	flag = true
	return
}
