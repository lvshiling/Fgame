package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	collectnpc "fgame/fgame/game/collect/npc"
	pbuitl "fgame/fgame/game/collect/pbutil"
	collecttypes "fgame/fgame/game/collect/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_CHOOSE_RESULT_TYPE), dispatch.HandlerFunc(handleSceneCollectChooseResult))
}

//处理采集信息
func handleSceneCollectChooseResult(s session.Session, msg interface{}) (err error) {
	log.Debug("collect choose result:处理采集消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSceneCollectChooseResult := msg.(*uipb.CSSceneCollectChooseResult)
	finishType := collecttypes.CollectChooseFinishType(csSceneCollectChooseResult.GetResult())
	if !finishType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"finishType": finishType,
				"error":      err,
			}).Warn("collect choose result:参数错误")
		return
	}

	err = collectChooseFinishNcp(tpl, finishType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("collect choose result:处理采集消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("collect choose result:处理采集消息完成")
	return nil
}

//处理采集信息逻辑
func collectChooseFinishNcp(pl player.Player, finishType collecttypes.CollectChooseFinishType) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"finishType": finishType,
			}).Warn("collect choose result:处理采集消息,场景为空")
		scMsg := pbuitl.BuildSCSceneCollectChooseResult(0, finishType, false)
		pl.SendMsg(scMsg)
		return
	}

	n, flag := pl.HasCollect()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"finishType": finishType,
			}).Warn("collect choose result:处理采集消息,不在采集中")
		playerlogic.SendSystemMessage(pl, lang.CollectChooseNoCollecting)
		scMsg := pbuitl.BuildSCSceneCollectChooseResult(0, finishType, false)
		pl.SendMsg(scMsg)
		return
	}

	npcId := n.GetId()
	cn, ok := n.(*collectnpc.CollectChooseNPC)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"npcId":      npcId,
				"finishType": finishType,
			}).Warn("collect choose result:处理采集消息,不是特殊采集物")
		playerlogic.SendSystemMessage(pl, lang.CollectChooseNoSpeaial)
		scMsg := pbuitl.BuildSCSceneCollectChooseResult(npcId, finishType, false)
		pl.SendMsg(scMsg)
		return
	}

	if cn.IsDead() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"npcId":      npcId,
				"finishType": finishType,
			}).Warn("collect choose result:该采集物已被采集过,请等待重生")
		playerlogic.SendSystemMessage(pl, lang.CommonCollectIsDead)
		scMsg := pbuitl.BuildSCSceneCollectChooseResult(npcId, finishType, false)
		pl.SendMsg(scMsg)
		return
	}

	//采集成功
	cn.CollectFinish(finishType)

	scMsg := pbuitl.BuildSCSceneCollectChooseResult(npcId, finishType, true)
	pl.SendMsg(scMsg)
	return
}
