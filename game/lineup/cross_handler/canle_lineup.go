package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	lineuplogic "fgame/fgame/game/lineup/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_LINEUP_CANCEL_TYPE), dispatch.HandlerFunc(handleLineupCancleLineUp))
}

//处理取消排队
func handleLineupCancleLineUp(s session.Session, msg interface{}) (err error) {
	log.Debug("lineup:处理取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = lineupCancleLineUp(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("lineup:处理取消排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lineup:处理取消排队,完成")
	return nil
}

//取消排队
func lineupCancleLineUp(pl player.Player) (err error) {
	lineuplogic.CancelCrossLineup(pl)
	return
}
