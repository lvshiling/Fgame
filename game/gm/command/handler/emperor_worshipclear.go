package handler

import (
	"fgame/fgame/game/emperor/emperor"
	"fgame/fgame/game/emperor/pbutil"
	playeremperor "fgame/fgame/game/emperor/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeEmperorWorshipClear, command.CommandHandlerFunc(handleEmperorWorshipClear))
}

//清空膜拜次数
func handleEmperorWorshipClear(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空膜拜次数")
	pl := p.(player.Player)
	err = emperorWorshipClear(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空膜拜次数,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空膜拜次数完成")
	return
}

func emperorWorshipClear(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerEmperorDataManagerType).(*playeremperor.PlayerEmperorDataManager)
	manager.GMClearWorshipNum()

	emperorObj := emperor.GetEmperorService().GetEmperorInfo()
	worshipNum := manager.GetWorshipNum()
	scEmperorGet := pbuitl.BuildSCEmperorGet(emperorObj, worshipNum)
	pl.SendMsg(scEmperorGet)
	return
}
