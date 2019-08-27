package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	chesslogic "fgame/fgame/game/chess/logic"
	chesstypes "fgame/fgame/game/chess/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHESS_ATTEND_TYPE), dispatch.HandlerFunc(handleChessAttend))

}

//破解苍龙棋局
func handleChessAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("chess:破解苍龙棋局")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csChessAttend := msg.(*uipb.CSChessAttend)
	typ := chesstypes.ChessType(csChessAttend.GetTyp())
	logTime := csChessAttend.GetLogTime()
	autoFlag := csChessAttend.GetAutoFlag()

	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"autoFlag": autoFlag,
			}).Warn("chess:破解棋局错误,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = chessAttend(tpl, typ, logTime, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"error":    err,
			}).Error("chess:处理破解苍龙棋局,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("chess:处理破解苍龙棋局完成")
	return nil

}

//破解苍龙棋局逻辑
func chessAttend(pl player.Player, typ chesstypes.ChessType, logTime int64, autoFlag bool) (err error) {
	return chesslogic.ChessAttend(pl, typ, logTime, autoFlag)
}
