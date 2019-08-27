package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/collect/pbutil"
	"fgame/fgame/cross/processor"
	collectnpc "fgame/fgame/game/collect/npc"
	collecttypes "fgame/fgame/game/collect/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_MIZANG_OPEN_TYPE), dispatch.HandlerFunc(handleSceneMiZangOpen))
}

//处理采集密藏
func handleSceneMiZangOpen(s session.Session, msg interface{}) (err error) {
	log.Debug("collect:处理采集密藏")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)

	csMsg := msg.(*uipb.CSSceneCollectMiZangOpen)
	npcId := csMsg.GetNpcId()
	openType := collecttypes.MiZangOpenType(csMsg.GetType())

	if !openType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
				"openType": openType,
			}).Warn("collect:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = miZangOpen(tpl, openType, npcId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
				"openType": openType,
				"error":    err,
			}).Error("collect:处理采集密藏,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
			"openType": openType,
		}).Info("collect:处理采集密藏完成")
	return nil
}

//处理采集密藏逻辑
func miZangOpen(pl scene.Player, openType collecttypes.MiZangOpenType, npcId int64) (err error) {
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
		playerlogic.SendSystemMessage(pl, lang.CollectMiZangDisappear)
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

	miZangTemplate := n.GetMiZangTemplate()
	miZangOpen := miZangTemplate.GetMiZang(openType)
	if miZangOpen == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理采集消息,没有采集密藏")
		playerlogic.SendSystemMessage(pl, lang.CollectMiZangOpenTypeWrong)
		return
	}

	//打开密藏
	flag := n.MiZangCollectFinish(pl)
	if !flag {
		panic("mizang:打开密藏应该成功")
	}
	biologyId := int32(n.GetBiologyTemplate().TemplateId())
	miZangId := int32(n.GetMiZangTemplate().TemplateId())
	scMsg := pbutil.BuildISCollectMiZang(npcId, biologyId, miZangId, int32(openType))
	pl.SendMsg(scMsg)
	return
}
