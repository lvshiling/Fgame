package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coretypes "fgame/fgame/core/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_ENTER_SCENE_TYPE), dispatch.HandlerFunc(handleEnter))
}

//处理进入场景
func handleEnter(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理进入场景")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	csEnterScene := msg.(*scenepb.CSEnterScene)
	mapId := csEnterScene.GetMapId()
	pos := csEnterScene.GetPos()
	//用户进入场景
	if pos == nil {
		err = playerEnterMap(tpl, mapId)
		if err != nil {
			return err
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Debug("scene:处理进入场景,完成")
	} else {
		aPos := coretypes.Position{
			X: float64(pos.GetPosX()),
			Y: float64(pos.GetPosY()),
			Z: float64(pos.GetPosZ()),
		}
		err = scenelogic.PlayerEnterMapWithPos(tpl, mapId, aPos)
		if err != nil {
			return
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
				"pos":      aPos,
			}).Debug("scene:处理进入场景,完成")
	}
	return nil
}

//玩家进入地图
func playerEnterMap(pl scene.Player, mapId int32) (err error) {
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
	bornPos := mapTemplate.GetBornPos()
	return scenelogic.PlayerEnterMapWithPos(pl, mapId, bornPos)
}
