package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arenapvp/arenapvp"
	"fgame/fgame/game/arenapvp/pbutil"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENAPVP_GUESS_INFO_TYPE), dispatch.HandlerFunc(handleArenapvpGuessInfo))
}

//处理竞猜信息
func handleArenapvpGuessInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("arenapvp:处理竞猜信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenapvpGuessInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arenapvp:处理竞猜信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arenapvp:处理竞猜信息,完成")
	return nil
}

func arenapvpGuessInfo(pl player.Player) (err error) {
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	logList := arenapvpManager.GetGuessLogList()
	guessLogObj := arenapvpManager.GetLastGuessLog()
	guessData := arenapvp.GetArenapvpService().GetArenapvpGuessData()

	scMsg := pbutil.BuildSCArenapvpGuessInfo(guessLogObj, guessData, logList)
	pl.SendMsg(scMsg)
	return
}
