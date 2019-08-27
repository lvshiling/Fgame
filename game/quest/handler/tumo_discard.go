package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_DISCARD_TYPE), dispatch.HandlerFunc(handleQuestTuMoDiscard))
}

//处理放弃屠魔
func handleQuestTuMoDiscard(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理放弃屠魔")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csQuestTuMoDiscard := msg.(*uipb.CSQuestTuMoDiscard)
	questId := csQuestTuMoDiscard.GetQuestId()

	err = questTuMoDiscard(tpl, questId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
				"error":    err,
			}).Error("quest:处理放弃屠魔,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"questId":  questId,
		}).Debug("quest:处理放弃屠魔,完成")
	return nil
}

//放弃屠魔
func questTuMoDiscard(pl player.Player, questId int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questObj := manager.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	tuMoObj := manager.GetTuMoQuestById(questId)
	if questObj == nil || tuMoObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("quest:无效参数")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//放弃任务
	manager.DiscardTuMoQuest(questId)
	quest := manager.GetQuestByIdAndState(questtypes.QuestStateDiscard, questId)
	scQuestUpdate := pbutil.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)
	scQuestTuMoDiscard := pbutil.BuildSCQuestTuMoDiscard(questId)
	pl.SendMsg(scQuestTuMoDiscard)
	return

}
