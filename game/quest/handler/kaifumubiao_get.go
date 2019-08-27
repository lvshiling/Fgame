package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_KAIFUMUBIAO_GET_TYPE), dispatch.HandlerFunc(handleQuestKaiFuMuBiaoGet))
}

//处理获取开服目标
func handleQuestKaiFuMuBiaoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理获取开服目标")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = questKaiFuMuBiaoGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("quest:处理获取开服目标,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理获取开服目标,完成")
	return nil
}

//获取开服目标
func questKaiFuMuBiaoGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	kaiFuMuBiaoMap := manager.GetKaiFuMuBiaoMap()
	scQuestKaiFuMuBiaoGet := pbutil.BuildSCQuestKaiFuMuBiaoGet(kaiFuMuBiaoMap)
	pl.SendMsg(scQuestKaiFuMuBiaoGet)
	return
}
