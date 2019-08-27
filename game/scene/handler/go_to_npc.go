package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GO_TO_NPC_TYPE), dispatch.HandlerFunc(handleGoToNPC))
}

//处理跳转到npc
func handleGoToNPC(s session.Session, msg interface{}) (err error) {
	log.Debug("scene:处理跳转到npc")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csGoToNPC := msg.(*uipb.CSGoToNPC)
	npcId := csGoToNPC.GetNpcId()
	err = goToNPC(tpl, npcId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
				"err":      err,
			}).Error("scene:处理跳转到npc,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Debug("scene:处理跳转到npc,完成")

	return nil
}

//处理跳转到npc
func goToNPC(pl player.Player, npcId int32) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	sceneTemplate := scenetemplate.GetSceneTemplateService().GetQuestNPC(npcId)
	if sceneTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("scene:处理跳转到npc,npc不存在")
		return
	}
	s := scene.GetSceneService().GetWorldSceneByMapId(sceneTemplate.SceneID)
	if s == nil {
		return
	}
	pos := sceneTemplate.GetPos()
	if pl.GetScene() != s {
		scenelogic.PlayerEnterScene(pl, s, pos)
	} else {
		scenelogic.FixPosition(pl, pos)
	}

	scGoToNPC := pbutil.BuildSCGoToNPC(npcId)
	err = pl.SendMsg(scGoToNPC)
	return
}
