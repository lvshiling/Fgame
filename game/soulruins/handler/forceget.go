package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	soulruinslogic "fgame/fgame/game/soulruins/logic"
	soulruinstypes "fgame/fgame/game/soulruins/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOULRUINS_FORCEGET_TYPE), dispatch.HandlerFunc(handleSoulRuinsForceGet))
}

//处理帝魂降临传功获得处理消息
func handleSoulRuinsForceGet(s session.Session, msg interface{}) (err error) {
	log.Debug("soulruins:处理获取帝魂降临传功获得处理消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = soulRuinsForceGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("soulruins:处理获取帝魂降临传功获得处理消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soulruins:处理获取帝魂降临传功获得处理消息完成")
	return nil

}

//获取帝魂降临传功获得处理的逻辑
func soulRuinsForceGet(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeFuBenSoulRuins {
		return
	}
	//判断状态
	sd := s.SceneDelegate()
	sceneData, ok := sd.(*soulruinslogic.SoulRuinsSceneData)
	if !ok {
		return
	}

	eventType := sceneData.GetEventType()
	if eventType != soulruinstypes.SoulRuinsEventTypeSoul {
		return
	}
	acceptFlag := sceneData.GetAccept()
	if !acceptFlag {
		return
	}
	stageType := sceneData.GetStageType()
	if stageType != soulruinstypes.SoulRuinsStageTypeSecond {
		return
	}

	s.Finish(true)
	soulruinslogic.GiveSoulRuinsSoulForceReward(pl, sceneData)
	return
}
