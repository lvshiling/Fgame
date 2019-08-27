package handler

import (
	"fgame/fgame/game/equipbaoku/equipbaoku"
	"fgame/fgame/game/equipbaoku/pbutil"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeEquipBaoKuClearLog, command.CommandHandlerFunc(handleEquipBaoKuClearLog))
}

func handleEquipBaoKuClearLog(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理宝库日志重置")

	err = equipBaoKuClearLog(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理宝库日志重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理宝库日志重置完成")
	return
}

func equipBaoKuClearLog(p scene.Player) (err error) {
	pl := p.(player.Player)
	equipbaoku.GetEquipBaoKuService().GMClearLog()

	manager := pl.GetPlayerDataManager(types.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)

	equipTyp := equipbaokutypes.BaoKuTypeEquip
	materialTyp := equipbaokutypes.BaoKuTypeMaterials
	equipObj := manager.GetEquipBaoKuObj(equipTyp)
	equipLogList := equipbaoku.GetEquipBaoKuService().GetLogByTime(0, equipTyp)

	materialObj := manager.GetEquipBaoKuObj(materialTyp)
	materialLogList := equipbaoku.GetEquipBaoKuService().GetLogByTime(0, materialTyp)

	shopBuyCountMap := manager.GetEquipBaoKuShopBuyAll()

	scEquipBaoKuInfoGet := pbutil.BuildSCEquipBaoKuInfoGet(equipObj, materialObj, equipLogList, materialLogList, shopBuyCountMap)
	pl.SendMsg(scEquipBaoKuInfoGet)

	return
}
