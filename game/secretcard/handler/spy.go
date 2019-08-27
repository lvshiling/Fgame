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
	secretcardlogic "fgame/fgame/game/secretcard/logic"
	"fgame/fgame/game/secretcard/pbutil"
	playersecretcard "fgame/fgame/game/secretcard/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_SPY_TYPE), dispatch.HandlerFunc(handleSecretSpy))
}

//处理天机牌窥探天机信息
func handleSecretSpy(s session.Session, msg interface{}) (err error) {
	log.Debug("secretcard:处理天机牌窥探天机消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = secretSpy(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("secretcard:处理天机牌窥探天机消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("secretcard:处理天机牌窥探天机消息完成")
	return nil

}

//天机牌窥探天机信息的逻辑
func secretSpy(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	secretCardObj := manager.GetSecretCard()
	cardId := secretCardObj.CardId
	curCardMap := secretCardObj.CardMap
	if cardId != 0 || len(curCardMap) != 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("secretcard:有任务在执行")
		playerlogic.SendSystemMessage(pl, lang.SecretCardHasExecute)
		return
	}
	flag := manager.IfCanSecretSpy()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("secretcard:天机牌次数已用完")
		playerlogic.SendSystemMessage(pl, lang.SecretCardNoEnough)
		return
	}

	cardMap := secretcardlogic.GetSecretCardSpy(pl)
	num := manager.SecretCardSyp(cardMap)
	scQuestSecretSpy := pbutil.BuildSCQuestSecretSpy(num, cardMap)
	pl.SendMsg(scQuestSecretSpy)
	return
}
