package battle

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	yuxiscene "fgame/fgame/game/yuxi/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterCheckFixPosHandler(scenetypes.SceneTypeYuXi, scene.CheckFixPosSceneHandlerFunc(checkFixPosScene))
}

//玉玺之战 位置修正
func checkFixPosScene(spl scene.Player, s scene.Scene) (flag bool) {
	pl, ok := spl.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 不是玩家")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	sd, ok := s.SceneDelegate().(yuxiscene.YuXiSceneData)
	if !ok {
		return
	}

	owner, _ := sd.GetOwerYuXinInfo()
	if owner == spl {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 持有玉玺，不能修正位置")
		playerlogic.SendSystemMessage(pl, lang.YuXiOwnerCanNotTransfer)
		return
	}

	flag = true
	return
}
