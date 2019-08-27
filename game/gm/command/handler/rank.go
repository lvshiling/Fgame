package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/rank/rank"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeRank, command.CommandHandlerFunc(handleRankUpdate))
}

//排行榜更新
func handleRankUpdate(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理排行榜更新")

	err = rankUpdate(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理排行榜更新,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理排行榜更新完成")
	return
}

//排行榜更新
func rankUpdate(pl scene.Player) (err error) {
	rank.GMRankUpdate()
	return
}
