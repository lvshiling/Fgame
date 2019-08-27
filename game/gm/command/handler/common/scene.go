package common

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeScene, command.CommandHandlerFunc(handleScene))
}

func handleScene(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理跳转场景")
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	sceneIdStr := c.Args[0]
	sceneId, err := strconv.ParseInt(sceneIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"sceneId": sceneIdStr,
				"error":   err,
			}).Warn("gm:处理跳转场景,场景id不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = changeScene(pl, int32(sceneId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理跳转场景,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理跳转场景,完成")
	return
}

func changeScene(pl scene.Player, mapId int32) (err error) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入场景,场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	var s scene.Scene
	//TODO 移除
	switch mapTemplate.GetMapType().MapType() {
	case scenetypes.MapTypeWorld:
		s = scene.GetSceneService().GetWorldSceneByMapId(mapId)
		break
		// case scenetypes.MapTypeWorldBoss:
		// 	s = scene.GetSceneService().GetWorldBossSceneByMapId(mapId)
		break
	case scenetypes.MapTypeTower:
		s = scene.GetSceneService().GetTowerSceneByMapId(mapId)
		break
		// case scenetypes.MapTypeUnrealBoss:
		// 	s = scene.GetSceneService().GetUnrealBossSceneByMapId(mapId)
		// 	break
		// case scenetypes.MapTypeOutlandBoss:
		// 	s = scene.GetSceneService().GetOutlandBossSceneByMapId(mapId)
		break
	default:
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入场景,场景暂时未实现")
		playerlogic.SendSystemMessage(pl, lang.SceneNotWorldScene)
		return
	}

	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入场景,场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	pos := s.MapTemplate().GetBornPos()
	scenelogic.PlayerEnterScene(pl, s, pos)

	return
}
