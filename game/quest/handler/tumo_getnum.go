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
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_TUMONUM_GET_TYPE), dispatch.HandlerFunc(handleQuestTuMoNumGet))
}

//处理获取屠魔次数
func handleQuestTuMoNumGet(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理获取屠魔次数")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = questTuMoNumGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("quest:处理获取屠魔次数,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理获取屠魔次数,完成")
	return nil
}

//获取屠魔次数
func questTuMoNumGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	num, buyNum, leftNum := manager.GetTuMoNum()
	scQuestTuMoNumGet := pbutil.BuildSCQuestTuMoNumGet(num, buyNum, leftNum)
	pl.SendMsg(scQuestTuMoNumGet)
	return

}
