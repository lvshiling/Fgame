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
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_PICKUP_TYPE), dispatch.HandlerFunc(handleSecretPickup))
}

//处理天机牌接取任务信息
func handleSecretPickup(s session.Session, msg interface{}) (err error) {
	log.Debug("secretcard:处理天机牌接取任务消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	secretPickUp := msg.(*uipb.CSQuestSecretPickUp)
	cardId := secretPickUp.GetCardId()
	err = secretPickup(tpl, cardId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"cardId":   cardId,
				"error":    err,
			}).Error("secretcard:处理天机牌接取任务消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("secretcard:处理天机牌接取任务消息完成")
	return nil

}

//天机牌接取任务信息的逻辑
func secretPickup(pl player.Player, cardId int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	flag := manager.IsValidSecretCard(cardId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"cardId":   cardId,
			}).Warn("secretcard:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	questId, flag := manager.PickUpSecretCard(cardId)
	if !flag {
		panic(fmt.Errorf("secretcard: secretPickup PickUpSecretCard should be ok,cardId:%d", cardId))
	}
	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	quest, flag := questManager.AcceptSecretCardQuest(questId)
	if !flag {
		panic(fmt.Errorf("secretcard: secretPickup AcceptSecretCardQuest should be ok,cardId:%d,questId:%d", cardId, questId))
	}

	scQuestUpdate := pbutilquest.BuildSCQuestUpdate(quest)
	pl.SendMsg(scQuestUpdate)
	scQuestSecretPickUp := pbutil.BuildSCQuestSecretPickUp(cardId)
	pl.SendMsg(scQuestSecretPickUp)
	return
}
