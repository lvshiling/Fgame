package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/secretcard/pbutil"
	playersecretcard "fgame/fgame/game/secretcard/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_CARD_GET_TYPE), dispatch.HandlerFunc(handleSecretCardGet))
}

//处理天机牌界面信息
func handleSecretCardGet(s session.Session, msg interface{}) (err error) {
	log.Debug("secretcard:处理获取天机牌界面消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = secretCardGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("secretcard:处理获取天机牌界面消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("secretcard:处理获取天机牌界面消息完成")
	return nil
}

//获取天机牌界面信息的逻辑
func secretCardGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	secretCard := manager.GetSecretCard()
	scQuestSecretCardGet := pbutil.BuildSCQuestSecretCardGet(secretCard)
	err = pl.SendMsg(scQuestSecretCardGet)
	return
}
