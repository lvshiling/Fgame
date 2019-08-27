package handler

import (
	"fgame/fgame/game/chess/chess"
	"fgame/fgame/game/chess/pbutil"
	playerchess "fgame/fgame/game/chess/player"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeChessClearLog, command.CommandHandlerFunc(handleChessClearLog))
}

func handleChessClearLog(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理棋局日志重置")

	err = clearChessLog(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理棋局日志重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理棋局日志重置完成")
	return
}

func clearChessLog(p scene.Player) (err error) {
	pl := p.(player.Player)
	chessManager := pl.GetPlayerDataManager(types.PlayerChessDataManagerType).(*playerchess.PlayerChessDataManager)
	chess.GetChessService().GMClearLog()

	logList := chess.GetChessService().GetLogByTime(0)
	chessMap := chessManager.GetChessMap()
	chessInfo := pbutil.BuildSCChessInfoGet(chessMap, logList)
	pl.SendMsg(chessInfo)

	return
}
