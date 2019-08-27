package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/xianfu/pbutil"
	playerxianfu "fgame/fgame/game/xianfu/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommondTypeXianfuResetTimes, command.CommandHandlerFunc(handleXianfuReset))
}

func handleXianfuReset(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理重置秘境仙府次数")
	pl := p.(player.Player)
	err = resetXianfu(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理重置秘境仙府次数,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理重置秘境仙府次数,完成")
	return
}

func resetXianfu(pl player.Player) (err error) {
	xfManager := pl.GetPlayerDataManager(playertypes.PlayerXianfuDtatManagerType).(*playerxianfu.PlayerXinafuDataManager)
	xfObjArr := xfManager.GetPlayerXianfuInfoList()
	for _, xfObj := range xfObjArr {
		xfObj.SetUseTimes(0)
		xfObj.SetModified()
	}

	xianfuArr := xfManager.GetPlayerXianfuInfoList()
	scXianfuGet := pbutil.BuildSCXianfuGet(xianfuArr)
	pl.SendMsg(scXianfuGet)

	return
}
