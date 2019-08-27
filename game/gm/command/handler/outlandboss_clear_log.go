package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/outlandboss/outlandboss"
	"fgame/fgame/game/outlandboss/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeOutlandBossClearDropRecords, command.CommandHandlerFunc(handleOutlandBossClearDropRecords))
}

func handleOutlandBossClearDropRecords(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理外域boss日志重置")

	err = outlandBossClearDropRecords(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理外域boss日志重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理外域boss日志重置完成")
	return
}

func outlandBossClearDropRecords(p scene.Player) (err error) {
	pl := p.(player.Player)
	outlandboss.GetOutlandBossService().GMClearDropRecords()

	logList := outlandboss.GetOutlandBossService().GetDropRecordsList()
	scMsg := pbutil.BuildSCOutlandBossDropRecordsGet(logList)
	pl.SendMsg(scMsg)
	return
}
