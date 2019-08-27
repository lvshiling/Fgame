package handler

import (
	"fgame/fgame/game/equipbaoku/pbutil"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeEquipBaoKuShopClearDayCount, command.CommandHandlerFunc(handleEquipBaoKuShopDayCountClear))
}

//清空宝库商店每日限购次数
func handleEquipBaoKuShopDayCountClear(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空宝库商店每日限购次数")

	err = clearEquipBaoKuShopDayCount(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空宝库商店每日限购次数,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空宝库商店每日限购次数完成")
	return
}

//宝库商店每日限购次数置为0
func clearEquipBaoKuShopDayCount(p scene.Player) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	manager.GmClearDayCount()

	shops := manager.GetEquipBaoKuShopBuyAll()
	scShopLimit := pbutil.BuildSCEquipBaoKuShopLimit(shops)
	err = pl.SendMsg(scShopLimit)
	return
}
