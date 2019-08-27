package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/chess/chess"
	"fgame/fgame/game/chess/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHESS_LOG_INCR_TYPE), dispatch.HandlerFunc(handleChessLogIncr))
}

//处理苍龙棋局日志请求
func handleChessLogIncr(s session.Session, msg interface{}) (err error) {
	log.Debug("chess:处理获取苍龙棋局日志请求")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csChessLogIncr := msg.(*uipb.CSChessLogIncr)
	logTime := csChessLogIncr.GetLogTime()

	err = chessLogIncr(tpl, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
				"error":    err,
			}).Error("chess:处理获取苍龙棋局日志请求,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"logTime":  logTime,
		}).Debug("chess:处理获取苍龙棋局日志请求完成")
	return nil

}

//获取苍龙棋局界面信息逻辑
func chessLogIncr(pl player.Player, logTime int64) (err error) {
	logList := chess.GetChessService().GetLogByTime(logTime)
	if len(logList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
			}).Warn("chess:处理获取苍龙棋局日志请求,日志增量列表为空")
		return
	}

	scChessLogIncr := pbutil.BuildSCChessLogIncr(logList)
	pl.SendMsg(scChessLogIncr)
	return
}
