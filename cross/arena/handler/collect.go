package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	arenascene "fgame/fgame/cross/arena/scene"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	playerlogic "fgame/fgame/game/player/logic"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_FOUR_GOD_SCENE_COLLECTING_TYPE), dispatch.HandlerFunc(handleCollect))
}

//处理四圣兽采集
func handleCollect(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理四圣兽采集")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	csArenaFourGodSceneCollecting := msg.(*uipb.CSArenaFourGodSceneCollecting)
	collectId := csArenaFourGodSceneCollecting.GetCollectId()
	err = arenaFourGodSceneCollect(tpl, collectId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理四圣兽采集,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理选择四圣兽")
	return nil

}

//处理四圣兽经验采集
func arenaFourGodSceneCollect(pl *player.Player, collectId int64) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arena:不在四圣兽场景")
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArenaShengShou {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arena:不在四圣兽场景")
		return
	}
	//判断是否可以采集
	//采集
	sd := s.SceneDelegate().(arenascene.FourGodSceneData)
	if sd == nil {
		return
	}

	flag := sd.IfCollectDistance(collectId, pl.GetPos())
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"collectId": collectId,
			}).Warn("arena:不在采集范围内")
		playerlogic.SendSystemMessage(pl, lang.CommonCollectNoDistance)
		return
	}

	flag = sd.IfCollect(collectId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"collectId": collectId,
			}).Warn("arena:别人正在采集")
		playerlogic.SendSystemMessage(pl, lang.ArenaOtherIsCollecting)
		return
	}
	flag = sd.Collect(pl, collectId)
	if !flag {
		panic(fmt.Errorf("arena:采集应该成功"))
	}
	return
}
