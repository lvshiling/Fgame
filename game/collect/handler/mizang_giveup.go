package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	collectnpc "fgame/fgame/game/collect/npc"
	"fgame/fgame/game/collect/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_MIZANG_GIVEUP_TYPE), dispatch.HandlerFunc(handleSceneMiZangGiveup))
}

//处理放弃密藏采集
func handleSceneMiZangGiveup(s session.Session, msg interface{}) (err error) {
	log.Debug("collect:处理放弃密藏采集")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSSceneCollectMiZangGiveup)
	npcId := csMsg.GetNpcId()

	err = miZangGiveup(tpl, npcId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
				"error":    err,
			}).Error("collect:处理放弃密藏采集,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Info("collect:处理放弃密藏采集完成")
	return nil
}

//处理放弃密藏采集逻辑
func miZangGiveup(pl player.Player, npcId int64) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("collect:处理采集消息,场景为空")
		return
	}

	// npc是否存在
	so := s.GetSceneObject(npcId)
	if so == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,生物不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	n, ok := so.(*collectnpc.CollecMiZangNPC)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,不是密藏")
		playerlogic.SendSystemMessage(pl, lang.CollectNotCollectNPC)
		return
	}

	//是否采集完成
	if !n.IfMiZangCanCollect(pl) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,没有采集密藏")
		playerlogic.SendSystemMessage(pl, lang.CollectMiZangNotCollect)
		return
	}

	//清除采集信息
	flag := n.MiZangCollectGiveUp(pl)
	if !flag {
		panic("mizang:放弃密藏应该成功")
	}

	scMsg := pbutil.BuildSCSceneCollectMiZangGiveup(npcId)
	pl.SendMsg(scMsg)
	return
}
