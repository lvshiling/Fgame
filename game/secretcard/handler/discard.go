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
	pbutilquest "fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/secretcard/pbutil"
	playersecretcard "fgame/fgame/game/secretcard/player"
	"fgame/fgame/game/secretcard/secretcard"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_DISCARD_TYPE), dispatch.HandlerFunc(handleSecretDiscard))
}

//处理放弃天机任务
func handleSecretDiscard(s session.Session, msg interface{}) (err error) {
	log.Debug("secretcard:处理放弃天机任务")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = secretDiscard(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("secretcard:处理放弃天机任务,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("secretcard:处理放弃天机任务,完成")
	return nil
}

//放弃天机任务
func secretDiscard(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	cardObj := manager.GetSecretCard()
	//当前没有天机牌任务(含翻牌未接取的)
	if cardObj.CardId == 0 && len(cardObj.CardMap) == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("secretcard:任务不存在")
		playerlogic.SendSystemMessage(pl, lang.SecretCardNoQuest)
		return
	}
	if cardObj.CardId == 0 && len(cardObj.CardMap) != 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("secretcard:请先接取任务")
		playerlogic.SendSystemMessage(pl, lang.SecretImmediateFinish)
		return
	}
	questId, _, flag := secretcard.GetSecretCardService().GetQuestIdByCardId(cardObj.CardId)
	if !flag {
		panic(fmt.Errorf("secretcard: secretDiscard get questId should be ok"))
	}
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	//放弃任务
	quest := questManager.DiscardQuest(questId)
	manager.DiscardQuest()
	scQuestUpdate := pbutilquest.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)
	scQuestSecretCardGet := pbutil.BuildSCQuestSecretCardGet(cardObj)
	pl.SendMsg(scQuestSecretCardGet)
	return

}
