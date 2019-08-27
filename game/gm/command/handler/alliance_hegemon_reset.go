package handler

import (
	"fgame/fgame/game/alliance/alliance"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeAllianceHegemonReset, command.CommandHandlerFunc(handleHegemonReset))
}

func handleHegemonReset(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理重置霸主信息")

	err = hegemonReset(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理重置霸主信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理重置霸主信息,完成")
	return
}

func hegemonReset(p scene.Player) (err error) {
	pl := p.(player.Player)
	err = alliance.GetAllianceService().GMAllianceHegemonReset()
	if err != nil {
		return
	}

	//生成新的主城雕像
	alliancelogic.LoadWinnerModel(0)
	alliancelogic.SendAllianceHegemonInfo(pl)
	return
}
