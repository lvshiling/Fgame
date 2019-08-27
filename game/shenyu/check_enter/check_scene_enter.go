package check_enter

import (
	"fgame/fgame/common/lang"
	coretypes "fgame/fgame/core/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	playershenyu "fgame/fgame/game/shenyu/player"
	shenyuscene "fgame/fgame/game/shenyu/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterCheckEnterHandler(scenetypes.SceneTypeShenYu, scene.CheckEnterSceneHandlerFunc(shenYuCheckEnterScene))
}

//神域之战 进入处理
func shenYuCheckEnterScene(spl scene.Player, s scene.Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	// 神域之战特殊处理
	enterPos = s.MapTemplate().GetMap().RandomPosition()

	pl, ok := spl.(player.Player)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 不是玩家")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// if enterType == scenetypes.SceneEnterTypeTrac {
	// 	log.WithFields(
	// 		log.Fields{}).Warn("shenyu: 本地图不支持传送")
	// 	playerlogic.SendSystemMessage(pl, lang.CommonMapNotTransfer)
	// 	return
	// }

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenYuZhiZhan) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 进入场景失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	sd, ok := s.SceneDelegate().(shenyuscene.ShenYuSceneData)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shenyu: 进入场景失败，不是神域之战场景数据")
		playerlogic.SendSystemMessage(pl, lang.ShenYuNotShenYu)
		return
	}

	shenYuManager := pl.GetPlayerDataManager(playertypes.PlayerShenYuDataManagerType).(*playershenyu.PlayerShenYuDataManager)

	// 掉线重登
	shenYuTemp := sd.GetShenYuTemplate()
	if pl.GetMapId() == shenYuTemp.MapId {
		flag = true
		return
	}

	if shenYuManager.IsAttendShenYu(sd.GetActivityEndTime()) {
		// 不是从神域进入神域
		if pl.GetScene().MapTemplate().GetMapType() != scenetypes.SceneTypeShenYu {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"endTime":  sd.GetActivityEndTime(),
				}).Warn("shenyu: 进入场景失败，已参与过神域之战")
			playerlogic.SendSystemMessage(pl, lang.ShenYuHadAttend)
			return
		}
	}
	flag = true
	return
}
