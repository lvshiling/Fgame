package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	collectlogic "fgame/fgame/game/collect/logic"
	collecttypes "fgame/fgame/game/collect/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SCENE_COLLECT_TYPE), dispatch.HandlerFunc(handleSceneCollect))
}

//处理采集信息
func handleSceneCollect(s session.Session, msg interface{}) (err error) {
	log.Debug("collect:处理采集消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSceneCollect := msg.(*uipb.CSSceneCollect)
	collectType := collecttypes.CollectType(csSceneCollect.GetCollectType())
	npcId := csSceneCollect.GetNpcId()
	if !collectType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"npcId":       npcId,
				"collectType": collectType,
				"error":       err,
			}).Warn("collect:参数错误")
		return
	}

	err = collectNpc(tpl, collectType, npcId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
				"error":    err,
			}).Error("collect:处理采集消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"npcId":    npcId,
		}).Debug("collect:处理采集消息完成")
	return nil
}

//处理采集信息逻辑
func collectNpc(pl player.Player, collectType collecttypes.CollectType, npcId int64) (err error) {
	return collectlogic.HandleCollectNpc(pl, collectType, npcId)
}
