package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/chess/chess"
	"fgame/fgame/game/chess/pbutil"
	playerchess "fgame/fgame/game/chess/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHESS_INFO_GET_TYPE), dispatch.HandlerFunc(handleChessGet))

}

//处理苍龙棋局信息
func handleChessGet(s session.Session, msg interface{}) (err error) {
	log.Debug("chess:处理获取苍龙棋局消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = chessGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("chess:处理获取苍龙棋局消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("chess:处理获取苍龙棋局消息完成")
	return nil

}

//获取苍龙棋局界面信息逻辑
func chessGet(pl player.Player) (err error) {
	chessManager := pl.GetPlayerDataManager(playertypes.PlayerChessDataManagerType).(*playerchess.PlayerChessDataManager)
	playerChessMap := chessManager.GetChessMap()
	logList := chess.GetChessService().GetLogByTime(0)
	if len(logList) > 10 {
		logList = logList[len(logList)-10:]
	}
	scChessInfoGet := pbutil.BuildSCChessInfoGet(playerChessMap, logList)
	pl.SendMsg(scChessInfoGet)
	return
}
