package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/chess/chess"
	"fgame/fgame/game/chess/pbutil"
	playerchess "fgame/fgame/game/chess/player"
	chesstypes "fgame/fgame/game/chess/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeChessReset, command.CommandHandlerFunc(handleChessReset))
}

func handleChessReset(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理棋局次数重置")

	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	typStr := c.Args[0]
	typ, _ := strconv.ParseInt(typStr, 10, 64)

	err = resetChess(pl, chesstypes.ChessType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理棋局次数重置,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理棋局次数重置完成")
	return
}

func resetChess(p scene.Player, typ chesstypes.ChessType) (err error) {
	pl := p.(player.Player)
	chessManager := pl.GetPlayerDataManager(types.PlayerChessDataManagerType).(*playerchess.PlayerChessDataManager)
	chessManager.GMResetTimes(typ)

	//发送仙盟个人信息
	logList := chess.GetChessService().GetLogByTime(0)
	chessMap := chessManager.GetChessMap()
	chessInfo := pbutil.BuildSCChessInfoGet(chessMap, logList)
	pl.SendMsg(chessInfo)

	return
}
