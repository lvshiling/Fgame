package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/major/pbutil"
	playermajor "fgame/fgame/game/major/player"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMajorClear, command.CommandHandlerFunc(handleMajorClear))
}

//清空双修次数
func handleMajorClear(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理清空双修次数")

	err = clearMajorClear(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理清空双修次数,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理清空双修次数完成")
	return
}

//双修次数置为0
func clearMajorClear(p scene.Player) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerMajorDataManagerType).(*playermajor.PlayerMajorDataManager)
	manager.GmClearMajorNum()

	for initType := majortypes.MinType; initType <= majortypes.MaxType; initType++ {
		scMajorNum := pbutil.BuildSCMajorNum(int32(initType), 0)
		pl.SendMsg(scMajorNum)
	}

	return
}
