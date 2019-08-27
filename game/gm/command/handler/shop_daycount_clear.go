package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/shop/pbutil"
	playershop "fgame/fgame/game/shop/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeClearDayCount, command.CommandHandlerFunc(handleShopDayCountClear))
}

//清空商店每日限购次数
func handleShopDayCountClear(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空商店每日限购次数")

	err = clearShopDayCount(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空商店每日限购次数,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空商店每日限购次数完成")
	return
}

//商店每日限购次数置为0
func clearShopDayCount(p scene.Player) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	manager.GmClearDayCount()

	shops := manager.GetShopBuyAll()
	scShopLimit := pbutil.BuildSCShopLimit(shops)
	err = pl.SendMsg(scShopLimit)
	return
}
