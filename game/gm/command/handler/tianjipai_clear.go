package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/secretcard/pbutil"
	playersecretcard "fgame/fgame/game/secretcard/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTianJiPaiClear, command.CommandHandlerFunc(handleTianJiPaiClear))
}

//天机牌次数置为0
func handleTianJiPaiClear(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空天机牌次数")
	pl := p.(player.Player)
	err = clearTianJiPaiNum(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空天机牌次数,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空天机牌次数完成")
	return
}

//天机牌次数置为0
func clearTianJiPaiNum(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	manager.GMClearNum()

	cardObj := manager.GetSecretCard()
	scQuestSecretCardGet := pbutil.BuildSCQuestSecretCardGet(cardObj)
	pl.SendMsg(scQuestSecretCardGet)
	return
}
