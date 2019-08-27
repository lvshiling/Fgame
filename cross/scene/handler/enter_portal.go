package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_ENTER_PORTAL_TYPE), dispatch.HandlerFunc(handleEnterPortal))
}

//处理进入场景
func handleEnterPortal(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理进入传送阵")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	csEnterPortal := msg.(*scenepb.CSEnterPortal)
	portalId := csEnterPortal.GetPortalId()

	err = playerEnterPortal(tpl, portalId)
	if err != nil {
		return
	}
	return nil
}

//玩家进入传送阵
func playerEnterPortal(pl scene.Player, portalId int32) (err error) {

	portalTemplate := scenetemplate.GetSceneTemplateService().GetPortal(portalId)
	if portalTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"portalId": portalId,
			}).Warn("scene:处理进入传送阵,传送阵不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}
	originS := pl.GetScene()
	if originS != nil {
		//同一个场景
		if int32(originS.MapTemplate().TemplateId()) == portalTemplate.MapId {
			scenelogic.FixPosition(pl, portalTemplate.GetPosition())
			return
		}
	}
	return scenelogic.PlayerEnterMapWithPortal(pl, portalTemplate.MapId, portalTemplate.GetPosition())
}
